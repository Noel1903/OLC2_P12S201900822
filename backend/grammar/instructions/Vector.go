package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
	"strconv"
)

type Vector struct {
	Id     string
	Type   Enviorement.TypeData
	Value  []interface{}
	Line   int
	Column int
}

func NewVector(id string, typeD Enviorement.TypeData, value []interface{}, line int, column int) *Vector {
	return &Vector{
		Id:     id,
		Type:   typeD,
		Value:  value,
		Line:   line,
		Column: column,
	}
}

func (v *Vector) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	variable := table.GetVariableThis(v.Id)
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if variable.Value != nil {
		err := Error.NewException("Variable ya declarada", v.Id, v.Line, v.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}

	var vector_value = []Enviorement.ReturnSymbol{}

	single_value := v.Value
	if single_value == nil {
		vector_result := Enviorement.ReturnSymbol{
			Type:  Enviorement.ARRAY,
			Value: vector_value,
		}
		table.SetVariable(v.Id, vector_result, true, v.Line, v.Column, false)

	} else {
		value := Enviorement.ReturnSymbol{
			Type:  Enviorement.ARRAY,
			Value: single_value,
		}
		var arr = []string{}
		for _, value := range single_value {
			value_exp := value.(Abstract.Expression).GetValue(table, ast)
			if value_exp.Type != v.Type {
				err := Error.NewException("Tipo de dato incorrecto", v.Id, v.Line, v.Column)
				return Enviorement.ReturnSymbol{
					Type:  Enviorement.ERROR,
					Value: err,
				}
			}
			arr = append(arr, value_exp.Value.(string))
		}

		result := table.SetArray(v.Id, value, v.Type, true, v.Line, v.Column, false)
		var tempPos string
		tempPos = strconv.Itoa(result.GetPos())
		if !result.GetIsGlobal() {
			tempPos = generator.AddTemporal()
			generator.AddExpression("P", strconv.Itoa(result.GetPos()), "+", tempPos)
		}
		temporalInitial := generator.AddTemporal()
		generator.AddAssign(temporalInitial, "H")
		generator.SetHeap("H", "1")
		generator.AddExpression("H", "1", "+", "H")
		generator.SetHeap("H", "0")
		generator.AddExpression("H", "1", "+", "H")
		generator.SetHeap("H", strconv.Itoa(len(single_value)))
		generator.AddExpression("H", "1", "+", "H")
		for _, value := range arr {

			generator.SetHeap("H", value)
			generator.AddExpression("H", "1", "+", "H")
		}
		generator.SetStack(tempPos, temporalInitial)

		/*for _, value := range single_value {
			value_exp := value.(Abstract.Expression).GetValue(table, ast)
			if value_exp.Type != v.Type {
				err := Error.NewException("Tipo de dato incorrecto", v.Id, v.Line, v.Column)
				return Enviorement.ReturnSymbol{
					Type:  Enviorement.ERROR,
					Value: err,
				}
			}
			vector_value = append(vector_value, value_exp)
		}
		vector_result := Enviorement.ReturnSymbol{
			Type:  Enviorement.ARRAY,
			Value: vector_value,
		}
		table.SetVariable(v.Id, vector_result, true, v.Line, v.Column, false)*/
		ast.UpdateSymbolTable("<tr><td>" + v.Id + "</td><td>Vector</td><td>" + strconv.Itoa(int(v.Type)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(v.Line) + "</td><td>" + strconv.Itoa(v.Column) + "</td></tr>")

	}

	return nil
}
