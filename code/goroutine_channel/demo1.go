package main

import "fmt"

func producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
	}
	close(queue)
}

func consumer(queue <-chan int, finish chan int) {
	for v := range queue {
		fmt.Printf("Get value: %d \n", v)
	}
	finish <- 1
}

func main() {
	queue := make(chan int)
	finish := make(chan int)
	go producer(queue)
	go consumer(queue, finish)
	<-finish
}
