# Browser game

## Overview

This demo provides a browser game which communicates
via a bidirectional stream with a Go server which
listens for input from the DS4 joystick and sends it
down to the browser

The browser then takes that input and animates the
player in real time


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

# Spins up an HTTP server which listens for input from DS4
go run client.go

# Open up your browser at localhost:8080
# Enjoy
```
 
## More info

- [Go SDL2](https://github.com/veandco/go-sdl2)
- [SDL](https://www.libsdl.org/)
- [Gobot Joystick](https://gobot.io/documentation/platforms/joystick/)

## Back to MDC19

[Home](https://github.com/steevehook/mdc19)
