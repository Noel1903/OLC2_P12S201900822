package expressions

import (
	errors "grammar/exceptions"
	symbol "grammar/symbol"
)

type Identifier struct {
	Identifier string
	Line       int
	Column     int
}

func NewIdentifier(id string, Line int, Column int) Identifier {
	return Identifier{Identifier: id, Line: Line, Column: Column}
}

func (i Identifier) GetValue(table symbol.SymbolTable, ast *symbol.AST) symbol.ReturnSymbol {
	variable := table.GetVariable(i.Identifier)
	if variable.Value == nil {
		newErr := errors.NewException("La variable "+i.Identifier+" no existe", table.GetName(), i.Line, i.Column)
		return symbol.ReturnSymbol{Type: symbol.ERROR, Value: newErr}
	}
	return symbol.ReturnSymbol{Type: variable.Type, Value: variable.Value}
}
