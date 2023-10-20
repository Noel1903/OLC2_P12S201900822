// Generated from e:/USAC2023/SegundoSemestre/Compiladores2/P1/OLC2_P12S201900822/backend/grammar/Swiftgramm.g4 by ANTLR 4.13.1

        import "grammar/expressions"
        import "grammar/instructions"
        import "grammar/symbol"
        import "grammar/abstract"

import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SwiftgrammParser}.
 */
public interface SwiftgrammListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#s}.
	 * @param ctx the parse tree
	 */
	void enterS(SwiftgrammParser.SContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#s}.
	 * @param ctx the parse tree
	 */
	void exitS(SwiftgrammParser.SContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(SwiftgrammParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(SwiftgrammParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#sentence}.
	 * @param ctx the parse tree
	 */
	void enterSentence(SwiftgrammParser.SentenceContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#sentence}.
	 * @param ctx the parse tree
	 */
	void exitSentence(SwiftgrammParser.SentenceContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#switch_bl}.
	 * @param ctx the parse tree
	 */
	void enterSwitch_bl(SwiftgrammParser.Switch_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#switch_bl}.
	 * @param ctx the parse tree
	 */
	void exitSwitch_bl(SwiftgrammParser.Switch_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#cases}.
	 * @param ctx the parse tree
	 */
	void enterCases(SwiftgrammParser.CasesContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#cases}.
	 * @param ctx the parse tree
	 */
	void exitCases(SwiftgrammParser.CasesContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#increment_bl}.
	 * @param ctx the parse tree
	 */
	void enterIncrement_bl(SwiftgrammParser.Increment_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#increment_bl}.
	 * @param ctx the parse tree
	 */
	void exitIncrement_bl(SwiftgrammParser.Increment_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#decrement_bl}.
	 * @param ctx the parse tree
	 */
	void enterDecrement_bl(SwiftgrammParser.Decrement_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#decrement_bl}.
	 * @param ctx the parse tree
	 */
	void exitDecrement_bl(SwiftgrammParser.Decrement_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#break_bl}.
	 * @param ctx the parse tree
	 */
	void enterBreak_bl(SwiftgrammParser.Break_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#break_bl}.
	 * @param ctx the parse tree
	 */
	void exitBreak_bl(SwiftgrammParser.Break_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#return_bl}.
	 * @param ctx the parse tree
	 */
	void enterReturn_bl(SwiftgrammParser.Return_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#return_bl}.
	 * @param ctx the parse tree
	 */
	void exitReturn_bl(SwiftgrammParser.Return_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#continue_bl}.
	 * @param ctx the parse tree
	 */
	void enterContinue_bl(SwiftgrammParser.Continue_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#continue_bl}.
	 * @param ctx the parse tree
	 */
	void exitContinue_bl(SwiftgrammParser.Continue_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#declare_let}.
	 * @param ctx the parse tree
	 */
	void enterDeclare_let(SwiftgrammParser.Declare_letContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#declare_let}.
	 * @param ctx the parse tree
	 */
	void exitDeclare_let(SwiftgrammParser.Declare_letContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#native_function}.
	 * @param ctx the parse tree
	 */
	void enterNative_function(SwiftgrammParser.Native_functionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#native_function}.
	 * @param ctx the parse tree
	 */
	void exitNative_function(SwiftgrammParser.Native_functionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#declare_var}.
	 * @param ctx the parse tree
	 */
	void enterDeclare_var(SwiftgrammParser.Declare_varContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#declare_var}.
	 * @param ctx the parse tree
	 */
	void exitDeclare_var(SwiftgrammParser.Declare_varContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#assign_bl}.
	 * @param ctx the parse tree
	 */
	void enterAssign_bl(SwiftgrammParser.Assign_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#assign_bl}.
	 * @param ctx the parse tree
	 */
	void exitAssign_bl(SwiftgrammParser.Assign_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#print_bl}.
	 * @param ctx the parse tree
	 */
	void enterPrint_bl(SwiftgrammParser.Print_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#print_bl}.
	 * @param ctx the parse tree
	 */
	void exitPrint_bl(SwiftgrammParser.Print_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#list_print}.
	 * @param ctx the parse tree
	 */
	void enterList_print(SwiftgrammParser.List_printContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#list_print}.
	 * @param ctx the parse tree
	 */
	void exitList_print(SwiftgrammParser.List_printContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#if_bl}.
	 * @param ctx the parse tree
	 */
	void enterIf_bl(SwiftgrammParser.If_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#if_bl}.
	 * @param ctx the parse tree
	 */
	void exitIf_bl(SwiftgrammParser.If_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#else_if}.
	 * @param ctx the parse tree
	 */
	void enterElse_if(SwiftgrammParser.Else_ifContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#else_if}.
	 * @param ctx the parse tree
	 */
	void exitElse_if(SwiftgrammParser.Else_ifContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#while_bl}.
	 * @param ctx the parse tree
	 */
	void enterWhile_bl(SwiftgrammParser.While_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#while_bl}.
	 * @param ctx the parse tree
	 */
	void exitWhile_bl(SwiftgrammParser.While_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#for_bl}.
	 * @param ctx the parse tree
	 */
	void enterFor_bl(SwiftgrammParser.For_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#for_bl}.
	 * @param ctx the parse tree
	 */
	void exitFor_bl(SwiftgrammParser.For_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#guard_bl}.
	 * @param ctx the parse tree
	 */
	void enterGuard_bl(SwiftgrammParser.Guard_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#guard_bl}.
	 * @param ctx the parse tree
	 */
	void exitGuard_bl(SwiftgrammParser.Guard_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#vector_bl}.
	 * @param ctx the parse tree
	 */
	void enterVector_bl(SwiftgrammParser.Vector_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#vector_bl}.
	 * @param ctx the parse tree
	 */
	void exitVector_bl(SwiftgrammParser.Vector_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#array_exp}.
	 * @param ctx the parse tree
	 */
	void enterArray_exp(SwiftgrammParser.Array_expContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#array_exp}.
	 * @param ctx the parse tree
	 */
	void exitArray_exp(SwiftgrammParser.Array_expContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#array_functions}.
	 * @param ctx the parse tree
	 */
	void enterArray_functions(SwiftgrammParser.Array_functionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#array_functions}.
	 * @param ctx the parse tree
	 */
	void exitArray_functions(SwiftgrammParser.Array_functionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#assign_vector}.
	 * @param ctx the parse tree
	 */
	void enterAssign_vector(SwiftgrammParser.Assign_vectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#assign_vector}.
	 * @param ctx the parse tree
	 */
	void exitAssign_vector(SwiftgrammParser.Assign_vectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#function_bl}.
	 * @param ctx the parse tree
	 */
	void enterFunction_bl(SwiftgrammParser.Function_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#function_bl}.
	 * @param ctx the parse tree
	 */
	void exitFunction_bl(SwiftgrammParser.Function_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#params}.
	 * @param ctx the parse tree
	 */
	void enterParams(SwiftgrammParser.ParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#params}.
	 * @param ctx the parse tree
	 */
	void exitParams(SwiftgrammParser.ParamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#extern_params}.
	 * @param ctx the parse tree
	 */
	void enterExtern_params(SwiftgrammParser.Extern_paramsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#extern_params}.
	 * @param ctx the parse tree
	 */
	void exitExtern_params(SwiftgrammParser.Extern_paramsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#call_function}.
	 * @param ctx the parse tree
	 */
	void enterCall_function(SwiftgrammParser.Call_functionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#call_function}.
	 * @param ctx the parse tree
	 */
	void exitCall_function(SwiftgrammParser.Call_functionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#list_exp}.
	 * @param ctx the parse tree
	 */
	void enterList_exp(SwiftgrammParser.List_expContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#list_exp}.
	 * @param ctx the parse tree
	 */
	void exitList_exp(SwiftgrammParser.List_expContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#call_function_exp}.
	 * @param ctx the parse tree
	 */
	void enterCall_function_exp(SwiftgrammParser.Call_function_expContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#call_function_exp}.
	 * @param ctx the parse tree
	 */
	void exitCall_function_exp(SwiftgrammParser.Call_function_expContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#declare_array_bl}.
	 * @param ctx the parse tree
	 */
	void enterDeclare_array_bl(SwiftgrammParser.Declare_array_blContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#declare_array_bl}.
	 * @param ctx the parse tree
	 */
	void exitDeclare_array_bl(SwiftgrammParser.Declare_array_blContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#exp_matriz}.
	 * @param ctx the parse tree
	 */
	void enterExp_matriz(SwiftgrammParser.Exp_matrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#exp_matriz}.
	 * @param ctx the parse tree
	 */
	void exitExp_matriz(SwiftgrammParser.Exp_matrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#type_matrix}.
	 * @param ctx the parse tree
	 */
	void enterType_matrix(SwiftgrammParser.Type_matrixContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#type_matrix}.
	 * @param ctx the parse tree
	 */
	void exitType_matrix(SwiftgrammParser.Type_matrixContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#definition_matrix}.
	 * @param ctx the parse tree
	 */
	void enterDefinition_matrix(SwiftgrammParser.Definition_matrixContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#definition_matrix}.
	 * @param ctx the parse tree
	 */
	void exitDefinition_matrix(SwiftgrammParser.Definition_matrixContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(SwiftgrammParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(SwiftgrammParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftgrammParser#datatype}.
	 * @param ctx the parse tree
	 */
	void enterDatatype(SwiftgrammParser.DatatypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftgrammParser#datatype}.
	 * @param ctx the parse tree
	 */
	void exitDatatype(SwiftgrammParser.DatatypeContext ctx);
}