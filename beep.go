package beeper

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

// Beep the buzzer.
func Beep(pin rpio.Pin, freq int, duration time.Duration) {
	pin.Pwm()
	pin.Freq(freq * 32)
	pin.DutyCycle(1, 32)
	rpio.StartPwm()
	time.Sleep(duration)
	rpio.StopPwm()
}
