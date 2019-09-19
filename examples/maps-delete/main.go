package main

import "fmt"

func main() {
	dictionary := map[string]string{
		"human": "being that has 1 life",
		"programmer": "being that has no life",
	}
	fmt.Printf("%+v\n", dictionary)
	delete(dictionary, "programmer")
	delete(dictionary, "non-existent") // no error
	fmt.Printf("%+v\n", dictionary)
}
