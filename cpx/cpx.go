package cpx

import (
	"tinygo-cpx/led"
	"tinygo-cpx/neo"
)

// Init initializes all peripherals on the Circuit Playground Express.
func Init() {
	led.Init()
	neo.Init()
}
