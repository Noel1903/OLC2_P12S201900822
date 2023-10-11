package instructions

import (
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Enviorment "grammar/symbol"
	Generator "grammar/symbol"
	"reflect"
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

	labelJump := generator.AddLabel()
	if condition.Type != Enviorment.BOOL {
		err := Error.NewException("Error: La condicion debe ser de tipo booleana", table.GetName(), i.Line, i.Column)
		return Enviorment.ReturnSymbol{
			Type:  Enviorment.ERROR,
			Value: err,
		}
	}

	newEnviorement := Enviorment.NewEnviorement("if", &table)
	//fmt.Println(condition.LabelTrue)
	for _, labels := range condition.LabelTrue {
		if reflect.TypeOf(labels).Kind() != reflect.Slice || reflect.TypeOf(labels).Elem().Kind() == reflect.Interface {
			generator.PutLabel(labels.(string))

		}
	}
	//generator.AddEnv(strconv.Itoa(table.GetSize()))
	generator.AddEnv(strconv.Itoa(table.GetSizeEnv()))

	for _, instr := range i.codeIf {
		instr.(Abstract.Instruction).Execute(newEnviorement, ast)

	}
	//generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	generator.ReturnEnv(strconv.Itoa(table.GetSizeEnv()))
	generator.AddGoto(labelJump)
	//generator.AddEnv(strconv.Itoa(table.GetSize()))
	newEnviorement01 := Enviorment.NewEnviorement("else", &table)
	for _, labels := range condition.LabelFalse {
		if reflect.TypeOf(labels).Kind() != reflect.Slice || reflect.TypeOf(labels).Elem().Kind() == reflect.Interface {
			generator.PutLabel(labels.(string))

		}

	}
	generator.AddEnv(strconv.Itoa(table.GetSizeEnv()))
	for _, instr := range i.codeElse {
		instr.(Abstract.Instruction).Execute(newEnviorement01, ast)
	}
	//generator.ReturnEnv(strconv.Itoa(table.GetSize()))
	generator.ReturnEnv(strconv.Itoa(table.GetSizeEnv()))
	generator.PutLabel(labelJump)

	return nil
}
