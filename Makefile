.PHONY: button ir led neo

button:
	tinygo flash -target=circuitplay-express examples/button/button.go

ir:
	tinygo flash -target=circuitplay-express examples/ir/ir.go

led:
	tinygo flash -target=circuitplay-express examples/led/led.go

neo:
	tinygo flash -target=circuitplay-express examples/neo/neo.go
