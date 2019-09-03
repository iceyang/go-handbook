package main

import (
	"fmt"
	"time"
)

func producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
	}
	close(queue)
}

func dealWithQueue(no int, ok bool, value int, queue *chan int) {
	if !ok {
		*queue = make(chan int)
		return
	}
	fmt.Printf("From Queue[%d]: Get value: %d\n", no, value)
}

func consumer(queues [3]chan int, finish chan int) {
	queue0, queue1, queue2 := queues[0], queues[1], queues[2]
	for {
		time.Sleep(time.Millisecond * 200)
		select {
		case v, ok := <-queue0:
			dealWithQueue(0, ok, v, &queue0)
		case v, ok := <-queue1:
			dealWithQueue(1, ok, v, &queue1)
		case v, ok := <-queue2:
			dealWithQueue(2, ok, v, &queue2)
		default:
			finish <- 1
		}
	}
}

func main() {
	queues := [3]chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}
	for _, queue := range queues {
		go producer(queue)
	}

	finish := make(chan int)
	go consumer(queues, finish)
	<-finish
}
