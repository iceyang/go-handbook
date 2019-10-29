package main

type Person struct {
	name string
	age  int
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) SetName2(name string) {
	p.name = name
}

type Adult struct {
	job    string
	person Person
}

type AnonymousField struct {
	a int
	b int
	int
}
