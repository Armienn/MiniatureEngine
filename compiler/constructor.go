package compiler

import (
	"strconv"

	"github.com/Armienn/MiniatureEngine/machine"
)

type constructor struct {
	alias      map[string]int
	operations []machine.Operation
	err        error
}

func newConstructor() *constructor {
	constructor := new(constructor)
	constructor.alias = make(map[string]int)
	constructor.operations = make([]machine.Operation, 0, 256)
	return constructor
}

//ConstructProgram constructs a list of operations based on a list of assembler commands
func ConstructProgram(lines []Line) ([]machine.Operation, error) {
	constructor := newConstructor()
	for _, line := range lines {
		constructor.processLine(line)
	}
	return constructor.operations, constructor.err
}

func (constructor *constructor) processLine(line Line) {
	if constructor.err != nil {
		return
	}
	switch line.CommandType {
	case Alias:
		constructor.processAlias(line)
	case LineTag:
		constructor.processLineTag(line)
	case Command:
		constructor.processCommand(line)
	}
}

func (constructor *constructor) processAlias(line Line) {
	value, err := strconv.Atoi(line.Op2)
	if err != nil {
		constructor.err = err
		return
	}
	constructor.alias[line.Op1] = value
}

func (constructor *constructor) processLineTag(line Line) {
	value := len(constructor.operations)
	constructor.alias[line.Op1] = value
}

func (constructor *constructor) processCommand(line Line) {
	var operation machine.Operation
	operation.Type = line.Command
	operation.FirstOperand, operation.SecondOperand = constructor.getValues(line)
	constructor.operations = append(constructor.operations, operation)
}

func (constructor *constructor) getValues(line Line) (byte, byte) {
	var err error
	var op1 int
	var op2 int
	if line.Op1 != "" {
		op1, err = constructor.getValue(line.Op1)
	}
	if line.Op2 != "" {
		op2, err = constructor.getValue(line.Op2)
	}
	constructor.err = err
	if op1 > 255 {
		return byte(op1 / 256), byte(op1 % 256)
	}
	return byte(op1 % 256), byte(op2 % 256)
}

func (constructor *constructor) getValue(source string) (int, error) {
	var err error
	value, ok := constructor.alias[source]
	if !ok {
		value, err = strconv.Atoi(source)
	}
	return value, err
}
