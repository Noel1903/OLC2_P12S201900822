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
		STRING_LITERAL=31, ID=32, CHARACTER_LITERAL=33, INCREMENT=34, DECREMENT=35, 
		RANGE=36, SUMMATION=37, SUBTRACTION=38, MULTIPLICATION=39, DIVISION=40, 
		MOD=41, QUESTION_MARK=42, OR=43, AND=44, NOT=45, EQUAL=46, NOT_EQUAL=47, 
		LESS_THAN=48, LESS_THAN_EQUAL=49, GREATER_THAN=50, GREATER_THAN_EQUAL=51, 
		ASSIGN=52, DOT=53, COMMA=54, COLON=55, SEMICOLON=56, OPEN_PARENTHESIS=57, 
		CLOSE_PARENTHESIS=58, OPEN_kEY=59, CLOSE_kEY=60, WHITESPACE=61, COMMENT=62, 
		LINE_COMMENT=63;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_sentence = 2, RULE_increment_bl = 3, 
		RULE_decrement_bl = 4, RULE_break_bl = 5, RULE_return_bl = 6, RULE_declare_let = 7, 
		RULE_declare_var = 8, RULE_print_bl = 9, RULE_if_bl = 10, RULE_else_if = 11, 
		RULE_while_bl = 12, RULE_for_bl = 13, RULE_expression = 14, RULE_datatype = 15;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "sentence", "increment_bl", "decrement_bl", "break_bl", 
			"return_bl", "declare_let", "declare_var", "print_bl", "if_bl", "else_if", 
			"while_bl", "for_bl", "expression", "datatype"
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
			"')'", "'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE", 
			"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK", 
			"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC", 
			"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "FLOATT", "STRING_LITERAL", 
			"ID", "CHARACTER_LITERAL", "INCREMENT", "DECREMENT", "RANGE", "SUMMATION", 
			"SUBTRACTION", "MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK", 
			"OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", 
			"GREATER_THAN", "GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON", 
			"SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", 
			"WHITESPACE", "COMMENT", "LINE_COMMENT"
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
			setState(32);
			((SContext)_localctx).block = block();
			setState(33);
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
			setState(37); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(36);
				((BlockContext)_localctx).sentence = sentence();
				((BlockContext)_localctx).instr.add(((BlockContext)_localctx).sentence);
				}
				}
				setState(39); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << PRINT) | (1L << IF) | (1L << BREAK) | (1L << WHILE) | (1L << FOR) | (1L << RETURN) | (1L << ID))) != 0) );

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
		public Break_blContext break_bl;
		public Return_blContext return_bl;
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
		public Break_blContext break_bl() {
			return getRuleContext(Break_blContext.class,0);
		}
		public Return_blContext return_bl() {
			return getRuleContext(Return_blContext.class,0);
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
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(43);
				((SentenceContext)_localctx).declare_var = declare_var();
				_localctx.instr = ((SentenceContext)_localctx).declare_var.instr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(46);
				((SentenceContext)_localctx).declare_let = declare_let();
				_localctx.instr = ((SentenceContext)_localctx).declare_let.instr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(49);
				((SentenceContext)_localctx).print_bl = print_bl();
				_localctx.instr = ((SentenceContext)_localctx).print_bl.instr
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(52);
				((SentenceContext)_localctx).if_bl = if_bl();
				_localctx.instr = ((SentenceContext)_localctx).if_bl.instr
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(55);
				((SentenceContext)_localctx).increment_bl = increment_bl();
				_localctx.instr = ((SentenceContext)_localctx).increment_bl.instr
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(58);
				((SentenceContext)_localctx).decrement_bl = decrement_bl();
				_localctx.instr = ((SentenceContext)_localctx).decrement_bl.instr
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(61);
				((SentenceContext)_localctx).while_bl = while_bl();
				_localctx.instr = ((SentenceContext)_localctx).while_bl.instr
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(64);
				((SentenceContext)_localctx).for_bl = for_bl();
				_localctx.instr = ((SentenceContext)_localctx).for_bl.instr
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(67);
				((SentenceContext)_localctx).break_bl = break_bl();
				_localctx.instr = ((SentenceContext)_localctx).break_bl.instr
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(70);
				((SentenceContext)_localctx).return_bl = return_bl();
				_localctx.instr = ((SentenceContext)_localctx).return_bl.instr
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
			setState(75);
			((Increment_blContext)_localctx).ID = match(ID);
			setState(76);
			((Increment_blContext)_localctx).INCREMENT = match(INCREMENT);
			setState(77);
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
			setState(80);
			((Decrement_blContext)_localctx).ID = match(ID);
			setState(81);
			((Decrement_blContext)_localctx).DECREMENT = match(DECREMENT);
			setState(82);
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
			setState(85);
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
			setState(94);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(88);
				match(RETURN);
				setState(89);
				((Return_blContext)_localctx).expression = expression(0);

				        _localctx.instr = instructions.NewReturn(((Return_blContext)_localctx).expression.p)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(92);
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
		enterRule(_localctx, 14, RULE_declare_let);
		try {
			setState(110);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(96);
				match(LET);
				setState(97);
				((Declare_letContext)_localctx).ID = match(ID);
				setState(98);
				match(COLON);
				setState(99);
				((Declare_letContext)_localctx).datatype = datatype();
				setState(100);
				match(ASSIGN);
				setState(101);
				((Declare_letContext)_localctx).expression = expression(0);

				                _localctx.instr = instructions.NewLet((((Declare_letContext)_localctx).ID!=null?((Declare_letContext)_localctx).ID.getText():null),((Declare_letContext)_localctx).datatype.td,((Declare_letContext)_localctx).expression.p)
				        
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(104);
				match(LET);
				setState(105);
				((Declare_letContext)_localctx).ID = match(ID);
				setState(106);
				match(ASSIGN);
				setState(107);
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
		enterRule(_localctx, 16, RULE_declare_var);
		try {
			setState(133);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(112);
				match(VAR);
				setState(113);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(114);
				match(COLON);
				setState(115);
				((Declare_varContext)_localctx).datatype = datatype();
				setState(116);
				match(ASSIGN);
				setState(117);
				((Declare_varContext)_localctx).expression = expression(0);

				                        _localctx.instr = instructions.NewDeclareWithValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),((Declare_varContext)_localctx).datatype.td,((Declare_varContext)_localctx).expression.p)
				                
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(120);
				match(VAR);
				setState(121);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(122);
				match(ASSIGN);
				setState(123);
				((Declare_varContext)_localctx).expression = expression(0);

				                        _localctx.instr = instructions.NewDeclareWithValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),symbol.UNDEFINED,((Declare_varContext)_localctx).expression.p)
				                
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(126);
				match(VAR);
				setState(127);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(128);
				match(COLON);
				setState(129);
				((Declare_varContext)_localctx).datatype = datatype();
				setState(130);
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
		enterRule(_localctx, 18, RULE_print_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(PRINT);
			setState(136);
			match(OPEN_PARENTHESIS);
			setState(137);
			((Print_blContext)_localctx).expression = expression(0);
			setState(138);
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
		enterRule(_localctx, 20, RULE_if_bl);
		try {
			setState(167);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(141);
				match(IF);
				setState(142);
				((If_blContext)_localctx).expression = expression(0);
				setState(143);
				match(OPEN_kEY);
				setState(144);
				((If_blContext)_localctx).ifblock = block();
				setState(145);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewIf(((If_blContext)_localctx).expression.p,((If_blContext)_localctx).ifblock.blk,nil)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(148);
				match(IF);
				setState(149);
				((If_blContext)_localctx).expression = expression(0);
				setState(150);
				match(OPEN_kEY);
				setState(151);
				((If_blContext)_localctx).ifblock = block();
				setState(152);
				match(CLOSE_kEY);
				setState(153);
				match(ELSE);
				setState(154);
				match(OPEN_kEY);
				setState(155);
				((If_blContext)_localctx).elseblock = block();
				setState(156);
				match(CLOSE_kEY);

				        _localctx.instr = instructions.NewIf(((If_blContext)_localctx).expression.p,((If_blContext)_localctx).ifblock.blk,((If_blContext)_localctx).elseblock.blk)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(159);
				match(IF);
				setState(160);
				((If_blContext)_localctx).expression = expression(0);
				setState(161);
				match(OPEN_kEY);
				setState(162);
				((If_blContext)_localctx).ifblock = block();
				setState(163);
				match(CLOSE_kEY);
				setState(164);
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
		enterRule(_localctx, 22, RULE_else_if);
		try {
			setState(198);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				match(ELSE);
				setState(170);
				match(IF);
				setState(171);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(172);
				match(OPEN_kEY);
				setState(173);
				((Else_ifContext)_localctx).ifblock = block();
				setState(174);
				match(CLOSE_kEY);

				        _localctx.instr = []interface{}{instructions.NewIf(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).ifblock.blk,nil)}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(177);
				match(ELSE);
				setState(178);
				match(IF);
				setState(179);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(180);
				match(OPEN_kEY);
				setState(181);
				((Else_ifContext)_localctx).ifblock = block();
				setState(182);
				match(CLOSE_kEY);
				setState(183);
				match(ELSE);
				setState(184);
				match(OPEN_kEY);
				setState(185);
				((Else_ifContext)_localctx).elseblock = block();
				setState(186);
				match(CLOSE_kEY);

				        _localctx.instr = []interface{}{instructions.NewIf(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).ifblock.blk,((Else_ifContext)_localctx).elseblock.blk)}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(189);
				match(ELSE);
				setState(190);
				match(IF);
				setState(191);
				((Else_ifContext)_localctx).expression = expression(0);
				setState(192);
				match(OPEN_kEY);
				setState(193);
				((Else_ifContext)_localctx).ifblock = block();
				setState(194);
				match(CLOSE_kEY);
				setState(195);
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
		enterRule(_localctx, 24, RULE_while_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
			match(WHILE);
			setState(201);
			((While_blContext)_localctx).expression = expression(0);
			setState(202);
			match(OPEN_kEY);
			setState(203);
			((While_blContext)_localctx).block = block();
			setState(204);
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
		enterRule(_localctx, 26, RULE_for_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			match(FOR);
			setState(208);
			((For_blContext)_localctx).ID = match(ID);
			setState(209);
			match(IN);
			setState(210);
			((For_blContext)_localctx).expression1 = expression(0);
			setState(211);
			match(RANGE);
			setState(212);
			((For_blContext)_localctx).expression2 = expression(0);
			setState(213);
			match(OPEN_kEY);
			setState(214);
			((For_blContext)_localctx).block = block();
			setState(215);
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

	public static class ExpressionContext extends ParserRuleContext {
		public abstract.Expression p;
		public ExpressionContext left;
		public Token oper;
		public ExpressionContext expression;
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
		int _startState = 28;
		enterRecursionRule(_localctx, 28, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				setState(219);
				((ExpressionContext)_localctx).oper = match(NOT);
				setState(220);
				((ExpressionContext)_localctx).expression = expression(10);

				                _localctx.p = expressions.NewLogicOperations(nil,((ExpressionContext)_localctx).expression.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
				        
				}
				break;
			case OPEN_PARENTHESIS:
				{
				setState(223);
				match(OPEN_PARENTHESIS);
				setState(224);
				((ExpressionContext)_localctx).expression = expression(0);
				setState(225);
				match(CLOSE_PARENTHESIS);
				}
				break;
			case NUMBER:
				{
				setState(227);
				((ExpressionContext)_localctx).NUMBER = match(NUMBER);

				                value,err := strconv.Atoi((((ExpressionContext)_localctx).NUMBER!=null?((ExpressionContext)_localctx).NUMBER.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.INT)
				        
				}
				break;
			case FLOATT:
				{
				setState(229);
				((ExpressionContext)_localctx).FLOATT = match(FLOATT);

				                value,err := strconv.ParseFloat((((ExpressionContext)_localctx).FLOATT!=null?((ExpressionContext)_localctx).FLOATT.getText():null),64)
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.FLOAT)
				        
				}
				break;
			case STRING_LITERAL:
				{
				setState(231);
				((ExpressionContext)_localctx).STRING_LITERAL = match(STRING_LITERAL);

				                value := (((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getText():null)[1:len((((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getText():null))-1]
				                _localctx.p = expressions.NewNative(value,symbol.STRING)
				        
				}
				break;
			case CHARACTER_LITERAL:
				{
				setState(233);
				((ExpressionContext)_localctx).CHARACTER_LITERAL = match(CHARACTER_LITERAL);

				                value := (((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getText():null)[1:len((((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getText():null))-1]
				                _localctx.p = expressions.NewNative(value,symbol.CHAR)
				        
				}
				break;
			case TRUE:
				{
				setState(235);
				((ExpressionContext)_localctx).TRUE = match(TRUE);

				                value,err := strconv.ParseBool((((ExpressionContext)_localctx).TRUE!=null?((ExpressionContext)_localctx).TRUE.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.BOOL) 
				        
				}
				break;
			case FALSE:
				{
				setState(237);
				((ExpressionContext)_localctx).FALSE = match(FALSE);

				                value,err := strconv.ParseBool((((ExpressionContext)_localctx).FALSE!=null?((ExpressionContext)_localctx).FALSE.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.BOOL) 
				        
				}
				break;
			case ID:
				{
				setState(239);
				((ExpressionContext)_localctx).ID = match(ID);

				                _localctx.p = expressions.NewIdentifier((((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getText():null))
				        
				}
				break;
			case NIL:
				{
				setState(241);
				match(NIL);

				                _localctx.p = expressions.NewNative(nil,symbol.NIL)
				        
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(282);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(280);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(245);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(246);
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
						setState(247);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(18);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(250);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(251);
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
						setState(252);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(17);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(255);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(256);
						((ExpressionContext)_localctx).oper = match(MOD);
						setState(257);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(16);

						                          _localctx.p = expressions.NewArithmeticOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(260);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(261);
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
						setState(262);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(15);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(265);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(266);
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
						setState(267);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(14);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 6:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(270);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(271);
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
						setState(272);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(13);

						                          _localctx.p = expressions.NewRelationalOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					case 7:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(275);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(276);
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
						setState(277);
						((ExpressionContext)_localctx).right = ((ExpressionContext)_localctx).expression = expression(12);

						                          _localctx.p = expressions.NewLogicOperations(((ExpressionContext)_localctx).left.p,((ExpressionContext)_localctx).right.p,(((ExpressionContext)_localctx).oper!=null?((ExpressionContext)_localctx).oper.getText():null))
						                  
						}
						break;
					}
					} 
				}
				setState(284);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
		public TerminalNode STRING_LITERAL() { return getToken(SwiftgrammParser.STRING_LITERAL, 0); }
		public TerminalNode BOOL() { return getToken(SwiftgrammParser.BOOL, 0); }
		public TerminalNode CHARACTER_LITERAL() { return getToken(SwiftgrammParser.CHARACTER_LITERAL, 0); }
		public DatatypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_datatype; }
	}

	public final DatatypeContext datatype() throws RecognitionException {
		DatatypeContext _localctx = new DatatypeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_datatype);
		try {
			setState(295);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(285);
				match(INT);

				                _localctx.td = symbol.INT
				        
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(287);
				match(FLOAT);

				                _localctx.td = symbol.FLOAT
				        
				}
				break;
			case STRING_LITERAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(289);
				match(STRING_LITERAL);

				                _localctx.td = symbol.STRING
				        
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(291);
				match(BOOL);

				                _localctx.td = symbol.BOOL
				        
				}
				break;
			case CHARACTER_LITERAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(293);
				match(CHARACTER_LITERAL);

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
		case 14:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 17);
		case 1:
			return precpred(_ctx, 16);
		case 2:
			return precpred(_ctx, 15);
		case 3:
			return precpred(_ctx, 14);
		case 4:
			return precpred(_ctx, 13);
		case 5:
			return precpred(_ctx, 12);
		case 6:
			return precpred(_ctx, 11);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3A\u012c\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\3\2\3\2\3"+
		"\2\3\2\3\3\6\3(\n\3\r\3\16\3)\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\5\4L\n\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3"+
		"\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\5\ba\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\tq\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u0088\n\n\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00aa"+
		"\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00c9\n\r\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\5\20\u00f6\n\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\7\20\u011b\n\20\f\20"+
		"\16\20\u011e\13\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5"+
		"\21\u012a\n\21\3\21\2\3\36\22\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 "+
		"\2\b\3\2)*\3\2\'(\3\2\62\63\3\2\64\65\3\2\60\61\3\2-.\2\u0141\2\"\3\2"+
		"\2\2\4\'\3\2\2\2\6K\3\2\2\2\bM\3\2\2\2\nR\3\2\2\2\fW\3\2\2\2\16`\3\2\2"+
		"\2\20p\3\2\2\2\22\u0087\3\2\2\2\24\u0089\3\2\2\2\26\u00a9\3\2\2\2\30\u00c8"+
		"\3\2\2\2\32\u00ca\3\2\2\2\34\u00d1\3\2\2\2\36\u00f5\3\2\2\2 \u0129\3\2"+
		"\2\2\"#\5\4\3\2#$\7\2\2\3$%\b\2\1\2%\3\3\2\2\2&(\5\6\4\2\'&\3\2\2\2()"+
		"\3\2\2\2)\'\3\2\2\2)*\3\2\2\2*+\3\2\2\2+,\b\3\1\2,\5\3\2\2\2-.\5\22\n"+
		"\2./\b\4\1\2/L\3\2\2\2\60\61\5\20\t\2\61\62\b\4\1\2\62L\3\2\2\2\63\64"+
		"\5\24\13\2\64\65\b\4\1\2\65L\3\2\2\2\66\67\5\26\f\2\678\b\4\1\28L\3\2"+
		"\2\29:\5\b\5\2:;\b\4\1\2;L\3\2\2\2<=\5\n\6\2=>\b\4\1\2>L\3\2\2\2?@\5\32"+
		"\16\2@A\b\4\1\2AL\3\2\2\2BC\5\34\17\2CD\b\4\1\2DL\3\2\2\2EF\5\f\7\2FG"+
		"\b\4\1\2GL\3\2\2\2HI\5\16\b\2IJ\b\4\1\2JL\3\2\2\2K-\3\2\2\2K\60\3\2\2"+
		"\2K\63\3\2\2\2K\66\3\2\2\2K9\3\2\2\2K<\3\2\2\2K?\3\2\2\2KB\3\2\2\2KE\3"+
		"\2\2\2KH\3\2\2\2L\7\3\2\2\2MN\7\"\2\2NO\7$\2\2OP\5\36\20\2PQ\b\5\1\2Q"+
		"\t\3\2\2\2RS\7\"\2\2ST\7%\2\2TU\5\36\20\2UV\b\6\1\2V\13\3\2\2\2WX\7\22"+
		"\2\2XY\b\7\1\2Y\r\3\2\2\2Z[\7\31\2\2[\\\5\36\20\2\\]\b\b\1\2]a\3\2\2\2"+
		"^_\7\31\2\2_a\b\b\1\2`Z\3\2\2\2`^\3\2\2\2a\17\3\2\2\2bc\7\f\2\2cd\7\""+
		"\2\2de\79\2\2ef\5 \21\2fg\7\66\2\2gh\5\36\20\2hi\b\t\1\2iq\3\2\2\2jk\7"+
		"\f\2\2kl\7\"\2\2lm\7\66\2\2mn\5\36\20\2no\b\t\1\2oq\3\2\2\2pb\3\2\2\2"+
		"pj\3\2\2\2q\21\3\2\2\2rs\7\13\2\2st\7\"\2\2tu\79\2\2uv\5 \21\2vw\7\66"+
		"\2\2wx\5\36\20\2xy\b\n\1\2y\u0088\3\2\2\2z{\7\13\2\2{|\7\"\2\2|}\7\66"+
		"\2\2}~\5\36\20\2~\177\b\n\1\2\177\u0088\3\2\2\2\u0080\u0081\7\13\2\2\u0081"+
		"\u0082\7\"\2\2\u0082\u0083\79\2\2\u0083\u0084\5 \21\2\u0084\u0085\7,\2"+
		"\2\u0085\u0086\b\n\1\2\u0086\u0088\3\2\2\2\u0087r\3\2\2\2\u0087z\3\2\2"+
		"\2\u0087\u0080\3\2\2\2\u0088\23\3\2\2\2\u0089\u008a\7\r\2\2\u008a\u008b"+
		"\7;\2\2\u008b\u008c\5\36\20\2\u008c\u008d\7<\2\2\u008d\u008e\b\13\1\2"+
		"\u008e\25\3\2\2\2\u008f\u0090\7\16\2\2\u0090\u0091\5\36\20\2\u0091\u0092"+
		"\7=\2\2\u0092\u0093\5\4\3\2\u0093\u0094\7>\2\2\u0094\u0095\b\f\1\2\u0095"+
		"\u00aa\3\2\2\2\u0096\u0097\7\16\2\2\u0097\u0098\5\36\20\2\u0098\u0099"+
		"\7=\2\2\u0099\u009a\5\4\3\2\u009a\u009b\7>\2\2\u009b\u009c\7\17\2\2\u009c"+
		"\u009d\7=\2\2\u009d\u009e\5\4\3\2\u009e\u009f\7>\2\2\u009f\u00a0\b\f\1"+
		"\2\u00a0\u00aa\3\2\2\2\u00a1\u00a2\7\16\2\2\u00a2\u00a3\5\36\20\2\u00a3"+
		"\u00a4\7=\2\2\u00a4\u00a5\5\4\3\2\u00a5\u00a6\7>\2\2\u00a6\u00a7\5\30"+
		"\r\2\u00a7\u00a8\b\f\1\2\u00a8\u00aa\3\2\2\2\u00a9\u008f\3\2\2\2\u00a9"+
		"\u0096\3\2\2\2\u00a9\u00a1\3\2\2\2\u00aa\27\3\2\2\2\u00ab\u00ac\7\17\2"+
		"\2\u00ac\u00ad\7\16\2\2\u00ad\u00ae\5\36\20\2\u00ae\u00af\7=\2\2\u00af"+
		"\u00b0\5\4\3\2\u00b0\u00b1\7>\2\2\u00b1\u00b2\b\r\1\2\u00b2\u00c9\3\2"+
		"\2\2\u00b3\u00b4\7\17\2\2\u00b4\u00b5\7\16\2\2\u00b5\u00b6\5\36\20\2\u00b6"+
		"\u00b7\7=\2\2\u00b7\u00b8\5\4\3\2\u00b8\u00b9\7>\2\2\u00b9\u00ba\7\17"+
		"\2\2\u00ba\u00bb\7=\2\2\u00bb\u00bc\5\4\3\2\u00bc\u00bd\7>\2\2\u00bd\u00be"+
		"\b\r\1\2\u00be\u00c9\3\2\2\2\u00bf\u00c0\7\17\2\2\u00c0\u00c1\7\16\2\2"+
		"\u00c1\u00c2\5\36\20\2\u00c2\u00c3\7=\2\2\u00c3\u00c4\5\4\3\2\u00c4\u00c5"+
		"\7>\2\2\u00c5\u00c6\5\30\r\2\u00c6\u00c7\b\r\1\2\u00c7\u00c9\3\2\2\2\u00c8"+
		"\u00ab\3\2\2\2\u00c8\u00b3\3\2\2\2\u00c8\u00bf\3\2\2\2\u00c9\31\3\2\2"+
		"\2\u00ca\u00cb\7\24\2\2\u00cb\u00cc\5\36\20\2\u00cc\u00cd\7=\2\2\u00cd"+
		"\u00ce\5\4\3\2\u00ce\u00cf\7>\2\2\u00cf\u00d0\b\16\1\2\u00d0\33\3\2\2"+
		"\2\u00d1\u00d2\7\25\2\2\u00d2\u00d3\7\"\2\2\u00d3\u00d4\7\26\2\2\u00d4"+
		"\u00d5\5\36\20\2\u00d5\u00d6\7&\2\2\u00d6\u00d7\5\36\20\2\u00d7\u00d8"+
		"\7=\2\2\u00d8\u00d9\5\4\3\2\u00d9\u00da\7>\2\2\u00da\u00db\b\17\1\2\u00db"+
		"\35\3\2\2\2\u00dc\u00dd\b\20\1\2\u00dd\u00de\7/\2\2\u00de\u00df\5\36\20"+
		"\f\u00df\u00e0\b\20\1\2\u00e0\u00f6\3\2\2\2\u00e1\u00e2\7;\2\2\u00e2\u00e3"+
		"\5\36\20\2\u00e3\u00e4\7<\2\2\u00e4\u00f6\3\2\2\2\u00e5\u00e6\7\37\2\2"+
		"\u00e6\u00f6\b\20\1\2\u00e7\u00e8\7 \2\2\u00e8\u00f6\b\20\1\2\u00e9\u00ea"+
		"\7!\2\2\u00ea\u00f6\b\20\1\2\u00eb\u00ec\7#\2\2\u00ec\u00f6\b\20\1\2\u00ed"+
		"\u00ee\7\b\2\2\u00ee\u00f6\b\20\1\2\u00ef\u00f0\7\t\2\2\u00f0\u00f6\b"+
		"\20\1\2\u00f1\u00f2\7\"\2\2\u00f2\u00f6\b\20\1\2\u00f3\u00f4\7\n\2\2\u00f4"+
		"\u00f6\b\20\1\2\u00f5\u00dc\3\2\2\2\u00f5\u00e1\3\2\2\2\u00f5\u00e5\3"+
		"\2\2\2\u00f5\u00e7\3\2\2\2\u00f5\u00e9\3\2\2\2\u00f5\u00eb\3\2\2\2\u00f5"+
		"\u00ed\3\2\2\2\u00f5\u00ef\3\2\2\2\u00f5\u00f1\3\2\2\2\u00f5\u00f3\3\2"+
		"\2\2\u00f6\u011c\3\2\2\2\u00f7\u00f8\f\23\2\2\u00f8\u00f9\t\2\2\2\u00f9"+
		"\u00fa\5\36\20\24\u00fa\u00fb\b\20\1\2\u00fb\u011b\3\2\2\2\u00fc\u00fd"+
		"\f\22\2\2\u00fd\u00fe\t\3\2\2\u00fe\u00ff\5\36\20\23\u00ff\u0100\b\20"+
		"\1\2\u0100\u011b\3\2\2\2\u0101\u0102\f\21\2\2\u0102\u0103\7+\2\2\u0103"+
		"\u0104\5\36\20\22\u0104\u0105\b\20\1\2\u0105\u011b\3\2\2\2\u0106\u0107"+
		"\f\20\2\2\u0107\u0108\t\4\2\2\u0108\u0109\5\36\20\21\u0109\u010a\b\20"+
		"\1\2\u010a\u011b\3\2\2\2\u010b\u010c\f\17\2\2\u010c\u010d\t\5\2\2\u010d"+
		"\u010e\5\36\20\20\u010e\u010f\b\20\1\2\u010f\u011b\3\2\2\2\u0110\u0111"+
		"\f\16\2\2\u0111\u0112\t\6\2\2\u0112\u0113\5\36\20\17\u0113\u0114\b\20"+
		"\1\2\u0114\u011b\3\2\2\2\u0115\u0116\f\r\2\2\u0116\u0117\t\7\2\2\u0117"+
		"\u0118\5\36\20\16\u0118\u0119\b\20\1\2\u0119\u011b\3\2\2\2\u011a\u00f7"+
		"\3\2\2\2\u011a\u00fc\3\2\2\2\u011a\u0101\3\2\2\2\u011a\u0106\3\2\2\2\u011a"+
		"\u010b\3\2\2\2\u011a\u0110\3\2\2\2\u011a\u0115\3\2\2\2\u011b\u011e\3\2"+
		"\2\2\u011c\u011a\3\2\2\2\u011c\u011d\3\2\2\2\u011d\37\3\2\2\2\u011e\u011c"+
		"\3\2\2\2\u011f\u0120\7\3\2\2\u0120\u012a\b\21\1\2\u0121\u0122\7\4\2\2"+
		"\u0122\u012a\b\21\1\2\u0123\u0124\7!\2\2\u0124\u012a\b\21\1\2\u0125\u0126"+
		"\7\6\2\2\u0126\u012a\b\21\1\2\u0127\u0128\7#\2\2\u0128\u012a\b\21\1\2"+
		"\u0129\u011f\3\2\2\2\u0129\u0121\3\2\2\2\u0129\u0123\3\2\2\2\u0129\u0125"+
		"\3\2\2\2\u0129\u0127\3\2\2\2\u012a!\3\2\2\2\r)K`p\u0087\u00a9\u00c8\u00f5"+
		"\u011a\u011c\u0129";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}