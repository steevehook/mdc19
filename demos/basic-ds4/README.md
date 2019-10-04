# Basic DS4 demo

## Overview

This demo provides basic I/O read
from a PlayStation DualShock 4 joystick and outputting
the telemetry data on the screen

For a more comprehensive example which builds on top
of this one, make sure to check out

[DJI Tello & DS4](https://github.com/steevehook/mdc19/tree/master/demos/dji-tello-ds4)


## Steps to run

```bash
# install SDL2 library for MacOS
brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config

# Download all go libraries used (just go-sdl2)
go get ./...

# Make sure you turn on your DS4
# Long press Home (PS) + Share buttons
# Connect from the Bluetooth options

# Compile & run your program
go run *.go

# Make sure you DON'T have both connections
# USB & from Bluetooth at the same time
# Make sure your DS4 shows up in CLI logs
# Enjoy
```
 
## More info

- [Go SDL2](https://github.com/veandco/go-sdl2)
- [SDL](https://www.libsdl.org/)
- [Gobot Joystick](https://gobot.io/documentation/platforms/joystick/)

## Back to MDC19

[Home](https://github.com/steevehook/mdc19)
