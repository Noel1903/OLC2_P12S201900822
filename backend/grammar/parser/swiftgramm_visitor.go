// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgramm
import "github.com/antlr4-go/antlr/v4"

import "fmt"

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

	// Visit a parse tree produced by SwiftgrammParser#print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#declare_var.
	VisitDeclare_var(ctx *Declare_varContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#declare_constant.
	VisitDeclare_constant(ctx *Declare_constantContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#assign_var.
	VisitAssign_var(ctx *Assign_varContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#if_sentence.
	VisitIf_sentence(ctx *If_sentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#switch_sentence.
	VisitSwitch_sentence(ctx *Switch_sentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#switch_cases.
	VisitSwitch_cases(ctx *Switch_casesContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#switch_case.
	VisitSwitch_case(ctx *Switch_caseContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#while_sentence.
	VisitWhile_sentence(ctx *While_sentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#for_sentence.
	VisitFor_sentence(ctx *For_sentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#guard_sentence.
	VisitGuard_sentence(ctx *Guard_sentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#break_sentence.
	VisitBreak_sentence(ctx *Break_sentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#continue_sentence.
	VisitContinue_sentence(ctx *Continue_sentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#return_sentence.
	VisitReturn_sentence(ctx *Return_sentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#datatype.
	VisitDatatype(ctx *DatatypeContext) interface{}
}
