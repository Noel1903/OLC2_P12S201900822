// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgramm
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "fmt"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftgrammParser struct {
	*antlr.BaseParser
}

var SwiftgrammParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammParserInit() {
	staticData := &SwiftgrammParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'",
		"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'",
		"'case'", "'break'", "'default'", "'while'", "'for'", "'in'", "'guard'",
		"'continue'", "'return'", "'func'", "'struct'", "'mutating'", "'self'",
		"'inout'", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'?'",
		"'||'", "'&&'", "'!'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
		"'='", "'.'", "','", "':'", "';'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE",
		"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK",
		"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC",
		"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "STRING_LITERAL", "ID",
		"CHARACTER_LITERAL", "SUMMATION", "SUBTRACTION", "MULTIPLICATION", "DIVISION",
		"MOD", "QUESTION_MARK", "OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", "LESS_THAN",
		"LESS_THAN_EQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL", "ASSIGN", "DOT",
		"COMMA", "COLON", "SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS",
		"OPEN_kEY", "CLOSE_kEY", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "sentences", "sentence", "print", "declare_var", "declare_constant",
		"assign_var", "if_sentence", "switch_sentence", "switch_cases", "switch_case",
		"while_sentence", "for_sentence", "guard_sentence", "break_sentence",
		"continue_sentence", "return_sentence", "expression", "datatype",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 204, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 1, 1, 5, 1, 45, 8, 1, 10, 1, 12, 1, 48, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 54, 8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 59, 8, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 84, 8, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 103, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 119, 8, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 131, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 141,
		8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 169, 8, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 180, 8, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 197, 8, 18, 10, 18, 12, 18, 200,
		9, 18, 1, 19, 1, 19, 1, 19, 0, 1, 36, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 6, 1, 0, 35, 36, 1, 0, 33,
		34, 1, 0, 44, 45, 1, 0, 46, 47, 1, 0, 42, 43, 1, 0, 1, 4, 204, 0, 40, 1,
		0, 0, 0, 2, 46, 1, 0, 0, 0, 4, 53, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 60,
		1, 0, 0, 0, 10, 83, 1, 0, 0, 0, 12, 102, 1, 0, 0, 0, 14, 104, 1, 0, 0,
		0, 16, 118, 1, 0, 0, 0, 18, 120, 1, 0, 0, 0, 20, 130, 1, 0, 0, 0, 22, 140,
		1, 0, 0, 0, 24, 142, 1, 0, 0, 0, 26, 148, 1, 0, 0, 0, 28, 156, 1, 0, 0,
		0, 30, 161, 1, 0, 0, 0, 32, 163, 1, 0, 0, 0, 34, 168, 1, 0, 0, 0, 36, 179,
		1, 0, 0, 0, 38, 201, 1, 0, 0, 0, 40, 41, 3, 2, 1, 0, 41, 42, 5, 0, 0, 1,
		42, 1, 1, 0, 0, 0, 43, 45, 3, 4, 2, 0, 44, 43, 1, 0, 0, 0, 45, 48, 1, 0,
		0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 3, 1, 0, 0, 0, 48, 46,
		1, 0, 0, 0, 49, 50, 3, 6, 3, 0, 50, 51, 3, 4, 2, 0, 51, 54, 1, 0, 0, 0,
		52, 54, 3, 6, 3, 0, 53, 49, 1, 0, 0, 0, 53, 52, 1, 0, 0, 0, 54, 5, 1, 0,
		0, 0, 55, 59, 3, 10, 5, 0, 56, 59, 3, 12, 6, 0, 57, 59, 3, 14, 7, 0, 58,
		55, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 7, 1, 0, 0,
		0, 60, 61, 5, 11, 0, 0, 61, 62, 5, 53, 0, 0, 62, 63, 3, 36, 18, 0, 63,
		64, 5, 54, 0, 0, 64, 9, 1, 0, 0, 0, 65, 66, 5, 9, 0, 0, 66, 67, 5, 31,
		0, 0, 67, 68, 5, 51, 0, 0, 68, 69, 3, 38, 19, 0, 69, 70, 5, 48, 0, 0, 70,
		71, 3, 36, 18, 0, 71, 72, 6, 5, -1, 0, 72, 84, 1, 0, 0, 0, 73, 74, 5, 9,
		0, 0, 74, 75, 5, 31, 0, 0, 75, 76, 5, 48, 0, 0, 76, 84, 3, 36, 18, 0, 77,
		78, 5, 9, 0, 0, 78, 79, 5, 31, 0, 0, 79, 80, 5, 51, 0, 0, 80, 81, 3, 38,
		19, 0, 81, 82, 5, 38, 0, 0, 82, 84, 1, 0, 0, 0, 83, 65, 1, 0, 0, 0, 83,
		73, 1, 0, 0, 0, 83, 77, 1, 0, 0, 0, 84, 11, 1, 0, 0, 0, 85, 86, 5, 10,
		0, 0, 86, 87, 5, 31, 0, 0, 87, 88, 5, 51, 0, 0, 88, 89, 3, 38, 19, 0, 89,
		90, 5, 48, 0, 0, 90, 91, 3, 36, 18, 0, 91, 103, 1, 0, 0, 0, 92, 93, 5,
		10, 0, 0, 93, 94, 5, 31, 0, 0, 94, 95, 5, 48, 0, 0, 95, 103, 3, 36, 18,
		0, 96, 97, 5, 10, 0, 0, 97, 98, 5, 31, 0, 0, 98, 99, 5, 51, 0, 0, 99, 100,
		3, 38, 19, 0, 100, 101, 5, 38, 0, 0, 101, 103, 1, 0, 0, 0, 102, 85, 1,
		0, 0, 0, 102, 92, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0, 103, 13, 1, 0, 0, 0,
		104, 105, 5, 31, 0, 0, 105, 106, 5, 48, 0, 0, 106, 107, 3, 36, 18, 0, 107,
		15, 1, 0, 0, 0, 108, 109, 5, 12, 0, 0, 109, 110, 3, 36, 18, 0, 110, 111,
		3, 4, 2, 0, 111, 119, 1, 0, 0, 0, 112, 113, 5, 12, 0, 0, 113, 114, 3, 36,
		18, 0, 114, 115, 3, 4, 2, 0, 115, 116, 5, 13, 0, 0, 116, 117, 3, 4, 2,
		0, 117, 119, 1, 0, 0, 0, 118, 108, 1, 0, 0, 0, 118, 112, 1, 0, 0, 0, 119,
		17, 1, 0, 0, 0, 120, 121, 5, 14, 0, 0, 121, 122, 3, 36, 18, 0, 122, 123,
		5, 55, 0, 0, 123, 124, 3, 20, 10, 0, 124, 125, 5, 56, 0, 0, 125, 19, 1,
		0, 0, 0, 126, 127, 3, 22, 11, 0, 127, 128, 3, 20, 10, 0, 128, 131, 1, 0,
		0, 0, 129, 131, 3, 22, 11, 0, 130, 126, 1, 0, 0, 0, 130, 129, 1, 0, 0,
		0, 131, 21, 1, 0, 0, 0, 132, 133, 5, 15, 0, 0, 133, 134, 3, 36, 18, 0,
		134, 135, 5, 51, 0, 0, 135, 136, 3, 4, 2, 0, 136, 141, 1, 0, 0, 0, 137,
		138, 5, 17, 0, 0, 138, 139, 5, 51, 0, 0, 139, 141, 3, 4, 2, 0, 140, 132,
		1, 0, 0, 0, 140, 137, 1, 0, 0, 0, 141, 23, 1, 0, 0, 0, 142, 143, 5, 18,
		0, 0, 143, 144, 3, 36, 18, 0, 144, 145, 5, 55, 0, 0, 145, 146, 3, 4, 2,
		0, 146, 147, 5, 56, 0, 0, 147, 25, 1, 0, 0, 0, 148, 149, 5, 19, 0, 0, 149,
		150, 5, 31, 0, 0, 150, 151, 5, 20, 0, 0, 151, 152, 3, 36, 18, 0, 152, 153,
		5, 55, 0, 0, 153, 154, 3, 4, 2, 0, 154, 155, 5, 56, 0, 0, 155, 27, 1, 0,
		0, 0, 156, 157, 5, 21, 0, 0, 157, 158, 3, 36, 18, 0, 158, 159, 5, 13, 0,
		0, 159, 160, 3, 4, 2, 0, 160, 29, 1, 0, 0, 0, 161, 162, 5, 16, 0, 0, 162,
		31, 1, 0, 0, 0, 163, 164, 5, 22, 0, 0, 164, 33, 1, 0, 0, 0, 165, 166, 5,
		23, 0, 0, 166, 169, 3, 36, 18, 0, 167, 169, 5, 23, 0, 0, 168, 165, 1, 0,
		0, 0, 168, 167, 1, 0, 0, 0, 169, 35, 1, 0, 0, 0, 170, 171, 6, 18, -1, 0,
		171, 172, 5, 53, 0, 0, 172, 173, 3, 36, 18, 0, 173, 174, 5, 54, 0, 0, 174,
		180, 1, 0, 0, 0, 175, 180, 5, 29, 0, 0, 176, 180, 5, 30, 0, 0, 177, 180,
		5, 32, 0, 0, 178, 180, 5, 31, 0, 0, 179, 170, 1, 0, 0, 0, 179, 175, 1,
		0, 0, 0, 179, 176, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 178, 1, 0, 0,
		0, 180, 198, 1, 0, 0, 0, 181, 182, 10, 10, 0, 0, 182, 183, 7, 0, 0, 0,
		183, 197, 3, 36, 18, 11, 184, 185, 10, 9, 0, 0, 185, 186, 7, 1, 0, 0, 186,
		197, 3, 36, 18, 10, 187, 188, 10, 8, 0, 0, 188, 189, 7, 2, 0, 0, 189, 197,
		3, 36, 18, 9, 190, 191, 10, 7, 0, 0, 191, 192, 7, 3, 0, 0, 192, 197, 3,
		36, 18, 8, 193, 194, 10, 6, 0, 0, 194, 195, 7, 4, 0, 0, 195, 197, 3, 36,
		18, 7, 196, 181, 1, 0, 0, 0, 196, 184, 1, 0, 0, 0, 196, 187, 1, 0, 0, 0,
		196, 190, 1, 0, 0, 0, 196, 193, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198,
		196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 37, 1, 0, 0, 0, 200, 198, 1,
		0, 0, 0, 201, 202, 7, 5, 0, 0, 202, 39, 1, 0, 0, 0, 12, 46, 53, 58, 83,
		102, 118, 130, 140, 168, 179, 196, 198,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftgrammParserInit initializes any static state used to implement SwiftgrammParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftgrammParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftgrammParserInit() {
	staticData := &SwiftgrammParserStaticData
	staticData.once.Do(swiftgrammParserInit)
}

// NewSwiftgrammParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftgrammParser(input antlr.TokenStream) *SwiftgrammParser {
	SwiftgrammParserInit()
	this := new(SwiftgrammParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftgrammParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Swiftgramm.g4"

	return this
}

// SwiftgrammParser tokens.
const (
	SwiftgrammParserEOF                = antlr.TokenEOF
	SwiftgrammParserINT                = 1
	SwiftgrammParserFLOAT              = 2
	SwiftgrammParserSTRING             = 3
	SwiftgrammParserBOOL               = 4
	SwiftgrammParserCHARACTER          = 5
	SwiftgrammParserTRUE               = 6
	SwiftgrammParserFALSE              = 7
	SwiftgrammParserNIL                = 8
	SwiftgrammParserVAR                = 9
	SwiftgrammParserLET                = 10
	SwiftgrammParserPRINT              = 11
	SwiftgrammParserIF                 = 12
	SwiftgrammParserELSE               = 13
	SwiftgrammParserSWITCH             = 14
	SwiftgrammParserCASE               = 15
	SwiftgrammParserBREAK              = 16
	SwiftgrammParserDEFAULT            = 17
	SwiftgrammParserWHILE              = 18
	SwiftgrammParserFOR                = 19
	SwiftgrammParserIN                 = 20
	SwiftgrammParserGUARD              = 21
	SwiftgrammParserCONTINUE           = 22
	SwiftgrammParserRETURN             = 23
	SwiftgrammParserFUNC               = 24
	SwiftgrammParserSTRUCT             = 25
	SwiftgrammParserMUTATING           = 26
	SwiftgrammParserSELF               = 27
	SwiftgrammParserINOUT              = 28
	SwiftgrammParserNUMBER             = 29
	SwiftgrammParserSTRING_LITERAL     = 30
	SwiftgrammParserID                 = 31
	SwiftgrammParserCHARACTER_LITERAL  = 32
	SwiftgrammParserSUMMATION          = 33
	SwiftgrammParserSUBTRACTION        = 34
	SwiftgrammParserMULTIPLICATION     = 35
	SwiftgrammParserDIVISION           = 36
	SwiftgrammParserMOD                = 37
	SwiftgrammParserQUESTION_MARK      = 38
	SwiftgrammParserOR                 = 39
	SwiftgrammParserAND                = 40
	SwiftgrammParserNOT                = 41
	SwiftgrammParserEQUAL              = 42
	SwiftgrammParserNOT_EQUAL          = 43
	SwiftgrammParserLESS_THAN          = 44
	SwiftgrammParserLESS_THAN_EQUAL    = 45
	SwiftgrammParserGREATER_THAN       = 46
	SwiftgrammParserGREATER_THAN_EQUAL = 47
	SwiftgrammParserASSIGN             = 48
	SwiftgrammParserDOT                = 49
	SwiftgrammParserCOMMA              = 50
	SwiftgrammParserCOLON              = 51
	SwiftgrammParserSEMICOLON          = 52
	SwiftgrammParserOPEN_PARENTHESIS   = 53
	SwiftgrammParserCLOSE_PARENTHESIS  = 54
	SwiftgrammParserOPEN_kEY           = 55
	SwiftgrammParserCLOSE_kEY          = 56
	SwiftgrammParserWHITESPACE         = 57
	SwiftgrammParserCOMMENT            = 58
	SwiftgrammParserLINE_COMMENT       = 59
)

// SwiftgrammParser rules.
const (
	SwiftgrammParserRULE_s                 = 0
	SwiftgrammParserRULE_block             = 1
	SwiftgrammParserRULE_sentences         = 2
	SwiftgrammParserRULE_sentence          = 3
	SwiftgrammParserRULE_print             = 4
	SwiftgrammParserRULE_declare_var       = 5
	SwiftgrammParserRULE_declare_constant  = 6
	SwiftgrammParserRULE_assign_var        = 7
	SwiftgrammParserRULE_if_sentence       = 8
	SwiftgrammParserRULE_switch_sentence   = 9
	SwiftgrammParserRULE_switch_cases      = 10
	SwiftgrammParserRULE_switch_case       = 11
	SwiftgrammParserRULE_while_sentence    = 12
	SwiftgrammParserRULE_for_sentence      = 13
	SwiftgrammParserRULE_guard_sentence    = 14
	SwiftgrammParserRULE_break_sentence    = 15
	SwiftgrammParserRULE_continue_sentence = 16
	SwiftgrammParserRULE_return_sentence   = 17
	SwiftgrammParserRULE_expression        = 18
	SwiftgrammParserRULE_datatype          = 19
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftgrammParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Block()
	}
	{
		p.SetState(41)
		p.Match(SwiftgrammParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSentences() []ISentencesContext
	Sentences(i int) ISentencesContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllSentences() []ISentencesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentencesContext); ok {
			len++
		}
	}

	tst := make([]ISentencesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentencesContext); ok {
			tst[i] = t.(ISentencesContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Sentences(i int) ISentencesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencesContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftgrammParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147485184) != 0 {
		{
			p.SetState(43)
			p.Sentences()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentencesContext is an interface to support dynamic dispatch.
type ISentencesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sentence() ISentenceContext
	Sentences() ISentencesContext

	// IsSentencesContext differentiates from other interfaces.
	IsSentencesContext()
}

type SentencesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentencesContext() *SentencesContext {
	var p = new(SentencesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_sentences
	return p
}

func InitEmptySentencesContext(p *SentencesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_sentences
}

func (*SentencesContext) IsSentencesContext() {}

func NewSentencesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentencesContext {
	var p = new(SentencesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_sentences

	return p
}

func (s *SentencesContext) GetParser() antlr.Parser { return s.parser }

func (s *SentencesContext) Sentence() ISentenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *SentencesContext) Sentences() ISentencesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencesContext)
}

func (s *SentencesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentencesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentencesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterSentences(s)
	}
}

func (s *SentencesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitSentences(s)
	}
}

