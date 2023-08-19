package expressions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
)

type LogicOperations struct {
	left  Abstract.Expression
	right Abstract.Expression
	op    string
}

func NewLogicOperations(left Abstract.Expression, right Abstract.Expression, op string) LogicOperations {
	return LogicOperations{
		left:  left,
		right: right,
		op:    op,
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
	}
	return result
}

func (l LogicOperations) GetValue(table Enviorement.SymbolTable) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if l.left != nil {
		expleft := l.left.GetValue(table).Value
		expright := l.right.GetValue(table).Value
		typeleft := l.left.GetValue(table).Type
		typeright := l.right.GetValue(table).Type

		switch l.op {
		case "&&":
			result = And(expleft, expright, typeleft, typeright)
		case "||":
			result = Or(expleft, expright, typeleft, typeright)
		}
	} else {
		expright := l.right.GetValue(table).Value
		typeright := l.right.GetValue(table).Type
		switch l.op {
		case "!":
			result = Not(expright, typeright)
		}
	}
	return result
}
