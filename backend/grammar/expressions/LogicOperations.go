package expressions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
)

type LogicOperations struct {
	left   Abstract.Expression
	right  Abstract.Expression
	op     string
	Line   int
	Column int
}

func NewLogicOperations(left Abstract.Expression, right Abstract.Expression, op string, line int, column int) LogicOperations {
	return LogicOperations{
		left:   left,
		right:  right,
		op:     op,
		Line:   line,
		Column: column,
	}
}

func (l LogicOperations) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if l.left != nil {
		expleft := l.left
		expright := l.right

		switch l.op {
		case "&&":
			left := expleft.GetValue(table, ast)
			for _, labels := range left.LabelTrue {
				generator.PutLabel(labels.(string))
			}
			right := expright.GetValue(table, ast)
			result = Enviorement.ReturnSymbol{
				Type:  Enviorement.BOOL,
				Value: "",
			}
			result.LabelTrue = append(right.LabelTrue, result.LabelTrue...)
			result.LabelFalse = append(left.LabelFalse, result.LabelFalse...)
			result.LabelFalse = append(right.LabelFalse, result.LabelFalse...)
			return result

		case "||":
			left := expleft.GetValue(table, ast)
			for _, labels := range left.LabelFalse {
				generator.PutLabel(labels.(string))
			}
			right := expright.GetValue(table, ast)
			result = Enviorement.ReturnSymbol{
				Type:  Enviorement.BOOL,
				Value: "",
			}
			result.LabelTrue = append(left.LabelTrue, result.LabelTrue...)
			result.LabelTrue = append(right.LabelTrue, result.LabelTrue...)
			result.LabelFalse = append(right.LabelFalse, result.LabelFalse...)
			return result

		default:
			err := Errors.NewException("Error: Operador no valido", table.GetName(), l.Line, l.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
	} else {
		//fmt.Println(l.right)
		right := l.right.GetValue(table, ast)
		if right.Type == Enviorement.BOOL {
			result = Enviorement.ReturnSymbol{
				Type:  Enviorement.BOOL,
				Value: "",
			}
			result.LabelTrue = append(right.LabelFalse, result.LabelTrue...)
			result.LabelFalse = append(right.LabelTrue, result.LabelFalse...)

			return result
		}
		err := Errors.NewException("Error: Operador no valido", table.GetName(), l.Line, l.Column)

		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}

	}
}
