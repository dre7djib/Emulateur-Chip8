package main

import ( // Assurez-vous d'importer votre package Chip8
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	setupKeys()
	chip8 = NewChip8()
	println(chip8)
	ebiten.SetMaxTPS(60)
	fmt.Println("Enter the file of the game :")
	var game string
	fmt.Scanln(&game)
	chip8.LoadROM("roms/" + game)
	if err := ebiten.Run(updt, 640, 320, 1, "Emulateur Chip8"); err != nil {
		panic(err)
	}
}
