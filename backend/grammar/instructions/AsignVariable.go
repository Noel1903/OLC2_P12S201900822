package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
	"reflect"
	"strconv"
)

type AsignVariable struct {
	Id         string
	Expression Abstract.Expression
	Line       int
	Column     int
}

func NewAsignVariable(Id string, Expression Abstract.Expression, line int, column int) *AsignVariable {
	return &AsignVariable{Id: Id, Expression: Expression, Line: line, Column: column}
}

func (a AsignVariable) Execute(table Enviorement.SymbolTable, tree *Enviorement.AST) interface{} {
	value := table.GetVariable(a.Id)
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if value.Value != nil {
		expresion := a.Expression.GetValue(table, tree)
		if value.Type != expresion.Type {
			err := Error.NewException("No se puede asignar el valor", a.Id, a.Line, a.Column)
			return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
		}
		variable := table.GetVar(a.Id)
		generator.SetStack(strconv.Itoa(variable.GetPos()), expresion.Value.(string))
		res := table.UpdateVariable(a.Id, expresion)
		if reflect.TypeOf(res) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
			if res.(Enviorement.ReturnSymbol).Type == Enviorement.ERROR {
				err := Error.NewException("No se puede asignar el valor", a.Id, a.Line, a.Column)
				return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
			}
		}
	}

	return nil
}
