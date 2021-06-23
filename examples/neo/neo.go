package main

import (
	"time"

	"github.com/dbridges/tinygo-cpx/cpx"
	"github.com/dbridges/tinygo-cpx/neo"
)

func main() {
	cpx.Init()

	for {
		for i := uint8(0); i < 10; i++ {
			neo.Clear()
			neo.SetPixel(i, 0, 2, 2)
			neo.Show()
			time.Sleep(500 * time.Millisecond)
		}
	}
}
