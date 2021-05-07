package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errorOne = errors.New("This is Error One")
var errorTwo = errors.New("This is Error Two")

func handleErrorWithType2(err error) {
	fmt.Print("handleErrorWithType2: ")
	switch err {
	case errorOne:
		fmt.Println("err is ErrorOne")
	case errorTwo:
		fmt.Println("err is ErrorTwo")
	}
}

func doSomething2() error {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	if num > 5 {
		return errorOne
	} else if num < 5 {
		return errorTwo
	}
	return nil
}

func handleErrorDemo2() {

	for i := 0; i < 10; i++ {
		if err := doSomething2(); err != nil {
			handleErrorWithType2(err)
		} else {
			fmt.Println("There'is no error")
		}
	}
}
