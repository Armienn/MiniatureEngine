package machine

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// CPU represents the content of a cpu
type CPU struct {
	Registers      [16]byte
	Memory         [65536]byte
	ProgramCounter int
}

//r12: low addr
//r13: high addr
//r14: stack pointer
//r15: div bits

// RunProgram runs the given program
func (cpu *CPU) RunProgram(program []Operation) {
	cpu.ProgramCounter = 0
	reader := bufio.NewReader(os.Stdin)
	for cpu.ProgramCounter = 0; cpu.ProgramCounter < len(program); cpu.ProgramCounter++ {
		operation := program[cpu.ProgramCounter]
		cpu.performOperation(operation, reader)
		//cpu.printState(operation)
	}
}

func (cpu *CPU) performOperation(operation Operation, reader *bufio.Reader) {
	temp := 0
	switch operation.Type {
	case NONE:
		time.Sleep(100 * time.Millisecond)

	case ADD:
		temp = int(cpu.Registers[operation.FirstOperand]) + int(operation.SecondOperand)
		cpu.setOverflow(temp > 255)
		cpu.Registers[operation.FirstOperand] = byte(temp)

	case ADDR:
		temp = int(cpu.Registers[operation.FirstOperand]) + int(cpu.Registers[operation.SecondOperand])
		cpu.setOverflow(temp > 255)
		cpu.Registers[operation.FirstOperand] = byte(temp)

	case SUB:
		temp = int(cpu.Registers[operation.FirstOperand]) - int(operation.SecondOperand)
		cpu.setOverflow(temp < 0)
		cpu.Registers[operation.FirstOperand] = byte(temp)

	case SUBR:
		temp = int(cpu.Registers[operation.FirstOperand]) + int(cpu.Registers[operation.SecondOperand])
		cpu.setOverflow(temp < 0)
		cpu.Registers[operation.FirstOperand] = byte(temp)

	case SET:
		cpu.Registers[operation.FirstOperand] = operation.SecondOperand
	case SETR:
		cpu.Registers[operation.FirstOperand] = cpu.Registers[operation.SecondOperand]

	case DEC:
		cpu.Registers[operation.FirstOperand]--
		cpu.setOverflow(cpu.Registers[operation.FirstOperand] == 255)
	case INC:
		cpu.Registers[operation.FirstOperand]++
		cpu.setOverflow(cpu.Registers[operation.FirstOperand] == 0)

	case JMP:
		cpu.ProgramCounter = int(operation.FirstOperand)*256 + int(operation.SecondOperand) - 1

	case IF:
		if !(cpu.Registers[operation.FirstOperand] > 0) {
			cpu.ProgramCounter++
		}
	case IFN:
		if !(cpu.Registers[operation.FirstOperand] == 0) {
			cpu.ProgramCounter++
		}
	case IFEQ:
		if cpu.Registers[operation.FirstOperand] != operation.SecondOperand {
			cpu.ProgramCounter++
		}
	case IFEQR:
		if cpu.Registers[operation.FirstOperand] != cpu.Registers[operation.SecondOperand] {
			cpu.ProgramCounter++
		}
	case IFOV:
		if !(cpu.Registers[15] > 0) {
			cpu.ProgramCounter++
		}

	case SHOW:
		fmt.Print(operation.FirstOperand)
	case SHOWR:
		fmt.Print(cpu.Registers[operation.FirstOperand])

	case RUNE:
		fmt.Print(string(operation.FirstOperand))
	case RUNER:
		fmt.Print(string(cpu.Registers[operation.FirstOperand]))
	case GRUNE:
		cpu.Registers[operation.FirstOperand], _ = reader.ReadByte() //this blocks until enter has been pressed
		//var err error
		//cpu.Registers[operation.FirstOperand], err = reader.ReadByte()
		//cpu.setOverflow(err != nil)

	case STR:
		temp = int(cpu.Registers[13])*256 + int(cpu.Registers[12])
		cpu.Memory[temp] = operation.FirstOperand
	case STRR:
		temp = int(cpu.Registers[13])*256 + int(cpu.Registers[12])
		cpu.Memory[temp] = cpu.Registers[operation.FirstOperand]
	case GET:
		temp = int(cpu.Registers[13])*256 + int(cpu.Registers[12])
		cpu.Registers[operation.FirstOperand] = cpu.Memory[temp]

	case PUSH:
		temp = 65535 - int(cpu.Registers[14])
		cpu.Memory[temp] = cpu.Registers[operation.FirstOperand]
		cpu.Registers[14]++
	case POP:
		temp = 65535 - int(cpu.Registers[14])
		cpu.Registers[operation.FirstOperand] = cpu.Memory[temp]
		cpu.Registers[14]--

	case CALL:
		temp = 65535 - int(cpu.Registers[14])
		cpu.Memory[temp] = byte(cpu.ProgramCounter)
		cpu.Registers[14]++
		temp = 65535 - int(cpu.Registers[14])
		cpu.Memory[temp] = byte(cpu.ProgramCounter / 256)
		cpu.Registers[14]++
		cpu.ProgramCounter = int(operation.FirstOperand)*256 + int(operation.SecondOperand) - 1
	case RET:
		temp = 65535 - int(cpu.Registers[14])
		cpu.ProgramCounter = int(cpu.Memory[temp]) * 256
		cpu.Registers[14]--
		temp = 65535 - int(cpu.Registers[14])
		cpu.ProgramCounter += int(cpu.Memory[temp])
		cpu.Registers[14]--
	}
}

func (cpu *CPU) setOverflow(value bool) {
	if value {
		cpu.Registers[15] = 1
	} else {
		cpu.Registers[15] = 0
	}
}

func (cpu *CPU) getOverflow() bool {
	return cpu.Registers[15] == 1
}

func (cpu *CPU) printState(operation Operation) {
	fmt.Printf("%d OP: %v \n", cpu.ProgramCounter, operation)
}
