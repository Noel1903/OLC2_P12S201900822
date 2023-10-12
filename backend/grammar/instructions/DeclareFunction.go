package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
	"reflect"
	"strconv"
)

type DeclareFunction struct {
	Id       string
	Type     Enviorement.TypeData
	ParamsIn []interface{}
	ParamsEx []interface{}
	Params   []interface{}
	codeF    []interface{}
	Line     int
	Column   int
}

func NewDeclareFunction(id string, typeD Enviorement.TypeData, Params []interface{}, codeF []interface{}, line int, column int) *DeclareFunction {
	return &DeclareFunction{
		Id:     id,
		Type:   typeD,
		Params: Params,
		codeF:  codeF,
		Line:   line,
		Column: column,
	}
}

func (f *DeclareFunction) GetSentences() []interface{} {
	return f.codeF
}

func (f *DeclareFunction) GetParamsIn() []interface{} {
	return f.ParamsIn
}

func (f *DeclareFunction) GetParamsEx() []interface{} {
	return f.ParamsEx
}

func (f *DeclareFunction) GetId() string {
	return f.Id
}

func (f *DeclareFunction) GetType() Enviorement.TypeData {
	return f.Type
}

func (f *DeclareFunction) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {

	variable := table.GetVariableThis(f.Id)
	if variable.Value != nil {
		err := Error.NewException("La funcion ya existe", table.GetName(), f.Line, f.Column)
		return Enviorement.ReturnSymbol{
			Type:  Enviorement.ERROR,
			Value: err,
		}
	}
	var paramsSave []interface{}
	for _, param := range f.Params {
		p := param.(*paramFunction)
		f.ParamsEx = append(f.ParamsEx, p.External)
		f.ParamsIn = append(f.ParamsIn, p.Internal)
		paramsSave = append(paramsSave, p.Internal)

	}
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	function := Enviorement.ReturnSymbol{
		Type:  f.Type,
		Value: f,
	}
	//tableGlobal := ast.GetGlobalTable()
	table.SetFunction(f.Id, function, true, f.Line, f.Column, false)

	newEnviorement := Enviorement.NewEnviorement(f.Id, &table)
	//fmt.Println(newEnviorement)
	for _, param := range paramsSave {
		temp := generator.AddTemporal()
		newEnviorement.SetVariable(param.(*DeclareVariable).Identifier, Enviorement.ReturnSymbol{Type: param.(*DeclareVariable).TypeD, Value: temp}, true, param.(*DeclareVariable).Line, param.(*DeclareVariable).Column, false)
	}
	generator.AddFunc(f.Id)
	labelExit := generator.AddLabel()
	for _, instr := range f.codeF {

		instr.(Abstract.Instruction).Execute(newEnviorement, ast)

		if reflect.TypeOf(instr) == reflect.TypeOf(Return{}) {
			fmt.Println("Return detectado")
			generator.AddGoto(labelExit)
		}
	}
	generator.PutLabel(labelExit)
	generator.EndFunc()

	ast.UpdateSymbolTable("<tr><td>" + f.Id + "</td><td>Funcion</td><td>" + strconv.Itoa(int(f.Type)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(f.Line) + "</td><td>" + strconv.Itoa(f.Column) + "</td></tr>")
	return nil
}
