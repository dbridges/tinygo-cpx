package cpx

import (
	"tinygo-cpx/button"
	"tinygo-cpx/ir"
	"tinygo-cpx/led"
	"tinygo-cpx/neo"
)

// Init initializes all peripherals on the Circuit Playground Express.
func Init() {
	button.Init()
	led.Init()
	neo.Init()
	ir.Init()
}
