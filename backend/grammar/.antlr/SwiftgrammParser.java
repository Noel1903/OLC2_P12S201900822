// Generated from e:\USAC2023\SegundoSemestre\Compiladores2\OLC2_P12S201900822\backend\grammar\Swiftgramm.g4 by ANTLR 4.9.2

        import "fmt"

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
		FUNC=24, STRUCT=25, MUTATING=26, SELF=27, INOUT=28, NUMBER=29, STRING_LITERAL=30, 
		ID=31, CHARACTER_LITERAL=32, SUMMATION=33, SUBTRACTION=34, MULTIPLICATION=35, 
		DIVISION=36, MOD=37, QUESTION_MARK=38, OR=39, AND=40, NOT=41, EQUAL=42, 
		NOT_EQUAL=43, LESS_THAN=44, LESS_THAN_EQUAL=45, GREATER_THAN=46, GREATER_THAN_EQUAL=47, 
		ASSIGN=48, DOT=49, COMMA=50, COLON=51, SEMICOLON=52, OPEN_PARENTHESIS=53, 
		CLOSE_PARENTHESIS=54, OPEN_kEY=55, CLOSE_kEY=56, WHITESPACE=57, COMMENT=58, 
		LINE_COMMENT=59;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_sentences = 2, RULE_sentence = 3, RULE_print = 4, 
		RULE_declare_var = 5, RULE_declare_constant = 6, RULE_assign_var = 7, 
		RULE_if_sentence = 8, RULE_switch_sentence = 9, RULE_switch_cases = 10, 
		RULE_switch_case = 11, RULE_while_sentence = 12, RULE_for_sentence = 13, 
		RULE_guard_sentence = 14, RULE_break_sentence = 15, RULE_continue_sentence = 16, 
		RULE_return_sentence = 17, RULE_expression = 18, RULE_datatype = 19;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "sentences", "sentence", "print", "declare_var", "declare_constant", 
			"assign_var", "if_sentence", "switch_sentence", "switch_cases", "switch_case", 
			"while_sentence", "for_sentence", "guard_sentence", "break_sentence", 
			"continue_sentence", "return_sentence", "expression", "datatype"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'break'", "'default'", "'while'", "'for'", "'in'", "'guard'", 
			"'continue'", "'return'", "'func'", "'struct'", "'mutating'", "'self'", 
			"'inout'", null, null, null, null, "'+'", "'-'", "'*'", "'/'", "'%'", 
			"'?'", "'||'", "'&&'", "'!'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", 
			"'='", "'.'", "','", "':'", "';'", "'('", "')'", "'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE", 
			"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK", 
			"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC", 
			"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "STRING_LITERAL", "ID", 
			"CHARACTER_LITERAL", "SUMMATION", "SUBTRACTION", "MULTIPLICATION", "DIVISION", 
			"MOD", "QUESTION_MARK", "OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", "LESS_THAN", 
			"LESS_THAN_EQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL", "ASSIGN", "DOT", 
			"COMMA", "COLON", "SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", 
			"OPEN_kEY", "CLOSE_kEY", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
			setState(40);
			block();
			setState(41);
			match(EOF);
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
		public List<SentencesContext> sentences() {
			return getRuleContexts(SentencesContext.class);
		}
		public SentencesContext sentences(int i) {
			return getRuleContext(SentencesContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << ID))) != 0)) {
				{
				{
				setState(43);
				sentences();
				}
				}
				setState(48);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	public static class SentencesContext extends ParserRuleContext {
		public SentenceContext sentence() {
			return getRuleContext(SentenceContext.class,0);
		}
		public SentencesContext sentences() {
			return getRuleContext(SentencesContext.class,0);
		}
		public SentencesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentences; }
	}

	public final SentencesContext sentences() throws RecognitionException {
		SentencesContext _localctx = new SentencesContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentences);
		try {
			setState(53);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(49);
				sentence();
				setState(50);
				sentences();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(52);
				sentence();
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

	public static class SentenceContext extends ParserRuleContext {
		public Declare_varContext declare_var() {
			return getRuleContext(Declare_varContext.class,0);
		}
		public Declare_constantContext declare_constant() {
			return getRuleContext(Declare_constantContext.class,0);
		}
		public Assign_varContext assign_var() {
			return getRuleContext(Assign_varContext.class,0);
		}
		public SentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentence; }
	}

	public final SentenceContext sentence() throws RecognitionException {
		SentenceContext _localctx = new SentenceContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_sentence);
		try {
			setState(58);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(55);
				declare_var();
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(56);
				declare_constant();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(57);
				assign_var();
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

	public static class PrintContext extends ParserRuleContext {
		public TerminalNode PRINT() { return getToken(SwiftgrammParser.PRINT, 0); }
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_print);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			match(PRINT);
			setState(61);
			match(OPEN_PARENTHESIS);
			setState(62);
			expression(0);
			setState(63);
			match(CLOSE_PARENTHESIS);
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
		enterRule(_localctx, 10, RULE_declare_var);
		try {
			setState(83);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				match(VAR);
				setState(66);
				match(ID);
				setState(67);
				match(COLON);
				setState(68);
				datatype();
				setState(69);
				match(ASSIGN);
				setState(70);
				expression(0);
				fmt.Println("Variable declaration: ");
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(73);
				match(VAR);
				setState(74);
				match(ID);
				setState(75);
				match(ASSIGN);
				setState(76);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(77);
				match(VAR);
				setState(78);
				match(ID);
				setState(79);
				match(COLON);
				setState(80);
				datatype();
				setState(81);
				match(QUESTION_MARK);
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

	public static class Declare_constantContext extends ParserRuleContext {
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
		public TerminalNode QUESTION_MARK() { return getToken(SwiftgrammParser.QUESTION_MARK, 0); }
		public Declare_constantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declare_constant; }
	}

	public final Declare_constantContext declare_constant() throws RecognitionException {
		Declare_constantContext _localctx = new Declare_constantContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_declare_constant);
		try {
			setState(102);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				match(LET);
				setState(86);
				match(ID);
				setState(87);
				match(COLON);
				setState(88);
				datatype();
				setState(89);
				match(ASSIGN);
				setState(90);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(92);
				match(LET);
				setState(93);
				match(ID);
				setState(94);
				match(ASSIGN);
				setState(95);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(96);
				match(LET);
				setState(97);
				match(ID);
				setState(98);
				match(COLON);
				setState(99);
				datatype();
				setState(100);
				match(QUESTION_MARK);
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

	public static class Assign_varContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(SwiftgrammParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Assign_varContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_var; }
	}

	public final Assign_varContext assign_var() throws RecognitionException {
		Assign_varContext _localctx = new Assign_varContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_assign_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(104);
			match(ID);
			setState(105);
			match(ASSIGN);
			setState(106);
			expression(0);
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

	public static class If_sentenceContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(SwiftgrammParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<SentencesContext> sentences() {
			return getRuleContexts(SentencesContext.class);
		}
		public SentencesContext sentences(int i) {
			return getRuleContext(SentencesContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(SwiftgrammParser.ELSE, 0); }
		public If_sentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_sentence; }
	}

	public final If_sentenceContext if_sentence() throws RecognitionException {
		If_sentenceContext _localctx = new If_sentenceContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_if_sentence);
		try {
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(108);
				match(IF);
				setState(109);
				expression(0);
				setState(110);
				sentences();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				match(IF);
				setState(113);
				expression(0);
				setState(114);
				sentences();
				setState(115);
				match(ELSE);
				setState(116);
				sentences();
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

	public static class Switch_sentenceContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(SwiftgrammParser.SWITCH, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode OPEN_kEY() { return getToken(SwiftgrammParser.OPEN_kEY, 0); }
		public Switch_casesContext switch_cases() {
			return getRuleContext(Switch_casesContext.class,0);
		}
		public TerminalNode CLOSE_kEY() { return getToken(SwiftgrammParser.CLOSE_kEY, 0); }
		public Switch_sentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_sentence; }
	}

	public final Switch_sentenceContext switch_sentence() throws RecognitionException {
		Switch_sentenceContext _localctx = new Switch_sentenceContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_switch_sentence);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(120);
			match(SWITCH);
			setState(121);
			expression(0);
			setState(122);
			match(OPEN_kEY);
			setState(123);
			switch_cases();
			setState(124);
			match(CLOSE_kEY);
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

	public static class Switch_casesContext extends ParserRuleContext {
		public Switch_caseContext switch_case() {
			return getRuleContext(Switch_caseContext.class,0);
		}
		public Switch_casesContext switch_cases() {
			return getRuleContext(Switch_casesContext.class,0);
		}
		public Switch_casesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_cases; }
	}

	public final Switch_casesContext switch_cases() throws RecognitionException {
		Switch_casesContext _localctx = new Switch_casesContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_switch_cases);
		try {
			setState(130);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(126);
				switch_case();
				setState(127);
				switch_cases();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(129);
				switch_case();
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

	public static class Switch_caseContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(SwiftgrammParser.CASE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COLON() { return getToken(SwiftgrammParser.COLON, 0); }
		public SentencesContext sentences() {
			return getRuleContext(SentencesContext.class,0);
		}
		public TerminalNode DEFAULT() { return getToken(SwiftgrammParser.DEFAULT, 0); }
		public Switch_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_case; }
	}

	public final Switch_caseContext switch_case() throws RecognitionException {
		Switch_caseContext _localctx = new Switch_caseContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_switch_case);
		try {
			setState(140);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CASE:
				enterOuterAlt(_localctx, 1);
				{
				setState(132);
				match(CASE);
				setState(133);
				expression(0);
				setState(134);
				match(COLON);
				setState(135);
				sentences();
				}
				break;
			case DEFAULT:
				enterOuterAlt(_localctx, 2);
				{
				setState(137);
				match(DEFAULT);
				setState(138);
				match(COLON);
				setState(139);
				sentences();
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

	public static class While_sentenceContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(SwiftgrammParser.WHILE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode OPEN_kEY() { return getToken(SwiftgrammParser.OPEN_kEY, 0); }
		public SentencesContext sentences() {
			return getRuleContext(SentencesContext.class,0);
		}
		public TerminalNode CLOSE_kEY() { return getToken(SwiftgrammParser.CLOSE_kEY, 0); }
		public While_sentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_sentence; }
	}

	public final While_sentenceContext while_sentence() throws RecognitionException {
		While_sentenceContext _localctx = new While_sentenceContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_while_sentence);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
			match(WHILE);
			setState(143);
			expression(0);
			setState(144);
			match(OPEN_kEY);
			setState(145);
			sentences();
			setState(146);
			match(CLOSE_kEY);
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

	public static class For_sentenceContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(SwiftgrammParser.FOR, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode IN() { return getToken(SwiftgrammParser.IN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode OPEN_kEY() { return getToken(SwiftgrammParser.OPEN_kEY, 0); }
		public SentencesContext sentences() {
			return getRuleContext(SentencesContext.class,0);
		}
		public TerminalNode CLOSE_kEY() { return getToken(SwiftgrammParser.CLOSE_kEY, 0); }
		public For_sentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_sentence; }
	}

	public final For_sentenceContext for_sentence() throws RecognitionException {
		For_sentenceContext _localctx = new For_sentenceContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_for_sentence);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(FOR);
			setState(149);
			match(ID);
			setState(150);
			match(IN);
			setState(151);
			expression(0);
			setState(152);
			match(OPEN_kEY);
			setState(153);
			sentences();
			setState(154);
			match(CLOSE_kEY);
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

	public static class Guard_sentenceContext extends ParserRuleContext {
		public TerminalNode GUARD() { return getToken(SwiftgrammParser.GUARD, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftgrammParser.ELSE, 0); }
		public SentencesContext sentences() {
			return getRuleContext(SentencesContext.class,0);
		}
		public Guard_sentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guard_sentence; }
	}

	public final Guard_sentenceContext guard_sentence() throws RecognitionException {
		Guard_sentenceContext _localctx = new Guard_sentenceContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_guard_sentence);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(156);
			match(GUARD);
			setState(157);
			expression(0);
			setState(158);
			match(ELSE);
			setState(159);
			sentences();
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

	public static class Break_sentenceContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(SwiftgrammParser.BREAK, 0); }
		public Break_sentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break_sentence; }
	}

	public final Break_sentenceContext break_sentence() throws RecognitionException {
		Break_sentenceContext _localctx = new Break_sentenceContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_break_sentence);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			match(BREAK);
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

	public static class Continue_sentenceContext extends ParserRuleContext {
		public TerminalNode CONTINUE() { return getToken(SwiftgrammParser.CONTINUE, 0); }
		public Continue_sentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue_sentence; }
	}

	public final Continue_sentenceContext continue_sentence() throws RecognitionException {
		Continue_sentenceContext _localctx = new Continue_sentenceContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_continue_sentence);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(163);
			match(CONTINUE);
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

	public static class Return_sentenceContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(SwiftgrammParser.RETURN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Return_sentenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_sentence; }
	}

	public final Return_sentenceContext return_sentence() throws RecognitionException {
		Return_sentenceContext _localctx = new Return_sentenceContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_return_sentence);
		try {
			setState(168);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(165);
				match(RETURN);
				setState(166);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(167);
				match(RETURN);
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

	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext left;
		public Token oper;
		public ExpressionContext right;
		public TerminalNode OPEN_PARENTHESIS() { return getToken(SwiftgrammParser.OPEN_PARENTHESIS, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CLOSE_PARENTHESIS() { return getToken(SwiftgrammParser.CLOSE_PARENTHESIS, 0); }
		public TerminalNode NUMBER() { return getToken(SwiftgrammParser.NUMBER, 0); }
		public TerminalNode STRING_LITERAL() { return getToken(SwiftgrammParser.STRING_LITERAL, 0); }
		public TerminalNode CHARACTER_LITERAL() { return getToken(SwiftgrammParser.CHARACTER_LITERAL, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public TerminalNode MULTIPLICATION() { return getToken(SwiftgrammParser.MULTIPLICATION, 0); }
		public TerminalNode DIVISION() { return getToken(SwiftgrammParser.DIVISION, 0); }
		public TerminalNode SUMMATION() { return getToken(SwiftgrammParser.SUMMATION, 0); }
		public TerminalNode SUBTRACTION() { return getToken(SwiftgrammParser.SUBTRACTION, 0); }
		public TerminalNode LESS_THAN() { return getToken(SwiftgrammParser.LESS_THAN, 0); }
		public TerminalNode LESS_THAN_EQUAL() { return getToken(SwiftgrammParser.LESS_THAN_EQUAL, 0); }
		public TerminalNode GREATER_THAN() { return getToken(SwiftgrammParser.GREATER_THAN, 0); }
		public TerminalNode GREATER_THAN_EQUAL() { return getToken(SwiftgrammParser.GREATER_THAN_EQUAL, 0); }
		public TerminalNode EQUAL() { return getToken(SwiftgrammParser.EQUAL, 0); }
		public TerminalNode NOT_EQUAL() { return getToken(SwiftgrammParser.NOT_EQUAL, 0); }
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
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OPEN_PARENTHESIS:
				{
				setState(171);
				match(OPEN_PARENTHESIS);
				setState(172);
				expression(0);
				setState(173);
				match(CLOSE_PARENTHESIS);
				}
				break;
			case NUMBER:
				{
				setState(175);
				match(NUMBER);
				}
				break;
			case STRING_LITERAL:
				{
				setState(176);
				match(STRING_LITERAL);
				}
				break;
			case CHARACTER_LITERAL:
				{
				setState(177);
				match(CHARACTER_LITERAL);
				}
				break;
			case ID:
				{
				setState(178);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(198);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(196);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(181);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(182);
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
						setState(183);
						((ExpressionContext)_localctx).right = expression(11);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(184);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(185);
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
						setState(186);
						((ExpressionContext)_localctx).right = expression(10);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(187);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(188);
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
						setState(189);
						((ExpressionContext)_localctx).right = expression(9);
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(190);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(191);
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
						setState(192);
						((ExpressionContext)_localctx).right = expression(8);
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(193);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(194);
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
						setState(195);
						((ExpressionContext)_localctx).right = expression(7);
						}
						break;
					}
					} 
				}
				setState(200);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
		public TerminalNode INT() { return getToken(SwiftgrammParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftgrammParser.FLOAT, 0); }
		public TerminalNode STRING() { return getToken(SwiftgrammParser.STRING, 0); }
		public TerminalNode BOOL() { return getToken(SwiftgrammParser.BOOL, 0); }
		public DatatypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_datatype; }
	}

	public final DatatypeContext datatype() throws RecognitionException {
		DatatypeContext _localctx = new DatatypeContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_datatype);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STRING) | (1L << BOOL))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 18:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 10);
		case 1:
			return precpred(_ctx, 9);
		case 2:
			return precpred(_ctx, 8);
		case 3:
			return precpred(_ctx, 7);
		case 4:
			return precpred(_ctx, 6);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3=\u00ce\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\2\3\3\7\3/\n\3\f\3\16\3\62\13"+
		"\3\3\4\3\4\3\4\3\4\5\48\n\4\3\5\3\5\3\5\5\5=\n\5\3\6\3\6\3\6\3\6\3\6\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\5\7V\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\5\bi\n\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\5\ny\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\5\f\u0085"+
		"\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u008f\n\r\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20"+
		"\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\23\5\23\u00ab\n\23\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u00b6\n\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\7\24\u00c7\n\24\f\24"+
		"\16\24\u00ca\13\24\3\25\3\25\3\25\2\3&\26\2\4\6\b\n\f\16\20\22\24\26\30"+
		"\32\34\36 \"$&(\2\b\3\2%&\3\2#$\3\2./\3\2\60\61\3\2,-\3\2\3\6\2\u00ce"+
		"\2*\3\2\2\2\4\60\3\2\2\2\6\67\3\2\2\2\b<\3\2\2\2\n>\3\2\2\2\fU\3\2\2\2"+
		"\16h\3\2\2\2\20j\3\2\2\2\22x\3\2\2\2\24z\3\2\2\2\26\u0084\3\2\2\2\30\u008e"+
		"\3\2\2\2\32\u0090\3\2\2\2\34\u0096\3\2\2\2\36\u009e\3\2\2\2 \u00a3\3\2"+
		"\2\2\"\u00a5\3\2\2\2$\u00aa\3\2\2\2&\u00b5\3\2\2\2(\u00cb\3\2\2\2*+\5"+
		"\4\3\2+,\7\2\2\3,\3\3\2\2\2-/\5\6\4\2.-\3\2\2\2/\62\3\2\2\2\60.\3\2\2"+
		"\2\60\61\3\2\2\2\61\5\3\2\2\2\62\60\3\2\2\2\63\64\5\b\5\2\64\65\5\6\4"+
		"\2\658\3\2\2\2\668\5\b\5\2\67\63\3\2\2\2\67\66\3\2\2\28\7\3\2\2\29=\5"+
		"\f\7\2:=\5\16\b\2;=\5\20\t\2<9\3\2\2\2<:\3\2\2\2<;\3\2\2\2=\t\3\2\2\2"+
		">?\7\r\2\2?@\7\67\2\2@A\5&\24\2AB\78\2\2B\13\3\2\2\2CD\7\13\2\2DE\7!\2"+
		"\2EF\7\65\2\2FG\5(\25\2GH\7\62\2\2HI\5&\24\2IJ\b\7\1\2JV\3\2\2\2KL\7\13"+
		"\2\2LM\7!\2\2MN\7\62\2\2NV\5&\24\2OP\7\13\2\2PQ\7!\2\2QR\7\65\2\2RS\5"+
		"(\25\2ST\7(\2\2TV\3\2\2\2UC\3\2\2\2UK\3\2\2\2UO\3\2\2\2V\r\3\2\2\2WX\7"+
		"\f\2\2XY\7!\2\2YZ\7\65\2\2Z[\5(\25\2[\\\7\62\2\2\\]\5&\24\2]i\3\2\2\2"+
		"^_\7\f\2\2_`\7!\2\2`a\7\62\2\2ai\5&\24\2bc\7\f\2\2cd\7!\2\2de\7\65\2\2"+
		"ef\5(\25\2fg\7(\2\2gi\3\2\2\2hW\3\2\2\2h^\3\2\2\2hb\3\2\2\2i\17\3\2\2"+
		"\2jk\7!\2\2kl\7\62\2\2lm\5&\24\2m\21\3\2\2\2no\7\16\2\2op\5&\24\2pq\5"+
		"\6\4\2qy\3\2\2\2rs\7\16\2\2st\5&\24\2tu\5\6\4\2uv\7\17\2\2vw\5\6\4\2w"+
		"y\3\2\2\2xn\3\2\2\2xr\3\2\2\2y\23\3\2\2\2z{\7\20\2\2{|\5&\24\2|}\79\2"+
		"\2}~\5\26\f\2~\177\7:\2\2\177\25\3\2\2\2\u0080\u0081\5\30\r\2\u0081\u0082"+
		"\5\26\f\2\u0082\u0085\3\2\2\2\u0083\u0085\5\30\r\2\u0084\u0080\3\2\2\2"+
		"\u0084\u0083\3\2\2\2\u0085\27\3\2\2\2\u0086\u0087\7\21\2\2\u0087\u0088"+
		"\5&\24\2\u0088\u0089\7\65\2\2\u0089\u008a\5\6\4\2\u008a\u008f\3\2\2\2"+
		"\u008b\u008c\7\23\2\2\u008c\u008d\7\65\2\2\u008d\u008f\5\6\4\2\u008e\u0086"+
		"\3\2\2\2\u008e\u008b\3\2\2\2\u008f\31\3\2\2\2\u0090\u0091\7\24\2\2\u0091"+
		"\u0092\5&\24\2\u0092\u0093\79\2\2\u0093\u0094\5\6\4\2\u0094\u0095\7:\2"+
		"\2\u0095\33\3\2\2\2\u0096\u0097\7\25\2\2\u0097\u0098\7!\2\2\u0098\u0099"+
		"\7\26\2\2\u0099\u009a\5&\24\2\u009a\u009b\79\2\2\u009b\u009c\5\6\4\2\u009c"+
		"\u009d\7:\2\2\u009d\35\3\2\2\2\u009e\u009f\7\27\2\2\u009f\u00a0\5&\24"+
		"\2\u00a0\u00a1\7\17\2\2\u00a1\u00a2\5\6\4\2\u00a2\37\3\2\2\2\u00a3\u00a4"+
		"\7\22\2\2\u00a4!\3\2\2\2\u00a5\u00a6\7\30\2\2\u00a6#\3\2\2\2\u00a7\u00a8"+
		"\7\31\2\2\u00a8\u00ab\5&\24\2\u00a9\u00ab\7\31\2\2\u00aa\u00a7\3\2\2\2"+
		"\u00aa\u00a9\3\2\2\2\u00ab%\3\2\2\2\u00ac\u00ad\b\24\1\2\u00ad\u00ae\7"+
		"\67\2\2\u00ae\u00af\5&\24\2\u00af\u00b0\78\2\2\u00b0\u00b6\3\2\2\2\u00b1"+
		"\u00b6\7\37\2\2\u00b2\u00b6\7 \2\2\u00b3\u00b6\7\"\2\2\u00b4\u00b6\7!"+
		"\2\2\u00b5\u00ac\3\2\2\2\u00b5\u00b1\3\2\2\2\u00b5\u00b2\3\2\2\2\u00b5"+
		"\u00b3\3\2\2\2\u00b5\u00b4\3\2\2\2\u00b6\u00c8\3\2\2\2\u00b7\u00b8\f\f"+
		"\2\2\u00b8\u00b9\t\2\2\2\u00b9\u00c7\5&\24\r\u00ba\u00bb\f\13\2\2\u00bb"+
		"\u00bc\t\3\2\2\u00bc\u00c7\5&\24\f\u00bd\u00be\f\n\2\2\u00be\u00bf\t\4"+
		"\2\2\u00bf\u00c7\5&\24\13\u00c0\u00c1\f\t\2\2\u00c1\u00c2\t\5\2\2\u00c2"+
		"\u00c7\5&\24\n\u00c3\u00c4\f\b\2\2\u00c4\u00c5\t\6\2\2\u00c5\u00c7\5&"+
		"\24\t\u00c6\u00b7\3\2\2\2\u00c6\u00ba\3\2\2\2\u00c6\u00bd\3\2\2\2\u00c6"+
		"\u00c0\3\2\2\2\u00c6\u00c3\3\2\2\2\u00c7\u00ca\3\2\2\2\u00c8\u00c6\3\2"+
		"\2\2\u00c8\u00c9\3\2\2\2\u00c9\'\3\2\2\2\u00ca\u00c8\3\2\2\2\u00cb\u00cc"+
		"\t\7\2\2\u00cc)\3\2\2\2\16\60\67<Uhx\u0084\u008e\u00aa\u00b5\u00c6\u00c8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}