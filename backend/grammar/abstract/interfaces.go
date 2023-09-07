package abstract

import (
	table "grammar/symbol"
)

type abstract struct {
	Line   int
	Column int
}

type Expression interface {
	GetValue(env table.SymbolTable, ast *table.AST) table.ReturnSymbol
}

type Instruction interface {
	Execute(env table.SymbolTable, ast *table.AST) interface{}
}
