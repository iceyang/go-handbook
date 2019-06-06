package main

func sliceExample() {
	// declare a slice variable and allocate memory address
	// length == capacity == 10
	_ = make([]int, 10)

	// initialize slice
	_ = []int{1, 2, 3, 4, 5}

	// create a new slice with specified length and capacity
	// length == 10, capacity == 20
	_ = make([]string, 10, 20)
}
