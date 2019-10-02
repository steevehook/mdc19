package main

import "fmt"

func marathon(runners ...runner) {
	for _, r := range runners {
		r.Run()
	}
}

func main() {
	steve := person{name: "Steve"}
	john := struct {
		name string
	}{
		name: "John",
	}
	marathon(steve)
	//marathon(steve, john) // incompatible types
	fmt.Println(john.name, "is running")
}
