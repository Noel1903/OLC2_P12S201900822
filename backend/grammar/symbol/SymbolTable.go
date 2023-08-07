package symbol

type SymbolTable struct {
	name          string
	previousTable *SymbolTable
	currentTable  map[string]interface{}
}

func newScoope(name string, previous *SymbolTable) SymbolTable {
	table := make(map[string]interface{})
	env := SymbolTable{name, previous, table}
	env.currentTable = make(map[string]interface{})
	return env
}

func (table *SymbolTable) setVariable(id string, value interface{}) {
	table.currentTable[id] = value
}
