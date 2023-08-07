// Generated from e:\USAC2023\SegundoSemestre\Compiladores2\OLC2_P12S201900822\backend\grammar\Swiftgramm.g4 by ANTLR 4.9.2
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
		DIVISION=36, QUESTION_MARK=37, OR=38, AND=39, NOT=40, EQUAL=41, NOT_EQUAL=42, 
		LESS_THAN=43, LESS_THAN_EQUAL=44, GREATER_THAN=45, GREATER_THAN_EQUAL=46, 
		ASSIGN=47, DOT=48, COMMA=49, COLON=50, SEMICOLON=51, OPEN_PARENTHESIS=52, 
		CLOSE_PARENTHESIS=53, OPEN_kEY=54, CLOSE_kEY=55, WHITESPACE=56, COMMENT=57, 
		LINE_COMMENT=58;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_sentences = 2, RULE_sentence = 3, RULE_declare_var = 4, 
		RULE_declare_constant = 5, RULE_assign_var = 6, RULE_expression = 7, RULE_datatype = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "sentences", "sentence", "declare_var", "declare_constant", 
			"assign_var", "expression", "datatype"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'break'", "'default'", "'while'", "'for'", "'in'", "'guard'", 
			"'continue'", "'return'", "'func'", "'struct'", "'mutating'", "'self'", 
			"'inout'", null, null, null, null, "'+'", "'-'", "'*'", "'/'", "'?'", 
			"'||'", "'&&'", "'!'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", 
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
			"QUESTION_MARK", "OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", "LESS_THAN", 
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
			setState(18);
			block();
			setState(19);
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
			setState(24);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << ID))) != 0)) {
				{
				{
				setState(21);
				sentences();
				}
				}
				setState(26);
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
			setState(31);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				sentence();
				setState(28);
				sentences();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(30);
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
			setState(36);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(33);
				declare_var();
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				declare_constant();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(35);
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
		enterRule(_localctx, 8, RULE_declare_var);
		try {
			setState(55);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(38);
				match(VAR);
				setState(39);
				match(ID);
				setState(40);
				match(COLON);
				setState(41);
				datatype();
				setState(42);
				match(ASSIGN);
				setState(43);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(45);
				match(VAR);
				setState(46);
				match(ID);
				setState(47);
				match(ASSIGN);
				setState(48);
				expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(49);
				match(VAR);
				setState(50);
				match(ID);
				setState(51);
				match(COLON);
				setState(52);
				datatype();
				setState(53);
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
		enterRule(_localctx, 10, RULE_declare_constant);
		try {
			setState(74);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(57);
				match(LET);
				setState(58);
				match(ID);
				setState(59);
				match(COLON);
				setState(60);
				datatype();
				setState(61);
				match(ASSIGN);
				setState(62);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(64);
				match(LET);
				setState(65);
				match(ID);
				setState(66);
				match(ASSIGN);
				setState(67);
				expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(68);
				match(LET);
				setState(69);
				match(ID);
				setState(70);
				match(COLON);
				setState(71);
				datatype();
				setState(72);
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
		enterRule(_localctx, 12, RULE_assign_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			match(ID);
			setState(77);
			match(ASSIGN);
			setState(78);
			expression();
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
		public TerminalNode NUMBER() { return getToken(SwiftgrammParser.NUMBER, 0); }
		public TerminalNode STRING_LITERAL() { return getToken(SwiftgrammParser.STRING_LITERAL, 0); }
		public TerminalNode CHARACTER_LITERAL() { return getToken(SwiftgrammParser.CHARACTER_LITERAL, 0); }
		public TerminalNode ID() { return getToken(SwiftgrammParser.ID, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NUMBER) | (1L << STRING_LITERAL) | (1L << ID) | (1L << CHARACTER_LITERAL))) != 0)) ) {
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
		enterRule(_localctx, 16, RULE_datatype);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3<W\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3\2\3\2\3"+
		"\3\7\3\31\n\3\f\3\16\3\34\13\3\3\4\3\4\3\4\3\4\5\4\"\n\4\3\5\3\5\3\5\5"+
		"\5\'\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\5\6:\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\5\7M\n\7\3\b\3\b\3\b\3\b\3\t\3\t\3\n\3\n\3\n\2\2\13\2\4"+
		"\6\b\n\f\16\20\22\2\4\3\2\37\"\3\2\3\6\2U\2\24\3\2\2\2\4\32\3\2\2\2\6"+
		"!\3\2\2\2\b&\3\2\2\2\n9\3\2\2\2\fL\3\2\2\2\16N\3\2\2\2\20R\3\2\2\2\22"+
		"T\3\2\2\2\24\25\5\4\3\2\25\26\7\2\2\3\26\3\3\2\2\2\27\31\5\6\4\2\30\27"+
		"\3\2\2\2\31\34\3\2\2\2\32\30\3\2\2\2\32\33\3\2\2\2\33\5\3\2\2\2\34\32"+
		"\3\2\2\2\35\36\5\b\5\2\36\37\5\6\4\2\37\"\3\2\2\2 \"\5\b\5\2!\35\3\2\2"+
		"\2! \3\2\2\2\"\7\3\2\2\2#\'\5\n\6\2$\'\5\f\7\2%\'\5\16\b\2&#\3\2\2\2&"+
		"$\3\2\2\2&%\3\2\2\2\'\t\3\2\2\2()\7\13\2\2)*\7!\2\2*+\7\64\2\2+,\5\22"+
		"\n\2,-\7\61\2\2-.\5\20\t\2.:\3\2\2\2/\60\7\13\2\2\60\61\7!\2\2\61\62\7"+
		"\61\2\2\62:\5\20\t\2\63\64\7\13\2\2\64\65\7!\2\2\65\66\7\64\2\2\66\67"+
		"\5\22\n\2\678\7\'\2\28:\3\2\2\29(\3\2\2\29/\3\2\2\29\63\3\2\2\2:\13\3"+
		"\2\2\2;<\7\f\2\2<=\7!\2\2=>\7\64\2\2>?\5\22\n\2?@\7\61\2\2@A\5\20\t\2"+
		"AM\3\2\2\2BC\7\f\2\2CD\7!\2\2DE\7\61\2\2EM\5\20\t\2FG\7\f\2\2GH\7!\2\2"+
		"HI\7\64\2\2IJ\5\22\n\2JK\7\'\2\2KM\3\2\2\2L;\3\2\2\2LB\3\2\2\2LF\3\2\2"+
		"\2M\r\3\2\2\2NO\7!\2\2OP\7\61\2\2PQ\5\20\t\2Q\17\3\2\2\2RS\t\2\2\2S\21"+
		"\3\2\2\2TU\t\3\2\2U\23\3\2\2\2\7\32!&9L";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}