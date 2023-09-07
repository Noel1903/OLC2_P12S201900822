package expressions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
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

func And(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.BOOL && typeright == Enviorement.BOOL {
		operation := left.(bool) && right.(bool)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else {
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: nil,
		}
	}
	return result
}

func Or(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.BOOL && typeright == Enviorement.BOOL {
		operation := left.(bool) || right.(bool)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else {
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: nil,
		}
	}
	return result
}

func Not(right interface{}, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeright == Enviorement.BOOL {
		operation := !right.(bool)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else {
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: nil,
		}
	}
	return result
}

func (l LogicOperations) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if l.left != nil {
		expleft := l.left.GetValue(table, ast).Value
		expright := l.right.GetValue(table, ast).Value
		typeleft := l.left.GetValue(table, ast).Type
		typeright := l.right.GetValue(table, ast).Type

		switch l.op {
		case "&&":
			result = And(expleft, expright, typeleft, typeright)
		case "||":
			result = Or(expleft, expright, typeleft, typeright)
		default:
			result = Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: nil,
			}
		}
	} else {
		expright := l.right.GetValue(table, ast).Value
		typeright := l.right.GetValue(table, ast).Type
		switch l.op {
		case "!":
			result = Not(expright, typeright)

		default:
			result = Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: nil,
			}
		}
	}

	if result.Type == Enviorement.ERROR {
		result.Value = Errors.NewException("Error en la operacion aritmetica", table.GetName(), l.Line, l.Column)
	}
	return result
}
