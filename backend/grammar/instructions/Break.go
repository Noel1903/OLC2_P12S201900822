package instructions

import (
	Enviorement "grammar/symbol"
)

type Break struct {
	identifier string
}

func NewBreak(identifier string) *Break {
	return &Break{
		identifier: identifier,
	}
}

func (b *Break) Execute(table Enviorement.SymbolTable) interface{} {
	return Enviorement.ReturnSymbol{
		Type:  Enviorement.UNDEFINED,
		Value: b.identifier,
	}
}