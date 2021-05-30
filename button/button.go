package button

import "machine"

// Init initializes the two push buttons for reading
func Init() {
	machine.BUTTONA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.BUTTONB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
}

// GetA returns the current value of the A push button.
func GetA() bool {
	return machine.BUTTONA.Get()
}

// GetB returns the current value of the B push button.
func GetB() bool {
	return machine.BUTTONB.Get()
}
