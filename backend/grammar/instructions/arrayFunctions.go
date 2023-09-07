package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
)

type arrayFunctions struct {
	Identifier string
	Id         string
	Expression interface{}
	Line       int
	Column     int
}

func NewArrayFunctions(Identifier string, Id string, Expression interface{}, line int, column int) *arrayFunctions {
	return &arrayFunctions{Identifier: Identifier, Id: Id, Expression: Expression, Line: line, Column: column}
}

func (n arrayFunctions) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	variable := table.GetVariable(n.Identifier)
	if variable.Value == nil {
		err := Error.NewException("Variable no encontrada", n.Identifier, n.Line, n.Column)

		return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
	}
	if n.Id == "append" {
		value := n.Expression.(Abstract.Expression).GetValue(table, ast)

		if variable.Type == Enviorement.ARRAY {
			firsP := variable.Value.([]Enviorement.ReturnSymbol)[0]
			if firsP.Type != value.Type {
				err := Error.NewException("No se puede agregar el valor", n.Identifier, n.Line, n.Column)
				return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
			}
			variable.Value = append(variable.Value.([]Enviorement.ReturnSymbol), value)
			table.UpdateVariable(n.Identifier, variable)
		}

	} else if n.Id == "removeLast" {
		if variable.Type == Enviorement.ARRAY {
			variable.Value = variable.Value.([]Enviorement.ReturnSymbol)[:len(variable.Value.([]Enviorement.ReturnSymbol))-1]
			table.UpdateVariable(n.Identifier, variable)
		} else {
			err := Error.NewException("No se puede remover el valor", n.Identifier, n.Line, n.Column)
			return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
		}
	} else if n.Id == "remove" {
		index := n.Expression.(Abstract.Expression).GetValue(table, ast)
		if variable.Type == Enviorement.ARRAY {
			if index.Type != Enviorement.INT {
				err := Error.NewException("No se puede remover el valor", n.Identifier, n.Line, n.Column)
				return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
			}
			if index.Value.(int) < 0 || index.Value.(int) > len(variable.Value.([]Enviorement.ReturnSymbol))-1 {
				err := Error.NewException("No se puede remover el valor", n.Identifier, n.Line, n.Column)
				return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
			}
			variable.Value = append(variable.Value.([]Enviorement.ReturnSymbol)[:index.Value.(int)], variable.Value.([]Enviorement.ReturnSymbol)[index.Value.(int)+1:]...)
			table.UpdateVariable(n.Identifier, variable)
		}
	}
	return nil
}
