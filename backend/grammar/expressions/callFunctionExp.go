package expressions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
)

type callFunctionExp struct {
	callFunc interface{}
}

func NewCallFunctionExp(callFunc interface{}) *callFunctionExp {
	return &callFunctionExp{
		callFunc: callFunc,
	}
}

func (c *callFunctionExp) GetValue(table Enviorement.SymbolTable) Enviorement.ReturnSymbol {
	value := c.callFunc.(Abstract.Instruction).Execute(table).(Enviorement.ReturnSymbol)

	if value.Value == nil {
		return Enviorement.ReturnSymbol{}
	} else {
		return value
	}
}
