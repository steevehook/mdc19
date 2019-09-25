package main

import "fmt"

func listen(messages chan message) {
	for {
		select {
		case msg := <-messages:
			fmt.Printf("%+v", msg)
		default:
			return
		}
	}
}
