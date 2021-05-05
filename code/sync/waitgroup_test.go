package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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
			fmt.Println(i)
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
	var counter int32
	n := 1000
	for i := 0; i < n; i++ {
		go func(i int) {
			wg.Add(1)
			atomic.AddInt32(&counter, 1)
			// counter += 1 // Dangerous!!
			wg.Done()
		}(i)
	}
	wg.Wait()
	if int(counter) != n {
		t.Errorf("Not all goroutine is finished. counter=%d，expected %d", counter, n)
	}
}
