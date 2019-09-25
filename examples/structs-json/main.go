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
	// always handle errors :D
	bs1, _ := json.Marshal(person{Name: "Steve", Age: 24, sex: true, Hobby: "developer"})
	bs2, _ := json.Marshal(person{Name: "Jane", sex: false, Hobby: "human"})
	fmt.Println(string(bs1)) // {"name":"Steve","age":24}
	fmt.Println(string(bs2)) // {"name":"Jane"}
}
