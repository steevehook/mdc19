package main

import "fmt"

func main() {
	names := []string{"John", "Mike", "Jane"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for index, name := range names {
		fmt.Println(index, name)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}
