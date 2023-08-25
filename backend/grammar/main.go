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
	code := "for i in 1...10{\nprint(i)\nif i==5{\nprint(\"Es 5\")\nreturn\n}\n}"
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
