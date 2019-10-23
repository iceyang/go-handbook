package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) GetName() string {
	return p.name
}

type Cat struct {
}

func (cat *Cat) Eat() {
	fmt.Println("Cat is eating")
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

type Dog struct {
}

func (cat *Dog) Eat() {
	fmt.Println("Dog is eating")
}

func (cat *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}
