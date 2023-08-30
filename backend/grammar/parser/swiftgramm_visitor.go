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

	// Visit a parse tree produced by SwiftgrammParser#sentence.
	VisitSentence(ctx *SentenceContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#increment_bl.
	VisitIncrement_bl(ctx *Increment_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#decrement_bl.
	VisitDecrement_bl(ctx *Decrement_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#break_bl.
	VisitBreak_bl(ctx *Break_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#return_bl.
	VisitReturn_bl(ctx *Return_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#continue_bl.
	VisitContinue_bl(ctx *Continue_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#declare_let.
	VisitDeclare_let(ctx *Declare_letContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#declare_var.
	VisitDeclare_var(ctx *Declare_varContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#print_bl.
	VisitPrint_bl(ctx *Print_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#if_bl.
	VisitIf_bl(ctx *If_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#else_if.
	VisitElse_if(ctx *Else_ifContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#while_bl.
	VisitWhile_bl(ctx *While_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#for_bl.
	VisitFor_bl(ctx *For_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#guard_bl.
	VisitGuard_bl(ctx *Guard_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#vector_bl.
	VisitVector_bl(ctx *Vector_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#array_exp.
	VisitArray_exp(ctx *Array_expContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#function_bl.
	VisitFunction_bl(ctx *Function_blContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#call_function.
	VisitCall_function(ctx *Call_functionContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#list_exp.
	VisitList_exp(ctx *List_expContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#call_function_exp.
	VisitCall_function_exp(ctx *Call_function_expContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SwiftgrammParser#datatype.
	VisitDatatype(ctx *DatatypeContext) interface{}
}
