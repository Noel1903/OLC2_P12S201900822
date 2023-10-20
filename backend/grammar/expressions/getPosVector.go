package expressions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
	"strconv"
)

type getPosVector struct {
	Identifier string
	Expression interface{}
	Line       int
	Column     int
}

func NewGetPosVector(Identifier string, Expression interface{}, line int, column int) *getPosVector {
	return &getPosVector{Identifier: Identifier, Expression: Expression, Line: line, Column: column}
}

func (n getPosVector) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	variable := table.GetVariable(n.Identifier)
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if variable.Value == nil {
		err := Errors.NewException("La variable no ha sido declarada", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	if variable.Type != Enviorement.ARRAY {
		err := Errors.NewException("La variable no es un vector", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	index := n.Expression.(Abstract.Expression).GetValue(table, ast)
	if index.Type != Enviorement.INT {
		err := Errors.NewException("El indice debe ser de tipo entero", table.GetName(), n.Line, n.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
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
	temp01 := generator.AddTemporal()
	temp02 := generator.AddTemporal()
	temp03 := generator.AddTemporal()
	tempPos := generator.AddTemporal()
	tempRes := generator.AddTemporal()
	tempRange := generator.AddTemporal()
	temp05 := generator.AddTemporal()
	temp06 := generator.AddTemporal()
	labelTrue := generator.AddLabel()
	labelFalse := generator.AddLabel()
	labelEnd := generator.AddLabel()
	generator.AddExpression(temporal02, "1", "+", temInferior)
	generator.AddExpression(temporal02, "2", "+", temSuperior)
	generator.GetHeap(temInferior, temp05)
	generator.GetHeap(temSuperior, temp06)
	generator.AddExpression(temp06, temp05, "-", tempRange)
	generator.AddIf("(int)"+index.Value.(string), "0", "<", labelTrue)
	generator.AddIf("(int)"+index.Value.(string), tempRange, ">=", labelTrue)
	generator.AddGoto(labelFalse)
	generator.PutLabel(labelTrue)
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
	generator.PutLabel(labelFalse)
	generator.GetHeap(temInferior, temp01)
	generator.AddExpression(index.Value.(string), temp01, "-", temp02)
	generator.AddExpression(temporal02, "3", "+", temp03)
	generator.AddExpression(temp02, temp03, "+", tempPos)
	generator.GetHeap(tempPos, tempRes)
	generator.PutLabel(labelEnd)
	return Enviorement.ReturnSymbol{
		Type:  varTemp.TypeArr,
		Value: tempRes,
	}
}
