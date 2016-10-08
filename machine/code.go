package machine

type OperationType int

type Operation struct {
	Type          OperationType
	FirstOperand  byte
	SecondOperand byte
}

const (
	NONE     OperationType = iota // NOOPERATION
	PLUS                          // PLUS     R1 123
	PLUSR                         // PLUS     R1 R2
	MINUS                         // MINUS    R1 123
	MINUSR                        // MINUS    R1 R2
	SÆT                           // SÆT      R1 123
	SÆTR                          // SÆT      R1 R2
	SÆNK                          // SÆNK     R1
	ØG                            // ØG       R1
	SPRING                        // SPRING   tag
	HVIS                          // HVIS     R1
	HVISIKKE                      // HVISIKKE R1
	HVISENS                       // HVISENS  R1 123
	HVISENSR                      // HVISENS  R1 R2
	VIS                           // VIS      123
	VISR                          // VIS      R1
	VISTEGN                       // VISTEGN  123
	VISTEGNR                      // VISTEGN  R1
	FÅTEGN                        // FÅTEGN   R1
	GEM                           // GEM      123
	GEMR                          // GEM      R1
	HENT                          // HENT     R1
	SKUB                          // SKUB		 R1
	TRÆK                          // TRÆK 		 R1
	OMVEJ                         // OMVEJ    tag
	RETUR                         // RETUR
)
