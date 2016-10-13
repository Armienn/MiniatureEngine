package machine

//OperationType represents the type of operation/command
type OperationType byte

//Operation represents the entirety of af command
type Operation struct {
	Type          OperationType
	FirstOperand  byte
	SecondOperand byte
}

const (
	//NONE
	NONE OperationType = iota
	//PLUS     R1 123
	ADD
	//PLUS     R1 R2
	ADDR
	//MINUS    R1 123
	SUB
	//MINUS    R1 R2
	SUBR
	//SÆT      R1 123
	SET
	//SÆT      R1 R2
	SETR
	//SÆNK     R1
	DEC
	//ØG       R1
	INC
	//SPRING   tag
	JMP
	//HVIS     R1
	IF
	//HVISIKKE R1
	IFN
	//HVISENS  R1 123
	IFEQ
	//HVISENS  R1 R2
	IFEQR
	//HVISOVER
	IFOV
	//VIS      123
	SHOW
	//VIS      R1
	SHOWR
	//VISTEGN  123
	RUNE
	//VISTEGN  R1
	RUNER
	//FÅTEGN   R1
	GRUNE
	//GEM      123
	STR
	//GEM      R1
	STRR
	//HENT     R1
	GET
	//SKUB		 R1
	PUSH
	//TRÆK 		 R1
	POP
	//OMVEJ    tag
	CALL
	//RETUR
	RET
)
