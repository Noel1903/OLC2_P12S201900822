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

func (v *BaseSwiftgrammVisitor) VisitSentences(ctx *SentencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitSentence(ctx *SentenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitDeclare_var(ctx *Declare_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitDeclare_constant(ctx *Declare_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitAssign_var(ctx *Assign_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftgrammVisitor) VisitDatatype(ctx *DatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}
