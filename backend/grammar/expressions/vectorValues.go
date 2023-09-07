package expressions

import (
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
)

type vectorValues struct {
	Identifier string
	Id         string
	Line       int
	Column     int
}

func NewVectorValues(Identifier string, Id string, line int, column int) *vectorValues {
	return &vectorValues{Identifier: Identifier, Id: Id, Line: line, Column: column}
}

func (n vectorValues) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	variable := table.GetVariable(n.Identifier)
	if variable.Value == nil {
		err := Errors.NewException("La variable no ha sido declarada", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	if n.Id == "IsEmpty" {
		if variable.Type != Enviorement.ARRAY {
			err := Errors.NewException("La variable no es un vector", table.GetName(), n.Line, n.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		size := len(variable.Value.([]Enviorement.ReturnSymbol))
		return Enviorement.ReturnSymbol{Type: Enviorement.BOOL, Value: size == 0}

	} else if n.Id == "count" {
		if variable.Type != Enviorement.ARRAY {
			err := Errors.NewException("La variable no es un vector", table.GetName(), n.Line, n.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		size := len(variable.Value.([]Enviorement.ReturnSymbol))
		return Enviorement.ReturnSymbol{Type: Enviorement.INT, Value: size}
	}
	return Enviorement.ReturnSymbol{}
}
