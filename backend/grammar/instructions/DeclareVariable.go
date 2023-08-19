package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	"grammar/symbol"
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

func NewDeclareWithoutValue(identifier string, typeD Enviorement.TypeData, value Abstract.Expression) *DeclareVariable {
	return &DeclareVariable{
		identifier: identifier,
		typeD:      typeD,
		value:      value,
	}
}

func (d *DeclareVariable) Execute(table Enviorement.SymbolTable) interface{} {

	variable := table.GetVariable(d.identifier)

	if variable == nil {
		/*if d.value != nil {
			value := d.value.GetValue(table)
			fmt.Println(value)
			table.SetVariable(d.identifier, value)
		} else {
			table.SetVariable(d.identifier, nil)
		}*/

		value := d.value.GetValue(table)
		if d.typeD == symbol.UNDEFINED {
			d.typeD = value.Type
		}
		if value.Type != d.typeD && value.Type != symbol.NIL {
			fmt.Println("Error, los tipos no coinciden")
			return nil
		}
		table.SetVariable(d.identifier, value)
	} else {
		fmt.Println("Error, la variable ya existe")
		return nil
	}

	return nil

}
