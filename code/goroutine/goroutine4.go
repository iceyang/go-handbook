package main

import "time"

func waitForTenSeconds() {
	time.Sleep(10 * time.Second)
}

func main() {
	for i := 1; i < 100000; i++ {
		go waitForTenSeconds()
	}
	time.Sleep(60 * time.Second)
}
