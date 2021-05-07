package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(time.Second)
}
