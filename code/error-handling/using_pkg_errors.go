package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "This is My Error"
}

func doSomething4() error {
	// return errors.Wrap(&MyError{}, "doSomething4 has an error")
	return errors.WithStack(&MyError{})
}

func pkgErrorsDemo() {
	err := doSomething4()
	fmt.Println(err)
	switch err := errors.Cause(err).(type) {
	case *MyError:
		fmt.Println(err)
	default:
		fmt.Println("unknown error")
	}
	// fmt.Println(err)
	// err = errors.WithStack(err, "new message")
	// errors.Cause
}
