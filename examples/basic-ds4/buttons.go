package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	// Types
	Cross    = 0
	Circle   = 1
	Square   = 2
	Triangle = 3
	Share    = 4
	Home     = 5
	Options  = 6
	L1       = 9
	R1       = 10
	Up       = 11
	Down     = 12
	Left     = 13
	Right    = 14
	Pad      = 15

	// States
	Pressed  = 0
	Released = 1
)

var buttonsMap = map[uint8]map[uint8]func(event *sdl.JoyButtonEvent){
	Cross: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("cross released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("cross pressed on joystick:", event.Which)
		},
	},
	Circle: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("circle released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("circle pressed on joystick:", event.Which)
		},
	},
	Square: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("square released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("square pressed on joystick:", event.Which)
		},
	},
	Triangle: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("triangle released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("triangle pressed on joystick:", event.Which)
		},
	},
	Share: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("share released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("share pressed on joystick:", event.Which)
		},
	},
	Home: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("home PS released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("home PS pressed on joystick:", event.Which)
		},
	},
	Options: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("options released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("options pressed on joystick:", event.Which)
		},
	},
	L1: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("L1 released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("L1 pressed on joystick:", event.Which)
		},
	},
	R1: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("R1 released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("R1 pressed on joystick:", event.Which)
		},
	},
	Up: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Up released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Up pressed on joystick:", event.Which)
		},
	},
	Down: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Down released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Down pressed on joystick:", event.Which)
		},
	},
	Left: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Left released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Left pressed on joystick:", event.Which)
		},
	},
	Right: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Right released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Right pressed on joystick:", event.Which)
		},
	},
	Pad: {
		Pressed: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Pad released on joystick:", event.Which)
		},
		Released: func(event *sdl.JoyButtonEvent) {
			fmt.Println("Pad pressed on joystick:", event.Which)
		},
	},
}

func ButtonSwitch(event *sdl.JoyButtonEvent) {
	buttonType, ok := buttonsMap[event.Button]
	if !ok {
		fmt.Println("no such button entry in buttons map:", event.Button)
		return
	}
	buttonType[event.State](event)
}
