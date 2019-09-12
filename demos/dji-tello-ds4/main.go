package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func logText(s string) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(s); err != nil {
		fmt.Println(err)
	}
}

func main() {
	drone := tello.NewDriver("8888")

	err := drone.On(tello.FlightDataEvent, func(d interface{}) {
		logText(fmt.Sprintf("%+v\n\n", d))
	})
	if err != nil {
		log.Fatal(err)
	}

	work := func() {
		err := drone.TakeOff()
		if err != nil {
			fmt.Println("Tello could not take off")
			log.Fatal(err)
		}
		time.Sleep(3 * time.Second)
		_ = drone.FrontFlip()
		time.Sleep(3 * time.Second)
		_ = drone.FrontFlip()
		time.Sleep(3 * time.Second)
		_ = drone.FrontFlip()

		// time.Sleep(3 * time.Second)
		// drone.Right(100)

		gobot.After(10*time.Second, func() {
			err := drone.Land()
			if err != nil {
				fmt.Println("Tello could not land")
				log.Fatal(err)
			}
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	err := robot.Start()
	if err != nil {
		log.Fatal("could not start robot", err)
	}
}
