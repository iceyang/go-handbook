package main

import (
	"github.com/pkg/errors"
)

func demo3() {
	err := errors.New("test")
	// err = errors.WithStack(err, "new message")
	// errors.Cause

	panic(err)
}