func (s *SentencesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitSentences(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Sentences() (localctx ISentencesContext) {
	localctx = NewSentencesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftgrammParserRULE_sentences)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Sentence()
		}
		{
			p.SetState(50)
			p.Sentences()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Sentence()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenceContext is an interface to support dynamic dispatch.
type ISentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declare_var() IDeclare_varContext
	Declare_constant() IDeclare_constantContext
	Assign_var() IAssign_varContext

	// IsSentenceContext differentiates from other interfaces.
	IsSentenceContext()
}

type SentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenceContext() *SentenceContext {
	var p = new(SentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_sentence
	return p
}

func InitEmptySentenceContext(p *SentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_sentence
}

func (*SentenceContext) IsSentenceContext() {}

func NewSentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenceContext {
	var p = new(SentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_sentence

	return p
}

func (s *SentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenceContext) Declare_var() IDeclare_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclare_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclare_varContext)
}

func (s *SentenceContext) Declare_constant() IDeclare_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclare_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclare_constantContext)
}

func (s *SentenceContext) Assign_var() IAssign_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_varContext)
}

func (s *SentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterSentence(s)
	}
}

func (s *SentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitSentence(s)
	}
}

func (s *SentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitSentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Sentence() (localctx ISentenceContext) {
	localctx = NewSentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftgrammParserRULE_sentence)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Declare_var()
		}

	case SwiftgrammParserLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Declare_constant()
		}

	case SwiftgrammParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Assign_var()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	OPEN_PARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	CLOSE_PARENTHESIS() antlr.TerminalNode

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserPRINT, 0)
}

