// Generated from e:\USAC2023\SegundoSemestre\Compiladores2\OLC2_P12S201900822\backend\grammar\Swiftgramm.g4 by ANTLR 4.9.2

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

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftgrammParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, STRING=3, BOOL=4, CHARACTER=5, TRUE=6, FALSE=7, NIL=8, 
		VAR=9, LET=10, PRINT=11, IF=12, ELSE=13, SWITCH=14, CASE=15, BREAK=16, 
		DEFAULT=17, WHILE=18, FOR=19, IN=20, GUARD=21, CONTINUE=22, RETURN=23, 
		FUNC=24, STRUCT=25, MUTATING=26, SELF=27, INOUT=28, NUMBER=29, FLOATT=30, 
		ID=31, CHARACTER_LITERAL=32, STRING_LITERAL=33, INCREMENT=34, DECREMENT=35, 
		RANGE=36, SUMMATION=37, SUBTRACTION=38, MULTIPLICATION=39, DIVISION=40, 
		MOD=41, QUESTION_MARK=42, OR=43, AND=44, NOT=45, EQUAL=46, NOT_EQUAL=47, 
		LESS_THAN=48, LESS_THAN_EQUAL=49, GREATER_THAN=50, GREATER_THAN_EQUAL=51, 
		ASSIGN=52, DOT=53, COMMA=54, COLON=55, SEMICOLON=56, OPEN_PARENTHESIS=57, 
		CLOSE_PARENTHESIS=58, OPEN_kEY=59, CLOSE_kEY=60, OPEN_BRACKET=61, CLOSE_BRACKET=62, 
		ARROW=63, WHITESPACE=64, COMMENT=65, LINE_COMMENT=66;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_sentence = 2, RULE_increment_bl = 3, 
		RULE_decrement_bl = 4, RULE_break_bl = 5, RULE_return_bl = 6, RULE_continue_bl = 7, 
		RULE_declare_let = 8, RULE_declare_var = 9, RULE_print_bl = 10, RULE_if_bl = 11, 
		RULE_else_if = 12, RULE_while_bl = 13, RULE_for_bl = 14, RULE_guard_bl = 15, 
		RULE_vector_bl = 16, RULE_array_exp = 17, RULE_function_bl = 18, RULE_params = 19, 
		RULE_call_function = 20, RULE_list_exp = 21, RULE_call_function_exp = 22, 
		RULE_expression = 23, RULE_datatype = 24;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "sentence", "increment_bl", "decrement_bl", "break_bl", 
			"return_bl", "continue_bl", "declare_let", "declare_var", "print_bl", 
			"if_bl", "else_if", "while_bl", "for_bl", "guard_bl", "vector_bl", "array_exp", 
			"function_bl", "params", "call_function", "list_exp", "call_function_exp", 
			"expression", "datatype"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'break'", "'default'", "'while'", "'for'", "'in'", "'guard'", 
			"'continue'", "'return'", "'func'", "'struct'", "'mutating'", "'self'", 
			"'inout'", null, null, null, null, null, "'+='", "'-='", "'...'", "'+'", 
			"'-'", "'*'", "'/'", "'%'", "'?'", "'||'", "'&&'", "'!'", "'=='", "'!='", 
			"'<'", "'<='", "'>'", "'>='", "'='", "'.'", "','", "':'", "';'", "'('", 
			"')'", "'{'", "'}'", "'['", "']'", "'->'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE", 
			"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK", 
			"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC", 
			"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "FLOATT", "ID", "CHARACTER_LITERAL", 
			"STRING_LITERAL", "INCREMENT", "DECREMENT", "RANGE", "SUMMATION", "SUBTRACTION", 
			"MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT", 
			"EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN", 
			"GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON", 
			"OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", "OPEN_BRACKET", 
			"CLOSE_BRACKET", "ARROW", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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

	public SwiftgrammParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

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
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(50);
			((SContext)_localctx).block = block();
			setState(51);
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
			setState(55); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(54);
				((BlockContext)_localctx).sentence = sentence();
				((BlockContext)_localctx).instr.add(((BlockContext)_localctx).sentence);
				}
				}
				setState(57); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << PRINT) | (1L << IF) | (1L << BREAK) | (1L << WHILE) | (1L << FOR) | (1L << GUARD) | (1L << CONTINUE) | (1L << RETURN) | (1L << FUNC) | (1L << ID))) != 0) );

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

	public static class SentenceContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Declare_varContext declare_var;
		public Declare_letContext declare_let;
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
		public Declare_varContext declare_var() {
			return getRuleContext(Declare_varContext.class,0);
		}
		public Declare_letContext declare_let() {
			return getRuleContext(Declare_letContext.class,0);
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
		public SentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentence; }
	}

	public final SentenceContext sentence() throws RecognitionException {
		SentenceContext _localctx = new SentenceContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentence);
		try {
			setState(106);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(61);
				((SentenceContext)_localctx).declare_var = declare_var();
				_localctx.instr = ((SentenceContext)_localctx).declare_var.instr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(64);
				((SentenceContext)_localctx).declare_let = declare_let();
				_localctx.instr = ((SentenceContext)_localctx).declare_let.instr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(67);
				((SentenceContext)_localctx).print_bl = print_bl();
				_localctx.instr = ((SentenceContext)_localctx).print_bl.instr
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(70);
				((SentenceContext)_localctx).if_bl = if_bl();
				_localctx.instr = ((SentenceContext)_localctx).if_bl.instr
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(73);
				((SentenceContext)_localctx).increment_bl = increment_bl();
				_localctx.instr = ((SentenceContext)_localctx).increment_bl.instr
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(76);
				((SentenceContext)_localctx).decrement_bl = decrement_bl();
				_localctx.instr = ((SentenceContext)_localctx).decrement_bl.instr
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(79);
				((SentenceContext)_localctx).while_bl = while_bl();
				_localctx.instr = ((SentenceContext)_localctx).while_bl.instr
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(82);
				((SentenceContext)_localctx).for_bl = for_bl();
				_localctx.instr = ((SentenceContext)_localctx).for_bl.instr
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(85);
				((SentenceContext)_localctx).guard_bl = guard_bl();
				_localctx.instr = ((SentenceContext)_localctx).guard_bl.instr
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(88);
				((SentenceContext)_localctx).break_bl = break_bl();
				_localctx.instr = ((SentenceContext)_localctx).break_bl.instr
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(91);
				((SentenceContext)_localctx).return_bl = return_bl();
				_localctx.instr = ((SentenceContext)_localctx).return_bl.instr
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(94);
				((SentenceContext)_localctx).continue_bl = continue_bl();
				_localctx.instr = ((SentenceContext)_localctx).continue_bl.instr
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(97);
				((SentenceContext)_localctx).vector_bl = vector_bl();
				_localctx.instr = ((SentenceContext)_localctx).vector_bl.instr
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(100);
				((SentenceContext)_localctx).function_bl = function_bl();
				_localctx.instr = ((SentenceContext)_localctx).function_bl.instr
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(103);
				((SentenceContext)_localctx).call_function = call_function();
				_localctx.instr = ((SentenceContext)_localctx).call_function.instr
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
	}

	public final Increment_blContext increment_bl() throws RecognitionException {
		Increment_blContext _localctx = new Increment_blContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_increment_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			((Increment_blContext)_localctx).ID = match(ID);
			setState(109);
			((Increment_blContext)_localctx).INCREMENT = match(INCREMENT);
			setState(110);
			((Increment_blContext)_localctx).expression = expression(0);

			        _localctx.instr = instructions.NewIncreDecrem((((Increment_blContext)_localctx).ID!=null?((Increment_blContext)_localctx).ID.getText():null),(((Increment_blContext)_localctx).INCREMENT!=null?((Increment_blContext)_localctx).INCREMENT.getText():null),((Increment_blContext)_localctx).expression.p)

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
	}

	public final Decrement_blContext decrement_bl() throws RecognitionException {
		Decrement_blContext _localctx = new Decrement_blContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_decrement_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(113);
			((Decrement_blContext)_localctx).ID = match(ID);
			setState(114);
			((Decrement_blContext)_localctx).DECREMENT = match(DECREMENT);
			setState(115);
			((Decrement_blContext)_localctx).expression = expression(0);

			        _localctx.instr = instructions.NewIncreDecrem((((Decrement_blContext)_localctx).ID!=null?((Decrement_blContext)_localctx).ID.getText():null),(((Decrement_blContext)_localctx).DECREMENT!=null?((Decrement_blContext)_localctx).DECREMENT.getText():null),((Decrement_blContext)_localctx).expression.p)

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

	public static class Break_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public TerminalNode BREAK() { return getToken(SwiftgrammParser.BREAK, 0); }
		public Break_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break_bl; }
	}

	public final Break_blContext break_bl() throws RecognitionException {
		Break_blContext _localctx = new Break_blContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_break_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(118);
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

	public static class Return_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public ExpressionContext expression;
		public TerminalNode RETURN() { return getToken(SwiftgrammParser.RETURN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Return_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_bl; }
	}

	public final Return_blContext return_bl() throws RecognitionException {
		Return_blContext _localctx = new Return_blContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_return_bl);
		try {
			setState(127);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(121);
				match(RETURN);
				setState(122);
				((Return_blContext)_localctx).expression = expression(0);

				        _localctx.instr = instructions.NewReturn(((Return_blContext)_localctx).expression.p)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(125);
				match(RETURN);

				        _localctx.instr = instructions.NewReturn(expressions.NewNative(nil,symbol.NIL))

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

	public static class Continue_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public TerminalNode CONTINUE() { return getToken(SwiftgrammParser.CONTINUE, 0); }
		public Continue_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue_bl; }
	}

	public final Continue_blContext continue_bl() throws RecognitionException {
		Continue_blContext _localctx = new Continue_blContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_continue_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
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
	}

	public final Declare_letContext declare_let() throws RecognitionException {
		Declare_letContext _localctx = new Declare_letContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_declare_let);
		try {
			setState(146);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(132);
				match(LET);
				setState(133);
				((Declare_letContext)_localctx).ID = match(ID);
				setState(134);
				match(COLON);
				setState(135);
				((Declare_letContext)_localctx).datatype = datatype();
				setState(136);
				match(ASSIGN);
				setState(137);
				((Declare_letContext)_localctx).expression = expression(0);

				                _localctx.instr = instructions.NewLet((((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getText():null),((Declare_letContext)_localctx).datatype.td,((Declare_letContext)_localctx).expression.p)
				        
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(140);
				match(LET);
				setState(141);
				((Declare_letContext)_localctx).ID = match(ID);
				setState(142);
				match(ASSIGN);
				setState(143);
				((Declare_letContext)_localctx).expression = expression(0);

				                _localctx.instr = instructions.NewLet((((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getText():null),symbol.UNDEFINED,((Declare_letContext)_localctx).expression.p)
				        
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

	public static class Declare_varContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public Token ID;
		public DatatypeContext datatype;
		public ExpressionContext expression;
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
	}

	public final Declare_varContext declare_var() throws RecognitionException {
		Declare_varContext _localctx = new Declare_varContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_declare_var);
		try {
			setState(169);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(148);
				match(VAR);
				setState(149);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(150);
				match(COLON);
				setState(151);
				((Declare_varContext)_localctx).datatype = datatype();
				setState(152);
				match(ASSIGN);
				setState(153);
				((Declare_varContext)_localctx).expression = expression(0);

				                        _localctx.instr = instructions.NewDeclareWithValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),((Declare_varContext)_localctx).datatype.td,((Declare_varContext)_localctx).expression.p)
				                
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(156);
				match(VAR);
				setState(157);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(158);
				match(ASSIGN);
				setState(159);
				((Declare_varContext)_localctx).expression = expression(0);

				                        _localctx.instr = instructions.NewDeclareWithValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),symbol.UNDEFINED,((Declare_varContext)_localctx).expression.p)
				                
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(162);
				match(VAR);
				setState(163);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(164);
				match(COLON);
				setState(165);
				((Declare_varContext)_localctx).datatype = datatype();
				setState(166);
				match(QUESTION_MARK);

				                        _localctx.instr = instructions.NewDeclareWithoutValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),((Declare_varContext)_localctx).datatype.td,expressions.NewNative(nil,symbol.NIL))
				                
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

	public static class Print_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
		public ExpressionContext expression;
		public TerminalNode PRINT() { return getToken(SwiftgrammParser.PRINT, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public Print_blContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print_bl; }
	}

	public final Print_blContext print_bl() throws RecognitionException {
		Print_blContext _localctx = new Print_blContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_print_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			match(PRINT);
			setState(172);
			match(OPEN_PARENTHESIS);
			setState(173);
			((Print_blContext)_localctx).expression = expression(0);
			setState(174);
			match(CLOSE_PARENTHESIS);

			        _localctx.instr = instructions.NewPrint(((Print_blContext)_localctx).expression.p)

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

	public static class If_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
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
	}

	public final If_blContext if_bl() throws RecognitionException {
		If_blContext _localctx = new If_blContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_if_bl);
		try {
			setState(203);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(177);
				match(IF);
				setState(178);
				((If_blContext)_localctx).expression = expression(0);
				setState(179);
				match(OPEN_kEY);
				setState(180);
				((If_blContext)_localctx).ifblock = block();
				setState(181);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewIf(((If_blContext)_localctx).expression.p,((If_blContext)_localctx).ifblock.blk,nil)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(184);
				match(IF);
				setState(185);
				((If_blContext)_localctx).expression = expression(0);
				setState(186);
				match(OPEN_kEY);
				setState(187);
				((If_blContext)_localctx).ifblock = block();
				setState(188);
				match(CLOSE_kEY);
				setState(189);
				match(ELSE);
				setState(190);
				match(OPEN_kEY);
				setState(191);
				((If_blContext)_localctx).elseblock = block();
				setState(192);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewIf(((If_blContext)_localctx).expression.p,((If_blContext)_localctx).ifblock.blk,((If_blContext)_localctx).elseblock.blk)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(195);
				match(IF);
				setState(196);
				((If_blContext)_localctx).expression = expression(0);
				setState(197);
				match(OPEN_kEY);
				setState(198);
				((If_blContext)_localctx).ifblock = block();
				setState(199);
				match(CLOSE_kEY);
				setState(200);
				((If_blContext)_localctx).else_if = else_if();

				        _localctx.instr = instructions.NewIf(((If_blContext)_localctx).expression.p,((If_blContext)_localctx).ifblock.blk,((If_blContext)_localctx).else_if.instr)

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

	public static class Else_ifContext extends ParserRuleContext {
		public []interface{} instr;
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
	}

	public final Else_ifContext else_if() throws RecognitionException {
		Else_ifContext _localctx = new Else_ifContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_else_if);
		try {
			setState(234);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(205);
				match(ELSE);
				setState(206);
				match(IF);
				setState(207);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(208);
				match(OPEN_kEY);
				setState(209);
				((Else_ifContext)_localctx).ifblock = block();
				setState(210);
				match(CLOSE_kEY);

				        _localctx.instr = []interface{}{instructions.NewIf(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).ifblock.blk,nil)}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(213);
				match(ELSE);
				setState(214);
				match(IF);
				setState(215);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(216);
				match(OPEN_kEY);
				setState(217);
				((Else_ifContext)_localctx).ifblock = block();
				setState(218);
				match(CLOSE_kEY);
				setState(219);
				match(ELSE);
				setState(220);
				match(OPEN_kEY);
				setState(221);
				((Else_ifContext)_localctx).elseblock = block();
				setState(222);
				match(CLOSE_kEY);

				        _localctx.instr = []interface{}{instructions.NewIf(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).ifblock.blk,((Else_ifContext)_localctx).elseblock.blk)}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(225);
				match(ELSE);
				setState(226);
				match(IF);
				setState(227);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(228);
				match(OPEN_kEY);
				setState(229);
				((Else_ifContext)_localctx).ifblock = block();
				setState(230);
				match(CLOSE_kEY);
				setState(231);
				((Else_ifContext)_localctx).else_if = else_if();

				        _localctx.instr = []interface{}{instructions.NewIf(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).ifblock.blk,((Else_ifContext)_localctx).else_if.instr)}

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

	public static class While_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
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
	}

	public final While_blContext while_bl() throws RecognitionException {
		While_blContext _localctx = new While_blContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_while_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			match(WHILE);
			setState(237);
			((While_blContext)_localctx).expression = expression(0);
			setState(238);
			match(OPEN_kEY);
			setState(239);
			((While_blContext)_localctx).block = block();
			setState(240);
			match(CLOSE_kEY);

			        _localctx.instr = instructions.NewWhile(((While_blContext)_localctx).expression.p,((While_blContext)_localctx).block.blk)

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
	}

	public final For_blContext for_bl() throws RecognitionException {
		For_blContext _localctx = new For_blContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_for_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			match(FOR);
			setState(244);
			((For_blContext)_localctx).ID = match(ID);
			setState(245);
			match(IN);
			setState(246);
			((For_blContext)_localctx).expression1 = expression(0);
			setState(247);
			match(RANGE);
			setState(248);
			((For_blContext)_localctx).expression2 = expression(0);
			setState(249);
			match(OPEN_kEY);
			setState(250);
			((For_blContext)_localctx).block = block();
			setState(251);
			match(CLOSE_kEY);

			        _localctx.instr = instructions.NewFor((((For_blContext)_localctx).ID!=null?((For_blContext)_localctx).ID.getText():null),((For_blContext)_localctx).expression1.p,((For_blContext)_localctx).expression2.p,((For_blContext)_localctx).block.blk)

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

	public static class Guard_blContext extends ParserRuleContext {
		public abstract.Instruction instr;
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
	}

	public final Guard_blContext guard_bl() throws RecognitionException {
		Guard_blContext _localctx = new Guard_blContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_guard_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(254);
			match(GUARD);
			setState(255);
			((Guard_blContext)_localctx).expression = expression(0);
			setState(256);
			match(ELSE);
			setState(257);
			match(OPEN_kEY);
			setState(258);
			((Guard_blContext)_localctx).block = block();
			setState(259);
			match(CLOSE_kEY);

			        _localctx.instr = instructions.NewGuard(((Guard_blContext)_localctx).expression.p,((Guard_blContext)_localctx).block.blk)

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
	}

	public final Vector_blContext vector_bl() throws RecognitionException {
		Vector_blContext _localctx = new Vector_blContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_vector_bl);
		try {
			setState(285);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(262);
				match(VAR);
				setState(263);
				((Vector_blContext)_localctx).ID = match(ID);
				setState(264);
				match(COLON);
				setState(265);
				match(OPEN_BRACKET);
				setState(266);
				((Vector_blContext)_localctx).datatype = datatype();
				setState(267);
				match(CLOSE_BRACKET);
				setState(268);
				match(ASSIGN);
				setState(269);
				match(OPEN_BRACKET);
				setState(270);
				((Vector_blContext)_localctx).array_exp = array_exp();
				setState(271);
				match(CLOSE_BRACKET);

				        
				        _localctx.instr = instructions.NewVector((((Vector_blContext)_localctx).ID!=null?((Vector_blContext)_localctx).ID.getText():null),((Vector_blContext)_localctx).datatype.td,((Vector_blContext)_localctx).array_exp.p)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(274);
				match(VAR);
				setState(275);
				((Vector_blContext)_localctx).ID = match(ID);
				setState(276);
				match(COLON);
				setState(277);
				match(OPEN_BRACKET);
				setState(278);
				((Vector_blContext)_localctx).datatype = datatype();
				setState(279);
				match(CLOSE_BRACKET);
				setState(280);
				match(ASSIGN);
				setState(281);
				match(OPEN_BRACKET);
				setState(282);
				match(CLOSE_BRACKET);

				        
				        _localctx.instr = instructions.NewVector((((Vector_blContext)_localctx).ID!=null?((Vector_blContext)_localctx).ID.getText():null),((Vector_blContext)_localctx).datatype.td,nil)

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
	}

	public final Array_expContext array_exp() throws RecognitionException {
		Array_expContext _localctx = new Array_expContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_array_exp);
		try {
			setState(295);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(287);
				((Array_expContext)_localctx).expression = expression(0);
				setState(288);
				match(COMMA);
				setState(289);
				((Array_expContext)_localctx).array_exp = array_exp();

				        _localctx.p = append(((Array_expContext)_localctx).array_exp.p, ((Array_expContext)_localctx).expression.p)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(292);
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
	}

	public final Function_blContext function_bl() throws RecognitionException {
		Function_blContext _localctx = new Function_blContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_function_bl);
		try {
			setState(329);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(297);
				match(FUNC);
				setState(298);
				((Function_blContext)_localctx).ID = match(ID);
				setState(299);
				match(OPEN_PARENTHESIS);
				setState(300);
				match(CLOSE_PARENTHESIS);
				setState(301);
				match(ARROW);
				setState(302);
				((Function_blContext)_localctx).datatype = datatype();
				setState(303);
				match(OPEN_kEY);
				setState(304);
				((Function_blContext)_localctx).block = block();
				setState(305);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewDeclareFunction((((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getText():null),((Function_blContext)_localctx).datatype.td,[]interface{}{},((Function_blContext)_localctx).block.blk)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(308);
				match(FUNC);
				setState(309);
				((Function_blContext)_localctx).ID = match(ID);
				setState(310);
				match(OPEN_PARENTHESIS);
				setState(311);
				match(CLOSE_PARENTHESIS);
				setState(312);
				match(OPEN_kEY);
				setState(313);
				((Function_blContext)_localctx).block = block();
				setState(314);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewDeclareFunction((((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getText():null),symbol.NIL,[]interface{}{},((Function_blContext)_localctx).block.blk)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(317);
				match(FUNC);
				setState(318);
				((Function_blContext)_localctx).ID = match(ID);
				setState(319);
				match(OPEN_PARENTHESIS);
				setState(320);
				((Function_blContext)_localctx).params = params();
				setState(321);
				match(CLOSE_PARENTHESIS);
				setState(322);
				match(ARROW);
				setState(323);
				((Function_blContext)_localctx).datatype = datatype();
				setState(324);
				match(OPEN_kEY);
				setState(325);
				((Function_blContext)_localctx).block = block();
				setState(326);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewDeclareFunction((((Function_blContext)_localctx).ID!=null?((Function_blContext)_localctx).ID.getText():null),((Function_blContext)_localctx).datatype.td,((Function_blContext)_localctx).params.p,((Function_blContext)_localctx).block.blk)

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

	public static class ParamsContext extends ParserRuleContext {
		public []interface{} p;
		public Token ID;
		public DatatypeContext datatype;
		public ParamsContext params;
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public DatatypeContext datatype() {
			return getRuleContext(DatatypeContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(SwiftgrammParser.COMMA, 0); }
		public ParamsContext params() {
			return getRuleContext(ParamsContext.class,0);
		}
		public ParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_params; }
	}

	public final ParamsContext params() throws RecognitionException {
		ParamsContext _localctx = new ParamsContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_params);
		try {
			setState(343);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(331);
				((ParamsContext)_localctx).ID = match(ID);
				setState(332);
				match(COLON);
				setState(333);
				((ParamsContext)_localctx).datatype = datatype();
				setState(334);
				match(COMMA);
				setState(335);
				((ParamsContext)_localctx).params = params();

				        _localctx.p = append(((ParamsContext)_localctx).params.p,instructions.NewDeclareWithoutValue((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,expressions.NewNative(nil,symbol.NIL)))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(338);
				((ParamsContext)_localctx).ID = match(ID);
				setState(339);
				match(COLON);
				setState(340);
				((ParamsContext)_localctx).datatype = datatype();

				        _localctx.p = []interface{}{instructions.NewDeclareWithoutValue((((ParamsContext)_localctx).ID!=null?((ParamsContext)_localctx).ID.getText():null),((ParamsContext)_localctx).datatype.td,expressions.NewNative(nil,symbol.NIL))}

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
	}

	public final Call_functionContext call_function() throws RecognitionException {
		Call_functionContext _localctx = new Call_functionContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_call_function);
		try {
			setState(355);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(345);
				((Call_functionContext)_localctx).ID = match(ID);
				setState(346);
				match(OPEN_PARENTHESIS);
				setState(347);
				match(CLOSE_PARENTHESIS);

				        _localctx.instr = instructions.NewCallFunction((((Call_functionContext)_localctx).ID!=null?((Call_functionContext)_localctx).ID.getText():null),[]interface{}{})

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(349);
				((Call_functionContext)_localctx).ID = match(ID);
				setState(350);
				match(OPEN_PARENTHESIS);
				setState(351);
				((Call_functionContext)_localctx).list_exp = list_exp();
				setState(352);
				match(CLOSE_PARENTHESIS);

				        _localctx.instr = instructions.NewCallFunction((((Call_functionContext)_localctx).ID!=null?((Call_functionContext)_localctx).ID.getText():null),((Call_functionContext)_localctx).list_exp.p)

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

	public static class List_expContext extends ParserRuleContext {
		public []interface{} p;
		public ExpressionContext expression;
		public List_expContext list_exp;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(SwiftgrammParser.COMMA, 0); }
		public List_expContext list_exp() {
			return getRuleContext(List_expContext.class,0);
		}
		public List_expContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_exp; }
	}

	public final List_expContext list_exp() throws RecognitionException {
		List_expContext _localctx = new List_expContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_list_exp);
		try {
			setState(365);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(357);
				((List_expContext)_localctx).expression = expression(0);
				setState(358);
				match(COMMA);
				setState(359);
				((List_expContext)_localctx).list_exp = list_exp();

				        _localctx.p = append(((List_expContext)_localctx).list_exp.p,((List_expContext)_localctx).expression.p)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(362);
				((List_expContext)_localctx).expression = expression(0);

				        _localctx.p = []interface{}{((List_expContext)_localctx).expression.p}

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
	}

	public final Call_function_expContext call_function_exp() throws RecognitionException {
		Call_function_expContext _localctx = new Call_function_expContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_call_function_exp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(367);
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

	public static class ExpressionContext extends ParserRuleContext {
		public abstract.Expression p;
		public ExpressionContext left;
		public Token oper;
		public ExpressionContext expression;
		public Call_function_expContext call_function_exp;
		public Token NUMBER;
		public Token FLOATT;
		public Token STRING_LITERAL;
		public Token CHARACTER_LITERAL;
		public Token TRUE;
		public Token FALSE;
		public Token ID;
		public ExpressionContext right;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode NOT() { return getToken(SwiftgrammParser.NOT, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public Call_function_expContext call_function_exp() {
			return getRuleContext(Call_function_expContext.class,0);
		}
		public TerminalNode NUMBER() { return getToken(SwiftgrammParser.NUMBER, 0); }
		public TerminalNode FLOATT() { return getToken(SwiftgrammParser.FLOATT, 0); }
		public TerminalNode STRING_LITERAL() { return getToken(SwiftgrammParser.STRING_LITERAL, 0); }
		public TerminalNode CHARACTER_LITERAL() { return getToken(SwiftgrammParser.CHARACTER_LITERAL, 0); }
		public TerminalNode TRUE() { return getToken(SwiftgrammParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(SwiftgrammParser.FALSE, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode NIL() { return getToken(SwiftgrammParser.NIL, 0); }
		public TerminalNode MULTIPLICATION() { return getToken(SwiftgrammParser.MULTIPLICATION, 0); }
		public TerminalNode DIVISION() { return getToken(SwiftgrammParser.DIVISION, 0); }
		public TerminalNode SUMMATION() { return getToken(SwiftgrammParser.SUMMATION, 0); }
		public TerminalNode SUBTRACTION() { return getToken(SwiftgrammParser.SUBTRACTION, 0); }
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
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(398);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				setState(371);
				((ExpressionContext)_localctx).oper = match(NOT);
				setState(372);
				((ExpressionContext)_localctx).expression = expression(11);

				                _localctx.p = expressions.NewLogicOperations(nil,((ExpressionContext)_localctx).expression.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
				        
				}
				break;
			case 2:
				{
				setState(375);
				match(OPEN_PARENTHESIS);
				setState(376);
				((ExpressionContext)_localctx).expression = expression(0);
				setState(377);
				match(CLOSE_PARENTHESIS);
				}
				break;
			case 3:
				{
				setState(379);
				((ExpressionContext)_localctx).call_function_exp = call_function_exp();

				                _localctx.p = ((ExpressionContext)_localctx).call_function_exp.p
				        
				}
				break;
			case 4:
				{
				setState(382);
				((ExpressionContext)_localctx).NUMBER = match(NUMBER);

				                value,err := strconv.Atoi((((ExpressionContext)_localctx).NUMBER!=null?((ExpressionContext)_localctx).NUMBER.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.INT)
				        
				}
				break;
			case 5:
				{
				setState(384);
				((ExpressionContext)_localctx).FLOATT = match(FLOATT);

				                value,err := strconv.ParseFloat((((ExpressionContext)_localctx).FLOATT!=null?((ExpressionContext)_localctx).FLOATT.getText():null),64)
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.FLOAT)
				        
				}
				break;
			case 6:
				{
				setState(386);
				((ExpressionContext)_localctx).STRING_LITERAL = match(STRING_LITERAL);

				                value := (((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getText():null)[1:len((((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getText():null))-1]
				                _localctx.p = expressions.NewNative(value,symbol.STRING)
				        
				}
				break;
			case 7:
				{
				setState(388);
				((ExpressionContext)_localctx).CHARACTER_LITERAL = match(CHARACTER_LITERAL);

				                value := (((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getText():null)[1:len((((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getText():null))-1]
				                _localctx.p = expressions.NewNative(value,symbol.CHAR)
				        
				}
				break;
			case 8:
				{
				setState(390);
				((ExpressionContext)_localctx).TRUE = match(TRUE);

				                value,err := strconv.ParseBool((((ExpressionContext)_localctx).TRUE!=null?((ExpressionContext)_localctx).TRUE.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.BOOL) 
				        
				}
				break;
			case 9:
				{
				setState(392);
				((ExpressionContext)_localctx).FALSE = match(FALSE);

				                value,err := strconv.ParseBool((((ExpressionContext)_localctx).FALSE!=null?((ExpressionContext)_localctx).FALSE.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.BOOL) 
				        
				}
				break;
			case 10:
				{
				setState(394);
				((ExpressionContext)_localctx).ID = match(ID);

				                _localctx.p = expressions.NewIdentifier((((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getText():null))
				        
				}
				break;
			case 11:
				{
				setState(396);
				match(NIL);

				                _localctx.p = expressions.NewNative(nil,symbol.NIL)
				        
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(437);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(435);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(400);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(401);
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
						setState(402);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(19);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(405);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(406);
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
						setState(407);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(18);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(410);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(411);
						((ExpressionContext)_localctx).oper = match(MOD);
						setState(412);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(17);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(415);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(416);
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
						setState(417);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(16);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(420);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(421);
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
						setState(422);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(15);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 6:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(425);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(426);
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
						setState(427);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(14);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 7:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(430);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(431);
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
						setState(432);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(13);

						                          _localctx.p = expressions.NewLogicOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					}
					} 
				}
				setState(439);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
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
	}

	public final DatatypeContext datatype() throws RecognitionException {
		DatatypeContext _localctx = new DatatypeContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_datatype);
		try {
			setState(450);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(440);
				match(INT);

				                _localctx.td = symbol.INT
				        
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(442);
				match(FLOAT);

				                _localctx.td = symbol.FLOAT
				        
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(444);
				match(STRING);

				                _localctx.td = symbol.STRING
				        
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(446);
				match(BOOL);

				                _localctx.td = symbol.BOOL
				        
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 5);
				{
				setState(448);
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
		case 23:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 18);
		case 1:
			return precpred(_ctx, 17);
		case 2:
			return precpred(_ctx, 16);
		case 3:
			return precpred(_ctx, 15);
		case 4:
			return precpred(_ctx, 14);
		case 5:
			return precpred(_ctx, 13);
		case 6:
			return precpred(_ctx, 12);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3D\u01c7\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\3\2\3\2\3\2\3\2\3\3\6\3:\n\3\r\3\16\3;\3\3\3\3\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4m\n\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u0082\n\b\3\t\3\t\3\t"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u0095\n\n"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u00ac\n\13\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00ce\n\r\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u00ed"+
		"\n\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0120\n\22\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\5\23\u012a\n\23\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u014c"+
		"\n\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\5\25"+
		"\u015a\n\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u0166"+
		"\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u0170\n\27\3\30\3\30"+
		"\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\5\31\u0191\n\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\7\31\u01b6\n\31"+
		"\f\31\16\31\u01b9\13\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3"+
		"\32\5\32\u01c5\n\32\3\32\2\3\60\33\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(*,.\60\62\2\b\3\2)*\3\2\'(\3\2\62\63\3\2\64\65\3\2\60\61\3\2"+
		"-.\2\u01e0\2\64\3\2\2\2\49\3\2\2\2\6l\3\2\2\2\bn\3\2\2\2\ns\3\2\2\2\f"+
		"x\3\2\2\2\16\u0081\3\2\2\2\20\u0083\3\2\2\2\22\u0094\3\2\2\2\24\u00ab"+
		"\3\2\2\2\26\u00ad\3\2\2\2\30\u00cd\3\2\2\2\32\u00ec\3\2\2\2\34\u00ee\3"+
		"\2\2\2\36\u00f5\3\2\2\2 \u0100\3\2\2\2\"\u011f\3\2\2\2$\u0129\3\2\2\2"+
		"&\u014b\3\2\2\2(\u0159\3\2\2\2*\u0165\3\2\2\2,\u016f\3\2\2\2.\u0171\3"+
		"\2\2\2\60\u0190\3\2\2\2\62\u01c4\3\2\2\2\64\65\5\4\3\2\65\66\7\2\2\3\66"+
		"\67\b\2\1\2\67\3\3\2\2\28:\5\6\4\298\3\2\2\2:;\3\2\2\2;9\3\2\2\2;<\3\2"+
		"\2\2<=\3\2\2\2=>\b\3\1\2>\5\3\2\2\2?@\5\24\13\2@A\b\4\1\2Am\3\2\2\2BC"+
		"\5\22\n\2CD\b\4\1\2Dm\3\2\2\2EF\5\26\f\2FG\b\4\1\2Gm\3\2\2\2HI\5\30\r"+
		"\2IJ\b\4\1\2Jm\3\2\2\2KL\5\b\5\2LM\b\4\1\2Mm\3\2\2\2NO\5\n\6\2OP\b\4\1"+
		"\2Pm\3\2\2\2QR\5\34\17\2RS\b\4\1\2Sm\3\2\2\2TU\5\36\20\2UV\b\4\1\2Vm\3"+
		"\2\2\2WX\5 \21\2XY\b\4\1\2Ym\3\2\2\2Z[\5\f\7\2[\\\b\4\1\2\\m\3\2\2\2]"+
		"^\5\16\b\2^_\b\4\1\2_m\3\2\2\2`a\5\20\t\2ab\b\4\1\2bm\3\2\2\2cd\5\"\22"+
		"\2de\b\4\1\2em\3\2\2\2fg\5&\24\2gh\b\4\1\2hm\3\2\2\2ij\5*\26\2jk\b\4\1"+
		"\2km\3\2\2\2l?\3\2\2\2lB\3\2\2\2lE\3\2\2\2lH\3\2\2\2lK\3\2\2\2lN\3\2\2"+
		"\2lQ\3\2\2\2lT\3\2\2\2lW\3\2\2\2lZ\3\2\2\2l]\3\2\2\2l`\3\2\2\2lc\3\2\2"+
		"\2lf\3\2\2\2li\3\2\2\2m\7\3\2\2\2no\7!\2\2op\7$\2\2pq\5\60\31\2qr\b\5"+
		"\1\2r\t\3\2\2\2st\7!\2\2tu\7%\2\2uv\5\60\31\2vw\b\6\1\2w\13\3\2\2\2xy"+
		"\7\22\2\2yz\b\7\1\2z\r\3\2\2\2{|\7\31\2\2|}\5\60\31\2}~\b\b\1\2~\u0082"+
		"\3\2\2\2\177\u0080\7\31\2\2\u0080\u0082\b\b\1\2\u0081{\3\2\2\2\u0081\177"+
		"\3\2\2\2\u0082\17\3\2\2\2\u0083\u0084\7\30\2\2\u0084\u0085\b\t\1\2\u0085"+
		"\21\3\2\2\2\u0086\u0087\7\f\2\2\u0087\u0088\7!\2\2\u0088\u0089\79\2\2"+
		"\u0089\u008a\5\62\32\2\u008a\u008b\7\66\2\2\u008b\u008c\5\60\31\2\u008c"+
		"\u008d\b\n\1\2\u008d\u0095\3\2\2\2\u008e\u008f\7\f\2\2\u008f\u0090\7!"+
		"\2\2\u0090\u0091\7\66\2\2\u0091\u0092\5\60\31\2\u0092\u0093\b\n\1\2\u0093"+
		"\u0095\3\2\2\2\u0094\u0086\3\2\2\2\u0094\u008e\3\2\2\2\u0095\23\3\2\2"+
		"\2\u0096\u0097\7\13\2\2\u0097\u0098\7!\2\2\u0098\u0099\79\2\2\u0099\u009a"+
		"\5\62\32\2\u009a\u009b\7\66\2\2\u009b\u009c\5\60\31\2\u009c\u009d\b\13"+
		"\1\2\u009d\u00ac\3\2\2\2\u009e\u009f\7\13\2\2\u009f\u00a0\7!\2\2\u00a0"+
		"\u00a1\7\66\2\2\u00a1\u00a2\5\60\31\2\u00a2\u00a3\b\13\1\2\u00a3\u00ac"+
		"\3\2\2\2\u00a4\u00a5\7\13\2\2\u00a5\u00a6\7!\2\2\u00a6\u00a7\79\2\2\u00a7"+
		"\u00a8\5\62\32\2\u00a8\u00a9\7,\2\2\u00a9\u00aa\b\13\1\2\u00aa\u00ac\3"+
		"\2\2\2\u00ab\u0096\3\2\2\2\u00ab\u009e\3\2\2\2\u00ab\u00a4\3\2\2\2\u00ac"+
		"\25\3\2\2\2\u00ad\u00ae\7\r\2\2\u00ae\u00af\7;\2\2\u00af\u00b0\5\60\31"+
		"\2\u00b0\u00b1\7<\2\2\u00b1\u00b2\b\f\1\2\u00b2\27\3\2\2\2\u00b3\u00b4"+
		"\7\16\2\2\u00b4\u00b5\5\60\31\2\u00b5\u00b6\7=\2\2\u00b6\u00b7\5\4\3\2"+
		"\u00b7\u00b8\7>\2\2\u00b8\u00b9\b\r\1\2\u00b9\u00ce\3\2\2\2\u00ba\u00bb"+
		"\7\16\2\2\u00bb\u00bc\5\60\31\2\u00bc\u00bd\7=\2\2\u00bd\u00be\5\4\3\2"+
		"\u00be\u00bf\7>\2\2\u00bf\u00c0\7\17\2\2\u00c0\u00c1\7=\2\2\u00c1\u00c2"+
		"\5\4\3\2\u00c2\u00c3\7>\2\2\u00c3\u00c4\b\r\1\2\u00c4\u00ce\3\2\2\2\u00c5"+
		"\u00c6\7\16\2\2\u00c6\u00c7\5\60\31\2\u00c7\u00c8\7=\2\2\u00c8\u00c9\5"+
		"\4\3\2\u00c9\u00ca\7>\2\2\u00ca\u00cb\5\32\16\2\u00cb\u00cc\b\r\1\2\u00cc"+
		"\u00ce\3\2\2\2\u00cd\u00b3\3\2\2\2\u00cd\u00ba\3\2\2\2\u00cd\u00c5\3\2"+
		"\2\2\u00ce\31\3\2\2\2\u00cf\u00d0\7\17\2\2\u00d0\u00d1\7\16\2\2\u00d1"+
		"\u00d2\5\60\31\2\u00d2\u00d3\7=\2\2\u00d3\u00d4\5\4\3\2\u00d4\u00d5\7"+
		">\2\2\u00d5\u00d6\b\16\1\2\u00d6\u00ed\3\2\2\2\u00d7\u00d8\7\17\2\2\u00d8"+
		"\u00d9\7\16\2\2\u00d9\u00da\5\60\31\2\u00da\u00db\7=\2\2\u00db\u00dc\5"+
		"\4\3\2\u00dc\u00dd\7>\2\2\u00dd\u00de\7\17\2\2\u00de\u00df\7=\2\2\u00df"+
		"\u00e0\5\4\3\2\u00e0\u00e1\7>\2\2\u00e1\u00e2\b\16\1\2\u00e2\u00ed\3\2"+
		"\2\2\u00e3\u00e4\7\17\2\2\u00e4\u00e5\7\16\2\2\u00e5\u00e6\5\60\31\2\u00e6"+
		"\u00e7\7=\2\2\u00e7\u00e8\5\4\3\2\u00e8\u00e9\7>\2\2\u00e9\u00ea\5\32"+
		"\16\2\u00ea\u00eb\b\16\1\2\u00eb\u00ed\3\2\2\2\u00ec\u00cf\3\2\2\2\u00ec"+
		"\u00d7\3\2\2\2\u00ec\u00e3\3\2\2\2\u00ed\33\3\2\2\2\u00ee\u00ef\7\24\2"+
		"\2\u00ef\u00f0\5\60\31\2\u00f0\u00f1\7=\2\2\u00f1\u00f2\5\4\3\2\u00f2"+
		"\u00f3\7>\2\2\u00f3\u00f4\b\17\1\2\u00f4\35\3\2\2\2\u00f5\u00f6\7\25\2"+
		"\2\u00f6\u00f7\7!\2\2\u00f7\u00f8\7\26\2\2\u00f8\u00f9\5\60\31\2\u00f9"+
		"\u00fa\7&\2\2\u00fa\u00fb\5\60\31\2\u00fb\u00fc\7=\2\2\u00fc\u00fd\5\4"+
		"\3\2\u00fd\u00fe\7>\2\2\u00fe\u00ff\b\20\1\2\u00ff\37\3\2\2\2\u0100\u0101"+
		"\7\27\2\2\u0101\u0102\5\60\31\2\u0102\u0103\7\17\2\2\u0103\u0104\7=\2"+
		"\2\u0104\u0105\5\4\3\2\u0105\u0106\7>\2\2\u0106\u0107\b\21\1\2\u0107!"+
		"\3\2\2\2\u0108\u0109\7\13\2\2\u0109\u010a\7!\2\2\u010a\u010b\79\2\2\u010b"+
		"\u010c\7?\2\2\u010c\u010d\5\62\32\2\u010d\u010e\7@\2\2\u010e\u010f\7\66"+
		"\2\2\u010f\u0110\7?\2\2\u0110\u0111\5$\23\2\u0111\u0112\7@\2\2\u0112\u0113"+
		"\b\22\1\2\u0113\u0120\3\2\2\2\u0114\u0115\7\13\2\2\u0115\u0116\7!\2\2"+
		"\u0116\u0117\79\2\2\u0117\u0118\7?\2\2\u0118\u0119\5\62\32\2\u0119\u011a"+
		"\7@\2\2\u011a\u011b\7\66\2\2\u011b\u011c\7?\2\2\u011c\u011d\7@\2\2\u011d"+
		"\u011e\b\22\1\2\u011e\u0120\3\2\2\2\u011f\u0108\3\2\2\2\u011f\u0114\3"+
		"\2\2\2\u0120#\3\2\2\2\u0121\u0122\5\60\31\2\u0122\u0123\78\2\2\u0123\u0124"+
		"\5$\23\2\u0124\u0125\b\23\1\2\u0125\u012a\3\2\2\2\u0126\u0127\5\60\31"+
		"\2\u0127\u0128\b\23\1\2\u0128\u012a\3\2\2\2\u0129\u0121\3\2\2\2\u0129"+
		"\u0126\3\2\2\2\u012a%\3\2\2\2\u012b\u012c\7\32\2\2\u012c\u012d\7!\2\2"+
		"\u012d\u012e\7;\2\2\u012e\u012f\7<\2\2\u012f\u0130\7A\2\2\u0130\u0131"+
		"\5\62\32\2\u0131\u0132\7=\2\2\u0132\u0133\5\4\3\2\u0133\u0134\7>\2\2\u0134"+
		"\u0135\b\24\1\2\u0135\u014c\3\2\2\2\u0136\u0137\7\32\2\2\u0137\u0138\7"+
		"!\2\2\u0138\u0139\7;\2\2\u0139\u013a\7<\2\2\u013a\u013b\7=\2\2\u013b\u013c"+
		"\5\4\3\2\u013c\u013d\7>\2\2\u013d\u013e\b\24\1\2\u013e\u014c\3\2\2\2\u013f"+
		"\u0140\7\32\2\2\u0140\u0141\7!\2\2\u0141\u0142\7;\2\2\u0142\u0143\5(\25"+
		"\2\u0143\u0144\7<\2\2\u0144\u0145\7A\2\2\u0145\u0146\5\62\32\2\u0146\u0147"+
		"\7=\2\2\u0147\u0148\5\4\3\2\u0148\u0149\7>\2\2\u0149\u014a\b\24\1\2\u014a"+
		"\u014c\3\2\2\2\u014b\u012b\3\2\2\2\u014b\u0136\3\2\2\2\u014b\u013f\3\2"+
		"\2\2\u014c\'\3\2\2\2\u014d\u014e\7!\2\2\u014e\u014f\79\2\2\u014f\u0150"+
		"\5\62\32\2\u0150\u0151\78\2\2\u0151\u0152\5(\25\2\u0152\u0153\b\25\1\2"+
		"\u0153\u015a\3\2\2\2\u0154\u0155\7!\2\2\u0155\u0156\79\2\2\u0156\u0157"+
		"\5\62\32\2\u0157\u0158\b\25\1\2\u0158\u015a\3\2\2\2\u0159\u014d\3\2\2"+
		"\2\u0159\u0154\3\2\2\2\u015a)\3\2\2\2\u015b\u015c\7!\2\2\u015c\u015d\7"+
		";\2\2\u015d\u015e\7<\2\2\u015e\u0166\b\26\1\2\u015f\u0160\7!\2\2\u0160"+
		"\u0161\7;\2\2\u0161\u0162\5,\27\2\u0162\u0163\7<\2\2\u0163\u0164\b\26"+
		"\1\2\u0164\u0166\3\2\2\2\u0165\u015b\3\2\2\2\u0165\u015f\3\2\2\2\u0166"+
		"+\3\2\2\2\u0167\u0168\5\60\31\2\u0168\u0169\78\2\2\u0169\u016a\5,\27\2"+
		"\u016a\u016b\b\27\1\2\u016b\u0170\3\2\2\2\u016c\u016d\5\60\31\2\u016d"+
		"\u016e\b\27\1\2\u016e\u0170\3\2\2\2\u016f\u0167\3\2\2\2\u016f\u016c\3"+
		"\2\2\2\u0170-\3\2\2\2\u0171\u0172\5*\26\2\u0172\u0173\b\30\1\2\u0173/"+
		"\3\2\2\2\u0174\u0175\b\31\1\2\u0175\u0176\7/\2\2\u0176\u0177\5\60\31\r"+
		"\u0177\u0178\b\31\1\2\u0178\u0191\3\2\2\2\u0179\u017a\7;\2\2\u017a\u017b"+
		"\5\60\31\2\u017b\u017c\7<\2\2\u017c\u0191\3\2\2\2\u017d\u017e\5.\30\2"+
		"\u017e\u017f\b\31\1\2\u017f\u0191\3\2\2\2\u0180\u0181\7\37\2\2\u0181\u0191"+
		"\b\31\1\2\u0182\u0183\7 \2\2\u0183\u0191\b\31\1\2\u0184\u0185\7#\2\2\u0185"+
		"\u0191\b\31\1\2\u0186\u0187\7\"\2\2\u0187\u0191\b\31\1\2\u0188\u0189\7"+
		"\b\2\2\u0189\u0191\b\31\1\2\u018a\u018b\7\t\2\2\u018b\u0191\b\31\1\2\u018c"+
		"\u018d\7!\2\2\u018d\u0191\b\31\1\2\u018e\u018f\7\n\2\2\u018f\u0191\b\31"+
		"\1\2\u0190\u0174\3\2\2\2\u0190\u0179\3\2\2\2\u0190\u017d\3\2\2\2\u0190"+
		"\u0180\3\2\2\2\u0190\u0182\3\2\2\2\u0190\u0184\3\2\2\2\u0190\u0186\3\2"+
		"\2\2\u0190\u0188\3\2\2\2\u0190\u018a\3\2\2\2\u0190\u018c\3\2\2\2\u0190"+
		"\u018e\3\2\2\2\u0191\u01b7\3\2\2\2\u0192\u0193\f\24\2\2\u0193\u0194\t"+
		"\2\2\2\u0194\u0195\5\60\31\25\u0195\u0196\b\31\1\2\u0196\u01b6\3\2\2\2"+
		"\u0197\u0198\f\23\2\2\u0198\u0199\t\3\2\2\u0199\u019a\5\60\31\24\u019a"+
		"\u019b\b\31\1\2\u019b\u01b6\3\2\2\2\u019c\u019d\f\22\2\2\u019d\u019e\7"+
		"+\2\2\u019e\u019f\5\60\31\23\u019f\u01a0\b\31\1\2\u01a0\u01b6\3\2\2\2"+
		"\u01a1\u01a2\f\21\2\2\u01a2\u01a3\t\4\2\2\u01a3\u01a4\5\60\31\22\u01a4"+
		"\u01a5\b\31\1\2\u01a5\u01b6\3\2\2\2\u01a6\u01a7\f\20\2\2\u01a7\u01a8\t"+
		"\5\2\2\u01a8\u01a9\5\60\31\21\u01a9\u01aa\b\31\1\2\u01aa\u01b6\3\2\2\2"+
		"\u01ab\u01ac\f\17\2\2\u01ac\u01ad\t\6\2\2\u01ad\u01ae\5\60\31\20\u01ae"+
		"\u01af\b\31\1\2\u01af\u01b6\3\2\2\2\u01b0\u01b1\f\16\2\2\u01b1\u01b2\t"+
		"\7\2\2\u01b2\u01b3\5\60\31\17\u01b3\u01b4\b\31\1\2\u01b4\u01b6\3\2\2\2"+
		"\u01b5\u0192\3\2\2\2\u01b5\u0197\3\2\2\2\u01b5\u019c\3\2\2\2\u01b5\u01a1"+
		"\3\2\2\2\u01b5\u01a6\3\2\2\2\u01b5\u01ab\3\2\2\2\u01b5\u01b0\3\2\2\2\u01b6"+
		"\u01b9\3\2\2\2\u01b7\u01b5\3\2\2\2\u01b7\u01b8\3\2\2\2\u01b8\61\3\2\2"+
		"\2\u01b9\u01b7\3\2\2\2\u01ba\u01bb\7\3\2\2\u01bb\u01c5\b\32\1\2\u01bc"+
		"\u01bd\7\4\2\2\u01bd\u01c5\b\32\1\2\u01be\u01bf\7\5\2\2\u01bf\u01c5\b"+
		"\32\1\2\u01c0\u01c1\7\6\2\2\u01c1\u01c5\b\32\1\2\u01c2\u01c3\7\7\2\2\u01c3"+
		"\u01c5\b\32\1\2\u01c4\u01ba\3\2\2\2\u01c4\u01bc\3\2\2\2\u01c4\u01be\3"+
		"\2\2\2\u01c4\u01c0\3\2\2\2\u01c4\u01c2\3\2\2\2\u01c5\63\3\2\2\2\23;l\u0081"+
		"\u0094\u00ab\u00cd\u00ec\u011f\u0129\u014b\u0159\u0165\u016f\u0190\u01b5"+
		"\u01b7\u01c4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}