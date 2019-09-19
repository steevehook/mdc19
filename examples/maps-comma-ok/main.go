package main

import "fmt"

func main() {
	dictionary := map[string]string{"some_key": "some_value"}
	val, ok := dictionary["another_key"]
	fmt.Printf("value: %v <=> ok: %v\n", val, ok)
	val, ok = dictionary["some_key"]
	fmt.Printf("value: %v <=> ok: %v\n", val, ok)
}
