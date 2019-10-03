package main

import "fmt"

func yikes() {
	panic("can't do anything")
}

func main() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	yikes()
}
