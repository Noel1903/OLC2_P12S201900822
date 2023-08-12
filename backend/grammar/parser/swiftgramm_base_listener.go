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

// EnterDeclare_var is called when production declare_var is entered.
func (s *BaseSwiftgrammListener) EnterDeclare_var(ctx *Declare_varContext) {}

// ExitDeclare_var is called when production declare_var is exited.
func (s *BaseSwiftgrammListener) ExitDeclare_var(ctx *Declare_varContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSwiftgrammListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSwiftgrammListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseSwiftgrammListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseSwiftgrammListener) ExitDatatype(ctx *DatatypeContext) {}
