package expressions

import (
	Scoope "grammar/symbol"
)

type Native struct {
	//Fila - columna
	Value interface{}
	TypeR Scoope.TypeData
}

func (n Native) GetValue(env Scoope.SymbolTable) Scoope.ReturnSymbol {

	return Scoope.ReturnSymbol{
		Type:  n.TypeR,
		Value: n.Value,
	}
}

func NewNative(value interface{}, typeR Scoope.TypeData) Native {
	s := Native{value, typeR}
	return s
}
