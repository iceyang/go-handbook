package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ErrorOne struct{}
type ErrorTwo struct{}

func (ErrorOne) Error() string {
	return "This is Error One"
}

func (ErrorTwo) Error() string {
	return "This is Error Two"
}

func handleErrorWithType1(err error) {
	fmt.Print("handleErrorWithType1: ")
	switch err.(type) {
	case *ErrorOne:
		fmt.Println("err is ErrorOne")
	case *ErrorTwo:
		fmt.Println("err is ErrorTwo")
	}
}

func doSomething1() error {
	num := rand.Intn(10)
	if num > 5 {
		return &ErrorOne{}
	} else if num < 5 {
		return &ErrorTwo{}
	}
	return nil
}

func handleErrorDemo1() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		if err := doSomething1(); err != nil {
			handleErrorWithType1(err)
		} else {
			fmt.Println("There'is no error")
		}
	}
}
