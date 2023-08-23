package symbol

import "fmt"

type SymbolTable struct {
	name          string
	previousTable *SymbolTable
	currentTable  map[string]Symbol
}

func NewEnviorement(name string, previous *SymbolTable) SymbolTable {
	env := SymbolTable{
		name:          name,
		previousTable: previous,
		currentTable:  make(map[string]Symbol),
	}
	return env
}

func (table *SymbolTable) SetVariable(id string, value ReturnSymbol, declare bool) {
	if declare {
		valueSymbol := Symbol{
			Value: value,
			Type:  MUTABLE,
		}
		table.currentTable[id] = valueSymbol
	} else {
		valueSymbol := Symbol{
			Value: value,
			Type:  UNMUTABLE,
		}
		table.currentTable[id] = valueSymbol
	}

}

func (table *SymbolTable) GetVariable(id string) ReturnSymbol {
	variable := table.currentTable[id].Value
	if variable.Value != nil {
		return variable
	} else {
		if table.previousTable != nil {
			return table.previousTable.GetVariable(id)
		} else {
			return ReturnSymbol{Type: UNDEFINED, Value: nil}
		}
	}
}

func (table *SymbolTable) UpdateVariable(id string, value ReturnSymbol) interface{} {
	variable := table.currentTable[id].Value
	typevar := table.currentTable[id].Type
	if variable.Value != nil {
		if typevar == MUTABLE {
			newValue := Symbol{
				Value: value,
				Type:  MUTABLE,
			}
			table.currentTable[id] = newValue
		} else {
			fmt.Println("Error, la variable es inmutable")
			return nil
		}
	} else {
		if table.previousTable != nil {
			return table.previousTable.UpdateVariable(id, value)
		} else {
			return Symbol{Type: 0, Value: ReturnSymbol{Type: UNDEFINED, Value: nil}}
		}
	}
	return nil
}
