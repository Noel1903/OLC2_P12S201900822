package instructions

import (
	Abstract "grammar/abstract"
	"grammar/symbol"
	Enviorement "grammar/symbol"
)

type DeclareLet struct {
	identifier string
	typeD      Enviorement.TypeData
	value      Abstract.Expression
}

func NewLet(identifier string, typeVar Enviorement.TypeData, value Abstract.Expression) *DeclareLet {
	return &DeclareLet{
		identifier: identifier,
		typeD:      typeVar,
		value:      value,
	}
}

func (l *DeclareLet) Execute(table Enviorement.SymbolTable) interface{} {

	variable := table.GetVariable(l.identifier)

	if variable.Value == nil {
		value := l.value.GetValue(table)
		if l.typeD == symbol.UNDEFINED {
			l.typeD = value.Type
		}
		if value.Type != l.typeD && value.Type != symbol.NIL {
			return nil
		}
		table.SetVariable(l.identifier, value, false)
	} else {
		return nil
	}

	return nil

}
