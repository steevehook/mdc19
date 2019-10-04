# DJI Tello & DS4

## Overview

This demo provides a fully featured example which
builds on top of
[Basic DS4](https://github.com/steevehook/mdc19/tree/master/demos/basic-ds4)
example

So in the end, you are able to control your DJI Tello drone
using a DS4 joystick

Before running this example, make sure that your DS4
key mappings are ok, and you actually input what you press.

Check out again the
[Basic DS4](https://github.com/steevehook/mdc19/tree/master/demos/basic-ds4)
example and adjust the mappings inside `buttons.go`


## Steps to run

```bash
# install SDL2 library for MacOS
brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config

# Download all go libraries used
go get ./...

# Make sure you turn on your DS4
# Long press Home (PS) + Share buttons
# Connect from the Bluetooth options

# Make sure to turn the DJI Tello drone ON
# Make sure to connect to its WiFi network

# Compile & run your program
go run *.go

# If it does not work, give the Tello a restart
# and rerun this program (IT SHOULD WORK)
```
 
## More info

- [Go SDL2](https://github.com/veandco/go-sdl2)
- [SDL](https://www.libsdl.org/)
- [Gobot Joystick](https://gobot.io/documentation/platforms/joystick/)
- [Gobot Tello](https://gobot.io/documentation/platforms/tello/)

## Back to MDC19

[Home](https://github.com/steevehook/mdc19)
