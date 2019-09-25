package main

import "time"

type message struct {
	id   string
	text string
}

func main() {
	messages := make(chan message)
	go listen(messages)
	messages <- message{id: "1", text: "first message"}
	messages <- message{id: "2", text: "second message"}
	time.Sleep(1 * time.Second)
}
