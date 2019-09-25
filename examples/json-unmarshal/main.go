package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  uint   `json:"age,omitempty"`
	//sex   bool `json:"sex"` // not gonna work
	sex   bool
	Hobby string `json:"-"`
}

func main() {
	j1 := `{"name": "Mary", "age": 21, "sex": true, "hobby": "human"}`
	j2 := `{"name": "Steve", "age": 24, "anything_else": "not_gonna_unmarshal"}`
	var p1 person
	// handle errors, you son of the bitch
	_ = json.Unmarshal([]byte(j1), &p1)
	var p2 person
	_ = json.Unmarshal([]byte(j2), &p2)
	fmt.Printf("p1:\n%+v\n\np2:\n%+v", p1, p2)
}
