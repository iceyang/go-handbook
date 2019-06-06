package main

import (
	"fmt"
)

func sliceExample() {
	var intSlice []int
	intSlice = make([]int, 10)
	fmt.Printf("intSlice: %+v, length: %d\n", intSlice, len(intSlice))

	intSlice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("intSlice2: %+v, length: %d\n", intSlice2, len(intSlice2))

	fmt.Println(append(intSlice2, 6))
}
