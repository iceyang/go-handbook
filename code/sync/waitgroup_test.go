package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup1(t *testing.T) {
	var wg sync.WaitGroup
	n := 10
	counter := 0
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			counter++
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}

/**
 * 会出现并发问题
 */
func TestWaitGroup2(t *testing.T) {
	var wg sync.WaitGroup
	n := 10
	counter := 0
	for i := 0; i < n; i++ {
		go func(i int) {
			wg.Add(1)
			counter++
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}
