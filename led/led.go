package led

import "machine"

var led = machine.LED

// Init initializes the built in red LED.
func Init() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

// Set sets the state of the buil in red LED.
func Set(val bool) {
	led.Set(val)
}
