package expressions

import (
	abstract "grammar/abstract"
	errors "grammar/exceptions"
	enviorement "grammar/symbol"
)

type Negative struct {
	Expression abstract.Expression
	Line       int
	Column     int
}

func NewNegative(exp abstract.Expression, Line int, Column int) Negative {
	return Negative{Expression: exp, Line: Line, Column: Column}
}

func (n Negative) GetValue(table enviorement.SymbolTable, ast *enviorement.AST) enviorement.ReturnSymbol {
	value := n.Expression.GetValue(table, ast)
	genAux := enviorement.NewGenerator()
	generator := genAux.GetInstance()
	if value.Type != enviorement.INT && value.Type != enviorement.FLOAT {
		newErr := errors.NewException("La expresion no es de tipo entero", table.GetName(), n.Line, n.Column)
		return enviorement.ReturnSymbol{Type: enviorement.ERROR, Value: newErr}
	}

	//fmt.Println(value.Value, value.Type)
	if value.Type == enviorement.INT {
		temporal := generator.AddTemporal()
		generator.AddExpression("0", value.Value.(string), "-", temporal)
		return enviorement.ReturnSymbol{Type: value.Type, Value: temporal}
	}
	if value.Type == enviorement.FLOAT {
		temporal := generator.AddTemporal()
		generator.AddExpression("0", value.Value.(string), "-", temporal)
		return enviorement.ReturnSymbol{Type: value.Type, Value: temporal}
	}

	return enviorement.ReturnSymbol{Type: enviorement.NIL, Value: nil}
}
