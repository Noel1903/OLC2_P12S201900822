// Generated from e:\USAC2023\SegundoSemestre\Compiladores2\OLC2_P12S201900822\backend\grammar\Swiftgramm.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftgrammLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE", "NIL", 
			"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK", "DEFAULT", 
			"WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC", "STRUCT", 
			"MUTATING", "SELF", "INOUT", "NUMBER", "STRING_LITERAL", "ID", "CHARACTER_LITERAL", 
			"SUMMATION", "SUBTRACTION", "MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK", 
			"OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", 
			"GREATER_THAN", "GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON", 
			"SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", 
			"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
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


	public SwiftgrammLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Swiftgramm.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2=\u0197\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7"+
		"\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13"+
		"\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16\3\16"+
		"\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34"+
		"\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\36\6\36\u011f\n\36"+
		"\r\36\16\36\u0120\3\36\3\36\6\36\u0125\n\36\r\36\16\36\u0126\5\36\u0129"+
		"\n\36\3\37\3\37\7\37\u012d\n\37\f\37\16\37\u0130\13\37\3\37\3\37\3 \3"+
		" \7 \u0136\n \f \16 \u0139\13 \3!\3!\3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3"+
		"&\3&\3\'\3\'\3(\3(\3(\3)\3)\3)\3*\3*\3+\3+\3+\3,\3,\3,\3-\3-\3.\3.\3."+
		"\3/\3/\3\60\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65"+
		"\3\66\3\66\3\67\3\67\38\38\39\39\3:\6:\u0176\n:\r:\16:\u0177\3:\3:\3;"+
		"\3;\3;\3;\7;\u0180\n;\f;\16;\u0183\13;\3;\3;\3;\3;\3;\3<\3<\3<\3<\7<\u018e"+
		"\n<\f<\16<\u0191\13<\3<\3<\3=\3=\3=\3\u0181\2>\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O"+
		")Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y\2\3\2\n\3\2"+
		"\62;\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\3\2))\5\2\13\f\17\17\"\"\4\2\f\f"+
		"\17\17\7\2%%*-/\60AB]_\2\u019d\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t"+
		"\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2"+
		"\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2"+
		"\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2"+
		"+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2"+
		"\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2"+
		"C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3"+
		"\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2"+
		"\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2"+
		"i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3"+
		"\2\2\2\2w\3\2\2\2\3{\3\2\2\2\5\177\3\2\2\2\7\u0085\3\2\2\2\t\u008c\3\2"+
		"\2\2\13\u0091\3\2\2\2\r\u009b\3\2\2\2\17\u00a0\3\2\2\2\21\u00a6\3\2\2"+
		"\2\23\u00aa\3\2\2\2\25\u00ae\3\2\2\2\27\u00b2\3\2\2\2\31\u00b8\3\2\2\2"+
		"\33\u00bb\3\2\2\2\35\u00c0\3\2\2\2\37\u00c7\3\2\2\2!\u00cc\3\2\2\2#\u00d2"+
		"\3\2\2\2%\u00da\3\2\2\2\'\u00e0\3\2\2\2)\u00e4\3\2\2\2+\u00e7\3\2\2\2"+
		"-\u00ed\3\2\2\2/\u00f6\3\2\2\2\61\u00fd\3\2\2\2\63\u0102\3\2\2\2\65\u0109"+
		"\3\2\2\2\67\u0112\3\2\2\29\u0117\3\2\2\2;\u011e\3\2\2\2=\u012a\3\2\2\2"+
		"?\u0133\3\2\2\2A\u013a\3\2\2\2C\u013e\3\2\2\2E\u0140\3\2\2\2G\u0142\3"+
		"\2\2\2I\u0144\3\2\2\2K\u0146\3\2\2\2M\u0148\3\2\2\2O\u014a\3\2\2\2Q\u014d"+
		"\3\2\2\2S\u0150\3\2\2\2U\u0152\3\2\2\2W\u0155\3\2\2\2Y\u0158\3\2\2\2["+
		"\u015a\3\2\2\2]\u015d\3\2\2\2_\u015f\3\2\2\2a\u0162\3\2\2\2c\u0164\3\2"+
		"\2\2e\u0166\3\2\2\2g\u0168\3\2\2\2i\u016a\3\2\2\2k\u016c\3\2\2\2m\u016e"+
		"\3\2\2\2o\u0170\3\2\2\2q\u0172\3\2\2\2s\u0175\3\2\2\2u\u017b\3\2\2\2w"+
		"\u0189\3\2\2\2y\u0194\3\2\2\2{|\7K\2\2|}\7p\2\2}~\7v\2\2~\4\3\2\2\2\177"+
		"\u0080\7H\2\2\u0080\u0081\7n\2\2\u0081\u0082\7q\2\2\u0082\u0083\7c\2\2"+
		"\u0083\u0084\7v\2\2\u0084\6\3\2\2\2\u0085\u0086\7U\2\2\u0086\u0087\7v"+
		"\2\2\u0087\u0088\7t\2\2\u0088\u0089\7k\2\2\u0089\u008a\7p\2\2\u008a\u008b"+
		"\7i\2\2\u008b\b\3\2\2\2\u008c\u008d\7D\2\2\u008d\u008e\7q\2\2\u008e\u008f"+
		"\7q\2\2\u008f\u0090\7n\2\2\u0090\n\3\2\2\2\u0091\u0092\7E\2\2\u0092\u0093"+
		"\7j\2\2\u0093\u0094\7c\2\2\u0094\u0095\7t\2\2\u0095\u0096\7c\2\2\u0096"+
		"\u0097\7e\2\2\u0097\u0098\7v\2\2\u0098\u0099\7g\2\2\u0099\u009a\7t\2\2"+
		"\u009a\f\3\2\2\2\u009b\u009c\7v\2\2\u009c\u009d\7t\2\2\u009d\u009e\7w"+
		"\2\2\u009e\u009f\7g\2\2\u009f\16\3\2\2\2\u00a0\u00a1\7h\2\2\u00a1\u00a2"+
		"\7c\2\2\u00a2\u00a3\7n\2\2\u00a3\u00a4\7u\2\2\u00a4\u00a5\7g\2\2\u00a5"+
		"\20\3\2\2\2\u00a6\u00a7\7p\2\2\u00a7\u00a8\7k\2\2\u00a8\u00a9\7n\2\2\u00a9"+
		"\22\3\2\2\2\u00aa\u00ab\7x\2\2\u00ab\u00ac\7c\2\2\u00ac\u00ad\7t\2\2\u00ad"+
		"\24\3\2\2\2\u00ae\u00af\7n\2\2\u00af\u00b0\7g\2\2\u00b0\u00b1\7v\2\2\u00b1"+
		"\26\3\2\2\2\u00b2\u00b3\7r\2\2\u00b3\u00b4\7t\2\2\u00b4\u00b5\7k\2\2\u00b5"+
		"\u00b6\7p\2\2\u00b6\u00b7\7v\2\2\u00b7\30\3\2\2\2\u00b8\u00b9\7k\2\2\u00b9"+
		"\u00ba\7h\2\2\u00ba\32\3\2\2\2\u00bb\u00bc\7g\2\2\u00bc\u00bd\7n\2\2\u00bd"+
		"\u00be\7u\2\2\u00be\u00bf\7g\2\2\u00bf\34\3\2\2\2\u00c0\u00c1\7u\2\2\u00c1"+
		"\u00c2\7y\2\2\u00c2\u00c3\7k\2\2\u00c3\u00c4\7v\2\2\u00c4\u00c5\7e\2\2"+
		"\u00c5\u00c6\7j\2\2\u00c6\36\3\2\2\2\u00c7\u00c8\7e\2\2\u00c8\u00c9\7"+
		"c\2\2\u00c9\u00ca\7u\2\2\u00ca\u00cb\7g\2\2\u00cb \3\2\2\2\u00cc\u00cd"+
		"\7d\2\2\u00cd\u00ce\7t\2\2\u00ce\u00cf\7g\2\2\u00cf\u00d0\7c\2\2\u00d0"+
		"\u00d1\7m\2\2\u00d1\"\3\2\2\2\u00d2\u00d3\7f\2\2\u00d3\u00d4\7g\2\2\u00d4"+
		"\u00d5\7h\2\2\u00d5\u00d6\7c\2\2\u00d6\u00d7\7w\2\2\u00d7\u00d8\7n\2\2"+
		"\u00d8\u00d9\7v\2\2\u00d9$\3\2\2\2\u00da\u00db\7y\2\2\u00db\u00dc\7j\2"+
		"\2\u00dc\u00dd\7k\2\2\u00dd\u00de\7n\2\2\u00de\u00df\7g\2\2\u00df&\3\2"+
		"\2\2\u00e0\u00e1\7h\2\2\u00e1\u00e2\7q\2\2\u00e2\u00e3\7t\2\2\u00e3(\3"+
		"\2\2\2\u00e4\u00e5\7k\2\2\u00e5\u00e6\7p\2\2\u00e6*\3\2\2\2\u00e7\u00e8"+
		"\7i\2\2\u00e8\u00e9\7w\2\2\u00e9\u00ea\7c\2\2\u00ea\u00eb\7t\2\2\u00eb"+
		"\u00ec\7f\2\2\u00ec,\3\2\2\2\u00ed\u00ee\7e\2\2\u00ee\u00ef\7q\2\2\u00ef"+
		"\u00f0\7p\2\2\u00f0\u00f1\7v\2\2\u00f1\u00f2\7k\2\2\u00f2\u00f3\7p\2\2"+
		"\u00f3\u00f4\7w\2\2\u00f4\u00f5\7g\2\2\u00f5.\3\2\2\2\u00f6\u00f7\7t\2"+
		"\2\u00f7\u00f8\7g\2\2\u00f8\u00f9\7v\2\2\u00f9\u00fa\7w\2\2\u00fa\u00fb"+
		"\7t\2\2\u00fb\u00fc\7p\2\2\u00fc\60\3\2\2\2\u00fd\u00fe\7h\2\2\u00fe\u00ff"+
		"\7w\2\2\u00ff\u0100\7p\2\2\u0100\u0101\7e\2\2\u0101\62\3\2\2\2\u0102\u0103"+
		"\7u\2\2\u0103\u0104\7v\2\2\u0104\u0105\7t\2\2\u0105\u0106\7w\2\2\u0106"+
		"\u0107\7e\2\2\u0107\u0108\7v\2\2\u0108\64\3\2\2\2\u0109\u010a\7o\2\2\u010a"+
		"\u010b\7w\2\2\u010b\u010c\7v\2\2\u010c\u010d\7c\2\2\u010d\u010e\7v\2\2"+
		"\u010e\u010f\7k\2\2\u010f\u0110\7p\2\2\u0110\u0111\7i\2\2\u0111\66\3\2"+
		"\2\2\u0112\u0113\7u\2\2\u0113\u0114\7g\2\2\u0114\u0115\7n\2\2\u0115\u0116"+
		"\7h\2\2\u01168\3\2\2\2\u0117\u0118\7k\2\2\u0118\u0119\7p\2\2\u0119\u011a"+
		"\7q\2\2\u011a\u011b\7w\2\2\u011b\u011c\7v\2\2\u011c:\3\2\2\2\u011d\u011f"+
		"\t\2\2\2\u011e\u011d\3\2\2\2\u011f\u0120\3\2\2\2\u0120\u011e\3\2\2\2\u0120"+
		"\u0121\3\2\2\2\u0121\u0128\3\2\2\2\u0122\u0124\7\60\2\2\u0123\u0125\t"+
		"\2\2\2\u0124\u0123\3\2\2\2\u0125\u0126\3\2\2\2\u0126\u0124\3\2\2\2\u0126"+
		"\u0127\3\2\2\2\u0127\u0129\3\2\2\2\u0128\u0122\3\2\2\2\u0128\u0129\3\2"+
		"\2\2\u0129<\3\2\2\2\u012a\u012e\7$\2\2\u012b\u012d\n\3\2\2\u012c\u012b"+
		"\3\2\2\2\u012d\u0130\3\2\2\2\u012e\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012f"+
		"\u0131\3\2\2\2\u0130\u012e\3\2\2\2\u0131\u0132\7$\2\2\u0132>\3\2\2\2\u0133"+
		"\u0137\t\4\2\2\u0134\u0136\t\5\2\2\u0135\u0134\3\2\2\2\u0136\u0139\3\2"+
		"\2\2\u0137\u0135\3\2\2\2\u0137\u0138\3\2\2\2\u0138@\3\2\2\2\u0139\u0137"+
		"\3\2\2\2\u013a\u013b\7$\2\2\u013b\u013c\n\6\2\2\u013c\u013d\7$\2\2\u013d"+
		"B\3\2\2\2\u013e\u013f\7-\2\2\u013fD\3\2\2\2\u0140\u0141\7/\2\2\u0141F"+
		"\3\2\2\2\u0142\u0143\7,\2\2\u0143H\3\2\2\2\u0144\u0145\7\61\2\2\u0145"+
		"J\3\2\2\2\u0146\u0147\7\'\2\2\u0147L\3\2\2\2\u0148\u0149\7A\2\2\u0149"+
		"N\3\2\2\2\u014a\u014b\7~\2\2\u014b\u014c\7~\2\2\u014cP\3\2\2\2\u014d\u014e"+
		"\7(\2\2\u014e\u014f\7(\2\2\u014fR\3\2\2\2\u0150\u0151\7#\2\2\u0151T\3"+
		"\2\2\2\u0152\u0153\7?\2\2\u0153\u0154\7?\2\2\u0154V\3\2\2\2\u0155\u0156"+
		"\7#\2\2\u0156\u0157\7?\2\2\u0157X\3\2\2\2\u0158\u0159\7>\2\2\u0159Z\3"+
		"\2\2\2\u015a\u015b\7>\2\2\u015b\u015c\7?\2\2\u015c\\\3\2\2\2\u015d\u015e"+
		"\7@\2\2\u015e^\3\2\2\2\u015f\u0160\7@\2\2\u0160\u0161\7?\2\2\u0161`\3"+
		"\2\2\2\u0162\u0163\7?\2\2\u0163b\3\2\2\2\u0164\u0165\7\60\2\2\u0165d\3"+
		"\2\2\2\u0166\u0167\7.\2\2\u0167f\3\2\2\2\u0168\u0169\7<\2\2\u0169h\3\2"+
		"\2\2\u016a\u016b\7=\2\2\u016bj\3\2\2\2\u016c\u016d\7*\2\2\u016dl\3\2\2"+
		"\2\u016e\u016f\7+\2\2\u016fn\3\2\2\2\u0170\u0171\7}\2\2\u0171p\3\2\2\2"+
		"\u0172\u0173\7\177\2\2\u0173r\3\2\2\2\u0174\u0176\t\7\2\2\u0175\u0174"+
		"\3\2\2\2\u0176\u0177\3\2\2\2\u0177\u0175\3\2\2\2\u0177\u0178\3\2\2\2\u0178"+
		"\u0179\3\2\2\2\u0179\u017a\b:\2\2\u017at\3\2\2\2\u017b\u017c\7\61\2\2"+
		"\u017c\u017d\7,\2\2\u017d\u0181\3\2\2\2\u017e\u0180\13\2\2\2\u017f\u017e"+
		"\3\2\2\2\u0180\u0183\3\2\2\2\u0181\u0182\3\2\2\2\u0181\u017f\3\2\2\2\u0182"+
		"\u0184\3\2\2\2\u0183\u0181\3\2\2\2\u0184\u0185\7,\2\2\u0185\u0186\7\61"+
		"\2\2\u0186\u0187\3\2\2\2\u0187\u0188\b;\2\2\u0188v\3\2\2\2\u0189\u018a"+
		"\7\61\2\2\u018a\u018b\7\61\2\2\u018b\u018f\3\2\2\2\u018c\u018e\n\b\2\2"+
		"\u018d\u018c\3\2\2\2\u018e\u0191\3\2\2\2\u018f\u018d\3\2\2\2\u018f\u0190"+
		"\3\2\2\2\u0190\u0192\3\2\2\2\u0191\u018f\3\2\2\2\u0192\u0193\b<\2\2\u0193"+
		"x\3\2\2\2\u0194\u0195\7^\2\2\u0195\u0196\t\t\2\2\u0196z\3\2\2\2\13\2\u0120"+
		"\u0126\u0128\u012e\u0137\u0177\u0181\u018f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}