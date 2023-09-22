package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	"reflect"
	"strconv"
)

type callFunction struct {
	Id         string
	execParams []interface{}
	ParamsEx   []interface{}
	ParamsIn   []interface{}
	Line       int
	Column     int
}

func NewCallFunction(id string, execParams []interface{}, line int, column int) *callFunction {
	return &callFunction{
		Id:         id,
		execParams: execParams,
		Line:       line,
		Column:     column,
	}
}

func (c *callFunction) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	value := table.GetVariable(c.Id)
	if value.Value == nil {
		fmt.Println("Entra aqui de no existe")
		err := Error.NewException("La funcion no existe", table.GetName(), c.Line, c.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	function := value.Value
	paramsList := function.(*DeclareFunction).GetParamsIn()
	instructionsList := function.(*DeclareFunction).GetSentences()

	for _, param := range c.execParams {
		p := param.(*paramCall)
		c.ParamsIn = append(c.ParamsIn, p.Internal)
		c.ParamsEx = append(c.ParamsEx, p.External)
	}

	/*if len(paramsList) != len(c.ParamsIn) {
		fmt.Println("Entra aqui de no coincide cantidad")
		err := Error.NewException("Error: La cantidad de parametros no coincide", table.GetName(), c.Line, c.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}*/

	newEnviorement := Enviorement.NewEnviorement(c.Id, &table)

	for index, param := range paramsList {
		paramE := function.(*DeclareFunction).GetParamsEx()[index]
		paramT := c.ParamsIn[index] //{"type": "native", "value": "hola"}

		if paramE.(string) != c.ParamsEx[index].(string) {
			fmt.Println("Entra aqui de no coincide nombre")
			err := Error.NewException("Error: El nombre de parametro no coincide", table.GetName(), c.Line, c.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		value := paramT.(Abstract.Expression).GetValue(table, ast) //"hola","string"
		//param.(DeclareVariable).Value = value.Value
		if reflect.TypeOf(param) == reflect.TypeOf(&DeclareVariable{}) {
			fmt.Println("Entra aqui de variable")
			newEnviorement.SetVariable(param.(*DeclareVariable).Identifier, value, true, c.Line, c.Column, false)
			ast.UpdateSymbolTable("<tr><td>" + param.(*DeclareVariable).Identifier + "</td><td>Funcion</td><td>" + strconv.Itoa(int(param.(*DeclareVariable).TypeD)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(param.(*DeclareVariable).Line) + "</td><td>" + strconv.Itoa(param.(*DeclareVariable).Column) + "</td></tr>")
		} else if reflect.TypeOf(param) == reflect.TypeOf(&Vector{}) {

			newEnviorement.SetVariable(param.(*Vector).Id, value, true, c.Line, c.Column, false)
			ast.UpdateSymbolTable("<tr><td>" + param.(*Vector).Id + "</td><td>Funcion</td><td>" + strconv.Itoa(int(param.(*Vector).Type)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(param.(*Vector).Line) + "</td><td>" + strconv.Itoa(param.(*Vector).Column) + "</td></tr>")
		}
		//mi tabla actual => saludo = "hola","string","mutable"

	}

	for _, instruction := range instructionsList {

		result := instruction.(Abstract.Instruction).Execute(newEnviorement, ast)

		if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
			if result.(Enviorement.ReturnSymbol).Value == "break" {
				break
			} else if result.(Enviorement.ReturnSymbol).Value == nil {
				//fmt.Println(result)
				return result
			} else if result.(Enviorement.ReturnSymbol).Value == "continue" {
				break
			} else if result.(Enviorement.ReturnSymbol).Value != nil {
				return result.(Enviorement.ReturnSymbol)

			}
		}
	}

	return Enviorement.ReturnSymbol{Type: Enviorement.NIL, Value: nil}
}
