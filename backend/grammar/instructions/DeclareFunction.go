package instructions

import (
	Enviorement "grammar/symbol"
)

type DeclareFunction struct {
	Id     string
	Type   Enviorement.TypeData
	Params []interface{}
	codeF  []interface{}
}

func NewDeclareFunction(id string, typeD Enviorement.TypeData, params []interface{}, codeF []interface{}) *DeclareFunction {
	return &DeclareFunction{
		Id:     id,
		Type:   typeD,
		Params: params,
		codeF:  codeF,
	}
}

func (f *DeclareFunction) GetSentences() []interface{} {
	return f.codeF
}

func (f *DeclareFunction) GetParams() []interface{} {
	return f.Params
}

func (f *DeclareFunction) GetId() string {
	return f.Id
}

func (f *DeclareFunction) GetType() Enviorement.TypeData {
	return f.Type
}

func (f *DeclareFunction) Execute(table Enviorement.SymbolTable) interface{} {
	function := Enviorement.ReturnSymbol{
		Type:  f.Type,
		Value: f,
	}
	table.SetVariable(f.Id, function, true)
	return nil
}
