package compiler

import (
	"bufio"
	"encoding/binary"
	"os"

	"github.com/Armienn/MiniatureEngine/machine"
)

func Compile(inputFile string, outputFile string) error {
	lines, err := ReadAndParseFile(inputFile)
	if err != nil {
		return err
	}

	operations, err := ConstructProgram(lines)
	if err != nil {
		return err
	}

	err = WriteProgramFile(operations, outputFile)
	return err
}

func ReadAndParseFile(inputFile string) ([]Line, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadAndParse(file), err
}

func WriteProgramFile(operations []machine.Operation, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return WriteProgram(operations, file)
}

func WriteProgram(operations []machine.Operation, file *os.File) error {
	writer := bufio.NewWriter(file)
	err := binary.Write(writer, binary.LittleEndian, operations)
	return err
}
