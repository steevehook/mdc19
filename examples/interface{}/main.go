package main

import "fmt"

func main() {
	anything := map[string]interface{}{
		"animals": []string{"lion", "bear"},
		"grades": []int{10, 8, 9},
		"people": []struct {
			name string
		}{
			{name: "Steve"},
			{name: "John"},
		},
	}
	fmt.Printf("Anything: %+v", anything)
}
