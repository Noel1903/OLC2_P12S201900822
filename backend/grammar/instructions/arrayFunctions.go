package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	"strconv"
)

type arrayFunctions struct {
	Identifier string
	Id         string
	Expression interface{}
	Line       int
	Column     int
}

func NewArrayFunctions(Identifier string, Id string, Expression interface{}, line int, column int) *arrayFunctions {
	return &arrayFunctions{Identifier: Identifier, Id: Id, Expression: Expression, Line: line, Column: column}
}

func (n arrayFunctions) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	variable := table.GetVariable(n.Identifier)
	genAux := Enviorement.NewGenerator()
	generator := genAux.GetInstance()

	if variable.Value == nil {
		err := Error.NewException("Variable no encontrada", n.Identifier, n.Line, n.Column)

		return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
	}
	if n.Id == "append" {
		value := n.Expression.(Abstract.Expression).GetValue(table, ast)
		tempVar := table.GetVar(n.Identifier)
		temporal01 := generator.AddTemporal()
		temporal02 := generator.AddTemporal()
		if tempVar.IsGlobal {
			generator.GetStack(strconv.Itoa(tempVar.GetPos()), temporal02)
		} else {
			generator.AddExpression("P", strconv.Itoa(tempVar.GetPos()), "+", temporal01)
			generator.GetStack(temporal01, temporal02)
		}
		tempFinal := generator.AddTemporal()
		generator.AddAssign(tempFinal, "H")

		//tempInitial := generator.AddTemporal()
		//generator.GetHeap(temporal02, tempInitial)

		tempSizeThis := generator.AddTemporal()
		tempCont := generator.AddTemporal()
		tempChar := generator.AddTemporal()
		tempNewChar := generator.AddTemporal()
		labelTrue := generator.AddLabel()
		labelFalse := generator.AddLabel()
		tempSizeNew := generator.AddTemporal()
		generator.AddExpression(temporal02, "2", "+", tempSizeThis)
		generator.GetHeap(tempSizeThis, tempSizeThis)
		generator.AddExpression(temporal02, "3", "+", tempChar)
		generator.AddExpression(tempSizeThis, "1", "+", tempSizeNew)
		generator.SetHeap("H", "1")
		generator.AddExpression("H", "1", "+", "H")
		generator.SetHeap("H", "0")
		generator.AddExpression("H", "1", "+", "H")
		generator.SetHeap("H", tempSizeNew)
		generator.AddExpression("H", "1", "+", "H")
		generator.PutLabel(labelFalse)
		generator.AddIf(tempCont, "(int)"+tempSizeThis, "==", labelTrue)
		generator.GetHeap(tempChar, tempNewChar)
		generator.SetHeap("H", tempNewChar)
		generator.AddExpression("H", "1", "+", "H")
		generator.AddExpression(tempChar, "1", "+", tempChar)
		generator.AddExpression(tempCont, "1", "+", tempCont)
		generator.AddGoto(labelFalse)
		generator.PutLabel(labelTrue)
		generator.SetHeap("H", value.Value.(string))
		generator.AddExpression("H", "1", "+", "H")
		generator.SetStack(strconv.Itoa(tempVar.GetPos()), tempFinal)
		return nil

	} else if n.Id == "removeLast" {
		if variable.Type == Enviorement.ARRAY {
			tempVar := table.GetVar(n.Identifier)
			temporal01 := generator.AddTemporal()
			temporal02 := generator.AddTemporal()
			if tempVar.IsGlobal {
				generator.GetStack(strconv.Itoa(tempVar.GetPos()), temporal02)
			} else {
				generator.AddExpression("P", strconv.Itoa(tempVar.GetPos()), "+", temporal01)
				generator.GetStack(temporal01, temporal02)
			}
			tempFinal := generator.AddTemporal()
			generator.AddAssign(tempFinal, "H")

			//tempInitial := generator.AddTemporal()
			//generator.GetHeap(temporal02, tempInitial)

			tempSizeThis := generator.AddTemporal()
			tempCont := generator.AddTemporal()
			tempChar := generator.AddTemporal()
			tempNewChar := generator.AddTemporal()
			labelTrue := generator.AddLabel()
			labelFalse := generator.AddLabel()
			tempSizeNew := generator.AddTemporal()
			generator.AddExpression(temporal02, "2", "+", tempSizeThis)
			generator.GetHeap(tempSizeThis, tempSizeThis)
			generator.AddExpression(temporal02, "3", "+", tempChar)
			generator.AddExpression(tempSizeThis, "1", "-", tempSizeNew)
			generator.SetHeap("H", "1")
			generator.AddExpression("H", "1", "+", "H")
			generator.SetHeap("H", "0")
			generator.AddExpression("H", "1", "+", "H")
			generator.SetHeap("H", tempSizeNew)
			generator.AddExpression("H", "1", "+", "H")
			generator.PutLabel(labelFalse)
			generator.AddIf(tempCont, "(int)"+tempSizeNew, "==", labelTrue)
			generator.GetHeap(tempChar, tempNewChar)
			generator.SetHeap("H", tempNewChar)
			generator.AddExpression("H", "1", "+", "H")
			generator.AddExpression(tempChar, "1", "+", tempChar)
			generator.AddExpression(tempCont, "1", "+", tempCont)
			generator.AddGoto(labelFalse)
			generator.PutLabel(labelTrue)

			generator.SetStack(strconv.Itoa(tempVar.GetPos()), tempFinal)
			return nil

		} else {
			err := Error.NewException("No se puede remover el valor", n.Identifier, n.Line, n.Column)
			return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
		}
	} else if n.Id == "remove" {
		index := n.Expression.(Abstract.Expression).GetValue(table, ast)
		if variable.Type == Enviorement.ARRAY {
			if index.Type != Enviorement.INT {
				err := Error.NewException("No se puede remover el valor", n.Identifier, n.Line, n.Column)
				return Enviorement.ReturnSymbol{Type: Enviorement.ERROR, Value: err}
			}
			varTemp := table.GetVar(n.Identifier)
			temporal01 := generator.AddTemporal()
			temporal02 := generator.AddTemporal()
			if varTemp.IsGlobal {
				generator.GetStack(strconv.Itoa(varTemp.GetPos()), temporal02)
			} else {
				generator.AddExpression("P", strconv.Itoa(varTemp.GetPos()), "+", temporal01)
				generator.GetStack(temporal01, temporal02)
			}
			temInferior := generator.AddTemporal()
			temSuperior := generator.AddTemporal()
			temp05 := generator.AddTemporal()
			temp06 := generator.AddTemporal()
			tempRange := generator.AddTemporal()
			tempChar := generator.AddTemporal()
			tempNewChar := generator.AddTemporal()
			tempSizeThis := generator.AddTemporal()
			tempSizeNew := generator.AddTemporal()
			tempCont := generator.AddTemporal()
			tempFinal := generator.AddTemporal()
			generator.AddAssign(tempFinal, "H")
			labelTrue := generator.AddLabel()
			labelFalse := generator.AddLabel()
			labelOmit := generator.AddLabel()
			labelTrue01 := generator.AddLabel()
			labelFalse01 := generator.AddLabel()
			labelEnd := generator.AddLabel()
			generator.AddExpression(temporal02, "1", "+", temInferior)
			generator.AddExpression(temporal02, "2", "+", temSuperior)
			generator.GetHeap(temInferior, temp05)
			generator.GetHeap(temSuperior, temp06)
			generator.AddExpression(temp06, temp05, "-", tempRange)
			generator.AddIf("(int)"+index.Value.(string), "0", "<", labelTrue01)
			generator.AddIf("(int)"+index.Value.(string), tempRange, ">=", labelTrue01)
			generator.AddGoto(labelFalse01)
			generator.PutLabel(labelTrue01)
			generator.AddPrint("c", "66")
			generator.AddPrint("c", "111")
			generator.AddPrint("c", "117")
			generator.AddPrint("c", "110")
			generator.AddPrint("c", "100")
			generator.AddPrint("c", "115")
			generator.AddPrint("c", "69")
			generator.AddPrint("c", "114")
			generator.AddPrint("c", "114")
			generator.AddPrint("c", "111")
			generator.AddPrint("c", "114")
			generator.Println()
			generator.AddGoto(labelEnd)
			generator.PutLabel(labelFalse01)
			generator.AddExpression(temporal02, "2", "+", tempSizeThis)
			generator.GetHeap(tempSizeThis, tempSizeThis)
			generator.AddExpression(temporal02, "3", "+", tempChar)
			generator.AddExpression(tempSizeThis, "1", "-", tempSizeNew)
			generator.SetHeap("H", "1")
			generator.AddExpression("H", "1", "+", "H")
			generator.SetHeap("H", "0")
			generator.AddExpression("H", "1", "+", "H")
			generator.SetHeap("H", tempSizeNew)
			generator.AddExpression("H", "1", "+", "H")
			generator.PutLabel(labelFalse)
			generator.AddIf(tempCont, "(int)"+tempSizeThis, "==", labelTrue)
			generator.AddIf(tempCont, "(int)"+index.Value.(string), "==", labelOmit)
			generator.GetHeap(tempChar, tempNewChar)
			generator.SetHeap("H", tempNewChar)
			generator.AddExpression("H", "1", "+", "H")
			generator.AddExpression(tempChar, "1", "+", tempChar)
			generator.AddExpression(tempCont, "1", "+", tempCont)
			generator.AddGoto(labelFalse)
			generator.PutLabel(labelOmit)
			generator.AddExpression(tempChar, "1", "+", tempChar)
			generator.AddExpression(tempCont, "1", "+", tempCont)
			generator.AddGoto(labelFalse)
			generator.PutLabel(labelTrue)
			generator.SetStack(strconv.Itoa(varTemp.GetPos()), tempFinal)
			generator.PutLabel(labelEnd)
			return nil

		}
	}
	return nil
}
