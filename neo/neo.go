package neo

import (
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

const bufferLength = 10 * 3
const pin = machine.D8

var buffer [bufferLength]byte
var ws ws2812.Device

// Init intializes the CPX's neo pixel driver.
func Init() {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws = ws2812.New(pin)
}

// SetPixel sets the specified pixel to the RGB value given.
func SetPixel(n, r, g, b uint8) {
	buffer[n*3] = g
	buffer[n*3+1] = r
	buffer[n*3+2] = b
}

// GetPixel returns the rgb color values for the given pixel index.
func GetPixel(n uint8) (r, g, b uint8) {
	return buffer[n*3+1], buffer[n*3], buffer[n*3+2]
}

// Show initiates the buffer transfer to display the leds.
func Show() error {
	_, err := ws.Write(buffer[:])
	if err != nil {
		return err
	}
	return nil
}

// Clear resets all pixel values to zero.
func Clear() {
	for i := 0; i < bufferLength; i++ {
		buffer[i] = 0x00
	}
}
