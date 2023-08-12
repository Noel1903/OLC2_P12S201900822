package symbol

type SymbolTable struct {
	name          string
	previousTable *SymbolTable
	currentTable  map[string]interface{}
}

func NewEnviorement(name string, previous *SymbolTable) SymbolTable {
	table := make(map[string]interface{})
	env := SymbolTable{name, previous, table}
	env.currentTable = make(map[string]interface{})
	return env
}

func (table *SymbolTable) SetVariable(id string, value interface{}) {
	table.currentTable[id] = value
}

func (table *SymbolTable) GetVariable(id string) interface{} {
	variable := table.currentTable[id]
	if variable != nil {
		return variable
	} else {
		if table.previousTable != nil {
			return table.previousTable.GetVariable(id)
		} else {
			return nil
		}
	}
}
