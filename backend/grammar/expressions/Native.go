package expressions

import (
	Generator "grammar/symbol"
	Scoope "grammar/symbol"
	"strconv"
)

type Native struct {
	//Fila - columna
	Value  interface{}
	TypeR  Scoope.TypeData
	Line   int
	Column int
}

func (n Native) GetValue(env Scoope.SymbolTable, ast *Scoope.AST) Scoope.ReturnSymbol {
	genAuxiliar := Generator.NewGenerator()
	generator := genAuxiliar.GetInstance()

	if n.TypeR == Scoope.INT {
		valueString := strconv.Itoa(n.Value.(int))
		return Scoope.ReturnSymbol{
			Type:  n.TypeR,
			Value: valueString,
		}
	} else if n.TypeR == Scoope.FLOAT {
		valueString := strconv.FormatFloat(n.Value.(float64), 'f', -1, 64)
		return Scoope.ReturnSymbol{
			Type:  n.TypeR,
			Value: valueString,
		}
	} else if n.TypeR == Scoope.STRING {
		temporal := generator.AddTemporal()
		generator.AddAssign(temporal, "H")
		for char := range n.Value.(string) {
			generator.SetHeap("H", strconv.Itoa(int(n.Value.(string)[char])))
			generator.NextHeap()
		}
		generator.SetHeap("H", "-1")
		generator.NextHeap()
		return Scoope.ReturnSymbol{
			Type:  n.TypeR,
			Value: temporal,
		}

	} else if n.TypeR == Scoope.BOOL {
		labelTrue := generator.AddLabel()
		labelFalse := generator.AddLabel()

		if n.Value.(bool) {
			generator.AddGoto(labelTrue)
		} else {
			generator.AddGoto(labelFalse)
		}
		returnS := Scoope.ReturnSymbol{
			Type:  n.TypeR,
			Value: "",
		}
		returnS.LabelTrue = append(returnS.LabelTrue, labelTrue)
		returnS.LabelFalse = append(returnS.LabelFalse, labelFalse)
		return returnS

	} else if n.TypeR == Scoope.CHAR {
		temporal := generator.AddTemporal()
		generator.AddAssign(temporal, "H")
		generator.SetHeap("H", strconv.Itoa(int(n.Value.(string)[0])))
		generator.NextHeap()
		generator.SetHeap("H", "-1")
		generator.NextHeap()
		return Scoope.ReturnSymbol{
			Type:  n.TypeR,
			Value: temporal,
		}
	} else if n.TypeR == Scoope.NIL {
		temporal := generator.AddTemporal()
		generator.AddAssign(temporal, "9999999827968.00")
		return Scoope.ReturnSymbol{
			Type:  n.TypeR,
			Value: temporal,
		}
	}

	return Scoope.ReturnSymbol{}
}

func NewNative(value interface{}, typeR Scoope.TypeData, line int, column int) Native {
	s := Native{value, typeR, line, column}
	return s
}
