// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgramm
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SwiftgrammParser.
type SwiftgrammVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SwiftgrammParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#sentences.
	VisitSentences(ctx *SentencesContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#sentence.
	VisitSentence(ctx *SentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#declare_var.
	VisitDeclare_var(ctx *Declare_varContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#declare_constant.
	VisitDeclare_constant(ctx *Declare_constantContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#assign_var.
	VisitAssign_var(ctx *Assign_varContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#datatype.
	VisitDatatype(ctx *DatatypeContext) interface{}
}
