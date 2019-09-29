# beeper

A Go library for using a piezo buzzer on the Raspberry Pi via the PWM output.

# Usage

Connect a piezo to the PWM output pin on the Pi (e.g. BCM pin 12, or pin 32 on the Raspberry Pi zero - see https://pinout.xyz for more details).


## Setup

```go
err := rpio.Open()
if err != nil {
  fmt.Printf("error: %v\n", err)
  os.Exit(1)
}
defer rpio.Close()
```

## Beep

```go
pin := rpio.Pin(12)
beeper.Beep(pin, 440, time.Millisecond*500)
```

## Play a tune

```go
pin := rpio.Pin(12)
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
```