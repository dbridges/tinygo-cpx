package main

import (
	"time"
	"tinygo-cpx/cpx"
	"tinygo-cpx/neo"
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
