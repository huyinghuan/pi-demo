package work

import (
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all" // This loads the RPi driver
)

//Blink 闪灯
//示意图见circuitdesign/blink.png
func Blink() {
	embd.InitGPIO()
	embd.SetDirection(17, embd.Out)
	for range time.Tick(time.Second * 1) {
		//AlarmByEmail()
		embd.DigitalWrite(17, embd.High)
		time.Sleep(100 * time.Millisecond)
		embd.DigitalWrite(17, embd.Low)
	}
}
