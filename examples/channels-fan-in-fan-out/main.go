package main

import (
	"fmt"
	"time"
)

func generate(numbers ...int) chan int {
	out := make(chan int)
	go func() {
		for _, num := range numbers {
			out <- num
		}
		close(out)
	}()
	return out
}

func main() {
	numbers := generate(3, 6, 9)
	c1 := fanOut(numbers)
	c2 := fanOut(numbers)
	result := fanIn(c1, c2)
	for i := range result {
		fmt.Println(i)
	}
	time.Sleep(1 * time.Second)
}
