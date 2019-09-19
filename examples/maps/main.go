package main

import "fmt"

var dictionary map[string]string // nil map (either init it later or avoid this)

func main() {
	//dictionary["some_key"] = "some_value" // panic: assignment to entry in nil map
	wiktionary := make(map[string]string)
	wiktionary["some_key"] = "some_value"
	grades := map[string]int{
		"Math":    10,
		"English": 10,
	}
	fmt.Printf("dictionary: %+v\nwiktionary: %+v\ngrades: %+v\n", dictionary, wiktionary, grades)
}
