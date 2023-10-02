package main

import "chip"

func main() {
	// Create a new Chip8 instance using the NewChip8 constructor function
	ch8 := chip.Chip8{}
	romFile := "roms/PONG.ch8"
	ch8.LoadROM(romFile)

}