func (s *PrintContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *PrintContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintContext) CLOSE_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_PARENTHESIS, 0)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftgrammParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(SwiftgrammParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(SwiftgrammParserOPEN_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.expression(0)
	}
	{
		p.SetState(63)
		p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclare_varContext is an interface to support dynamic dispatch.
type IDeclare_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Datatype() IDatatypeContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	QUESTION_MARK() antlr.TerminalNode

	// IsDeclare_varContext differentiates from other interfaces.
	IsDeclare_varContext()
}

type Declare_varContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclare_varContext() *Declare_varContext {
	var p = new(Declare_varContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_declare_var
	return p
}

func InitEmptyDeclare_varContext(p *Declare_varContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_declare_var
}

func (*Declare_varContext) IsDeclare_varContext() {}

func NewDeclare_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declare_varContext {
	var p = new(Declare_varContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_declare_var

	return p
}

func (s *Declare_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Declare_varContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserVAR, 0)
}

func (s *Declare_varContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Declare_varContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *Declare_varContext) Datatype() IDatatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *Declare_varContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserASSIGN, 0)
}

func (s *Declare_varContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Declare_varContext) QUESTION_MARK() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserQUESTION_MARK, 0)
}

func (s *Declare_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declare_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declare_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterDeclare_var(s)
	}
}

