package main

import (
	"tinygo-cpx/button"
	"tinygo-cpx/cpx"
	"tinygo-cpx/led"
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
