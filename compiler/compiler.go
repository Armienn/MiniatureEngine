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

	lines := readAndParse(file)
	operations := constructProgram(lines)
	err = writeProgram(operations)
	return err
}

func writeProgram([]machine.Operation) error {
	return nil
}

func constructProgram(lines []Line) []machine.Operation {
	return nil
}
