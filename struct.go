package main

const (
	height = byte(0x20)
	width  = byte(0x40)
)

type Chip8 struct {
	memory     [4096]uint8
	i          uint16
	pc         uint16
	stack      [16]uint16
	sp         uint16
	v          [16]byte
	delayTimer byte
	soundTimer byte
	keypad     [16]byte
	video      [height][width]uint32
	draw       bool
	register   byte
	flag       bool
}
