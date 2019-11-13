package main

import (
	"fmt"
	"testing"
)

func checkAnimalType(animal Animal) {
	switch animal.(type) {
	case *Cat:
		fmt.Println("The animal is Cat")
	case *Dog:
		fmt.Println("The animal is Dog")
	}
}

func TestInterface1(t *testing.T) {
	var animal Animal
	animal = &Cat{}
	animal.Eat()

	animal = &Dog{}
	animal.Sleep()
}

func TestInterface2(t *testing.T) {
	var animal Animal
	animal = &Cat{}
	checkAnimalType(animal)

	animal = &Dog{}
	checkAnimalType(animal)
}
