package compiler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Armienn/MiniatureEngine/machine"
)

//Line represents a line of assembler code
type Line struct {
	CommandType byte
	Command     machine.OperationType
	Op1         string
	Op2         string
}

const (
	//NoCommand is the type of a line which has no meaning
	NoCommand = iota
	//Alias is a line which defines a string alias for some value
	Alias
	//LineTag is a line which defines a string which is an alias for the current program line
	LineTag
	//Command is a line which represents an actual program command
	Command
)

//ReadAndParse reads the content of a file and constructs a list of lines with meaningful commands
func ReadAndParse(file *os.File) []Line {
	lines := make([]Line, 0, 256)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := parseLine(scanner.Text())
		if line.CommandType != NoCommand {
			lines = append(lines, line)
		}
	}
	return lines
}

func parseLine(input string) (line Line) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		line.CommandType = NoCommand
		return
	}
	line = parseFields(fields)
	return
}

func parseFields(fields []string) (line Line) {
	defer func() {
		if r := recover(); r != nil {
			line.CommandType = NoCommand
			return
		}
		if line.CommandType == NoCommand {
			fmt.Printf("Couldn't parse %v \n", fields)
		}
	}()

	line.CommandType = parseLineType(fields[0])
	if line.CommandType == Command {
		parseCommand(&line, fields[0])
	}
	parseOperators(&line, fields)
	return
}

func parseLineType(field string) byte {
	switch field {
	case "$":
		return Alias
	case ":":
		return LineTag
	default:
		return Command
	}
}

func parseOperators(line *Line, fields []string) {
	nextField := 1
	if len(fields) > nextField {
		line.Op1 = fields[nextField]
		nextField++
		if len(fields) > nextField {
			line.Op2 = fields[nextField]
		}
	}
}

func parseCommand(line *Line, command string) {
	switch command {
	case "NONE":
		line.Command = machine.NONE
	case "PLUS", "ADD":
		line.Command = machine.ADD
	case "PLUSR", "ADDR":
		line.Command = machine.ADDR
	case "MINUS", "SUB":
		line.Command = machine.SUB
	case "MINUSR", "SUBR":
		line.Command = machine.SUBR
	case "SÆT", "SET":
		line.Command = machine.SET
	case "SÆTR", "SETR":
		line.Command = machine.SETR
	case "SÆNK", "DEC":
		line.Command = machine.DEC
	case "ØG", "INC":
		line.Command = machine.INC
	case "SPRING", "JMP":
		line.Command = machine.JMP
	case "HVIS", "IF":
		line.Command = machine.IF
	case "HVISIKKE", "IFN":
		line.Command = machine.IFN
	case "HVISENS", "IFEQ":
		line.Command = machine.IFEQ
	case "HVISENSR", "IFEQR":
		line.Command = machine.IFEQR
	case "HVISOVER", "IFOV":
		line.Command = machine.IFOV
	case "VIS", "SHOW":
		line.Command = machine.SHOW
	case "VISR", "SHOWR":
		line.Command = machine.SHOWR
	case "VISTEGN", "RUNE":
		line.Command = machine.RUNE
	case "VISTEGNR", "RUNER":
		line.Command = machine.RUNER
	case "FÅTEGN", "GRUNE":
		line.Command = machine.GRUNE
	case "GEM", "STR":
		line.Command = machine.STR
	case "GEMR", "STRR":
		line.Command = machine.STRR
	case "HENT", "GET":
		line.Command = machine.GET
	case "SKUB", "PUSH":
		line.Command = machine.PUSH
	case "TRÆK", "POP":
		line.Command = machine.POP
	case "OMVEJ", "CALL":
		line.Command = machine.CALL
	case "RETUR", "RET":
		line.Command = machine.RET
	default:
		line.CommandType = NoCommand
	}
}