func (s *Declare_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitDeclare_var(s)
	}
}

func (s *Declare_varContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitDeclare_var(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Declare_var() (localctx IDeclare_varContext) {
	localctx = NewDeclare_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftgrammParserRULE_declare_var)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Datatype()
		}
		{
			p.SetState(69)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.expression(0)
		}
		fmt.Println("Variable declaration: ")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Datatype()
		}
		{
			p.SetState(81)
			p.Match(SwiftgrammParserQUESTION_MARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclare_constantContext is an interface to support dynamic dispatch.
type IDeclare_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Datatype() IDatatypeContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	QUESTION_MARK() antlr.TerminalNode

	// IsDeclare_constantContext differentiates from other interfaces.
	IsDeclare_constantContext()
}

type Declare_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclare_constantContext() *Declare_constantContext {
	var p = new(Declare_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_declare_constant
	return p
}

func InitEmptyDeclare_constantContext(p *Declare_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_declare_constant
}

func (*Declare_constantContext) IsDeclare_constantContext() {}

func NewDeclare_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declare_constantContext {
	var p = new(Declare_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_declare_constant

	return p
}

func (s *Declare_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Declare_constantContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserLET, 0)
}

func (s *Declare_constantContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Declare_constantContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *Declare_constantContext) Datatype() IDatatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *Declare_constantContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserASSIGN, 0)
}

