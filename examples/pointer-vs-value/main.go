package main

import "fmt"

type person struct {
	name string
}

func (p person) changeNameValue(name string) {
	p.name = name
}
func (p *person) changeNamePointer(name string) {
	p.name = name
}

func main() {
	p1, p2 := person{name: "Steve"}, &person{name: "George"}
	p1.changeNameValue("John")
	fmt.Println(p1.name)
	p1.changeNamePointer("John")
	fmt.Println(p1.name)
	p2.changeNameValue("Unknown")
	fmt.Println(p2.name)
}
