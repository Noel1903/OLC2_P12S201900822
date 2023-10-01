package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
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
	fmt.Println(&table, "Tabla que viene del ambito anterior")
	function := value.Value
	paramsList := function.(*DeclareFunction).GetParamsIn()
	//instructionsList := function.(*DeclareFunction).GetSentences()

	for _, param := range c.execParams {
		p := param.(*paramCall)
		c.ParamsIn = append(c.ParamsIn, p.Internal)
		c.ParamsEx = append(c.ParamsEx, p.External)
	}
	newEnviorement := Enviorement.NewEnviorement(c.Id, &table)
	//fmt.Println(newEnviorement)
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
		value := paramT.(Abstract.Expression).GetValue(table, ast)

		valueParam := param.(Enviorement.Symbol)
		position := valueParam.GetPos()
		fmt.Println(valueParam.Value.Value, " Valor de parametro")
		fmt.Println(value.Value, "valor a asignar")
		//generator.AddAssign(valueParam.Value.Value.(string), value.Value.(string))
		generator.SetStack(strconv.Itoa(position), value.Value.(string))

		/*if reflect.TypeOf(param) == reflect.TypeOf(&DeclareVariable{}) {

			res := newEnviorement.GetVar(param.(*DeclareVariable).Identifier)
			fmt.Println(newEnviorement.GetVariable(param.(*DeclareVariable).Identifier))
			position := res.GetPos()
			generator.AddComment("Asignacion de parametro")
			generator.AddAssign(res.Value.Value.(string), value.Value.(string))
			generator.SetStack(strconv.Itoa(position), value.Value.(string))

			table.SetVariable(param.(*DeclareVariable).Identifier, value, true, c.Line, c.Column, false)

			ast.UpdateSymbolTable("<tr><td>" + param.(*DeclareVariable).Identifier + "</td><td>Funcion</td><td>" + strconv.Itoa(int(param.(*DeclareVariable).TypeD)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(param.(*DeclareVariable).Line) + "</td><td>" + strconv.Itoa(param.(*DeclareVariable).Column) + "</td></tr>")
		} else if reflect.TypeOf(param) == reflect.TypeOf(&Vector{}) {

			table.SetVariable(param.(*Vector).Id, value, true, c.Line, c.Column, false)
			ast.UpdateSymbolTable("<tr><td>" + param.(*Vector).Id + "</td><td>Funcion</td><td>" + strconv.Itoa(int(param.(*Vector).Type)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(param.(*Vector).Line) + "</td><td>" + strconv.Itoa(param.(*Vector).Column) + "</td></tr>")
		}*/

	}

	generator.AddComment("Llamada a funcion")
	generator.AddEnv(strconv.Itoa(newEnviorement.GetSize()))
	generator.AddCall(c.Id)
	temporal01 := generator.AddTemporal()
	generator.AddComment("Obtencion de retorno")
	generator.GetStack("P", temporal01)
	generator.ReturnEnv(strconv.Itoa(newEnviorement.GetSize()))
	return Enviorement.ReturnSymbol{Type: value.Type, Value: temporal01}

	/**/

	//return Enviorement.ReturnSymbol{Type: Enviorement.NIL, Value: nil}
}
