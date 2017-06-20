package main

import (
	"log"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all" // This loads the RPi driver
)

func main() {
	embd.InitGPIO()
	defer embd.CloseGPIO()
	embd.SetDirection(17, embd.Out)
	pwm, err := embd.NewPWMPin(17)
	if err != nil {
		log.Fatal(err)
		return
	}

	pwm.SetPeriod(131)
	pwm.SetPeriod(50)
	pwm.SetMicroseconds(1000)
}
