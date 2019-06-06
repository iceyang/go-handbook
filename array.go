package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	No   int
}

func arrayExample() {
	var intArr [10]int
	intArr[4] = 10
	fmt.Println("intArr:", intArr)

	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers:", numbers)

	var studentArr [4]Student
	fmt.Println("studentArr:", studentArr)

	var twoDimention [5][5]int
	fmt.Println("twoDimention:", twoDimention)

	names := [5]string{0: "Andy", 2: "Tim", 1: "Jason"}
	fmt.Println("names:", names)
}

func typeOf() {
	fmt.Println("type of [...]int{1,2,3,4,5}: ", reflect.TypeOf([...]int{1, 2, 3, 4, 5}))
	fmt.Println("type of []int{1,2,3,4,5}: ", reflect.TypeOf([]int{1, 2, 3, 4, 5}))
}
