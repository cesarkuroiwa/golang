package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func main() {
	var p Person
	p.SetName("João")
	p.SetAge(24)

	p1 := Person{ "João", 30 }
	p2 := Person{ age: 30, name: "João" }
	p3 := Person{ name: "Joao" }
	fmt.Println(p.Name(), p.Age())
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}