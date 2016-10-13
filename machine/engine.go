package machine

import (
	"encoding/binary"
	"os"
)

type MiniatureEngine struct {
	Cpu     CPU
	Program [65536]Operation
}

func (engine *MiniatureEngine) LoadProgram(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; err == nil; i++ {
		err = binary.Read(file, binary.LittleEndian, &(engine.Program[i]))
	}
	return err
}

func (engine *MiniatureEngine) Run() {
	engine.Cpu.RunProgram(engine.Program[:])
}
