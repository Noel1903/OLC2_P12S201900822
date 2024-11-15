package abstract

import (
	table "grammar/symbol"
)

type abstract struct {
	Line   int
	Column int
}

type Expression interface {
	GetValue(env table.SymbolTable) table.ReturnSymbol
}

type Instruction interface {
	GetInstruction(env table.SymbolTable) interface{}
}
