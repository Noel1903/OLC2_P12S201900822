package expressions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
)

type RelationalOperations struct {
	left   Abstract.Expression
	right  Abstract.Expression
	op     string
	Line   int
	Column int
}

func NewRelationalOperations(left Abstract.Expression, right Abstract.Expression, op string, line int, column int) RelationalOperations {
	return RelationalOperations{
		left:   left,
		right:  right,
		op:     op,
		Line:   line,
		Column: column,
	}
}

func LessThan(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) < right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) < right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) < float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) < right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		operation := left.(string) < right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.CHAR && typeright == Enviorement.CHAR {
		operation := left.(string) < right.(string)
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

func GreaterThan(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) > right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) > right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) > float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) > right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		operation := left.(string) > right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.CHAR && typeright == Enviorement.CHAR {
		operation := left.(string) > right.(string)
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

func LessThanEquals(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) <= right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) <= right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) <= float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) <= right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		operation := left.(string) <= right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.CHAR && typeright == Enviorement.CHAR {
		operation := left.(string) <= right.(string)
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

func GreaterThanEquals(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) >= right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) >= right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) >= float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) >= right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		operation := left.(string) >= right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.CHAR && typeright == Enviorement.CHAR {
		operation := left.(string) >= right.(string)
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

func Equals(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) == right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) == right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) == float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) == right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		operation := left.(string) == right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.CHAR && typeright == Enviorement.CHAR {
		operation := left.(string) == right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.BOOL && typeright == Enviorement.BOOL {
		operation := left.(bool) == right.(bool)
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

func NotEquals(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) != right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) != right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) != float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) != right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		operation := left.(string) != right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.CHAR && typeright == Enviorement.CHAR {
		operation := left.(string) != right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: operation,
		}
	} else if typeleft == Enviorement.BOOL && typeright == Enviorement.BOOL {
		operation := left.(bool) != right.(bool)
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

func (r RelationalOperations) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	expleft := r.left.GetValue(table, ast).Value
	expright := r.right.GetValue(table, ast).Value
	typeleft := r.left.GetValue(table, ast).Type
	typeright := r.right.GetValue(table, ast).Type
	switch r.op {
	case "<":
		result = LessThan(expleft, expright, typeleft, typeright)
	case ">":
		result = GreaterThan(expleft, expright, typeleft, typeright)
	case "<=":
		result = LessThanEquals(expleft, expright, typeleft, typeright)
	case ">=":
		result = GreaterThanEquals(expleft, expright, typeleft, typeright)
	case "==":
		result = Equals(expleft, expright, typeleft, typeright)
	case "!=":
		result = NotEquals(expleft, expright, typeleft, typeright)
	default:
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: nil,
		}
	}

	if result.Type == Enviorement.ERROR {
		result.Value = Errors.NewException("Error en la operacion aritmetica", table.GetName(), r.Line, r.Column)
	}
	return result

}
