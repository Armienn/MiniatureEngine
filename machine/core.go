package machine

import "fmt"

type Cpu struct {
	Registers [16]byte
	Memory    [65536]byte
}

func (*Cpu) RunProgram(program []Operation) {
	for i, operation := range program {
		fmt.Println("line ", i, operation)
	}
}
