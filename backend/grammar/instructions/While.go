package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	"reflect"
)

type While struct {
	condition Abstract.Expression
	codeWhile []interface{}
	Line      int
	Column    int
}

func NewWhile(condition Abstract.Expression, codeWhile []interface{}, line int, column int) *While {
	return &While{
		condition: condition,
		codeWhile: codeWhile,
		Line:      line,
		Column:    column,
	}
}

func (w *While) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	condition := w.condition.GetValue(table, ast)
	if condition.Type != Enviorement.BOOL {
		err := Error.NewException("La condicion no es booleana", "While", w.Line, w.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}

	newEnviorement := Enviorement.NewEnviorement("while", &table)
	var instruc string
	for condition.Value.(bool) {
		for _, instr := range w.codeWhile {
			result := instr.(Abstract.Instruction).Execute(newEnviorement, ast) //ejecuta instrucciones
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
		condition = w.condition.GetValue(table, ast)
		if condition.Type != Enviorement.BOOL {
			err := Error.NewException("La condicion no es booleana", "While", w.Line, w.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
	}

	return nil
}
