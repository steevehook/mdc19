package main

import (
	"fmt"
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var joysticks [16]*sdl.Joystick

func run() int {
	var event sdl.Event
	var running bool

	err := sdl.Init(sdl.INIT_JOYSTICK)
	if err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	sdl.JoystickEventState(sdl.ENABLE)

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.JoyAxisEvent:
				AxisSwitch(t)
			case *sdl.JoyButtonEvent:
				ButtonSwitch(t)
			case *sdl.JoyBallEvent:
				fmt.Printf("[%d ms] Ball:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Ball, t.XRel, t.YRel)
			case *sdl.JoyHatEvent:
				fmt.Printf("[%d ms] Hat:%d\tvalue:%d\n", t.Timestamp, t.Hat, t.Value)
			case *sdl.JoyDeviceAddedEvent:
				joysticks[int(t.Which)] = sdl.JoystickOpen(int(t.Which))
				if joysticks[int(t.Which)] != nil {
					fmt.Printf("Joystick %d connected\n", t.Which)
				}
			case *sdl.JoyDeviceRemovedEvent:
				if joystick := joysticks[int(t.Which)]; joystick != nil {
					joystick.Close()
				}
				fmt.Printf("Joystick %d disconnected\n", int(t.Which))
			default:
				fmt.Printf("Unknown event\n")
			}
		}

		sdl.Delay(16)
	}

	return 0
}

func main() {
	os.Exit(run())
}
