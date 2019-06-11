package main

import (
	"testing"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error()
	}
}

func TestAdd2(t *testing.T) {
	if add2(1, 2) != 3 {
		t.Error()
	}
}
