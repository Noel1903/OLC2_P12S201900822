package instructions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
)

type IncreDecrem struct {
	identifier string
	operation  string
	value      Abstract.Expression
}

func NewIncreDecrem(identifier string, operation string, value Abstract.Expression) *IncreDecrem {
	return &IncreDecrem{
		identifier: identifier,
		operation:  operation,
		value:      value,
	}
}

func (inde *IncreDecrem) Execute(table Enviorement.SymbolTable) interface{} {
	value := inde.value.GetValue(table)
	variable := table.GetVariable(inde.identifier)
	if value.Type == Enviorement.INT && variable.Type == Enviorement.INT {
		if inde.operation == "+=" {
			valueResult := Enviorement.ReturnSymbol{
				Type:  Enviorement.INT,
				Value: variable.Value.(int) + value.Value.(int),
			}
			table.UpdateVariable(inde.identifier, valueResult)
		}
		if inde.operation == "-=" {
			valueResult := Enviorement.ReturnSymbol{
				Type:  Enviorement.INT,
				Value: variable.Value.(int) - value.Value.(int),
			}
			table.UpdateVariable(inde.identifier, valueResult)
		}

	} else if value.Type == Enviorement.FLOAT && variable.Type == Enviorement.FLOAT {
		if inde.operation == "+=" {
			valueResult := Enviorement.ReturnSymbol{
				Type:  Enviorement.FLOAT,
				Value: variable.Value.(float64) + value.Value.(float64),
			}
			table.UpdateVariable(inde.identifier, valueResult)
		}
		if inde.operation == "-=" {
			valueResult := Enviorement.ReturnSymbol{
				Type:  Enviorement.FLOAT,
				Value: variable.Value.(float64) - value.Value.(float64),
			}
			table.UpdateVariable(inde.identifier, valueResult)
		}
	}

	return nil
}
