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

	// EnterSentence is called when entering the sentence production.
	EnterSentence(c *SentenceContext)

	// EnterIncrement_bl is called when entering the increment_bl production.
	EnterIncrement_bl(c *Increment_blContext)

	// EnterDecrement_bl is called when entering the decrement_bl production.
	EnterDecrement_bl(c *Decrement_blContext)

	// EnterBreak_bl is called when entering the break_bl production.
	EnterBreak_bl(c *Break_blContext)

	// EnterReturn_bl is called when entering the return_bl production.
	EnterReturn_bl(c *Return_blContext)

	// EnterContinue_bl is called when entering the continue_bl production.
	EnterContinue_bl(c *Continue_blContext)

	// EnterDeclare_let is called when entering the declare_let production.
	EnterDeclare_let(c *Declare_letContext)

	// EnterNative_function is called when entering the native_function production.
	EnterNative_function(c *Native_functionContext)

	// EnterDeclare_var is called when entering the declare_var production.
	EnterDeclare_var(c *Declare_varContext)

	// EnterAssign_bl is called when entering the assign_bl production.
	EnterAssign_bl(c *Assign_blContext)

	// EnterPrint_bl is called when entering the print_bl production.
	EnterPrint_bl(c *Print_blContext)

	// EnterList_print is called when entering the list_print production.
	EnterList_print(c *List_printContext)

	// EnterIf_bl is called when entering the if_bl production.
	EnterIf_bl(c *If_blContext)

	// EnterElse_if is called when entering the else_if production.
	EnterElse_if(c *Else_ifContext)

	// EnterWhile_bl is called when entering the while_bl production.
	EnterWhile_bl(c *While_blContext)

	// EnterFor_bl is called when entering the for_bl production.
	EnterFor_bl(c *For_blContext)

	// EnterGuard_bl is called when entering the guard_bl production.
	EnterGuard_bl(c *Guard_blContext)

	// EnterVector_bl is called when entering the vector_bl production.
	EnterVector_bl(c *Vector_blContext)

	// EnterArray_exp is called when entering the array_exp production.
	EnterArray_exp(c *Array_expContext)

	// EnterArray_functions is called when entering the array_functions production.
	EnterArray_functions(c *Array_functionsContext)

	// EnterAssign_vector is called when entering the assign_vector production.
	EnterAssign_vector(c *Assign_vectorContext)

	// EnterFunction_bl is called when entering the function_bl production.
	EnterFunction_bl(c *Function_blContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterExtern_params is called when entering the extern_params production.
	EnterExtern_params(c *Extern_paramsContext)

	// EnterCall_function is called when entering the call_function production.
	EnterCall_function(c *Call_functionContext)

	// EnterList_exp is called when entering the list_exp production.
	EnterList_exp(c *List_expContext)

	// EnterCall_function_exp is called when entering the call_function_exp production.
	EnterCall_function_exp(c *Call_function_expContext)

	// EnterDeclare_array_bl is called when entering the declare_array_bl production.
	EnterDeclare_array_bl(c *Declare_array_blContext)

	// EnterExp_matriz is called when entering the exp_matriz production.
	EnterExp_matriz(c *Exp_matrizContext)

	// EnterType_matrix is called when entering the type_matrix production.
	EnterType_matrix(c *Type_matrixContext)

	// EnterDefinition_matrix is called when entering the definition_matrix production.
	EnterDefinition_matrix(c *Definition_matrixContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDatatype is called when entering the datatype production.
	EnterDatatype(c *DatatypeContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

	// ExitIncrement_bl is called when exiting the increment_bl production.
	ExitIncrement_bl(c *Increment_blContext)

	// ExitDecrement_bl is called when exiting the decrement_bl production.
	ExitDecrement_bl(c *Decrement_blContext)

	// ExitBreak_bl is called when exiting the break_bl production.
	ExitBreak_bl(c *Break_blContext)

	// ExitReturn_bl is called when exiting the return_bl production.
	ExitReturn_bl(c *Return_blContext)

	// ExitContinue_bl is called when exiting the continue_bl production.
	ExitContinue_bl(c *Continue_blContext)

	// ExitDeclare_let is called when exiting the declare_let production.
	ExitDeclare_let(c *Declare_letContext)

	// ExitNative_function is called when exiting the native_function production.
	ExitNative_function(c *Native_functionContext)

	// ExitDeclare_var is called when exiting the declare_var production.
	ExitDeclare_var(c *Declare_varContext)

	// ExitAssign_bl is called when exiting the assign_bl production.
	ExitAssign_bl(c *Assign_blContext)

	// ExitPrint_bl is called when exiting the print_bl production.
	ExitPrint_bl(c *Print_blContext)

	// ExitList_print is called when exiting the list_print production.
	ExitList_print(c *List_printContext)

	// ExitIf_bl is called when exiting the if_bl production.
	ExitIf_bl(c *If_blContext)

	// ExitElse_if is called when exiting the else_if production.
	ExitElse_if(c *Else_ifContext)

	// ExitWhile_bl is called when exiting the while_bl production.
	ExitWhile_bl(c *While_blContext)

	// ExitFor_bl is called when exiting the for_bl production.
	ExitFor_bl(c *For_blContext)

	// ExitGuard_bl is called when exiting the guard_bl production.
	ExitGuard_bl(c *Guard_blContext)

	// ExitVector_bl is called when exiting the vector_bl production.
	ExitVector_bl(c *Vector_blContext)

	// ExitArray_exp is called when exiting the array_exp production.
	ExitArray_exp(c *Array_expContext)

	// ExitArray_functions is called when exiting the array_functions production.
	ExitArray_functions(c *Array_functionsContext)

	// ExitAssign_vector is called when exiting the assign_vector production.
	ExitAssign_vector(c *Assign_vectorContext)

	// ExitFunction_bl is called when exiting the function_bl production.
	ExitFunction_bl(c *Function_blContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitExtern_params is called when exiting the extern_params production.
	ExitExtern_params(c *Extern_paramsContext)

	// ExitCall_function is called when exiting the call_function production.
	ExitCall_function(c *Call_functionContext)

	// ExitList_exp is called when exiting the list_exp production.
	ExitList_exp(c *List_expContext)

	// ExitCall_function_exp is called when exiting the call_function_exp production.
	ExitCall_function_exp(c *Call_function_expContext)

	// ExitDeclare_array_bl is called when exiting the declare_array_bl production.
	ExitDeclare_array_bl(c *Declare_array_blContext)

	// ExitExp_matriz is called when exiting the exp_matriz production.
	ExitExp_matriz(c *Exp_matrizContext)

	// ExitType_matrix is called when exiting the type_matrix production.
	ExitType_matrix(c *Type_matrixContext)

	// ExitDefinition_matrix is called when exiting the definition_matrix production.
	ExitDefinition_matrix(c *Definition_matrixContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)
}
