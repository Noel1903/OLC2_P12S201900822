package expressions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
)

type ArithmeticOperations struct {
	left  Abstract.Expression
	right Abstract.Expression
	op    string
}

func NewArithmeticOperations(left Abstract.Expression, right Abstract.Expression, op string) ArithmeticOperations {
	return ArithmeticOperations{
		left:  left,
		right: right,
		op:    op,
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
	}
	if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) + right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) + float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) + right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		operation := left.(string) + right.(string)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.STRING,
			Value: operation,
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
	}
	if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) - right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) - float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) - right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
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
	}
	if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) * right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) * float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) * right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
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
	}
	if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		operation := float64(left.(int)) / right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		operation := left.(float64) / float64(right.(int))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
		}
	}
	if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		operation := left.(float64) / right.(float64)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: operation,
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
	}
	return result

}

func (a ArithmeticOperations) GetValue(table Enviorement.SymbolTable) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	expleft := a.left.GetValue(table).Value
	expright := a.right.GetValue(table).Value
	typeleft := a.left.GetValue(table).Type
	typeright := a.right.GetValue(table).Type
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
	}
	return result
}
