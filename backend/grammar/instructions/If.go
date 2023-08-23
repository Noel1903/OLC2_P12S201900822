package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Enviorment "grammar/symbol"
)

type If struct {
	condition Abstract.Expression
	codeIf    []interface{}
	codeElse  []interface{}
}

func NewIf(condition Abstract.Expression, codeIf []interface{}, codeElse []interface{}) *If {
	return &If{
		condition: condition,
		codeIf:    codeIf,
		codeElse:  codeElse,
	}
}

func (i *If) Execute(table Enviorment.SymbolTable) interface{} {
	condition := i.condition.GetValue(table)
	if condition.Type != Enviorment.BOOL {
		fmt.Println("Error: La condicion no es booleana")
		return nil
	}

	newEnviorement := Enviorment.NewEnviorement("if", &table)

	if condition.Value.(bool) {
		for _, instr := range i.codeIf {
			instr.(Abstract.Instruction).Execute(newEnviorement)
		}
	} else {
		for _, instr := range i.codeElse {
			instr.(Abstract.Instruction).Execute(newEnviorement)
		}
	}

	return nil
}
