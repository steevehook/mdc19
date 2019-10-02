package main

import "fmt"

type person struct {
	name string
}

func (p person) Run() {
	fmt.Println(p.name, "is running")
}

type runner interface {
	Run()
}
