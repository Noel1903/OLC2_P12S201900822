package expressions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
)

type ArithmeticOperations struct {
	left   Abstract.Expression
	right  Abstract.Expression
	op     string
	Line   int
	Column int
	Name   string
}

func NewArithmeticOperations(left Abstract.Expression, right Abstract.Expression, op string, line int, column int) ArithmeticOperations {
	return ArithmeticOperations{
		left:   left,
		right:  right,
		op:     op,
		Line:   line,
		Column: column,
	}
}

func Addition(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) + right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) + right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) + float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) + right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		operation := left.(string) + right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.STRING,
			Value: operation,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.CHAR {
		operation := left.(string) + right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.STRING,
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

func Subtraction(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) - right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) - right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) - float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) - right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
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

func Multiplication(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) * right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) * right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) * float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) * right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
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

func Division(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) / right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: operation,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) / right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) / float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) / right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
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

func Module(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		operation := left.(int) % right.(int)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
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

func (a ArithmeticOperations) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	expleft := a.left.GetValue(table, ast).Value
	expright := a.right.GetValue(table, ast).Value
	typeleft := a.left.GetValue(table, ast).Type
	typeright := a.right.GetValue(table, ast).Type

	a.Name = table.GetName()
	switch a.op {
	case "+":
		result = Addition(expleft, expright, typeleft, typeright)
	case "-":
		result = Subtraction(expleft, expright, typeleft, typeright)
	case "*":
		result = Multiplication(expleft, expright, typeleft, typeright)
	case "/":
		result = Division(expleft, expright, typeleft, typeright)

	case "%":
		result = Module(expleft, expright, typeleft, typeright)

	default:
		err := Errors.NewException("Operador no valido", table.GetName(), a.Line, a.Column)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}

	}

	if result.Type == Enviorement.ERROR {
		result.Value = Errors.NewException("Error en la operacion aritmetica", table.GetName(), a.Line, a.Column)
	}
	return result
}
