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
	err = WriteProgram(operations, outputFile)
	return err
}

func WriteProgram(operations []machine.Operation, outputFile string) error {
	return nil
}
