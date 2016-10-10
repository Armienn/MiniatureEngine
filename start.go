package main

import (
	"fmt"

	"github.com/Armienn/MiniatureEngine/machine"
)

func main() {
	engine := new(machine.MiniatureEngine)
	engine.LoadProgram(".gitignore")
	for _, e := range engine.Program {
		if e.FirstOperand == 0 &&
			e.SecondOperand == 0 &&
			e.Type == 0 {
			break
		}
		fmt.Println(e)
	}
	/*
		program := []machine.Operation{
			{machine.RUNE, 87, 2},
			{machine.RUNE, 114, 2},
			{machine.RUNE, 105, 2},
			{machine.RUNE, 116, 2},
			{machine.RUNE, 101, 2},
			{machine.RUNE, 32, 2},
			{machine.RUNE, 115, 2},
			{machine.RUNE, 111, 2},
			{machine.RUNE, 109, 2},
			{machine.RUNE, 101, 2},
			{machine.RUNE, 116, 2},
			{machine.RUNE, 104, 2},
			{machine.RUNE, 105, 2},
			{machine.RUNE, 110, 2},
			{machine.RUNE, 103, 2},
			{machine.RUNE, 58, 2},
			{machine.RUNE, '\n', 2},
			{machine.GRUNE, 1, 8},
			{machine.RUNER, 1, 8},
			{machine.JMP, 0, 17},
		}

		cpu := new(machine.CPU)
		cpu.RunProgram(program)*/
}
