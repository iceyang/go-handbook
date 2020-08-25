package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func doSomething3() error {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	if num > 5 {
		return &ErrorOne{}
	} else if num < 5 {
		return &ErrorTwo{}
	}
	return errors.New("Unknown Error")
}

func handleErrorWithType3(err error) {
	fmt.Print("handleErrorWithType3: ")
	errStr := err.Error()
	if strings.Contains(errStr, "Error One") {
		fmt.Println("err is ErrorOne")
		return
	}
	if strings.Contains(errStr, "Error Two") {
		fmt.Println("err is ErrorTwo")
		return
	}
	fmt.Println("unknown err")
}

func handleErrorDemo3() {
	for i := 0; i < 10; i++ {
		if err := doSomething3(); err != nil {
			handleErrorWithType3(err)
		} else {
			fmt.Println("There'is no error")
		}
	}
}
