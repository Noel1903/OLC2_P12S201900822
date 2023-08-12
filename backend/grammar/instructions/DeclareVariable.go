package instructions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
)

type DeclareVariable struct {
	identifier string
	typeD      Enviorement.TypeData
	value      Abstract.Expression
}

func NewDeclareWithValue(identifier string, typeVar Enviorement.TypeData, value Abstract.Expression) *DeclareVariable {
	return &DeclareVariable{
		identifier: identifier,
		typeD:      typeVar,
		value:      value,
	}
}

func NewDeclareWithoutValue(identifier string, typeD Enviorement.TypeData) *DeclareVariable {
	return &DeclareVariable{
		identifier: identifier,
		typeD:      typeD,
		value:      nil,
	}
}

func (d *DeclareVariable) Execute(table Enviorement.SymbolTable) interface{} {

	variable := table.GetVariable(d.identifier)
	if variable == nil {
		if d.value != nil {
			value := d.value.GetValue(table)
			table.SetVariable(d.identifier, value)
		} else {
			table.SetVariable(d.identifier, nil)
		}
	}

	return nil

}
