package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
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
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if value.Value == nil {
		fmt.Println("Entra aqui de no existe")
		err := Error.NewException("La funcion no existe", table.GetName(), c.Line, c.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	//fmt.Println(&table, "Tabla que viene del ambito anterior")
	function := value.Value
	paramsList := function.(*DeclareFunction).GetParamsIn()
	//instructionsList := function.(*DeclareFunction).GetSentences()

	for _, param := range c.execParams {
		p := param.(*paramCall)
		c.ParamsIn = append(c.ParamsIn, p.Internal)
		c.ParamsEx = append(c.ParamsEx, p.External)
	}
	// := Enviorement.NewEnviorement(c.Id, &table)
	//fmt.Println(newEnviorement)
	/*temp := generator.AddTemporal()
	generator.AddExpression("P", strconv.Itoa(table.GetSize()), "+", temp)*/
	cont := 0
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
		var value Enviorement.ReturnSymbol
		//fmt.Println(param, reflect.TypeOf(param))

		value = paramT.(Abstract.Expression).GetValue(table, ast)

		if reflect.TypeOf(param) == reflect.TypeOf(&DeclareVariable{}) {
			//variable := table.GetVar(param.(*DeclareVariable).Identifier)

			table.UpdateVariable(param.(*DeclareVariable).Identifier, value)
			/*generator.SetStack(temp, value.Value.(string))
			cont++
			if cont != len(paramsList) {
				generator.AddExpression(temp, "1", "+", temp)
			}*/
			temp := generator.AddTemporal()
			generator.AddExpression("P", strconv.Itoa(table.GetSize()+cont), "+", temp)
			generator.SetStack(temp, value.Value.(string))
			cont++

		}
		//fmt.Println(value, " Valor de parametro")
		//fmt.Println(valueParam, " Valor de parametro")

		//fmt.Println(valueParam.Value.Value, " Valor de parametro")

	}

	generator.AddComment("Llamada a funcion")
	generator.AddEnv(strconv.Itoa(table.GetSize()))
	generator.AddCall(c.Id)
	temporal01 := generator.AddTemporal()
	generator.AddComment("Obtencion de retorno")
	generator.GetStack("P", temporal01)
	generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	return Enviorement.ReturnSymbol{Type: value.Type, Value: temporal01}

	/**/

	//return Enviorement.ReturnSymbol{Type: Enviorement.NIL, Value: nil}
}
