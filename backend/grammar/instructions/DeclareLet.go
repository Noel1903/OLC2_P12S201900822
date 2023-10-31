package instructions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	"grammar/symbol"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
	"reflect"
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

	var variable symbol.ReturnSymbol
	variable = table.GetVariableThis(l.identifier)
	genAuxiliar := Generator.NewGenerator()
	generator := genAuxiliar.GetInstance()
	var tempPos string
	if variable.Value == nil {
		value := l.value.GetValue(table, ast)
		if l.typeD == symbol.UNDEFINED {
			l.typeD = value.Type
		}
		if value.Type != l.typeD && value.Type != symbol.NIL {
			//fmt.Println("Error, los tipos no coinciden")
			err := Errors.NewException("Error, los tipos no coinciden", table.GetName(), l.Line, l.Column)
			return symbol.ReturnSymbol{
				Type:  symbol.ERROR,
				Value: err,
			}
		}
		result := table.SetVariable(l.identifier, value, false, l.Line, l.Column, false)
		tempPos = strconv.Itoa(result.GetPos())
		if !result.GetIsGlobal() {
			tempPos = generator.AddTemporal()
			generator.AddExpression("P", strconv.Itoa(result.GetPos()), "+", tempPos)
		} else {

		}
		if value.Type == symbol.INT {
			generator.SetStack(tempPos, value.GetValue().(string))
		} else if value.Type == symbol.FLOAT {
			generator.SetStack(tempPos, value.GetValue().(string))
		} else if value.Type == symbol.STRING {
			generator.SetStack(tempPos, value.GetValue().(string))
		} else if value.Type == symbol.BOOL {
			newLabel := generator.AddLabel()

			for _, labels := range value.LabelTrue {
				if reflect.TypeOf(labels).Kind() != reflect.Slice || reflect.TypeOf(labels).Elem().Kind() == reflect.Interface {
					generator.PutLabel(labels.(string))

				}
			}
			generator.SetStack(tempPos, "1")
			generator.AddGoto(newLabel)
			for _, labels := range value.LabelFalse {
				if reflect.TypeOf(labels).Kind() != reflect.Slice || reflect.TypeOf(labels).Elem().Kind() == reflect.Interface {
					generator.PutLabel(labels.(string))

				}

			}
			generator.SetStack(tempPos, "0")
			generator.PutLabel(newLabel)

		} else if value.Type == symbol.CHAR {
			generator.SetStack(tempPos, value.GetValue().(string))
		} else if value.Type == symbol.NIL {
			generator.SetStack(tempPos, value.GetValue().(string))
		}
		ast.UpdateSymbolTable("<tr><td>" + l.identifier + "</td><td>Constante</td><td>" + strconv.Itoa(int(l.typeD)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(l.Line) + "</td><td>" + strconv.Itoa(l.Column) + "</td></tr>")
	} else {
		err := Errors.NewException("La constante ya ha sido declarada", table.GetName(), l.Line, l.Column)

		return symbol.ReturnSymbol{
			Type:  symbol.ERROR,
			Value: err,
		}
	}

	return nil

}
