package instructions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
)

type Return struct {
	value Abstract.Expression
}

func NewReturn(value Abstract.Expression) *Return {
	return &Return{
		value: value,
	}
}

func (r *Return) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	valueReturn := r.value.GetValue(table, ast)
	return Enviorement.ReturnSymbol{
		Type:  valueReturn.Type,
		Value: valueReturn.Value,
	}

}
