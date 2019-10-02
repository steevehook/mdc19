package main

import "errors"

func identity(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("the number is less than zero")
	}
	return n, nil
}







