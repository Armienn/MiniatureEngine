package compiler

import (
	"os"

	"github.com/Armienn/MiniatureEngine/machine"
)

func Compile(inputFile string, outputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := ReadAndParse(file)
	operations := ConstructProgram(lines)
	err = WriteProgram(operations)
	return err
}

func ConstructProgram(lines []Line) []machine.Operation {
	return nil
}

func WriteProgram([]machine.Operation) error {
	return nil
}
