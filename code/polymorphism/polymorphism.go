package polymorphism

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Cat struct {
	Name string
}

func (c Cat) Eat() {
	fmt.Printf("Cat %s is eating.\n", c.Name)
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping.\n", c.Name)
}

type Dog struct {
	Name string
}

func (d Dog) Eat() {
	fmt.Printf("Dog %s is eating.\n", d.Name)
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleep.\n", d.Name)
}

func AnimalEat(animal Animal) {
	animal.Eat()
}
