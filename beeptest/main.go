package main

import (
	"fmt"
	"os"
	"time"

	"github.com/a-h/beeper"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// See https://pinout.xyz to select pins.
	// Need a PWM pin
	pin := rpio.Pin(12)
	beeper.Beep(pin, 440, time.Millisecond*500)
	time.Sleep(time.Second)
	playTune(pin)
}

func playTune(pin rpio.Pin) {
	bpm := 140
	m := beeper.NewMusic(pin, bpm)
	m.Note("A5", beeper.Quaver)
	m.Note("B5", beeper.Quaver)
	m.Note("D5", beeper.Quaver)
	m.Note("B5", beeper.Quaver)
	m.Note("F#5", beeper.Crotchet)
	m.Note("F#5", beeper.Crotchet)
	m.Note("E5", beeper.Crotchet)
	m.Rest(beeper.Crotchet)
	m.Note("A5", beeper.Quaver)
	m.Note("B5", beeper.Quaver)
	m.Note("D5", beeper.Quaver)
	m.Note("B5", beeper.Quaver)
	m.Note("E5", beeper.Crotchet)
	m.Note("E5", beeper.Crotchet)
	m.Note("D5", beeper.Quaver)
	m.Note("C#5", beeper.Quaver)
	m.Note("B4", beeper.Quaver)
	m.Rest(beeper.Crotchet)
	m.Note("A5", beeper.Quaver)
	m.Note("B5", beeper.Quaver)
	m.Note("D5", beeper.Quaver)
	m.Note("B5", beeper.Quaver)
	m.Note("D5", beeper.Crotchet)
	m.Note("E5", beeper.Quaver)
	m.Note("C#5", beeper.Quaver)
	m.Rest(beeper.Quaver)
	m.Note("A4", beeper.Quaver)
	m.Note("E5", beeper.Crotchet)
	m.Note("D5", beeper.Crotchet)
}
