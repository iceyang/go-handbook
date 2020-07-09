package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
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
	fmt.Println("counter:", counter)
}
