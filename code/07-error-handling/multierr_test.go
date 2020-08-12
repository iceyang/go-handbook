package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/multierr"
)

func TestMultiErr(t *testing.T) {
	assert.NoError(t, multierr.Combine(nil, nil))

	assert.Error(t, multierr.Combine(nil, errors.New("Error occurs"), nil))
}
