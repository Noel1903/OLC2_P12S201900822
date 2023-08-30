// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgramm
import "github.com/antlr4-go/antlr/v4"

type BaseSwiftgrammVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSwiftgrammVisitor) VisitS(ctx *SContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitSentence(ctx *SentenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitIncrement_bl(ctx *Increment_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitDecrement_bl(ctx *Decrement_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitBreak_bl(ctx *Break_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitReturn_bl(ctx *Return_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitContinue_bl(ctx *Continue_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitDeclare_let(ctx *Declare_letContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitDeclare_var(ctx *Declare_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitPrint_bl(ctx *Print_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitIf_bl(ctx *If_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitElse_if(ctx *Else_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitWhile_bl(ctx *While_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitFor_bl(ctx *For_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitGuard_bl(ctx *Guard_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitVector_bl(ctx *Vector_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitArray_exp(ctx *Array_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitFunction_bl(ctx *Function_blContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitCall_function(ctx *Call_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitList_exp(ctx *List_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitCall_function_exp(ctx *Call_function_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitDatatype(ctx *DatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}