func (s *Declare_constantContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Declare_constantContext) QUESTION_MARK() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserQUESTION_MARK, 0)
}

func (s *Declare_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declare_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declare_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterDeclare_constant(s)
	}
}

func (s *Declare_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitDeclare_constant(s)
	}
}

func (s *Declare_constantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitDeclare_constant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Declare_constant() (localctx IDeclare_constantContext) {
	localctx = NewDeclare_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftgrammParserRULE_declare_constant)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Datatype()
		}
		{
			p.SetState(89)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Datatype()
		}
		{
			p.SetState(100)
			p.Match(SwiftgrammParserQUESTION_MARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssign_varContext is an interface to support dynamic dispatch.
type IAssign_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssign_varContext differentiates from other interfaces.
	IsAssign_varContext()
}

type Assign_varContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_varContext() *Assign_varContext {
	var p = new(Assign_varContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_assign_var
	return p
}

func InitEmptyAssign_varContext(p *Assign_varContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_assign_var
}

func (*Assign_varContext) IsAssign_varContext() {}

func NewAssign_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_varContext {
	var p = new(Assign_varContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_assign_var

	return p
}

func (s *Assign_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_varContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Assign_varContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserASSIGN, 0)
}

func (s *Assign_varContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assign_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterAssign_var(s)
	}
}

func (s *Assign_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitAssign_var(s)
	}
}

func (s *Assign_varContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitAssign_var(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Assign_var() (localctx IAssign_varContext) {
	localctx = NewAssign_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftgrammParserRULE_assign_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(SwiftgrammParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(SwiftgrammParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_sentenceContext is an interface to support dynamic dispatch.
type IIf_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllSentences() []ISentencesContext
	Sentences(i int) ISentencesContext
	ELSE() antlr.TerminalNode

	// IsIf_sentenceContext differentiates from other interfaces.
	IsIf_sentenceContext()
}

type If_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_sentenceContext() *If_sentenceContext {
	var p = new(If_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_if_sentence
	return p
}

func InitEmptyIf_sentenceContext(p *If_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_if_sentence
}

func (*If_sentenceContext) IsIf_sentenceContext() {}

func NewIf_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_sentenceContext {
	var p = new(If_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_if_sentence

	return p
}

func (s *If_sentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *If_sentenceContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserIF, 0)
}

func (s *If_sentenceContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_sentenceContext) AllSentences() []ISentencesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentencesContext); ok {
			len++
		}
	}

	tst := make([]ISentencesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentencesContext); ok {
			tst[i] = t.(ISentencesContext)
			i++
		}
	}

	return tst
}

func (s *If_sentenceContext) Sentences(i int) ISentencesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencesContext)
}

func (s *If_sentenceContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserELSE, 0)
}

func (s *If_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterIf_sentence(s)
	}
}

func (s *If_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitIf_sentence(s)
	}
}

func (s *If_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitIf_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) If_sentence() (localctx IIf_sentenceContext) {
	localctx = NewIf_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftgrammParserRULE_if_sentence)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.expression(0)
		}
		{
			p.SetState(110)
			p.Sentences()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.expression(0)
		}
		{
			p.SetState(114)
			p.Sentences()
		}
		{
			p.SetState(115)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Sentences()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_sentenceContext is an interface to support dynamic dispatch.
type ISwitch_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expression() IExpressionContext
	OPEN_kEY() antlr.TerminalNode
	Switch_cases() ISwitch_casesContext
	CLOSE_kEY() antlr.TerminalNode

	// IsSwitch_sentenceContext differentiates from other interfaces.
	IsSwitch_sentenceContext()
}

type Switch_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_sentenceContext() *Switch_sentenceContext {
	var p = new(Switch_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_switch_sentence
	return p
}

func InitEmptySwitch_sentenceContext(p *Switch_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_switch_sentence
}

func (*Switch_sentenceContext) IsSwitch_sentenceContext() {}

func NewSwitch_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_sentenceContext {
	var p = new(Switch_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_switch_sentence

	return p
}

func (s *Switch_sentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_sentenceContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSWITCH, 0)
}

func (s *Switch_sentenceContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Switch_sentenceContext) OPEN_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, 0)
}

