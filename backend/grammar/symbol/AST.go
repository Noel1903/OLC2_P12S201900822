package symbol

type AST struct {
	GoblalTable       SymbolTable
	Instructions      []interface{}
	ConsoleOut        string
	reportSymbolTable string
}

func NewAST(instructions []interface{}) *AST {
	return &AST{Instructions: instructions}
}

func (ast *AST) GetConsoleOut() string {
	return ast.ConsoleOut
}

func (ast *AST) SetConsoleOut(consoleOut string) {
	ast.ConsoleOut = consoleOut
}

func (ast *AST) SetGlobalTable(globalTable SymbolTable) {
	ast.GoblalTable = globalTable
}

func (ast *AST) UpdateConsoleOut(consoleOut string) {
	ast.ConsoleOut += consoleOut
}

func (ast *AST) GetGlobalTable() *SymbolTable {
	return &ast.GoblalTable
}

func (ast *AST) GetReportSymbolTable() string {
	return ast.reportSymbolTable
}

func (ast *AST) SetReportSymbolTable(reportSymbolTable string) {
	ast.reportSymbolTable = reportSymbolTable
}

func (ast *AST) UpdateSymbolTable(table string) {
	ast.reportSymbolTable += table
}
