package expressions

import (
	"fmt"
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
)

type callFunctionExp struct {
	callFunc interface{}
}

func NewCallFunctionExp(callFunc interface{}) *callFunctionExp {
	return &callFunctionExp{
		callFunc: callFunc,
	}
}

func (c *callFunctionExp) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	value := c.callFunc.(Abstract.Instruction).Execute(table, ast).(Enviorement.ReturnSymbol)

	if value.Value == nil {
		err := Errors.NewException("No se puede llamar a la funcion", table.GetName(), 0, 0)
		fmt.Println("Entra aqui")
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	} else {
		return value
	}
}