func (s *Switch_sentenceContext) Switch_cases() ISwitch_casesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_casesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_casesContext)
}

func (s *Switch_sentenceContext) CLOSE_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, 0)
}

func (s *Switch_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterSwitch_sentence(s)
	}
}

func (s *Switch_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitSwitch_sentence(s)
	}
}

func (s *Switch_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitSwitch_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Switch_sentence() (localctx ISwitch_sentenceContext) {
	localctx = NewSwitch_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftgrammParserRULE_switch_sentence)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(SwiftgrammParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.expression(0)
	}
	{
		p.SetState(122)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Switch_cases()
	}
	{
		p.SetState(124)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_casesContext is an interface to support dynamic dispatch.
type ISwitch_casesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Switch_case() ISwitch_caseContext
	Switch_cases() ISwitch_casesContext

	// IsSwitch_casesContext differentiates from other interfaces.
	IsSwitch_casesContext()
}

type Switch_casesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_casesContext() *Switch_casesContext {
	var p = new(Switch_casesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_switch_cases
	return p
}

func InitEmptySwitch_casesContext(p *Switch_casesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_switch_cases
}

func (*Switch_casesContext) IsSwitch_casesContext() {}

func NewSwitch_casesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_casesContext {
	var p = new(Switch_casesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_switch_cases

	return p
}

func (s *Switch_casesContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_casesContext) Switch_case() ISwitch_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_caseContext)
}

func (s *Switch_casesContext) Switch_cases() ISwitch_casesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_casesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_casesContext)
}

func (s *Switch_casesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_casesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_casesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterSwitch_cases(s)
	}
}

func (s *Switch_casesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitSwitch_cases(s)
	}
}

func (s *Switch_casesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitSwitch_cases(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Switch_cases() (localctx ISwitch_casesContext) {
	localctx = NewSwitch_casesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftgrammParserRULE_switch_cases)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Switch_case()
		}
		{
			p.SetState(127)
			p.Switch_cases()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Switch_case()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	Sentences() ISentencesContext
	DEFAULT() antlr.TerminalNode

	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCASE, 0)
}

func (s *Switch_caseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Switch_caseContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *Switch_caseContext) Sentences() ISentencesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencesContext)
}

func (s *Switch_caseContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserDEFAULT, 0)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterSwitch_case(s)
	}
}

func (s *Switch_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitSwitch_case(s)
	}
}

func (s *Switch_caseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitSwitch_case(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftgrammParserRULE_switch_case)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserCASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(SwiftgrammParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.expression(0)
		}
		{
			p.SetState(134)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Sentences()
		}

	case SwiftgrammParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(SwiftgrammParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Sentences()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_sentenceContext is an interface to support dynamic dispatch.
type IWhile_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	OPEN_kEY() antlr.TerminalNode
	Sentences() ISentencesContext
	CLOSE_kEY() antlr.TerminalNode

	// IsWhile_sentenceContext differentiates from other interfaces.
	IsWhile_sentenceContext()
}

type While_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_sentenceContext() *While_sentenceContext {
	var p = new(While_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_while_sentence
	return p
}

func InitEmptyWhile_sentenceContext(p *While_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_while_sentence
}

func (*While_sentenceContext) IsWhile_sentenceContext() {}

func NewWhile_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_sentenceContext {
	var p = new(While_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_while_sentence

	return p
}

func (s *While_sentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *While_sentenceContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserWHILE, 0)
}

func (s *While_sentenceContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *While_sentenceContext) OPEN_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, 0)
}

func (s *While_sentenceContext) Sentences() ISentencesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencesContext)
}

func (s *While_sentenceContext) CLOSE_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, 0)
}

func (s *While_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterWhile_sentence(s)
	}
}

func (s *While_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitWhile_sentence(s)
	}
}

func (s *While_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitWhile_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) While_sentence() (localctx IWhile_sentenceContext) {
	localctx = NewWhile_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftgrammParserRULE_while_sentence)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(SwiftgrammParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.expression(0)
	}
	{
		p.SetState(144)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Sentences()
	}
	{
		p.SetState(146)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_sentenceContext is an interface to support dynamic dispatch.
type IFor_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	Expression() IExpressionContext
	OPEN_kEY() antlr.TerminalNode
	Sentences() ISentencesContext
	CLOSE_kEY() antlr.TerminalNode

	// IsFor_sentenceContext differentiates from other interfaces.
	IsFor_sentenceContext()
}

type For_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_sentenceContext() *For_sentenceContext {
	var p = new(For_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_for_sentence
	return p
}

func InitEmptyFor_sentenceContext(p *For_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_for_sentence
}

func (*For_sentenceContext) IsFor_sentenceContext() {}

func NewFor_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_sentenceContext {
	var p = new(For_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_for_sentence

	return p
}

func (s *For_sentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *For_sentenceContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserFOR, 0)
}

func (s *For_sentenceContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *For_sentenceContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserIN, 0)
}

