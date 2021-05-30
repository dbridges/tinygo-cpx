.PHONY: all led neo

all: led

led:
	tinygo flash -target=circuitplay-express examples/led/led.go

neo:
	tinygo flash -target=circuitplay-express examples/neo/neo.go
