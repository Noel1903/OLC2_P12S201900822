package instructions

import (
	Abstract "grammar/abstract"
	Scoope "grammar/symbol"
)

type DeclareVariable struct {
	typeD Scoope.TypeData
	value Abstract.Expression
}

func newDeclareWithValue(typeVar Scoope.TypeData, value Abstract.Expression) *DeclareVariable {
	return &DeclareVariable{
		typeD: typeVar,
		value: value,
	}
}

func newDeclareWithoutValue(typeVar Scoope.TypeData) *DeclareVariable {
	return &DeclareVariable{
		typeD: typeVar,
		value: nil,
	}
}

/*func (d *DeclareVariable) Execute(env Scoope.SymbolTable)interface{}{

}*/
