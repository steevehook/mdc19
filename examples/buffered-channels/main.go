package main

import (
	"fmt"
	"time"
)

func main() {
	buffered := make(chan int, 2)
	buffered <- 1 // will not block - buffer = 1
	buffered <- 2 // will not block - buffer = 2
	//buffered <- 3 // WILL BLOCK - buffer > 2
	go func() {
		for i := range buffered {
			fmt.Println(i)
		}
	}()
	buffered <- 3
	time.Sleep(1 * time.Second)
}
