package expressions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
)

type getPosVector struct {
	Identifier string
	Expression interface{}
	Line       int
	Column     int
}

func NewGetPosVector(Identifier string, Expression interface{}, line int, column int) *getPosVector {
	return &getPosVector{Identifier: Identifier, Expression: Expression, Line: line, Column: column}
}

func (n getPosVector) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	variable := table.GetVariable(n.Identifier)
	if variable.Value == nil {
		err := Errors.NewException("La variable no ha sido declarada", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	if variable.Type != Enviorement.ARRAY {
		err := Errors.NewException("La variable no es un vector", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	index := n.Expression.(Abstract.Expression).GetValue(table, ast)
	if index.Type != Enviorement.INT {
		err := Errors.NewException("El indice debe ser de tipo entero", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}

	if index.Value.(int) < 0 || index.Value.(int) > len(variable.Value.([]Enviorement.ReturnSymbol))-1 {
		err := Errors.NewException("El indice esta fuera de rango", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	return variable.Value.([]Enviorement.ReturnSymbol)[index.Value.(int)]
}
