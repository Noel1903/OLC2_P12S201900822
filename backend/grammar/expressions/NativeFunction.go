package expressions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	"strconv"
)

type NativeFunction struct {
	Id         string
	Expression Abstract.Expression
}

func NewNativeFunction(Id string, Expression Abstract.Expression) *NativeFunction {
	return &NativeFunction{Id: Id, Expression: Expression}
}

func (n NativeFunction) Int(table Enviorement.SymbolTable, tree *Enviorement.AST) Enviorement.ReturnSymbol {
	genAux := Enviorement.NewGenerator()
	generator := genAux.GetInstance()

	value := n.Expression.GetValue(table, tree)
	if value.Type == Enviorement.STRING {
		generator.AddEnv(strconv.Itoa(table.GetSizeEnv()))
		generator.AddFuncStringToNumber()
		generator.SetStack("P", value.Value.(string))
		generator.AddCall("stringToNumber")
		temporal := generator.AddTemporal()
		generator.GetStack("P", temporal)
		generator.ReturnEnv(strconv.Itoa(table.GetSizeEnv()))
		return Enviorement.ReturnSymbol{Type: Enviorement.INT, Value: temporal}
	} else if value.Type == Enviorement.FLOAT {
		valueReturn := int(value.Value.(float64))
		return Enviorement.ReturnSymbol{Type: Enviorement.INT, Value: valueReturn}
	} else if value.Type == Enviorement.INT {
		return value
	} else if value.Type == Enviorement.BOOL {
		if value.Value.(bool) {
			return Enviorement.ReturnSymbol{Type: Enviorement.INT, Value: 1}
		} else {
			return Enviorement.ReturnSymbol{Type: Enviorement.INT, Value: 0}
		}
	} else if value.Type == Enviorement.CHAR {
		return Enviorement.ReturnSymbol{Type: Enviorement.INT, Value: int(value.Value.(byte))}
	}
	return Enviorement.ReturnSymbol{}
}

func (n NativeFunction) Float(table Enviorement.SymbolTable, tree *Enviorement.AST) Enviorement.ReturnSymbol {
	genAux := Enviorement.NewGenerator()
	generator := genAux.GetInstance()
	value := n.Expression.GetValue(table, tree)
	if value.Type == Enviorement.STRING {
		generator.AddEnv(strconv.Itoa(table.GetSizeEnv()))
		generator.AddFuncStringToNumber()
		generator.SetStack("P", value.Value.(string))
		generator.AddCall("stringToNumber")
		temporal := generator.AddTemporal()
		generator.GetStack("P", temporal)
		generator.ReturnEnv(strconv.Itoa(table.GetSizeEnv()))
		return Enviorement.ReturnSymbol{Type: Enviorement.FLOAT, Value: temporal}
	} else if value.Type == Enviorement.FLOAT {
		return value
	} else if value.Type == Enviorement.INT {
		valueReturn := float64(value.Value.(int))
		return Enviorement.ReturnSymbol{Type: Enviorement.FLOAT, Value: valueReturn}
	} else if value.Type == Enviorement.BOOL {
		if value.Value.(bool) {
			return Enviorement.ReturnSymbol{Type: Enviorement.FLOAT, Value: 1.0}
		} else {
			return Enviorement.ReturnSymbol{Type: Enviorement.FLOAT, Value: 0.0}
		}
	} else if value.Type == Enviorement.CHAR {
		return Enviorement.ReturnSymbol{Type: Enviorement.FLOAT, Value: float64(value.Value.(byte))}
	}
	return Enviorement.ReturnSymbol{}
}

func (n NativeFunction) String(table Enviorement.SymbolTable, tree *Enviorement.AST) Enviorement.ReturnSymbol {
	value := n.Expression.GetValue(table, tree)
	if value.Type == Enviorement.STRING {
		return value
	} else if value.Type == Enviorement.FLOAT {
		valueReturn := strconv.FormatFloat(value.Value.(float64), 'f', -1, 64)
		return Enviorement.ReturnSymbol{Type: Enviorement.STRING, Value: valueReturn}
	} else if value.Type == Enviorement.INT {
		valueReturn := strconv.Itoa(value.Value.(int))
		return Enviorement.ReturnSymbol{Type: Enviorement.STRING, Value: valueReturn}
	} else if value.Type == Enviorement.BOOL {
		if value.Value.(bool) {
			return Enviorement.ReturnSymbol{Type: Enviorement.STRING, Value: "true"}
		} else {
			return Enviorement.ReturnSymbol{Type: Enviorement.STRING, Value: "false"}
		}
	} else if value.Type == Enviorement.CHAR {
		return Enviorement.ReturnSymbol{Type: Enviorement.STRING, Value: string(value.Value.(byte))}
	}
	return Enviorement.ReturnSymbol{}
}

func (n NativeFunction) GetValue(table Enviorement.SymbolTable, tree *Enviorement.AST) Enviorement.ReturnSymbol {
	switch n.Id {
	case "Int":
		return n.Int(table, tree)
	case "Float":
		return n.Float(table, tree)
	case "String":
		return n.String(table, tree)
	}

	return Enviorement.ReturnSymbol{}
}