func (s *For_sentenceContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *For_sentenceContext) OPEN_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, 0)
}

func (s *For_sentenceContext) Sentences() ISentencesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencesContext)
}

func (s *For_sentenceContext) CLOSE_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, 0)
}

func (s *For_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterFor_sentence(s)
	}
}

func (s *For_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitFor_sentence(s)
	}
}

func (s *For_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitFor_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) For_sentence() (localctx IFor_sentenceContext) {
	localctx = NewFor_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftgrammParserRULE_for_sentence)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(SwiftgrammParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(SwiftgrammParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(SwiftgrammParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.expression(0)
	}
	{
		p.SetState(152)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Sentences()
	}
	{
		p.SetState(154)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuard_sentenceContext is an interface to support dynamic dispatch.
type IGuard_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expression() IExpressionContext
	ELSE() antlr.TerminalNode
	Sentences() ISentencesContext

	// IsGuard_sentenceContext differentiates from other interfaces.
	IsGuard_sentenceContext()
}

type Guard_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuard_sentenceContext() *Guard_sentenceContext {
	var p = new(Guard_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_guard_sentence
	return p
}

func InitEmptyGuard_sentenceContext(p *Guard_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_guard_sentence
}

func (*Guard_sentenceContext) IsGuard_sentenceContext() {}

func NewGuard_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Guard_sentenceContext {
	var p = new(Guard_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_guard_sentence

	return p
}

func (s *Guard_sentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Guard_sentenceContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserGUARD, 0)
}

func (s *Guard_sentenceContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Guard_sentenceContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserELSE, 0)
}

func (s *Guard_sentenceContext) Sentences() ISentencesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencesContext)
}

func (s *Guard_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Guard_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Guard_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterGuard_sentence(s)
	}
}

func (s *Guard_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitGuard_sentence(s)
	}
}

func (s *Guard_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitGuard_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Guard_sentence() (localctx IGuard_sentenceContext) {
	localctx = NewGuard_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftgrammParserRULE_guard_sentence)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(SwiftgrammParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.expression(0)
	}
	{
		p.SetState(158)
		p.Match(SwiftgrammParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Sentences()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreak_sentenceContext is an interface to support dynamic dispatch.
type IBreak_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreak_sentenceContext differentiates from other interfaces.
	IsBreak_sentenceContext()
}

type Break_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_sentenceContext() *Break_sentenceContext {
	var p = new(Break_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_break_sentence
	return p
}

func InitEmptyBreak_sentenceContext(p *Break_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_break_sentence
}

func (*Break_sentenceContext) IsBreak_sentenceContext() {}

func NewBreak_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_sentenceContext {
	var p = new(Break_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_break_sentence

	return p
}

func (s *Break_sentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Break_sentenceContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserBREAK, 0)
}

func (s *Break_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterBreak_sentence(s)
	}
}

func (s *Break_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitBreak_sentence(s)
	}
}

func (s *Break_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitBreak_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Break_sentence() (localctx IBreak_sentenceContext) {
	localctx = NewBreak_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftgrammParserRULE_break_sentence)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(SwiftgrammParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinue_sentenceContext is an interface to support dynamic dispatch.
type IContinue_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinue_sentenceContext differentiates from other interfaces.
	IsContinue_sentenceContext()
}

type Continue_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_sentenceContext() *Continue_sentenceContext {
	var p = new(Continue_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_continue_sentence
	return p
}

func InitEmptyContinue_sentenceContext(p *Continue_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_continue_sentence
}

func (*Continue_sentenceContext) IsContinue_sentenceContext() {}

func NewContinue_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_sentenceContext {
	var p = new(Continue_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_continue_sentence

	return p
}

func (s *Continue_sentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Continue_sentenceContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCONTINUE, 0)
}

func (s *Continue_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterContinue_sentence(s)
	}
}

func (s *Continue_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitContinue_sentence(s)
	}
}

func (s *Continue_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitContinue_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Continue_sentence() (localctx IContinue_sentenceContext) {
	localctx = NewContinue_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftgrammParserRULE_continue_sentence)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(SwiftgrammParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_sentenceContext is an interface to support dynamic dispatch.
type IReturn_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturn_sentenceContext differentiates from other interfaces.
	IsReturn_sentenceContext()
}

type Return_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_sentenceContext() *Return_sentenceContext {
	var p = new(Return_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_return_sentence
	return p
}

func InitEmptyReturn_sentenceContext(p *Return_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_return_sentence
}

func (*Return_sentenceContext) IsReturn_sentenceContext() {}

func NewReturn_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_sentenceContext {
	var p = new(Return_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_return_sentence

	return p
}

func (s *Return_sentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_sentenceContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserRETURN, 0)
}

func (s *Return_sentenceContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterReturn_sentence(s)
	}
}

func (s *Return_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitReturn_sentence(s)
	}
}

func (s *Return_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitReturn_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Return_sentence() (localctx IReturn_sentenceContext) {
	localctx = NewReturn_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftgrammParserRULE_return_sentence)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(SwiftgrammParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(SwiftgrammParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOper returns the oper token.
	GetOper() antlr.Token

	// SetOper sets the oper token.
	SetOper(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// Getter signatures
	OPEN_PARENTHESIS() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	CLOSE_PARENTHESIS() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	CHARACTER_LITERAL() antlr.TerminalNode
	ID() antlr.TerminalNode
	MULTIPLICATION() antlr.TerminalNode
	DIVISION() antlr.TerminalNode
	SUMMATION() antlr.TerminalNode
	SUBTRACTION() antlr.TerminalNode
	LESS_THAN() antlr.TerminalNode
	LESS_THAN_EQUAL() antlr.TerminalNode
	GREATER_THAN() antlr.TerminalNode
	GREATER_THAN_EQUAL() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpressionContext
	oper   antlr.Token
	right  IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetOper() antlr.Token { return s.oper }

func (s *ExpressionContext) SetOper(v antlr.Token) { s.oper = v }

func (s *ExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *ExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ExpressionContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) CLOSE_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_PARENTHESIS, 0)
}

func (s *ExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserNUMBER, 0)
}

func (s *ExpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSTRING_LITERAL, 0)
}

func (s *ExpressionContext) CHARACTER_LITERAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCHARACTER_LITERAL, 0)
}

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *ExpressionContext) MULTIPLICATION() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserMULTIPLICATION, 0)
}

