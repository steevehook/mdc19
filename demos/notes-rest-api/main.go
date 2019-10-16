package main

import (
	"log"
	"os"
	"syscall"

	"github.com/steevehook/mdc19/notes-rest-api/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err := application.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
	app.ListenForSignals([]os.Signal{syscall.SIGINT, syscall.SIGTERM}, application)
}
