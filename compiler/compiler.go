package compiler

import (
	"encoding/binary"
	"os"

	"github.com/Armienn/MiniatureEngine/machine"
)

//Compile compiles the code found in the given input files and writes the result to the output file
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

//ReadAndParseFile reads the content of the input file and parses it into a list of Line
func ReadAndParseFile(inputFile string) ([]Line, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadAndParse(file), err
}

//WriteProgramFile writes a binary program into the output file based on the given operations
func WriteProgramFile(operations []machine.Operation, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return WriteProgram(operations, file)
}

//WriteProgram writes a binary program into the output file based on the given operations
func WriteProgram(operations []machine.Operation, file *os.File) error {
	for _, v := range operations {
		err := binary.Write(file, binary.LittleEndian, v)
		if err != nil {
			return err
		}
	}
	return nil
}
