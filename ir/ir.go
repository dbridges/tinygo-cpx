// Package ir provides a basic interface to the ir peripherals.
// Currently only NEC decoding is supported.
package ir

import (
	"machine"
	"time"
)

const irTxPin = machine.PA23
const irRxPin = machine.PA12

const (
	stateIdle         uint8 = 0
	stateWaitForStart       = 1
	stateReceiving          = 3
)

// All units in nanoseconds
const pulseLength = 562500
const highBitPulseWidthCutoff = 3 * pulseLength
const frameTimeout = 6 * pulseLength
const startPulseWidth = 9 * pulseLength
const startPulseMargin = 100000

var bit uint8
var data uint32
var state uint8
var lastPulse = time.Now().UnixNano()

// Init initializes the ir transmitter and receiver/decoder.
func Init() {
	irTxPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	irRxPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	irRxPin.SetInterrupt(machine.PinRising, handleRxFall)
}

// HasData returns true or false if there is data available.
func HasData() bool {
	return data != 0 && state == stateIdle
}

// GetData returns the current data value
func GetData() uint32 {
	return data
}

// EnableReceive resets data and waits for a new signal to arrive.
func EnableReceive() {
	data = 0
	state = stateWaitForStart
}

func handleRxFall(rx machine.Pin) {
	t := time.Now().UnixNano()
	dt := t - lastPulse

	switch state {
	case stateIdle:
		break
	case stateWaitForStart:
		if dt > startPulseWidth-startPulseMargin && dt < startPulseWidth+startPulseMargin {
			state = stateReceiving
			bit = 0
			data = 0
		}
	case stateReceiving:
		if dt > frameTimeout {
			state = stateWaitForStart
		} else if dt > highBitPulseWidthCutoff {
			data |= (1 << (31 - bit))
		}
		if bit >= 31 {
			state = stateIdle
		}
		bit++
	}

	lastPulse = t
}
