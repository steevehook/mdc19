package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	L2 = 4
	R2 = 5

	LeftY  = 0
	LeftX  = 1

	RightY = 2
	RightX = 3
)

var axisMap = map[uint8]func(*sdl.JoyAxisEvent){
	L2: func(event *sdl.JoyAxisEvent) {
		fmt.Println("L2 pressed with value:", event.Value)
	},
	R2: func(event *sdl.JoyAxisEvent) {
		fmt.Println("R2 pressed with value:", event.Value)
	},
	LeftY: func(event *sdl.JoyAxisEvent) {
		fmt.Println("Left Y axis: value:", event.Value)
	},
	LeftX: func(event *sdl.JoyAxisEvent) {
		fmt.Println("Left X axis: value:", event.Value)
	},
	RightY: func(event *sdl.JoyAxisEvent) {
		fmt.Println("Right Y axis: value:", event.Value)
	},
	RightX: func(event *sdl.JoyAxisEvent) {
		fmt.Println("Right X axis: value:", event.Value)
	},
}

func AxisSwitch(event *sdl.JoyAxisEvent) {
	axisType, ok := axisMap[event.Axis]
	if !ok {
		fmt.Println("no such axis entry in axis map:", event.Axis)
		return
	}
	axisType(event)
}
