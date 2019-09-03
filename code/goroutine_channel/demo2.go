package main

import (
	"fmt"
	"time"
)

func producer(queue chan<- int) {
	for i := 0; i < 20; i++ {
		queue <- i
	}
	close(queue)
}

func consumer(queue <-chan int, finish chan int) {
	for v := range queue {
		fmt.Printf("Get value: %d \n", v)
		time.Sleep(time.Millisecond * 200)
	}
	finish <- 1
}

func main() {
	queue := make(chan int)
	finish := make(chan int)
	consumerCount := 3
	go producer(queue)
	for i := 0; i < consumerCount; i++ {
		go consumer(queue, finish)
	}
	for i := 0; i < consumerCount; i++ {
		<-finish
	}
}
