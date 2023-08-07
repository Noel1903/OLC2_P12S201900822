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

// EnterSentences is called when production sentences is entered.
func (s *BaseSwiftgrammListener) EnterSentences(ctx *SentencesContext) {}

// ExitSentences is called when production sentences is exited.
func (s *BaseSwiftgrammListener) ExitSentences(ctx *SentencesContext) {}

// EnterSentence is called when production sentence is entered.
func (s *BaseSwiftgrammListener) EnterSentence(ctx *SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (s *BaseSwiftgrammListener) ExitSentence(ctx *SentenceContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseSwiftgrammListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseSwiftgrammListener) ExitPrint(ctx *PrintContext) {}

// EnterDeclare_var is called when production declare_var is entered.
func (s *BaseSwiftgrammListener) EnterDeclare_var(ctx *Declare_varContext) {}

// ExitDeclare_var is called when production declare_var is exited.
func (s *BaseSwiftgrammListener) ExitDeclare_var(ctx *Declare_varContext) {}

// EnterDeclare_constant is called when production declare_constant is entered.
func (s *BaseSwiftgrammListener) EnterDeclare_constant(ctx *Declare_constantContext) {}

// ExitDeclare_constant is called when production declare_constant is exited.
func (s *BaseSwiftgrammListener) ExitDeclare_constant(ctx *Declare_constantContext) {}

// EnterAssign_var is called when production assign_var is entered.
func (s *BaseSwiftgrammListener) EnterAssign_var(ctx *Assign_varContext) {}

// ExitAssign_var is called when production assign_var is exited.
func (s *BaseSwiftgrammListener) ExitAssign_var(ctx *Assign_varContext) {}

// EnterIf_sentence is called when production if_sentence is entered.
func (s *BaseSwiftgrammListener) EnterIf_sentence(ctx *If_sentenceContext) {}

// ExitIf_sentence is called when production if_sentence is exited.
func (s *BaseSwiftgrammListener) ExitIf_sentence(ctx *If_sentenceContext) {}

// EnterSwitch_sentence is called when production switch_sentence is entered.
func (s *BaseSwiftgrammListener) EnterSwitch_sentence(ctx *Switch_sentenceContext) {}

// ExitSwitch_sentence is called when production switch_sentence is exited.
func (s *BaseSwiftgrammListener) ExitSwitch_sentence(ctx *Switch_sentenceContext) {}

// EnterSwitch_cases is called when production switch_cases is entered.
func (s *BaseSwiftgrammListener) EnterSwitch_cases(ctx *Switch_casesContext) {}

// ExitSwitch_cases is called when production switch_cases is exited.
func (s *BaseSwiftgrammListener) ExitSwitch_cases(ctx *Switch_casesContext) {}

// EnterSwitch_case is called when production switch_case is entered.
func (s *BaseSwiftgrammListener) EnterSwitch_case(ctx *Switch_caseContext) {}

// ExitSwitch_case is called when production switch_case is exited.
func (s *BaseSwiftgrammListener) ExitSwitch_case(ctx *Switch_caseContext) {}

// EnterWhile_sentence is called when production while_sentence is entered.
func (s *BaseSwiftgrammListener) EnterWhile_sentence(ctx *While_sentenceContext) {}

// ExitWhile_sentence is called when production while_sentence is exited.
func (s *BaseSwiftgrammListener) ExitWhile_sentence(ctx *While_sentenceContext) {}

// EnterFor_sentence is called when production for_sentence is entered.
func (s *BaseSwiftgrammListener) EnterFor_sentence(ctx *For_sentenceContext) {}

// ExitFor_sentence is called when production for_sentence is exited.
func (s *BaseSwiftgrammListener) ExitFor_sentence(ctx *For_sentenceContext) {}

// EnterGuard_sentence is called when production guard_sentence is entered.
func (s *BaseSwiftgrammListener) EnterGuard_sentence(ctx *Guard_sentenceContext) {}

// ExitGuard_sentence is called when production guard_sentence is exited.
func (s *BaseSwiftgrammListener) ExitGuard_sentence(ctx *Guard_sentenceContext) {}

// EnterBreak_sentence is called when production break_sentence is entered.
func (s *BaseSwiftgrammListener) EnterBreak_sentence(ctx *Break_sentenceContext) {}

// ExitBreak_sentence is called when production break_sentence is exited.
func (s *BaseSwiftgrammListener) ExitBreak_sentence(ctx *Break_sentenceContext) {}

// EnterContinue_sentence is called when production continue_sentence is entered.
func (s *BaseSwiftgrammListener) EnterContinue_sentence(ctx *Continue_sentenceContext) {}

// ExitContinue_sentence is called when production continue_sentence is exited.
func (s *BaseSwiftgrammListener) ExitContinue_sentence(ctx *Continue_sentenceContext) {}

// EnterReturn_sentence is called when production return_sentence is entered.
func (s *BaseSwiftgrammListener) EnterReturn_sentence(ctx *Return_sentenceContext) {}

// ExitReturn_sentence is called when production return_sentence is exited.
func (s *BaseSwiftgrammListener) ExitReturn_sentence(ctx *Return_sentenceContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSwiftgrammListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSwiftgrammListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseSwiftgrammListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseSwiftgrammListener) ExitDatatype(ctx *DatatypeContext) {}
