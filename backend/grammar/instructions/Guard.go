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
	genAux := Enviorement.NewGenerator()
	generator := genAux.GetInstance()

	condition := g.condition.GetValue(table, ast)
	if condition.Type != Enviorement.BOOL {
		err := Error.NewException("Error: La condicion debe ser de tipo booleana", table.GetName(), g.Line, g.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	newEnviorement := Enviorement.NewEnviorement("guard", &table)
	for _, labels := range condition.LabelFalse {
		if reflect.TypeOf(labels).Kind() != reflect.Slice || reflect.TypeOf(labels).Elem().Kind() == reflect.Interface {
			generator.PutLabel(labels.(string))
		}
	}

	for _, instr := range g.codeGuard {
		result := instr.(Abstract.Instruction).Execute(newEnviorement, ast)
		//fmt.Println("ingresa a guard")
		if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
			if result.(Enviorement.ReturnSymbol).Type == Enviorement.BREAK {
				generator.AddGoto(table.GetFalseLabel())
			}
			if result.(Enviorement.ReturnSymbol).Type == Enviorement.CONTINUE {
				generator.AddGoto(table.GetTrueLabel())
			}
		}

	}
	for _, labels := range condition.LabelTrue {
		if reflect.TypeOf(labels).Kind() != reflect.Slice || reflect.TypeOf(labels).Elem().Kind() == reflect.Interface {
			generator.PutLabel(labels.(string))
		}
	}
	return nil
}
