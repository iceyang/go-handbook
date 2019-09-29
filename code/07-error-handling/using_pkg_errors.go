package main

import (
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("test")
	err = errors.WithStack(err, "new message")
	// errors.Cause

	panic(err)
}
