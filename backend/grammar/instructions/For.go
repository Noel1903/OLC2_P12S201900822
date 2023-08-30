package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	"reflect"
)

type For struct {
	identifier string
	Initial    Abstract.Expression
	Final      Abstract.Expression
	codeFor    []interface{}
}

func NewFor(identifier string, Initial Abstract.Expression, Final Abstract.Expression, codeFor []interface{}) *For {
	return &For{
		identifier: identifier,
		Initial:    Initial,
		Final:      Final,
		codeFor:    codeFor,
	}
}

func (f *For) Execute(table Enviorement.SymbolTable) interface{} {
	valueChange := table.GetVariable(f.identifier)
	valueInitial := f.Initial.GetValue(table)
	valueFinal := f.Final.GetValue(table)

	newEnviorement := Enviorement.NewEnviorement("for", &table)
	var instruc string
	if valueChange.Value == nil {
		newVariable := Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: valueInitial.Value,
		}
		newEnviorement.SetVariable(f.identifier, newVariable, true)

		if valueInitial.Type != Enviorement.INT || valueFinal.Type != Enviorement.INT {
			fmt.Println("Error: For only accepts numbers")
			return nil
		}

		if valueInitial.Value.(int) > valueFinal.Value.(int) {
			fmt.Println("Error: El valor inicial no puede ser mayor al valor final")
			return nil
		}

		for newVariable.Value.(int) <= valueFinal.Value.(int) {
			for _, instr := range f.codeFor {

				result := instr.(Abstract.Instruction).Execute(newEnviorement)

				if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
					if result.(Enviorement.ReturnSymbol).Value == "break" {
						instruc = "break"
						break
					} else if result.(Enviorement.ReturnSymbol).Value == nil {
						return nil
					} else if result.(Enviorement.ReturnSymbol).Value == "continue" {
						break
					} else if result.(Enviorement.ReturnSymbol).Value != nil {
						return result.(Enviorement.ReturnSymbol)
					}
				}
			}
			if instruc == "break" {
				break
			}
			newVariable.Value = newVariable.Value.(int) + 1
			newEnviorement.UpdateVariable(f.identifier, newVariable)
		}

	} else {
		fmt.Println("Error: La variable ya existe")
		return nil
	}

	return nil
}
