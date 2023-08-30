package instructions

import (
	Enviorement "grammar/symbol"
)

type Continue struct {
	identifier string
}

func NewContinue(identifier string) *Continue {
	return &Continue{
		identifier: identifier,
	}
}

func (b *Continue) Execute(table Enviorement.SymbolTable) interface{} {
	return Enviorement.ReturnSymbol{
		Type:  Enviorement.UNDEFINED,
		Value: b.identifier,
	}
}
