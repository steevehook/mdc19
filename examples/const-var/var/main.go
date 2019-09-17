package main

import "fmt"

var s1 string = "s1"
var s2 = "s2"
var s3 string

// group
var (
	g1 = 1
	g2 = 2
)

func main() {
	s3 = "s3"
	s4 := "s4"
	fmt.Printf("%s\n%s\n%s\n%s\n%d\n%d", s1, s2, s3, s4, g1, g2)
}
