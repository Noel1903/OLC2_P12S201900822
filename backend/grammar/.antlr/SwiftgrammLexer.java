// Generated from e:\USAC2023\SegundoSemestre\Compiladores2\P1\OLC2_P12S201900822\backend\grammar\Swiftgramm.g4 by ANTLR 4.9.2

        import "grammar/expressions"
        import "grammar/instructions"
        import "grammar/symbol"
        import "grammar/abstract"

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
		FUNC=24, STRUCT=25, MUTATING=26, SELF=27, INOUT=28, APPEND=29, REMOVELAST=30, 
		REMOVE=31, AT=32, ISEMPTY=33, COUNT=34, NUMBER=35, FLOATT=36, ID=37, CHARACTER_LITERAL=38, 
		STRING_LITERAL=39, INCREMENT=40, DECREMENT=41, RANGE=42, SUMMATION=43, 
		SUBTRACTION=44, MULTIPLICATION=45, DIVISION=46, MOD=47, QUESTION_MARK=48, 
		OR=49, AND=50, NOT=51, EQUAL=52, NOT_EQUAL=53, LESS_THAN=54, LESS_THAN_EQUAL=55, 
		GREATER_THAN=56, GREATER_THAN_EQUAL=57, ASSIGN=58, DOT=59, COMMA=60, COLON=61, 
		SEMICOLON=62, OPEN_PARENTHESIS=63, CLOSE_PARENTHESIS=64, OPEN_kEY=65, 
		CLOSE_kEY=66, OPEN_BRACKET=67, CLOSE_BRACKET=68, ARROW=69, UNDERSCORE=70, 
		WHITESPACE=71, COMMENT=72, LINE_COMMENT=73;
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
			"MUTATING", "SELF", "INOUT", "APPEND", "REMOVELAST", "REMOVE", "AT", 
			"ISEMPTY", "COUNT", "NUMBER", "FLOATT", "ID", "CHARACTER_LITERAL", "STRING_LITERAL", 
			"INCREMENT", "DECREMENT", "RANGE", "SUMMATION", "SUBTRACTION", "MULTIPLICATION", 
			"DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", 
			"LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL", 
			"ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", 
			"OPEN_kEY", "CLOSE_kEY", "OPEN_BRACKET", "CLOSE_BRACKET", "ARROW", "UNDERSCORE", 
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


	        var CountM int = 0 


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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2K\u01f5\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3"+
		"\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\16"+
		"\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25"+
		"\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\36"+
		"\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\""+
		"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\6$\u0165\n$\r$\16$\u0166\3%\6%\u016a\n%"+
		"\r%\16%\u016b\3%\3%\6%\u0170\n%\r%\16%\u0171\5%\u0174\n%\3&\3&\7&\u0178"+
		"\n&\f&\16&\u017b\13&\3\'\3\'\3\'\3\'\3(\3(\7(\u0183\n(\f(\16(\u0186\13"+
		"(\3(\3(\3)\3)\3)\3*\3*\3*\3+\3+\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60"+
		"\3\61\3\61\3\62\3\62\3\62\3\63\3\63\3\63\3\64\3\64\3\65\3\65\3\65\3\66"+
		"\3\66\3\66\3\67\3\67\38\38\38\39\39\3:\3:\3:\3;\3;\3<\3<\3=\3=\3>\3>\3"+
		"?\3?\3@\3@\3A\3A\3B\3B\3C\3C\3D\3D\3E\3E\3F\3F\3F\3G\3G\3H\6H\u01d4\n"+
		"H\rH\16H\u01d5\3H\3H\3I\3I\3I\3I\7I\u01de\nI\fI\16I\u01e1\13I\3I\3I\3"+
		"I\3I\3I\3J\3J\3J\3J\7J\u01ec\nJ\fJ\16J\u01ef\13J\3J\3J\3K\3K\3K\3\u01df"+
		"\2L\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35"+
		"\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36"+
		";\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67"+
		"m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008d"+
		"H\u008fI\u0091J\u0093K\u0095\2\3\2\t\3\2\62;\5\2C\\aac|\6\2\62;C\\aac"+
		"|\3\2$$\5\2\13\f\17\17\"\"\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u01fc"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3"+
		"\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2"+
		"\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2"+
		"{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085"+
		"\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2"+
		"\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2\2\3\u0097\3\2\2\2\5\u009b"+
		"\3\2\2\2\7\u00a1\3\2\2\2\t\u00a8\3\2\2\2\13\u00ad\3\2\2\2\r\u00b7\3\2"+
		"\2\2\17\u00bc\3\2\2\2\21\u00c2\3\2\2\2\23\u00c6\3\2\2\2\25\u00ca\3\2\2"+
		"\2\27\u00ce\3\2\2\2\31\u00d4\3\2\2\2\33\u00d7\3\2\2\2\35\u00dc\3\2\2\2"+
		"\37\u00e3\3\2\2\2!\u00e8\3\2\2\2#\u00ee\3\2\2\2%\u00f6\3\2\2\2\'\u00fc"+
		"\3\2\2\2)\u0100\3\2\2\2+\u0103\3\2\2\2-\u0109\3\2\2\2/\u0112\3\2\2\2\61"+
		"\u0119\3\2\2\2\63\u011e\3\2\2\2\65\u0125\3\2\2\2\67\u012e\3\2\2\29\u0133"+
		"\3\2\2\2;\u0139\3\2\2\2=\u0140\3\2\2\2?\u014b\3\2\2\2A\u0152\3\2\2\2C"+
		"\u0155\3\2\2\2E\u015d\3\2\2\2G\u0164\3\2\2\2I\u0169\3\2\2\2K\u0175\3\2"+
		"\2\2M\u017c\3\2\2\2O\u0180\3\2\2\2Q\u0189\3\2\2\2S\u018c\3\2\2\2U\u018f"+
		"\3\2\2\2W\u0193\3\2\2\2Y\u0195\3\2\2\2[\u0197\3\2\2\2]\u0199\3\2\2\2_"+
		"\u019b\3\2\2\2a\u019d\3\2\2\2c\u019f\3\2\2\2e\u01a2\3\2\2\2g\u01a5\3\2"+
		"\2\2i\u01a7\3\2\2\2k\u01aa\3\2\2\2m\u01ad\3\2\2\2o\u01af\3\2\2\2q\u01b2"+
		"\3\2\2\2s\u01b4\3\2\2\2u\u01b7\3\2\2\2w\u01b9\3\2\2\2y\u01bb\3\2\2\2{"+
		"\u01bd\3\2\2\2}\u01bf\3\2\2\2\177\u01c1\3\2\2\2\u0081\u01c3\3\2\2\2\u0083"+
		"\u01c5\3\2\2\2\u0085\u01c7\3\2\2\2\u0087\u01c9\3\2\2\2\u0089\u01cb\3\2"+
		"\2\2\u008b\u01cd\3\2\2\2\u008d\u01d0\3\2\2\2\u008f\u01d3\3\2\2\2\u0091"+
		"\u01d9\3\2\2\2\u0093\u01e7\3\2\2\2\u0095\u01f2\3\2\2\2\u0097\u0098\7K"+
		"\2\2\u0098\u0099\7p\2\2\u0099\u009a\7v\2\2\u009a\4\3\2\2\2\u009b\u009c"+
		"\7H\2\2\u009c\u009d\7n\2\2\u009d\u009e\7q\2\2\u009e\u009f\7c\2\2\u009f"+
		"\u00a0\7v\2\2\u00a0\6\3\2\2\2\u00a1\u00a2\7U\2\2\u00a2\u00a3\7v\2\2\u00a3"+
		"\u00a4\7t\2\2\u00a4\u00a5\7k\2\2\u00a5\u00a6\7p\2\2\u00a6\u00a7\7i\2\2"+
		"\u00a7\b\3\2\2\2\u00a8\u00a9\7D\2\2\u00a9\u00aa\7q\2\2\u00aa\u00ab\7q"+
		"\2\2\u00ab\u00ac\7n\2\2\u00ac\n\3\2\2\2\u00ad\u00ae\7E\2\2\u00ae\u00af"+
		"\7j\2\2\u00af\u00b0\7c\2\2\u00b0\u00b1\7t\2\2\u00b1\u00b2\7c\2\2\u00b2"+
		"\u00b3\7e\2\2\u00b3\u00b4\7v\2\2\u00b4\u00b5\7g\2\2\u00b5\u00b6\7t\2\2"+
		"\u00b6\f\3\2\2\2\u00b7\u00b8\7v\2\2\u00b8\u00b9\7t\2\2\u00b9\u00ba\7w"+
		"\2\2\u00ba\u00bb\7g\2\2\u00bb\16\3\2\2\2\u00bc\u00bd\7h\2\2\u00bd\u00be"+
		"\7c\2\2\u00be\u00bf\7n\2\2\u00bf\u00c0\7u\2\2\u00c0\u00c1\7g\2\2\u00c1"+
		"\20\3\2\2\2\u00c2\u00c3\7p\2\2\u00c3\u00c4\7k\2\2\u00c4\u00c5\7n\2\2\u00c5"+
		"\22\3\2\2\2\u00c6\u00c7\7x\2\2\u00c7\u00c8\7c\2\2\u00c8\u00c9\7t\2\2\u00c9"+
		"\24\3\2\2\2\u00ca\u00cb\7n\2\2\u00cb\u00cc\7g\2\2\u00cc\u00cd\7v\2\2\u00cd"+
		"\26\3\2\2\2\u00ce\u00cf\7r\2\2\u00cf\u00d0\7t\2\2\u00d0\u00d1\7k\2\2\u00d1"+
		"\u00d2\7p\2\2\u00d2\u00d3\7v\2\2\u00d3\30\3\2\2\2\u00d4\u00d5\7k\2\2\u00d5"+
		"\u00d6\7h\2\2\u00d6\32\3\2\2\2\u00d7\u00d8\7g\2\2\u00d8\u00d9\7n\2\2\u00d9"+
		"\u00da\7u\2\2\u00da\u00db\7g\2\2\u00db\34\3\2\2\2\u00dc\u00dd\7u\2\2\u00dd"+
		"\u00de\7y\2\2\u00de\u00df\7k\2\2\u00df\u00e0\7v\2\2\u00e0\u00e1\7e\2\2"+
		"\u00e1\u00e2\7j\2\2\u00e2\36\3\2\2\2\u00e3\u00e4\7e\2\2\u00e4\u00e5\7"+
		"c\2\2\u00e5\u00e6\7u\2\2\u00e6\u00e7\7g\2\2\u00e7 \3\2\2\2\u00e8\u00e9"+
		"\7d\2\2\u00e9\u00ea\7t\2\2\u00ea\u00eb\7g\2\2\u00eb\u00ec\7c\2\2\u00ec"+
		"\u00ed\7m\2\2\u00ed\"\3\2\2\2\u00ee\u00ef\7f\2\2\u00ef\u00f0\7g\2\2\u00f0"+
		"\u00f1\7h\2\2\u00f1\u00f2\7c\2\2\u00f2\u00f3\7w\2\2\u00f3\u00f4\7n\2\2"+
		"\u00f4\u00f5\7v\2\2\u00f5$\3\2\2\2\u00f6\u00f7\7y\2\2\u00f7\u00f8\7j\2"+
		"\2\u00f8\u00f9\7k\2\2\u00f9\u00fa\7n\2\2\u00fa\u00fb\7g\2\2\u00fb&\3\2"+
		"\2\2\u00fc\u00fd\7h\2\2\u00fd\u00fe\7q\2\2\u00fe\u00ff\7t\2\2\u00ff(\3"+
		"\2\2\2\u0100\u0101\7k\2\2\u0101\u0102\7p\2\2\u0102*\3\2\2\2\u0103\u0104"+
		"\7i\2\2\u0104\u0105\7w\2\2\u0105\u0106\7c\2\2\u0106\u0107\7t\2\2\u0107"+
		"\u0108\7f\2\2\u0108,\3\2\2\2\u0109\u010a\7e\2\2\u010a\u010b\7q\2\2\u010b"+
		"\u010c\7p\2\2\u010c\u010d\7v\2\2\u010d\u010e\7k\2\2\u010e\u010f\7p\2\2"+
		"\u010f\u0110\7w\2\2\u0110\u0111\7g\2\2\u0111.\3\2\2\2\u0112\u0113\7t\2"+
		"\2\u0113\u0114\7g\2\2\u0114\u0115\7v\2\2\u0115\u0116\7w\2\2\u0116\u0117"+
		"\7t\2\2\u0117\u0118\7p\2\2\u0118\60\3\2\2\2\u0119\u011a\7h\2\2\u011a\u011b"+
		"\7w\2\2\u011b\u011c\7p\2\2\u011c\u011d\7e\2\2\u011d\62\3\2\2\2\u011e\u011f"+
		"\7u\2\2\u011f\u0120\7v\2\2\u0120\u0121\7t\2\2\u0121\u0122\7w\2\2\u0122"+
		"\u0123\7e\2\2\u0123\u0124\7v\2\2\u0124\64\3\2\2\2\u0125\u0126\7o\2\2\u0126"+
		"\u0127\7w\2\2\u0127\u0128\7v\2\2\u0128\u0129\7c\2\2\u0129\u012a\7v\2\2"+
		"\u012a\u012b\7k\2\2\u012b\u012c\7p\2\2\u012c\u012d\7i\2\2\u012d\66\3\2"+
		"\2\2\u012e\u012f\7u\2\2\u012f\u0130\7g\2\2\u0130\u0131\7n\2\2\u0131\u0132"+
		"\7h\2\2\u01328\3\2\2\2\u0133\u0134\7k\2\2\u0134\u0135\7p\2\2\u0135\u0136"+
		"\7q\2\2\u0136\u0137\7w\2\2\u0137\u0138\7v\2\2\u0138:\3\2\2\2\u0139\u013a"+
		"\7c\2\2\u013a\u013b\7r\2\2\u013b\u013c\7r\2\2\u013c\u013d\7g\2\2\u013d"+
		"\u013e\7p\2\2\u013e\u013f\7f\2\2\u013f<\3\2\2\2\u0140\u0141\7t\2\2\u0141"+
		"\u0142\7g\2\2\u0142\u0143\7o\2\2\u0143\u0144\7q\2\2\u0144\u0145\7x\2\2"+
		"\u0145\u0146\7g\2\2\u0146\u0147\7N\2\2\u0147\u0148\7c\2\2\u0148\u0149"+
		"\7u\2\2\u0149\u014a\7v\2\2\u014a>\3\2\2\2\u014b\u014c\7t\2\2\u014c\u014d"+
		"\7g\2\2\u014d\u014e\7o\2\2\u014e\u014f\7q\2\2\u014f\u0150\7x\2\2\u0150"+
		"\u0151\7g\2\2\u0151@\3\2\2\2\u0152\u0153\7c\2\2\u0153\u0154\7v\2\2\u0154"+
		"B\3\2\2\2\u0155\u0156\7K\2\2\u0156\u0157\7u\2\2\u0157\u0158\7G\2\2\u0158"+
		"\u0159\7o\2\2\u0159\u015a\7r\2\2\u015a\u015b\7v\2\2\u015b\u015c\7{\2\2"+
		"\u015cD\3\2\2\2\u015d\u015e\7e\2\2\u015e\u015f\7q\2\2\u015f\u0160\7w\2"+
		"\2\u0160\u0161\7p\2\2\u0161\u0162\7v\2\2\u0162F\3\2\2\2\u0163\u0165\t"+
		"\2\2\2\u0164\u0163\3\2\2\2\u0165\u0166\3\2\2\2\u0166\u0164\3\2\2\2\u0166"+
		"\u0167\3\2\2\2\u0167H\3\2\2\2\u0168\u016a\t\2\2\2\u0169\u0168\3\2\2\2"+
		"\u016a\u016b\3\2\2\2\u016b\u0169\3\2\2\2\u016b\u016c\3\2\2\2\u016c\u0173"+
		"\3\2\2\2\u016d\u016f\7\60\2\2\u016e\u0170\t\2\2\2\u016f\u016e\3\2\2\2"+
		"\u0170\u0171\3\2\2\2\u0171\u016f\3\2\2\2\u0171\u0172\3\2\2\2\u0172\u0174"+
		"\3\2\2\2\u0173\u016d\3\2\2\2\u0173\u0174\3\2\2\2\u0174J\3\2\2\2\u0175"+
		"\u0179\t\3\2\2\u0176\u0178\t\4\2\2\u0177\u0176\3\2\2\2\u0178\u017b\3\2"+
		"\2\2\u0179\u0177\3\2\2\2\u0179\u017a\3\2\2\2\u017aL\3\2\2\2\u017b\u0179"+
		"\3\2\2\2\u017c\u017d\7$\2\2\u017d\u017e\n\5\2\2\u017e\u017f\7$\2\2\u017f"+
		"N\3\2\2\2\u0180\u0184\7$\2\2\u0181\u0183\n\5\2\2\u0182\u0181\3\2\2\2\u0183"+
		"\u0186\3\2\2\2\u0184\u0182\3\2\2\2\u0184\u0185\3\2\2\2\u0185\u0187\3\2"+
		"\2\2\u0186\u0184\3\2\2\2\u0187\u0188\7$\2\2\u0188P\3\2\2\2\u0189\u018a"+
		"\7-\2\2\u018a\u018b\7?\2\2\u018bR\3\2\2\2\u018c\u018d\7/\2\2\u018d\u018e"+
		"\7?\2\2\u018eT\3\2\2\2\u018f\u0190\7\60\2\2\u0190\u0191\7\60\2\2\u0191"+
		"\u0192\7\60\2\2\u0192V\3\2\2\2\u0193\u0194\7-\2\2\u0194X\3\2\2\2\u0195"+
		"\u0196\7/\2\2\u0196Z\3\2\2\2\u0197\u0198\7,\2\2\u0198\\\3\2\2\2\u0199"+
		"\u019a\7\61\2\2\u019a^\3\2\2\2\u019b\u019c\7\'\2\2\u019c`\3\2\2\2\u019d"+
		"\u019e\7A\2\2\u019eb\3\2\2\2\u019f\u01a0\7~\2\2\u01a0\u01a1\7~\2\2\u01a1"+
		"d\3\2\2\2\u01a2\u01a3\7(\2\2\u01a3\u01a4\7(\2\2\u01a4f\3\2\2\2\u01a5\u01a6"+
		"\7#\2\2\u01a6h\3\2\2\2\u01a7\u01a8\7?\2\2\u01a8\u01a9\7?\2\2\u01a9j\3"+
		"\2\2\2\u01aa\u01ab\7#\2\2\u01ab\u01ac\7?\2\2\u01acl\3\2\2\2\u01ad\u01ae"+
		"\7>\2\2\u01aen\3\2\2\2\u01af\u01b0\7>\2\2\u01b0\u01b1\7?\2\2\u01b1p\3"+
		"\2\2\2\u01b2\u01b3\7@\2\2\u01b3r\3\2\2\2\u01b4\u01b5\7@\2\2\u01b5\u01b6"+
		"\7?\2\2\u01b6t\3\2\2\2\u01b7\u01b8\7?\2\2\u01b8v\3\2\2\2\u01b9\u01ba\7"+
		"\60\2\2\u01bax\3\2\2\2\u01bb\u01bc\7.\2\2\u01bcz\3\2\2\2\u01bd\u01be\7"+
		"<\2\2\u01be|\3\2\2\2\u01bf\u01c0\7=\2\2\u01c0~\3\2\2\2\u01c1\u01c2\7*"+
		"\2\2\u01c2\u0080\3\2\2\2\u01c3\u01c4\7+\2\2\u01c4\u0082\3\2\2\2\u01c5"+
		"\u01c6\7}\2\2\u01c6\u0084\3\2\2\2\u01c7\u01c8\7\177\2\2\u01c8\u0086\3"+
		"\2\2\2\u01c9\u01ca\7]\2\2\u01ca\u0088\3\2\2\2\u01cb\u01cc\7_\2\2\u01cc"+
		"\u008a\3\2\2\2\u01cd\u01ce\7/\2\2\u01ce\u01cf\7@\2\2\u01cf\u008c\3\2\2"+
		"\2\u01d0\u01d1\7a\2\2\u01d1\u008e\3\2\2\2\u01d2\u01d4\t\6\2\2\u01d3\u01d2"+
		"\3\2\2\2\u01d4\u01d5\3\2\2\2\u01d5\u01d3\3\2\2\2\u01d5\u01d6\3\2\2\2\u01d6"+
		"\u01d7\3\2\2\2\u01d7\u01d8\bH\2\2\u01d8\u0090\3\2\2\2\u01d9\u01da\7\61"+
		"\2\2\u01da\u01db\7,\2\2\u01db\u01df\3\2\2\2\u01dc\u01de\13\2\2\2\u01dd"+
		"\u01dc\3\2\2\2\u01de\u01e1\3\2\2\2\u01df\u01e0\3\2\2\2\u01df\u01dd\3\2"+
		"\2\2\u01e0\u01e2\3\2\2\2\u01e1\u01df\3\2\2\2\u01e2\u01e3\7,\2\2\u01e3"+
		"\u01e4\7\61\2\2\u01e4\u01e5\3\2\2\2\u01e5\u01e6\bI\2\2\u01e6\u0092\3\2"+
		"\2\2\u01e7\u01e8\7\61\2\2\u01e8\u01e9\7\61\2\2\u01e9\u01ed\3\2\2\2\u01ea"+
		"\u01ec\n\7\2\2\u01eb\u01ea\3\2\2\2\u01ec\u01ef\3\2\2\2\u01ed\u01eb\3\2"+
		"\2\2\u01ed\u01ee\3\2\2\2\u01ee\u01f0\3\2\2\2\u01ef\u01ed\3\2\2\2\u01f0"+
		"\u01f1\bJ\2\2\u01f1\u0094\3\2\2\2\u01f2\u01f3\7^\2\2\u01f3\u01f4\t\b\2"+
		"\2\u01f4\u0096\3\2\2\2\f\2\u0166\u016b\u0171\u0173\u0179\u0184\u01d5\u01df"+
		"\u01ed\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}