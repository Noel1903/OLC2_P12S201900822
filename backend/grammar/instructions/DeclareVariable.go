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

type DeclareVariable struct {
	Identifier string
	TypeD      Enviorement.TypeData
	Value      Abstract.Expression
	Line       int
	Column     int
}

func NewDeclareWithValue(identifier string, typeVar Enviorement.TypeData, value Abstract.Expression, line int, column int) *DeclareVariable {
	return &DeclareVariable{
		Identifier: identifier,
		TypeD:      typeVar,
		Value:      value,
		Line:       line,
		Column:     column,
	}
}

func NewDeclareWithoutValue(identifier string, typeD Enviorement.TypeData, value Abstract.Expression, line int, column int) *DeclareVariable {
	return &DeclareVariable{
		Identifier: identifier,
		TypeD:      typeD,
		Value:      value,
		Line:       line,
		Column:     column,
	}
}

func (d *DeclareVariable) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	var variable symbol.ReturnSymbol
	variable = table.GetVariableThis(d.Identifier)
	genAuxiliar := Generator.NewGenerator()
	generator := genAuxiliar.GetInstance()
	var tempPos string

	if variable.Value == nil {
		value := d.Value.GetValue(table, ast)
		if d.TypeD == symbol.UNDEFINED {
			d.TypeD = value.Type
		}
		if value.Type != d.TypeD && value.Type != symbol.NIL {
			//fmt.Println("Error, los tipos no coinciden")
			err := Errors.NewException("Error, los tipos no coinciden", table.GetName(), d.Line, d.Column)
			return symbol.ReturnSymbol{
				Type:  symbol.ERROR,
				Value: err,
			}
		}
		result := table.SetVariable(d.Identifier, value, true, d.Line, d.Column, false)
		tempPos = strconv.Itoa(result.GetPos())
		if !result.GetIsGlobal() {
			tempPos = generator.AddTemporal()
			generator.AddExpression("P", strconv.Itoa(result.GetPos()), "+", tempPos)
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
		ast.UpdateSymbolTable("<tr><td>" + d.Identifier + "</td><td>Variable</td><td>" + strconv.Itoa(int(d.TypeD)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(d.Line) + "</td><td>" + strconv.Itoa(d.Column) + "</td></tr>")
	} else {
		err := Errors.NewException("La variable ya ha sido declarada", table.GetName(), d.Line, d.Column)

		return symbol.ReturnSymbol{
			Type:  symbol.ERROR,
			Value: err,
		}
	}

	return nil

}
