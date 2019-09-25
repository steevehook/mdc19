package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	grades := []int{10, 20, 30, 40}
	people := []struct {
		Name string `json:"name"`
	}{{"Steve"}, {"John"}, {"Jane"}, {"Mary"}}
	// handle errors you bustard
	gradesBS, _ := json.Marshal(grades)
	peopleBS, _ := json.Marshal(people)
	//[10,20,30,40]
	//[{"name":"Steve"},{"name":"John"},{"name":"Jane"},{"name":"Mary"}]
	fmt.Printf("grades:\n%s\n\npeople:\n%s", string(gradesBS), string(peopleBS))
}
