package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	cache := cache{}
	listener, err := net.Listen(
		"tcp",
		":8080",
	)
	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		// handle errors, you bustard
		conn, _ := listener.Accept()
		fmt.Fprintf(
			conn,
			"Welcome to CACHE 1.0\n-> ",
		)
		go listen(conn, cache)
	}
}
