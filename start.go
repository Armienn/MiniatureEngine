package main

import (
	"fmt"
	"os"

	"github.com/Armienn/MiniatureEngine/compiler"
	"github.com/Armienn/MiniatureEngine/machine"
)

func main() {
	if len(os.Args) < 2 {
		engine := new(machine.MiniatureEngine)
		err := engine.LoadProgram("program.bin")
		if err != nil {
			fmt.Println(err)
			return
		}
		engine.Run()
		return
	}

	switch os.Args[1] {

	case "compile":
		if len(os.Args) == 2 {
			err := compiler.Compile("program.txt", "program.bin")
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		if len(os.Args) == 3 {
			err := compiler.Compile(os.Args[2], "program.bin")
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		err := compiler.Compile(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println(err)
		}
		return

	case "run":
		if len(os.Args) == 2 {
			engine := new(machine.MiniatureEngine)
			err := engine.LoadProgram("program.bin")
			if err != nil {
				fmt.Println(err)
				return
			}
			engine.Run()
			return
		}
		engine := new(machine.MiniatureEngine)
		err := engine.LoadProgram(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		engine.Run()
		return

	default:
		fmt.Println("Invalid argument")
	}
}
