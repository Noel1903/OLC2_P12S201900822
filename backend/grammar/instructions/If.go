package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	Enviorment "grammar/symbol"
	"reflect"
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
			result := instr.(Abstract.Instruction).Execute(newEnviorement)
			if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
				if result.(Enviorement.ReturnSymbol).Value == "break" {
					return result
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
	} else {
		for _, instr := range i.codeElse {
			result := instr.(Abstract.Instruction).Execute(newEnviorement)
			if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
				if result.(Enviorement.ReturnSymbol).Value == "break" {
					return result
				} else if result.(Enviorement.ReturnSymbol).Value == nil {
					return nil
				} else if result.(Enviorement.ReturnSymbol).Value == "continue" {
					break
				} else if result.(Enviorement.ReturnSymbol).Value != nil {
					return result.(Enviorement.ReturnSymbol)
				}
			}
		}
	}

	return nil
}
