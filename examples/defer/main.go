package main

import "fmt"

type operation struct{}

func (op operation) SideEffects() {
	fmt.Println("doing some side effects")
}
func (op operation) CleanUp() {
	fmt.Println("cleaning up after the side effects")
}

func main() {
	o := operation{}
	o.SideEffects()
	defer o.CleanUp()
	fmt.Println("other operations")
	panic("some serious error")
}
