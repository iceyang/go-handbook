package main

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
}

func (person *Person) String() string {
	return "Person: " + person.name
}

func TestFmt(t *testing.T) {
	person := &Person{"Justin"}
	fmt.Println(person)
}
