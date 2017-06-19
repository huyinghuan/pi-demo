package main

import (
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all" // This loads the RPi driver
)

func main() {
	embd.InitGPIO()
	defer embd.CloseGPIO()
	embd.SetDirection(17, embd.Out)
	for range time.Tick(time.Second * 1) {
		//AlarmByEmail()
		embd.DigitalWrite(17, embd.High)
		time.Sleep(100 * time.Millisecond)
		embd.DigitalWrite(17, embd.Low)
	}
}
