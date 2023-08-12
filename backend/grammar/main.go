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
	code := "print(\"Hola\")\nvar a : Int = 5\nprint(a)\n"
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
