package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"gobot.io/x/gobot/platforms/dji/tello"
	"math"
)

const (
	L2 = 4
	R2 = 5

	Yaw    = 0
	UpDown = 1

	LeftRight       = 2
	ForwardBackward = 3
)

var axisMap = map[uint8]func(event *sdl.JoyAxisEvent, drone *tello.Driver){
	L2: func(event *sdl.JoyAxisEvent, drone *tello.Driver) {
		fmt.Println("L2 pressed with value:", event.Value)
	},
	R2: func(event *sdl.JoyAxisEvent, drone *tello.Driver) {
		fmt.Println("R2 pressed with value:", event.Value)
	},
	UpDown: func(event *sdl.JoyAxisEvent, drone *tello.Driver) {
		// positive - up
		// negative - down
		convertedUnit := int16To100(event.Value)
		if convertedUnit > 0 {
			fmt.Println("up", convertedUnit)
			drone.Up(int(convertedUnit))
		} else {
			absVal := math.Abs(float64(convertedUnit))
			fmt.Println("down", absVal)
			drone.Down(int(absVal))
		}
	},
	Yaw: func(event *sdl.JoyAxisEvent, drone *tello.Driver) {
		// positive - yaw left
		// negative - yaw right
		convertedUnit := float64(int16To100(event.Value))
		if convertedUnit > 0 {
			drone.CounterClockwise(int(convertedUnit))
		} else {
			absVal := math.Abs(float64(convertedUnit))
			drone.Clockwise(int(absVal))
		}
	},
	LeftRight: func(event *sdl.JoyAxisEvent, drone *tello.Driver) {
		// positive - left
		// negative - right
		convertedUnit := int16To100(event.Value)
		if convertedUnit > 0 {
			fmt.Println("left", convertedUnit)
			drone.Left(int(convertedUnit))
		} else {
			absVal := math.Abs(float64(convertedUnit))
			fmt.Println("right", absVal)
			drone.Right(int(absVal))
		}
	},
	ForwardBackward: func(event *sdl.JoyAxisEvent, drone *tello.Driver) {
		// positive - forward
		// negative - backward
		convertedUnit := int16To100(event.Value)
		if convertedUnit > 0 {
			fmt.Println("forward", convertedUnit)
			drone.Forward(int(convertedUnit))
		} else {
			absVal := math.Abs(float64(convertedUnit))
			fmt.Println("backward", absVal)
			drone.Backward(int(absVal))
		}
	},
}

func AxisSwitch(event *sdl.JoyAxisEvent, drone *tello.Driver) {
	axisType, ok := axisMap[event.Axis]
	if !ok {
		fmt.Println("no such axis entry in axis map:", event.Axis)
		return
	}
	axisType(event, drone)
}
