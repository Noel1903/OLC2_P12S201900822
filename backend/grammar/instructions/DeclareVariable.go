package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	"grammar/symbol"
	Enviorement "grammar/symbol"
)

type DeclareVariable struct {
	Identifier string
	TypeD      Enviorement.TypeData
	Value      Abstract.Expression
}

func NewDeclareWithValue(identifier string, typeVar Enviorement.TypeData, value Abstract.Expression) *DeclareVariable {
	return &DeclareVariable{
		Identifier: identifier,
		TypeD:      typeVar,
		Value:      value,
	}
}

func NewDeclareWithoutValue(identifier string, typeD Enviorement.TypeData, value Abstract.Expression) *DeclareVariable {
	return &DeclareVariable{
		Identifier: identifier,
		TypeD:      typeD,
		Value:      value,
	}
}

func (d *DeclareVariable) Execute(table Enviorement.SymbolTable) interface{} {
	var variable symbol.ReturnSymbol
	variable = table.GetVariable(d.Identifier)

	if variable.Value == nil {
		/*if d.value != nil {
			value := d.value.GetValue(table)
			fmt.Println(value)
			table.SetVariable(d.identifier, value)
		} else {
			table.SetVariable(d.identifier, nil)
		}*/

		value := d.Value.GetValue(table)
		if d.TypeD == symbol.UNDEFINED {
			d.TypeD = value.Type
		}
		if value.Type != d.TypeD && value.Type != symbol.NIL {
			fmt.Println("Error, los tipos no coinciden")
			return nil
		}
		table.SetVariable(d.Identifier, value, true)
	} else {
		fmt.Println("Error, la variable ya existe")
		return nil
	}

	return nil

}
