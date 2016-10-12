package compiler

import (
	"bufio"
	"os"
	"strings"
)

type line struct {
	commandType byte
	command     byte
	op1         string
	op2         string
}

const (
	noCommand = iota
	alias
	lineTag
	command
)

func Compile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := make([]line, 0, 256)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, parseLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseLine(input string) (parsedLine line) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		parsedLine.commandType = noCommand
		return
	}
	parsedLine.commandType = parseLineType(fields[0])
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
