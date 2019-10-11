package main

import "log"

func handle(err error) {
	log.Fatal(err)
}

func main() {
	handle(CustomErr{Message: "some error"})
}
