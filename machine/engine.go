package machine

import (
	"encoding/binary"
	"os"
)

//MiniatureEngine represents a virtual computer
type MiniatureEngine struct {
	//Cpu is the CPU of the MiniatureEngine
	Cpu CPU
	//Program is the program memory of the MiniatureEngine
	Program [65536]Operation
}

//LoadProgram loads a program into a MiniatureEngine
func (engine *MiniatureEngine) LoadProgram(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; err == nil; i++ {
		err = binary.Read(file, binary.LittleEndian, &(engine.Program[i]))
	}
	return nil
}

//Run runs the currently loaded program
func (engine *MiniatureEngine) Run() {
	engine.Cpu.RunProgram(engine.Program[:])
}
