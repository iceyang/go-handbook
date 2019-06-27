package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// declare a slice variable and allocate memory address
	// length == capacity == 10
	_ = make([]int, 10)

	// initialize slice
	_ = []int{1, 2, 3, 4, 5}

	// create a new slice with specified length and capacity
	// length == 10, capacity == 20
	slice := make([]string, 10, 20)

	_ = append(slice, "str")
	_ = append(slice, []string{"1", "2", "3"}...)

	fmt.Println(make([]int, 0))
	fmt.Println(make([]int, 10))
	fmt.Println(make([]int, 0, 20))
	fmt.Println(make([]int, 10, 20))
}
