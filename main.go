package mcp23017

import (
	"periph.io/x/conn/v3/i2c"
)

type Pin uint8
type Bank uint8

const (
	BankA Bank = iota
	BankB
)

const (
	PinA0 Pin = iota
	PinA1
	PinA2
	PinA3
	PinA4
	PinA5
	PinA6
	PinA7
	PinB0
	PinB1
	PinB2
	PinB3
	PinB4
	PinB5
	PinB6
	PinB7
)

const (
	gpioA  byte = 0x12
	gpioB  byte = 0x13
	olataA byte = 0x14
	olataB byte = 0x15
	iodirA byte = 0x00
	iodirB byte = 0x01
	gppuA  byte = 0x0C
	gppuB  byte = 0x0D
)

type MCP23017 struct {
	Dev *i2c.Dev
}

func NewMCP23017(bus i2c.Bus, addr uint16) *MCP23017 {
	return &MCP23017{
		Dev: &i2c.Dev{Bus: bus, Addr: addr},
	}
}

func (m *MCP23017) readRegister(reg byte) (byte, error) {
	write := []byte{reg}
	read := make([]byte, 1)
	if err := m.Dev.Tx(write, read); err != nil {
		return 0, err
	}
	return read[0], nil
}

func (m *MCP23017) writeRegister(reg byte, value byte) error {
	write := []byte{reg, value}
	return m.Dev.Tx(write, nil)
}

func (m *MCP23017) SetPinMode(pin Pin, isInput bool) error {
	var iodirReg byte
	if pin <= 7 {
		iodirReg = iodirA
	} else {
		iodirReg = iodirB
	}
	current, err := m.readRegister(iodirReg)
	if err != nil {
		return err
	}
	if isInput {
		current |= (1 << (uint8(pin) % 8))
	} else {
		current &^= (1 << (uint8(pin) % 8))
	}
	return m.writeRegister(iodirReg, current)
}

func (m *MCP23017) DigitalWrite(pin Pin, value bool) error {
	var gpioReg byte
	if pin <= 7 {
		gpioReg = olataA
	} else {
		gpioReg = olataB
	}
	current, err := m.readRegister(gpioReg)
	if err != nil {
		return err
	}
	if value {
		current |= (1 << (uint8(pin) % 8))
	} else {
		current &^= (1 << (uint8(pin) % 8))
	}
	return m.writeRegister(gpioReg, current)
}

func (m *MCP23017) DigitalReadBank(bank Bank) (byte, error) {
	var gpioReg byte
	if bank == BankA {
		gpioReg = gpioA
	} else {
		gpioReg = gpioB
	}
	return m.readRegister(gpioReg)
}

func (m *MCP23017) DigitalReadAll() (byte, byte, error) {
	gpioAVal, err := m.readRegister(gpioA)
	if err != nil {
		return 0, 0, err
	}
	gpioBVal, err := m.readRegister(gpioB)
	if err != nil {
		return 0, 0, err
	}
	return gpioAVal, gpioBVal, nil
}

func (m *MCP23017) DigitalRead(pin Pin) (bool, error) {
	var gpioReg byte
	if pin <= 7 {
		gpioReg = gpioA
	} else {
		gpioReg = gpioB
	}
	current, err := m.readRegister(gpioReg)
	if err != nil {
		return false, err
	}
	return (current&(1<<(uint8(pin)%8)) != 0), nil
}

func (m *MCP23017) SetPullUp(pin Pin, enable bool) error {
	var gppuReg byte
	if pin <= 7 {
		gppuReg = gppuA
	} else {
		gppuReg = gppuB
	}
	current, err := m.readRegister(gppuReg)
	if err != nil {
		return err
	}
	if enable {
		current |= (1 << (uint8(pin) % 8))
	} else {
		current &^= (1 << (uint8(pin) % 8))
	}
	return m.writeRegister(gppuReg, current)
}

// func main() {
// 	host.Init()
// 	b, err := i2creg.Open("")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer b.Close()

// 	d := &i2c.Dev{Addr: 0x20, Bus: b}

// 	write := []byte{0x10}
// 	read := make([]byte, 5)
// 	if err := d.Tx(write, read); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%v\n", read)
// }
