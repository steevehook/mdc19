package main

import (
	"fmt"
	"time"
)

func main() {
	n := 10
	go func() {
		n = 15
	}()
	go func() {
		n = 12
	}()
	go func() {
		fmt.Println(n)
	}()
	time.Sleep(1 * time.Second)
}
