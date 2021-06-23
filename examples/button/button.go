package main

import (
	"github.com/dbridges/tinygo-cpx/button"
	"github.com/dbridges/tinygo-cpx/cpx"
	"github.com/dbridges/tinygo-cpx/led"
)

func main() {
	cpx.Init()

	for {
		if button.GetA() || button.GetB() {
			led.Set(true)
		} else {
			led.Set(false)
		}
	}
}
