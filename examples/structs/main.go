package main

import "fmt"

type person struct {
	Name string
	Age  uint
}

func main() {
	people := []struct {
		name string
		age  uint
	}{
		{name: "John", age: 21},
		{"Jane", 19}, // don't use this
	}
	fmt.Printf("%+v\n", person{Name: "Steve", Age: 24})
	fmt.Printf("%+v", people)
}
