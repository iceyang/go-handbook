package main

import "fmt"

type ErrorOne struct{}
type ErrorTwo struct{}

func (ErrorOne) Error() string {
	return "This is Error One"
}

func (ErrorTwo) Error() string {
	return "This is Error Two"
}

func checkErrorType(err error) {
	switch err.(type) {
	case *ErrorOne:
		fmt.Println("ErrorOne")
	case *ErrorTwo:
		fmt.Println("ErrorTwo")
	}
}

func demo2() {
	var err error

	err = &ErrorOne{}
	checkErrorType(err)

	err = &ErrorTwo{}
	checkErrorType(err)
}
