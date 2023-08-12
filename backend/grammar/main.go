package main

import (
	"grammar/abstract"
	"grammar/parser"
	"grammar/symbol"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseSwiftgrammListener
	Code []interface{}
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func main() {
	code := "var a:Int = 5"
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftlex(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSwiftgrammParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	table := symbol.NewEnviorement("global", nil)
	for _, instr := range Code {
		instr.(abstract.Instruction).Execute(table)
	}

}
