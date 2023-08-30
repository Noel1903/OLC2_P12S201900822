package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	"reflect"
)

type Guard struct {
	condition Abstract.Expression
	codeGuard []interface{}
}

func NewGuard(condition Abstract.Expression, codeGuard []interface{}) *Guard {
	return &Guard{
		condition: condition,
		codeGuard: codeGuard,
	}
}

func (g *Guard) Execute(table Enviorement.SymbolTable) interface{} {
	condition := g.condition.GetValue(table)
	if condition.Type != Enviorement.BOOL {
		fmt.Println("Error: La condicion no es booleana")
		return nil
	}
	newEnviorement := Enviorement.NewEnviorement("guard", &table)
	if condition.Value.(bool) == false {
		for _, instr := range g.codeGuard {
			result := instr.(Abstract.Instruction).Execute(newEnviorement)
			//fmt.Println("ingresa a guard")
			if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
				if result.(Enviorement.ReturnSymbol).Value == "break" {
					return result
				} else if result.(Enviorement.ReturnSymbol).Value == nil {
					//fmt.Println(result)
					return result
				} else if result.(Enviorement.ReturnSymbol).Value == "continue" {
					return result
				} else if result.(Enviorement.ReturnSymbol).Value != nil {
					return result.(Enviorement.ReturnSymbol)
				}
			}
		}
	}
	return nil
}
