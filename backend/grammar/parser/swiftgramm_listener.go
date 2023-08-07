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

	// EnterDeclare_var is called when entering the declare_var production.
	EnterDeclare_var(c *Declare_varContext)

	// EnterDeclare_constant is called when entering the declare_constant production.
	EnterDeclare_constant(c *Declare_constantContext)

	// EnterAssign_var is called when entering the assign_var production.
	EnterAssign_var(c *Assign_varContext)

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

	// ExitDeclare_var is called when exiting the declare_var production.
	ExitDeclare_var(c *Declare_varContext)

	// ExitDeclare_constant is called when exiting the declare_constant production.
	ExitDeclare_constant(c *Declare_constantContext)

	// ExitAssign_var is called when exiting the assign_var production.
	ExitAssign_var(c *Assign_varContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)
}
