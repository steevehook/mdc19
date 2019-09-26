package main

import "fmt"

func main() {
	var a, b *int
	iPtr := intPtr(10)
	a = iPtr
	b = iPtr
	*a = 10
	fmt.Printf("a: %d\nb: %d\niPtr: %d", *a, *b, *iPtr)
}
