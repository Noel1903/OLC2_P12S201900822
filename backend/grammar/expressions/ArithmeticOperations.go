package expressions

import (
	"fmt"
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
	"reflect"
	"strconv"
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

func Addition(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData, table Enviorement.SymbolTable) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		//operation := left.(int) + right.(int)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "+", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		//operation := float64(left.(int)) + right.(float64)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "+", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		//operation := left.(float64) + float64(right.(int))
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "+", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		//operation := left.(float64) + right.(float64)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "+", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.FLOAT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.STRING {
		temporal := generator.AddTemporal()
		temporal2 := generator.AddTemporal()
		generator.AddAdditionString()

		generator.AddAssign(temporal, "H")
		generator.AddEnv(strconv.Itoa(table.GetSizeEnv()))
		generator.AddAssign(temporal2, "P")
		generator.SetStack(temporal2, left.(string))
		generator.AddExpression(temporal2, "1", "+", temporal2)
		generator.SetStack(temporal2, right.(string))
		generator.AddCall("AdditionString")
		generator.ReturnEnv(strconv.Itoa(table.GetSizeEnv()))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.STRING,
			Value: temporal,
		}
	} else if typeleft == Enviorement.STRING && typeright == Enviorement.CHAR {
		temporal := generator.AddTemporal()
		temporal2 := generator.AddTemporal()
		generator.AddAdditionString()

		generator.AddAssign(temporal, "H")
		generator.AddEnv(strconv.Itoa(table.GetSizeEnv()))
		generator.AddAssign(temporal2, "P")
		generator.SetStack(temporal2, left.(string))
		generator.AddExpression(temporal2, "1", "+", temporal2)
		generator.SetStack(temporal2, right.(string))
		generator.AddCall("AdditionString")
		generator.ReturnEnv(strconv.Itoa(table.GetSizeEnv()))
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.STRING,
			Value: temporal,
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
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		//operation := left.(int) - right.(int)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "-", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		//operation := float64(left.(int)) - right.(float64)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "-", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		//operation := left.(float64) - float64(right.(int))

		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "-", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		//operation := left.(float64) - right.(float64)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "-", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
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
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		//operation := left.(int) * right.(int)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "*", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		//operation := float64(left.(int)) * right.(float64)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "*", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		//operation := left.(float64) * float64(right.(int))
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "*", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		//operation := left.(float64) * right.(float64)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "*", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
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
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		//operation := left.(int) / right.(int)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "/", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.INT && typeright == Enviorement.FLOAT {
		//operation := float64(left.(int)) / right.(float64)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "/", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.INT {
		//operation := left.(float64) / float64(right.(int))
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "/", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
		}
	} else if typeleft == Enviorement.FLOAT && typeright == Enviorement.FLOAT {
		//operation := left.(float64) / right.(float64)
		temporal := generator.AddTemporal()
		generator.AddExpression(left.(string), right.(string), "/", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
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
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if typeleft == Enviorement.INT && typeright == Enviorement.INT {
		//operation := left.(int) % right.(int)
		temporal := generator.AddTemporal()
		generator.AddExpression("(int)"+left.(string), "(int)"+right.(string), "%", temporal)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.INT,
			Value: temporal,
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

	a.Name = table.GetName()
	switch a.op {
	case "+":
		expleft := a.left.GetValue(table, ast)
		expright := a.right.GetValue(table, ast)
		typeleft := expleft.Type
		typeright := expright.Type
		left := expleft.Value
		right := expright.Value
		result = Addition(left, right, typeleft, typeright, table)
	case "-":
		expleft := a.left.GetValue(table, ast)
		expright := a.right.GetValue(table, ast)
		typeleft := expleft.Type
		typeright := expright.Type
		left := expleft.Value
		right := expright.Value
		result = Subtraction(left, right, typeleft, typeright)
	case "*":
		/*expleft := a.left.GetValue(table, ast)
		expright := a.right.GetValue(table, ast)
		typeleft := expleft.Type
		typeright := expright.Type
		left := expleft.Value
		right := expright.Value
		result = Multiplication(left, right, typeleft, typeright)*/
		fmt.Println(reflect.TypeOf(a.right))
		if reflect.TypeOf(a.right) == reflect.TypeOf(&callFunctionExp{}) {
			expright := a.right.GetValue(table, ast)
			right := expright.Value
			expleft := a.left.GetValue(table, ast)
			left := expleft.Value
			typeleft := expleft.Type
			typeright := expright.Type
			result = Multiplication(left, right, typeleft, typeright)
		} else {
			expleft := a.left.GetValue(table, ast)
			expright := a.right.GetValue(table, ast)
			typeleft := expleft.Type
			typeright := expright.Type
			left := expleft.Value
			right := expright.Value
			result = Multiplication(left, right, typeleft, typeright)
		}
	case "/":
		expleft := a.left.GetValue(table, ast)
		expright := a.right.GetValue(table, ast)
		typeleft := expleft.Type
		typeright := expright.Type
		left := expleft.Value
		right := expright.Value
		result = Division(left, right, typeleft, typeright)

	case "%":
		expleft := a.left.GetValue(table, ast)
		expright := a.right.GetValue(table, ast)
		typeleft := expleft.Type
		typeright := expright.Type
		left := expleft.Value
		right := expright.Value
		result = Module(left, right, typeleft, typeright)

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
