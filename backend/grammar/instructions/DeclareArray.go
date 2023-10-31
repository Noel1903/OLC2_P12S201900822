package instructions

import (
	Abstract "grammar/abstract"
	"grammar/expressions"
	Enviorement "grammar/symbol"
	Generator "grammar/symbol"
	"reflect"
	"strconv"
)

type DeclareArray struct {
	Id    string
	Type  Enviorement.TypeData
	Value []interface{}
}

func NewDeclareArray(id string, typeD Enviorement.TypeData, value []interface{}) *DeclareArray {
	return &DeclareArray{
		Id:    id,
		Type:  typeD,
		Value: value,
	}
}

func RowsData(value []interface{}, contRow int, table Enviorement.SymbolTable, ast *Enviorement.AST) {
	//var data interface{}
	var contData int = 0
	var data []interface{}

	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	for _, v := range value {
		if reflect.TypeOf(v.(expressions.Native).TypeR) == reflect.TypeOf(Enviorement.ARRAY) && reflect.TypeOf(v.(expressions.Native).Value) == reflect.TypeOf([]interface{}{}) {
			//fmt.Println("Es un array")
			//fmt.Println(v.(expressions.Native).Value)
			RowsData(v.(expressions.Native).Value.([]interface{}), contRow, table, ast)
		} else {
			valueIndex := v.(Abstract.Expression).GetValue(table, ast)
			data = append(data, valueIndex.Value.(string))
			contData++
		}
	}
	if data != nil {
		generator.SetHeap("H", "0")
		generator.AddExpression("H", "1", "+", "H")
		generator.SetHeap("H", strconv.Itoa(contData-1))
		generator.AddExpression("H", "1", "+", "H")
		generator.SetHeap("H", strconv.Itoa(contData))
		generator.AddExpression("H", "1", "+", "H")
		for _, value := range data {

			generator.SetHeap("H", value.(string))
			generator.AddExpression("H", "1", "+", "H")
		}
	}

}

func (a *DeclareArray) Execute(table Enviorement.SymbolTable, ast *Enviorement.AST) interface{} {
	var contRow int = 0
	//var contCol int = 0
	genAux := Generator.NewGenerator()
	generator := genAux.GetInstance()
	//fmt.Println("Valor : ", a.Value)
	temporalInitial := generator.AddTemporal()
	generator.AddAssign(temporalInitial, "H")
	RowsData(a.Value, contRow, table, ast)
	var tempPos string
	value := Enviorement.ReturnSymbol{
		Type:  Enviorement.ARRAY,
		Value: a.Value,
	}
	result := table.SetArray(a.Id, value, a.Type, true, 1, 1, false)

	generator.SetHeap("H", "-1")
	generator.AddExpression("H", "1", "+", "H")
	tempPos = strconv.Itoa(result.GetPos())
	if !result.GetIsGlobal() {
		tempPos = generator.AddTemporal()
		generator.AddExpression("P", strconv.Itoa(result.GetPos()), "+", tempPos)
	}
	generator.SetStack(tempPos, temporalInitial)
	return nil
}
