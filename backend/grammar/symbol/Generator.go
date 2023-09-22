package symbol

import (
	"strconv"
	"sync"
)

type Generator struct {
	contTemp        int
	contLabel       int
	temporals       []string
	code            string
	funcs           string
	natives         string
	inFunc          bool
	inNative        bool
	printStringFunc bool
}

func (g *Generator) ClearAll() {
	g.contTemp = 0
	g.contLabel = 0
	g.temporals = nil
	g.code = ""
	g.funcs = ""
	g.natives = ""
	g.inFunc = false
	g.inNative = false
	g.printStringFunc = false

}

func NewGenerator() *Generator {
	return &Generator{
		contTemp:        0,
		contLabel:       0,
		temporals:       nil,
		code:            "",
		funcs:           "",
		natives:         "",
		inFunc:          false,
		inNative:        false,
		printStringFunc: false,
	}
}

var instance *Generator
var once sync.Once

func (g *Generator) GetInstance() *Generator {
	once.Do(func() {
		instance = NewGenerator()
	})
	return instance
}

func Imports() string {
	return "#include <stdio.h>\n"
}
func (g *Generator) GetTemporals() []string {
	return g.temporals
}

func (g *Generator) Header() string {
	code := "#include <stdio.h>\n"
	if len(g.temporals) > 0 {
		code += "double "
		for i := 0; i < len(g.temporals); i++ {
			code += g.temporals[i]
			if i != len(g.temporals)-1 {
				code += ", "
			}
		}

		code += ";\n"
	}
	code += "double P,H;\n"
	code += "double Stack[1000000];\n"
	code += "double Heap[1000000];\n"
	return code
}

// **************************************COMMENT**************************************
func (g *Generator) AddComment(comment string) {
	g.CodeIn("//" + comment + "\n")
}

func (g *Generator) GetCode() string {

	return g.Header() + g.funcs + g.natives + "\nint main(){\n" + g.code + "\nreturn 0;\n}"
}

//*************************************PRINT*************************************

func (g *Generator) AddPrint(type_ string, value string) {
	if type_ == "c" {
		g.CodeIn("printf(\"%" + type_ + "\",(char)" + value + ");\n")
	} else if type_ == "f" {
		g.CodeIn("printf(\"%" + type_ + "\"," + value + ");\n")
	} else if type_ == "d" {
		g.CodeIn("printf(\"%" + type_ + "\",(int)" + value + ");\n")
	}

}

// *************************************PRINT STRING*************************************
func (g *Generator) AddPrintString() {
	if g.printStringFunc {
		return
	}
	g.printStringFunc = true
	g.inNative = true
	g.AddFunc("printString")
	labelCompare := g.AddLabel()
	labelJump := g.AddLabel()
	temp1 := g.AddTemporal()
	temp2 := g.AddTemporal()
	g.GetStack("P", temp1)
	g.PutLabel(labelJump)
	g.GetHeap(temp1, temp2)
	g.AddIf(temp2, "-1", "==", labelCompare)
	g.AddPrint("c", temp2)
	g.AddExpression(temp1, "1", "+", temp1)
	g.AddGoto(labelJump)
	g.PutLabel(labelCompare)
	g.CodeIn("return;\n}\n")
	g.inNative = false
}

// ****************************************Insersión de código****************************************
func (g *Generator) CodeIn(code string) {
	if g.inNative {
		g.natives = g.natives + "\t" + code
	} else if g.inFunc {
		g.funcs = g.funcs + "\t" + code
	} else {
		g.code = g.code + "\t" + code
	}

}

// ***************************************Añadir Temporales y Labels***************************************
func (g *Generator) AddTemporal() string {
	temp := "t" + strconv.Itoa(g.contTemp)
	g.contTemp++
	g.temporals = append(g.temporals, temp)
	return temp
}

func (g *Generator) AddLabel() string {
	label := "L" + strconv.Itoa(g.contLabel)
	g.contLabel++
	return label
}

func (g *Generator) AddGoto(label string) {
	g.CodeIn("goto " + label + ";\n")
}

// ****************************************Insertar Labels****************************************
func (g *Generator) PutLabel(label string) {
	g.CodeIn(label + ":\n")
}

// *****************************************Expresiones*****************************************
func (g *Generator) AddExpression(left string, right string, operator string, result string) {
	g.CodeIn(result + " = " + left + " " + operator + " " + right + ";\n")
}

func (g *Generator) AddAssign(id string, value string) {
	g.CodeIn(id + " = " + value + ";\n")
}

// *****************************************STACK*****************************************
func (g *Generator) SetStack(position string, value string) {
	g.CodeIn("Stack[(int)" + position + "] = " + value + ";\n")
}

func (g *Generator) GetStack(position string, result string) {
	g.CodeIn(result + " = Stack[(int)" + position + "];\n")
}

// ******************************************HEAP******************************************
func (g *Generator) SetHeap(position string, value string) {
	g.CodeIn("Heap[(int)" + position + "] = " + value + ";\n")
}

func (g *Generator) GetHeap(position string, result string) {
	g.CodeIn(result + " = Heap[(int)" + position + "];\n")
}

func (g *Generator) NextHeap() {
	g.CodeIn("H = H + 1;\n")
}

// ****************************************ENVIOREMENT****************************************
func (g *Generator) AddEnv(env string) {
	g.CodeIn("P = P + " + env + ";\n")
}

func (g *Generator) ReturnEnv(env string) {
	g.CodeIn("P = P - " + env + ";\n")
}

// *******************************************Funciones*******************************************
func (g *Generator) AddFunc(id string) {

	if !g.inNative {
		g.inFunc = true
	}
	g.CodeIn("void " + id + "(){\n")
}

func (g *Generator) EndFunc() {
	g.CodeIn("return;\n}\n")
	g.inFunc = false
}

func (g *Generator) AddCall(id string) {
	g.CodeIn(id + "();\n")
}

// ***************************************IF***************************************
func (g *Generator) AddIf(left string, right string, operator string, label string) {
	g.CodeIn("if(" + left + " " + operator + " " + right + ") goto " + label + ";\n")
}
