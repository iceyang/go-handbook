package main

import (
	"fmt"
)

type Student struct {
	Name string
	No   int
}

func arrayExample() {
	// declare a array variable
	var intArr [10]int
	intArr[4] = 10

	// initialize array
	_ = [5]int{1, 2, 3, 4, 5}

	// struct array
	var _ [4]Student

	// two dimentional array
	var _ [5][5]int

	// using index to initialize
	_ = [5]string{0: "Andy", 2: "Tim", 1: "Jason"}
}

func typeOf() {
	fmt.Printf("type of [...]int{1,2,3,4,5}: %T\n", [...]int{1, 2, 3, 4, 5})
	fmt.Printf("type of []int{1,2,3,4,5}: %T\n", []int{1, 2, 3, 4, 5})
}
