package main

import "fmt"

var robot map[string]func() // nil map

func init() {
	robot = map[string]func(){
		"start": func() {
			fmt.Println("robot started")
		},
	}
}

func main() {
	robot["start"]()
	// panic: runtime error: invalid memory address or nil pointer dereference
	robot["stop"]()
}
