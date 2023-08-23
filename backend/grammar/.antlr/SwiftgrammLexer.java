// Generated from e:\USAC2023\SegundoSemestre\Compiladores2\OLC2_P12S201900822\backend\grammar\Swiftgramm.g4 by ANTLR 4.9.2

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
		FUNC=24, STRUCT=25, MUTATING=26, SELF=27, INOUT=28, NUMBER=29, FLOATT=30, 
		STRING_LITERAL=31, ID=32, CHARACTER_LITERAL=33, INCREMENT=34, DECREMENT=35, 
		SUMMATION=36, SUBTRACTION=37, MULTIPLICATION=38, DIVISION=39, MOD=40, 
		QUESTION_MARK=41, OR=42, AND=43, NOT=44, EQUAL=45, NOT_EQUAL=46, LESS_THAN=47, 
		LESS_THAN_EQUAL=48, GREATER_THAN=49, GREATER_THAN_EQUAL=50, ASSIGN=51, 
		DOT=52, COMMA=53, COLON=54, SEMICOLON=55, OPEN_PARENTHESIS=56, CLOSE_PARENTHESIS=57, 
		OPEN_kEY=58, CLOSE_kEY=59, WHITESPACE=60, COMMENT=61, LINE_COMMENT=62;
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
			"MUTATING", "SELF", "INOUT", "NUMBER", "FLOATT", "STRING_LITERAL", "ID", 
			"CHARACTER_LITERAL", "INCREMENT", "DECREMENT", "SUMMATION", "SUBTRACTION", 
			"MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT", 
			"EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN", 
			"GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON", 
			"OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'break'", "'default'", "'while'", "'for'", "'in'", "'guard'", 
			"'continue'", "'return'", "'func'", "'struct'", "'mutating'", "'self'", 
			"'inout'", null, null, null, null, null, "'+='", "'-='", "'+'", "'-'", 
			"'*'", "'/'", "'%'", "'?'", "'||'", "'&&'", "'!'", "'=='", "'!='", "'<'", 
			"'<='", "'>'", "'>='", "'='", "'.'", "','", "':'", "';'", "'('", "')'", 
			"'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE", 
			"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK", 
			"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC", 
			"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "FLOATT", "STRING_LITERAL", 
			"ID", "CHARACTER_LITERAL", "INCREMENT", "DECREMENT", "SUMMATION", "SUBTRACTION", 
			"MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT", 
			"EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN", 
			"GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON", 
			"OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2@\u01a8\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n"+
		"\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3"+
		"\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3"+
		"\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3"+
		"\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3"+
		"\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3"+
		"\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3"+
		"\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3"+
		"\36\6\36\u0125\n\36\r\36\16\36\u0126\3\37\6\37\u012a\n\37\r\37\16\37\u012b"+
		"\3\37\3\37\6\37\u0130\n\37\r\37\16\37\u0131\5\37\u0134\n\37\3 \3 \7 \u0138"+
		"\n \f \16 \u013b\13 \3 \3 \3!\3!\7!\u0141\n!\f!\16!\u0144\13!\3\"\3\""+
		"\3\"\3\"\3#\3#\3#\3$\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3"+
		"+\3+\3,\3,\3,\3-\3-\3.\3.\3.\3/\3/\3/\3\60\3\60\3\61\3\61\3\61\3\62\3"+
		"\62\3\63\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\38\38\39\3"+
		"9\3:\3:\3;\3;\3<\3<\3=\6=\u0187\n=\r=\16=\u0188\3=\3=\3>\3>\3>\3>\7>\u0191"+
		"\n>\f>\16>\u0194\13>\3>\3>\3>\3>\3>\3?\3?\3?\3?\7?\u019f\n?\f?\16?\u01a2"+
		"\13?\3?\3?\3@\3@\3@\3\u0192\2A\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61"+
		"a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177\2\3\2\n\3\2\62;\3\2$$\5"+
		"\2C\\aac|\6\2\62;C\\aac|\3\2))\5\2\13\f\17\17\"\"\4\2\f\f\17\17\t\2\""+
		"#%%--/\60<<BB]_\2\u01af\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2"+
		"\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25"+
		"\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2"+
		"\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2"+
		"\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3"+
		"\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2"+
		"\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2"+
		"Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3"+
		"\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2"+
		"\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2"+
		"w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\3\u0081\3\2\2\2\5\u0085\3\2"+
		"\2\2\7\u008b\3\2\2\2\t\u0092\3\2\2\2\13\u0097\3\2\2\2\r\u00a1\3\2\2\2"+
		"\17\u00a6\3\2\2\2\21\u00ac\3\2\2\2\23\u00b0\3\2\2\2\25\u00b4\3\2\2\2\27"+
		"\u00b8\3\2\2\2\31\u00be\3\2\2\2\33\u00c1\3\2\2\2\35\u00c6\3\2\2\2\37\u00cd"+
		"\3\2\2\2!\u00d2\3\2\2\2#\u00d8\3\2\2\2%\u00e0\3\2\2\2\'\u00e6\3\2\2\2"+
		")\u00ea\3\2\2\2+\u00ed\3\2\2\2-\u00f3\3\2\2\2/\u00fc\3\2\2\2\61\u0103"+
		"\3\2\2\2\63\u0108\3\2\2\2\65\u010f\3\2\2\2\67\u0118\3\2\2\29\u011d\3\2"+
		"\2\2;\u0124\3\2\2\2=\u0129\3\2\2\2?\u0135\3\2\2\2A\u013e\3\2\2\2C\u0145"+
		"\3\2\2\2E\u0149\3\2\2\2G\u014c\3\2\2\2I\u014f\3\2\2\2K\u0151\3\2\2\2M"+
		"\u0153\3\2\2\2O\u0155\3\2\2\2Q\u0157\3\2\2\2S\u0159\3\2\2\2U\u015b\3\2"+
		"\2\2W\u015e\3\2\2\2Y\u0161\3\2\2\2[\u0163\3\2\2\2]\u0166\3\2\2\2_\u0169"+
		"\3\2\2\2a\u016b\3\2\2\2c\u016e\3\2\2\2e\u0170\3\2\2\2g\u0173\3\2\2\2i"+
		"\u0175\3\2\2\2k\u0177\3\2\2\2m\u0179\3\2\2\2o\u017b\3\2\2\2q\u017d\3\2"+
		"\2\2s\u017f\3\2\2\2u\u0181\3\2\2\2w\u0183\3\2\2\2y\u0186\3\2\2\2{\u018c"+
		"\3\2\2\2}\u019a\3\2\2\2\177\u01a5\3\2\2\2\u0081\u0082\7K\2\2\u0082\u0083"+
		"\7p\2\2\u0083\u0084\7v\2\2\u0084\4\3\2\2\2\u0085\u0086\7H\2\2\u0086\u0087"+
		"\7n\2\2\u0087\u0088\7q\2\2\u0088\u0089\7c\2\2\u0089\u008a\7v\2\2\u008a"+
		"\6\3\2\2\2\u008b\u008c\7U\2\2\u008c\u008d\7v\2\2\u008d\u008e\7t\2\2\u008e"+
		"\u008f\7k\2\2\u008f\u0090\7p\2\2\u0090\u0091\7i\2\2\u0091\b\3\2\2\2\u0092"+
		"\u0093\7D\2\2\u0093\u0094\7q\2\2\u0094\u0095\7q\2\2\u0095\u0096\7n\2\2"+
		"\u0096\n\3\2\2\2\u0097\u0098\7E\2\2\u0098\u0099\7j\2\2\u0099\u009a\7c"+
		"\2\2\u009a\u009b\7t\2\2\u009b\u009c\7c\2\2\u009c\u009d\7e\2\2\u009d\u009e"+
		"\7v\2\2\u009e\u009f\7g\2\2\u009f\u00a0\7t\2\2\u00a0\f\3\2\2\2\u00a1\u00a2"+
		"\7v\2\2\u00a2\u00a3\7t\2\2\u00a3\u00a4\7w\2\2\u00a4\u00a5\7g\2\2\u00a5"+
		"\16\3\2\2\2\u00a6\u00a7\7h\2\2\u00a7\u00a8\7c\2\2\u00a8\u00a9\7n\2\2\u00a9"+
		"\u00aa\7u\2\2\u00aa\u00ab\7g\2\2\u00ab\20\3\2\2\2\u00ac\u00ad\7p\2\2\u00ad"+
		"\u00ae\7k\2\2\u00ae\u00af\7n\2\2\u00af\22\3\2\2\2\u00b0\u00b1\7x\2\2\u00b1"+
		"\u00b2\7c\2\2\u00b2\u00b3\7t\2\2\u00b3\24\3\2\2\2\u00b4\u00b5\7n\2\2\u00b5"+
		"\u00b6\7g\2\2\u00b6\u00b7\7v\2\2\u00b7\26\3\2\2\2\u00b8\u00b9\7r\2\2\u00b9"+
		"\u00ba\7t\2\2\u00ba\u00bb\7k\2\2\u00bb\u00bc\7p\2\2\u00bc\u00bd\7v\2\2"+
		"\u00bd\30\3\2\2\2\u00be\u00bf\7k\2\2\u00bf\u00c0\7h\2\2\u00c0\32\3\2\2"+
		"\2\u00c1\u00c2\7g\2\2\u00c2\u00c3\7n\2\2\u00c3\u00c4\7u\2\2\u00c4\u00c5"+
		"\7g\2\2\u00c5\34\3\2\2\2\u00c6\u00c7\7u\2\2\u00c7\u00c8\7y\2\2\u00c8\u00c9"+
		"\7k\2\2\u00c9\u00ca\7v\2\2\u00ca\u00cb\7e\2\2\u00cb\u00cc\7j\2\2\u00cc"+
		"\36\3\2\2\2\u00cd\u00ce\7e\2\2\u00ce\u00cf\7c\2\2\u00cf\u00d0\7u\2\2\u00d0"+
		"\u00d1\7g\2\2\u00d1 \3\2\2\2\u00d2\u00d3\7d\2\2\u00d3\u00d4\7t\2\2\u00d4"+
		"\u00d5\7g\2\2\u00d5\u00d6\7c\2\2\u00d6\u00d7\7m\2\2\u00d7\"\3\2\2\2\u00d8"+
		"\u00d9\7f\2\2\u00d9\u00da\7g\2\2\u00da\u00db\7h\2\2\u00db\u00dc\7c\2\2"+
		"\u00dc\u00dd\7w\2\2\u00dd\u00de\7n\2\2\u00de\u00df\7v\2\2\u00df$\3\2\2"+
		"\2\u00e0\u00e1\7y\2\2\u00e1\u00e2\7j\2\2\u00e2\u00e3\7k\2\2\u00e3\u00e4"+
		"\7n\2\2\u00e4\u00e5\7g\2\2\u00e5&\3\2\2\2\u00e6\u00e7\7h\2\2\u00e7\u00e8"+
		"\7q\2\2\u00e8\u00e9\7t\2\2\u00e9(\3\2\2\2\u00ea\u00eb\7k\2\2\u00eb\u00ec"+
		"\7p\2\2\u00ec*\3\2\2\2\u00ed\u00ee\7i\2\2\u00ee\u00ef\7w\2\2\u00ef\u00f0"+
		"\7c\2\2\u00f0\u00f1\7t\2\2\u00f1\u00f2\7f\2\2\u00f2,\3\2\2\2\u00f3\u00f4"+
		"\7e\2\2\u00f4\u00f5\7q\2\2\u00f5\u00f6\7p\2\2\u00f6\u00f7\7v\2\2\u00f7"+
		"\u00f8\7k\2\2\u00f8\u00f9\7p\2\2\u00f9\u00fa\7w\2\2\u00fa\u00fb\7g\2\2"+
		"\u00fb.\3\2\2\2\u00fc\u00fd\7t\2\2\u00fd\u00fe\7g\2\2\u00fe\u00ff\7v\2"+
		"\2\u00ff\u0100\7w\2\2\u0100\u0101\7t\2\2\u0101\u0102\7p\2\2\u0102\60\3"+
		"\2\2\2\u0103\u0104\7h\2\2\u0104\u0105\7w\2\2\u0105\u0106\7p\2\2\u0106"+
		"\u0107\7e\2\2\u0107\62\3\2\2\2\u0108\u0109\7u\2\2\u0109\u010a\7v\2\2\u010a"+
		"\u010b\7t\2\2\u010b\u010c\7w\2\2\u010c\u010d\7e\2\2\u010d\u010e\7v\2\2"+
		"\u010e\64\3\2\2\2\u010f\u0110\7o\2\2\u0110\u0111\7w\2\2\u0111\u0112\7"+
		"v\2\2\u0112\u0113\7c\2\2\u0113\u0114\7v\2\2\u0114\u0115\7k\2\2\u0115\u0116"+
		"\7p\2\2\u0116\u0117\7i\2\2\u0117\66\3\2\2\2\u0118\u0119\7u\2\2\u0119\u011a"+
		"\7g\2\2\u011a\u011b\7n\2\2\u011b\u011c\7h\2\2\u011c8\3\2\2\2\u011d\u011e"+
		"\7k\2\2\u011e\u011f\7p\2\2\u011f\u0120\7q\2\2\u0120\u0121\7w\2\2\u0121"+
		"\u0122\7v\2\2\u0122:\3\2\2\2\u0123\u0125\t\2\2\2\u0124\u0123\3\2\2\2\u0125"+
		"\u0126\3\2\2\2\u0126\u0124\3\2\2\2\u0126\u0127\3\2\2\2\u0127<\3\2\2\2"+
		"\u0128\u012a\t\2\2\2\u0129\u0128\3\2\2\2\u012a\u012b\3\2\2\2\u012b\u0129"+
		"\3\2\2\2\u012b\u012c\3\2\2\2\u012c\u0133\3\2\2\2\u012d\u012f\7\60\2\2"+
		"\u012e\u0130\t\2\2\2\u012f\u012e\3\2\2\2\u0130\u0131\3\2\2\2\u0131\u012f"+
		"\3\2\2\2\u0131\u0132\3\2\2\2\u0132\u0134\3\2\2\2\u0133\u012d\3\2\2\2\u0133"+
		"\u0134\3\2\2\2\u0134>\3\2\2\2\u0135\u0139\7$\2\2\u0136\u0138\n\3\2\2\u0137"+
		"\u0136\3\2\2\2\u0138\u013b\3\2\2\2\u0139\u0137\3\2\2\2\u0139\u013a\3\2"+
		"\2\2\u013a\u013c\3\2\2\2\u013b\u0139\3\2\2\2\u013c\u013d\7$\2\2\u013d"+
		"@\3\2\2\2\u013e\u0142\t\4\2\2\u013f\u0141\t\5\2\2\u0140\u013f\3\2\2\2"+
		"\u0141\u0144\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0143\3\2\2\2\u0143B\3"+
		"\2\2\2\u0144\u0142\3\2\2\2\u0145\u0146\7$\2\2\u0146\u0147\n\6\2\2\u0147"+
		"\u0148\7$\2\2\u0148D\3\2\2\2\u0149\u014a\7-\2\2\u014a\u014b\7?\2\2\u014b"+
		"F\3\2\2\2\u014c\u014d\7/\2\2\u014d\u014e\7?\2\2\u014eH\3\2\2\2\u014f\u0150"+
		"\7-\2\2\u0150J\3\2\2\2\u0151\u0152\7/\2\2\u0152L\3\2\2\2\u0153\u0154\7"+
		",\2\2\u0154N\3\2\2\2\u0155\u0156\7\61\2\2\u0156P\3\2\2\2\u0157\u0158\7"+
		"\'\2\2\u0158R\3\2\2\2\u0159\u015a\7A\2\2\u015aT\3\2\2\2\u015b\u015c\7"+
		"~\2\2\u015c\u015d\7~\2\2\u015dV\3\2\2\2\u015e\u015f\7(\2\2\u015f\u0160"+
		"\7(\2\2\u0160X\3\2\2\2\u0161\u0162\7#\2\2\u0162Z\3\2\2\2\u0163\u0164\7"+
		"?\2\2\u0164\u0165\7?\2\2\u0165\\\3\2\2\2\u0166\u0167\7#\2\2\u0167\u0168"+
		"\7?\2\2\u0168^\3\2\2\2\u0169\u016a\7>\2\2\u016a`\3\2\2\2\u016b\u016c\7"+
		">\2\2\u016c\u016d\7?\2\2\u016db\3\2\2\2\u016e\u016f\7@\2\2\u016fd\3\2"+
		"\2\2\u0170\u0171\7@\2\2\u0171\u0172\7?\2\2\u0172f\3\2\2\2\u0173\u0174"+
		"\7?\2\2\u0174h\3\2\2\2\u0175\u0176\7\60\2\2\u0176j\3\2\2\2\u0177\u0178"+
		"\7.\2\2\u0178l\3\2\2\2\u0179\u017a\7<\2\2\u017an\3\2\2\2\u017b\u017c\7"+
		"=\2\2\u017cp\3\2\2\2\u017d\u017e\7*\2\2\u017er\3\2\2\2\u017f\u0180\7+"+
		"\2\2\u0180t\3\2\2\2\u0181\u0182\7}\2\2\u0182v\3\2\2\2\u0183\u0184\7\177"+
		"\2\2\u0184x\3\2\2\2\u0185\u0187\t\7\2\2\u0186\u0185\3\2\2\2\u0187\u0188"+
		"\3\2\2\2\u0188\u0186\3\2\2\2\u0188\u0189\3\2\2\2\u0189\u018a\3\2\2\2\u018a"+
		"\u018b\b=\2\2\u018bz\3\2\2\2\u018c\u018d\7\61\2\2\u018d\u018e\7,\2\2\u018e"+
		"\u0192\3\2\2\2\u018f\u0191\13\2\2\2\u0190\u018f\3\2\2\2\u0191\u0194\3"+
		"\2\2\2\u0192\u0193\3\2\2\2\u0192\u0190\3\2\2\2\u0193\u0195\3\2\2\2\u0194"+
		"\u0192\3\2\2\2\u0195\u0196\7,\2\2\u0196\u0197\7\61\2\2\u0197\u0198\3\2"+
		"\2\2\u0198\u0199\b>\2\2\u0199|\3\2\2\2\u019a\u019b\7\61\2\2\u019b\u019c"+
		"\7\61\2\2\u019c\u01a0\3\2\2\2\u019d\u019f\n\b\2\2\u019e\u019d\3\2\2\2"+
		"\u019f\u01a2\3\2\2\2\u01a0\u019e\3\2\2\2\u01a0\u01a1\3\2\2\2\u01a1\u01a3"+
		"\3\2\2\2\u01a2\u01a0\3\2\2\2\u01a3\u01a4\b?\2\2\u01a4~\3\2\2\2\u01a5\u01a6"+
		"\7^\2\2\u01a6\u01a7\t\t\2\2\u01a7\u0080\3\2\2\2\f\2\u0126\u012b\u0131"+
		"\u0133\u0139\u0142\u0188\u0192\u01a0\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}