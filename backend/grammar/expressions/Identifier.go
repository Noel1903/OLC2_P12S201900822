package expressions

import (
	symbol "grammar/symbol"
)

type Identifier struct {
	Identifier string
}

func NewIdentifier(id string) Identifier {
	return Identifier{Identifier: id}
}

func (i Identifier) GetValue(table symbol.SymbolTable) symbol.ReturnSymbol {
	variable := table.GetVariable(i.Identifier)
	return symbol.ReturnSymbol{Type: variable.Type, Value: variable.Value}
}
