package main

import (
	"machine"
	"time"
)

func main() {
	// Configure the integrated led.
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Configure the ADC0 pin on the Arduino board.
	machine.InitADC()
	sensor := machine.ADC{Pin: machine.ADC0}
	sensor.Configure(machine.ADCConfig{})

	/*
		Official values of "Grove Moisture Sensor" in a 0-1024 scale:

		Output Value           Min     Max
		-------------------   -----   -----
		Sensor in dry soil       0     300
		Sensor in humid soil   300     700
		Sensor in water        700     950
	*/
	for {
		// Divide by 64 to "convert" from tinygo adc scale (0xffff, 65535) to moisture sensor scale (0-1024).
		moisture := sensor.Get() / 64
		println("The value of the moisture sensor is: ", moisture)
		// If the moisture sensor returns a value equal or under 300, we consider it dry.
		// We will turn on the LED to indicate this, and keep that state until the plant is watered.
		if moisture <= 300 {
			led.High()
		} else {
			led.Low()
		}
		// Wait a second to read the sensor value again.
		time.Sleep(1 * time.Second)
	}
}
