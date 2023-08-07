package symbol

type SymbolTable struct {
	previousTable *SymbolTable
	currentTable  map[string]interface{}
}
