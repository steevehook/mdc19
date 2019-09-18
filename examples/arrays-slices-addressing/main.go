package main

import "fmt"

func main() {
	s := make([]int, 10, 12)
	fmt.Println("s:", s) // [0 0 0 0 0 0 0 0 0 0]
	s[0], s[5] = 10, 15
	fmt.Printf("s[0]:%d\ns[5]:%d\ns[9]:%d\n", s[0], s[5], s[9])
	fmt.Println("cap(s):", cap(s)) // 12
	s = append(s, 1, 2, 3)
	fmt.Println("cap(s):", cap(s)) // 24
	//fmt.Println(s[11]) // index out of range
}
