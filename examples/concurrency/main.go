package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("routine 1")
	}()
	go func() {
		fmt.Println("routine 2")
	}()
	go func() {
		fmt.Println("routine 3")
	}()
	time.Sleep(1 * time.Second)
}
