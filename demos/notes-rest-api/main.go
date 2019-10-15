package main

import (
	"log"
	"os"
	"syscall"

	"github.com/steevehook/mdc19/dummy-notes-rest-api/app"
)

func main() {
	application := app.New()
	go func() {
		err := application.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
	app.ListenForSignals([]os.Signal{syscall.SIGINT, syscall.SIGTERM}, application)
}
