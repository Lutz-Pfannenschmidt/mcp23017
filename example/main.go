package main

import (
	"log"

	"github.com/Lutz-Pfannenschmidt/mcp23017"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

func main() {
	host.Init()
	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	defer b.Close()

	m1 := mcp23017.NewMCP23017(b, 0x20)

	err = m1.SetPinMode(mcp23017.PinA0, false) // Set PinA0 as Output
	_ = err                                    // ignore error handling for this example

	err = m1.DigitalWrite(mcp23017.PinA0, true) // Set PinA0 High
	_ = err                                     // ignore error handling for this example

	val, err := m1.DigitalRead(mcp23017.PinA0) // Read PinA0, returns a boolean
	_ = err                                    // ignore error handling for this example

	log.Printf("PinA0 Value: %v", val)
}
