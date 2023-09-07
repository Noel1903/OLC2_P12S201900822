package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	"grammar/symbol"
	Enviorement "grammar/symbol"
	"strconv"
)

type DeclareLet struct {
	identifier string
	typeD      Enviorement.TypeData
	value      Abstract.Expression
	Line       int
	Column     int
}

func NewLet(identifier string, typeVar Enviorement.TypeData, value Abstract.Expression, line int, column int) *DeclareLet {
	return &DeclareLet{
		identifier: identifier,
		typeD:      typeVar,
		value:      value,
		Line:       line,
		Column:     column,
	}
}

func (l *DeclareLet) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {

	variable := table.GetVariable(l.identifier)

	if variable.Value == nil {
		value := l.value.GetValue(table, ast)
		if l.typeD == symbol.UNDEFINED {
			l.typeD = value.Type
		}
		if value.Type != l.typeD && value.Type != symbol.NIL {
			err := Error.NewException("Error: No se puede asignar el valor a la variable", table.GetName(), l.Line, l.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		table.SetVariable(l.identifier, value, false, l.Line, l.Column)
		ast.UpdateSymbolTable("<tr><td>" + l.identifier + "</td><td>Constante</td><td>" + strconv.Itoa(int(l.typeD)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(l.Line) + "</td><td>" + strconv.Itoa(l.Column) + "</td></tr>")
	} else {
		err := Error.NewException("Error: La variable ya existe", table.GetName(), l.Line, l.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}

	return nil

}
