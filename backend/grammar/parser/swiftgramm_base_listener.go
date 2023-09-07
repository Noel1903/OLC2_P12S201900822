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

// EnterBreak_bl is called when production break_bl is entered.
func (s *BaseSwiftgrammListener) EnterBreak_bl(ctx *Break_blContext) {}

// ExitBreak_bl is called when production break_bl is exited.
func (s *BaseSwiftgrammListener) ExitBreak_bl(ctx *Break_blContext) {}

// EnterReturn_bl is called when production return_bl is entered.
func (s *BaseSwiftgrammListener) EnterReturn_bl(ctx *Return_blContext) {}

// ExitReturn_bl is called when production return_bl is exited.
func (s *BaseSwiftgrammListener) ExitReturn_bl(ctx *Return_blContext) {}

// EnterContinue_bl is called when production continue_bl is entered.
func (s *BaseSwiftgrammListener) EnterContinue_bl(ctx *Continue_blContext) {}

// ExitContinue_bl is called when production continue_bl is exited.
func (s *BaseSwiftgrammListener) ExitContinue_bl(ctx *Continue_blContext) {}

// EnterDeclare_let is called when production declare_let is entered.
func (s *BaseSwiftgrammListener) EnterDeclare_let(ctx *Declare_letContext) {}

// ExitDeclare_let is called when production declare_let is exited.
func (s *BaseSwiftgrammListener) ExitDeclare_let(ctx *Declare_letContext) {}

// EnterNative_function is called when production native_function is entered.
func (s *BaseSwiftgrammListener) EnterNative_function(ctx *Native_functionContext) {}

// ExitNative_function is called when production native_function is exited.
func (s *BaseSwiftgrammListener) ExitNative_function(ctx *Native_functionContext) {}

// EnterDeclare_var is called when production declare_var is entered.
func (s *BaseSwiftgrammListener) EnterDeclare_var(ctx *Declare_varContext) {}

// ExitDeclare_var is called when production declare_var is exited.
func (s *BaseSwiftgrammListener) ExitDeclare_var(ctx *Declare_varContext) {}

// EnterAssign_bl is called when production assign_bl is entered.
func (s *BaseSwiftgrammListener) EnterAssign_bl(ctx *Assign_blContext) {}

// ExitAssign_bl is called when production assign_bl is exited.
func (s *BaseSwiftgrammListener) ExitAssign_bl(ctx *Assign_blContext) {}

// EnterPrint_bl is called when production print_bl is entered.
func (s *BaseSwiftgrammListener) EnterPrint_bl(ctx *Print_blContext) {}

// ExitPrint_bl is called when production print_bl is exited.
func (s *BaseSwiftgrammListener) ExitPrint_bl(ctx *Print_blContext) {}

// EnterList_print is called when production list_print is entered.
func (s *BaseSwiftgrammListener) EnterList_print(ctx *List_printContext) {}

// ExitList_print is called when production list_print is exited.
func (s *BaseSwiftgrammListener) ExitList_print(ctx *List_printContext) {}

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

// EnterFor_bl is called when production for_bl is entered.
func (s *BaseSwiftgrammListener) EnterFor_bl(ctx *For_blContext) {}

// ExitFor_bl is called when production for_bl is exited.
func (s *BaseSwiftgrammListener) ExitFor_bl(ctx *For_blContext) {}

// EnterGuard_bl is called when production guard_bl is entered.
func (s *BaseSwiftgrammListener) EnterGuard_bl(ctx *Guard_blContext) {}

// ExitGuard_bl is called when production guard_bl is exited.
func (s *BaseSwiftgrammListener) ExitGuard_bl(ctx *Guard_blContext) {}

// EnterVector_bl is called when production vector_bl is entered.
func (s *BaseSwiftgrammListener) EnterVector_bl(ctx *Vector_blContext) {}

// ExitVector_bl is called when production vector_bl is exited.
func (s *BaseSwiftgrammListener) ExitVector_bl(ctx *Vector_blContext) {}

// EnterArray_exp is called when production array_exp is entered.
func (s *BaseSwiftgrammListener) EnterArray_exp(ctx *Array_expContext) {}

