package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  uint   `json:"age,omitempty"`
}

func main() {
	people := map[string]person{
		"fa1b4f6f-beab-5e55-acdf-6da3a03aabec": {
			Name: "Steve", Age: 24,
		},
		"fa1b4f6f-beab-5e55-acdf-6da3a03aabed": {
			Name: "John",
		},
	}
	// handle errors, PLEASE :D
	bs, _ := json.Marshal(people)
	// {
	// 	"fa1b4f6f-beab-5e55-acdf-6da3a03aabec":{"name":"Steve","age":24},
	// 	"fa1b4f6f-beab-5e55-acdf-6da3a03aabed":{"name":"John"}
	// }
	fmt.Println(string(bs))
}
