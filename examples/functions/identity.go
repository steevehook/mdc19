package main

import "errors"

func identity(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("number is lower than 0")
	}
	return num, nil
}
