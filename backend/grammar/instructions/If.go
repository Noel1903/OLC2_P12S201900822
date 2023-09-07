package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Enviorment "grammar/symbol"
	"reflect"
)

type If struct {
	condition Abstract.Expression
	codeIf    []interface{}
	codeElse  []interface{}
	Line      int
	Column    int
}

func NewIf(condition Abstract.Expression, codeIf []interface{}, codeElse []interface{}, line int, column int) *If {
	return &If{
		condition: condition,
		codeIf:    codeIf,
		codeElse:  codeElse,
		Line:      line,
		Column:    column,
	}
}

func (i *If) Execute(table Enviorment.SymbolTable, ast *Enviorement.AST) interface{} {
	condition := i.condition.GetValue(table, ast)
	if condition.Type != Enviorment.BOOL {
		err := Error.NewException("Error: La condicion debe ser de tipo booleana", table.GetName(), i.Line, i.Column)
		return Enviorment.ReturnSymbol{
			Type:  Enviorment.ERROR,
			Value: err,
		}
	}

	newEnviorement := Enviorment.NewEnviorement("if", &table)

	if condition.Value.(bool) {
		for _, instr := range i.codeIf {
			result := instr.(Abstract.Instruction).Execute(newEnviorement, ast)
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
			result := instr.(Abstract.Instruction).Execute(newEnviorement, ast)
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
