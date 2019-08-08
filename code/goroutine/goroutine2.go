package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go fmt.Print(i, " ")
	}
	time.Sleep(time.Second)
}
