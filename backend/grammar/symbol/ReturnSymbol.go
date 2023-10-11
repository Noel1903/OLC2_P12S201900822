package symbol

type TypeData int

const (
	INT TypeData = iota
	FLOAT
	STRING
	BOOL
	CHAR
	ARRAY
	VOID
	NIL
	UNDEFINED
	ERROR
)

type ReturnSymbol struct {
	Type       TypeData
	Value      interface{}
	LabelTrue  []interface{}
	LabelFalse []interface{}
}

func (this *ReturnSymbol) GetValue() interface{} {
	return this.Value
}

func (this *ReturnSymbol) GetType() TypeData {
	return this.Type
}

func (this *ReturnSymbol) SetValue(value interface{}) {
	this.Value = value
}

func (this *ReturnSymbol) SetType(Type TypeData) {
	this.Type = Type
}
