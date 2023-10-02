package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

var (
	chip8 *Chip8
)

var keyMap map[ebiten.Key]byte

var audioPlayer *audio.Player

func setupKeys() {
	keyMap = make(map[ebiten.Key]byte)
	keyMap[ebiten.Key1] = 0x01
	keyMap[ebiten.Key2] = 0x02
	keyMap[ebiten.Key3] = 0x03
	keyMap[ebiten.Key4] = 0x0C
	keyMap[ebiten.KeyQ] = 0x04
	keyMap[ebiten.KeyW] = 0x05
	keyMap[ebiten.KeyE] = 0x06
	keyMap[ebiten.KeyR] = 0x0D
	keyMap[ebiten.KeyA] = 0x07
	keyMap[ebiten.KeyS] = 0x08
	keyMap[ebiten.KeyD] = 0x09
	keyMap[ebiten.KeyF] = 0x0E
	keyMap[ebiten.KeyZ] = 0x0A
	keyMap[ebiten.KeyX] = 0x00
	keyMap[ebiten.KeyC] = 0x0B
	keyMap[ebiten.KeyV] = 0x0F
}

var (
	square *ebiten.Image
)

func init() {
	square, _ = ebiten.NewImage(10, 10, ebiten.FilterNearest)
	square.Fill(color.White)
}

func getInput() bool {
	for key, value := range keyMap {
		if ebiten.IsKeyPressed(key) {
			chip8.keypad[value] = 0x01
			return true
		}
	}
	return false
}

func update(screen *ebiten.Image) error {

	// fill screen
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	for i := 0; i < 10; i++ {

		chip8.draw = false
		chip8.flag = false
		gotInput := true
		chip8.Run()

		if chip8.flag {
			gotInput = getInput()
			if !gotInput {
				chip8.pc = chip8.pc - 2
			}
		}

		if chip8.draw || !gotInput {
			for i := 0; i < 32; i++ {
				for j := 0; j < 64; j++ {
					if chip8.video[i][j] == 0x01 {

						opts := &ebiten.DrawImageOptions{}

						opts.GeoM.Translate(float64(j*10), float64(i*10))

						screen.DrawImage(square, opts)
					}
				}
			}
		}
		for key, value := range keyMap {
			if ebiten.IsKeyPressed(key) {
				chip8.keypad[value] = 0x01
			} else {
				chip8.keypad[value] = 0x00
			}
		}

		if chip8.soundTimer > 0 {
			audioPlayer.Play()
			audioPlayer.Rewind()
		}

	}

	return nil
}
