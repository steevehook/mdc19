package main

import (
	"fmt"
	"log"
)

func main() {
	nums := []int{10, -15}
	for _, n := range nums {
		val, err := identity(n)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("identity:", val)
	}
}
