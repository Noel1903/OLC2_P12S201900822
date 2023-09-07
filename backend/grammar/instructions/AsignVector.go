package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
)

type AsignVector struct {
	Id         string
	Position   Abstract.Expression
	Expression Abstract.Expression
	Line       int
	Column     int
}

func NewAsignVector(Id string, Position Abstract.Expression, Expression Abstract.Expression, line int, column int) *AsignVector {
	return &AsignVector{Id: Id, Position: Position, Expression: Expression, Line: line, Column: column}
}

func (a AsignVector) Execute(table Enviorement.SymbolTable, tree *Enviorement.AST) interface{} {
	value := table.GetVariable(a.Id)
	if value.Value != nil {
		expresion := a.Expression.GetValue(table, tree)
		valueArr := value.Value.(Abstract.Expression).GetValue(table, tree)

		if valueArr.Type != expresion.Type {
			fmt.Println("Error no son del mismo tipo")
			err := Error.NewException("Error no son del mismo tipo", table.GetName(), a.Line, a.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		if value.Type != Enviorement.ARRAY {
			err := Error.NewException("Error no es un arreglo", table.GetName(), a.Line, a.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		pos := a.Position.GetValue(table, tree)
		if pos.Type != Enviorement.INT {
			err := Error.NewException("Error no es un entero", table.GetName(), a.Line, a.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		if pos.Value.(int) < 0 || pos.Value.(int) > len(valueArr.Value.([]Enviorement.ReturnSymbol))-1 {
			err := Error.NewException("Error fuera de rango", table.GetName(), a.Line, a.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}

		value.Value.([]Enviorement.ReturnSymbol)[pos.Value.(int)] = expresion
		table.UpdateVariable(a.Id, value)
	}

	return nil
}
