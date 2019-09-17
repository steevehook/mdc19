package main

import "fmt"

func main() {
	arrInt, arrString := [3]int{1, 2, 3}, [2]string{"Hello", "World"}
	//arrIntErr := [2]int{10, 20, 30} // index out of bounds: 2
	arrVar := [...]string{
		"Some", "very", "long", "string", "array", "Heck", "knows", "the", "end", "of", "it",
	}
	arrVarCap := cap(arrVar)
	fmt.Printf("%d\n%s\n%s\n", arrInt, arrString, arrVar)
	fmt.Printf("The type of arrVar: %T", arrVar) // [11]string
	fmt.Printf("The capacity of arrVar: %d", arrVarCap)
}
