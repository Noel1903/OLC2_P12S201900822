package expressions

import (
	Abstract "grammar/abstract"
	Errors "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
)

type LogicOperations struct {
	left   Abstract.Expression
	right  Abstract.Expression
	op     string
	Line   int
	Column int
}

func NewLogicOperations(left Abstract.Expression, right Abstract.Expression, op string, line int, column int) LogicOperations {
	return LogicOperations{
		left:   left,
		right:  right,
		op:     op,
		Line:   line,
		Column: column,
	}
}

func And(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()

	if typeleft == Enviorement.BOOL && typeright == Enviorement.BOOL {
		temporal := generator.AddTemporal()
		labeltrue := generator.AddLabel()
		labelfalse := generator.AddLabel()
		labelJump := generator.AddLabel()

		generator.AddIf(left.(string), right.(string), "&&", labeltrue)
		generator.AddGoto(labelfalse)
		generator.PutLabel(labeltrue)
		generator.AddAssign(temporal, "1")
		generator.AddGoto(labelJump)
		generator.PutLabel(labelfalse)
		generator.AddAssign(temporal, "0")
		generator.PutLabel(labelJump)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: temporal,
		}
	} else {
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: nil,
		}
	}
	return result
}

func Or(left interface{}, right interface{}, typeleft Enviorement.TypeData, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	if typeleft == Enviorement.BOOL && typeright == Enviorement.BOOL {
		temporal := generator.AddTemporal()
		labeltrue := generator.AddLabel()
		labelfalse := generator.AddLabel()
		labelJump := generator.AddLabel()

		generator.AddIf(left.(string), right.(string), "||", labeltrue)
		generator.AddGoto(labelfalse)
		generator.PutLabel(labeltrue)
		generator.AddAssign(temporal, "1")
		generator.AddGoto(labelJump)
		generator.PutLabel(labelfalse)
		generator.AddAssign(temporal, "0")
		generator.PutLabel(labelJump)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: temporal,
		}
	} else {
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: nil,
		}
	}
	return result
}

func Not(right interface{}, typeright Enviorement.TypeData) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()

	if typeright == Enviorement.BOOL {
		temporal := generator.AddTemporal()
		labeltrue := generator.AddLabel()
		labelfalse := generator.AddLabel()
		labelJump := generator.AddLabel()
		generator.AddIf(right.(string), "1", "==", labeltrue)
		generator.AddGoto(labelfalse)
		generator.PutLabel(labeltrue)
		generator.AddAssign(temporal, "0")
		generator.AddGoto(labelJump)
		generator.PutLabel(labelfalse)
		generator.AddAssign(temporal, "1")
		generator.PutLabel(labelJump)
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.BOOL,
			Value: temporal,
		}
	} else {
		result = Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: nil,
		}
	}
	return result
}

func (l LogicOperations) GetValue(table Enviorement.SymbolTable, ast *Enviorement.AST) Enviorement.ReturnSymbol {
	var result Enviorement.ReturnSymbol
	if l.left != nil {
		expleft := l.left.GetValue(table, ast).Value
		expright := l.right.GetValue(table, ast).Value
		typeleft := l.left.GetValue(table, ast).Type
		typeright := l.right.GetValue(table, ast).Type

		switch l.op {
		case "&&":
			result = And(expleft, expright, typeleft, typeright)
		case "||":
			result = Or(expleft, expright, typeleft, typeright)
		default:
			result = Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: nil,
			}
		}
	} else {
		expright := l.right.GetValue(table, ast).Value
		typeright := l.right.GetValue(table, ast).Type
		switch l.op {
		case "!":
			result = Not(expright, typeright)

		default:
			result = Enviorement.ReturnSymbol{
				Type:  Enviorement.ERROR,
				Value: nil,
			}
		}
	}

	if result.Type == Enviorement.ERROR {
		result.Value = Errors.NewException("Error en la operacion aritmetica", table.GetName(), l.Line, l.Column)
	}
	return result
}
