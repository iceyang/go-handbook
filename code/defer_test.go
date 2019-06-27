package main

import (
	"testing"
)

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Do something finally.")
	}()
	t.Log("here we go.")
}
