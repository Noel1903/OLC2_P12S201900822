package expressions

import (
	Scoope "grammar/symbol"
)

type Native struct {
	//Fila - columna
	Value  interface{}
	TypeR  Scoope.TypeData
	Line   int
	Column int
}

func (n Native) GetValue(env Scoope.SymbolTable, ast *Scoope.AST) Scoope.ReturnSymbol {

	return Scoope.ReturnSymbol{
		Type:  n.TypeR,
		Value: n.Value,
	}
}

func NewNative(value interface{}, typeR Scoope.TypeData, line int, column int) Native {
	s := Native{value, typeR, line, column}
	return s
}
