package expressions

import (
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
	"strconv"
)

type vectorValues struct {
	Identifier string
	Id         string
	Line       int
	Column     int
}

func NewVectorValues(Identifier string, Id string, line int, column int) *vectorValues {
	return &vectorValues{Identifier: Identifier, Id: Id, Line: line, Column: column}
}

func (n vectorValues) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	variable := table.GetVariable(n.Identifier)
	genAux := Enviorement.NewGenerator()
	generator := genAux.GetInstance()
	if variable.Value == nil {
		err := Errors.NewException("La variable no ha sido declarada", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	if n.Id == "IsEmpty" {
		if variable.Type != Enviorement.ARRAY {
			err := Errors.NewException("La variable no es un vector", table.GetName(), n.Line, n.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		tempVar := table.GetVar(n.Identifier)
		temporal01 := generator.AddTemporal()
		temporal02 := generator.AddTemporal()
		tempSizeThis := generator.AddTemporal()
		labelTrue := generator.AddLabel()
		labelFalse := generator.AddLabel()
		if tempVar.IsGlobal {
			generator.GetStack(strconv.Itoa(tempVar.GetPos()), temporal02)
		} else {
			generator.AddExpression("P", strconv.Itoa(tempVar.GetPos()), "+", temporal01)
			generator.GetStack(temporal01, temporal02)
		}
		generator.AddExpression(temporal02, "2", "+", tempSizeThis)
		generator.GetHeap(tempSizeThis, tempSizeThis)
		generator.AddIf(tempSizeThis, "0", "==", labelTrue)
		generator.AddGoto(labelFalse)
		returnSym := Enviorement.ReturnSymbol{Type: Enviorement.BOOL, Value: ""}
		returnSym.LabelTrue = append(returnSym.LabelTrue, labelTrue)
		returnSym.LabelFalse = append(returnSym.LabelFalse, labelFalse)
		return returnSym

	} else if n.Id == "count" {
		if variable.Type != Enviorement.ARRAY {
			err := Errors.NewException("La variable no es un vector", table.GetName(), n.Line, n.Column)
			return Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: err,
			}
		}
		tempVar := table.GetVar(n.Identifier)
		temporal01 := generator.AddTemporal()
		temporal02 := generator.AddTemporal()
		tempSizeThis := generator.AddTemporal()
		if tempVar.IsGlobal {
			generator.GetStack(strconv.Itoa(tempVar.GetPos()), temporal02)
		} else {
			generator.AddExpression("P", strconv.Itoa(tempVar.GetPos()), "+", temporal01)
			generator.GetStack(temporal01, temporal02)
		}
		generator.AddExpression(temporal02, "2", "+", tempSizeThis)
		generator.GetHeap(tempSizeThis, tempSizeThis)

		returnSym := Enviorement.ReturnSymbol{Type: Enviorement.INT, Value: tempSizeThis}

		return returnSym
	}
	return Enviorement.ReturnSymbol{}
}
