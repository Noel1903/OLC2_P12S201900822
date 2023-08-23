package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
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
	}

	newEnviorement := Enviorement.NewEnviorement("while", &table)

	for condition.Value.(bool) {
		for _, instr := range w.codeWhile {
			instr.(Abstract.Instruction).Execute(newEnviorement)
		}
		condition = w.condition.GetValue(table)
		if condition.Type != Enviorement.BOOL {
			fmt.Println("Error: La condicion no es booleana")
		}
	}

	return nil
}
