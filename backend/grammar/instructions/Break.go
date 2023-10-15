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

func (b *Break) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	return Enviorement.ReturnSymbol{
		Type:  Enviorement.BREAK,
		Value: b.identifier,
	}
}
