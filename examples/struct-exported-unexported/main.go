package main

import (
	"fmt"
	"github.com/steevehook/mdc19/examples/struct-exported-unexported/person"
)

func main() {
	p1 := struct {
		name string
		age  uint
	}{name: "Steve", age: 24}
	p2 := person.NewPerson("John", 19, "developer")
	p3 := person.Person{Hobby: "human"}
	fmt.Printf("p1:\nname: %s\nage: %d\n\n", p1.name, p1.age)
	fmt.Printf("p2:\nname: %s\nage: %d\nHobby: %s\n\n", p2.Name(), p2.Age(), p2.Hobby)
	fmt.Printf("p3:\nname: %s\nage: %d\nHobby: %s\n", p3.Name(), p3.Age(), p3.Hobby)
}
