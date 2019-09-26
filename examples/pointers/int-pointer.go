package main

func intPtr(i int) *int {
	var n int
	n = i
	return &n
}
