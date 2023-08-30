package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	"reflect"
)

type callFunction struct {
	Id         string
	execParams []interface{}
}

func NewCallFunction(id string, execParams []interface{}) *callFunction {
	return &callFunction{
		Id:         id,
		execParams: execParams,
	}
}

func (c *callFunction) Execute(table Enviorement.SymbolTable) interface{} {
	value := table.GetVariable(c.Id)
	if value.Value == nil {
		return nil
	}
	function := value.Value
	paramsList := function.(*DeclareFunction).GetParams()
	instructionsList := function.(*DeclareFunction).GetSentences()
	if len(paramsList) != len(c.execParams) {
		fmt.Println("Error: La cantidad de parametros no coincide")
		return nil
	}

	newEnviorement := Enviorement.NewEnviorement(c.Id, &table)

	for index, param := range paramsList {
		paramT := c.execParams[index]
		/*if paramT.(expressions.Native).TypeR != param.(*DeclareVariable).TypeD {
			fmt.Println("Error: El tipo de dato no coincide")
			return nil
		}*/
		value := paramT.(Abstract.Expression).GetValue(table)
		//param.(DeclareVariable).Value = value.Value

		newEnviorement.SetVariable(param.(*DeclareVariable).Identifier, value, true)

	}

	for _, instruction := range instructionsList {

		result := instruction.(Abstract.Instruction).Execute(newEnviorement)
		if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
			if result.(Enviorement.ReturnSymbol).Value == "break" {
				break
			} else if result.(Enviorement.ReturnSymbol).Value == nil {
				//fmt.Println(result)
				return result
			} else if result.(Enviorement.ReturnSymbol).Value == "continue" {
				break
			} else if result.(Enviorement.ReturnSymbol).Value != nil {
				if result.(Enviorement.ReturnSymbol).Type == function.(*DeclareFunction).GetType() {
					return result
				} else {
					fmt.Println("Error: El tipo de retorno no coincide")
					return nil
				}

			}
		}
	}

	return nil
}
