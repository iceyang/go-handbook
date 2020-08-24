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
