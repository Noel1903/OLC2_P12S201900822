package symbol

type TypeData int

const (
	INT TypeData = iota
	FLOAT
	STRING
	BOOL
	CHAR
	NIL
	UNDEFINED
)

type ReturnSymbol struct {
	Type  TypeData
	Value interface{}
}
