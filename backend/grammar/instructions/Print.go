package instructions

import (
	abstract "grammar/abstract"
	Generator "grammar/symbol"
	symbol "grammar/symbol"

	"strconv"
)

type Print struct {
	Expression []interface{}
}

func NewPrint(Expression []interface{}) *Print {
	return &Print{Expression: Expression}
}

func (p Print) StringArray(value symbol.ReturnSymbol, table symbol.SymbolTable, ast *symbol.AST) string {
	var print string = ""
	valueR := value.Value

	for _, element := range valueR.([]symbol.ReturnSymbol) {
		if element.Type == symbol.INT {
			print += strconv.Itoa(element.Value.(int))
		} else if element.Type == symbol.FLOAT {
			print += strconv.FormatFloat(element.Value.(float64), 'f', -1, 64)
		} else if element.Type == symbol.CHAR {
			print += string(element.Value.(byte))
		} else if element.Type == symbol.STRING {
			print += element.Value.(string)
		} else if element.Type == symbol.BOOL {
			if element.Value.(bool) {
				print += "true"
			} else {
				print += "false"
			}
		} else if element.Type == symbol.NIL {
			print += "nil"
		}
		print += ","
	}
	if len(print) > 0 {
		print = print[:len(print)-1]
	}

	return "[" + print + "]"
}

func (p Print) String(value symbol.ReturnSymbol) string {
	var print string = ""
	if value.Type == symbol.INT {
		print = strconv.Itoa(value.Value.(int))
	} else if value.Type == symbol.FLOAT {
		print = strconv.FormatFloat(value.Value.(float64), 'f', -1, 64)
	} else if value.Type == symbol.CHAR {
		print = value.Value.(string)
	} else if value.Type == symbol.STRING {
		print = value.Value.(string)
	} else if value.Type == symbol.BOOL {
		if value.Value.(bool) {
			print = "true"
		} else {
			print = "false"
		}
	} else if value.Type == symbol.NIL {
		print = "nil"
	}
	return print
}

func (p Print) Execute(table symbol.SymbolTable, ast *symbol.AST) interface{} {
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	value := p.Expression[0].(abstract.Expression).GetValue(table, ast)
	if value.Type == symbol.INT {
		generator.AddPrint("d", value.Value.(string))
		generator.Println()
	} else if value.Type == symbol.FLOAT {
		generator.AddPrint("f", value.Value.(string))
		generator.Println()
	} else if value.Type == symbol.CHAR {
		generator.AddPrintString()
		generator.AddEnv(strconv.Itoa(table.GetSize()))
		generator.SetStack("P", value.Value.(string))
		generator.AddComment("Aqui se debe llamar a la funcion printString")
		generator.Println()
		generator.AddCall("printString")
		generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	} else if value.Type == symbol.STRING {
		generator.AddPrintString()
		generator.AddEnv(strconv.Itoa(table.GetSize()))
		generator.SetStack("P", value.Value.(string))
		generator.AddComment("Aqui se debe llamar a la funcion printString")
		generator.AddCall("printString")
		generator.Println()
		generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	} else if value.Type == symbol.BOOL {
		generator.AddPrint("d", value.Value.(string))
		generator.Println()

	}
	return nil
	/*
		if len(p.Expression) == 0 {
			ast.UpdateConsoleOut("\n")
			return nil
		} else if len(p.Expression) == 1 {
			value := p.Expression[0].(abstract.Expression).GetValue(table, ast)

			if value.Type == symbol.ERROR {
				return value
			}
			var print string = ""

			if value.Type == symbol.ARRAY {
				print = p.StringArray(value, table, ast)
				ast.UpdateConsoleOut(print + "\n")
				return nil
			} else if value.Value == nil {
				ast.UpdateConsoleOut("nil\n")
				return nil
			} else {
				print = p.String(value)
				ast.UpdateConsoleOut(print + "\n")
				return nil
			}

		} else if len(p.Expression) > 1 {
			for _, expression := range p.Expression {
				value := expression.(abstract.Expression).GetValue(table, ast)
				var print string = ""
				if value.Type == symbol.CHAR {
					print = string(value.Value.(byte))
					ast.UpdateConsoleOut(print + " ")

				}
				if value.Type == symbol.ARRAY {
					print = p.StringArray(value, table, ast)
					ast.UpdateConsoleOut(print + " ")

				} else if value.Value == nil {
					ast.UpdateConsoleOut("nil")
				} else {
					print = p.String(value)
					ast.UpdateConsoleOut(print + " ")

				}
			}
			ast.UpdateConsoleOut("\n")
			return nil
		} else if p.Expression == nil {
			ast.UpdateConsoleOut("nil\n")
		}
		return nil*/
}
