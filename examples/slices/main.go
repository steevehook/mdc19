package main

import "fmt"

var sComplex []complex128

func main() {
	sInt, sString := []int{1,2,3}, []string{"Hello", "World"}
	s1, s2, s3 := make([]int, 100), make([]int, 0, 100), make([]int, 10, 100)
	sComplex = append(sComplex, 2 + 3i, 4+2i, 5, 6)
	newIntSlice := append([]int{10, 20}, sInt...)
	fmt.Printf("%d\n%s\n%v\n%d\n", sInt, sString, sComplex, newIntSlice)
	fmt.Printf("s1|len:%d<=>cap:%d\ns2|len:%d<=>cap:%d\ns3|len:%d<=>cap:%d\n",
		len(s1), cap(s1), len(s2), cap(s2), len(s3), cap(s3),
	)
}
