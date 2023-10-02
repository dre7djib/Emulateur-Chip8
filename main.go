package main

import ( // Assurez-vous d'importer votre package Chip8
	"github.com/hajimehoshi/ebiten"
)

func main() {
	setupKeys()
	chip8 = NewChip8()
	println(chip8)
	chip8.LoadROM("roms/invaders.ch8")
	if err := ebiten.Run(update, 640, 320, 1, "JEU"); err != nil {
		panic(err)
	}
}
