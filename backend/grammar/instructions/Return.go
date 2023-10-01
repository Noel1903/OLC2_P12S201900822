package instructions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
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

	//fmt.Println("Valor de retorno: ", valueReturn.Value)
	genAuxiliar := Generator.NewGenerator()
	generator := genAuxiliar.GetInstance()
	if valueReturn.Value == nil {
		generator.AddReturn("")
	} else {
		generator.AddExpression("P", "1", "+", "P")
		generator.SetStack("P", valueReturn.Value.(string))
	}
	return Enviorement.ReturnSymbol{
		Type:  valueReturn.Type,
		Value: valueReturn.Value,
	}

}
