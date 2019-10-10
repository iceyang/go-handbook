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

func handleErrorWithType(err error) {
	switch err.(type) {
	case *ErrorOne:
		fmt.Println("err is ErrorOne")
	case *ErrorTwo:
		fmt.Println("err is ErrorTwo")
	}
}

func doSomething() error {
	num := rand.Intn(10)
	if num > 5 {
		return &ErrorOne{}
	} else if num < 5 {
		return &ErrorTwo{}
	}
	return nil
}

func demo2() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		if err := doSomething(); err != nil {
			handleErrorWithType(err)
		} else {
			fmt.Println("There'is no error")
		}
	}
}
