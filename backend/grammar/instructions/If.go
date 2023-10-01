package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Enviorment "grammar/symbol"
	Generator "grammar/symbol"
	"strconv"
)

type If struct {
	condition Abstract.Expression
	codeIf    []interface{}
	codeElse  []interface{}
	Line      int
	Column    int
}

func NewIf(condition Abstract.Expression, codeIf []interface{}, codeElse []interface{}, line int, column int) *If {
	return &If{
		condition: condition,
		codeIf:    codeIf,
		codeElse:  codeElse,
		Line:      line,
		Column:    column,
	}
}

func (i *If) Execute(table Enviorment.SymbolTable, ast *Enviorement.AST) interface{} {
	condition := i.condition.GetValue(table, ast)
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	labelTrue := generator.AddLabel()
	labelFalse := generator.AddLabel()
	labelJump := generator.AddLabel()
	if condition.Type != Enviorment.BOOL {
		err := Error.NewException("Error: La condicion debe ser de tipo booleana", table.GetName(), i.Line, i.Column)
		return Enviorment.ReturnSymbol{
			Type:  Enviorment.ERROR,
			Value: err,
		}
	}

	newEnviorement := Enviorment.NewEnviorement("if", &table)
	generator.AddIf(condition.Value.(string), "1", "==", labelTrue)
	generator.AddGoto(labelFalse)
	generator.PutLabel(labelTrue)
	generator.AddEnv(strconv.Itoa(table.GetSize()))

	for _, instr := range i.codeIf {
		instr.(Abstract.Instruction).Execute(newEnviorement, ast)

	}
	generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	generator.AddGoto(labelJump)
	generator.PutLabel(labelFalse)
	generator.AddEnv(strconv.Itoa(table.GetSize()))

	for _, instr := range i.codeElse {
		instr.(Abstract.Instruction).Execute(newEnviorement, ast)
	}
	generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	generator.PutLabel(labelJump)

	return nil
}
