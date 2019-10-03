package main

import "fmt"

func clean(i int) {
	fmt.Printf("cleanup for operation: %d\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		defer clean(i)
	}
}
