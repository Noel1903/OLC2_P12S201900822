package instructions

import (
	Error "grammar/exceptions"
	Enviorement "grammar/symbol"
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
	for _, param := range f.Params {
		p := param.(*paramFunction)
		f.ParamsEx = append(f.ParamsEx, p.External)
		f.ParamsIn = append(f.ParamsIn, p.Internal)

	}

	function := Enviorement.ReturnSymbol{
		Type:  f.Type,
		Value: f,
	}
	table.SetVariable(f.Id, function, true, f.Line, f.Column)
	ast.UpdateSymbolTable("<tr><td>" + f.Id + "</td><td>Funcion</td><td>" + strconv.Itoa(int(f.Type)) + "</td><td>" + table.GetName() + "</td><td>" + strconv.Itoa(f.Line) + "</td><td>" + strconv.Itoa(f.Column) + "</td></tr>")
	return nil
}
