package main

import (
	"time"
	"tinygo-cpx/cpx"
	"tinygo-cpx/ir"
)

func main() {
	cpx.Init()

	ir.EnableReceive()

	for {
		if ir.HasData() {
			println(ir.GetData())
			time.Sleep(100 * time.Millisecond)
			ir.EnableReceive()
		}
		time.Sleep(100 * time.Millisecond)
	}
}
