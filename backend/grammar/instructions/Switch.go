package instructions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	"reflect"
)

type Switch struct {
	Expression Abstract.Expression
	Cases      []interface{}
	Line       int
	Column     int
}

func NewSwitch(Expression Abstract.Expression, Cases []interface{}, line int, column int) *Switch {
	return &Switch{
		Expression: Expression,
		Cases:      Cases,
		Line:       line,
		Column:     column,
	}
}

func (s *Switch) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	genAux := Enviorement.NewGenerator()
	generator := genAux.GetInstance()
	generator.AddComment("Inicia Switch")
	labelInit := generator.AddLabel()
	labelEnd := generator.AddLabel()
	generator.PutLabel(labelInit)
	value := s.Expression.GetValue(table, ast)
	var labels []string
	for _, instr := range s.Cases {

		expr1 := instr.(*CaseSwitch).GetOption()

		label := generator.AddLabel()
		labels = append(labels, label)
		if expr1 != nil {
			expr := expr1.GetValue(table, ast)
			generator.AddIf(value.Value.(string), expr.Value.(string), "==", label)
		} else {
			generator.AddGoto(label)
		}

	}
	for i, instr := range s.Cases {
		generator.PutLabel(labels[i])
		for _, ins := range instr.(*CaseSwitch).GetInstructions() {
			result := ins.(Abstract.Instruction).Execute(table, ast)
			if reflect.TypeOf(result) == reflect.TypeOf(Enviorement.ReturnSymbol{}) {
				if result.(Enviorement.ReturnSymbol).Type == Enviorement.BREAK {
					generator.AddGoto(labelEnd)
				}
			}
		}
	}
	generator.PutLabel(labelEnd)
	return nil
}
