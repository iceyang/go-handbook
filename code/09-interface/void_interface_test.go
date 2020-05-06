package main

import (
	"fmt"
	"testing"
)

func show(entry interface{}) {
	fmt.Println(entry)
}

func TestVoidInterface(t *testing.T) {
	show(1)
	show("Hello World")
	show(Person{"Justin"})
}

func TestVoidInterface2(t *testing.T) {
	var a interface{} = 100
	var b int = a.(int)
	fmt.Println(b)
	_, ok := a.(string)
	if !ok {
		fmt.Println("a is not a string")
	}
}
