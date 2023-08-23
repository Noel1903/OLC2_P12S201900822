// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgramm
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftgrammListener is a complete listener for a parse tree produced by SwiftgrammParser.
type BaseSwiftgrammListener struct{}

var _ SwiftgrammListener = &BaseSwiftgrammListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftgrammListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftgrammListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftgrammListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftgrammListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftgrammListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftgrammListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftgrammListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftgrammListener) ExitBlock(ctx *BlockContext) {}

// EnterSentence is called when production sentence is entered.
func (s *BaseSwiftgrammListener) EnterSentence(ctx *SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (s *BaseSwiftgrammListener) ExitSentence(ctx *SentenceContext) {}

// EnterIncrement_bl is called when production increment_bl is entered.
func (s *BaseSwiftgrammListener) EnterIncrement_bl(ctx *Increment_blContext) {}

// ExitIncrement_bl is called when production increment_bl is exited.
func (s *BaseSwiftgrammListener) ExitIncrement_bl(ctx *Increment_blContext) {}

// EnterDecrement_bl is called when production decrement_bl is entered.
func (s *BaseSwiftgrammListener) EnterDecrement_bl(ctx *Decrement_blContext) {}

// ExitDecrement_bl is called when production decrement_bl is exited.
func (s *BaseSwiftgrammListener) ExitDecrement_bl(ctx *Decrement_blContext) {}

// EnterDeclare_let is called when production declare_let is entered.
func (s *BaseSwiftgrammListener) EnterDeclare_let(ctx *Declare_letContext) {}

// ExitDeclare_let is called when production declare_let is exited.
func (s *BaseSwiftgrammListener) ExitDeclare_let(ctx *Declare_letContext) {}

// EnterDeclare_var is called when production declare_var is entered.
func (s *BaseSwiftgrammListener) EnterDeclare_var(ctx *Declare_varContext) {}

// ExitDeclare_var is called when production declare_var is exited.
func (s *BaseSwiftgrammListener) ExitDeclare_var(ctx *Declare_varContext) {}

// EnterPrint_bl is called when production print_bl is entered.
func (s *BaseSwiftgrammListener) EnterPrint_bl(ctx *Print_blContext) {}

// ExitPrint_bl is called when production print_bl is exited.
func (s *BaseSwiftgrammListener) ExitPrint_bl(ctx *Print_blContext) {}

// EnterIf_bl is called when production if_bl is entered.
func (s *BaseSwiftgrammListener) EnterIf_bl(ctx *If_blContext) {}

// ExitIf_bl is called when production if_bl is exited.
func (s *BaseSwiftgrammListener) ExitIf_bl(ctx *If_blContext) {}

// EnterElse_if is called when production else_if is entered.
func (s *BaseSwiftgrammListener) EnterElse_if(ctx *Else_ifContext) {}

// ExitElse_if is called when production else_if is exited.
func (s *BaseSwiftgrammListener) ExitElse_if(ctx *Else_ifContext) {}

// EnterWhile_bl is called when production while_bl is entered.
func (s *BaseSwiftgrammListener) EnterWhile_bl(ctx *While_blContext) {}

// ExitWhile_bl is called when production while_bl is exited.
func (s *BaseSwiftgrammListener) ExitWhile_bl(ctx *While_blContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSwiftgrammListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSwiftgrammListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseSwiftgrammListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseSwiftgrammListener) ExitDatatype(ctx *DatatypeContext) {}
