package instructions

import (
	abstract "grammar/abstract"
	Generator "grammar/symbol"
	symbol "grammar/symbol"

	"strconv"
)

type Print struct {
	Expression []interface{}
}

func NewPrint(Expression []interface{}) *Print {
	return &Print{Expression: Expression}
}

func (p Print) StringArray(value symbol.ReturnSymbol, table symbol.SymbolTable, ast *symbol.AST) string {
	var print string = ""
	valueR := value.Value

	for _, element := range valueR.([]symbol.ReturnSymbol) {
		if element.Type == symbol.INT {
			print += strconv.Itoa(element.Value.(int))
		} else if element.Type == symbol.FLOAT {
			print += strconv.FormatFloat(element.Value.(float64), 'f', -1, 64)
		} else if element.Type == symbol.CHAR {
			print += string(element.Value.(byte))
		} else if element.Type == symbol.STRING {
			print += element.Value.(string)
		} else if element.Type == symbol.BOOL {
			if element.Value.(bool) {
				print += "true"
			} else {
				print += "false"
			}
		} else if element.Type == symbol.NIL {
			print += "nil"
		}
		print += ","
	}
	if len(print) > 0 {
		print = print[:len(print)-1]
	}

	return "[" + print + "]"
}

func (p Print) String(value symbol.ReturnSymbol) string {
	var print string = ""
	if value.Type == symbol.INT {
		print = strconv.Itoa(value.Value.(int))
	} else if value.Type == symbol.FLOAT {
		print = strconv.FormatFloat(value.Value.(float64), 'f', -1, 64)
	} else if value.Type == symbol.CHAR {
		print = value.Value.(string)
	} else if value.Type == symbol.STRING {
		print = value.Value.(string)
	} else if value.Type == symbol.BOOL {
		if value.Value.(bool) {
			print = "true"
		} else {
			print = "false"
		}
	} else if value.Type == symbol.NIL {
		print = "nil"
	}
	return print
}

func (p Print) Execute(table symbol.SymbolTable, ast *symbol.AST) interface{} {
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	label01 := generator.AddLabel()
	label02 := generator.AddLabel()
	label03 := generator.AddLabel()
	value := p.Expression[0].(abstract.Expression).GetValue(table, ast)

	if value.Type == symbol.INT {
		generator.AddIf(value.Value.(string), "9999999827968.00", "==", label01)
		generator.AddGoto(label02)
		generator.PutLabel(label01)
		generator.AddPrint("c", "110")
		generator.AddPrint("c", "105")
		generator.AddPrint("c", "108")
		generator.Println()
		generator.AddGoto(label03)
		generator.PutLabel(label02)
		generator.AddPrint("d", value.Value.(string))
		generator.Println()
	} else if value.Type == symbol.FLOAT {
		generator.AddIf(value.Value.(string), "9999999827968.00", "==", label01)
		generator.AddGoto(label02)
		generator.PutLabel(label01)
		generator.AddPrint("c", "110")
		generator.AddPrint("c", "105")
		generator.AddPrint("c", "108")
		generator.Println()
		generator.AddGoto(label03)
		generator.PutLabel(label02)
		generator.AddPrint("f", value.Value.(string))
		generator.Println()
	} else if value.Type == symbol.CHAR {
		generator.AddIf(value.Value.(string), "9999999827968.00", "==", label01)
		generator.AddGoto(label02)
		generator.PutLabel(label01)
		generator.AddPrint("c", "110")
		generator.AddPrint("c", "105")
		generator.AddPrint("c", "108")
		generator.Println()
		generator.AddGoto(label03)
		generator.PutLabel(label02)
		generator.AddPrintString()
		generator.AddEnv(strconv.Itoa(table.GetSize()))
		generator.SetStack("P", value.Value.(string))
		generator.AddComment("Aqui se debe llamar a la funcion printString")
		generator.AddCall("printString")
		generator.Println()
		generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	} else if value.Type == symbol.STRING {
		generator.AddIf(value.Value.(string), "9999999827968.00", "==", label01)
		generator.AddGoto(label02)
		generator.PutLabel(label01)
		generator.AddPrint("c", "110")
		generator.AddPrint("c", "105")
		generator.AddPrint("c", "108")
		generator.Println()
		generator.AddGoto(label03)
		generator.PutLabel(label02)
		generator.AddPrintString()
		generator.AddEnv(strconv.Itoa(table.GetSize()))
		generator.SetStack("P", value.Value.(string))
		generator.AddComment("Aqui se debe llamar a la funcion printString")
		generator.AddCall("printString")
		generator.Println()
		generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	} else if value.Type == symbol.BOOL {
		labelExit := generator.AddLabel()
		for _, label := range value.LabelTrue {
			generator.PutLabel(label.(string))
		}
		generator.AddPrint("c", "116")
		generator.AddPrint("c", "114")
		generator.AddPrint("c", "117")
		generator.AddPrint("c", "101")
		generator.AddGoto(labelExit)
		for _, label := range value.LabelFalse {
			generator.PutLabel(label.(string))
		}
		generator.AddPrint("c", "102")
		generator.AddPrint("c", "97")
		generator.AddPrint("c", "108")
		generator.AddPrint("c", "115")
		generator.AddPrint("c", "101")
		generator.PutLabel(labelExit)
		generator.Println()
	} else if value.Type == symbol.ARRAY {
		tempVar := table.GetVar(value.Value.(string))
		temporal01 := generator.AddTemporal()
		temporal02 := generator.AddTemporal()
		tempSizeThis := generator.AddTemporal()
		tempChar := generator.AddTemporal()
		tempNewChar := generator.AddTemporal()
		tempCont := generator.AddTemporal()
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
		generator.AddExpression(temporal02, "3", "+", tempChar)
		generator.PutLabel(labelFalse)
		generator.AddIf(tempCont, "(int)"+tempSizeThis, "==", labelTrue)
		generator.GetHeap(tempChar, tempNewChar)
		if tempVar.TypeArr == symbol.INT {
			generator.AddPrint("d", tempNewChar)
		} else if tempVar.TypeArr == symbol.FLOAT {
			generator.AddPrint("f", tempNewChar)
		} else if tempVar.TypeArr == symbol.CHAR {
			generator.AddPrint("c", tempNewChar)
		} else if tempVar.TypeArr == symbol.STRING {
			generator.AddPrintString()
			generator.AddEnv(strconv.Itoa(table.GetSize()))
			generator.SetStack("P", tempNewChar)
			generator.AddComment("Aqui se debe llamar a la funcion printString")
			generator.AddCall("printString")
			generator.ReturnEnv(strconv.Itoa(table.GetSize()))
		}
		generator.AddExpression(tempChar, "1", "+", tempChar)
		generator.AddExpression(tempCont, "1", "+", tempCont)
		generator.AddIf(tempCont, "(int)"+tempSizeThis, "==", labelFalse)
		generator.AddPrint("c", "44")
		generator.AddGoto(labelFalse)
		generator.PutLabel(labelTrue)
		generator.Println()

	}
	generator.PutLabel(label03)
	return nil

}
