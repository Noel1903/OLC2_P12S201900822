package symbol

import "fmt"

type SymbolTable struct {
	name          string
	previousTable *SymbolTable
	currentTable  map[string]Symbol
	Position      int
}

func (t *SymbolTable) GetName() string {
	return t.name
}

func (t *SymbolTable) GetSize() int {

	if t.previousTable != nil {
		return t.previousTable.GetSize()
	}
	return t.Position
}

func (t *SymbolTable) SetSize(size int) {
	t.Position = size + 1
}

func NewEnviorement(name string, previous *SymbolTable) SymbolTable {
	position := 0

	if previous != nil {
		position = previous.GetSize()
	}

	env := SymbolTable{
		name:          name,
		previousTable: previous,
		currentTable:  make(map[string]Symbol),
		Position:      position,
	}
	return env
}

func (table *SymbolTable) SetVariable(id string, value ReturnSymbol, declare bool, line int, column int, InHeap bool) Symbol {
	if declare {

		valueSymbol := Symbol{
			Value:    value,
			Type:     MUTABLE,
			Line:     line,
			Column:   column,
			Position: table.GetSize(),
			IsGlobal: table.previousTable == nil,
			InHeap:   InHeap,
		}
		valueSymbol.SetPos(table.GetSize())
		table.currentTable[id] = valueSymbol
		table.SetSize(table.GetSize())

	} else {
		valueSymbol := Symbol{
			Value:  value,
			Type:   UNMUTABLE,
			Line:   line,
			Column: column,
		}
		table.currentTable[id] = valueSymbol
	}
	return table.currentTable[id]
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

func (table *SymbolTable) GetVar(id string) Symbol {
	variable := table.currentTable[id]
	if variable.Value.Value != nil {
		return variable
	} else {
		if table.previousTable != nil {
			return table.previousTable.GetVar(id)
		} else {
			return Symbol{Type: 0, Value: ReturnSymbol{Type: UNDEFINED, Value: nil}}
		}
	}
}

func (table *SymbolTable) GetVariableThis(id string) ReturnSymbol {
	variable := table.currentTable[id].Value
	if variable.Value != nil {
		return variable
	} else {
		return ReturnSymbol{Type: UNDEFINED, Value: nil}
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
