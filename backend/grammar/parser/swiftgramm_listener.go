// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgramm
import "github.com/antlr4-go/antlr/v4"

// SwiftgrammListener is a complete listener for a parse tree produced by SwiftgrammParser.
type SwiftgrammListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterSentences is called when entering the sentences production.
	EnterSentences(c *SentencesContext)

	// EnterSentence is called when entering the sentence production.
	EnterSentence(c *SentenceContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterDeclare_var is called when entering the declare_var production.
	EnterDeclare_var(c *Declare_varContext)

	// EnterDeclare_constant is called when entering the declare_constant production.
	EnterDeclare_constant(c *Declare_constantContext)

	// EnterAssign_var is called when entering the assign_var production.
	EnterAssign_var(c *Assign_varContext)

	// EnterIf_sentence is called when entering the if_sentence production.
	EnterIf_sentence(c *If_sentenceContext)

	// EnterSwitch_sentence is called when entering the switch_sentence production.
	EnterSwitch_sentence(c *Switch_sentenceContext)

	// EnterSwitch_cases is called when entering the switch_cases production.
	EnterSwitch_cases(c *Switch_casesContext)

	// EnterSwitch_case is called when entering the switch_case production.
	EnterSwitch_case(c *Switch_caseContext)

	// EnterWhile_sentence is called when entering the while_sentence production.
	EnterWhile_sentence(c *While_sentenceContext)

	// EnterFor_sentence is called when entering the for_sentence production.
	EnterFor_sentence(c *For_sentenceContext)

	// EnterGuard_sentence is called when entering the guard_sentence production.
	EnterGuard_sentence(c *Guard_sentenceContext)

	// EnterBreak_sentence is called when entering the break_sentence production.
	EnterBreak_sentence(c *Break_sentenceContext)

	// EnterContinue_sentence is called when entering the continue_sentence production.
	EnterContinue_sentence(c *Continue_sentenceContext)

	// EnterReturn_sentence is called when entering the return_sentence production.
	EnterReturn_sentence(c *Return_sentenceContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDatatype is called when entering the datatype production.
	EnterDatatype(c *DatatypeContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSentences is called when exiting the sentences production.
	ExitSentences(c *SentencesContext)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitDeclare_var is called when exiting the declare_var production.
	ExitDeclare_var(c *Declare_varContext)

	// ExitDeclare_constant is called when exiting the declare_constant production.
	ExitDeclare_constant(c *Declare_constantContext)

	// ExitAssign_var is called when exiting the assign_var production.
	ExitAssign_var(c *Assign_varContext)

	// ExitIf_sentence is called when exiting the if_sentence production.
	ExitIf_sentence(c *If_sentenceContext)

	// ExitSwitch_sentence is called when exiting the switch_sentence production.
	ExitSwitch_sentence(c *Switch_sentenceContext)

	// ExitSwitch_cases is called when exiting the switch_cases production.
	ExitSwitch_cases(c *Switch_casesContext)

	// ExitSwitch_case is called when exiting the switch_case production.
	ExitSwitch_case(c *Switch_caseContext)

	// ExitWhile_sentence is called when exiting the while_sentence production.
	ExitWhile_sentence(c *While_sentenceContext)

	// ExitFor_sentence is called when exiting the for_sentence production.
	ExitFor_sentence(c *For_sentenceContext)

	// ExitGuard_sentence is called when exiting the guard_sentence production.
	ExitGuard_sentence(c *Guard_sentenceContext)

	// ExitBreak_sentence is called when exiting the break_sentence production.
	ExitBreak_sentence(c *Break_sentenceContext)

	// ExitContinue_sentence is called when exiting the continue_sentence production.
	ExitContinue_sentence(c *Continue_sentenceContext)

	// ExitReturn_sentence is called when exiting the return_sentence production.
	ExitReturn_sentence(c *Return_sentenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)
}
