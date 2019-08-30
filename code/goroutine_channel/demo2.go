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

func consumer(queues [3]chan int, finish chan int) {
	queue0, queue1, queue2 := queues[0], queues[1], queues[2]
	for {
		time.Sleep(time.Millisecond * 200)
		select {
		case v, ok := <-queue0:
			if !ok {
				queue0 = make(chan int)
				break
			}
			fmt.Printf("From %s: Get value: %d\n", "queue[0]", v)
		case v, ok := <-queue1:
			if !ok {
				queue1 = make(chan int)
				break
			}
			fmt.Printf("From %s: Get value: %d\n", "queue[1]", v)
		case v, ok := <-queue2:
			if !ok {
				queue2 = make(chan int)
				break
			}
			fmt.Printf("From %s: Get value: %d\n", "queue[2]", v)
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
