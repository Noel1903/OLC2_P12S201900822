package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
)

type IncreDecrem struct {
	identifier string
	operation  string
	value      Abstract.Expression
	Line       int
	Column     int
}

func NewIncreDecrem(identifier string, operation string, value Abstract.Expression, line int, column int) *IncreDecrem {
	return &IncreDecrem{
		identifier: identifier,
		operation:  operation,
		value:      value,
		Line:       line,
		Column:     column,
	}
}

func (inde *IncreDecrem) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	value := inde.value.GetValue(table, ast)
	variable := table.GetVariable(inde.identifier)
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()

	if variable.Value == nil {
		err := Error.NewException("Variable no declarada", inde.identifier, inde.Line, inde.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	if value.Type == Enviorement.INT && variable.Type == Enviorement.INT {
		if inde.operation == "+=" {
			temporal := generator.AddTemporal()
			generator.AddExpression("(int)"+variable.Value.(string), "(int)"+value.Value.(string), "+", temporal)
			valueResult := Enviorement.ReturnSymbol{
				Type:  Enviorement.INT,
				Value: temporal,
			}
			table.UpdateVariable(inde.identifier, valueResult)
		}
		if inde.operation == "-=" {
			temporal := generator.AddTemporal()
			generator.AddExpression("(int)"+variable.Value.(string), "(int)"+value.Value.(string), "-", temporal)
			valueResult := Enviorement.ReturnSymbol{
				Type:  Enviorement.INT,
				Value: temporal,
			}
			table.UpdateVariable(inde.identifier, valueResult)
		}

	} else if value.Type == Enviorement.FLOAT && variable.Type == Enviorement.FLOAT {
		if inde.operation == "+=" {
			temporal := generator.AddTemporal()
			generator.AddExpression(variable.Value.(string), value.Value.(string), "+", temporal)
			valueResult := Enviorement.ReturnSymbol{
				Type:  Enviorement.INT,
				Value: temporal,
			}
			table.UpdateVariable(inde.identifier, valueResult)
		}
		if inde.operation == "-=" {
			temporal := generator.AddTemporal()
			generator.AddExpression(variable.Value.(string), value.Value.(string), "-", temporal)
			valueResult := Enviorement.ReturnSymbol{
				Type:  Enviorement.INT,
				Value: temporal,
			}
			table.UpdateVariable(inde.identifier, valueResult)
		}
	} else {
		err := Error.NewException("Tipos incompatibles", inde.identifier, inde.Line, inde.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}

	return nil
}
