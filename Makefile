.PHONY: button led neo

button:
	tinygo flash -target=circuitplay-express examples/button/button.go

led:
	tinygo flash -target=circuitplay-express examples/led/led.go

neo:
	tinygo flash -target=circuitplay-express examples/neo/neo.go
