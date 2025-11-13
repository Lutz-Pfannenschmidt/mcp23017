# MCP23017 GPIO Expander Library in Go

This repository provides a Go library for interfacing with the MCP23017 GPIO expander over I2C. The MCP23017 is a 16-bit I/O expander that communicates via the I2C bus, allowing you to add additional GPIO pins to your microcontroller or single-board computer.

## Features

-   [x] Control up to 16 GPIO pins
-   [x] Configure pins as input or output, with optional pull-up resistors
-   [x] Read and write pin states
-   [ ] WIP support for interrupts

## Installation

To install the library, use the following command:

```bash
go get github.com/Lutz-Pfannenschmidt/mcp23017
```

## Usage

Here is a simple example of how to use the MCP23017 library:

```go
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
```

## Documentation

For detailed documentation, please refer to the [GoDoc page](https://pkg.go.dev/github.com/Lutz-Pfannenschmidt/mcp23017).
