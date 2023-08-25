package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	"reflect"
)

type While struct {
	condition Abstract.Expression
	codeWhile []interface{}
}

func NewWhile(condition Abstract.Expression, codeWhile []interface{}) *While {
	return &While{
		condition: condition,
		codeWhile: codeWhile,
	}
}

func (w *While) Execute(table Enviorement.SymbolTable) interface{} {
	condition := w.condition.GetValue(table)
	if condition.Type != Enviorement.BOOL {
		fmt.Println("Error: La condicion no es booleana")
		return nil
	}

	newEnviorement := Enviorement.NewEnviorement("while", &table)
	var instruc string
	for condition.Value.(bool) {
		for _, instr := range w.codeWhile {
			result := instr.(Abstract.Instruction).Execute(newEnviorement) //ejecuta instrucciones
			if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
				if result.(Enviorement.ReturnSymbol).Value == "break" {
					instruc = "break"
					break
				}
			}
		}
		if instruc == "break" {
			break
		}
		condition = w.condition.GetValue(table)
		if condition.Type != Enviorement.BOOL {
			fmt.Println("Error: La condicion no es booleana")
		}
	}

	return nil
}
