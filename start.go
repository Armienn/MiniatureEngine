package main

import "github.com/Armienn/MiniatureEngine/machine"

func main() {
	program := []machine.Operation{
		{machine.RUNE, 61, 2},
		{machine.NONE, 62, 8},
		{machine.GRUNE, 1, 8},
		{machine.RUNER, 1, 8},
		{machine.JMP, 0, 2},
	}

	cpu := new(machine.CPU)
	cpu.RunProgram(program)
}
