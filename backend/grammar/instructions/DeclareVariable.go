package instructions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	"grammar/symbol"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
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
		tempPos := result.GetPos()
		if !result.GetIsGlobal() {
			tempPos := generator.AddTemporal()
			generator.AddExpression(tempPos, "P", strconv.Itoa(result.GetPos()), "+")
		}
		if value.Type == symbol.INT {
			generator.SetStack(strconv.Itoa(tempPos), value.GetValue().(string))
		} else if value.Type == symbol.FLOAT {
			generator.SetStack(strconv.Itoa(tempPos), value.GetValue().(string))
		} else if value.Type == symbol.STRING {
			generator.SetStack(strconv.Itoa(tempPos), value.GetValue().(string))
		} else if value.Type == symbol.BOOL {
			generator.SetStack(strconv.Itoa(tempPos), value.GetValue().(string))
		} else if value.Type == symbol.CHAR {
			generator.SetStack(strconv.Itoa(tempPos), value.GetValue().(string))
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
