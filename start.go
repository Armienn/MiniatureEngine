package main

import (
	"fmt"

	"github.com/Armienn/MiniatureEngine/machine"
	"github.com/Armienn/MiniatureEngine/mulle"
)

func main() {
	fmt.Print("hahaee")
	b := make(map[string]int)
	fmt.Print(b["moskus"])
	b["moskus"] = 23
	fmt.Print(b["moskus"])
	fmt.Println(mulle.DoThing())

	var ints []int
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	ints = arr[:]
	fmt.Print("arr")
	fmt.Println(arr)
	fmt.Print("ints")
	fmt.Println(ints)
	ints[0] = 999
	fmt.Print("arr")
	fmt.Println(arr)
	fmt.Print("ints")
	fmt.Println(ints)
	fmt.Println()

	ints = ints[:5]
	fmt.Print("arr")
	fmt.Println(arr)
	fmt.Print("ints")
	fmt.Println(ints)
	ints[2] = 999
	fmt.Print("arr")
	fmt.Println(arr)
	fmt.Print("ints")
	fmt.Println(ints)
	fmt.Println()

	ints = ints[:7]
	fmt.Print("arr")
	fmt.Println(arr)
	fmt.Print("ints")
	fmt.Println(ints)
	ints[3] = 999
	fmt.Print("arr")
	fmt.Println(arr)
	fmt.Print("ints")
	fmt.Println(ints)
	fmt.Println()

	ints = append(ints, 23, 43, 23, 534)
	fmt.Print("arr")
	fmt.Println(arr)
	fmt.Print("ints")
	fmt.Println(ints)
	ints[1] = 999
	fmt.Print("arr")
	fmt.Println(arr)
	fmt.Print("ints")
	fmt.Println(ints)
	fmt.Println()

	program := []machine.Operation{
		{machine.GEM, 1, 2},
		{machine.HENT, 7, 8},
	}

	cpu := new(machine.Cpu)
	cpu.RunProgram(program)
}
