package main

import "fmt"

const (
	hello       = "你好"
	exclamation = "!"
)

var world string

func init() {
	world = "世界"
}

func main() {
	fmt.Println(hello + world + exclamation)
}