func (s *ExpressionContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserDIVISION, 0)
}

func (s *ExpressionContext) SUMMATION() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSUMMATION, 0)
}

func (s *ExpressionContext) SUBTRACTION() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSUBTRACTION, 0)
}

func (s *ExpressionContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserLESS_THAN, 0)
}

func (s *ExpressionContext) LESS_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserLESS_THAN_EQUAL, 0)
}

func (s *ExpressionContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserGREATER_THAN, 0)
}

func (s *ExpressionContext) GREATER_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserGREATER_THAN_EQUAL, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserEQUAL, 0)
}

func (s *ExpressionContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserNOT_EQUAL, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SwiftgrammParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, SwiftgrammParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserOPEN_PARENTHESIS:
		{
			p.SetState(171)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.expression(0)
		}
		{
			p.SetState(173)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftgrammParserNUMBER:
		{
			p.SetState(175)
			p.Match(SwiftgrammParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftgrammParserSTRING_LITERAL:
		{
			p.SetState(176)
			p.Match(SwiftgrammParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftgrammParserCHARACTER_LITERAL:
		{
			p.SetState(177)
			p.Match(SwiftgrammParserCHARACTER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftgrammParserID:
		{
			p.SetState(178)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(181)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(182)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).oper = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftgrammParserMULTIPLICATION || _la == SwiftgrammParserDIVISION) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).oper = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(183)

					var _x = p.expression(11)

					localctx.(*ExpressionContext).right = _x
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(185)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).oper = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftgrammParserSUMMATION || _la == SwiftgrammParserSUBTRACTION) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).oper = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(186)

					var _x = p.expression(10)

					localctx.(*ExpressionContext).right = _x
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(188)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).oper = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftgrammParserLESS_THAN || _la == SwiftgrammParserLESS_THAN_EQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).oper = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(189)

					var _x = p.expression(9)

					localctx.(*ExpressionContext).right = _x
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(191)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).oper = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftgrammParserGREATER_THAN || _la == SwiftgrammParserGREATER_THAN_EQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).oper = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(192)

					var _x = p.expression(8)

					localctx.(*ExpressionContext).right = _x
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(194)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).oper = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftgrammParserEQUAL || _la == SwiftgrammParserNOT_EQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).oper = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(195)

					var _x = p.expression(7)

					localctx.(*ExpressionContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatatypeContext is an interface to support dynamic dispatch.
type IDatatypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode

	// IsDatatypeContext differentiates from other interfaces.
	IsDatatypeContext()
}

type DatatypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatatypeContext() *DatatypeContext {
	var p = new(DatatypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_datatype
	return p
}

func InitEmptyDatatypeContext(p *DatatypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_datatype
}

func (*DatatypeContext) IsDatatypeContext() {}

func NewDatatypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatatypeContext {
	var p = new(DatatypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_datatype

	return p
}

func (s *DatatypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatatypeContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserINT, 0)
}

func (s *DatatypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserFLOAT, 0)
}

func (s *DatatypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSTRING, 0)
}

func (s *DatatypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserBOOL, 0)
}

func (s *DatatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatatypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterDatatype(s)
	}
}

func (s *DatatypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitDatatype(s)
	}
}

func (s *DatatypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitDatatype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Datatype() (localctx IDatatypeContext) {
	localctx = NewDatatypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftgrammParserRULE_datatype)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftgrammParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftgrammParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
