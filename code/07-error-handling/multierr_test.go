package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/multierr"
)

func TestMultiErr(t *testing.T) {
	assert.NoError(t, multierr.Combine(nil, nil))

	assert.Error(t, multierr.Combine(nil, errors.New("Error occurs"), nil))
}

func foo() error {
	return multierr.Combine(
		doSomething1(),
		doSomething2(),
		doSomething3(),
	)
}

func TestMultiErr2(t *testing.T) {
	err := foo()
	errors := multierr.Errors(err)
	if len(errors) > 0 {
		fmt.Println("The following errors occurred:")
		for _, e := range errors {
			fmt.Println(e)
		}
	}
}

type Handler func() error

func Execute(handlers ...Handler) error {
	for _, handler := range handlers {
		if err := handler(); err != nil {
			return err
		}
	}
	return nil
}

func TestMultiErr3(t *testing.T) {
	err := Execute(
		func() error {
			fmt.Println("Here we are1")
			return nil
		},
		func() error {
			fmt.Println("Here we are2")
			return errors.New("error occurs")
		},
		func() error {
			fmt.Println("Here we are3")
			return nil
		},
	)
	fmt.Println(err)
}
