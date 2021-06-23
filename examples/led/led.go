package main

import (
	"time"

	"github.com/dbridges/tinygo-cpx/cpx"
	"github.com/dbridges/tinygo-cpx/led"
)

func main() {
	cpx.Init()

	for {
		led.Set(true)
		time.Sleep(500 * time.Millisecond)
		led.Set(false)
		time.Sleep(500 * time.Millisecond)
	}
}
