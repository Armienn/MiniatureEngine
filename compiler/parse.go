package compiler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Armienn/MiniatureEngine/machine"
)

type line struct {
	commandType byte
	command     machine.OperationType
	op1         string
	op2         string
}

const (
	noCommand = iota
	alias
	lineTag
	command
)

func readAndParse(file *os.File) []line {
	lines := make([]line, 0, 256)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := parseLine(scanner.Text())
		if line.commandType != noCommand {
			lines = append(lines, line)
		}
	}
	return lines
}

func parseLine(input string) (parsedLine line) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		parsedLine.commandType = noCommand
		return
	}
	parsedLine = parseFields(fields)
	return
}

func parseFields(fields []string) (parsedLine line) {
	defer func() {
		if r := recover(); r != nil {
			parsedLine.commandType = noCommand
			return
		}
		if parsedLine.commandType == noCommand {
			fmt.Printf("Couldn't parse %v \n", fields)
		}
	}()

	parsedLine.commandType = parseLineType(fields[0])
	if parsedLine.commandType == command {
		parseCommand(&parsedLine, fields[1])
	}
	parseOperators(&parsedLine, fields)
	return
}

func parseLineType(field string) byte {
	switch field {
	case "$":
		return alias
	case ":":
		return lineTag
	default:
		return command
	}
}

func parseOperators(parsedLine *line, fields []string) {
	nextField := 1
	if parsedLine.commandType == command {
		nextField = 2
	}
	if len(fields) > nextField {
		parsedLine.op1 = fields[nextField]
		nextField++
		if len(fields) > nextField {
			parsedLine.op2 = fields[nextField]
		}
	}
}

func parseCommand(parsedLine *line, command string) {
	switch command {
	case "NONE":
		parsedLine.command = machine.NONE
	case "PLUS", "ADD":
		parsedLine.command = machine.ADD
	case "PLUSR", "ADDR":
		parsedLine.command = machine.ADDR
	case "MINUS", "SUB":
		parsedLine.command = machine.SUB
	case "MINUSR", "SUBR":
		parsedLine.command = machine.SUBR
	case "SÆT", "SET":
		parsedLine.command = machine.SET
	case "SÆTR", "SETR":
		parsedLine.command = machine.SETR
	case "SÆNK", "DEC":
		parsedLine.command = machine.DEC
	case "ØG", "INC":
		parsedLine.command = machine.INC
	case "SPRING", "JMP":
		parsedLine.command = machine.JMP
	case "HVIS", "IF":
		parsedLine.command = machine.IF
	case "HVISIKKE", "IFN":
		parsedLine.command = machine.IFN
	case "HVISENS", "IFEQ":
		parsedLine.command = machine.IFEQ
	case "HVISENSR", "IFEQR":
		parsedLine.command = machine.IFEQR
	case "HVISOVER", "IFOV":
		parsedLine.command = machine.IFOV
	case "VIS", "SHOW":
		parsedLine.command = machine.SHOW
	case "VISR", "SHOWR":
		parsedLine.command = machine.SHOWR
	case "VISTEGN", "RUNE":
		parsedLine.command = machine.RUNE
	case "VISTEGNR", "RUNER":
		parsedLine.command = machine.RUNER
	case "FÅTEGN", "GRUNE":
		parsedLine.command = machine.GRUNE
	case "GEM", "STR":
		parsedLine.command = machine.STR
	case "GEMR", "STRR":
		parsedLine.command = machine.STRR
	case "HENT", "GET":
		parsedLine.command = machine.GET
	case "SKUB", "PUSH":
		parsedLine.command = machine.PUSH
	case "TRÆK", "POP":
		parsedLine.command = machine.POP
	case "OMVEJ", "CALL":
		parsedLine.command = machine.CALL
	case "RETUR", "RET":
		parsedLine.command = machine.RET
	default:
		parsedLine.commandType = noCommand
	}
}
