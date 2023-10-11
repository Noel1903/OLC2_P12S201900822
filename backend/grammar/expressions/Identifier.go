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
	temporal01 := generator.AddTemporal()
	temporal02 := generator.AddTemporal()
	if varTemp.IsGlobal {
		generator.GetStack(strconv.Itoa(varTemp.GetPos()), temporal02)
	} else {
		generator.AddExpression("P", strconv.Itoa(varTemp.GetPos()), "+", temporal01)
		generator.GetStack(temporal01, temporal02)
	}
	//temporal := generator.AddTemporal()
	//varPosition := varTemp.GetPos()

	//generator.GetStack(strconv.Itoa(varPosition), temporal)
	if variable.Type == symbol.BOOL {
		LabelTrue := generator.AddLabel()
		LabelFalse := generator.AddLabel()
		generator.AddIf(temporal02, "1", "==", LabelTrue)
		generator.AddGoto(LabelFalse)
		result := symbol.ReturnSymbol{
			Type:  symbol.BOOL,
			Value: "",
		}
		result.LabelTrue = append(result.LabelTrue, LabelTrue)
		result.LabelFalse = append(result.LabelFalse, LabelFalse)
		return result
	} else {
		return symbol.ReturnSymbol{Type: variable.Type, Value: temporal02}
	}

}
