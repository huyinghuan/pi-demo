package main

import (
	"pi/work"

	_ "github.com/kidoman/embd/host/all" // This loads the RPi driver
)

func main() {
	work.ShowTemperature()
}
