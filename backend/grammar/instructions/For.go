package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	"reflect"
)

type For struct {
	identifier string
	Initial    Abstract.Expression
	Final      Abstract.Expression
	codeFor    []interface{}
	Line       int
	Column     int
}

func NewFor(identifier string, Initial Abstract.Expression, Final Abstract.Expression, codeFor []interface{}, line int, column int) *For {
	return &For{
		identifier: identifier,
		Initial:    Initial,
		Final:      Final,
		codeFor:    codeFor,
		Line:       line,
		Column:     column,
	}
}

func (f *For) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	valueChange := table.GetVariable(f.identifier)
	valueInitial := f.Initial.GetValue(table, ast)
	valueFinal := f.Final.GetValue(table, ast)

	//newEnviorement := Enviorement.NewEnviorement("for", &table)
	var instruc string
	if valueChange.Value == nil {
		newVariable := Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: valueInitial.Value,
		}
		table.SetVariable(f.identifier, newVariable, true, f.Line, f.Column, false)
		fmt.Println(valueFinal)
		if valueInitial.Type != Enviorement.INT || valueFinal.Type != Enviorement.INT {
			err := Error.NewException("Los valores deben ser de tipo entero", table.GetName(), f.Line, f.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}

		if valueInitial.Value.(int) > valueFinal.Value.(int) {
			err := Error.NewException("El valor inicial no puede ser mayor al valor final", table.GetName(), f.Line, f.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}

		for newVariable.Value.(int) <= valueFinal.Value.(int) {
			newEnviorement := Enviorement.NewEnviorement("for", &table)
			for _, instr := range f.codeFor {

				result := instr.(Abstract.Instruction).Execute(newEnviorement, ast)

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
		err := Error.NewException("La variable ya existe", table.GetName(), f.Line, f.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}

	return nil
}
