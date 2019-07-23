package main

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	var chan0 chan int
	t.Log(chan0)

	chan1 := make(chan int, 5)

	// send 1 & 2 to channel chan1
	chan1 <- 1
	chan1 <- 2

	elem1 := <-chan1
	t.Log(elem1)
	close(chan1)

	chan2 := make(chan string, 3)
	// the element type must be string
	// chan2 <- 1
	chan2 <- "hello"
	chan2 <- "go channel"
	t.Log(<-chan2, <-chan2)
	close(chan2)
}

func TestNoBufferChannel(t *testing.T) {
	chan1 := make(chan int)
	go func() {
		chan1 <- 3
	}()
	// it will be blocked until chan1 has element
	elem1 := <-chan1
	t.Log(elem1)
	close(chan1)
}

func TestPanic(t *testing.T) {
	chan1 := make(chan int, 3)

	chan1 <- 1
	chan1 <- 2
	chan1 <- 3
	elem := <-chan1
	t.Log(elem)

	close(chan1)
	// close a closed channel will cause panic>
	// close(chan1)

	// send element to a closed channel will cause panic.
	// chan1 <- 4

	// it doesn't matter that we received the element from a closed channel.
	elem2, ok := <-chan1
	t.Log(elem2, ok)
}

func TestOneWayChannel(t *testing.T) {
	chan1 := make(chan<- int, 5)
	chan1 <- 1

	// chan1 is a send-only channel. Elements can be send to it, but can't be out.
	// it can not be compiled.
	// elem1 := <-chan1
	close(chan1)

	// chan2 is a receive-only channel.
	// it can not be compiled.
	// chan2 := make(<-chan int, 5)
	// chan2 <- 1

	chan3 := make(chan int, 5)
	chan3 <- 1

	receivedOnlyChannel(chan3)
}

func receivedOnlyChannel(ch <-chan int) {
	fmt.Println(<-ch)
}
