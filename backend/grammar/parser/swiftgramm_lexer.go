// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SwiftgrammLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SwiftgrammLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammlexerLexerInit() {
	staticData := &SwiftgrammLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'",
		"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'",
		"'case'", "'break'", "'default'", "'while'", "'for'", "'in'", "'guard'",
		"'continue'", "'return'", "'func'", "'struct'", "'mutating'", "'self'",
		"'inout'", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "'?'", "'||'",
		"'&&'", "'!'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'='",
		"'.'", "','", "':'", "';'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE",
		"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK",
		"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC",
		"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "STRING_LITERAL", "ID",
		"CHARACTER_LITERAL", "SUMMATION", "SUBTRACTION", "MULTIPLICATION", "DIVISION",
		"QUESTION_MARK", "OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", "LESS_THAN",
		"LESS_THAN_EQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL", "ASSIGN", "DOT",
		"COMMA", "COLON", "SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS",
		"OPEN_kEY", "CLOSE_kEY", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE", "NIL",
		"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK", "DEFAULT",
		"WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC", "STRUCT",
		"MUTATING", "SELF", "INOUT", "NUMBER", "STRING_LITERAL", "ID", "CHARACTER_LITERAL",
		"SUMMATION", "SUBTRACTION", "MULTIPLICATION", "DIVISION", "QUESTION_MARK",
		"OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL",
		"GREATER_THAN", "GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON",
		"SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY",
		"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 58, 401, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 28, 4, 28, 283, 8, 28, 11, 28, 12, 28, 284, 1, 28, 1, 28, 4, 28,
		289, 8, 28, 11, 28, 12, 28, 290, 3, 28, 293, 8, 28, 1, 29, 1, 29, 5, 29,
		297, 8, 29, 10, 29, 12, 29, 300, 9, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5,
		30, 306, 8, 30, 10, 30, 12, 30, 309, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1,
		37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40,
		1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49,
		1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1,
		55, 4, 55, 368, 8, 55, 11, 55, 12, 55, 369, 1, 55, 1, 55, 1, 56, 1, 56,
		1, 56, 1, 56, 5, 56, 378, 8, 56, 10, 56, 12, 56, 381, 9, 56, 1, 56, 1,
		56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 5, 57, 392, 8, 57,
		10, 57, 12, 57, 395, 9, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 1, 379,
		0, 59, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46,
		93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109,
		55, 111, 56, 113, 57, 115, 58, 117, 0, 1, 0, 8, 1, 0, 48, 57, 1, 0, 34,
		34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		1, 0, 39, 39, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 5, 0,
		35, 35, 40, 43, 45, 46, 63, 64, 91, 93, 407, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65,
		1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0,
		73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0,
		0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0,
		0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0,
		0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0,
		0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 1, 119, 1,
		0, 0, 0, 3, 123, 1, 0, 0, 0, 5, 129, 1, 0, 0, 0, 7, 136, 1, 0, 0, 0, 9,
		141, 1, 0, 0, 0, 11, 151, 1, 0, 0, 0, 13, 156, 1, 0, 0, 0, 15, 162, 1,
		0, 0, 0, 17, 166, 1, 0, 0, 0, 19, 170, 1, 0, 0, 0, 21, 174, 1, 0, 0, 0,
		23, 180, 1, 0, 0, 0, 25, 183, 1, 0, 0, 0, 27, 188, 1, 0, 0, 0, 29, 195,
		1, 0, 0, 0, 31, 200, 1, 0, 0, 0, 33, 206, 1, 0, 0, 0, 35, 214, 1, 0, 0,
		0, 37, 220, 1, 0, 0, 0, 39, 224, 1, 0, 0, 0, 41, 227, 1, 0, 0, 0, 43, 233,
		1, 0, 0, 0, 45, 242, 1, 0, 0, 0, 47, 249, 1, 0, 0, 0, 49, 254, 1, 0, 0,
		0, 51, 261, 1, 0, 0, 0, 53, 270, 1, 0, 0, 0, 55, 275, 1, 0, 0, 0, 57, 282,
		1, 0, 0, 0, 59, 294, 1, 0, 0, 0, 61, 303, 1, 0, 0, 0, 63, 310, 1, 0, 0,
		0, 65, 314, 1, 0, 0, 0, 67, 316, 1, 0, 0, 0, 69, 318, 1, 0, 0, 0, 71, 320,
		1, 0, 0, 0, 73, 322, 1, 0, 0, 0, 75, 324, 1, 0, 0, 0, 77, 327, 1, 0, 0,
		0, 79, 330, 1, 0, 0, 0, 81, 332, 1, 0, 0, 0, 83, 335, 1, 0, 0, 0, 85, 338,
		1, 0, 0, 0, 87, 340, 1, 0, 0, 0, 89, 343, 1, 0, 0, 0, 91, 345, 1, 0, 0,
		0, 93, 348, 1, 0, 0, 0, 95, 350, 1, 0, 0, 0, 97, 352, 1, 0, 0, 0, 99, 354,
		1, 0, 0, 0, 101, 356, 1, 0, 0, 0, 103, 358, 1, 0, 0, 0, 105, 360, 1, 0,
		0, 0, 107, 362, 1, 0, 0, 0, 109, 364, 1, 0, 0, 0, 111, 367, 1, 0, 0, 0,
		113, 373, 1, 0, 0, 0, 115, 387, 1, 0, 0, 0, 117, 398, 1, 0, 0, 0, 119,
		120, 5, 73, 0, 0, 120, 121, 5, 110, 0, 0, 121, 122, 5, 116, 0, 0, 122,
		2, 1, 0, 0, 0, 123, 124, 5, 70, 0, 0, 124, 125, 5, 108, 0, 0, 125, 126,
		5, 111, 0, 0, 126, 127, 5, 97, 0, 0, 127, 128, 5, 116, 0, 0, 128, 4, 1,
		0, 0, 0, 129, 130, 5, 83, 0, 0, 130, 131, 5, 116, 0, 0, 131, 132, 5, 114,
		0, 0, 132, 133, 5, 105, 0, 0, 133, 134, 5, 110, 0, 0, 134, 135, 5, 103,
		0, 0, 135, 6, 1, 0, 0, 0, 136, 137, 5, 66, 0, 0, 137, 138, 5, 111, 0, 0,
		138, 139, 5, 111, 0, 0, 139, 140, 5, 108, 0, 0, 140, 8, 1, 0, 0, 0, 141,
		142, 5, 67, 0, 0, 142, 143, 5, 104, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145,
		5, 114, 0, 0, 145, 146, 5, 97, 0, 0, 146, 147, 5, 99, 0, 0, 147, 148, 5,
		116, 0, 0, 148, 149, 5, 101, 0, 0, 149, 150, 5, 114, 0, 0, 150, 10, 1,
		0, 0, 0, 151, 152, 5, 116, 0, 0, 152, 153, 5, 114, 0, 0, 153, 154, 5, 117,
		0, 0, 154, 155, 5, 101, 0, 0, 155, 12, 1, 0, 0, 0, 156, 157, 5, 102, 0,
		0, 157, 158, 5, 97, 0, 0, 158, 159, 5, 108, 0, 0, 159, 160, 5, 115, 0,
		0, 160, 161, 5, 101, 0, 0, 161, 14, 1, 0, 0, 0, 162, 163, 5, 110, 0, 0,
		163, 164, 5, 105, 0, 0, 164, 165, 5, 108, 0, 0, 165, 16, 1, 0, 0, 0, 166,
		167, 5, 118, 0, 0, 167, 168, 5, 97, 0, 0, 168, 169, 5, 114, 0, 0, 169,
		18, 1, 0, 0, 0, 170, 171, 5, 108, 0, 0, 171, 172, 5, 101, 0, 0, 172, 173,
		5, 116, 0, 0, 173, 20, 1, 0, 0, 0, 174, 175, 5, 112, 0, 0, 175, 176, 5,
		114, 0, 0, 176, 177, 5, 105, 0, 0, 177, 178, 5, 110, 0, 0, 178, 179, 5,
		116, 0, 0, 179, 22, 1, 0, 0, 0, 180, 181, 5, 105, 0, 0, 181, 182, 5, 102,
		0, 0, 182, 24, 1, 0, 0, 0, 183, 184, 5, 101, 0, 0, 184, 185, 5, 108, 0,
		0, 185, 186, 5, 115, 0, 0, 186, 187, 5, 101, 0, 0, 187, 26, 1, 0, 0, 0,
		188, 189, 5, 115, 0, 0, 189, 190, 5, 119, 0, 0, 190, 191, 5, 105, 0, 0,
		191, 192, 5, 116, 0, 0, 192, 193, 5, 99, 0, 0, 193, 194, 5, 104, 0, 0,
		194, 28, 1, 0, 0, 0, 195, 196, 5, 99, 0, 0, 196, 197, 5, 97, 0, 0, 197,
		198, 5, 115, 0, 0, 198, 199, 5, 101, 0, 0, 199, 30, 1, 0, 0, 0, 200, 201,
		5, 98, 0, 0, 201, 202, 5, 114, 0, 0, 202, 203, 5, 101, 0, 0, 203, 204,
		5, 97, 0, 0, 204, 205, 5, 107, 0, 0, 205, 32, 1, 0, 0, 0, 206, 207, 5,
		100, 0, 0, 207, 208, 5, 101, 0, 0, 208, 209, 5, 102, 0, 0, 209, 210, 5,
		97, 0, 0, 210, 211, 5, 117, 0, 0, 211, 212, 5, 108, 0, 0, 212, 213, 5,
		116, 0, 0, 213, 34, 1, 0, 0, 0, 214, 215, 5, 119, 0, 0, 215, 216, 5, 104,
		0, 0, 216, 217, 5, 105, 0, 0, 217, 218, 5, 108, 0, 0, 218, 219, 5, 101,
		0, 0, 219, 36, 1, 0, 0, 0, 220, 221, 5, 102, 0, 0, 221, 222, 5, 111, 0,
		0, 222, 223, 5, 114, 0, 0, 223, 38, 1, 0, 0, 0, 224, 225, 5, 105, 0, 0,
		225, 226, 5, 110, 0, 0, 226, 40, 1, 0, 0, 0, 227, 228, 5, 103, 0, 0, 228,
		229, 5, 117, 0, 0, 229, 230, 5, 97, 0, 0, 230, 231, 5, 114, 0, 0, 231,
		232, 5, 100, 0, 0, 232, 42, 1, 0, 0, 0, 233, 234, 5, 99, 0, 0, 234, 235,
		5, 111, 0, 0, 235, 236, 5, 110, 0, 0, 236, 237, 5, 116, 0, 0, 237, 238,
		5, 105, 0, 0, 238, 239, 5, 110, 0, 0, 239, 240, 5, 117, 0, 0, 240, 241,
		5, 101, 0, 0, 241, 44, 1, 0, 0, 0, 242, 243, 5, 114, 0, 0, 243, 244, 5,
		101, 0, 0, 244, 245, 5, 116, 0, 0, 245, 246, 5, 117, 0, 0, 246, 247, 5,
		114, 0, 0, 247, 248, 5, 110, 0, 0, 248, 46, 1, 0, 0, 0, 249, 250, 5, 102,
		0, 0, 250, 251, 5, 117, 0, 0, 251, 252, 5, 110, 0, 0, 252, 253, 5, 99,
		0, 0, 253, 48, 1, 0, 0, 0, 254, 255, 5, 115, 0, 0, 255, 256, 5, 116, 0,
		0, 256, 257, 5, 114, 0, 0, 257, 258, 5, 117, 0, 0, 258, 259, 5, 99, 0,
		0, 259, 260, 5, 116, 0, 0, 260, 50, 1, 0, 0, 0, 261, 262, 5, 109, 0, 0,
		262, 263, 5, 117, 0, 0, 263, 264, 5, 116, 0, 0, 264, 265, 5, 97, 0, 0,
		265, 266, 5, 116, 0, 0, 266, 267, 5, 105, 0, 0, 267, 268, 5, 110, 0, 0,
		268, 269, 5, 103, 0, 0, 269, 52, 1, 0, 0, 0, 270, 271, 5, 115, 0, 0, 271,
		272, 5, 101, 0, 0, 272, 273, 5, 108, 0, 0, 273, 274, 5, 102, 0, 0, 274,
		54, 1, 0, 0, 0, 275, 276, 5, 105, 0, 0, 276, 277, 5, 110, 0, 0, 277, 278,
		5, 111, 0, 0, 278, 279, 5, 117, 0, 0, 279, 280, 5, 116, 0, 0, 280, 56,
		1, 0, 0, 0, 281, 283, 7, 0, 0, 0, 282, 281, 1, 0, 0, 0, 283, 284, 1, 0,
		0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 292, 1, 0, 0, 0,
		286, 288, 5, 46, 0, 0, 287, 289, 7, 0, 0, 0, 288, 287, 1, 0, 0, 0, 289,
		290, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 293,
		1, 0, 0, 0, 292, 286, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 58, 1, 0,
		0, 0, 294, 298, 5, 34, 0, 0, 295, 297, 8, 1, 0, 0, 296, 295, 1, 0, 0, 0,
		297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299,
		301, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 302, 5, 34, 0, 0, 302, 60,
		1, 0, 0, 0, 303, 307, 7, 2, 0, 0, 304, 306, 7, 3, 0, 0, 305, 304, 1, 0,
		0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0,
		308, 62, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 311, 5, 34, 0, 0, 311,
		312, 8, 4, 0, 0, 312, 313, 5, 34, 0, 0, 313, 64, 1, 0, 0, 0, 314, 315,
		5, 43, 0, 0, 315, 66, 1, 0, 0, 0, 316, 317, 5, 45, 0, 0, 317, 68, 1, 0,
		0, 0, 318, 319, 5, 42, 0, 0, 319, 70, 1, 0, 0, 0, 320, 321, 5, 47, 0, 0,
		321, 72, 1, 0, 0, 0, 322, 323, 5, 63, 0, 0, 323, 74, 1, 0, 0, 0, 324, 325,
		5, 124, 0, 0, 325, 326, 5, 124, 0, 0, 326, 76, 1, 0, 0, 0, 327, 328, 5,
		38, 0, 0, 328, 329, 5, 38, 0, 0, 329, 78, 1, 0, 0, 0, 330, 331, 5, 33,
		0, 0, 331, 80, 1, 0, 0, 0, 332, 333, 5, 61, 0, 0, 333, 334, 5, 61, 0, 0,
		334, 82, 1, 0, 0, 0, 335, 336, 5, 33, 0, 0, 336, 337, 5, 61, 0, 0, 337,
		84, 1, 0, 0, 0, 338, 339, 5, 60, 0, 0, 339, 86, 1, 0, 0, 0, 340, 341, 5,
		60, 0, 0, 341, 342, 5, 61, 0, 0, 342, 88, 1, 0, 0, 0, 343, 344, 5, 62,
		0, 0, 344, 90, 1, 0, 0, 0, 345, 346, 5, 62, 0, 0, 346, 347, 5, 61, 0, 0,
		347, 92, 1, 0, 0, 0, 348, 349, 5, 61, 0, 0, 349, 94, 1, 0, 0, 0, 350, 351,
		5, 46, 0, 0, 351, 96, 1, 0, 0, 0, 352, 353, 5, 44, 0, 0, 353, 98, 1, 0,
		0, 0, 354, 355, 5, 58, 0, 0, 355, 100, 1, 0, 0, 0, 356, 357, 5, 59, 0,
		0, 357, 102, 1, 0, 0, 0, 358, 359, 5, 40, 0, 0, 359, 104, 1, 0, 0, 0, 360,
		361, 5, 41, 0, 0, 361, 106, 1, 0, 0, 0, 362, 363, 5, 123, 0, 0, 363, 108,
		1, 0, 0, 0, 364, 365, 5, 125, 0, 0, 365, 110, 1, 0, 0, 0, 366, 368, 7,
		5, 0, 0, 367, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 367, 1, 0, 0,
		0, 369, 370, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 372, 6, 55, 0, 0, 372,
		112, 1, 0, 0, 0, 373, 374, 5, 47, 0, 0, 374, 375, 5, 42, 0, 0, 375, 379,
		1, 0, 0, 0, 376, 378, 9, 0, 0, 0, 377, 376, 1, 0, 0, 0, 378, 381, 1, 0,
		0, 0, 379, 380, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 380, 382, 1, 0, 0, 0,
		381, 379, 1, 0, 0, 0, 382, 383, 5, 42, 0, 0, 383, 384, 5, 47, 0, 0, 384,
		385, 1, 0, 0, 0, 385, 386, 6, 56, 0, 0, 386, 114, 1, 0, 0, 0, 387, 388,
		5, 47, 0, 0, 388, 389, 5, 47, 0, 0, 389, 393, 1, 0, 0, 0, 390, 392, 8,
		6, 0, 0, 391, 390, 1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393, 391, 1, 0, 0,
		0, 393, 394, 1, 0, 0, 0, 394, 396, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 396,
		397, 6, 57, 0, 0, 397, 116, 1, 0, 0, 0, 398, 399, 5, 92, 0, 0, 399, 400,
		7, 7, 0, 0, 400, 118, 1, 0, 0, 0, 9, 0, 284, 290, 292, 298, 307, 369, 379,
		393, 1, 6, 0, 0,
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

