package main

import "fmt"

type age int
type person struct {
	name string
	age  uint
}

func main() {
	var p person
	p1 := struct {
		name string
		age  uint
	}{name: "Jane", age: 19}
	p = p1
	var a1 int
	var a2 age
	a1 = int(a2)
	//a1 = a2 // won't work
	fmt.Printf("person: %+v\nage: %d\n\n", person{name: "Steve", age: 24}, age(1))
	fmt.Printf("p: %+v\na1:%d", p, a1)
}
