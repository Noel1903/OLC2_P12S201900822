package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	"reflect"
)

type Guard struct {
	condition Abstract.Expression
	codeGuard []interface{}
	Line      int
	Column    int
}

func NewGuard(condition Abstract.Expression, codeGuard []interface{}, line int, column int) *Guard {
	return &Guard{
		condition: condition,
		codeGuard: codeGuard,
		Line:      line,
		Column:    column,
	}
}

func (g *Guard) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	condition := g.condition.GetValue(table, ast)
	if condition.Type != Enviorement.BOOL {
		err := Error.NewException("Error: La condicion debe ser de tipo booleana", table.GetName(), g.Line, g.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	newEnviorement := Enviorement.NewEnviorement("guard", &table)
	if condition.Value.(bool) == false {
		for _, instr := range g.codeGuard {
			result := instr.(Abstract.Instruction).Execute(newEnviorement, ast)
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
