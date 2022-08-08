package main

import (
	"machine"
	"time"
)

func main() {
	machine.InitADC()

	// TODO(lopezator): No led for now
	// led := machine.LED
	// led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Configure the ADC0 pin is the analog pin on the Arduino board.
	sensor := machine.ADC{Pin: machine.ADC0}
	sensor.Configure(machine.ADCConfig{})

	for {
		val := sensor.Get()
		/*
			    Official values of "Grove Moisture Sensor" in a 0-1024 scale:

				Output Value           Min     Max
				Sensor in dry soil       0     300
				Sensor in humid soil   300     700
				Sensor in water        700     950
		*/
		// TODO(lopezator): Why this operation doesn't work and always return zero?
		// val = val * 1024 / 65535
		println("The value of the moisture sensor is: ", val)
		time.Sleep(time.Millisecond * 500)
	}
}
