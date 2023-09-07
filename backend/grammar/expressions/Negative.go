package expressions

import (
	"fmt"
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
	if value.Type != enviorement.INT && value.Type != enviorement.FLOAT {
		newErr := errors.NewException("La expresion no es de tipo entero", table.GetName(), n.Line, n.Column)
		return enviorement.ReturnSymbol{Type: enviorement.ERROR, Value: newErr}
	}

	fmt.Println(value.Value, value.Type)
	if value.Type == enviorement.INT {
		return enviorement.ReturnSymbol{Type: value.Type, Value: -value.Value.(int)}
	}
	if value.Type == enviorement.FLOAT {
		return enviorement.ReturnSymbol{Type: value.Type, Value: -value.Value.(float64)}
	}

	return enviorement.ReturnSymbol{Type: enviorement.NIL, Value: nil}
}
