package main

import (
	"fmt"
	"testing"
)

func TestAssert1(t *testing.T) {
	var animal Animal
	animal = &Cat{}
	cat, ok := animal.(*Cat)
	fmt.Println(cat, ",", ok)
	dog, ok := animal.(*Dog)
	fmt.Println(dog, ",", ok)
}

func TestAssert2(t *testing.T) {
	var animal interface{}
	animal = &Cat{}
	cat, ok := animal.(*Cat)
	fmt.Println(cat, ",", ok)
	dog, ok := animal.(*Dog)
	fmt.Println(dog, ",", ok)
}

func TestAssert3(t *testing.T) {
	var x interface{}
	x = "Hello World"
	v, ok := x.(string)
	fmt.Println(v, ",", ok)
	v2, ok := x.(int)
	fmt.Println(v2, ",", ok)
}
