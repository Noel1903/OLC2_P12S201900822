package instructions

import (
	"fmt"
	abstract "grammar/abstract"
	symbol "grammar/symbol"
)

type Print struct {
	Expression abstract.Expression
}

func NewPrint(Expression abstract.Expression) *Print {
	return &Print{Expression: Expression}
}

func (p Print) Execute(table symbol.SymbolTable) interface{} {
	value := p.Expression.GetValue(table)
	fmt.Println(value.Value)
	return nil
}