// SwiftgrammLexerInit initializes any static state used to implement SwiftgrammLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSwiftgrammLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftgrammLexerInit() {
	staticData := &SwiftgrammLexerLexerStaticData
	staticData.once.Do(swiftgrammlexerLexerInit)
}

// NewSwiftgrammLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSwiftgrammLexer(input antlr.CharStream) *SwiftgrammLexer {
	SwiftgrammLexerInit()
	l := new(SwiftgrammLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SwiftgrammLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Swiftgramm.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftgrammLexer tokens.
const (
	SwiftgrammLexerINT                = 1
	SwiftgrammLexerFLOAT              = 2
	SwiftgrammLexerSTRING             = 3
	SwiftgrammLexerBOOL               = 4
	SwiftgrammLexerCHARACTER          = 5
	SwiftgrammLexerTRUE               = 6
	SwiftgrammLexerFALSE              = 7
	SwiftgrammLexerNIL                = 8
	SwiftgrammLexerVAR                = 9
	SwiftgrammLexerLET                = 10
	SwiftgrammLexerPRINT              = 11
	SwiftgrammLexerIF                 = 12
	SwiftgrammLexerELSE               = 13
	SwiftgrammLexerSWITCH             = 14
	SwiftgrammLexerCASE               = 15
	SwiftgrammLexerBREAK              = 16
	SwiftgrammLexerDEFAULT            = 17
	SwiftgrammLexerWHILE              = 18
	SwiftgrammLexerFOR                = 19
	SwiftgrammLexerIN                 = 20
	SwiftgrammLexerGUARD              = 21
	SwiftgrammLexerCONTINUE           = 22
	SwiftgrammLexerRETURN             = 23
	SwiftgrammLexerFUNC               = 24
	SwiftgrammLexerSTRUCT             = 25
	SwiftgrammLexerMUTATING           = 26
	SwiftgrammLexerSELF               = 27
	SwiftgrammLexerINOUT              = 28
	SwiftgrammLexerNUMBER             = 29
	SwiftgrammLexerSTRING_LITERAL     = 30
	SwiftgrammLexerID                 = 31
	SwiftgrammLexerCHARACTER_LITERAL  = 32
	SwiftgrammLexerSUMMATION          = 33
	SwiftgrammLexerSUBTRACTION        = 34
	SwiftgrammLexerMULTIPLICATION     = 35
	SwiftgrammLexerDIVISION           = 36
	SwiftgrammLexerQUESTION_MARK      = 37
	SwiftgrammLexerOR                 = 38
	SwiftgrammLexerAND                = 39
	SwiftgrammLexerNOT                = 40
	SwiftgrammLexerEQUAL              = 41
	SwiftgrammLexerNOT_EQUAL          = 42
	SwiftgrammLexerLESS_THAN          = 43
	SwiftgrammLexerLESS_THAN_EQUAL    = 44
	SwiftgrammLexerGREATER_THAN       = 45
	SwiftgrammLexerGREATER_THAN_EQUAL = 46
	SwiftgrammLexerASSIGN             = 47
	SwiftgrammLexerDOT                = 48
	SwiftgrammLexerCOMMA              = 49
	SwiftgrammLexerCOLON              = 50
	SwiftgrammLexerSEMICOLON          = 51
	SwiftgrammLexerOPEN_PARENTHESIS   = 52
	SwiftgrammLexerCLOSE_PARENTHESIS  = 53
	SwiftgrammLexerOPEN_kEY           = 54
	SwiftgrammLexerCLOSE_kEY          = 55
	SwiftgrammLexerWHITESPACE         = 56
	SwiftgrammLexerCOMMENT            = 57
	SwiftgrammLexerLINE_COMMENT       = 58
)