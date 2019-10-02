package main

import "fmt"

func main() {
	grades := map[string]int{
		"Math":    8,
		"English": 10,
		"Physics": 9,
	}
	// the order may always be different (because it's hashmap)
	for subject, grade := range grades {
		fmt.Printf("%s ==> %d\n", subject, grade)
	}
}
