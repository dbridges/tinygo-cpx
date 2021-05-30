# tinygo-cpx

A utility library for [TinyGo](https://tinygo.org) and Adafruit's [Circuit Playground Express (CPX)](https://www.adafruit.com/product/3333). This library wraps many of the [TinyGo drivers](https://github.com/tinygo-org/drivers) for peripherals used on the CPX in an easy to use interface.

Peripherals currently supported:

* Red LED
* NeoPixels

## Usage

```
go get dbridges/tinygo-cpx
```

See the [examples](https://github.com/dbridges/tinygo-cpx/tree/main/examples) directory for information on each peripheral.

A minimal example, flashing the built-in red LED:

```go
package main

import (
	"time"
	"tinygo-cpx/cpx"
	"tinygo-cpx/led"
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
```
