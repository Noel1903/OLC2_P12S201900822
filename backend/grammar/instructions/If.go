package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Enviorment "grammar/symbol"
	Generator "grammar/symbol"
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
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()

	labelJump := generator.AddLabel()
	if condition.Type != Enviorment.BOOL {
		err := Error.NewException("Error: La condicion debe ser de tipo booleana", table.GetName(), i.Line, i.Column)
		return Enviorment.ReturnSymbol{
			Type:  Enviorment.ERROR,
			Value: err,
		}
	}

	newEnviorement := Enviorment.NewEnviorement("if", &table)
	newEnviorement.SetTrueLabel(table.GetTrueLabel())
	newEnviorement.SetFalseLabel(table.GetFalseLabel())
	for _, labels := range condition.LabelTrue {

		if reflect.TypeOf(labels).Kind() != reflect.Slice || reflect.TypeOf(labels).Elem().Kind() == reflect.Interface {
			generator.PutLabel(labels.(string))

		}
	}

	for _, instr := range i.codeIf {
		result := instr.(Abstract.Instruction).Execute(newEnviorement, ast)
		if reflect.TypeOf(result) == reflect.TypeOf(Enviorment.ReturnSymbol{}) {
			if result.(Enviorment.ReturnSymbol).Type == Enviorment.BREAK {
				generator.AddGoto(table.GetFalseLabel())
			}
			if result.(Enviorment.ReturnSymbol).Type == Enviorment.CONTINUE {
				generator.AddGoto(table.GetTrueLabel())
			}
		}

	}
	generator.AddGoto(labelJump)
	newEnviorement01 := Enviorment.NewEnviorement("else", &table)
	newEnviorement01.SetTrueLabel(table.GetTrueLabel())
	newEnviorement01.SetFalseLabel(table.GetFalseLabel())
	for _, labels := range condition.LabelFalse {
		if reflect.TypeOf(labels).Kind() != reflect.Slice || reflect.TypeOf(labels).Elem().Kind() == reflect.Interface {
			generator.PutLabel(labels.(string))
		}

	}
	for _, instr := range i.codeElse {
		result := instr.(Abstract.Instruction).Execute(newEnviorement01, ast)
		if reflect.TypeOf(result) == reflect.TypeOf(Enviorment.ReturnSymbol{}) {
			if result.(Enviorment.ReturnSymbol).Type == Enviorment.BREAK {
				generator.AddGoto(table.GetFalseLabel())
			}
			if result.(Enviorment.ReturnSymbol).Type == Enviorment.CONTINUE {
				generator.AddGoto(table.GetTrueLabel())
			}
		}
	}
	generator.PutLabel(labelJump)

	return nil
}
