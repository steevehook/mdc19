package main

import (
	"fmt"
	"log"
)

func main() {
	nums := []int{10, 15, -2, 4}
	for _, n := range nums {
		i, err := identity(n)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i)
	}
}
