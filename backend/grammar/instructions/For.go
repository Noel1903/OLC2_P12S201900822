package instructions

import (
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
	"strconv"
)

type For struct {
	identifier string
	Initial    Abstract.Expression
	Final      Abstract.Expression
	codeFor    []interface{}
	Line       int
	Column     int
}

func NewFor(identifier string, Initial Abstract.Expression, Final Abstract.Expression, codeFor []interface{}, line int, column int) *For {
	return &For{
		identifier: identifier,
		Initial:    Initial,
		Final:      Final,
		codeFor:    codeFor,
		Line:       line,
		Column:     column,
	}
}

func (f *For) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	valueInitial := f.Initial.GetValue(table, ast)
	valueFinal := f.Final.GetValue(table, ast)

	value, _ := strconv.Atoi(valueInitial.Value.(string))
	newVariable := Enviorement.ReturnSymbol{
		Type:  Enviorement.INT,
		Value: value,
	}
	genAux := Enviorement.NewGenerator()
	generator := genAux.GetInstance()

	newLabel := generator.AddLabel()
	exitLabel := generator.AddLabel()

	generator.AddEnv(strconv.Itoa(table.GetSizeEnv()))
	newEnviorement := Enviorement.NewEnviorement("for", &table)
	variableChange := newEnviorement.SetVariable(f.identifier, valueInitial, true, f.Line, f.Column, false)
	tempPos := generator.AddTemporal()
	generator.AddExpression("P", strconv.Itoa(variableChange.Position), "+", tempPos)
	generator.SetStack(tempPos, valueInitial.Value.(string))
	generator.PutLabel(newLabel)
	for _, instr := range f.codeFor {
		instr.(Abstract.Instruction).Execute(newEnviorement, ast)
	}
	newVariable.Value = newVariable.Value.(int) + 1
	newEnviorement.UpdateVariable(f.identifier, newVariable)
	temp := generator.AddTemporal()
	generator.GetStack(strconv.Itoa(variableChange.Position), temp)
	generator.AddExpression(temp, "1", "+", temp)
	generator.SetStack(strconv.Itoa(variableChange.Position), temp)
	generator.AddIf(temp, valueFinal.Value.(string), "<=", newLabel)
	generator.AddGoto(exitLabel)
	generator.PutLabel(exitLabel)
	generator.ReturnEnv(strconv.Itoa(table.GetSizeEnv()))

	return nil
}
