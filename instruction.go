package main

import (
	"math/rand"
	"time"
)

func (ch8 *Chip8) Run() {
	ch8.Instruction()

	if ch8.delayTimer > 0 {
		ch8.delayTimer = ch8.delayTimer - 1
	}

}
func (ch8 *Chip8) Instruction() {
	opcode := uint16(ch8.memory[ch8.pc])<<8 | uint16(ch8.memory[ch8.pc+1])
	ch8.pc = ch8.pc + 2
	switch opcode & 0xF000 {
	case 0x0000:
		switch opcode & 0x000F {
		case 0x0000:
			ch8.CLS()
		case 0x000E:
			ch8.pc = ch8.stack[ch8.sp-1]
			ch8.sp = ch8.sp - 1
		}
	case 0x1000:
		ch8.pc = opcode & 0x0FFF
	case 0x2000:
		ch8.stack[ch8.sp] = ch8.pc
		ch8.sp = ch8.sp + 1
		ch8.pc = opcode & 0x0FFF
	case 0x3000:
		compareTo := byte(opcode & 0x00FF)
		register := (opcode & 0x0F00) >> 8
		if ch8.v[register] == compareTo {
			ch8.pc = ch8.pc + 2
		}
	case 0x4000:
		compareTo := byte(opcode & 0x00FF)
		register := (opcode & 0x0F00) >> 8
		if ch8.v[register] != compareTo {
			ch8.pc = ch8.pc + 2
		}
	case 0x5000:
		registerX := (opcode & 0x0F00) >> 8
		registerY := (opcode & 0x00F0) >> 4
		if ch8.v[registerX] == ch8.v[registerY] {
			ch8.pc = ch8.pc + 2
		}
	case 0x6000:
		register := byte((opcode & 0x0F00) >> 8)
		ch8.v[register] = byte(opcode & 0x00FF)
	case 0x7000:
		register := byte((opcode & 0x0F00) >> 8)
		value := byte(opcode & 0x00FF)
		ch8.v[register] = ch8.v[register] + value
	case 0x8000:
		switch opcode & 0x000F {
		case 0x0000:
			registerX := (opcode & 0x0F00) >> 8
			registerY := (opcode & 0x00F0) >> 4
			ch8.v[registerX] = ch8.v[registerY]
		case 0x0001:
			registerX := (opcode & 0x0F00) >> 8
			registerY := (opcode & 0x00F0) >> 4
			ch8.v[registerX] = ch8.v[registerX] | ch8.v[registerY]
		case 0x0002:
			registerX := (opcode & 0x0F00) >> 8
			registerY := (opcode & 0x00F0) >> 4
			ch8.v[registerX] = ch8.v[registerX] & ch8.v[registerY]
		case 0x0003:
			registerX := (opcode & 0x0F00) >> 8
			registerY := (opcode & 0x00F0) >> 4
			ch8.v[registerX] = ch8.v[registerX] ^ ch8.v[registerY]
		case 0x0004:
			registerX := byte((opcode & 0x0F00) >> 8)
			registerY := byte((opcode & 0x00F0) >> 4)
			ch8.v[registerX] = ch8.v[registerX] + ch8.v[registerY]
			if uint16(ch8.v[registerX])+uint16(ch8.v[registerY]) > 0xFF {
				ch8.v[0xF] = 1
			} else {
				ch8.v[0xF] = 0
			}
		case 0x0005:
			registerX := (opcode & 0x0F00) >> 8
			registerY := (opcode & 0x00F0) >> 4
			if ch8.v[registerX] > ch8.v[registerY] {
				ch8.v[0xF] = 1
			} else {
				ch8.v[0xF] = 0
			}
			ch8.v[registerX] = ch8.v[registerX] - ch8.v[registerY]
		case 0x0006:
			registerX := (opcode & 0x0F00) >> 8
			if ch8.v[registerX]&0x1 == 1 {
				ch8.v[0xF] = 1
			} else {
				ch8.v[0xF] = 0
			}
			ch8.v[registerX] = ch8.v[registerX] >> 1
		case 0x0007:
			registerX := (opcode & 0x0F00) >> 8
			registerY := (opcode & 0x00F0) >> 4
			if ch8.v[registerY] > ch8.v[registerX] {
				ch8.v[0xF] = 1
			} else {
				ch8.v[0xF] = 0
			}
			ch8.v[registerX] = ch8.v[registerY] - ch8.v[registerX]
		case 0x000E:
			registerX := (opcode & 0x0F00) >> 8
			if ch8.v[registerX]&0x80 == 0x80 {
				ch8.v[0xF] = 1
			} else {
				ch8.v[0xF] = 0
			}
			ch8.v[registerX] = ch8.v[registerX] << 1
		}
	case 0x9000:
		registerX := (opcode & 0x0F00) >> 8
		registerY := (opcode & 0x00F0) >> 4
		if ch8.v[registerX] != ch8.v[registerY] {
			ch8.pc = ch8.pc + 2
		}
	case 0xA000:
		ch8.i = (opcode & 0x0FFF)
	case 0xB000:
		ch8.pc = (opcode & 0x0FFF) + uint16(ch8.v[0x0])
	case 0xC000:
		registerX := (opcode & 0x0F00) >> 8
		value := byte(opcode & 0x00FF)
		rand.Seed(time.Now().Unix())
		ch8.v[registerX] = byte(rand.Intn(256)) & value
	case 0xD000:
		registerX := (opcode & 0x0F00) >> 8
		registerY := (opcode & 0x00F0) >> 4
		nibble := byte(opcode & 0x000F)
		x := ch8.v[registerX]
		y := ch8.v[registerY]
		ch8.v[0xF] = 0x00
		for i := y; i < y+nibble; i++ {
			for j := x; j < x+8; j++ {
				bit := (ch8.memory[ch8.i+uint16(i-y)] >> (7 - j + x)) & 0x01
				xIndex, yIndex := j, i
				if j >= byte(0x40) {
					xIndex = j - byte(0x40)
				}
				if i >= byte(0x20) {
					yIndex = i - byte(0x20)
				}
				if bit == 0x01 && ch8.video[yIndex][xIndex] == 0x01 {
					ch8.v[0xF] = 0x01
				}
				ch8.video[yIndex][xIndex] = ch8.video[yIndex][xIndex] ^ uint32(bit)
			}
		}
		ch8.draw = true
	case 0xE000:
		switch opcode & 0x00FF {
		case 0x009E:
			register := (opcode & 0x0F00) >> 8
			if ch8.keypad[ch8.v[register]] == 0x01 {
				ch8.pc = ch8.pc + 2
			}
		case 0x00A1:
			register := (opcode & 0x0F00) >> 8
			if ch8.keypad[ch8.v[register]] == 0x00 {
				ch8.pc = ch8.pc + 2
			}
		}
	case 0xF000:
		switch opcode & 0x00FF {
		case 0x007:
			register := (opcode & 0x0F00) >> 8
			ch8.v[register] = ch8.delayTimer
		case 0x0015:
			register := (opcode & 0x0F00) >> 8
			ch8.delayTimer = ch8.v[register]
		case 0x0018:
			register := (opcode & 0x0F00) >> 8
			ch8.soundTimer = ch8.v[register]
		case 0x000A:
			register := (opcode & 0x0F00) >> 8
			ch8.flag = true
			ch8.register = byte(register)
		case 0x001E:
			register := (opcode & 0x0F00) >> 8
			ch8.i = ch8.i + uint16(ch8.v[register])
		case 0x0029:
			register := (opcode & 0x0F00) >> 8
			ch8.i = uint16(ch8.v[register] * 0x5)
		case 0x0033:
			register := (opcode & 0x0F00) >> 8
			number := ch8.v[register]
			ch8.memory[ch8.i] = (number / 100) % 10
			ch8.memory[ch8.i+1] = (number / 10) % 10
			ch8.memory[ch8.i+2] = number % 10
		case 0x0055:
			register := (opcode & 0x0F00) >> 8
			for i := uint16(0x00); i <= register; i++ {
				ch8.memory[ch8.i+i] = ch8.v[i]
			}
		case 0x0065:
			register := (opcode & 0x0F00) >> 8
			for i := uint16(0x00); i <= register; i++ {
				ch8.v[i] = ch8.memory[ch8.i+i]
			}
		}
	}
}

func (ch8 *Chip8) CLS() {
	for x := 0x00; x < 0x20; x++ {
		for y := 0x00; y < 0x40; y++ {
			ch8.video[x][y] = 0x00
		}
	}
}
