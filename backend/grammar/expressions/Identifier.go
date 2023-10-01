package expressions

import (
	errors "grammar/exceptions"
	Generator "grammar/symbol"
	symbol "grammar/symbol"
	"strconv"
)

type Identifier struct {
	Identifier string
	Line       int
	Column     int
}

func NewIdentifier(id string, Line int, Column int) Identifier {
	return Identifier{Identifier: id, Line: Line, Column: Column}
}

func (i Identifier) GetValue(table symbol.SymbolTable, ast *symbol.AST) symbol.ReturnSymbol {
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	variable := table.GetVariable(i.Identifier)

	if variable.Value == nil {
		newErr := errors.NewException("La variable "+i.Identifier+" no existe", table.GetName(), i.Line, i.Column)
		return symbol.ReturnSymbol{Type: symbol.ERROR, Value: newErr}
	}
	varTemp := table.GetVar(i.Identifier)
	temporal := generator.AddTemporal()
	varPosition := varTemp.GetPos()

	if !varTemp.GetIsGlobal() {
		tempPos := generator.AddTemporal()
		generator.AddExpression("P", strconv.Itoa(varPosition), "+", tempPos)
		generator.GetStack(strconv.Itoa(varPosition), tempPos)
	} else {
		generator.GetStack(strconv.Itoa(varPosition), temporal)
	}

	return symbol.ReturnSymbol{Type: variable.Type, Value: temporal}
}
