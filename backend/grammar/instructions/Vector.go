package instructions

import (
	"fmt"
	Abstract "grammar/abstract"
	Enviorement "grammar/symbol"
)

type Vector struct {
	Id    string
	Type  Enviorement.TypeData
	Value []interface{}
}

func NewVector(id string, typeD Enviorement.TypeData, value []interface{}) *Vector {
	return &Vector{
		Id:    id,
		Type:  typeD,
		Value: value,
	}
}

func (v *Vector) Execute(table Enviorement.SymbolTable) interface{} {
	var vector_value = []Enviorement.ReturnSymbol{}

	single_value := v.Value
	if single_value == nil {
		vector_result := Enviorement.ReturnSymbol{
			Type:  Enviorement.ARRAY,
			Value: vector_value,
		}
		table.SetVariable(v.Id, vector_result, true)

	} else {

		for _, value := range single_value {
			value_exp := value.(Abstract.Expression).GetValue(table)
			if value_exp.Type != v.Type {
				fmt.Println("Error: El tipo de dato no coincide")
				return nil
			}
			vector_value = append(vector_value, value_exp)
		}
		vector_result := Enviorement.ReturnSymbol{
			Type:  Enviorement.ARRAY,
			Value: vector_value,
		}
		table.SetVariable(v.Id, vector_result, true)

	}

	return nil
}
