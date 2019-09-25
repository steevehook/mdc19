package main

import "fmt"

type person struct {
	name string
	age  uint
}
type developer struct {
	person
	//age uint // this will override the previous field
	language string
}

func main() {
	d := developer{
		language: "Go",
		person:   person{name: "Steve", age: 24},
	}
	fmt.Printf("developer info\n---\nname: %s\nage: %d\nlanguage: %s", d.name, d.person.age, d.language)
}
