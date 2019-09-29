package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second)
}
