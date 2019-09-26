package main

import "fmt"

type age uint

func (a *age) BirthDay() *age {
	*a += 1
	return a
}
func (a age) IsDead() bool {
	return a > 150
}

func main() {
	a := age(148)
	a.BirthDay().BirthDay()
	fmt.Printf("age: %d\nis dead: %t\n\n", a, a.IsDead())
	a.BirthDay()
	fmt.Printf("age: %d\nis dead: %t", a, a.IsDead())
}
