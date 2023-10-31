package expressions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
	"strconv"
)

type GetPosMatrix struct {
	Id     string
	Indexs []interface{}
	Line   int
	Column int
}

func NewGetPosMatrix(id string, indexs []interface{}, line int, column int) *GetPosMatrix {
	return &GetPosMatrix{
		Id:     id,
		Indexs: indexs,
		Line:   line,
		Column: column,
	}
}

func GetPos(indexs interface{}, table Enviorement.SymbolTable, arbol *Enviorement.AST) string {
	value := indexs.(Abstract.Expression).GetValue(table, arbol)
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	Nn := generator.AddTemporal()
	//pointer := generator.AddTemporal()
	generator.AddExpression("2", Nn, "+", Nn)
	//generator.AddAssign(pointer, Nn)
	generator.GetHeap(Nn, Nn)
	tempo := generator.AddTemporal()
	generator.AddExpression(value.Value.(string), Nn, "*", tempo)
	//generator.AddExpression()
	//generator.AddExpression(pointer, Nn, "+", Nn)
	return tempo
}

func (g *GetPosMatrix) GetValue(table Enviorement.SymbolTable, arbol *Enviorement.AST) Enviorement.ReturnSymbol {
	variable := table.GetVariableThis(g.Id)
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if variable.Value == nil {
		err := Error.NewException("Variable no esta declarada", g.Id, g.Line, g.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	//fmt.Println("Indexs : ", g.Indexs)
	varTemp := table.GetVar(g.Id)
	temporal01 := generator.AddTemporal()
	temporal02 := generator.AddTemporal()
	if varTemp.IsGlobal {
		generator.GetStack(strconv.Itoa(varTemp.GetPos()), temporal02)
	} else {
		generator.AddExpression("P", strconv.Itoa(varTemp.GetPos()), "+", temporal01)
		generator.GetStack(temporal01, temporal02)
	}
	var temp string
	generator.GetHeap(temporal02, temporal02)
	for i := 0; i < len(g.Indexs); i++ {
		temp = GetPos(g.Indexs[i], table, arbol)
		if (i + 1) <= (len(g.Indexs) - 1) {
			value := g.Indexs[i+1].(Abstract.Expression).GetValue(table, arbol)
			//fmt.Println("Valor : ", value)
			generator.AddExpression(value.Value.(string), temp, "+", temp)

		}
		if i+1 == len(g.Indexs)-1 {
			break
		}

	}
	//generator.GetHeap(temp, temp)
	labelInit := generator.AddLabel()
	labelEnd := generator.AddLabel()
	labelInter := generator.AddLabel()
	tempInit := generator.AddTemporal()
	tempValue := generator.AddTemporal()
	tempSize := generator.AddTemporal()
	tempValues := generator.AddTemporal()
	tempIndex := generator.AddTemporal()

	generator.AddAssign(tempInit, temporal02)

	generator.PutLabel(labelInit)
	generator.GetHeap(tempInit, tempValue)
	generator.AddIf(tempValue, "-1", "==", labelEnd)
	generator.AddExpression(tempInit, "2", "+", tempSize)
	generator.GetHeap(tempSize, tempSize)
	generator.AddExpression(tempInit, "3", "+", tempValues)
	generator.PutLabel(labelInter)
	generator.AddIf(tempIndex, temp, "==", labelEnd)
	generator.AddExpression(tempValues, "1", "+", tempValues)
	generator.AddExpression(tempIndex, "1", "+", tempIndex)
	generator.AddExpression(tempInit, "1", "+", tempInit)
	generator.AddGoto(labelInter)
	generator.AddIf(tempIndex, tempSize, "==", labelInit)
	generator.PutLabel(labelEnd)
	generator.GetHeap(tempValues, tempValues)

	return Enviorement.ReturnSymbol{
		Type:  Enviorement.INT,
		Value: tempValues,
	}
}
