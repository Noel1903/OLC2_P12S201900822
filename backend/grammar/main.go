package main

import (
	"fmt"
	"grammar/abstract"
	"grammar/parser"
	"grammar/symbol"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseSwiftgrammListener
	Code []interface{}
}

func main() {
	//code := "\nprint(\"Hola\"+\"Mundo\")\nprint(!true)\nprint(5!=8)\nprint(\"a\"==\"a\")\nprint(45.10<20.05)"
	//code := "if 10<10{\nprint(\"Hola\")\n} else if 10>10 {\nprint(\"Mundo\")\n}else{\nprint(\"Adios\")\n}"
	//code := "var i:Int = 10\nprint(i)\nlet numero = 85\nprint(numero)\nnumero+=2\nprint(i)"
	//code := "var i:Int = 0\nwhile i<=25{\nprint(i)\ni+=1\nif i<5{\nprint(\"Numero menor a 5\")\n}else{\nbreak\n}\n}"
	//code := "for i in 1...10{\nprint(i)\nif i==5{\nprint(\"Es 5\")\nreturn\n}\n}"
	//code := "var i = 2\nwhile i<=10{\nguard i%2 == 0 else {\ni+=1\ncontinue\n}\nprint(i)\ni+=1\n}"
	//code := "var vector1:[Character] = []\nprint(vector1)"
	code := "func factorial(x:Int,y:Int)->Int{\nif x == 0{\nreturn y+1\n}else if x>0 && y==0{return factorial(x-1,1)\n}else{factorial(x-1,factorial(x,y-1))\n}\n}\nprint(factorial(3,4))\n"
	fmt.Println(code)
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftlex(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSwiftgrammParser(tokens)
	p.BuildParseTrees = true

	tree := p.S()
	//fmt.Println(tree.ToStringTree([]string{}, p))
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	table := symbol.NewEnviorement("global", nil)
	for _, instr := range Code {
		//fmt.Println(instr)
		instr.(abstract.Instruction).Execute(table)
	}

}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}