// ExitArray_exp is called when production array_exp is exited.
func (s *BaseSwiftgrammListener) ExitArray_exp(ctx *Array_expContext) {}

// EnterArray_functions is called when production array_functions is entered.
func (s *BaseSwiftgrammListener) EnterArray_functions(ctx *Array_functionsContext) {}

// ExitArray_functions is called when production array_functions is exited.
func (s *BaseSwiftgrammListener) ExitArray_functions(ctx *Array_functionsContext) {}

// EnterAssign_vector is called when production assign_vector is entered.
func (s *BaseSwiftgrammListener) EnterAssign_vector(ctx *Assign_vectorContext) {}

// ExitAssign_vector is called when production assign_vector is exited.
func (s *BaseSwiftgrammListener) ExitAssign_vector(ctx *Assign_vectorContext) {}

// EnterFunction_bl is called when production function_bl is entered.
func (s *BaseSwiftgrammListener) EnterFunction_bl(ctx *Function_blContext) {}

// ExitFunction_bl is called when production function_bl is exited.
func (s *BaseSwiftgrammListener) ExitFunction_bl(ctx *Function_blContext) {}

// EnterParams is called when production params is entered.
func (s *BaseSwiftgrammListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseSwiftgrammListener) ExitParams(ctx *ParamsContext) {}

// EnterExtern_params is called when production extern_params is entered.
func (s *BaseSwiftgrammListener) EnterExtern_params(ctx *Extern_paramsContext) {}

// ExitExtern_params is called when production extern_params is exited.
func (s *BaseSwiftgrammListener) ExitExtern_params(ctx *Extern_paramsContext) {}

// EnterCall_function is called when production call_function is entered.
func (s *BaseSwiftgrammListener) EnterCall_function(ctx *Call_functionContext) {}

// ExitCall_function is called when production call_function is exited.
func (s *BaseSwiftgrammListener) ExitCall_function(ctx *Call_functionContext) {}

// EnterList_exp is called when production list_exp is entered.
func (s *BaseSwiftgrammListener) EnterList_exp(ctx *List_expContext) {}

// ExitList_exp is called when production list_exp is exited.
func (s *BaseSwiftgrammListener) ExitList_exp(ctx *List_expContext) {}

// EnterCall_function_exp is called when production call_function_exp is entered.
func (s *BaseSwiftgrammListener) EnterCall_function_exp(ctx *Call_function_expContext) {}

// ExitCall_function_exp is called when production call_function_exp is exited.
func (s *BaseSwiftgrammListener) ExitCall_function_exp(ctx *Call_function_expContext) {}

// EnterDeclare_array_bl is called when production declare_array_bl is entered.
func (s *BaseSwiftgrammListener) EnterDeclare_array_bl(ctx *Declare_array_blContext) {}

// ExitDeclare_array_bl is called when production declare_array_bl is exited.
func (s *BaseSwiftgrammListener) ExitDeclare_array_bl(ctx *Declare_array_blContext) {}

// EnterExp_matriz is called when production exp_matriz is entered.
func (s *BaseSwiftgrammListener) EnterExp_matriz(ctx *Exp_matrizContext) {}

// ExitExp_matriz is called when production exp_matriz is exited.
func (s *BaseSwiftgrammListener) ExitExp_matriz(ctx *Exp_matrizContext) {}

// EnterType_matrix is called when production type_matrix is entered.
func (s *BaseSwiftgrammListener) EnterType_matrix(ctx *Type_matrixContext) {}

// ExitType_matrix is called when production type_matrix is exited.
func (s *BaseSwiftgrammListener) ExitType_matrix(ctx *Type_matrixContext) {}

// EnterDefinition_matrix is called when production definition_matrix is entered.
func (s *BaseSwiftgrammListener) EnterDefinition_matrix(ctx *Definition_matrixContext) {}

// ExitDefinition_matrix is called when production definition_matrix is exited.
func (s *BaseSwiftgrammListener) ExitDefinition_matrix(ctx *Definition_matrixContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSwiftgrammListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSwiftgrammListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseSwiftgrammListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseSwiftgrammListener) ExitDatatype(ctx *DatatypeContext) {}
