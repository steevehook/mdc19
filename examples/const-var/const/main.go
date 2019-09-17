package main

import "fmt"

const c1 = 10
const s string = "Hello"

// group
const (
	g1      = 10
	g2 int8 = 32
)

func main() {
	fmt.Println(c1, s, g1, g2)
}
