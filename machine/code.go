package machine

type OperationType int

type Operation struct {
	Type          OperationType
	FirstOperand  byte
	SecondOperand byte
}

const (
	NONE  OperationType = iota // NOOPERATION
	ADD                        // PLUS     R1 123
	ADDR                       // PLUS     R1 R2
	SUB                        // MINUS    R1 123
	SUBR                       // MINUS    R1 R2
	SET                        // SÆT      R1 123
	SETR                       // SÆT      R1 R2
	DEC                        // SÆNK     R1
	INC                        // ØG       R1
	JMP                        // SPRING   tag
	IF                         // HVIS     R1
	IFN                        // HVISIKKE R1
	IFEQ                       // HVISENS  R1 123
	IFEQR                      // HVISENS  R1 R2
	IFOV                       // HVISOVER
	SHOW                       // VIS      123
	SHOWR                      // VIS      R1
	RUNE                       // VISTEGN  123
	RUNER                      // VISTEGN  R1
	GRUNE                      // FÅTEGN   R1
	STR                        // GEM      123
	STRR                       // GEM      R1
	GET                        // HENT     R1
	PUSH                       // SKUB		 R1
	POP                        // TRÆK 		 R1
	CALL                       // OMVEJ    tag
	RET                        // RETUR
)
