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
		STRING_LITERAL=31, ID=32, CHARACTER_LITERAL=33, SUMMATION=34, SUBTRACTION=35, 
		MULTIPLICATION=36, DIVISION=37, MOD=38, QUESTION_MARK=39, OR=40, AND=41, 
		NOT=42, EQUAL=43, NOT_EQUAL=44, LESS_THAN=45, LESS_THAN_EQUAL=46, GREATER_THAN=47, 
		GREATER_THAN_EQUAL=48, ASSIGN=49, DOT=50, COMMA=51, COLON=52, SEMICOLON=53, 
		OPEN_PARENTHESIS=54, CLOSE_PARENTHESIS=55, OPEN_kEY=56, CLOSE_kEY=57, 
		WHITESPACE=58, COMMENT=59, LINE_COMMENT=60;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_sentence = 2, RULE_declare_var = 3, RULE_print_bl = 4, 
		RULE_expression = 5, RULE_datatype = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "sentence", "declare_var", "print_bl", "expression", "datatype"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'break'", "'default'", "'while'", "'for'", "'in'", "'guard'", 
			"'continue'", "'return'", "'func'", "'struct'", "'mutating'", "'self'", 
			"'inout'", null, null, null, null, null, "'+'", "'-'", "'*'", "'/'", 
			"'%'", "'?'", "'||'", "'&&'", "'!'", "'=='", "'!='", "'<'", "'<='", "'>'", 
			"'>='", "'='", "'.'", "','", "':'", "';'", "'('", "')'", "'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE", 
			"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK", 
			"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC", 
			"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "FLOATT", "STRING_LITERAL", 
			"ID", "CHARACTER_LITERAL", "SUMMATION", "SUBTRACTION", "MULTIPLICATION", 
			"DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", 
			"LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL", 
			"ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", 
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
			setState(14);
			((SContext)_localctx).block = block();
			setState(15);
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
			setState(19); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(18);
				((BlockContext)_localctx).sentence = sentence();
				((BlockContext)_localctx).instr.add(((BlockContext)_localctx).sentence);
				}
				}
				setState(21); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==VAR || _la==PRINT );

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
		public Print_blContext print_bl;
		public Declare_varContext declare_var() {
			return getRuleContext(Declare_varContext.class,0);
		}
		public Print_blContext print_bl() {
			return getRuleContext(Print_blContext.class,0);
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
			setState(31);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(25);
				((SentenceContext)_localctx).declare_var = declare_var();
				_localctx.instr = ((SentenceContext)_localctx).declare_var.instr
				}
				break;
			case PRINT:
				enterOuterAlt(_localctx, 2);
				{
				setState(28);
				((SentenceContext)_localctx).print_bl = print_bl();
				_localctx.instr = ((SentenceContext)_localctx).print_bl.instr
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
		enterRule(_localctx, 6, RULE_declare_var);
		try {
			setState(51);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(33);
				match(VAR);
				setState(34);
				((Declare_varContext)_localctx).ID = match(ID);
				setState(35);
				match(COLON);
				setState(36);
				((Declare_varContext)_localctx).datatype = datatype();
				setState(37);
				match(ASSIGN);
				setState(38);
				((Declare_varContext)_localctx).expression = expression(0);

				                        fmt.Println("Variable declaration: ");
				                        _localctx.instr = instructions.NewDeclareWithValue((((Declare_varContext)_localctx).ID!=null?((Declare_varContext)_localctx).ID.getText():null),((Declare_varContext)_localctx).datatype.td,((Declare_varContext)_localctx).expression.p)
				                
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(41);
				match(VAR);
				setState(42);
				match(ID);
				setState(43);
				match(ASSIGN);
				setState(44);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(45);
				match(VAR);
				setState(46);
				match(ID);
				setState(47);
				match(COLON);
				setState(48);
				datatype();
				setState(49);
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
		enterRule(_localctx, 8, RULE_print_bl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			match(PRINT);
			setState(54);
			match(OPEN_PARENTHESIS);
			setState(55);
			((Print_blContext)_localctx).expression = expression(0);
			setState(56);
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

	public static class ExpressionContext extends ParserRuleContext {
		public abstract.Expression p;
		public ExpressionContext left;
		public Token NUMBER;
		public Token STRING_LITERAL;
		public Token CHARACTER_LITERAL;
		public Token TRUE;
		public Token FALSE;
		public Token ID;
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
		public TerminalNode TRUE() { return getToken(SwiftgrammParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(SwiftgrammParser.FALSE, 0); }
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
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OPEN_PARENTHESIS:
				{
				setState(60);
				match(OPEN_PARENTHESIS);
				setState(61);
				expression(0);
				setState(62);
				match(CLOSE_PARENTHESIS);
				}
				break;
			case NUMBER:
				{
				setState(64);
				((ExpressionContext)_localctx).NUMBER = match(NUMBER);

				                value,err := strconv.Atoi((((ExpressionContext)_localctx).NUMBER!=null?((ExpressionContext)_localctx).NUMBER.getText():null))
				                if err != nil{
				                        fmt.Println(err)
				                }
				                _localctx.p = expressions.NewNative(value,symbol.INT)
				        
				}
				break;
			case STRING_LITERAL:
				{
				setState(66);
				((ExpressionContext)_localctx).STRING_LITERAL = match(STRING_LITERAL);

				                value := (((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getText():null)[1:len((((ExpressionContext)_localctx).STRING_LITERAL!=null?((ExpressionContext)_localctx).STRING_LITERAL.getText():null))-1]
				                _localctx.p = expressions.NewNative(value,symbol.STRING)
				        
				}
				break;
			case CHARACTER_LITERAL:
				{
				setState(68);
				((ExpressionContext)_localctx).CHARACTER_LITERAL = match(CHARACTER_LITERAL);

				                value := (((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getText():null)[1:len((((ExpressionContext)_localctx).CHARACTER_LITERAL!=null?((ExpressionContext)_localctx).CHARACTER_LITERAL.getText():null))-1]
				                _localctx.p = expressions.NewNative(value,symbol.CHAR)
				        
				}
				break;
			case TRUE:
				{
				setState(70);
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
				setState(72);
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
				setState(74);
				((ExpressionContext)_localctx).ID = match(ID);

				                _localctx.p = expressions.NewIdentifier((((ExpressionContext)_localctx).ID!=null?((ExpressionContext)_localctx).ID.getText():null))
				        
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(95);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(93);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(78);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(79);
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
						setState(80);
						((ExpressionContext)_localctx).right = expression(13);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(81);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(82);
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
						setState(83);
						((ExpressionContext)_localctx).right = expression(12);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(84);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(85);
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
						setState(86);
						((ExpressionContext)_localctx).right = expression(11);
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(87);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(88);
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
						setState(89);
						((ExpressionContext)_localctx).right = expression(10);
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(90);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(91);
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
						setState(92);
						((ExpressionContext)_localctx).right = expression(9);
						}
						break;
					}
					} 
				}
				setState(97);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
		public TerminalNode FLOATT() { return getToken(SwiftgrammParser.FLOATT, 0); }
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
		enterRule(_localctx, 12, RULE_datatype);
		try {
			setState(108);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(98);
				match(INT);

				                _localctx.td = symbol.INT
				        
				}
				break;
			case FLOATT:
				enterOuterAlt(_localctx, 2);
				{
				setState(100);
				match(FLOATT);

				                _localctx.td = symbol.FLOAT
				        
				}
				break;
			case STRING_LITERAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(102);
				match(STRING_LITERAL);

				                _localctx.td = symbol.STRING
				        
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(104);
				match(BOOL);

				                _localctx.td = symbol.BOOL
				        
				}
				break;
			case CHARACTER_LITERAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(106);
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
		case 5:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 12);
		case 1:
			return precpred(_ctx, 11);
		case 2:
			return precpred(_ctx, 10);
		case 3:
			return precpred(_ctx, 9);
		case 4:
			return precpred(_ctx, 8);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3>q\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\3\2\3\2\3\3\6\3\26\n"+
		"\3\r\3\16\3\27\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\5\4\"\n\4\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5\66\n\5"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\5\7O\n\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\7\7`\n\7\f\7\16\7c\13\7\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\5\bo\n\b\3\b\2\3\f\t\2\4\6\b\n\f\16\2\7\3\2&\'\3\2$%"+
		"\3\2/\60\3\2\61\62\3\2-.\2|\2\20\3\2\2\2\4\25\3\2\2\2\6!\3\2\2\2\b\65"+
		"\3\2\2\2\n\67\3\2\2\2\fN\3\2\2\2\16n\3\2\2\2\20\21\5\4\3\2\21\22\7\2\2"+
		"\3\22\23\b\2\1\2\23\3\3\2\2\2\24\26\5\6\4\2\25\24\3\2\2\2\26\27\3\2\2"+
		"\2\27\25\3\2\2\2\27\30\3\2\2\2\30\31\3\2\2\2\31\32\b\3\1\2\32\5\3\2\2"+
		"\2\33\34\5\b\5\2\34\35\b\4\1\2\35\"\3\2\2\2\36\37\5\n\6\2\37 \b\4\1\2"+
		" \"\3\2\2\2!\33\3\2\2\2!\36\3\2\2\2\"\7\3\2\2\2#$\7\13\2\2$%\7\"\2\2%"+
		"&\7\66\2\2&\'\5\16\b\2\'(\7\63\2\2()\5\f\7\2)*\b\5\1\2*\66\3\2\2\2+,\7"+
		"\13\2\2,-\7\"\2\2-.\7\63\2\2.\66\5\f\7\2/\60\7\13\2\2\60\61\7\"\2\2\61"+
		"\62\7\66\2\2\62\63\5\16\b\2\63\64\7)\2\2\64\66\3\2\2\2\65#\3\2\2\2\65"+
		"+\3\2\2\2\65/\3\2\2\2\66\t\3\2\2\2\678\7\r\2\289\78\2\29:\5\f\7\2:;\7"+
		"9\2\2;<\b\6\1\2<\13\3\2\2\2=>\b\7\1\2>?\78\2\2?@\5\f\7\2@A\79\2\2AO\3"+
		"\2\2\2BC\7\37\2\2CO\b\7\1\2DE\7!\2\2EO\b\7\1\2FG\7#\2\2GO\b\7\1\2HI\7"+
		"\b\2\2IO\b\7\1\2JK\7\t\2\2KO\b\7\1\2LM\7\"\2\2MO\b\7\1\2N=\3\2\2\2NB\3"+
		"\2\2\2ND\3\2\2\2NF\3\2\2\2NH\3\2\2\2NJ\3\2\2\2NL\3\2\2\2Oa\3\2\2\2PQ\f"+
		"\16\2\2QR\t\2\2\2R`\5\f\7\17ST\f\r\2\2TU\t\3\2\2U`\5\f\7\16VW\f\f\2\2"+
		"WX\t\4\2\2X`\5\f\7\rYZ\f\13\2\2Z[\t\5\2\2[`\5\f\7\f\\]\f\n\2\2]^\t\6\2"+
		"\2^`\5\f\7\13_P\3\2\2\2_S\3\2\2\2_V\3\2\2\2_Y\3\2\2\2_\\\3\2\2\2`c\3\2"+
		"\2\2a_\3\2\2\2ab\3\2\2\2b\r\3\2\2\2ca\3\2\2\2de\7\3\2\2eo\b\b\1\2fg\7"+
		" \2\2go\b\b\1\2hi\7!\2\2io\b\b\1\2jk\7\6\2\2ko\b\b\1\2lm\7#\2\2mo\b\b"+
		"\1\2nd\3\2\2\2nf\3\2\2\2nh\3\2\2\2nj\3\2\2\2nl\3\2\2\2o\17\3\2\2\2\t\27"+
		"!\65N_an";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}