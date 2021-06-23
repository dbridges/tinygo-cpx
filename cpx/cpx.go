package cpx

import (
	"github.com/dbridges/tinygo-cpx/button"
	"github.com/dbridges/tinygo-cpx/ir"
	"github.com/dbridges/tinygo-cpx/led"
	"github.com/dbridges/tinygo-cpx/neo"
)

// Init initializes all peripherals on the Circuit Playground Express.
func Init() {
	button.Init()
	led.Init()
	neo.Init()
	ir.Init()
}
