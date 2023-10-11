// Generated from e:/USAC2023/SegundoSemestre/Compiladores2/P1/OLC2_P12S201900822/backend/grammar/Swiftgramm.g4 by ANTLR 4.13.1

        import "grammar/expressions"
        import "grammar/instructions"
        import "grammar/symbol"
        import "grammar/abstract"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class SwiftgrammParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, STRING=3, BOOL=4, CHARACTER=5, TRUE=6, FALSE=7, NIL=8, 
		VAR=9, LET=10, PRINT=11, IF=12, ELSE=13, SWITCH=14, CASE=15, BREAK=16, 
		DEFAULT=17, WHILE=18, FOR=19, IN=20, GUARD=21, CONTINUE=22, RETURN=23, 
		FUNC=24, STRUCT=25, MUTATING=26, SELF=27, INOUT=28, APPEND=29, REMOVELAST=30, 
		REMOVE=31, AT=32, ISEMPTY=33, COUNT=34, NUMBER=35, FLOATT=36, ID=37, CHARACTER_LITERAL=38, 
		STRING_LITERAL=39, INCREMENT=40, DECREMENT=41, RANGE=42, SUMMATION=43, 
		SUBTRACTION=44, MULTIPLICATION=45, DIVISION=46, MOD=47, QUESTION_MARK=48, 
		OR=49, AND=50, NOT=51, EQUAL=52, NOT_EQUAL=53, LESS_THAN=54, LESS_THAN_EQUAL=55, 
		GREATER_THAN=56, GREATER_THAN_EQUAL=57, ASSIGN=58, DOT=59, COMMA=60, COLON=61, 
		SEMICOLON=62, OPEN_PARENTHESIS=63, CLOSE_PARENTHESIS=64, OPEN_kEY=65, 
		CLOSE_kEY=66, OPEN_BRACKET=67, CLOSE_BRACKET=68, ARROW=69, UNDERSCORE=70, 
		WHITESPACE=71, COMMENT=72, LINE_COMMENT=73;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_sentence = 2, RULE_increment_bl = 3, 
		RULE_decrement_bl = 4, RULE_break_bl = 5, RULE_return_bl = 6, RULE_continue_bl = 7, 
		RULE_declare_let = 8, RULE_native_function = 9, RULE_declare_var = 10, 
		RULE_assign_bl = 11, RULE_print_bl = 12, RULE_list_print = 13, RULE_if_bl = 14, 
		RULE_else_if = 15, RULE_while_bl = 16, RULE_for_bl = 17, RULE_guard_bl = 18, 
		RULE_vector_bl = 19, RULE_array_exp = 20, RULE_array_functions = 21, RULE_assign_vector = 22, 
		RULE_function_bl = 23, RULE_params = 24, RULE_extern_params = 25, RULE_call_function = 26, 
		RULE_list_exp = 27, RULE_call_function_exp = 28, RULE_declare_array_bl = 29, 
		RULE_exp_matriz = 30, RULE_type_matrix = 31, RULE_definition_matrix = 32, 
		RULE_expression = 33, RULE_datatype = 34;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "sentence", "increment_bl", "decrement_bl", "break_bl", 
			"return_bl", "continue_bl", "declare_let", "native_function", "declare_var", 
			"assign_bl", "print_bl", "list_print", "if_bl", "else_if", "while_bl", 
			"for_bl", "guard_bl", "vector_bl", "array_exp", "array_functions", "assign_vector", 
			"function_bl", "params", "extern_params", "call_function", "list_exp", 
			"call_function_exp", "declare_array_bl", "exp_matriz", "type_matrix", 
			"definition_matrix", "expression", "datatype"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'break'", "'default'", "'while'", "'for'", "'in'", "'guard'", 
			"'continue'", "'return'", "'func'", "'struct'", "'mutating'", "'self'", 
			"'inout'", "'append'", "'removeLast'", "'remove'", "'at'", "'IsEmpty'", 
			"'count'", null, null, null, null, null, "'+='", "'-='", "'...'", "'+'", 
			"'-'", "'*'", "'/'", "'%'", "'?'", "'||'", "'&&'", "'!'", "'=='", "'!='", 
			"'<'", "'<='", "'>'", "'>='", "'='", "'.'", "','", "':'", "';'", "'('", 
			"')'", "'{'", "'}'", "'['", "']'", "'->'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE", 
			"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK", 
			"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC", 
			"STRUCT", "MUTATING", "SELF", "INOUT", "APPEND", "REMOVELAST", "REMOVE", 
			"AT", "ISEMPTY", "COUNT", "NUMBER", "FLOATT", "ID", "CHARACTER_LITERAL", 
			"STRING_LITERAL", "INCREMENT", "DECREMENT", "RANGE", "SUMMATION", "SUBTRACTION", 
			"MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT", 
			"EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN", 
			"GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON", 
			"OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", "OPEN_BRACKET", 
			"CLOSE_BRACKET", "ARROW", "UNDERSCORE", "WHITESPACE", "COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Swiftgramm.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }


	        var CountM int = 0 

	public SwiftgrammParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SContext extends ParserRuleContext {
		public []interface{} code;
		public BlockContext block;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SwiftgrammParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterS(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitS(this);
		}
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			((SContext)_localctx).block = block();
			setState(71);
			match(EOF);

			        _localctx.code = ((SContext)_localctx).block.blk

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public []interface{} blk;
		public SentenceContext sentence;
		public List<SentenceContext> instr = new ArrayList<SentenceContext>();
		public List<SentenceContext> sentence() {
			return getRuleContexts(SentenceContext.class);
		}
		public SentenceContext sentence(int i) {
			return getRuleContext(SentenceContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitBlock(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);

		        _localctx.blk = []interface{}{}
		        var listInt []ISentenceContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(74);
				((BlockContext)_localctx).sentence = sentence();
				((BlockContext)_localctx).instr.add(((BlockContext)_localctx).sentence);
				}
				}
				setState(77); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 137471270400L) != 0) );

			        listInt = localctx.(*BlockContext).GetInstr()
			        for _,e := range listInt{
			                _localctx.blk = append(_localctx.blk,e.GetInstr())
			        }

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SentenceContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Declare_varContext declare_var;
		public Declare_letContext declare_let;
		public Assign_blContext assign_bl;
		public Assign_vectorContext assign_vector;
		public Print_blContext print_bl;
		public If_blContext if_bl;
		public Increment_blContext increment_bl;
		public Decrement_blContext decrement_bl;
		public While_blContext while_bl;
		public For_blContext for_bl;
		public Guard_blContext guard_bl;
		public Break_blContext break_bl;
		public Return_blContext return_bl;
		public Continue_blContext continue_bl;
		public Vector_blContext vector_bl;
		public Function_blContext function_bl;
		public Call_functionContext call_function;
		public Array_functionsContext array_functions;
		public Declare_array_blContext declare_array_bl;
		public Declare_varContext declare_var() {
			return getRuleContext(Declare_varContext.class,0);
		}
		public Declare_letContext declare_let() {
			return getRuleContext(Declare_letContext.class,0);
		}
		public Assign_blContext assign_bl() {
			return getRuleContext(Assign_blContext.class,0);
		}
		public Assign_vectorContext assign_vector() {
			return getRuleContext(Assign_vectorContext.class,0);
		}
		public Print_blContext print_bl() {
			return getRuleContext(Print_blContext.class,0);
		}
		public If_blContext if_bl() {
			return getRuleContext(If_blContext.class,0);
		}
		public Increment_blContext increment_bl() {
			return getRuleContext(Increment_blContext.class,0);
		}
		public Decrement_blContext decrement_bl() {
			return getRuleContext(Decrement_blContext.class,0);
		}
		public While_blContext while_bl() {
			return getRuleContext(While_blContext.class,0);
		}
		public For_blContext for_bl() {
			return getRuleContext(For_blContext.class,0);
		}
		public Guard_blContext guard_bl() {
			return getRuleContext(Guard_blContext.class,0);
		}
		public Break_blContext break_bl() {
			return getRuleContext(Break_blContext.class,0);
		}
		public Return_blContext return_bl() {
			return getRuleContext(Return_blContext.class,0);
		}
		public Continue_blContext continue_bl() {
			return getRuleContext(Continue_blContext.class,0);
		}
		public Vector_blContext vector_bl() {
			return getRuleContext(Vector_blContext.class,0);
		}
		public Function_blContext function_bl() {
			return getRuleContext(Function_blContext.class,0);
		}
		public Call_functionContext call_function() {
			return getRuleContext(Call_functionContext.class,0);
		}
		public Array_functionsContext array_functions() {
			return getRuleContext(Array_functionsContext.class,0);
		}
		public Declare_array_blContext declare_array_bl() {
			return getRuleContext(Declare_array_blContext.class,0);
		}
		public SentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentence; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterSentence(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitSentence(this);
		}
	}

	public final SentenceContext sentence() throws RecognitionException {
		SentenceContext _localctx = new SentenceContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentence);
		try {
			setState(138);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				((SentenceContext)_localctx).declare_var = declare_var();
				_localctx.instr = ((SentenceContext)_localctx).declare_var.instr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(84);
				((SentenceContext)_localctx).declare_let = declare_let();
				_localctx.instr = ((SentenceContext)_localctx).declare_let.instr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(87);
				((SentenceContext)_localctx).assign_bl = assign_bl();
				_localctx.instr = ((SentenceContext)_localctx).assign_bl.instr
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(90);
				((SentenceContext)_localctx).assign_vector = assign_vector();
				_localctx.instr = ((SentenceContext)_localctx).assign_vector.instr
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(93);
				((SentenceContext)_localctx).print_bl = print_bl();
				_localctx.instr = ((SentenceContext)_localctx).print_bl.instr
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(96);
				((SentenceContext)_localctx).if_bl = if_bl();
				_localctx.instr = ((SentenceContext)_localctx).if_bl.instr
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(99);
				((SentenceContext)_localctx).increment_bl = increment_bl();
				_localctx.instr = ((SentenceContext)_localctx).increment_bl.instr
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(102);
				((SentenceContext)_localctx).decrement_bl = decrement_bl();
				_localctx.instr = ((SentenceContext)_localctx).decrement_bl.instr
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(105);
				((SentenceContext)_localctx).while_bl = while_bl();
				_localctx.instr = ((SentenceContext)_localctx).while_bl.instr
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(108);
				((SentenceContext)_localctx).for_bl = for_bl();
				_localctx.instr = ((SentenceContext)_localctx).for_bl.instr
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(111);
				((SentenceContext)_localctx).guard_bl = guard_bl();
				_localctx.instr = ((SentenceContext)_localctx).guard_bl.instr
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(114);
				((SentenceContext)_localctx).break_bl = break_bl();
				_localctx.instr = ((SentenceContext)_localctx).break_bl.instr
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(117);
				((SentenceContext)_localctx).return_bl = return_bl();
				_localctx.instr = ((SentenceContext)_localctx).return_bl.instr
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(120);
				((SentenceContext)_localctx).continue_bl = continue_bl();
				_localctx.instr = ((SentenceContext)_localctx).continue_bl.instr
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(123);
				((SentenceContext)_localctx).vector_bl = vector_bl();
				_localctx.instr = ((SentenceContext)_localctx).vector_bl.instr
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(126);
				((SentenceContext)_localctx).function_bl = function_bl();
				_localctx.instr = ((SentenceContext)_localctx).function_bl.instr
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(129);
				((SentenceContext)_localctx).call_function = call_function();
				_localctx.instr = ((SentenceContext)_localctx).call_function.instr
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(132);
				((SentenceContext)_localctx).array_functions = array_functions();
				_localctx.instr = ((SentenceContext)_localctx).array_functions.instr
				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(135);
				((SentenceContext)_localctx).declare_array_bl = declare_array_bl();
				 _localctx.instr = ((SentenceContext)_localctx).declare_array_bl.instr
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Increment_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public Token INCREMENT;
		public ExpressionContext expression;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode INCREMENT() { return getToken(SwiftgrammParser.INCREMENT, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Increment_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_increment_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterIncrement_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitIncrement_bl(this);
		}
	}

	public final Increment_blContext increment_bl() throws RecognitionException {
		Increment_blContext _localctx = new Increment_blContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_increment_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			((Increment_blContext)_localctx).ID = match(ID);
			setState(141);
			((Increment_blContext)_localctx).INCREMENT = match(INCREMENT);
			setState(142);
			((Increment_blContext)_localctx).expression = expression(0);

			        _localctx.instr = instructions.NewIncreDecrem((((Increment_blContext)_localctx).ID!=null?((Increment_blContext)_localctx).ID.getText():null),(((Increment_blContext)_localctx).INCREMENT!=null?((Increment_blContext)_localctx).INCREMENT.getText():null),((Increment_blContext)_localctx).expression.p,(((Increment_blContext)_localctx).ID!=null?((Increment_blContext)_localctx).ID.getLine():0),(((Increment_blContext)_localctx).ID!=null?((Increment_blContext)_localctx).ID.getCharPositionInLine():0))

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Decrement_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public Token DECREMENT;
		public ExpressionContext expression;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode DECREMENT() { return getToken(SwiftgrammParser.DECREMENT, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Decrement_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decrement_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterDecrement_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitDecrement_bl(this);
		}
	}

	public final Decrement_blContext decrement_bl() throws RecognitionException {
		Decrement_blContext _localctx = new Decrement_blContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_decrement_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(145);
			((Decrement_blContext)_localctx).ID = match(ID);
			setState(146);
			((Decrement_blContext)_localctx).DECREMENT = match(DECREMENT);
			setState(147);
			((Decrement_blContext)_localctx).expression = expression(0);

			        _localctx.instr = instructions.NewIncreDecrem((((Decrement_blContext)_localctx).ID!=null?((Decrement_blContext)_localctx).ID.getText():null),(((Decrement_blContext)_localctx).DECREMENT!=null?((Decrement_blContext)_localctx).DECREMENT.getText():null),((Decrement_blContext)_localctx).expression.p,(((Decrement_blContext)_localctx).ID!=null?((Decrement_blContext)_localctx).ID.getLine():0),(((Decrement_blContext)_localctx).ID!=null?((Decrement_blContext)_localctx).ID.getCharPositionInLine():0))

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Break_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public TerminalNode BREAK() { return getToken(SwiftgrammParser.BREAK, 0); }
		public Break_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterBreak_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitBreak_bl(this);
		}
	}

	public final Break_blContext break_bl() throws RecognitionException {
		Break_blContext _localctx = new Break_blContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_break_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			match(BREAK);

			        _localctx.instr = instructions.NewBreak("break")

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Return_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public ExpressionContext expression;
		public Token RETURN;
		public TerminalNode RETURN() { return getToken(SwiftgrammParser.RETURN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Return_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterReturn_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitReturn_bl(this);
		}
	}

	public final Return_blContext return_bl() throws RecognitionException {
		Return_blContext _localctx = new Return_blContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_return_bl);
		try {
			setState(159);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(153);
				match(RETURN);
				setState(154);
				((Return_blContext)_localctx).expression = expression(0);

				        _localctx.instr = instructions.NewReturn(((Return_blContext)_localctx).expression.p)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(157);
				((Return_blContext)_localctx).RETURN = match(RETURN);

				        _localctx.instr = instructions.NewReturn(expressions.NewNative(nil,symbol.NIL,(((Return_blContext)_localctx).RETURN!=null?((Return_blContext)_localctx).RETURN.getLine():0),(((Return_blContext)_localctx).RETURN!=null?((Return_blContext)_localctx).RETURN.getCharPositionInLine():0)))

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Continue_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public TerminalNode CONTINUE() { return getToken(SwiftgrammParser.CONTINUE, 0); }
		public Continue_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterContinue_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitContinue_bl(this);
		}
	}

	public final Continue_blContext continue_bl() throws RecognitionException {
		Continue_blContext _localctx = new Continue_blContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_continue_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			match(CONTINUE);

			        _localctx.instr = instructions.NewContinue("continue")

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Declare_letContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public DatatypeContext datatype;
		public ExpressionContext expression;
		public TerminalNode LET() { return getToken(SwiftgrammParser.LET, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public DatatypeContext datatype() {
			return getRuleContext(DatatypeContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(SwiftgrammParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Declare_letContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declare_let; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterDeclare_let(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitDeclare_let(this);
		}
	}

	public final Declare_letContext declare_let() throws RecognitionException {
		Declare_letContext _localctx = new Declare_letContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_declare_let);
		try {
			setState(178);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(164);
				match(LET);
				setState(165);
				((Declare_letContext)_localctx).ID = match(ID);
				setState(166);
				match(COLON);
				setState(167);
				((Declare_letContext)_localctx).datatype = datatype();
				setState(168);
				match(ASSIGN);
				setState(169);
				((Declare_letContext)_localctx).expression = expression(0);

				                _localctx.instr = instructions.NewLet((((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getText():null),((Declare_letContext)_localctx).datatype.td,((Declare_letContext)_localctx).expression.p,(((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getLine():0),(((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getCharPositionInLine():0))
				        
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(172);
				match(LET);
				setState(173);
				((Declare_letContext)_localctx).ID = match(ID);
				setState(174);
				match(ASSIGN);
				setState(175);
				((Declare_letContext)_localctx).expression = expression(0);

				                _localctx.instr = instructions.NewLet((((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getText():null),symbol.UNDEFINED,((Declare_letContext)_localctx).expression.p,(((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getLine():0),(((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getCharPositionInLine():0))
				        
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Native_functionContext extends ParserRuleContext {
		public abstract.Expression p;
		public Token STRING;
		public ExpressionContext expression;
		public Token INT;
		public Token FLOAT;
		public TerminalNode STRING() { return getToken(SwiftgrammParser.STRING, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public TerminalNode INT() { return getToken(SwiftgrammParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftgrammParser.FLOAT, 0); }
		public Native_functionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_native_function; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterNative_function(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitNative_function(this);
		}
	}

	public final Native_functionContext native_function() throws RecognitionException {
		Native_functionContext _localctx = new Native_functionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_native_function);
		try {
			setState(198);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(180);
				((Native_functionContext)_localctx).STRING = match(STRING);
				setState(181);
				match(OPEN_PARENTHESIS);
				setState(182);
				((Native_functionContext)_localctx).expression = expression(0);
				setState(183);
				match(CLOSE_PARENTHESIS);

				        _localctx.p = expressions.NewNativeFunction((((Native_functionContext)_localctx).STRING!=null?((Native_functionContext)_localctx).STRING.getText():null),((Native_functionContext)_localctx).expression.p)

				}
				break;
			case INT:
				enterOuterAlt(_localctx, 2);
				{
				setState(186);
				((Native_functionContext)_localctx).INT = match(INT);
				setState(187);
				match(OPEN_PARENTHESIS);
				setState(188);
				((Native_functionContext)_localctx).expression = expression(0);
				setState(189);
				match(CLOSE_PARENTHESIS);

				        _localctx.p = expressions.NewNativeFunction((((Native_functionContext)_localctx).INT!=null?((Native_functionContext)_localctx).INT.getText():null),((Native_functionContext)_localctx).expression.p)

				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 3);
				{
				setState(192);
				((Native_functionContext)_localctx).FLOAT = match(FLOAT);
				setState(193);
				match(OPEN_PARENTHESIS);
				setState(194);
				((Native_functionContext)_localctx).expression = expression(0);
				setState(195);
				match(CLOSE_PARENTHESIS);

				        _localctx.p = expressions.NewNativeFunction((((Native_functionContext)_localctx).FLOAT!=null?((Native_functionContext)_localctx).FLOAT.getText():null),((Native_functionContext)_localctx).expression.p)

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Declare_varContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public DatatypeContext datatype;
		public ExpressionContext expression;
		public Token QUESTION_MARK;
		public TerminalNode VAR() { return getToken(SwiftgrammParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public DatatypeContext datatype() {
			return getRuleContext(DatatypeContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(SwiftgrammParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode QUESTION_MARK() { return getToken(SwiftgrammParser.QUESTION_MARK, 0); }
		public Declare_varContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declare_var; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterDeclare_var(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitDeclare_var(this);
		}
	}

	public final Declare_varContext declare_var() throws RecognitionException {
		Declare_varContext _localctx = new Declare_varContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_declare_var);
		try {
			setState(221);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(200);
				match(VAR);
				setState(201);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(202);
				match(COLON);
				setState(203);
				((Declare_varContext)_localctx).datatype = datatype();
				setState(204);
				match(ASSIGN);
				setState(205);
				((Declare_varContext)_localctx).expression = expression(0);

				                        _localctx.instr = instructions.NewDeclareWithValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),((Declare_varContext)_localctx).datatype.td,((Declare_varContext)_localctx).expression.p,(((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getLine():0),(((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getCharPositionInLine():0))
				                
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(208);
				match(VAR);
				setState(209);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(210);
				match(ASSIGN);
				setState(211);
				((Declare_varContext)_localctx).expression = expression(0);

				                        _localctx.instr = instructions.NewDeclareWithValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),symbol.UNDEFINED,((Declare_varContext)_localctx).expression.p,(((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getLine():0),(((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getCharPositionInLine():0))
				                
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(214);
				match(VAR);
				setState(215);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(216);
				match(COLON);
				setState(217);
				((Declare_varContext)_localctx).datatype = datatype();
				setState(218);
				((Declare_varContext)_localctx).QUESTION_MARK = match(QUESTION_MARK);

				                        _localctx.instr = instructions.NewDeclareWithoutValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),((Declare_varContext)_localctx).datatype.td,expressions.NewNative(nil,symbol.NIL,(((Declare_varContext)_localctx).QUESTION_MARK!=null?((Declare_varContext)_localctx).QUESTION_MARK.getLine():0),(((Declare_varContext)_localctx).QUESTION_MARK!=null?((Declare_varContext)_localctx).QUESTION_MARK.getCharPositionInLine():0)),(((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getLine():0),(((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getCharPositionInLine():0))
				                
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public ExpressionContext expression;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(SwiftgrammParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Assign_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterAssign_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitAssign_bl(this);
		}
	}

	public final Assign_blContext assign_bl() throws RecognitionException {
		Assign_blContext _localctx = new Assign_blContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_assign_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			((Assign_blContext)_localctx).ID = match(ID);
			setState(224);
			match(ASSIGN);
			setState(225);
			((Assign_blContext)_localctx).expression = expression(0);

			        _localctx.instr = instructions.NewAsignVariable((((Assign_blContext)_localctx).ID!=null?((Assign_blContext)_localctx).ID.getText():null),((Assign_blContext)_localctx).expression.p,(((Assign_blContext)_localctx).ID!=null?((Assign_blContext)_localctx).ID.getLine():0),(((Assign_blContext)_localctx).ID!=null?((Assign_blContext)_localctx).ID.getCharPositionInLine():0))

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Print_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public List_printContext list_print;
		public TerminalNode PRINT() { return getToken(SwiftgrammParser.PRINT, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public List_printContext list_print() {
			return getRuleContext(List_printContext.class,0);
		}
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public Print_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterPrint_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitPrint_bl(this);
		}
	}

	public final Print_blContext print_bl() throws RecognitionException {
		Print_blContext _localctx = new Print_blContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_print_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			match(PRINT);
			setState(229);
			match(OPEN_PARENTHESIS);
			setState(230);
			((Print_blContext)_localctx).list_print = list_print();
			setState(231);
			match(CLOSE_PARENTHESIS);

			        _localctx.instr = instructions.NewPrint(((Print_blContext)_localctx).list_print.p)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_printContext extends ParserRuleContext {
		public []interface{} p;
		public ExpressionContext expression;
		public List_printContext list_print;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(SwiftgrammParser.COMMA, 0); }
		public List_printContext list_print() {
			return getRuleContext(List_printContext.class,0);
		}
		public List_printContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_print; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterList_print(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitList_print(this);
		}
	}

	public final List_printContext list_print() throws RecognitionException {
		List_printContext _localctx = new List_printContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_list_print);
		try {
			setState(242);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(234);
				((List_printContext)_localctx).expression = expression(0);
				setState(235);
				match(COMMA);
				setState(236);
				((List_printContext)_localctx).list_print = list_print();

				        _localctx.p = append([]interface{}{((List_printContext)_localctx).expression.p},((List_printContext)_localctx).list_print.p...)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(239);
				((List_printContext)_localctx).expression = expression(0);

				        _localctx.p = []interface{}{((List_printContext)_localctx).expression.p}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token IF;
		public ExpressionContext expression;
		public BlockContext ifblock;
		public BlockContext elseblock;
		public Else_ifContext else_if;
		public TerminalNode IF() { return getToken(SwiftgrammParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<TerminalNode> OPEN_kEY() { return getTokens(SwiftgrammParser.OPEN_kEY); }
		public TerminalNode OPEN_kEY(int i) {
			return getToken(SwiftgrammParser.OPEN_kEY, i);
		}
		public List<TerminalNode> CLOSE_kEY() { return getTokens(SwiftgrammParser.CLOSE_kEY); }
		public TerminalNode CLOSE_kEY(int i) {
			return getToken(SwiftgrammParser.CLOSE_kEY, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(SwiftgrammParser.ELSE, 0); }
		public Else_ifContext else_if() {
			return getRuleContext(Else_ifContext.class,0);
		}
		public If_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterIf_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitIf_bl(this);
		}
	}

	public final If_blContext if_bl() throws RecognitionException {
		If_blContext _localctx = new If_blContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_if_bl);
		try {
			setState(270);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(244);
				((If_blContext)_localctx).IF = match(IF);
				setState(245);
				((If_blContext)_localctx).expression = expression(0);
				setState(246);
				match(OPEN_kEY);
				setState(247);
				((If_blContext)_localctx).ifblock = block();
				setState(248);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewIf(((If_blContext)_localctx).expression.p,((If_blContext)_localctx).ifblock.blk,nil,(((If_blContext)_localctx).IF!=null?((If_blContext)_localctx).IF.getLine():0),(((If_blContext)_localctx).IF!=null?((If_blContext)_localctx).IF.getCharPositionInLine():0))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(251);
				((If_blContext)_localctx).IF = match(IF);
				setState(252);
				((If_blContext)_localctx).expression = expression(0);
				setState(253);
				match(OPEN_kEY);
				setState(254);
				((If_blContext)_localctx).ifblock = block();
				setState(255);
				match(CLOSE_kEY);
				setState(256);
				match(ELSE);
				setState(257);
				match(OPEN_kEY);
				setState(258);
				((If_blContext)_localctx).elseblock = block();
				setState(259);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewIf(((If_blContext)_localctx).expression.p,((If_blContext)_localctx).ifblock.blk,((If_blContext)_localctx).elseblock.blk,(((If_blContext)_localctx).IF!=null?((If_blContext)_localctx).IF.getLine():0),(((If_blContext)_localctx).IF!=null?((If_blContext)_localctx).IF.getCharPositionInLine():0))

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(262);
				((If_blContext)_localctx).IF = match(IF);
				setState(263);
				((If_blContext)_localctx).expression = expression(0);
				setState(264);
				match(OPEN_kEY);
				setState(265);
				((If_blContext)_localctx).ifblock = block();
				setState(266);
				match(CLOSE_kEY);
				setState(267);
				((If_blContext)_localctx).else_if = else_if();

				        _localctx.instr = instructions.NewIf(((If_blContext)_localctx).expression.p,((If_blContext)_localctx).ifblock.blk,((If_blContext)_localctx).else_if.instr,(((If_blContext)_localctx).IF!=null?((If_blContext)_localctx).IF.getLine():0),(((If_blContext)_localctx).IF!=null?((If_blContext)_localctx).IF.getCharPositionInLine():0))

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Else_ifContext extends ParserRuleContext {
		public []interface{} instr;
		public Token ELSE;
		public ExpressionContext expression;
		public BlockContext ifblock;
		public BlockContext elseblock;
		public Else_ifContext else_if;
		public List<TerminalNode> ELSE() { return getTokens(SwiftgrammParser.ELSE); }
		public TerminalNode ELSE(int i) {
			return getToken(SwiftgrammParser.ELSE, i);
		}
		public TerminalNode IF() { return getToken(SwiftgrammParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<TerminalNode> OPEN_kEY() { return getTokens(SwiftgrammParser.OPEN_kEY); }
		public TerminalNode OPEN_kEY(int i) {
			return getToken(SwiftgrammParser.OPEN_kEY, i);
		}
		public List<TerminalNode> CLOSE_kEY() { return getTokens(SwiftgrammParser.CLOSE_kEY); }
		public TerminalNode CLOSE_kEY(int i) {
			return getToken(SwiftgrammParser.CLOSE_kEY, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public Else_ifContext else_if() {
			return getRuleContext(Else_ifContext.class,0);
		}
		public Else_ifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_if; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterElse_if(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitElse_if(this);
		}
	}

	public final Else_ifContext else_if() throws RecognitionException {
		Else_ifContext _localctx = new Else_ifContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_else_if);
		try {
			setState(301);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(272);
				((Else_ifContext)_localctx).ELSE = match(ELSE);
				setState(273);
				match(IF);
				setState(274);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(275);
				match(OPEN_kEY);
				setState(276);
				((Else_ifContext)_localctx).ifblock = block();
				setState(277);
				match(CLOSE_kEY);

				        _localctx.instr = []interface{}{instructions.NewIf(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).ifblock.blk,nil,(((Else_ifContext)_localctx).ELSE!=null?((Else_ifContext)_localctx).ELSE.getLine():0),(((Else_ifContext)_localctx).ELSE!=null?((Else_ifContext)_localctx).ELSE.getCharPositionInLine():0))}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(280);
				((Else_ifContext)_localctx).ELSE = match(ELSE);
				setState(281);
				match(IF);
				setState(282);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(283);
				match(OPEN_kEY);
				setState(284);
				((Else_ifContext)_localctx).ifblock = block();
				setState(285);
				match(CLOSE_kEY);
				setState(286);
				((Else_ifContext)_localctx).ELSE = match(ELSE);
				setState(287);
				match(OPEN_kEY);
				setState(288);
				((Else_ifContext)_localctx).elseblock = block();
				setState(289);
				match(CLOSE_kEY);

				        _localctx.instr = []interface{}{instructions.NewIf(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).ifblock.blk,((Else_ifContext)_localctx).elseblock.blk,(((Else_ifContext)_localctx).ELSE!=null?((Else_ifContext)_localctx).ELSE.getLine():0),(((Else_ifContext)_localctx).ELSE!=null?((Else_ifContext)_localctx).ELSE.getCharPositionInLine():0))}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(292);
				((Else_ifContext)_localctx).ELSE = match(ELSE);
				setState(293);
				match(IF);
				setState(294);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(295);
				match(OPEN_kEY);
				setState(296);
				((Else_ifContext)_localctx).ifblock = block();
				setState(297);
				match(CLOSE_kEY);
				setState(298);
				((Else_ifContext)_localctx).else_if = else_if();

				        _localctx.instr = []interface{}{instructions.NewIf(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).ifblock.blk,((Else_ifContext)_localctx).else_if.instr,(((Else_ifContext)_localctx).ELSE!=null?((Else_ifContext)_localctx).ELSE.getLine():0),(((Else_ifContext)_localctx).ELSE!=null?((Else_ifContext)_localctx).ELSE.getCharPositionInLine():0))}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class While_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token WHILE;
		public ExpressionContext expression;
		public BlockContext block;
		public TerminalNode WHILE() { return getToken(SwiftgrammParser.WHILE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode OPEN_kEY() { return getToken(SwiftgrammParser.OPEN_kEY, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode CLOSE_kEY() { return getToken(SwiftgrammParser.CLOSE_kEY, 0); }
		public While_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterWhile_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitWhile_bl(this);
		}
	}

	public final While_blContext while_bl() throws RecognitionException {
		While_blContext _localctx = new While_blContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_while_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(303);
			((While_blContext)_localctx).WHILE = match(WHILE);
			setState(304);
			((While_blContext)_localctx).expression = expression(0);
			setState(305);
			match(OPEN_kEY);
			setState(306);
			((While_blContext)_localctx).block = block();
			setState(307);
			match(CLOSE_kEY);

			        _localctx.instr = instructions.NewWhile(((While_blContext)_localctx).expression.p,((While_blContext)_localctx).block.blk,(((While_blContext)_localctx).WHILE!=null?((While_blContext)_localctx).WHILE.getLine():0),(((While_blContext)_localctx).WHILE!=null?((While_blContext)_localctx).WHILE.getCharPositionInLine():0))

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public ExpressionContext expression1;
		public ExpressionContext expression2;
		public BlockContext block;
		public TerminalNode FOR() { return getToken(SwiftgrammParser.FOR, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode IN() { return getToken(SwiftgrammParser.IN, 0); }
		public TerminalNode RANGE() { return getToken(SwiftgrammParser.RANGE, 0); }
		public TerminalNode OPEN_kEY() { return getToken(SwiftgrammParser.OPEN_kEY, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode CLOSE_kEY() { return getToken(SwiftgrammParser.CLOSE_kEY, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public For_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterFor_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitFor_bl(this);
		}
	}

	public final For_blContext for_bl() throws RecognitionException {
		For_blContext _localctx = new For_blContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_for_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(310);
			match(FOR);
			setState(311);
			((For_blContext)_localctx).ID = match(ID);
			setState(312);
			match(IN);
			setState(313);
			((For_blContext)_localctx).expression1 = expression(0);
			setState(314);
			match(RANGE);
			setState(315);
			((For_blContext)_localctx).expression2 = expression(0);
			setState(316);
			match(OPEN_kEY);
			setState(317);
			((For_blContext)_localctx).block = block();
			setState(318);
			match(CLOSE_kEY);

			        _localctx.instr = instructions.NewFor((((For_blContext)_localctx).ID!=null?((For_blContext)_localctx).ID.getText():null),((For_blContext)_localctx).expression1.p,((For_blContext)_localctx).expression2.p,((For_blContext)_localctx).block.blk,(((For_blContext)_localctx).ID!=null?((For_blContext)_localctx).ID.getLine():0),(((For_blContext)_localctx).ID!=null?((For_blContext)_localctx).ID.getCharPositionInLine():0))

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Guard_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token GUARD;
		public ExpressionContext expression;
		public BlockContext block;
		public TerminalNode GUARD() { return getToken(SwiftgrammParser.GUARD, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftgrammParser.ELSE, 0); }
		public TerminalNode OPEN_kEY() { return getToken(SwiftgrammParser.OPEN_kEY, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode CLOSE_kEY() { return getToken(SwiftgrammParser.CLOSE_kEY, 0); }
		public Guard_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guard_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterGuard_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitGuard_bl(this);
		}
	}

	public final Guard_blContext guard_bl() throws RecognitionException {
		Guard_blContext _localctx = new Guard_blContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_guard_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			((Guard_blContext)_localctx).GUARD = match(GUARD);
			setState(322);
			((Guard_blContext)_localctx).expression = expression(0);
			setState(323);
			match(ELSE);
			setState(324);
			match(OPEN_kEY);
			setState(325);
			((Guard_blContext)_localctx).block = block();
			setState(326);
			match(CLOSE_kEY);

			        _localctx.instr = instructions.NewGuard(((Guard_blContext)_localctx).expression.p,((Guard_blContext)_localctx).block.blk,(((Guard_blContext)_localctx).GUARD!=null?((Guard_blContext)_localctx).GUARD.getLine():0),(((Guard_blContext)_localctx).GUARD!=null?((Guard_blContext)_localctx).GUARD.getCharPositionInLine():0))

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Vector_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public DatatypeContext datatype;
		public Array_expContext array_exp;
		public TerminalNode VAR() { return getToken(SwiftgrammParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public List<TerminalNode> OPEN_BRACKET() { return getTokens(SwiftgrammParser.OPEN_BRACKET); }
		public TerminalNode OPEN_BRACKET(int i) {
			return getToken(SwiftgrammParser.OPEN_BRACKET, i);
		}
		public DatatypeContext datatype() {
			return getRuleContext(DatatypeContext.class,0);
		}
		public List<TerminalNode> CLOSE_BRACKET() { return getTokens(SwiftgrammParser.CLOSE_BRACKET); }
		public TerminalNode CLOSE_BRACKET(int i) {
			return getToken(SwiftgrammParser.CLOSE_BRACKET, i);
		}
		public TerminalNode ASSIGN() { return getToken(SwiftgrammParser.ASSIGN, 0); }
		public Array_expContext array_exp() {
			return getRuleContext(Array_expContext.class,0);
		}
		public Vector_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vector_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterVector_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitVector_bl(this);
		}
	}

	public final Vector_blContext vector_bl() throws RecognitionException {
		Vector_blContext _localctx = new Vector_blContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_vector_bl);
		try {
			setState(352);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(329);
				match(VAR);
				setState(330);
				((Vector_blContext)_localctx).ID = match(ID);
				setState(331);
				match(COLON);
				setState(332);
				match(OPEN_BRACKET);
				setState(333);
				((Vector_blContext)_localctx).datatype = datatype();
				setState(334);
				match(CLOSE_BRACKET);
				setState(335);
				match(ASSIGN);
				setState(336);
				match(OPEN_BRACKET);
				setState(337);
				((Vector_blContext)_localctx).array_exp = array_exp();
				setState(338);
				match(CLOSE_BRACKET);

				        
				        _localctx.instr = instructions.NewVector((((Vector_blContext)_localctx).ID!=null?((Vector_blContext)_localctx).ID.getText():null),((Vector_blContext)_localctx).datatype.td,((Vector_blContext)_localctx).array_exp.p,(((Vector_blContext)_localctx).ID!=null?((Vector_blContext)_localctx).ID.getLine():0),(((Vector_blContext)_localctx).ID!=null?((Vector_blContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(341);
				match(VAR);
				setState(342);
				((Vector_blContext)_localctx).ID = match(ID);
				setState(343);
				match(COLON);
				setState(344);
				match(OPEN_BRACKET);
				setState(345);
				((Vector_blContext)_localctx).datatype = datatype();
				setState(346);
				match(CLOSE_BRACKET);
				setState(347);
				match(ASSIGN);
				setState(348);
				match(OPEN_BRACKET);
				setState(349);
				match(CLOSE_BRACKET);

				        
				        _localctx.instr = instructions.NewVector((((Vector_blContext)_localctx).ID!=null?((Vector_blContext)_localctx).ID.getText():null),((Vector_blContext)_localctx).datatype.td,nil,(((Vector_blContext)_localctx).ID!=null?((Vector_blContext)_localctx).ID.getLine():0),(((Vector_blContext)_localctx).ID!=null?((Vector_blContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Array_expContext extends ParserRuleContext {
		public []interface{} p;
		public ExpressionContext expression;
		public Array_expContext array_exp;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(SwiftgrammParser.COMMA, 0); }
		public Array_expContext array_exp() {
			return getRuleContext(Array_expContext.class,0);
		}
		public Array_expContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_exp; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterArray_exp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitArray_exp(this);
		}
	}

	public final Array_expContext array_exp() throws RecognitionException {
		Array_expContext _localctx = new Array_expContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_array_exp);
		try {
			setState(362);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(354);
				((Array_expContext)_localctx).expression = expression(0);
				setState(355);
				match(COMMA);
				setState(356);
				((Array_expContext)_localctx).array_exp = array_exp();

				        _localctx.p = append([]interface{}{((Array_expContext)_localctx).expression.p},((Array_expContext)_localctx).array_exp.p...)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(359);
				((Array_expContext)_localctx).expression = expression(0);

				        _localctx.p = []interface{}{((Array_expContext)_localctx).expression.p}
				        

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Array_functionsContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public Token APPEND;
		public ExpressionContext expression;
		public Token REMOVELAST;
		public Token REMOVE;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode DOT() { return getToken(SwiftgrammParser.DOT, 0); }
		public TerminalNode APPEND() { return getToken(SwiftgrammParser.APPEND, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public TerminalNode REMOVELAST() { return getToken(SwiftgrammParser.REMOVELAST, 0); }
		public TerminalNode REMOVE() { return getToken(SwiftgrammParser.REMOVE, 0); }
		public TerminalNode AT() { return getToken(SwiftgrammParser.AT, 0); }
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public Array_functionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_functions; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterArray_functions(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitArray_functions(this);
		}
	}

	public final Array_functionsContext array_functions() throws RecognitionException {
		Array_functionsContext _localctx = new Array_functionsContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_array_functions);
		try {
			setState(388);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(364);
				((Array_functionsContext)_localctx).ID = match(ID);
				setState(365);
				match(DOT);
				setState(366);
				((Array_functionsContext)_localctx).APPEND = match(APPEND);
				setState(367);
				match(OPEN_PARENTHESIS);
				setState(368);
				((Array_functionsContext)_localctx).expression = expression(0);
				setState(369);
				match(CLOSE_PARENTHESIS);

				        _localctx.instr = instructions.NewArrayFunctions((((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getText():null),(((Array_functionsContext)_localctx).APPEND!=null?((Array_functionsContext)_localctx).APPEND.getText():null),((Array_functionsContext)_localctx).expression.p,(((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getLine():0),(((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(372);
				((Array_functionsContext)_localctx).ID = match(ID);
				setState(373);
				match(DOT);
				setState(374);
				((Array_functionsContext)_localctx).REMOVELAST = match(REMOVELAST);
				setState(375);
				match(OPEN_PARENTHESIS);
				setState(376);
				match(CLOSE_PARENTHESIS);

				        _localctx.instr = instructions.NewArrayFunctions((((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getText():null),(((Array_functionsContext)_localctx).REMOVELAST!=null?((Array_functionsContext)_localctx).REMOVELAST.getText():null),nil,(((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getLine():0),(((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(378);
				((Array_functionsContext)_localctx).ID = match(ID);
				setState(379);
				match(DOT);
				setState(380);
				((Array_functionsContext)_localctx).REMOVE = match(REMOVE);
				setState(381);
				match(OPEN_PARENTHESIS);
				setState(382);
				match(AT);
				setState(383);
				match(COLON);
				setState(384);
				((Array_functionsContext)_localctx).expression = expression(0);
				setState(385);
				match(CLOSE_PARENTHESIS);

				        _localctx.instr = instructions.NewArrayFunctions((((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getText():null),(((Array_functionsContext)_localctx).REMOVE!=null?((Array_functionsContext)_localctx).REMOVE.getText():null),((Array_functionsContext)_localctx).expression.p,(((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getLine():0),(((Array_functionsContext)_localctx).ID!=null?((Array_functionsContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_vectorContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public ExpressionContext expression;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode OPEN_BRACKET() { return getToken(SwiftgrammParser.OPEN_BRACKET, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CLOSE_BRACKET() { return getToken(SwiftgrammParser.CLOSE_BRACKET, 0); }
		public TerminalNode ASSIGN() { return getToken(SwiftgrammParser.ASSIGN, 0); }
		public Assign_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_vector; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterAssign_vector(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitAssign_vector(this);
		}
	}

	public final Assign_vectorContext assign_vector() throws RecognitionException {
		Assign_vectorContext _localctx = new Assign_vectorContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_assign_vector);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(390);
			((Assign_vectorContext)_localctx).ID = match(ID);
			setState(391);
			match(OPEN_BRACKET);
			setState(392);
			((Assign_vectorContext)_localctx).expression = expression(0);
			setState(393);
			match(CLOSE_BRACKET);
			setState(394);
			match(ASSIGN);
			setState(395);
			((Assign_vectorContext)_localctx).expression = expression(0);

			        _localctx.instr = instructions.NewAsignVector((((Assign_vectorContext)_localctx).ID!=null?((Assign_vectorContext)_localctx).ID.getText():null),((Assign_vectorContext)_localctx).expression.p,((Assign_vectorContext)_localctx).expression.p,(((Assign_vectorContext)_localctx).ID!=null?((Assign_vectorContext)_localctx).ID.getLine():0),(((Assign_vectorContext)_localctx).ID!=null?((Assign_vectorContext)_localctx).ID.getCharPositionInLine():0))

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public DatatypeContext datatype;
		public BlockContext block;
		public ParamsContext params;
		public TerminalNode FUNC() { return getToken(SwiftgrammParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public TerminalNode ARROW() { return getToken(SwiftgrammParser.ARROW, 0); }
		public DatatypeContext datatype() {
			return getRuleContext(DatatypeContext.class,0);
		}
		public TerminalNode OPEN_kEY() { return getToken(SwiftgrammParser.OPEN_kEY, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode CLOSE_kEY() { return getToken(SwiftgrammParser.CLOSE_kEY, 0); }
		public ParamsContext params() {
			return getRuleContext(ParamsContext.class,0);
		}
		public Function_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterFunction_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitFunction_bl(this);
		}
	}

	public final Function_blContext function_bl() throws RecognitionException {
		Function_blContext _localctx = new Function_blContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_function_bl);
		try {
			setState(440);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(398);
				match(FUNC);
				setState(399);
				((Function_blContext)_localctx).ID = match(ID);
				setState(400);
				match(OPEN_PARENTHESIS);
				setState(401);
				match(CLOSE_PARENTHESIS);
				setState(402);
				match(ARROW);
				setState(403);
				((Function_blContext)_localctx).datatype = datatype();
				setState(404);
				match(OPEN_kEY);
				setState(405);
				((Function_blContext)_localctx).block = block();
				setState(406);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewDeclareFunction((((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getText():null),((Function_blContext)_localctx).datatype.td,[]interface{}{},((Function_blContext)_localctx).block.blk,(((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getLine():0),(((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(409);
				match(FUNC);
				setState(410);
				((Function_blContext)_localctx).ID = match(ID);
				setState(411);
				match(OPEN_PARENTHESIS);
				setState(412);
				match(CLOSE_PARENTHESIS);
				setState(413);
				match(OPEN_kEY);
				setState(414);
				((Function_blContext)_localctx).block = block();
				setState(415);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewDeclareFunction((((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getText():null),symbol.NIL,[]interface{}{},((Function_blContext)_localctx).block.blk,(((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getLine():0),(((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(418);
				match(FUNC);
				setState(419);
				((Function_blContext)_localctx).ID = match(ID);
				setState(420);
				match(OPEN_PARENTHESIS);
				setState(421);
				((Function_blContext)_localctx).params = params();
				setState(422);
				match(CLOSE_PARENTHESIS);
				setState(423);
				match(ARROW);
				setState(424);
				((Function_blContext)_localctx).datatype = datatype();
				setState(425);
				match(OPEN_kEY);
				setState(426);
				((Function_blContext)_localctx).block = block();
				setState(427);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewDeclareFunction((((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getText():null),((Function_blContext)_localctx).datatype.td,((Function_blContext)_localctx).params.p,((Function_blContext)_localctx).block.blk,(((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getLine():0),(((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(430);
				match(FUNC);
				setState(431);
				((Function_blContext)_localctx).ID = match(ID);
				setState(432);
				match(OPEN_PARENTHESIS);
				setState(433);
				((Function_blContext)_localctx).params = params();
				setState(434);
				match(CLOSE_PARENTHESIS);
				setState(435);
				match(OPEN_kEY);
				setState(436);
				((Function_blContext)_localctx).block = block();
				setState(437);
				match(CLOSE_kEY);

				        //fmt.Println(((Function_blContext)_localctx).params.p)
				        _localctx.instr = instructions.NewDeclareFunction((((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getText():null),symbol.NIL,((Function_blContext)_localctx).params.p,((Function_blContext)_localctx).block.blk,(((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getLine():0),(((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamsContext extends ParserRuleContext {
		public []interface{} p;
		public Token ID;
		public DatatypeContext datatype;
		public ParamsContext params;
		public Extern_paramsContext extern_params;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public DatatypeContext datatype() {
			return getRuleContext(DatatypeContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(SwiftgrammParser.COMMA, 0); }
		public ParamsContext params() {
			return getRuleContext(ParamsContext.class,0);
		}
		public Extern_paramsContext extern_params() {
			return getRuleContext(Extern_paramsContext.class,0);
		}
		public TerminalNode OPEN_BRACKET() { return getToken(SwiftgrammParser.OPEN_BRACKET, 0); }
		public TerminalNode CLOSE_BRACKET() { return getToken(SwiftgrammParser.CLOSE_BRACKET, 0); }
		public ParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_params; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterParams(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitParams(this);
		}
	}

	public final ParamsContext params() throws RecognitionException {
		ParamsContext _localctx = new ParamsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_params);
		try {
			setState(502);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(442);
				((ParamsContext)_localctx).ID = match(ID);
				setState(443);
				match(COLON);
				setState(444);
				((ParamsContext)_localctx).datatype = datatype();
				setState(445);
				match(COMMA);
				setState(446);
				((ParamsContext)_localctx).params = params();

				        value:=instructions.NewParamFunction(instructions.NewDeclareWithoutValue((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,expressions.NewNative(nil,symbol.NIL,(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null))
				        _localctx.p = append([]interface{}{value},((ParamsContext)_localctx).params.p...)


				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(449);
				((ParamsContext)_localctx).ID = match(ID);
				setState(450);
				match(COLON);
				setState(451);
				((ParamsContext)_localctx).datatype = datatype();

				        value := instructions.NewParamFunction(instructions.NewDeclareWithoutValue((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,expressions.NewNative(nil,symbol.NIL,(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null))
				        _localctx.p = []interface{}{value}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(454);
				((ParamsContext)_localctx).extern_params = extern_params();
				setState(455);
				((ParamsContext)_localctx).ID = match(ID);
				setState(456);
				match(COLON);
				setState(457);
				((ParamsContext)_localctx).datatype = datatype();
				setState(458);
				match(COMMA);
				setState(459);
				((ParamsContext)_localctx).params = params();

				        value:=instructions.NewParamFunction(instructions.NewDeclareWithoutValue((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,expressions.NewNative(nil,symbol.NIL,(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).extern_params!=null?_input.getText(((ParamsContext)_localctx).extern_params.start,((ParamsContext)_localctx).extern_params.stop):null))
				        _localctx.p = append([]interface{}{value},((ParamsContext)_localctx).params.p...)


				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(462);
				((ParamsContext)_localctx).extern_params = extern_params();
				setState(463);
				((ParamsContext)_localctx).ID = match(ID);
				setState(464);
				match(COLON);
				setState(465);
				((ParamsContext)_localctx).datatype = datatype();

				        value := instructions.NewParamFunction(instructions.NewDeclareWithoutValue((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,expressions.NewNative(nil,symbol.NIL,(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).extern_params!=null?_input.getText(((ParamsContext)_localctx).extern_params.start,((ParamsContext)_localctx).extern_params.stop):null))
				        _localctx.p = []interface{}{value}

				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(468);
				((ParamsContext)_localctx).ID = match(ID);
				setState(469);
				match(COLON);
				setState(470);
				match(OPEN_BRACKET);
				setState(471);
				((ParamsContext)_localctx).datatype = datatype();
				setState(472);
				match(CLOSE_BRACKET);
				setState(473);
				match(COMMA);
				setState(474);
				((ParamsContext)_localctx).params = params();

				        value:=instructions.NewParamFunction(instructions.NewVector((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,nil,(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null))
				        _localctx.p = append([]interface{}{value},((ParamsContext)_localctx).params.p...)


				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(477);
				((ParamsContext)_localctx).ID = match(ID);
				setState(478);
				match(COLON);
				setState(479);
				match(OPEN_BRACKET);
				setState(480);
				((ParamsContext)_localctx).datatype = datatype();
				setState(481);
				match(CLOSE_BRACKET);

				        value := instructions.NewParamFunction(instructions.NewVector((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,nil,(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null))
				        _localctx.p = []interface{}{value}

				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(484);
				((ParamsContext)_localctx).extern_params = extern_params();
				setState(485);
				((ParamsContext)_localctx).ID = match(ID);
				setState(486);
				match(COLON);
				setState(487);
				match(OPEN_BRACKET);
				setState(488);
				((ParamsContext)_localctx).datatype = datatype();
				setState(489);
				match(CLOSE_BRACKET);
				setState(490);
				match(COMMA);
				setState(491);
				((ParamsContext)_localctx).params = params();

				        value:=instructions.NewParamFunction(instructions.NewVector((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,nil,(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).extern_params!=null?_input.getText(((ParamsContext)_localctx).extern_params.start,((ParamsContext)_localctx).extern_params.stop):null))
				        _localctx.p = append([]interface{}{value},((ParamsContext)_localctx).params.p...)


				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(494);
				((ParamsContext)_localctx).extern_params = extern_params();
				setState(495);
				((ParamsContext)_localctx).ID = match(ID);
				setState(496);
				match(COLON);
				setState(497);
				match(OPEN_BRACKET);
				setState(498);
				((ParamsContext)_localctx).datatype = datatype();
				setState(499);
				match(CLOSE_BRACKET);

				        value := instructions.NewParamFunction(instructions.NewVector((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,nil,(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getLine():0),(((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getCharPositionInLine():0)),(((ParamsContext)_localctx).extern_params!=null?_input.getText(((ParamsContext)_localctx).extern_params.start,((ParamsContext)_localctx).extern_params.stop):null))
				        _localctx.p = []interface{}{value}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Extern_paramsContext extends ParserRuleContext {
		public string p;
		public Token ID;
		public Token UNDERSCORE;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode UNDERSCORE() { return getToken(SwiftgrammParser.UNDERSCORE, 0); }
		public Extern_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extern_params; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterExtern_params(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitExtern_params(this);
		}
	}

	public final Extern_paramsContext extern_params() throws RecognitionException {
		Extern_paramsContext _localctx = new Extern_paramsContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_extern_params);
		try {
			setState(508);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(504);
				((Extern_paramsContext)_localctx).ID = match(ID);

				        _localctx.p = (((Extern_paramsContext)_localctx).ID!=null?((Extern_paramsContext)_localctx).ID.getText():null)

				}
				break;
			case UNDERSCORE:
				enterOuterAlt(_localctx, 2);
				{
				setState(506);
				((Extern_paramsContext)_localctx).UNDERSCORE = match(UNDERSCORE);

				        _localctx.p = (((Extern_paramsContext)_localctx).UNDERSCORE!=null?((Extern_paramsContext)_localctx).UNDERSCORE.getText():null)

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Call_functionContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public List_expContext list_exp;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public List_expContext list_exp() {
			return getRuleContext(List_expContext.class,0);
		}
		public Call_functionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_call_function; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterCall_function(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitCall_function(this);
		}
	}

	public final Call_functionContext call_function() throws RecognitionException {
		Call_functionContext _localctx = new Call_functionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_call_function);
		try {
			setState(520);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(510);
				((Call_functionContext)_localctx).ID = match(ID);
				setState(511);
				match(OPEN_PARENTHESIS);
				setState(512);
				match(CLOSE_PARENTHESIS);

				        _localctx.instr = instructions.NewCallFunction((((Call_functionContext)_localctx).ID!=null?((Call_functionContext)_localctx).ID.getText():null),[]interface{}{},(((Call_functionContext)_localctx).ID!=null?((Call_functionContext)_localctx).ID.getLine():0),(((Call_functionContext)_localctx).ID!=null?((Call_functionContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(514);
				((Call_functionContext)_localctx).ID = match(ID);
				setState(515);
				match(OPEN_PARENTHESIS);
				setState(516);
				((Call_functionContext)_localctx).list_exp = list_exp();
				setState(517);
				match(CLOSE_PARENTHESIS);

				        _localctx.instr = instructions.NewCallFunction((((Call_functionContext)_localctx).ID!=null?((Call_functionContext)_localctx).ID.getText():null),((Call_functionContext)_localctx).list_exp.p,(((Call_functionContext)_localctx).ID!=null?((Call_functionContext)_localctx).ID.getLine():0),(((Call_functionContext)_localctx).ID!=null?((Call_functionContext)_localctx).ID.getCharPositionInLine():0))

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_expContext extends ParserRuleContext {
		public []interface{} p;
		public ExpressionContext expression;
		public List_expContext list_exp;
		public Token ID;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(SwiftgrammParser.COMMA, 0); }
		public List_expContext list_exp() {
			return getRuleContext(List_expContext.class,0);
		}
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public List_expContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_exp; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterList_exp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitList_exp(this);
		}
	}

	public final List_expContext list_exp() throws RecognitionException {
		List_expContext _localctx = new List_expContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_list_exp);
		try {
			setState(542);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(522);
				((List_expContext)_localctx).expression = expression(0);
				setState(523);
				match(COMMA);
				setState(524);
				((List_expContext)_localctx).list_exp = list_exp();

				        value := instructions.NewParamCall(((List_expContext)_localctx).expression.p,"_")
				        _localctx.p = append([]interface{}{value},((List_expContext)_localctx).list_exp.p...)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(527);
				((List_expContext)_localctx).expression = expression(0);

				        value := instructions.NewParamCall(((List_expContext)_localctx).expression.p,"_")
				        _localctx.p = []interface{}{value}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(530);
				((List_expContext)_localctx).ID = match(ID);
				setState(531);
				match(COLON);
				setState(532);
				((List_expContext)_localctx).expression = expression(0);
				setState(533);
				match(COMMA);
				setState(534);
				((List_expContext)_localctx).list_exp = list_exp();

				        value := instructions.NewParamCall(((List_expContext)_localctx).expression.p,(((List_expContext)_localctx).ID!=null?((List_expContext)_localctx).ID.getText():null))
				        _localctx.p = append([]interface{}{value},((List_expContext)_localctx).list_exp.p...)

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(537);
				((List_expContext)_localctx).ID = match(ID);
				setState(538);
				match(COLON);
				setState(539);
				((List_expContext)_localctx).expression = expression(0);

				        value := instructions.NewParamCall(((List_expContext)_localctx).expression.p,(((List_expContext)_localctx).ID!=null?((List_expContext)_localctx).ID.getText():null))
				        _localctx.p = []interface{}{value}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Call_function_expContext extends ParserRuleContext {
		public abstract.Expression p;
		public Call_functionContext call_function;
		public Call_functionContext call_function() {
			return getRuleContext(Call_functionContext.class,0);
		}
		public Call_function_expContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_call_function_exp; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterCall_function_exp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitCall_function_exp(this);
		}
	}

	public final Call_function_expContext call_function_exp() throws RecognitionException {
		Call_function_expContext _localctx = new Call_function_expContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_call_function_exp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(544);
			((Call_function_expContext)_localctx).call_function = call_function();

			        _localctx.p = expressions.NewCallFunctionExp(((Call_function_expContext)_localctx).call_function.instr)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Declare_array_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public Type_matrixContext type_matrix;
		public Exp_matrizContext exp_matriz;
		public TerminalNode VAR() { return getToken(SwiftgrammParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public Type_matrixContext type_matrix() {
			return getRuleContext(Type_matrixContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(SwiftgrammParser.ASSIGN, 0); }
		public TerminalNode OPEN_BRACKET() { return getToken(SwiftgrammParser.OPEN_BRACKET, 0); }
		public Exp_matrizContext exp_matriz() {
			return getRuleContext(Exp_matrizContext.class,0);
		}
		public TerminalNode CLOSE_BRACKET() { return getToken(SwiftgrammParser.CLOSE_BRACKET, 0); }
		public Declare_array_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declare_array_bl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterDeclare_array_bl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitDeclare_array_bl(this);
		}
	}

	public final Declare_array_blContext declare_array_bl() throws RecognitionException {
		Declare_array_blContext _localctx = new Declare_array_blContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_declare_array_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(547);
			match(VAR);
			setState(548);
			((Declare_array_blContext)_localctx).ID = match(ID);
			setState(549);
			match(COLON);
			setState(550);
			((Declare_array_blContext)_localctx).type_matrix = type_matrix();
			setState(551);
			match(ASSIGN);
			setState(552);
			match(OPEN_BRACKET);
			setState(553);
			((Declare_array_blContext)_localctx).exp_matriz = exp_matriz();
			setState(554);
			match(CLOSE_BRACKET);

			        _localctx.instr = instructions.NewDeclareArray((((Declare_array_blContext)_localctx).ID!=null?((Declare_array_blContext)_localctx).ID.getText():null),((Declare_array_blContext)_localctx).type_matrix.td,CountM,((Declare_array_blContext)_localctx).exp_matriz.p)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Exp_matrizContext extends ParserRuleContext {
		public []interface{} p;
		public ExpressionContext expression;
		public Exp_matrizContext exp_matriz;
		public TerminalNode OPEN_BRACKET() { return getToken(SwiftgrammParser.OPEN_BRACKET, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CLOSE_BRACKET() { return getToken(SwiftgrammParser.CLOSE_BRACKET, 0); }
		public TerminalNode COMMA() { return getToken(SwiftgrammParser.COMMA, 0); }
		public Exp_matrizContext exp_matriz() {
			return getRuleContext(Exp_matrizContext.class,0);
		}
		public Exp_matrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp_matriz; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterExp_matriz(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitExp_matriz(this);
		}
	}

	public final Exp_matrizContext exp_matriz() throws RecognitionException {
		Exp_matrizContext _localctx = new Exp_matrizContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_exp_matriz);
		try {
			setState(569);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(557);
				match(OPEN_BRACKET);
				setState(558);
				((Exp_matrizContext)_localctx).expression = expression(0);
				setState(559);
				match(CLOSE_BRACKET);

				        value := ((Exp_matrizContext)_localctx).expression.p
				        _localctx.p = []interface{}{value}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(562);
				match(OPEN_BRACKET);
				setState(563);
				((Exp_matrizContext)_localctx).expression = expression(0);
				setState(564);
				match(CLOSE_BRACKET);
				setState(565);
				match(COMMA);
				setState(566);
				((Exp_matrizContext)_localctx).exp_matriz = exp_matriz();
				 
				        value := ((Exp_matrizContext)_localctx).expression.p
				        _localctx.p = append([]interface{}{value},((Exp_matrizContext)_localctx).exp_matriz.p...)

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Type_matrixContext extends ParserRuleContext {
		public symbol.TypeData td;
		public Type_matrixContext type_matrix;
		public DatatypeContext datatype;
		public TerminalNode OPEN_BRACKET() { return getToken(SwiftgrammParser.OPEN_BRACKET, 0); }
		public Type_matrixContext type_matrix() {
			return getRuleContext(Type_matrixContext.class,0);
		}
		public TerminalNode CLOSE_BRACKET() { return getToken(SwiftgrammParser.CLOSE_BRACKET, 0); }
		public DatatypeContext datatype() {
			return getRuleContext(DatatypeContext.class,0);
		}
		public Type_matrixContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_matrix; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterType_matrix(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitType_matrix(this);
		}
	}

	public final Type_matrixContext type_matrix() throws RecognitionException {
		Type_matrixContext _localctx = new Type_matrixContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_type_matrix);
		try {
			setState(581);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(571);
				match(OPEN_BRACKET);
				setState(572);
				((Type_matrixContext)_localctx).type_matrix = type_matrix();
				setState(573);
				match(CLOSE_BRACKET);

				        CountM++;
				        _localctx.td = ((Type_matrixContext)_localctx).type_matrix.td

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(576);
				match(OPEN_BRACKET);
				setState(577);
				((Type_matrixContext)_localctx).datatype = datatype();
				setState(578);
				match(CLOSE_BRACKET);

				        CountM++;
				        _localctx.td = ((Type_matrixContext)_localctx).datatype.td

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Definition_matrixContext extends ParserRuleContext {
		public [][]interface{} p;
		public Type_matrixContext type_matrix() {
			return getRuleContext(Type_matrixContext.class,0);
		}
		public Definition_matrixContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_definition_matrix; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterDefinition_matrix(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitDefinition_matrix(this);
		}
	}

	public final Definition_matrixContext definition_matrix() throws RecognitionException {
		Definition_matrixContext _localctx = new Definition_matrixContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_definition_matrix);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(583);
			type_matrix();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public abstract.Expression p;
		public ExpressionContext left;
		public Token oper;
		public ExpressionContext expression;
		public Call_function_expContext call_function_exp;
		public Native_functionContext native_function;
		public Token ID;
		public Token OPEN_BRACKET;
		public Token COUNT;
		public Token ISEMPTY;
		public Array_expContext array_exp;
		public Token NUMBER;
		public Token FLOATT;
		public Token STRING_LITERAL;
		public Token CHARACTER_LITERAL;
		public Token TRUE;
		public Token FALSE;
		public Token NIL;
		public ExpressionContext right;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode NOT() { return getToken(SwiftgrammParser.NOT, 0); }
		public TerminalNode SUBTRACTION() { return getToken(SwiftgrammParser.SUBTRACTION, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public Call_function_expContext call_function_exp() {
			return getRuleContext(Call_function_expContext.class,0);
		}
		public Native_functionContext native_function() {
			return getRuleContext(Native_functionContext.class,0);
		}
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode OPEN_BRACKET() { return getToken(SwiftgrammParser.OPEN_BRACKET, 0); }
		public TerminalNode CLOSE_BRACKET() { return getToken(SwiftgrammParser.CLOSE_BRACKET, 0); }
		public TerminalNode DOT() { return getToken(SwiftgrammParser.DOT, 0); }
		public TerminalNode COUNT() { return getToken(SwiftgrammParser.COUNT, 0); }
		public TerminalNode ISEMPTY() { return getToken(SwiftgrammParser.ISEMPTY, 0); }
		public Array_expContext array_exp() {
			return getRuleContext(Array_expContext.class,0);
		}
		public TerminalNode NUMBER() { return getToken(SwiftgrammParser.NUMBER, 0); }
		public TerminalNode FLOATT() { return getToken(SwiftgrammParser.FLOATT, 0); }
		public TerminalNode STRING_LITERAL() { return getToken(SwiftgrammParser.STRING_LITERAL, 0); }
		public TerminalNode CHARACTER_LITERAL() { return getToken(SwiftgrammParser.CHARACTER_LITERAL, 0); }
		public TerminalNode TRUE() { return getToken(SwiftgrammParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(SwiftgrammParser.FALSE, 0); }
		public TerminalNode NIL() { return getToken(SwiftgrammParser.NIL, 0); }
		public TerminalNode MULTIPLICATION() { return getToken(SwiftgrammParser.MULTIPLICATION, 0); }
		public TerminalNode DIVISION() { return getToken(SwiftgrammParser.DIVISION, 0); }
		public TerminalNode SUMMATION() { return getToken(SwiftgrammParser.SUMMATION, 0); }
		public TerminalNode MOD() { return getToken(SwiftgrammParser.MOD, 0); }
		public TerminalNode LESS_THAN() { return getToken(SwiftgrammParser.LESS_THAN, 0); }
		public TerminalNode LESS_THAN_EQUAL() { return getToken(SwiftgrammParser.LESS_THAN_EQUAL, 0); }
		public TerminalNode GREATER_THAN() { return getToken(SwiftgrammParser.GREATER_THAN, 0); }
		public TerminalNode GREATER_THAN_EQUAL() { return getToken(SwiftgrammParser.GREATER_THAN_EQUAL, 0); }
		public TerminalNode EQUAL() { return getToken(SwiftgrammParser.EQUAL, 0); }
		public TerminalNode NOT_EQUAL() { return getToken(SwiftgrammParser.NOT_EQUAL, 0); }
		public TerminalNode AND() { return getToken(SwiftgrammParser.AND, 0); }
		public TerminalNode OR() { return getToken(SwiftgrammParser.OR, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitExpression(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 66;
		enterRecursionRule(_localctx, 66, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(640);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				{
				setState(586);
				((ExpressionContext)_localctx).oper = match(NOT);
				setState(587);
				((ExpressionContext)_localctx).expression = expression(17);

				                _localctx.p = expressions.NewLogicOperations(nil,((ExpressionContext)_localctx).expression.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
				        
				}
				break;
			case 2:
				{
				setState(590);
				((ExpressionContext)_localctx).oper = match(SUBTRACTION);
				setState(591);
				((ExpressionContext)_localctx).expression = expression(16);

				                _localctx.p = expressions.NewNegative(((ExpressionContext)_localctx).expression.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
				        
				}
				break;
			case 3:
				{
				setState(594);
				match(OPEN_PARENTHESIS);
				setState(595);
				((ExpressionContext)_localctx).expression = expression(0);
				setState(596);
				match(CLOSE_PARENTHESIS);

				                _localctx.p = ((ExpressionContext)_localctx).expression.p
				        
				}
				break;
			case 4:
				{
				setState(599);
				((ExpressionContext)_localctx).call_function_exp = call_function_exp();

				                _localctx.p = ((ExpressionContext)_localctx).call_function_exp.p
				        
				}
				break;
			case 5:
				{
				setState(602);
				((ExpressionContext)_localctx).native_function = native_function();

				                _localctx.p = ((ExpressionContext)_localctx).native_function.p
				        
				}
				break;
			case 6:
				{
				setState(605);
				((ExpressionContext)_localctx).ID = match(ID);
				setState(606);
				((ExpressionContext)_localctx).OPEN_BRACKET = match(OPEN_BRACKET);
				setState(607);
				((ExpressionContext)_localctx).expression = expression(0);
				setState(608);
				match(CLOSE_BRACKET);

				                _localctx.p = expressions.NewGetPosVector((((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getText():null),((ExpressionContext)_localctx).expression.p,(((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getLine():0),(((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getCharPositionInLine():0))
				        
				}
				break;
			case 7:
				{
				setState(611);
				((ExpressionContext)_localctx).ID = match(ID);
				setState(612);
				match(DOT);
				setState(613);
				((ExpressionContext)_localctx).COUNT = match(COUNT);

				                _localctx.p = expressions.NewVectorValues((((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getText():null),(((ExpressionContext)_localctx).COUNT!=null?((ExpressionContext)_localctx).COUNT.getText():null),(((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getLine():0),(((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getCharPositionInLine():0))
				        
				}
				break;
			case 8:
				{
				setState(615);
				((ExpressionContext)_localctx).ID = match(ID);
				setState(616);
				match(DOT);
				setState(617);
				((ExpressionContext)_localctx).ISEMPTY = match(ISEMPTY);

				                _localctx.p = expressions.NewVectorValues((((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getText():null),(((ExpressionContext)_localctx).ISEMPTY!=null?((ExpressionContext)_localctx).ISEMPTY.getText():null),(((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getLine():0),(((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getCharPositionInLine():0))
				        
				}
				break;
			case 9:
				{
				setState(619);
				((ExpressionContext)_localctx).OPEN_BRACKET = match(OPEN_BRACKET);
				setState(620);
				((ExpressionContext)_localctx).array_exp = array_exp();
				setState(621);
				match(CLOSE_BRACKET);

				                _localctx.p = expressions.NewNative(((ExpressionContext)_localctx).array_exp.p,symbol.ARRAY,(((ExpressionContext)_localctx).OPEN_BRACKET!=null?((ExpressionContext)_localctx).OPEN_BRACKET.getLine():0),(((ExpressionContext)_localctx).OPEN_BRACKET!=null?((ExpressionContext)_localctx).OPEN_BRACKET.getCharPositionInLine():0))
				        
				}
				break;
			case 10:
				{
				setState(624);
				((ExpressionContext)_localctx).NUMBER = match(NUMBER);

				                value,err := strconv.Atoi((((ExpressionContext)_localctx).NUMBER!=null?((ExpressionContext)_localctx).NUMBER.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.INT,(((ExpressionContext)_localctx).NUMBER!=null?((ExpressionContext)_localctx).NUMBER.getLine():0),(((ExpressionContext)_localctx).NUMBER!=null?((ExpressionContext)_localctx).NUMBER.getCharPositionInLine():0))
				        
				}
				break;
			case 11:
				{
				setState(626);
				((ExpressionContext)_localctx).FLOATT = match(FLOATT);

				                value,err := strconv.ParseFloat((((ExpressionContext)_localctx).FLOATT!=null?((ExpressionContext)_localctx).FLOATT.getText():null),64)
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.FLOAT,(((ExpressionContext)_localctx).FLOATT!=null?((ExpressionContext)_localctx).FLOATT.getLine():0),(((ExpressionContext)_localctx).FLOATT!=null?((ExpressionContext)_localctx).FLOATT.getCharPositionInLine():0))
				        
				}
				break;
			case 12:
				{
				setState(628);
				((ExpressionContext)_localctx).STRING_LITERAL = match(STRING_LITERAL);

				                value := (((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getText():null)[1:len((((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getText():null))-1]
				                _localctx.p = expressions.NewNative(value,symbol.STRING,(((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getLine():0),(((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getCharPositionInLine():0))
				        
				}
				break;
			case 13:
				{
				setState(630);
				((ExpressionContext)_localctx).CHARACTER_LITERAL = match(CHARACTER_LITERAL);

				                value := (((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getText():null)[1:len((((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getText():null))-1]
				                _localctx.p = expressions.NewNative(value,symbol.CHAR,(((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getLine():0),(((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getCharPositionInLine():0))
				        
				}
				break;
			case 14:
				{
				setState(632);
				((ExpressionContext)_localctx).TRUE = match(TRUE);

				                value,err := strconv.ParseBool((((ExpressionContext)_localctx).TRUE!=null?((ExpressionContext)_localctx).TRUE.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.BOOL,(((ExpressionContext)_localctx).TRUE!=null?((ExpressionContext)_localctx).TRUE.getLine():0),(((ExpressionContext)_localctx).TRUE!=null?((ExpressionContext)_localctx).TRUE.getCharPositionInLine():0)) 
				        
				}
				break;
			case 15:
				{
				setState(634);
				((ExpressionContext)_localctx).FALSE = match(FALSE);

				                value,err := strconv.ParseBool((((ExpressionContext)_localctx).FALSE!=null?((ExpressionContext)_localctx).FALSE.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.BOOL,(((ExpressionContext)_localctx).FALSE!=null?((ExpressionContext)_localctx).FALSE.getLine():0),(((ExpressionContext)_localctx).FALSE!=null?((ExpressionContext)_localctx).FALSE.getCharPositionInLine():0)) 
				        
				}
				break;
			case 16:
				{
				setState(636);
				((ExpressionContext)_localctx).ID = match(ID);

				                _localctx.p = expressions.NewIdentifier((((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getText():null),(((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getLine():0),(((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getCharPositionInLine():0))
				        
				}
				break;
			case 17:
				{
				setState(638);
				((ExpressionContext)_localctx).NIL = match(NIL);

				                _localctx.p = expressions.NewNative(nil,symbol.NIL,(((ExpressionContext)_localctx).NIL!=null?((ExpressionContext)_localctx).NIL.getLine():0),(((ExpressionContext)_localctx).NIL!=null?((ExpressionContext)_localctx).NIL.getCharPositionInLine():0))
				        
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(679);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(677);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(642);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(643);
						((ExpressionContext)_localctx).oper = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MULTIPLICATION || _la==DIVISION) ) {
							((ExpressionContext)_localctx).oper = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(644);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(25);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
						                  
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(647);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(648);
						((ExpressionContext)_localctx).oper = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==SUMMATION || _la==SUBTRACTION) ) {
							((ExpressionContext)_localctx).oper = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(649);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(24);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
						                  
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(652);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(653);
						((ExpressionContext)_localctx).oper = match(MOD);
						setState(654);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(23);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
						                  
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(657);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(658);
						((ExpressionContext)_localctx).oper = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==LESS_THAN || _la==LESS_THAN_EQUAL) ) {
							((ExpressionContext)_localctx).oper = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(659);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(22);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
						                  
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(662);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(663);
						((ExpressionContext)_localctx).oper = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==GREATER_THAN || _la==GREATER_THAN_EQUAL) ) {
							((ExpressionContext)_localctx).oper = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(664);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(21);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
						                  
						}
						break;
					case 6:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(667);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(668);
						((ExpressionContext)_localctx).oper = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQUAL || _la==NOT_EQUAL) ) {
							((ExpressionContext)_localctx).oper = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(669);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(20);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
						                  
						}
						break;
					case 7:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(672);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(673);
						((ExpressionContext)_localctx).oper = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OR || _la==AND) ) {
							((ExpressionContext)_localctx).oper = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(674);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(19);

						                          _localctx.p = expressions.NewLogicOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getLine():0),(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getCharPositionInLine():0))
						                  
						}
						break;
					}
					} 
				}
				setState(681);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DatatypeContext extends ParserRuleContext {
		public symbol.TypeData td;
		public TerminalNode INT() { return getToken(SwiftgrammParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftgrammParser.FLOAT, 0); }
		public TerminalNode STRING() { return getToken(SwiftgrammParser.STRING, 0); }
		public TerminalNode BOOL() { return getToken(SwiftgrammParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftgrammParser.CHARACTER, 0); }
		public DatatypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_datatype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).enterDatatype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftgrammListener ) ((SwiftgrammListener)listener).exitDatatype(this);
		}
	}

	public final DatatypeContext datatype() throws RecognitionException {
		DatatypeContext _localctx = new DatatypeContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_datatype);
		try {
			setState(692);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(682);
				match(INT);

				                _localctx.td = symbol.INT
				        
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(684);
				match(FLOAT);

				                _localctx.td = symbol.FLOAT
				        
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(686);
				match(STRING);

				                _localctx.td = symbol.STRING
				        
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(688);
				match(BOOL);

				                _localctx.td = symbol.BOOL
				        
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 5);
				{
				setState(690);
				match(CHARACTER);

				                _localctx.td = symbol.CHAR
				        
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 33:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 24);
		case 1:
			return precpred(_ctx, 23);
		case 2:
			return precpred(_ctx, 22);
		case 3:
			return precpred(_ctx, 21);
		case 4:
			return precpred(_ctx, 20);
		case 5:
			return precpred(_ctx, 19);
		case 6:
			return precpred(_ctx, 18);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001I\u02b7\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0004\u0001L\b"+
		"\u0001\u000b\u0001\f\u0001M\u0001\u0001\u0001\u0001\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002\u008b\b\u0002\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u00a0"+
		"\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0003\b\u00b3\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0003\t\u00c7\b\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003"+
		"\n\u00de\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u00f3\b\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0003\u000e\u010f\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0003\u000f\u012e\b\u000f\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013"+
		"\u0161\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u016b\b\u0014\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015"+
		"\u0185\b\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u01b9\b\u0017\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018"+
		"\u01f7\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003\u0019"+
		"\u01fd\b\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0003\u001a"+
		"\u0209\b\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u021f\b\u001b\u0001\u001c"+
		"\u0001\u001c\u0001\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0003\u001e\u023a\b\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0003\u001f\u0246\b\u001f\u0001 \u0001 \u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0003!\u0281\b!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0005!\u02a6\b!\n!"+
		"\f!\u02a9\t!\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0003\"\u02b5\b\"\u0001\"\u0000\u0001B#\u0000\u0002"+
		"\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e"+
		" \"$&(*,.02468:<>@BD\u0000\u0006\u0001\u0000-.\u0001\u0000+,\u0001\u0000"+
		"67\u0001\u000089\u0001\u000045\u0001\u000012\u02e1\u0000F\u0001\u0000"+
		"\u0000\u0000\u0002K\u0001\u0000\u0000\u0000\u0004\u008a\u0001\u0000\u0000"+
		"\u0000\u0006\u008c\u0001\u0000\u0000\u0000\b\u0091\u0001\u0000\u0000\u0000"+
		"\n\u0096\u0001\u0000\u0000\u0000\f\u009f\u0001\u0000\u0000\u0000\u000e"+
		"\u00a1\u0001\u0000\u0000\u0000\u0010\u00b2\u0001\u0000\u0000\u0000\u0012"+
		"\u00c6\u0001\u0000\u0000\u0000\u0014\u00dd\u0001\u0000\u0000\u0000\u0016"+
		"\u00df\u0001\u0000\u0000\u0000\u0018\u00e4\u0001\u0000\u0000\u0000\u001a"+
		"\u00f2\u0001\u0000\u0000\u0000\u001c\u010e\u0001\u0000\u0000\u0000\u001e"+
		"\u012d\u0001\u0000\u0000\u0000 \u012f\u0001\u0000\u0000\u0000\"\u0136"+
		"\u0001\u0000\u0000\u0000$\u0141\u0001\u0000\u0000\u0000&\u0160\u0001\u0000"+
		"\u0000\u0000(\u016a\u0001\u0000\u0000\u0000*\u0184\u0001\u0000\u0000\u0000"+
		",\u0186\u0001\u0000\u0000\u0000.\u01b8\u0001\u0000\u0000\u00000\u01f6"+
		"\u0001\u0000\u0000\u00002\u01fc\u0001\u0000\u0000\u00004\u0208\u0001\u0000"+
		"\u0000\u00006\u021e\u0001\u0000\u0000\u00008\u0220\u0001\u0000\u0000\u0000"+
		":\u0223\u0001\u0000\u0000\u0000<\u0239\u0001\u0000\u0000\u0000>\u0245"+
		"\u0001\u0000\u0000\u0000@\u0247\u0001\u0000\u0000\u0000B\u0280\u0001\u0000"+
		"\u0000\u0000D\u02b4\u0001\u0000\u0000\u0000FG\u0003\u0002\u0001\u0000"+
		"GH\u0005\u0000\u0000\u0001HI\u0006\u0000\uffff\uffff\u0000I\u0001\u0001"+
		"\u0000\u0000\u0000JL\u0003\u0004\u0002\u0000KJ\u0001\u0000\u0000\u0000"+
		"LM\u0001\u0000\u0000\u0000MK\u0001\u0000\u0000\u0000MN\u0001\u0000\u0000"+
		"\u0000NO\u0001\u0000\u0000\u0000OP\u0006\u0001\uffff\uffff\u0000P\u0003"+
		"\u0001\u0000\u0000\u0000QR\u0003\u0014\n\u0000RS\u0006\u0002\uffff\uffff"+
		"\u0000S\u008b\u0001\u0000\u0000\u0000TU\u0003\u0010\b\u0000UV\u0006\u0002"+
		"\uffff\uffff\u0000V\u008b\u0001\u0000\u0000\u0000WX\u0003\u0016\u000b"+
		"\u0000XY\u0006\u0002\uffff\uffff\u0000Y\u008b\u0001\u0000\u0000\u0000"+
		"Z[\u0003,\u0016\u0000[\\\u0006\u0002\uffff\uffff\u0000\\\u008b\u0001\u0000"+
		"\u0000\u0000]^\u0003\u0018\f\u0000^_\u0006\u0002\uffff\uffff\u0000_\u008b"+
		"\u0001\u0000\u0000\u0000`a\u0003\u001c\u000e\u0000ab\u0006\u0002\uffff"+
		"\uffff\u0000b\u008b\u0001\u0000\u0000\u0000cd\u0003\u0006\u0003\u0000"+
		"de\u0006\u0002\uffff\uffff\u0000e\u008b\u0001\u0000\u0000\u0000fg\u0003"+
		"\b\u0004\u0000gh\u0006\u0002\uffff\uffff\u0000h\u008b\u0001\u0000\u0000"+
		"\u0000ij\u0003 \u0010\u0000jk\u0006\u0002\uffff\uffff\u0000k\u008b\u0001"+
		"\u0000\u0000\u0000lm\u0003\"\u0011\u0000mn\u0006\u0002\uffff\uffff\u0000"+
		"n\u008b\u0001\u0000\u0000\u0000op\u0003$\u0012\u0000pq\u0006\u0002\uffff"+
		"\uffff\u0000q\u008b\u0001\u0000\u0000\u0000rs\u0003\n\u0005\u0000st\u0006"+
		"\u0002\uffff\uffff\u0000t\u008b\u0001\u0000\u0000\u0000uv\u0003\f\u0006"+
		"\u0000vw\u0006\u0002\uffff\uffff\u0000w\u008b\u0001\u0000\u0000\u0000"+
		"xy\u0003\u000e\u0007\u0000yz\u0006\u0002\uffff\uffff\u0000z\u008b\u0001"+
		"\u0000\u0000\u0000{|\u0003&\u0013\u0000|}\u0006\u0002\uffff\uffff\u0000"+
		"}\u008b\u0001\u0000\u0000\u0000~\u007f\u0003.\u0017\u0000\u007f\u0080"+
		"\u0006\u0002\uffff\uffff\u0000\u0080\u008b\u0001\u0000\u0000\u0000\u0081"+
		"\u0082\u00034\u001a\u0000\u0082\u0083\u0006\u0002\uffff\uffff\u0000\u0083"+
		"\u008b\u0001\u0000\u0000\u0000\u0084\u0085\u0003*\u0015\u0000\u0085\u0086"+
		"\u0006\u0002\uffff\uffff\u0000\u0086\u008b\u0001\u0000\u0000\u0000\u0087"+
		"\u0088\u0003:\u001d\u0000\u0088\u0089\u0006\u0002\uffff\uffff\u0000\u0089"+
		"\u008b\u0001\u0000\u0000\u0000\u008aQ\u0001\u0000\u0000\u0000\u008aT\u0001"+
		"\u0000\u0000\u0000\u008aW\u0001\u0000\u0000\u0000\u008aZ\u0001\u0000\u0000"+
		"\u0000\u008a]\u0001\u0000\u0000\u0000\u008a`\u0001\u0000\u0000\u0000\u008a"+
		"c\u0001\u0000\u0000\u0000\u008af\u0001\u0000\u0000\u0000\u008ai\u0001"+
		"\u0000\u0000\u0000\u008al\u0001\u0000\u0000\u0000\u008ao\u0001\u0000\u0000"+
		"\u0000\u008ar\u0001\u0000\u0000\u0000\u008au\u0001\u0000\u0000\u0000\u008a"+
		"x\u0001\u0000\u0000\u0000\u008a{\u0001\u0000\u0000\u0000\u008a~\u0001"+
		"\u0000\u0000\u0000\u008a\u0081\u0001\u0000\u0000\u0000\u008a\u0084\u0001"+
		"\u0000\u0000\u0000\u008a\u0087\u0001\u0000\u0000\u0000\u008b\u0005\u0001"+
		"\u0000\u0000\u0000\u008c\u008d\u0005%\u0000\u0000\u008d\u008e\u0005(\u0000"+
		"\u0000\u008e\u008f\u0003B!\u0000\u008f\u0090\u0006\u0003\uffff\uffff\u0000"+
		"\u0090\u0007\u0001\u0000\u0000\u0000\u0091\u0092\u0005%\u0000\u0000\u0092"+
		"\u0093\u0005)\u0000\u0000\u0093\u0094\u0003B!\u0000\u0094\u0095\u0006"+
		"\u0004\uffff\uffff\u0000\u0095\t\u0001\u0000\u0000\u0000\u0096\u0097\u0005"+
		"\u0010\u0000\u0000\u0097\u0098\u0006\u0005\uffff\uffff\u0000\u0098\u000b"+
		"\u0001\u0000\u0000\u0000\u0099\u009a\u0005\u0017\u0000\u0000\u009a\u009b"+
		"\u0003B!\u0000\u009b\u009c\u0006\u0006\uffff\uffff\u0000\u009c\u00a0\u0001"+
		"\u0000\u0000\u0000\u009d\u009e\u0005\u0017\u0000\u0000\u009e\u00a0\u0006"+
		"\u0006\uffff\uffff\u0000\u009f\u0099\u0001\u0000\u0000\u0000\u009f\u009d"+
		"\u0001\u0000\u0000\u0000\u00a0\r\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005"+
		"\u0016\u0000\u0000\u00a2\u00a3\u0006\u0007\uffff\uffff\u0000\u00a3\u000f"+
		"\u0001\u0000\u0000\u0000\u00a4\u00a5\u0005\n\u0000\u0000\u00a5\u00a6\u0005"+
		"%\u0000\u0000\u00a6\u00a7\u0005=\u0000\u0000\u00a7\u00a8\u0003D\"\u0000"+
		"\u00a8\u00a9\u0005:\u0000\u0000\u00a9\u00aa\u0003B!\u0000\u00aa\u00ab"+
		"\u0006\b\uffff\uffff\u0000\u00ab\u00b3\u0001\u0000\u0000\u0000\u00ac\u00ad"+
		"\u0005\n\u0000\u0000\u00ad\u00ae\u0005%\u0000\u0000\u00ae\u00af\u0005"+
		":\u0000\u0000\u00af\u00b0\u0003B!\u0000\u00b0\u00b1\u0006\b\uffff\uffff"+
		"\u0000\u00b1\u00b3\u0001\u0000\u0000\u0000\u00b2\u00a4\u0001\u0000\u0000"+
		"\u0000\u00b2\u00ac\u0001\u0000\u0000\u0000\u00b3\u0011\u0001\u0000\u0000"+
		"\u0000\u00b4\u00b5\u0005\u0003\u0000\u0000\u00b5\u00b6\u0005?\u0000\u0000"+
		"\u00b6\u00b7\u0003B!\u0000\u00b7\u00b8\u0005@\u0000\u0000\u00b8\u00b9"+
		"\u0006\t\uffff\uffff\u0000\u00b9\u00c7\u0001\u0000\u0000\u0000\u00ba\u00bb"+
		"\u0005\u0001\u0000\u0000\u00bb\u00bc\u0005?\u0000\u0000\u00bc\u00bd\u0003"+
		"B!\u0000\u00bd\u00be\u0005@\u0000\u0000\u00be\u00bf\u0006\t\uffff\uffff"+
		"\u0000\u00bf\u00c7\u0001\u0000\u0000\u0000\u00c0\u00c1\u0005\u0002\u0000"+
		"\u0000\u00c1\u00c2\u0005?\u0000\u0000\u00c2\u00c3\u0003B!\u0000\u00c3"+
		"\u00c4\u0005@\u0000\u0000\u00c4\u00c5\u0006\t\uffff\uffff\u0000\u00c5"+
		"\u00c7\u0001\u0000\u0000\u0000\u00c6\u00b4\u0001\u0000\u0000\u0000\u00c6"+
		"\u00ba\u0001\u0000\u0000\u0000\u00c6\u00c0\u0001\u0000\u0000\u0000\u00c7"+
		"\u0013\u0001\u0000\u0000\u0000\u00c8\u00c9\u0005\t\u0000\u0000\u00c9\u00ca"+
		"\u0005%\u0000\u0000\u00ca\u00cb\u0005=\u0000\u0000\u00cb\u00cc\u0003D"+
		"\"\u0000\u00cc\u00cd\u0005:\u0000\u0000\u00cd\u00ce\u0003B!\u0000\u00ce"+
		"\u00cf\u0006\n\uffff\uffff\u0000\u00cf\u00de\u0001\u0000\u0000\u0000\u00d0"+
		"\u00d1\u0005\t\u0000\u0000\u00d1\u00d2\u0005%\u0000\u0000\u00d2\u00d3"+
		"\u0005:\u0000\u0000\u00d3\u00d4\u0003B!\u0000\u00d4\u00d5\u0006\n\uffff"+
		"\uffff\u0000\u00d5\u00de\u0001\u0000\u0000\u0000\u00d6\u00d7\u0005\t\u0000"+
		"\u0000\u00d7\u00d8\u0005%\u0000\u0000\u00d8\u00d9\u0005=\u0000\u0000\u00d9"+
		"\u00da\u0003D\"\u0000\u00da\u00db\u00050\u0000\u0000\u00db\u00dc\u0006"+
		"\n\uffff\uffff\u0000\u00dc\u00de\u0001\u0000\u0000\u0000\u00dd\u00c8\u0001"+
		"\u0000\u0000\u0000\u00dd\u00d0\u0001\u0000\u0000\u0000\u00dd\u00d6\u0001"+
		"\u0000\u0000\u0000\u00de\u0015\u0001\u0000\u0000\u0000\u00df\u00e0\u0005"+
		"%\u0000\u0000\u00e0\u00e1\u0005:\u0000\u0000\u00e1\u00e2\u0003B!\u0000"+
		"\u00e2\u00e3\u0006\u000b\uffff\uffff\u0000\u00e3\u0017\u0001\u0000\u0000"+
		"\u0000\u00e4\u00e5\u0005\u000b\u0000\u0000\u00e5\u00e6\u0005?\u0000\u0000"+
		"\u00e6\u00e7\u0003\u001a\r\u0000\u00e7\u00e8\u0005@\u0000\u0000\u00e8"+
		"\u00e9\u0006\f\uffff\uffff\u0000\u00e9\u0019\u0001\u0000\u0000\u0000\u00ea"+
		"\u00eb\u0003B!\u0000\u00eb\u00ec\u0005<\u0000\u0000\u00ec\u00ed\u0003"+
		"\u001a\r\u0000\u00ed\u00ee\u0006\r\uffff\uffff\u0000\u00ee\u00f3\u0001"+
		"\u0000\u0000\u0000\u00ef\u00f0\u0003B!\u0000\u00f0\u00f1\u0006\r\uffff"+
		"\uffff\u0000\u00f1\u00f3\u0001\u0000\u0000\u0000\u00f2\u00ea\u0001\u0000"+
		"\u0000\u0000\u00f2\u00ef\u0001\u0000\u0000\u0000\u00f3\u001b\u0001\u0000"+
		"\u0000\u0000\u00f4\u00f5\u0005\f\u0000\u0000\u00f5\u00f6\u0003B!\u0000"+
		"\u00f6\u00f7\u0005A\u0000\u0000\u00f7\u00f8\u0003\u0002\u0001\u0000\u00f8"+
		"\u00f9\u0005B\u0000\u0000\u00f9\u00fa\u0006\u000e\uffff\uffff\u0000\u00fa"+
		"\u010f\u0001\u0000\u0000\u0000\u00fb\u00fc\u0005\f\u0000\u0000\u00fc\u00fd"+
		"\u0003B!\u0000\u00fd\u00fe\u0005A\u0000\u0000\u00fe\u00ff\u0003\u0002"+
		"\u0001\u0000\u00ff\u0100\u0005B\u0000\u0000\u0100\u0101\u0005\r\u0000"+
		"\u0000\u0101\u0102\u0005A\u0000\u0000\u0102\u0103\u0003\u0002\u0001\u0000"+
		"\u0103\u0104\u0005B\u0000\u0000\u0104\u0105\u0006\u000e\uffff\uffff\u0000"+
		"\u0105\u010f\u0001\u0000\u0000\u0000\u0106\u0107\u0005\f\u0000\u0000\u0107"+
		"\u0108\u0003B!\u0000\u0108\u0109\u0005A\u0000\u0000\u0109\u010a\u0003"+
		"\u0002\u0001\u0000\u010a\u010b\u0005B\u0000\u0000\u010b\u010c\u0003\u001e"+
		"\u000f\u0000\u010c\u010d\u0006\u000e\uffff\uffff\u0000\u010d\u010f\u0001"+
		"\u0000\u0000\u0000\u010e\u00f4\u0001\u0000\u0000\u0000\u010e\u00fb\u0001"+
		"\u0000\u0000\u0000\u010e\u0106\u0001\u0000\u0000\u0000\u010f\u001d\u0001"+
		"\u0000\u0000\u0000\u0110\u0111\u0005\r\u0000\u0000\u0111\u0112\u0005\f"+
		"\u0000\u0000\u0112\u0113\u0003B!\u0000\u0113\u0114\u0005A\u0000\u0000"+
		"\u0114\u0115\u0003\u0002\u0001\u0000\u0115\u0116\u0005B\u0000\u0000\u0116"+
		"\u0117\u0006\u000f\uffff\uffff\u0000\u0117\u012e\u0001\u0000\u0000\u0000"+
		"\u0118\u0119\u0005\r\u0000\u0000\u0119\u011a\u0005\f\u0000\u0000\u011a"+
		"\u011b\u0003B!\u0000\u011b\u011c\u0005A\u0000\u0000\u011c\u011d\u0003"+
		"\u0002\u0001\u0000\u011d\u011e\u0005B\u0000\u0000\u011e\u011f\u0005\r"+
		"\u0000\u0000\u011f\u0120\u0005A\u0000\u0000\u0120\u0121\u0003\u0002\u0001"+
		"\u0000\u0121\u0122\u0005B\u0000\u0000\u0122\u0123\u0006\u000f\uffff\uffff"+
		"\u0000\u0123\u012e\u0001\u0000\u0000\u0000\u0124\u0125\u0005\r\u0000\u0000"+
		"\u0125\u0126\u0005\f\u0000\u0000\u0126\u0127\u0003B!\u0000\u0127\u0128"+
		"\u0005A\u0000\u0000\u0128\u0129\u0003\u0002\u0001\u0000\u0129\u012a\u0005"+
		"B\u0000\u0000\u012a\u012b\u0003\u001e\u000f\u0000\u012b\u012c\u0006\u000f"+
		"\uffff\uffff\u0000\u012c\u012e\u0001\u0000\u0000\u0000\u012d\u0110\u0001"+
		"\u0000\u0000\u0000\u012d\u0118\u0001\u0000\u0000\u0000\u012d\u0124\u0001"+
		"\u0000\u0000\u0000\u012e\u001f\u0001\u0000\u0000\u0000\u012f\u0130\u0005"+
		"\u0012\u0000\u0000\u0130\u0131\u0003B!\u0000\u0131\u0132\u0005A\u0000"+
		"\u0000\u0132\u0133\u0003\u0002\u0001\u0000\u0133\u0134\u0005B\u0000\u0000"+
		"\u0134\u0135\u0006\u0010\uffff\uffff\u0000\u0135!\u0001\u0000\u0000\u0000"+
		"\u0136\u0137\u0005\u0013\u0000\u0000\u0137\u0138\u0005%\u0000\u0000\u0138"+
		"\u0139\u0005\u0014\u0000\u0000\u0139\u013a\u0003B!\u0000\u013a\u013b\u0005"+
		"*\u0000\u0000\u013b\u013c\u0003B!\u0000\u013c\u013d\u0005A\u0000\u0000"+
		"\u013d\u013e\u0003\u0002\u0001\u0000\u013e\u013f\u0005B\u0000\u0000\u013f"+
		"\u0140\u0006\u0011\uffff\uffff\u0000\u0140#\u0001\u0000\u0000\u0000\u0141"+
		"\u0142\u0005\u0015\u0000\u0000\u0142\u0143\u0003B!\u0000\u0143\u0144\u0005"+
		"\r\u0000\u0000\u0144\u0145\u0005A\u0000\u0000\u0145\u0146\u0003\u0002"+
		"\u0001\u0000\u0146\u0147\u0005B\u0000\u0000\u0147\u0148\u0006\u0012\uffff"+
		"\uffff\u0000\u0148%\u0001\u0000\u0000\u0000\u0149\u014a\u0005\t\u0000"+
		"\u0000\u014a\u014b\u0005%\u0000\u0000\u014b\u014c\u0005=\u0000\u0000\u014c"+
		"\u014d\u0005C\u0000\u0000\u014d\u014e\u0003D\"\u0000\u014e\u014f\u0005"+
		"D\u0000\u0000\u014f\u0150\u0005:\u0000\u0000\u0150\u0151\u0005C\u0000"+
		"\u0000\u0151\u0152\u0003(\u0014\u0000\u0152\u0153\u0005D\u0000\u0000\u0153"+
		"\u0154\u0006\u0013\uffff\uffff\u0000\u0154\u0161\u0001\u0000\u0000\u0000"+
		"\u0155\u0156\u0005\t\u0000\u0000\u0156\u0157\u0005%\u0000\u0000\u0157"+
		"\u0158\u0005=\u0000\u0000\u0158\u0159\u0005C\u0000\u0000\u0159\u015a\u0003"+
		"D\"\u0000\u015a\u015b\u0005D\u0000\u0000\u015b\u015c\u0005:\u0000\u0000"+
		"\u015c\u015d\u0005C\u0000\u0000\u015d\u015e\u0005D\u0000\u0000\u015e\u015f"+
		"\u0006\u0013\uffff\uffff\u0000\u015f\u0161\u0001\u0000\u0000\u0000\u0160"+
		"\u0149\u0001\u0000\u0000\u0000\u0160\u0155\u0001\u0000\u0000\u0000\u0161"+
		"\'\u0001\u0000\u0000\u0000\u0162\u0163\u0003B!\u0000\u0163\u0164\u0005"+
		"<\u0000\u0000\u0164\u0165\u0003(\u0014\u0000\u0165\u0166\u0006\u0014\uffff"+
		"\uffff\u0000\u0166\u016b\u0001\u0000\u0000\u0000\u0167\u0168\u0003B!\u0000"+
		"\u0168\u0169\u0006\u0014\uffff\uffff\u0000\u0169\u016b\u0001\u0000\u0000"+
		"\u0000\u016a\u0162\u0001\u0000\u0000\u0000\u016a\u0167\u0001\u0000\u0000"+
		"\u0000\u016b)\u0001\u0000\u0000\u0000\u016c\u016d\u0005%\u0000\u0000\u016d"+
		"\u016e\u0005;\u0000\u0000\u016e\u016f\u0005\u001d\u0000\u0000\u016f\u0170"+
		"\u0005?\u0000\u0000\u0170\u0171\u0003B!\u0000\u0171\u0172\u0005@\u0000"+
		"\u0000\u0172\u0173\u0006\u0015\uffff\uffff\u0000\u0173\u0185\u0001\u0000"+
		"\u0000\u0000\u0174\u0175\u0005%\u0000\u0000\u0175\u0176\u0005;\u0000\u0000"+
		"\u0176\u0177\u0005\u001e\u0000\u0000\u0177\u0178\u0005?\u0000\u0000\u0178"+
		"\u0179\u0005@\u0000\u0000\u0179\u0185\u0006\u0015\uffff\uffff\u0000\u017a"+
		"\u017b\u0005%\u0000\u0000\u017b\u017c\u0005;\u0000\u0000\u017c\u017d\u0005"+
		"\u001f\u0000\u0000\u017d\u017e\u0005?\u0000\u0000\u017e\u017f\u0005 \u0000"+
		"\u0000\u017f\u0180\u0005=\u0000\u0000\u0180\u0181\u0003B!\u0000\u0181"+
		"\u0182\u0005@\u0000\u0000\u0182\u0183\u0006\u0015\uffff\uffff\u0000\u0183"+
		"\u0185\u0001\u0000\u0000\u0000\u0184\u016c\u0001\u0000\u0000\u0000\u0184"+
		"\u0174\u0001\u0000\u0000\u0000\u0184\u017a\u0001\u0000\u0000\u0000\u0185"+
		"+\u0001\u0000\u0000\u0000\u0186\u0187\u0005%\u0000\u0000\u0187\u0188\u0005"+
		"C\u0000\u0000\u0188\u0189\u0003B!\u0000\u0189\u018a\u0005D\u0000\u0000"+
		"\u018a\u018b\u0005:\u0000\u0000\u018b\u018c\u0003B!\u0000\u018c\u018d"+
		"\u0006\u0016\uffff\uffff\u0000\u018d-\u0001\u0000\u0000\u0000\u018e\u018f"+
		"\u0005\u0018\u0000\u0000\u018f\u0190\u0005%\u0000\u0000\u0190\u0191\u0005"+
		"?\u0000\u0000\u0191\u0192\u0005@\u0000\u0000\u0192\u0193\u0005E\u0000"+
		"\u0000\u0193\u0194\u0003D\"\u0000\u0194\u0195\u0005A\u0000\u0000\u0195"+
		"\u0196\u0003\u0002\u0001\u0000\u0196\u0197\u0005B\u0000\u0000\u0197\u0198"+
		"\u0006\u0017\uffff\uffff\u0000\u0198\u01b9\u0001\u0000\u0000\u0000\u0199"+
		"\u019a\u0005\u0018\u0000\u0000\u019a\u019b\u0005%\u0000\u0000\u019b\u019c"+
		"\u0005?\u0000\u0000\u019c\u019d\u0005@\u0000\u0000\u019d\u019e\u0005A"+
		"\u0000\u0000\u019e\u019f\u0003\u0002\u0001\u0000\u019f\u01a0\u0005B\u0000"+
		"\u0000\u01a0\u01a1\u0006\u0017\uffff\uffff\u0000\u01a1\u01b9\u0001\u0000"+
		"\u0000\u0000\u01a2\u01a3\u0005\u0018\u0000\u0000\u01a3\u01a4\u0005%\u0000"+
		"\u0000\u01a4\u01a5\u0005?\u0000\u0000\u01a5\u01a6\u00030\u0018\u0000\u01a6"+
		"\u01a7\u0005@\u0000\u0000\u01a7\u01a8\u0005E\u0000\u0000\u01a8\u01a9\u0003"+
		"D\"\u0000\u01a9\u01aa\u0005A\u0000\u0000\u01aa\u01ab\u0003\u0002\u0001"+
		"\u0000\u01ab\u01ac\u0005B\u0000\u0000\u01ac\u01ad\u0006\u0017\uffff\uffff"+
		"\u0000\u01ad\u01b9\u0001\u0000\u0000\u0000\u01ae\u01af\u0005\u0018\u0000"+
		"\u0000\u01af\u01b0\u0005%\u0000\u0000\u01b0\u01b1\u0005?\u0000\u0000\u01b1"+
		"\u01b2\u00030\u0018\u0000\u01b2\u01b3\u0005@\u0000\u0000\u01b3\u01b4\u0005"+
		"A\u0000\u0000\u01b4\u01b5\u0003\u0002\u0001\u0000\u01b5\u01b6\u0005B\u0000"+
		"\u0000\u01b6\u01b7\u0006\u0017\uffff\uffff\u0000\u01b7\u01b9\u0001\u0000"+
		"\u0000\u0000\u01b8\u018e\u0001\u0000\u0000\u0000\u01b8\u0199\u0001\u0000"+
		"\u0000\u0000\u01b8\u01a2\u0001\u0000\u0000\u0000\u01b8\u01ae\u0001\u0000"+
		"\u0000\u0000\u01b9/\u0001\u0000\u0000\u0000\u01ba\u01bb\u0005%\u0000\u0000"+
		"\u01bb\u01bc\u0005=\u0000\u0000\u01bc\u01bd\u0003D\"\u0000\u01bd\u01be"+
		"\u0005<\u0000\u0000\u01be\u01bf\u00030\u0018\u0000\u01bf\u01c0\u0006\u0018"+
		"\uffff\uffff\u0000\u01c0\u01f7\u0001\u0000\u0000\u0000\u01c1\u01c2\u0005"+
		"%\u0000\u0000\u01c2\u01c3\u0005=\u0000\u0000\u01c3\u01c4\u0003D\"\u0000"+
		"\u01c4\u01c5\u0006\u0018\uffff\uffff\u0000\u01c5\u01f7\u0001\u0000\u0000"+
		"\u0000\u01c6\u01c7\u00032\u0019\u0000\u01c7\u01c8\u0005%\u0000\u0000\u01c8"+
		"\u01c9\u0005=\u0000\u0000\u01c9\u01ca\u0003D\"\u0000\u01ca\u01cb\u0005"+
		"<\u0000\u0000\u01cb\u01cc\u00030\u0018\u0000\u01cc\u01cd\u0006\u0018\uffff"+
		"\uffff\u0000\u01cd\u01f7\u0001\u0000\u0000\u0000\u01ce\u01cf\u00032\u0019"+
		"\u0000\u01cf\u01d0\u0005%\u0000\u0000\u01d0\u01d1\u0005=\u0000\u0000\u01d1"+
		"\u01d2\u0003D\"\u0000\u01d2\u01d3\u0006\u0018\uffff\uffff\u0000\u01d3"+
		"\u01f7\u0001\u0000\u0000\u0000\u01d4\u01d5\u0005%\u0000\u0000\u01d5\u01d6"+
		"\u0005=\u0000\u0000\u01d6\u01d7\u0005C\u0000\u0000\u01d7\u01d8\u0003D"+
		"\"\u0000\u01d8\u01d9\u0005D\u0000\u0000\u01d9\u01da\u0005<\u0000\u0000"+
		"\u01da\u01db\u00030\u0018\u0000\u01db\u01dc\u0006\u0018\uffff\uffff\u0000"+
		"\u01dc\u01f7\u0001\u0000\u0000\u0000\u01dd\u01de\u0005%\u0000\u0000\u01de"+
		"\u01df\u0005=\u0000\u0000\u01df\u01e0\u0005C\u0000\u0000\u01e0\u01e1\u0003"+
		"D\"\u0000\u01e1\u01e2\u0005D\u0000\u0000\u01e2\u01e3\u0006\u0018\uffff"+
		"\uffff\u0000\u01e3\u01f7\u0001\u0000\u0000\u0000\u01e4\u01e5\u00032\u0019"+
		"\u0000\u01e5\u01e6\u0005%\u0000\u0000\u01e6\u01e7\u0005=\u0000\u0000\u01e7"+
		"\u01e8\u0005C\u0000\u0000\u01e8\u01e9\u0003D\"\u0000\u01e9\u01ea\u0005"+
		"D\u0000\u0000\u01ea\u01eb\u0005<\u0000\u0000\u01eb\u01ec\u00030\u0018"+
		"\u0000\u01ec\u01ed\u0006\u0018\uffff\uffff\u0000\u01ed\u01f7\u0001\u0000"+
		"\u0000\u0000\u01ee\u01ef\u00032\u0019\u0000\u01ef\u01f0\u0005%\u0000\u0000"+
		"\u01f0\u01f1\u0005=\u0000\u0000\u01f1\u01f2\u0005C\u0000\u0000\u01f2\u01f3"+
		"\u0003D\"\u0000\u01f3\u01f4\u0005D\u0000\u0000\u01f4\u01f5\u0006\u0018"+
		"\uffff\uffff\u0000\u01f5\u01f7\u0001\u0000\u0000\u0000\u01f6\u01ba\u0001"+
		"\u0000\u0000\u0000\u01f6\u01c1\u0001\u0000\u0000\u0000\u01f6\u01c6\u0001"+
		"\u0000\u0000\u0000\u01f6\u01ce\u0001\u0000\u0000\u0000\u01f6\u01d4\u0001"+
		"\u0000\u0000\u0000\u01f6\u01dd\u0001\u0000\u0000\u0000\u01f6\u01e4\u0001"+
		"\u0000\u0000\u0000\u01f6\u01ee\u0001\u0000\u0000\u0000\u01f71\u0001\u0000"+
		"\u0000\u0000\u01f8\u01f9\u0005%\u0000\u0000\u01f9\u01fd\u0006\u0019\uffff"+
		"\uffff\u0000\u01fa\u01fb\u0005F\u0000\u0000\u01fb\u01fd\u0006\u0019\uffff"+
		"\uffff\u0000\u01fc\u01f8\u0001\u0000\u0000\u0000\u01fc\u01fa\u0001\u0000"+
		"\u0000\u0000\u01fd3\u0001\u0000\u0000\u0000\u01fe\u01ff\u0005%\u0000\u0000"+
		"\u01ff\u0200\u0005?\u0000\u0000\u0200\u0201\u0005@\u0000\u0000\u0201\u0209"+
		"\u0006\u001a\uffff\uffff\u0000\u0202\u0203\u0005%\u0000\u0000\u0203\u0204"+
		"\u0005?\u0000\u0000\u0204\u0205\u00036\u001b\u0000\u0205\u0206\u0005@"+
		"\u0000\u0000\u0206\u0207\u0006\u001a\uffff\uffff\u0000\u0207\u0209\u0001"+
		"\u0000\u0000\u0000\u0208\u01fe\u0001\u0000\u0000\u0000\u0208\u0202\u0001"+
		"\u0000\u0000\u0000\u02095\u0001\u0000\u0000\u0000\u020a\u020b\u0003B!"+
		"\u0000\u020b\u020c\u0005<\u0000\u0000\u020c\u020d\u00036\u001b\u0000\u020d"+
		"\u020e\u0006\u001b\uffff\uffff\u0000\u020e\u021f\u0001\u0000\u0000\u0000"+
		"\u020f\u0210\u0003B!\u0000\u0210\u0211\u0006\u001b\uffff\uffff\u0000\u0211"+
		"\u021f\u0001\u0000\u0000\u0000\u0212\u0213\u0005%\u0000\u0000\u0213\u0214"+
		"\u0005=\u0000\u0000\u0214\u0215\u0003B!\u0000\u0215\u0216\u0005<\u0000"+
		"\u0000\u0216\u0217\u00036\u001b\u0000\u0217\u0218\u0006\u001b\uffff\uffff"+
		"\u0000\u0218\u021f\u0001\u0000\u0000\u0000\u0219\u021a\u0005%\u0000\u0000"+
		"\u021a\u021b\u0005=\u0000\u0000\u021b\u021c\u0003B!\u0000\u021c\u021d"+
		"\u0006\u001b\uffff\uffff\u0000\u021d\u021f\u0001\u0000\u0000\u0000\u021e"+
		"\u020a\u0001\u0000\u0000\u0000\u021e\u020f\u0001\u0000\u0000\u0000\u021e"+
		"\u0212\u0001\u0000\u0000\u0000\u021e\u0219\u0001\u0000\u0000\u0000\u021f"+
		"7\u0001\u0000\u0000\u0000\u0220\u0221\u00034\u001a\u0000\u0221\u0222\u0006"+
		"\u001c\uffff\uffff\u0000\u02229\u0001\u0000\u0000\u0000\u0223\u0224\u0005"+
		"\t\u0000\u0000\u0224\u0225\u0005%\u0000\u0000\u0225\u0226\u0005=\u0000"+
		"\u0000\u0226\u0227\u0003>\u001f\u0000\u0227\u0228\u0005:\u0000\u0000\u0228"+
		"\u0229\u0005C\u0000\u0000\u0229\u022a\u0003<\u001e\u0000\u022a\u022b\u0005"+
		"D\u0000\u0000\u022b\u022c\u0006\u001d\uffff\uffff\u0000\u022c;\u0001\u0000"+
		"\u0000\u0000\u022d\u022e\u0005C\u0000\u0000\u022e\u022f\u0003B!\u0000"+
		"\u022f\u0230\u0005D\u0000\u0000\u0230\u0231\u0006\u001e\uffff\uffff\u0000"+
		"\u0231\u023a\u0001\u0000\u0000\u0000\u0232\u0233\u0005C\u0000\u0000\u0233"+
		"\u0234\u0003B!\u0000\u0234\u0235\u0005D\u0000\u0000\u0235\u0236\u0005"+
		"<\u0000\u0000\u0236\u0237\u0003<\u001e\u0000\u0237\u0238\u0006\u001e\uffff"+
		"\uffff\u0000\u0238\u023a\u0001\u0000\u0000\u0000\u0239\u022d\u0001\u0000"+
		"\u0000\u0000\u0239\u0232\u0001\u0000\u0000\u0000\u023a=\u0001\u0000\u0000"+
		"\u0000\u023b\u023c\u0005C\u0000\u0000\u023c\u023d\u0003>\u001f\u0000\u023d"+
		"\u023e\u0005D\u0000\u0000\u023e\u023f\u0006\u001f\uffff\uffff\u0000\u023f"+
		"\u0246\u0001\u0000\u0000\u0000\u0240\u0241\u0005C\u0000\u0000\u0241\u0242"+
		"\u0003D\"\u0000\u0242\u0243\u0005D\u0000\u0000\u0243\u0244\u0006\u001f"+
		"\uffff\uffff\u0000\u0244\u0246\u0001\u0000\u0000\u0000\u0245\u023b\u0001"+
		"\u0000\u0000\u0000\u0245\u0240\u0001\u0000\u0000\u0000\u0246?\u0001\u0000"+
		"\u0000\u0000\u0247\u0248\u0003>\u001f\u0000\u0248A\u0001\u0000\u0000\u0000"+
		"\u0249\u024a\u0006!\uffff\uffff\u0000\u024a\u024b\u00053\u0000\u0000\u024b"+
		"\u024c\u0003B!\u0011\u024c\u024d\u0006!\uffff\uffff\u0000\u024d\u0281"+
		"\u0001\u0000\u0000\u0000\u024e\u024f\u0005,\u0000\u0000\u024f\u0250\u0003"+
		"B!\u0010\u0250\u0251\u0006!\uffff\uffff\u0000\u0251\u0281\u0001\u0000"+
		"\u0000\u0000\u0252\u0253\u0005?\u0000\u0000\u0253\u0254\u0003B!\u0000"+
		"\u0254\u0255\u0005@\u0000\u0000\u0255\u0256\u0006!\uffff\uffff\u0000\u0256"+
		"\u0281\u0001\u0000\u0000\u0000\u0257\u0258\u00038\u001c\u0000\u0258\u0259"+
		"\u0006!\uffff\uffff\u0000\u0259\u0281\u0001\u0000\u0000\u0000\u025a\u025b"+
		"\u0003\u0012\t\u0000\u025b\u025c\u0006!\uffff\uffff\u0000\u025c\u0281"+
		"\u0001\u0000\u0000\u0000\u025d\u025e\u0005%\u0000\u0000\u025e\u025f\u0005"+
		"C\u0000\u0000\u025f\u0260\u0003B!\u0000\u0260\u0261\u0005D\u0000\u0000"+
		"\u0261\u0262\u0006!\uffff\uffff\u0000\u0262\u0281\u0001\u0000\u0000\u0000"+
		"\u0263\u0264\u0005%\u0000\u0000\u0264\u0265\u0005;\u0000\u0000\u0265\u0266"+
		"\u0005\"\u0000\u0000\u0266\u0281\u0006!\uffff\uffff\u0000\u0267\u0268"+
		"\u0005%\u0000\u0000\u0268\u0269\u0005;\u0000\u0000\u0269\u026a\u0005!"+
		"\u0000\u0000\u026a\u0281\u0006!\uffff\uffff\u0000\u026b\u026c\u0005C\u0000"+
		"\u0000\u026c\u026d\u0003(\u0014\u0000\u026d\u026e\u0005D\u0000\u0000\u026e"+
		"\u026f\u0006!\uffff\uffff\u0000\u026f\u0281\u0001\u0000\u0000\u0000\u0270"+
		"\u0271\u0005#\u0000\u0000\u0271\u0281\u0006!\uffff\uffff\u0000\u0272\u0273"+
		"\u0005$\u0000\u0000\u0273\u0281\u0006!\uffff\uffff\u0000\u0274\u0275\u0005"+
		"\'\u0000\u0000\u0275\u0281\u0006!\uffff\uffff\u0000\u0276\u0277\u0005"+
		"&\u0000\u0000\u0277\u0281\u0006!\uffff\uffff\u0000\u0278\u0279\u0005\u0006"+
		"\u0000\u0000\u0279\u0281\u0006!\uffff\uffff\u0000\u027a\u027b\u0005\u0007"+
		"\u0000\u0000\u027b\u0281\u0006!\uffff\uffff\u0000\u027c\u027d\u0005%\u0000"+
		"\u0000\u027d\u0281\u0006!\uffff\uffff\u0000\u027e\u027f\u0005\b\u0000"+
		"\u0000\u027f\u0281\u0006!\uffff\uffff\u0000\u0280\u0249\u0001\u0000\u0000"+
		"\u0000\u0280\u024e\u0001\u0000\u0000\u0000\u0280\u0252\u0001\u0000\u0000"+
		"\u0000\u0280\u0257\u0001\u0000\u0000\u0000\u0280\u025a\u0001\u0000\u0000"+
		"\u0000\u0280\u025d\u0001\u0000\u0000\u0000\u0280\u0263\u0001\u0000\u0000"+
		"\u0000\u0280\u0267\u0001\u0000\u0000\u0000\u0280\u026b\u0001\u0000\u0000"+
		"\u0000\u0280\u0270\u0001\u0000\u0000\u0000\u0280\u0272\u0001\u0000\u0000"+
		"\u0000\u0280\u0274\u0001\u0000\u0000\u0000\u0280\u0276\u0001\u0000\u0000"+
		"\u0000\u0280\u0278\u0001\u0000\u0000\u0000\u0280\u027a\u0001\u0000\u0000"+
		"\u0000\u0280\u027c\u0001\u0000\u0000\u0000\u0280\u027e\u0001\u0000\u0000"+
		"\u0000\u0281\u02a7\u0001\u0000\u0000\u0000\u0282\u0283\n\u0018\u0000\u0000"+
		"\u0283\u0284\u0007\u0000\u0000\u0000\u0284\u0285\u0003B!\u0019\u0285\u0286"+
		"\u0006!\uffff\uffff\u0000\u0286\u02a6\u0001\u0000\u0000\u0000\u0287\u0288"+
		"\n\u0017\u0000\u0000\u0288\u0289\u0007\u0001\u0000\u0000\u0289\u028a\u0003"+
		"B!\u0018\u028a\u028b\u0006!\uffff\uffff\u0000\u028b\u02a6\u0001\u0000"+
		"\u0000\u0000\u028c\u028d\n\u0016\u0000\u0000\u028d\u028e\u0005/\u0000"+
		"\u0000\u028e\u028f\u0003B!\u0017\u028f\u0290\u0006!\uffff\uffff\u0000"+
		"\u0290\u02a6\u0001\u0000\u0000\u0000\u0291\u0292\n\u0015\u0000\u0000\u0292"+
		"\u0293\u0007\u0002\u0000\u0000\u0293\u0294\u0003B!\u0016\u0294\u0295\u0006"+
		"!\uffff\uffff\u0000\u0295\u02a6\u0001\u0000\u0000\u0000\u0296\u0297\n"+
		"\u0014\u0000\u0000\u0297\u0298\u0007\u0003\u0000\u0000\u0298\u0299\u0003"+
		"B!\u0015\u0299\u029a\u0006!\uffff\uffff\u0000\u029a\u02a6\u0001\u0000"+
		"\u0000\u0000\u029b\u029c\n\u0013\u0000\u0000\u029c\u029d\u0007\u0004\u0000"+
		"\u0000\u029d\u029e\u0003B!\u0014\u029e\u029f\u0006!\uffff\uffff\u0000"+
		"\u029f\u02a6\u0001\u0000\u0000\u0000\u02a0\u02a1\n\u0012\u0000\u0000\u02a1"+
		"\u02a2\u0007\u0005\u0000\u0000\u02a2\u02a3\u0003B!\u0013\u02a3\u02a4\u0006"+
		"!\uffff\uffff\u0000\u02a4\u02a6\u0001\u0000\u0000\u0000\u02a5\u0282\u0001"+
		"\u0000\u0000\u0000\u02a5\u0287\u0001\u0000\u0000\u0000\u02a5\u028c\u0001"+
		"\u0000\u0000\u0000\u02a5\u0291\u0001\u0000\u0000\u0000\u02a5\u0296\u0001"+
		"\u0000\u0000\u0000\u02a5\u029b\u0001\u0000\u0000\u0000\u02a5\u02a0\u0001"+
		"\u0000\u0000\u0000\u02a6\u02a9\u0001\u0000\u0000\u0000\u02a7\u02a5\u0001"+
		"\u0000\u0000\u0000\u02a7\u02a8\u0001\u0000\u0000\u0000\u02a8C\u0001\u0000"+
		"\u0000\u0000\u02a9\u02a7\u0001\u0000\u0000\u0000\u02aa\u02ab\u0005\u0001"+
		"\u0000\u0000\u02ab\u02b5\u0006\"\uffff\uffff\u0000\u02ac\u02ad\u0005\u0002"+
		"\u0000\u0000\u02ad\u02b5\u0006\"\uffff\uffff\u0000\u02ae\u02af\u0005\u0003"+
		"\u0000\u0000\u02af\u02b5\u0006\"\uffff\uffff\u0000\u02b0\u02b1\u0005\u0004"+
		"\u0000\u0000\u02b1\u02b5\u0006\"\uffff\uffff\u0000\u02b2\u02b3\u0005\u0005"+
		"\u0000\u0000\u02b3\u02b5\u0006\"\uffff\uffff\u0000\u02b4\u02aa\u0001\u0000"+
		"\u0000\u0000\u02b4\u02ac\u0001\u0000\u0000\u0000\u02b4\u02ae\u0001\u0000"+
		"\u0000\u0000\u02b4\u02b0\u0001\u0000\u0000\u0000\u02b4\u02b2\u0001\u0000"+
		"\u0000\u0000\u02b5E\u0001\u0000\u0000\u0000\u0017M\u008a\u009f\u00b2\u00c6"+
		"\u00dd\u00f2\u010e\u012d\u0160\u016a\u0184\u01b8\u01f6\u01fc\u0208\u021e"+
		"\u0239\u0245\u0280\u02a5\u02a7\u02b4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}