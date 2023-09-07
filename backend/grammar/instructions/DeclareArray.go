package instructions

import (
	"fmt"
	Enviorement "grammar/symbol"
)

type DeclareArray struct {
	Id    string
	Type  Enviorement.TypeData
	Size  int
	Value []interface{}
}

func NewDeclareArray(id string, typeD Enviorement.TypeData, size int, value []interface{}) *DeclareArray {
	return &DeclareArray{
		Id:    id,
		Type:  typeD,
		Size:  size,
		Value: value,
	}
}

func (a *DeclareArray) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	var contM int = 0
	fmt.Println("Declare Array")
	//for _, value := range a.Value {

	//var vector_value = []Enviorement.ReturnSymbol{}
	/*
		for _, valueI := range value {
			value_exp := valueI.(Abstract.Expression).GetValue(table, ast)
			fmt.Println("Value: ", value_exp.Value)
			if value_exp.Type == Enviorement.ARRAY {
				contM++
			}
			/*if value_exp.Type != a.Type {
				fmt.Println("Error: El tipo de dato no coincide")
				return nil
			}
			//vector_value = append(vector_value, value_exp)
		}*/
	/*vector_result := Enviorement.ReturnSymbol{
		Type:  Enviorement.ARRAY,
		Value: vector_value,
	}
	table.SetVariable(a.Id, vector_result, true)*/
	//}
	fmt.Println("Contador: ", contM)
	fmt.Println("Size: ", a.Size)
	return nil
}
