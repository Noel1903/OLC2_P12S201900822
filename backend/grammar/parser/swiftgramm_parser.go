// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgramm
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "grammar/expressions"
import "grammar/instructions"
import "grammar/symbol"
import "grammar/abstract"

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
		"'inout'", "'append'", "'removeLast'", "'remove'", "'at'", "'IsEmpty'",
		"'count'", "", "", "", "", "", "'+='", "'-='", "'...'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'?'", "'||'", "'&&'", "'!'", "'=='", "'!='", "'<'",
		"'<='", "'>'", "'>='", "'='", "'.'", "','", "':'", "';'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'->'", "'_'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE",
		"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK",
		"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC",
		"STRUCT", "MUTATING", "SELF", "INOUT", "APPEND", "REMOVELAST", "REMOVE",
		"AT", "ISEMPTY", "COUNT", "NUMBER", "FLOATT", "ID", "CHARACTER_LITERAL",
		"STRING_LITERAL", "INCREMENT", "DECREMENT", "RANGE", "SUMMATION", "SUBTRACTION",
		"MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT",
		"EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN",
		"GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON",
		"OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", "OPEN_BRACKET",
		"CLOSE_BRACKET", "ARROW", "UNDERSCORE", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "sentence", "switch_bl", "cases", "increment_bl", "decrement_bl",
		"break_bl", "return_bl", "continue_bl", "declare_let", "native_function",
		"declare_var", "assign_bl", "print_bl", "list_print", "if_bl", "else_if",
		"while_bl", "for_bl", "guard_bl", "vector_bl", "array_exp", "array_functions",
		"assign_vector", "function_bl", "params", "extern_params", "call_function",
		"list_exp", "call_function_exp", "declare_array_bl", "exp_matriz", "type_matrix",
		"definition_matrix", "expression", "datatype",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 73, 729, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 80, 8, 1, 11, 1, 12, 1, 81, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 146, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 173, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 194, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 213, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 3, 11, 233, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 3, 12, 256, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 277, 8, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 3, 16, 305, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 336, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 3, 21, 387, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 397, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 423,
		8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 475,
		8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 537,
		8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 543, 8, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 555, 8, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29,
		577, 8, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 604, 8, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 616,
		8, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 675, 8, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 5, 35, 712, 8, 35, 10, 35, 12, 35, 715, 9, 35, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 727, 8,
		36, 1, 36, 0, 1, 70, 37, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 0, 6, 1, 0, 45, 46, 1, 0, 43, 44, 1, 0, 54, 55,
		1, 0, 56, 57, 1, 0, 52, 53, 1, 0, 49, 50, 772, 0, 74, 1, 0, 0, 0, 2, 79,
		1, 0, 0, 0, 4, 145, 1, 0, 0, 0, 6, 147, 1, 0, 0, 0, 8, 172, 1, 0, 0, 0,
		10, 174, 1, 0, 0, 0, 12, 179, 1, 0, 0, 0, 14, 184, 1, 0, 0, 0, 16, 193,
		1, 0, 0, 0, 18, 195, 1, 0, 0, 0, 20, 212, 1, 0, 0, 0, 22, 232, 1, 0, 0,
		0, 24, 255, 1, 0, 0, 0, 26, 257, 1, 0, 0, 0, 28, 262, 1, 0, 0, 0, 30, 276,
		1, 0, 0, 0, 32, 304, 1, 0, 0, 0, 34, 335, 1, 0, 0, 0, 36, 337, 1, 0, 0,
		0, 38, 344, 1, 0, 0, 0, 40, 355, 1, 0, 0, 0, 42, 386, 1, 0, 0, 0, 44, 396,
		1, 0, 0, 0, 46, 422, 1, 0, 0, 0, 48, 424, 1, 0, 0, 0, 50, 474, 1, 0, 0,
		0, 52, 536, 1, 0, 0, 0, 54, 542, 1, 0, 0, 0, 56, 554, 1, 0, 0, 0, 58, 576,
		1, 0, 0, 0, 60, 578, 1, 0, 0, 0, 62, 581, 1, 0, 0, 0, 64, 603, 1, 0, 0,
		0, 66, 615, 1, 0, 0, 0, 68, 617, 1, 0, 0, 0, 70, 674, 1, 0, 0, 0, 72, 726,
		1, 0, 0, 0, 74, 75, 3, 2, 1, 0, 75, 76, 5, 0, 0, 1, 76, 77, 6, 0, -1, 0,
		77, 1, 1, 0, 0, 0, 78, 80, 3, 4, 2, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0,
		0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84,
		6, 1, -1, 0, 84, 3, 1, 0, 0, 0, 85, 86, 3, 24, 12, 0, 86, 87, 6, 2, -1,
		0, 87, 146, 1, 0, 0, 0, 88, 89, 3, 20, 10, 0, 89, 90, 6, 2, -1, 0, 90,
		146, 1, 0, 0, 0, 91, 92, 3, 26, 13, 0, 92, 93, 6, 2, -1, 0, 93, 146, 1,
		0, 0, 0, 94, 95, 3, 48, 24, 0, 95, 96, 6, 2, -1, 0, 96, 146, 1, 0, 0, 0,
		97, 98, 3, 28, 14, 0, 98, 99, 6, 2, -1, 0, 99, 146, 1, 0, 0, 0, 100, 101,
		3, 32, 16, 0, 101, 102, 6, 2, -1, 0, 102, 146, 1, 0, 0, 0, 103, 104, 3,
		10, 5, 0, 104, 105, 6, 2, -1, 0, 105, 146, 1, 0, 0, 0, 106, 107, 3, 12,
		6, 0, 107, 108, 6, 2, -1, 0, 108, 146, 1, 0, 0, 0, 109, 110, 3, 36, 18,
		0, 110, 111, 6, 2, -1, 0, 111, 146, 1, 0, 0, 0, 112, 113, 3, 38, 19, 0,
		113, 114, 6, 2, -1, 0, 114, 146, 1, 0, 0, 0, 115, 116, 3, 40, 20, 0, 116,
		117, 6, 2, -1, 0, 117, 146, 1, 0, 0, 0, 118, 119, 3, 14, 7, 0, 119, 120,
		6, 2, -1, 0, 120, 146, 1, 0, 0, 0, 121, 122, 3, 16, 8, 0, 122, 123, 6,
		2, -1, 0, 123, 146, 1, 0, 0, 0, 124, 125, 3, 18, 9, 0, 125, 126, 6, 2,
		-1, 0, 126, 146, 1, 0, 0, 0, 127, 128, 3, 42, 21, 0, 128, 129, 6, 2, -1,
		0, 129, 146, 1, 0, 0, 0, 130, 131, 3, 50, 25, 0, 131, 132, 6, 2, -1, 0,
		132, 146, 1, 0, 0, 0, 133, 134, 3, 56, 28, 0, 134, 135, 6, 2, -1, 0, 135,
		146, 1, 0, 0, 0, 136, 137, 3, 46, 23, 0, 137, 138, 6, 2, -1, 0, 138, 146,
		1, 0, 0, 0, 139, 140, 3, 62, 31, 0, 140, 141, 6, 2, -1, 0, 141, 146, 1,
		0, 0, 0, 142, 143, 3, 6, 3, 0, 143, 144, 6, 2, -1, 0, 144, 146, 1, 0, 0,
		0, 145, 85, 1, 0, 0, 0, 145, 88, 1, 0, 0, 0, 145, 91, 1, 0, 0, 0, 145,
		94, 1, 0, 0, 0, 145, 97, 1, 0, 0, 0, 145, 100, 1, 0, 0, 0, 145, 103, 1,
		0, 0, 0, 145, 106, 1, 0, 0, 0, 145, 109, 1, 0, 0, 0, 145, 112, 1, 0, 0,
		0, 145, 115, 1, 0, 0, 0, 145, 118, 1, 0, 0, 0, 145, 121, 1, 0, 0, 0, 145,
		124, 1, 0, 0, 0, 145, 127, 1, 0, 0, 0, 145, 130, 1, 0, 0, 0, 145, 133,
		1, 0, 0, 0, 145, 136, 1, 0, 0, 0, 145, 139, 1, 0, 0, 0, 145, 142, 1, 0,
		0, 0, 146, 5, 1, 0, 0, 0, 147, 148, 5, 14, 0, 0, 148, 149, 3, 70, 35, 0,
		149, 150, 5, 65, 0, 0, 150, 151, 3, 8, 4, 0, 151, 152, 5, 66, 0, 0, 152,
		153, 6, 3, -1, 0, 153, 7, 1, 0, 0, 0, 154, 155, 5, 15, 0, 0, 155, 156,
		3, 70, 35, 0, 156, 157, 5, 61, 0, 0, 157, 158, 3, 2, 1, 0, 158, 159, 3,
		8, 4, 0, 159, 160, 6, 4, -1, 0, 160, 173, 1, 0, 0, 0, 161, 162, 5, 15,
		0, 0, 162, 163, 3, 70, 35, 0, 163, 164, 5, 61, 0, 0, 164, 165, 3, 2, 1,
		0, 165, 166, 6, 4, -1, 0, 166, 173, 1, 0, 0, 0, 167, 168, 5, 17, 0, 0,
		168, 169, 5, 61, 0, 0, 169, 170, 3, 2, 1, 0, 170, 171, 6, 4, -1, 0, 171,
		173, 1, 0, 0, 0, 172, 154, 1, 0, 0, 0, 172, 161, 1, 0, 0, 0, 172, 167,
		1, 0, 0, 0, 173, 9, 1, 0, 0, 0, 174, 175, 5, 37, 0, 0, 175, 176, 5, 40,
		0, 0, 176, 177, 3, 70, 35, 0, 177, 178, 6, 5, -1, 0, 178, 11, 1, 0, 0,
		0, 179, 180, 5, 37, 0, 0, 180, 181, 5, 41, 0, 0, 181, 182, 3, 70, 35, 0,
		182, 183, 6, 6, -1, 0, 183, 13, 1, 0, 0, 0, 184, 185, 5, 16, 0, 0, 185,
		186, 6, 7, -1, 0, 186, 15, 1, 0, 0, 0, 187, 188, 5, 23, 0, 0, 188, 189,
		3, 70, 35, 0, 189, 190, 6, 8, -1, 0, 190, 194, 1, 0, 0, 0, 191, 192, 5,
		23, 0, 0, 192, 194, 6, 8, -1, 0, 193, 187, 1, 0, 0, 0, 193, 191, 1, 0,
		0, 0, 194, 17, 1, 0, 0, 0, 195, 196, 5, 22, 0, 0, 196, 197, 6, 9, -1, 0,
		197, 19, 1, 0, 0, 0, 198, 199, 5, 10, 0, 0, 199, 200, 5, 37, 0, 0, 200,
		201, 5, 61, 0, 0, 201, 202, 3, 72, 36, 0, 202, 203, 5, 58, 0, 0, 203, 204,
		3, 70, 35, 0, 204, 205, 6, 10, -1, 0, 205, 213, 1, 0, 0, 0, 206, 207, 5,
		10, 0, 0, 207, 208, 5, 37, 0, 0, 208, 209, 5, 58, 0, 0, 209, 210, 3, 70,
		35, 0, 210, 211, 6, 10, -1, 0, 211, 213, 1, 0, 0, 0, 212, 198, 1, 0, 0,
		0, 212, 206, 1, 0, 0, 0, 213, 21, 1, 0, 0, 0, 214, 215, 5, 3, 0, 0, 215,
		216, 5, 63, 0, 0, 216, 217, 3, 70, 35, 0, 217, 218, 5, 64, 0, 0, 218, 219,
		6, 11, -1, 0, 219, 233, 1, 0, 0, 0, 220, 221, 5, 1, 0, 0, 221, 222, 5,
		63, 0, 0, 222, 223, 3, 70, 35, 0, 223, 224, 5, 64, 0, 0, 224, 225, 6, 11,
		-1, 0, 225, 233, 1, 0, 0, 0, 226, 227, 5, 2, 0, 0, 227, 228, 5, 63, 0,
		0, 228, 229, 3, 70, 35, 0, 229, 230, 5, 64, 0, 0, 230, 231, 6, 11, -1,
		0, 231, 233, 1, 0, 0, 0, 232, 214, 1, 0, 0, 0, 232, 220, 1, 0, 0, 0, 232,
		226, 1, 0, 0, 0, 233, 23, 1, 0, 0, 0, 234, 235, 5, 9, 0, 0, 235, 236, 5,
		37, 0, 0, 236, 237, 5, 61, 0, 0, 237, 238, 3, 72, 36, 0, 238, 239, 5, 58,
		0, 0, 239, 240, 3, 70, 35, 0, 240, 241, 6, 12, -1, 0, 241, 256, 1, 0, 0,
		0, 242, 243, 5, 9, 0, 0, 243, 244, 5, 37, 0, 0, 244, 245, 5, 58, 0, 0,
		245, 246, 3, 70, 35, 0, 246, 247, 6, 12, -1, 0, 247, 256, 1, 0, 0, 0, 248,
		249, 5, 9, 0, 0, 249, 250, 5, 37, 0, 0, 250, 251, 5, 61, 0, 0, 251, 252,
		3, 72, 36, 0, 252, 253, 5, 48, 0, 0, 253, 254, 6, 12, -1, 0, 254, 256,
		1, 0, 0, 0, 255, 234, 1, 0, 0, 0, 255, 242, 1, 0, 0, 0, 255, 248, 1, 0,
		0, 0, 256, 25, 1, 0, 0, 0, 257, 258, 5, 37, 0, 0, 258, 259, 5, 58, 0, 0,
		259, 260, 3, 70, 35, 0, 260, 261, 6, 13, -1, 0, 261, 27, 1, 0, 0, 0, 262,
		263, 5, 11, 0, 0, 263, 264, 5, 63, 0, 0, 264, 265, 3, 30, 15, 0, 265, 266,
		5, 64, 0, 0, 266, 267, 6, 14, -1, 0, 267, 29, 1, 0, 0, 0, 268, 269, 3,
		70, 35, 0, 269, 270, 5, 60, 0, 0, 270, 271, 3, 30, 15, 0, 271, 272, 6,
		15, -1, 0, 272, 277, 1, 0, 0, 0, 273, 274, 3, 70, 35, 0, 274, 275, 6, 15,
		-1, 0, 275, 277, 1, 0, 0, 0, 276, 268, 1, 0, 0, 0, 276, 273, 1, 0, 0, 0,
		277, 31, 1, 0, 0, 0, 278, 279, 5, 12, 0, 0, 279, 280, 3, 70, 35, 0, 280,
		281, 5, 65, 0, 0, 281, 282, 3, 2, 1, 0, 282, 283, 5, 66, 0, 0, 283, 284,
		6, 16, -1, 0, 284, 305, 1, 0, 0, 0, 285, 286, 5, 12, 0, 0, 286, 287, 3,
		70, 35, 0, 287, 288, 5, 65, 0, 0, 288, 289, 3, 2, 1, 0, 289, 290, 5, 66,
		0, 0, 290, 291, 5, 13, 0, 0, 291, 292, 5, 65, 0, 0, 292, 293, 3, 2, 1,
		0, 293, 294, 5, 66, 0, 0, 294, 295, 6, 16, -1, 0, 295, 305, 1, 0, 0, 0,
		296, 297, 5, 12, 0, 0, 297, 298, 3, 70, 35, 0, 298, 299, 5, 65, 0, 0, 299,
		300, 3, 2, 1, 0, 300, 301, 5, 66, 0, 0, 301, 302, 3, 34, 17, 0, 302, 303,
		6, 16, -1, 0, 303, 305, 1, 0, 0, 0, 304, 278, 1, 0, 0, 0, 304, 285, 1,
		0, 0, 0, 304, 296, 1, 0, 0, 0, 305, 33, 1, 0, 0, 0, 306, 307, 5, 13, 0,
		0, 307, 308, 5, 12, 0, 0, 308, 309, 3, 70, 35, 0, 309, 310, 5, 65, 0, 0,
		310, 311, 3, 2, 1, 0, 311, 312, 5, 66, 0, 0, 312, 313, 6, 17, -1, 0, 313,
		336, 1, 0, 0, 0, 314, 315, 5, 13, 0, 0, 315, 316, 5, 12, 0, 0, 316, 317,
		3, 70, 35, 0, 317, 318, 5, 65, 0, 0, 318, 319, 3, 2, 1, 0, 319, 320, 5,
		66, 0, 0, 320, 321, 5, 13, 0, 0, 321, 322, 5, 65, 0, 0, 322, 323, 3, 2,
		1, 0, 323, 324, 5, 66, 0, 0, 324, 325, 6, 17, -1, 0, 325, 336, 1, 0, 0,
		0, 326, 327, 5, 13, 0, 0, 327, 328, 5, 12, 0, 0, 328, 329, 3, 70, 35, 0,
		329, 330, 5, 65, 0, 0, 330, 331, 3, 2, 1, 0, 331, 332, 5, 66, 0, 0, 332,
		333, 3, 34, 17, 0, 333, 334, 6, 17, -1, 0, 334, 336, 1, 0, 0, 0, 335, 306,
		1, 0, 0, 0, 335, 314, 1, 0, 0, 0, 335, 326, 1, 0, 0, 0, 336, 35, 1, 0,
		0, 0, 337, 338, 5, 18, 0, 0, 338, 339, 3, 70, 35, 0, 339, 340, 5, 65, 0,
		0, 340, 341, 3, 2, 1, 0, 341, 342, 5, 66, 0, 0, 342, 343, 6, 18, -1, 0,
		343, 37, 1, 0, 0, 0, 344, 345, 5, 19, 0, 0, 345, 346, 5, 37, 0, 0, 346,
		347, 5, 20, 0, 0, 347, 348, 3, 70, 35, 0, 348, 349, 5, 42, 0, 0, 349, 350,
		3, 70, 35, 0, 350, 351, 5, 65, 0, 0, 351, 352, 3, 2, 1, 0, 352, 353, 5,
		66, 0, 0, 353, 354, 6, 19, -1, 0, 354, 39, 1, 0, 0, 0, 355, 356, 5, 21,
		0, 0, 356, 357, 3, 70, 35, 0, 357, 358, 5, 13, 0, 0, 358, 359, 5, 65, 0,
		0, 359, 360, 3, 2, 1, 0, 360, 361, 5, 66, 0, 0, 361, 362, 6, 20, -1, 0,
		362, 41, 1, 0, 0, 0, 363, 364, 5, 9, 0, 0, 364, 365, 5, 37, 0, 0, 365,
		366, 5, 61, 0, 0, 366, 367, 5, 67, 0, 0, 367, 368, 3, 72, 36, 0, 368, 369,
		5, 68, 0, 0, 369, 370, 5, 58, 0, 0, 370, 371, 5, 67, 0, 0, 371, 372, 3,
		44, 22, 0, 372, 373, 5, 68, 0, 0, 373, 374, 6, 21, -1, 0, 374, 387, 1,
		0, 0, 0, 375, 376, 5, 9, 0, 0, 376, 377, 5, 37, 0, 0, 377, 378, 5, 61,
		0, 0, 378, 379, 5, 67, 0, 0, 379, 380, 3, 72, 36, 0, 380, 381, 5, 68, 0,
		0, 381, 382, 5, 58, 0, 0, 382, 383, 5, 67, 0, 0, 383, 384, 5, 68, 0, 0,
		384, 385, 6, 21, -1, 0, 385, 387, 1, 0, 0, 0, 386, 363, 1, 0, 0, 0, 386,
		375, 1, 0, 0, 0, 387, 43, 1, 0, 0, 0, 388, 389, 3, 70, 35, 0, 389, 390,
		5, 60, 0, 0, 390, 391, 3, 44, 22, 0, 391, 392, 6, 22, -1, 0, 392, 397,
		1, 0, 0, 0, 393, 394, 3, 70, 35, 0, 394, 395, 6, 22, -1, 0, 395, 397, 1,
		0, 0, 0, 396, 388, 1, 0, 0, 0, 396, 393, 1, 0, 0, 0, 397, 45, 1, 0, 0,
		0, 398, 399, 5, 37, 0, 0, 399, 400, 5, 59, 0, 0, 400, 401, 5, 29, 0, 0,
		401, 402, 5, 63, 0, 0, 402, 403, 3, 70, 35, 0, 403, 404, 5, 64, 0, 0, 404,
		405, 6, 23, -1, 0, 405, 423, 1, 0, 0, 0, 406, 407, 5, 37, 0, 0, 407, 408,
		5, 59, 0, 0, 408, 409, 5, 30, 0, 0, 409, 410, 5, 63, 0, 0, 410, 411, 5,
		64, 0, 0, 411, 423, 6, 23, -1, 0, 412, 413, 5, 37, 0, 0, 413, 414, 5, 59,
		0, 0, 414, 415, 5, 31, 0, 0, 415, 416, 5, 63, 0, 0, 416, 417, 5, 32, 0,
		0, 417, 418, 5, 61, 0, 0, 418, 419, 3, 70, 35, 0, 419, 420, 5, 64, 0, 0,
		420, 421, 6, 23, -1, 0, 421, 423, 1, 0, 0, 0, 422, 398, 1, 0, 0, 0, 422,
		406, 1, 0, 0, 0, 422, 412, 1, 0, 0, 0, 423, 47, 1, 0, 0, 0, 424, 425, 5,
		37, 0, 0, 425, 426, 5, 67, 0, 0, 426, 427, 3, 70, 35, 0, 427, 428, 5, 68,
		0, 0, 428, 429, 5, 58, 0, 0, 429, 430, 3, 70, 35, 0, 430, 431, 6, 24, -1,
		0, 431, 49, 1, 0, 0, 0, 432, 433, 5, 24, 0, 0, 433, 434, 5, 37, 0, 0, 434,
		435, 5, 63, 0, 0, 435, 436, 5, 64, 0, 0, 436, 437, 5, 69, 0, 0, 437, 438,
		3, 72, 36, 0, 438, 439, 5, 65, 0, 0, 439, 440, 3, 2, 1, 0, 440, 441, 5,
		66, 0, 0, 441, 442, 6, 25, -1, 0, 442, 475, 1, 0, 0, 0, 443, 444, 5, 24,
		0, 0, 444, 445, 5, 37, 0, 0, 445, 446, 5, 63, 0, 0, 446, 447, 5, 64, 0,
		0, 447, 448, 5, 65, 0, 0, 448, 449, 3, 2, 1, 0, 449, 450, 5, 66, 0, 0,
		450, 451, 6, 25, -1, 0, 451, 475, 1, 0, 0, 0, 452, 453, 5, 24, 0, 0, 453,
		454, 5, 37, 0, 0, 454, 455, 5, 63, 0, 0, 455, 456, 3, 52, 26, 0, 456, 457,
		5, 64, 0, 0, 457, 458, 5, 69, 0, 0, 458, 459, 3, 72, 36, 0, 459, 460, 5,
		65, 0, 0, 460, 461, 3, 2, 1, 0, 461, 462, 5, 66, 0, 0, 462, 463, 6, 25,
		-1, 0, 463, 475, 1, 0, 0, 0, 464, 465, 5, 24, 0, 0, 465, 466, 5, 37, 0,
		0, 466, 467, 5, 63, 0, 0, 467, 468, 3, 52, 26, 0, 468, 469, 5, 64, 0, 0,
		469, 470, 5, 65, 0, 0, 470, 471, 3, 2, 1, 0, 471, 472, 5, 66, 0, 0, 472,
		473, 6, 25, -1, 0, 473, 475, 1, 0, 0, 0, 474, 432, 1, 0, 0, 0, 474, 443,
		1, 0, 0, 0, 474, 452, 1, 0, 0, 0, 474, 464, 1, 0, 0, 0, 475, 51, 1, 0,
		0, 0, 476, 477, 5, 37, 0, 0, 477, 478, 5, 61, 0, 0, 478, 479, 3, 72, 36,
		0, 479, 480, 5, 60, 0, 0, 480, 481, 3, 52, 26, 0, 481, 482, 6, 26, -1,
		0, 482, 537, 1, 0, 0, 0, 483, 484, 5, 37, 0, 0, 484, 485, 5, 61, 0, 0,
		485, 486, 3, 72, 36, 0, 486, 487, 6, 26, -1, 0, 487, 537, 1, 0, 0, 0, 488,
		489, 3, 54, 27, 0, 489, 490, 5, 37, 0, 0, 490, 491, 5, 61, 0, 0, 491, 492,
		3, 72, 36, 0, 492, 493, 5, 60, 0, 0, 493, 494, 3, 52, 26, 0, 494, 495,
		6, 26, -1, 0, 495, 537, 1, 0, 0, 0, 496, 497, 3, 54, 27, 0, 497, 498, 5,
		37, 0, 0, 498, 499, 5, 61, 0, 0, 499, 500, 3, 72, 36, 0, 500, 501, 6, 26,
		-1, 0, 501, 537, 1, 0, 0, 0, 502, 503, 5, 37, 0, 0, 503, 504, 5, 61, 0,
		0, 504, 505, 5, 67, 0, 0, 505, 506, 3, 72, 36, 0, 506, 507, 5, 68, 0, 0,
		507, 508, 5, 60, 0, 0, 508, 509, 3, 52, 26, 0, 509, 510, 6, 26, -1, 0,
		510, 537, 1, 0, 0, 0, 511, 512, 5, 37, 0, 0, 512, 513, 5, 61, 0, 0, 513,
		514, 5, 67, 0, 0, 514, 515, 3, 72, 36, 0, 515, 516, 5, 68, 0, 0, 516, 517,
		6, 26, -1, 0, 517, 537, 1, 0, 0, 0, 518, 519, 3, 54, 27, 0, 519, 520, 5,
		37, 0, 0, 520, 521, 5, 61, 0, 0, 521, 522, 5, 67, 0, 0, 522, 523, 3, 72,
		36, 0, 523, 524, 5, 68, 0, 0, 524, 525, 5, 60, 0, 0, 525, 526, 3, 52, 26,
		0, 526, 527, 6, 26, -1, 0, 527, 537, 1, 0, 0, 0, 528, 529, 3, 54, 27, 0,
		529, 530, 5, 37, 0, 0, 530, 531, 5, 61, 0, 0, 531, 532, 5, 67, 0, 0, 532,
		533, 3, 72, 36, 0, 533, 534, 5, 68, 0, 0, 534, 535, 6, 26, -1, 0, 535,
		537, 1, 0, 0, 0, 536, 476, 1, 0, 0, 0, 536, 483, 1, 0, 0, 0, 536, 488,
		1, 0, 0, 0, 536, 496, 1, 0, 0, 0, 536, 502, 1, 0, 0, 0, 536, 511, 1, 0,
		0, 0, 536, 518, 1, 0, 0, 0, 536, 528, 1, 0, 0, 0, 537, 53, 1, 0, 0, 0,
		538, 539, 5, 37, 0, 0, 539, 543, 6, 27, -1, 0, 540, 541, 5, 70, 0, 0, 541,
		543, 6, 27, -1, 0, 542, 538, 1, 0, 0, 0, 542, 540, 1, 0, 0, 0, 543, 55,
		1, 0, 0, 0, 544, 545, 5, 37, 0, 0, 545, 546, 5, 63, 0, 0, 546, 547, 5,
		64, 0, 0, 547, 555, 6, 28, -1, 0, 548, 549, 5, 37, 0, 0, 549, 550, 5, 63,
		0, 0, 550, 551, 3, 58, 29, 0, 551, 552, 5, 64, 0, 0, 552, 553, 6, 28, -1,
		0, 553, 555, 1, 0, 0, 0, 554, 544, 1, 0, 0, 0, 554, 548, 1, 0, 0, 0, 555,
		57, 1, 0, 0, 0, 556, 557, 3, 70, 35, 0, 557, 558, 5, 60, 0, 0, 558, 559,
		3, 58, 29, 0, 559, 560, 6, 29, -1, 0, 560, 577, 1, 0, 0, 0, 561, 562, 3,
		70, 35, 0, 562, 563, 6, 29, -1, 0, 563, 577, 1, 0, 0, 0, 564, 565, 5, 37,
		0, 0, 565, 566, 5, 61, 0, 0, 566, 567, 3, 70, 35, 0, 567, 568, 5, 60, 0,
		0, 568, 569, 3, 58, 29, 0, 569, 570, 6, 29, -1, 0, 570, 577, 1, 0, 0, 0,
		571, 572, 5, 37, 0, 0, 572, 573, 5, 61, 0, 0, 573, 574, 3, 70, 35, 0, 574,
		575, 6, 29, -1, 0, 575, 577, 1, 0, 0, 0, 576, 556, 1, 0, 0, 0, 576, 561,
		1, 0, 0, 0, 576, 564, 1, 0, 0, 0, 576, 571, 1, 0, 0, 0, 577, 59, 1, 0,
		0, 0, 578, 579, 3, 56, 28, 0, 579, 580, 6, 30, -1, 0, 580, 61, 1, 0, 0,
		0, 581, 582, 5, 9, 0, 0, 582, 583, 5, 37, 0, 0, 583, 584, 5, 61, 0, 0,
		584, 585, 3, 66, 33, 0, 585, 586, 5, 58, 0, 0, 586, 587, 5, 67, 0, 0, 587,
		588, 3, 64, 32, 0, 588, 589, 5, 68, 0, 0, 589, 590, 6, 31, -1, 0, 590,
		63, 1, 0, 0, 0, 591, 592, 5, 67, 0, 0, 592, 593, 3, 70, 35, 0, 593, 594,
		5, 68, 0, 0, 594, 595, 6, 32, -1, 0, 595, 604, 1, 0, 0, 0, 596, 597, 5,
		67, 0, 0, 597, 598, 3, 70, 35, 0, 598, 599, 5, 68, 0, 0, 599, 600, 5, 60,
		0, 0, 600, 601, 3, 64, 32, 0, 601, 602, 6, 32, -1, 0, 602, 604, 1, 0, 0,
		0, 603, 591, 1, 0, 0, 0, 603, 596, 1, 0, 0, 0, 604, 65, 1, 0, 0, 0, 605,
		606, 5, 67, 0, 0, 606, 607, 3, 66, 33, 0, 607, 608, 5, 68, 0, 0, 608, 609,
		6, 33, -1, 0, 609, 616, 1, 0, 0, 0, 610, 611, 5, 67, 0, 0, 611, 612, 3,
		72, 36, 0, 612, 613, 5, 68, 0, 0, 613, 614, 6, 33, -1, 0, 614, 616, 1,
		0, 0, 0, 615, 605, 1, 0, 0, 0, 615, 610, 1, 0, 0, 0, 616, 67, 1, 0, 0,
		0, 617, 618, 3, 66, 33, 0, 618, 69, 1, 0, 0, 0, 619, 620, 6, 35, -1, 0,
		620, 621, 5, 51, 0, 0, 621, 622, 3, 70, 35, 17, 622, 623, 6, 35, -1, 0,
		623, 675, 1, 0, 0, 0, 624, 625, 5, 44, 0, 0, 625, 626, 3, 70, 35, 16, 626,
		627, 6, 35, -1, 0, 627, 675, 1, 0, 0, 0, 628, 629, 5, 63, 0, 0, 629, 630,
		3, 70, 35, 0, 630, 631, 5, 64, 0, 0, 631, 632, 6, 35, -1, 0, 632, 675,
		1, 0, 0, 0, 633, 634, 3, 60, 30, 0, 634, 635, 6, 35, -1, 0, 635, 675, 1,
		0, 0, 0, 636, 637, 3, 22, 11, 0, 637, 638, 6, 35, -1, 0, 638, 675, 1, 0,
		0, 0, 639, 640, 5, 37, 0, 0, 640, 641, 5, 67, 0, 0, 641, 642, 3, 70, 35,
		0, 642, 643, 5, 68, 0, 0, 643, 644, 6, 35, -1, 0, 644, 675, 1, 0, 0, 0,
		645, 646, 5, 37, 0, 0, 646, 647, 5, 59, 0, 0, 647, 648, 5, 34, 0, 0, 648,
		675, 6, 35, -1, 0, 649, 650, 5, 37, 0, 0, 650, 651, 5, 59, 0, 0, 651, 652,
		5, 33, 0, 0, 652, 675, 6, 35, -1, 0, 653, 654, 5, 67, 0, 0, 654, 655, 3,
		44, 22, 0, 655, 656, 5, 68, 0, 0, 656, 657, 6, 35, -1, 0, 657, 675, 1,
		0, 0, 0, 658, 659, 5, 35, 0, 0, 659, 675, 6, 35, -1, 0, 660, 661, 5, 36,
		0, 0, 661, 675, 6, 35, -1, 0, 662, 663, 5, 39, 0, 0, 663, 675, 6, 35, -1,
		0, 664, 665, 5, 38, 0, 0, 665, 675, 6, 35, -1, 0, 666, 667, 5, 6, 0, 0,
		667, 675, 6, 35, -1, 0, 668, 669, 5, 7, 0, 0, 669, 675, 6, 35, -1, 0, 670,
		671, 5, 37, 0, 0, 671, 675, 6, 35, -1, 0, 672, 673, 5, 8, 0, 0, 673, 675,
		6, 35, -1, 0, 674, 619, 1, 0, 0, 0, 674, 624, 1, 0, 0, 0, 674, 628, 1,
		0, 0, 0, 674, 633, 1, 0, 0, 0, 674, 636, 1, 0, 0, 0, 674, 639, 1, 0, 0,
		0, 674, 645, 1, 0, 0, 0, 674, 649, 1, 0, 0, 0, 674, 653, 1, 0, 0, 0, 674,
		658, 1, 0, 0, 0, 674, 660, 1, 0, 0, 0, 674, 662, 1, 0, 0, 0, 674, 664,
		1, 0, 0, 0, 674, 666, 1, 0, 0, 0, 674, 668, 1, 0, 0, 0, 674, 670, 1, 0,
		0, 0, 674, 672, 1, 0, 0, 0, 675, 713, 1, 0, 0, 0, 676, 677, 10, 24, 0,
		0, 677, 678, 7, 0, 0, 0, 678, 679, 3, 70, 35, 25, 679, 680, 6, 35, -1,
		0, 680, 712, 1, 0, 0, 0, 681, 682, 10, 23, 0, 0, 682, 683, 7, 1, 0, 0,
		683, 684, 3, 70, 35, 24, 684, 685, 6, 35, -1, 0, 685, 712, 1, 0, 0, 0,
		686, 687, 10, 22, 0, 0, 687, 688, 5, 47, 0, 0, 688, 689, 3, 70, 35, 23,
		689, 690, 6, 35, -1, 0, 690, 712, 1, 0, 0, 0, 691, 692, 10, 21, 0, 0, 692,
		693, 7, 2, 0, 0, 693, 694, 3, 70, 35, 22, 694, 695, 6, 35, -1, 0, 695,
		712, 1, 0, 0, 0, 696, 697, 10, 20, 0, 0, 697, 698, 7, 3, 0, 0, 698, 699,
		3, 70, 35, 21, 699, 700, 6, 35, -1, 0, 700, 712, 1, 0, 0, 0, 701, 702,
		10, 19, 0, 0, 702, 703, 7, 4, 0, 0, 703, 704, 3, 70, 35, 20, 704, 705,
		6, 35, -1, 0, 705, 712, 1, 0, 0, 0, 706, 707, 10, 18, 0, 0, 707, 708, 7,
		5, 0, 0, 708, 709, 3, 70, 35, 19, 709, 710, 6, 35, -1, 0, 710, 712, 1,
		0, 0, 0, 711, 676, 1, 0, 0, 0, 711, 681, 1, 0, 0, 0, 711, 686, 1, 0, 0,
		0, 711, 691, 1, 0, 0, 0, 711, 696, 1, 0, 0, 0, 711, 701, 1, 0, 0, 0, 711,
		706, 1, 0, 0, 0, 712, 715, 1, 0, 0, 0, 713, 711, 1, 0, 0, 0, 713, 714,
		1, 0, 0, 0, 714, 71, 1, 0, 0, 0, 715, 713, 1, 0, 0, 0, 716, 717, 5, 1,
		0, 0, 717, 727, 6, 36, -1, 0, 718, 719, 5, 2, 0, 0, 719, 727, 6, 36, -1,
		0, 720, 721, 5, 3, 0, 0, 721, 727, 6, 36, -1, 0, 722, 723, 5, 4, 0, 0,
		723, 727, 6, 36, -1, 0, 724, 725, 5, 5, 0, 0, 725, 727, 6, 36, -1, 0, 726,
		716, 1, 0, 0, 0, 726, 718, 1, 0, 0, 0, 726, 720, 1, 0, 0, 0, 726, 722,
		1, 0, 0, 0, 726, 724, 1, 0, 0, 0, 727, 73, 1, 0, 0, 0, 24, 81, 145, 172,
		193, 212, 232, 255, 276, 304, 335, 386, 396, 422, 474, 536, 542, 554, 576,
		603, 615, 674, 711, 713, 726,
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

// Note that '@members' cannot be changed now, but this should have been 'globals'
// If you are looking to have variables for each instance, use '@structmembers'

var CountM int = 0

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
	SwiftgrammParserAPPEND             = 29
	SwiftgrammParserREMOVELAST         = 30
	SwiftgrammParserREMOVE             = 31
	SwiftgrammParserAT                 = 32
	SwiftgrammParserISEMPTY            = 33
	SwiftgrammParserCOUNT              = 34
	SwiftgrammParserNUMBER             = 35
	SwiftgrammParserFLOATT             = 36
	SwiftgrammParserID                 = 37
	SwiftgrammParserCHARACTER_LITERAL  = 38
	SwiftgrammParserSTRING_LITERAL     = 39
	SwiftgrammParserINCREMENT          = 40
	SwiftgrammParserDECREMENT          = 41
	SwiftgrammParserRANGE              = 42
	SwiftgrammParserSUMMATION          = 43
	SwiftgrammParserSUBTRACTION        = 44
	SwiftgrammParserMULTIPLICATION     = 45
	SwiftgrammParserDIVISION           = 46
	SwiftgrammParserMOD                = 47
	SwiftgrammParserQUESTION_MARK      = 48
	SwiftgrammParserOR                 = 49
	SwiftgrammParserAND                = 50
	SwiftgrammParserNOT                = 51
	SwiftgrammParserEQUAL              = 52
	SwiftgrammParserNOT_EQUAL          = 53
	SwiftgrammParserLESS_THAN          = 54
	SwiftgrammParserLESS_THAN_EQUAL    = 55
	SwiftgrammParserGREATER_THAN       = 56
	SwiftgrammParserGREATER_THAN_EQUAL = 57
	SwiftgrammParserASSIGN             = 58
	SwiftgrammParserDOT                = 59
	SwiftgrammParserCOMMA              = 60
	SwiftgrammParserCOLON              = 61
	SwiftgrammParserSEMICOLON          = 62
	SwiftgrammParserOPEN_PARENTHESIS   = 63
	SwiftgrammParserCLOSE_PARENTHESIS  = 64
	SwiftgrammParserOPEN_kEY           = 65
	SwiftgrammParserCLOSE_kEY          = 66
	SwiftgrammParserOPEN_BRACKET       = 67
	SwiftgrammParserCLOSE_BRACKET      = 68
	SwiftgrammParserARROW              = 69
	SwiftgrammParserUNDERSCORE         = 70
	SwiftgrammParserWHITESPACE         = 71
	SwiftgrammParserCOMMENT            = 72
	SwiftgrammParserLINE_COMMENT       = 73
)

// SwiftgrammParser rules.
const (
	SwiftgrammParserRULE_s                 = 0
	SwiftgrammParserRULE_block             = 1
	SwiftgrammParserRULE_sentence          = 2
	SwiftgrammParserRULE_switch_bl         = 3
	SwiftgrammParserRULE_cases             = 4
	SwiftgrammParserRULE_increment_bl      = 5
	SwiftgrammParserRULE_decrement_bl      = 6
	SwiftgrammParserRULE_break_bl          = 7
	SwiftgrammParserRULE_return_bl         = 8
	SwiftgrammParserRULE_continue_bl       = 9
	SwiftgrammParserRULE_declare_let       = 10
	SwiftgrammParserRULE_native_function   = 11
	SwiftgrammParserRULE_declare_var       = 12
	SwiftgrammParserRULE_assign_bl         = 13
	SwiftgrammParserRULE_print_bl          = 14
	SwiftgrammParserRULE_list_print        = 15
	SwiftgrammParserRULE_if_bl             = 16
	SwiftgrammParserRULE_else_if           = 17
	SwiftgrammParserRULE_while_bl          = 18
	SwiftgrammParserRULE_for_bl            = 19
	SwiftgrammParserRULE_guard_bl          = 20
	SwiftgrammParserRULE_vector_bl         = 21
	SwiftgrammParserRULE_array_exp         = 22
	SwiftgrammParserRULE_array_functions   = 23
	SwiftgrammParserRULE_assign_vector     = 24
	SwiftgrammParserRULE_function_bl       = 25
	SwiftgrammParserRULE_params            = 26
	SwiftgrammParserRULE_extern_params     = 27
	SwiftgrammParserRULE_call_function     = 28
	SwiftgrammParserRULE_list_exp          = 29
	SwiftgrammParserRULE_call_function_exp = 30
	SwiftgrammParserRULE_declare_array_bl  = 31
	SwiftgrammParserRULE_exp_matriz        = 32
	SwiftgrammParserRULE_type_matrix       = 33
	SwiftgrammParserRULE_definition_matrix = 34
	SwiftgrammParserRULE_expression        = 35
	SwiftgrammParserRULE_datatype          = 36
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
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

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

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
		p.SetState(74)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(75)
		p.Match(SwiftgrammParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

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

	// Get_sentence returns the _sentence rule contexts.
	Get_sentence() ISentenceContext

	// Set_sentence sets the _sentence rule contexts.
	Set_sentence(ISentenceContext)

	// GetInstr returns the instr rule context list.
	GetInstr() []ISentenceContext

	// SetInstr sets the instr rule context list.
	SetInstr([]ISentenceContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllSentence() []ISentenceContext
	Sentence(i int) ISentenceContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	blk       []interface{}
	_sentence ISentenceContext
	instr     []ISentenceContext
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

func (s *BlockContext) Get_sentence() ISentenceContext { return s._sentence }

func (s *BlockContext) Set_sentence(v ISentenceContext) { s._sentence = v }

func (s *BlockContext) GetInstr() []ISentenceContext { return s.instr }

func (s *BlockContext) SetInstr(v []ISentenceContext) { s.instr = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllSentence() []ISentenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentenceContext); ok {
			len++
		}
	}

	tst := make([]ISentenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentenceContext); ok {
			tst[i] = t.(ISentenceContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Sentence(i int) ISentenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenceContext); ok {
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

	return t.(ISentenceContext)
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

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []ISentenceContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137471286784) != 0) {
		{
			p.SetState(78)

			var _x = p.Sentence()

			localctx.(*BlockContext)._sentence = _x
		}
		localctx.(*BlockContext).instr = append(localctx.(*BlockContext).instr, localctx.(*BlockContext)._sentence)

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockContext).GetInstr()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInstr())
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

	// Get_declare_var returns the _declare_var rule contexts.
	Get_declare_var() IDeclare_varContext

	// Get_declare_let returns the _declare_let rule contexts.
	Get_declare_let() IDeclare_letContext

	// Get_assign_bl returns the _assign_bl rule contexts.
	Get_assign_bl() IAssign_blContext

	// Get_assign_vector returns the _assign_vector rule contexts.
	Get_assign_vector() IAssign_vectorContext

	// Get_print_bl returns the _print_bl rule contexts.
	Get_print_bl() IPrint_blContext

	// Get_if_bl returns the _if_bl rule contexts.
	Get_if_bl() IIf_blContext

	// Get_increment_bl returns the _increment_bl rule contexts.
	Get_increment_bl() IIncrement_blContext

	// Get_decrement_bl returns the _decrement_bl rule contexts.
	Get_decrement_bl() IDecrement_blContext

	// Get_while_bl returns the _while_bl rule contexts.
	Get_while_bl() IWhile_blContext

	// Get_for_bl returns the _for_bl rule contexts.
	Get_for_bl() IFor_blContext

	// Get_guard_bl returns the _guard_bl rule contexts.
	Get_guard_bl() IGuard_blContext

	// Get_break_bl returns the _break_bl rule contexts.
	Get_break_bl() IBreak_blContext

	// Get_return_bl returns the _return_bl rule contexts.
	Get_return_bl() IReturn_blContext

	// Get_continue_bl returns the _continue_bl rule contexts.
	Get_continue_bl() IContinue_blContext

	// Get_vector_bl returns the _vector_bl rule contexts.
	Get_vector_bl() IVector_blContext

	// Get_function_bl returns the _function_bl rule contexts.
	Get_function_bl() IFunction_blContext

	// Get_call_function returns the _call_function rule contexts.
	Get_call_function() ICall_functionContext

	// Get_array_functions returns the _array_functions rule contexts.
	Get_array_functions() IArray_functionsContext

	// Get_declare_array_bl returns the _declare_array_bl rule contexts.
	Get_declare_array_bl() IDeclare_array_blContext

	// Get_switch_bl returns the _switch_bl rule contexts.
	Get_switch_bl() ISwitch_blContext

	// Set_declare_var sets the _declare_var rule contexts.
	Set_declare_var(IDeclare_varContext)

	// Set_declare_let sets the _declare_let rule contexts.
	Set_declare_let(IDeclare_letContext)

	// Set_assign_bl sets the _assign_bl rule contexts.
	Set_assign_bl(IAssign_blContext)

	// Set_assign_vector sets the _assign_vector rule contexts.
	Set_assign_vector(IAssign_vectorContext)

	// Set_print_bl sets the _print_bl rule contexts.
	Set_print_bl(IPrint_blContext)

	// Set_if_bl sets the _if_bl rule contexts.
	Set_if_bl(IIf_blContext)

	// Set_increment_bl sets the _increment_bl rule contexts.
	Set_increment_bl(IIncrement_blContext)

	// Set_decrement_bl sets the _decrement_bl rule contexts.
	Set_decrement_bl(IDecrement_blContext)

	// Set_while_bl sets the _while_bl rule contexts.
	Set_while_bl(IWhile_blContext)

	// Set_for_bl sets the _for_bl rule contexts.
	Set_for_bl(IFor_blContext)

	// Set_guard_bl sets the _guard_bl rule contexts.
	Set_guard_bl(IGuard_blContext)

	// Set_break_bl sets the _break_bl rule contexts.
	Set_break_bl(IBreak_blContext)

	// Set_return_bl sets the _return_bl rule contexts.
	Set_return_bl(IReturn_blContext)

	// Set_continue_bl sets the _continue_bl rule contexts.
	Set_continue_bl(IContinue_blContext)

	// Set_vector_bl sets the _vector_bl rule contexts.
	Set_vector_bl(IVector_blContext)

	// Set_function_bl sets the _function_bl rule contexts.
	Set_function_bl(IFunction_blContext)

	// Set_call_function sets the _call_function rule contexts.
	Set_call_function(ICall_functionContext)

	// Set_array_functions sets the _array_functions rule contexts.
	Set_array_functions(IArray_functionsContext)

	// Set_declare_array_bl sets the _declare_array_bl rule contexts.
	Set_declare_array_bl(IDeclare_array_blContext)

	// Set_switch_bl sets the _switch_bl rule contexts.
	Set_switch_bl(ISwitch_blContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	Declare_var() IDeclare_varContext
	Declare_let() IDeclare_letContext
	Assign_bl() IAssign_blContext
	Assign_vector() IAssign_vectorContext
	Print_bl() IPrint_blContext
	If_bl() IIf_blContext
	Increment_bl() IIncrement_blContext
	Decrement_bl() IDecrement_blContext
	While_bl() IWhile_blContext
	For_bl() IFor_blContext
	Guard_bl() IGuard_blContext
	Break_bl() IBreak_blContext
	Return_bl() IReturn_blContext
	Continue_bl() IContinue_blContext
	Vector_bl() IVector_blContext
	Function_bl() IFunction_blContext
	Call_function() ICall_functionContext
	Array_functions() IArray_functionsContext
	Declare_array_bl() IDeclare_array_blContext
	Switch_bl() ISwitch_blContext

	// IsSentenceContext differentiates from other interfaces.
	IsSentenceContext()
}

type SentenceContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	instr             abstract.Instruction
	_declare_var      IDeclare_varContext
	_declare_let      IDeclare_letContext
	_assign_bl        IAssign_blContext
	_assign_vector    IAssign_vectorContext
	_print_bl         IPrint_blContext
	_if_bl            IIf_blContext
	_increment_bl     IIncrement_blContext
	_decrement_bl     IDecrement_blContext
	_while_bl         IWhile_blContext
	_for_bl           IFor_blContext
	_guard_bl         IGuard_blContext
	_break_bl         IBreak_blContext
	_return_bl        IReturn_blContext
	_continue_bl      IContinue_blContext
	_vector_bl        IVector_blContext
	_function_bl      IFunction_blContext
	_call_function    ICall_functionContext
	_array_functions  IArray_functionsContext
	_declare_array_bl IDeclare_array_blContext
	_switch_bl        ISwitch_blContext
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

func (s *SentenceContext) Get_declare_var() IDeclare_varContext { return s._declare_var }

func (s *SentenceContext) Get_declare_let() IDeclare_letContext { return s._declare_let }

func (s *SentenceContext) Get_assign_bl() IAssign_blContext { return s._assign_bl }

func (s *SentenceContext) Get_assign_vector() IAssign_vectorContext { return s._assign_vector }

func (s *SentenceContext) Get_print_bl() IPrint_blContext { return s._print_bl }

func (s *SentenceContext) Get_if_bl() IIf_blContext { return s._if_bl }

func (s *SentenceContext) Get_increment_bl() IIncrement_blContext { return s._increment_bl }

func (s *SentenceContext) Get_decrement_bl() IDecrement_blContext { return s._decrement_bl }

func (s *SentenceContext) Get_while_bl() IWhile_blContext { return s._while_bl }

func (s *SentenceContext) Get_for_bl() IFor_blContext { return s._for_bl }

func (s *SentenceContext) Get_guard_bl() IGuard_blContext { return s._guard_bl }

func (s *SentenceContext) Get_break_bl() IBreak_blContext { return s._break_bl }

func (s *SentenceContext) Get_return_bl() IReturn_blContext { return s._return_bl }

func (s *SentenceContext) Get_continue_bl() IContinue_blContext { return s._continue_bl }

func (s *SentenceContext) Get_vector_bl() IVector_blContext { return s._vector_bl }

func (s *SentenceContext) Get_function_bl() IFunction_blContext { return s._function_bl }

func (s *SentenceContext) Get_call_function() ICall_functionContext { return s._call_function }

func (s *SentenceContext) Get_array_functions() IArray_functionsContext { return s._array_functions }

func (s *SentenceContext) Get_declare_array_bl() IDeclare_array_blContext { return s._declare_array_bl }

func (s *SentenceContext) Get_switch_bl() ISwitch_blContext { return s._switch_bl }

func (s *SentenceContext) Set_declare_var(v IDeclare_varContext) { s._declare_var = v }

func (s *SentenceContext) Set_declare_let(v IDeclare_letContext) { s._declare_let = v }

func (s *SentenceContext) Set_assign_bl(v IAssign_blContext) { s._assign_bl = v }

func (s *SentenceContext) Set_assign_vector(v IAssign_vectorContext) { s._assign_vector = v }

func (s *SentenceContext) Set_print_bl(v IPrint_blContext) { s._print_bl = v }

func (s *SentenceContext) Set_if_bl(v IIf_blContext) { s._if_bl = v }

func (s *SentenceContext) Set_increment_bl(v IIncrement_blContext) { s._increment_bl = v }

func (s *SentenceContext) Set_decrement_bl(v IDecrement_blContext) { s._decrement_bl = v }

func (s *SentenceContext) Set_while_bl(v IWhile_blContext) { s._while_bl = v }

func (s *SentenceContext) Set_for_bl(v IFor_blContext) { s._for_bl = v }

func (s *SentenceContext) Set_guard_bl(v IGuard_blContext) { s._guard_bl = v }

func (s *SentenceContext) Set_break_bl(v IBreak_blContext) { s._break_bl = v }

func (s *SentenceContext) Set_return_bl(v IReturn_blContext) { s._return_bl = v }

func (s *SentenceContext) Set_continue_bl(v IContinue_blContext) { s._continue_bl = v }

func (s *SentenceContext) Set_vector_bl(v IVector_blContext) { s._vector_bl = v }

func (s *SentenceContext) Set_function_bl(v IFunction_blContext) { s._function_bl = v }

func (s *SentenceContext) Set_call_function(v ICall_functionContext) { s._call_function = v }

func (s *SentenceContext) Set_array_functions(v IArray_functionsContext) { s._array_functions = v }

func (s *SentenceContext) Set_declare_array_bl(v IDeclare_array_blContext) { s._declare_array_bl = v }

func (s *SentenceContext) Set_switch_bl(v ISwitch_blContext) { s._switch_bl = v }

func (s *SentenceContext) GetInstr() abstract.Instruction { return s.instr }

func (s *SentenceContext) SetInstr(v abstract.Instruction) { s.instr = v }

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

func (s *SentenceContext) Declare_let() IDeclare_letContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclare_letContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclare_letContext)
}

func (s *SentenceContext) Assign_bl() IAssign_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_blContext)
}

func (s *SentenceContext) Assign_vector() IAssign_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_vectorContext)
}

func (s *SentenceContext) Print_bl() IPrint_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrint_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrint_blContext)
}

func (s *SentenceContext) If_bl() IIf_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_blContext)
}

func (s *SentenceContext) Increment_bl() IIncrement_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncrement_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncrement_blContext)
}

func (s *SentenceContext) Decrement_bl() IDecrement_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecrement_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecrement_blContext)
}

func (s *SentenceContext) While_bl() IWhile_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_blContext)
}

func (s *SentenceContext) For_bl() IFor_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_blContext)
}

func (s *SentenceContext) Guard_bl() IGuard_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuard_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuard_blContext)
}

func (s *SentenceContext) Break_bl() IBreak_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_blContext)
}

func (s *SentenceContext) Return_bl() IReturn_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_blContext)
}

func (s *SentenceContext) Continue_bl() IContinue_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_blContext)
}

func (s *SentenceContext) Vector_bl() IVector_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_blContext)
}

func (s *SentenceContext) Function_bl() IFunction_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_blContext)
}

func (s *SentenceContext) Call_function() ICall_functionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_functionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_functionContext)
}

func (s *SentenceContext) Array_functions() IArray_functionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_functionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_functionsContext)
}

func (s *SentenceContext) Declare_array_bl() IDeclare_array_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclare_array_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclare_array_blContext)
}

func (s *SentenceContext) Switch_bl() ISwitch_blContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_blContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_blContext)
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
	p.EnterRule(localctx, 4, SwiftgrammParserRULE_sentence)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)

			var _x = p.Declare_var()

			localctx.(*SentenceContext)._declare_var = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_declare_var().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)

			var _x = p.Declare_let()

			localctx.(*SentenceContext)._declare_let = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_declare_let().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)

			var _x = p.Assign_bl()

			localctx.(*SentenceContext)._assign_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_assign_bl().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)

			var _x = p.Assign_vector()

			localctx.(*SentenceContext)._assign_vector = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_assign_vector().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)

			var _x = p.Print_bl()

			localctx.(*SentenceContext)._print_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_print_bl().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)

			var _x = p.If_bl()

			localctx.(*SentenceContext)._if_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_if_bl().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(103)

			var _x = p.Increment_bl()

			localctx.(*SentenceContext)._increment_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_increment_bl().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(106)

			var _x = p.Decrement_bl()

			localctx.(*SentenceContext)._decrement_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_decrement_bl().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(109)

			var _x = p.While_bl()

			localctx.(*SentenceContext)._while_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_while_bl().GetInstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(112)

			var _x = p.For_bl()

			localctx.(*SentenceContext)._for_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_for_bl().GetInstr()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(115)

			var _x = p.Guard_bl()

			localctx.(*SentenceContext)._guard_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_guard_bl().GetInstr()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(118)

			var _x = p.Break_bl()

			localctx.(*SentenceContext)._break_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_break_bl().GetInstr()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(121)

			var _x = p.Return_bl()

			localctx.(*SentenceContext)._return_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_return_bl().GetInstr()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(124)

			var _x = p.Continue_bl()

			localctx.(*SentenceContext)._continue_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_continue_bl().GetInstr()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(127)

			var _x = p.Vector_bl()

			localctx.(*SentenceContext)._vector_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_vector_bl().GetInstr()

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(130)

			var _x = p.Function_bl()

			localctx.(*SentenceContext)._function_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_function_bl().GetInstr()

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(133)

			var _x = p.Call_function()

			localctx.(*SentenceContext)._call_function = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_call_function().GetInstr()

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(136)

			var _x = p.Array_functions()

			localctx.(*SentenceContext)._array_functions = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_array_functions().GetInstr()

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(139)

			var _x = p.Declare_array_bl()

			localctx.(*SentenceContext)._declare_array_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_declare_array_bl().GetInstr()

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(142)

			var _x = p.Switch_bl()

			localctx.(*SentenceContext)._switch_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_switch_bl().GetInstr()

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

// ISwitch_blContext is an interface to support dynamic dispatch.
type ISwitch_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SWITCH returns the _SWITCH token.
	Get_SWITCH() antlr.Token

	// Set_SWITCH sets the _SWITCH token.
	Set_SWITCH(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_cases returns the _cases rule contexts.
	Get_cases() ICasesContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_cases sets the _cases rule contexts.
	Set_cases(ICasesContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expression() IExpressionContext
	OPEN_kEY() antlr.TerminalNode
	Cases() ICasesContext
	CLOSE_kEY() antlr.TerminalNode

	// IsSwitch_blContext differentiates from other interfaces.
	IsSwitch_blContext()
}

type Switch_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_SWITCH     antlr.Token
	_expression IExpressionContext
	_cases      ICasesContext
}

func NewEmptySwitch_blContext() *Switch_blContext {
	var p = new(Switch_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_switch_bl
	return p
}

func InitEmptySwitch_blContext(p *Switch_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_switch_bl
}

func (*Switch_blContext) IsSwitch_blContext() {}

func NewSwitch_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_blContext {
	var p = new(Switch_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_switch_bl

	return p
}

func (s *Switch_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_blContext) Get_SWITCH() antlr.Token { return s._SWITCH }

func (s *Switch_blContext) Set_SWITCH(v antlr.Token) { s._SWITCH = v }

func (s *Switch_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *Switch_blContext) Get_cases() ICasesContext { return s._cases }

func (s *Switch_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Switch_blContext) Set_cases(v ICasesContext) { s._cases = v }

func (s *Switch_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Switch_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Switch_blContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSWITCH, 0)
}

func (s *Switch_blContext) Expression() IExpressionContext {
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

func (s *Switch_blContext) OPEN_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, 0)
}

func (s *Switch_blContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *Switch_blContext) CLOSE_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, 0)
}

func (s *Switch_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterSwitch_bl(s)
	}
}

func (s *Switch_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitSwitch_bl(s)
	}
}

func (s *Switch_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitSwitch_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Switch_bl() (localctx ISwitch_blContext) {
	localctx = NewSwitch_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftgrammParserRULE_switch_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)

		var _m = p.Match(SwiftgrammParserSWITCH)

		localctx.(*Switch_blContext)._SWITCH = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)

		var _x = p.expression(0)

		localctx.(*Switch_blContext)._expression = _x
	}
	{
		p.SetState(149)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)

		var _x = p.Cases()

		localctx.(*Switch_blContext)._cases = _x
	}
	{
		p.SetState(151)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*Switch_blContext).instr = instructions.NewSwitch(localctx.(*Switch_blContext).Get_expression().GetP(), localctx.(*Switch_blContext).Get_cases().GetP(), (func() int {
		if localctx.(*Switch_blContext).Get_SWITCH() == nil {
			return 0
		} else {
			return localctx.(*Switch_blContext).Get_SWITCH().GetLine()
		}
	}()), (func() int {
		if localctx.(*Switch_blContext).Get_SWITCH() == nil {
			return 0
		} else {
			return localctx.(*Switch_blContext).Get_SWITCH().GetColumn()
		}
	}()))

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

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Get_DEFAULT returns the _DEFAULT token.
	Get_DEFAULT() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// Set_DEFAULT sets the _DEFAULT token.
	Set_DEFAULT(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_cases returns the _cases rule contexts.
	Get_cases() ICasesContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_cases sets the _cases rule contexts.
	Set_cases(ICasesContext)

	// GetP returns the p attribute.
	GetP() []interface{}

	// SetP sets the p attribute.
	SetP([]interface{})

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	Block() IBlockContext
	Cases() ICasesContext
	DEFAULT() antlr.TerminalNode

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           []interface{}
	_CASE       antlr.Token
	_expression IExpressionContext
	_block      IBlockContext
	_cases      ICasesContext
	_DEFAULT    antlr.Token
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_cases
	return p
}

func InitEmptyCasesContext(p *CasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_cases
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) Get_CASE() antlr.Token { return s._CASE }

func (s *CasesContext) Get_DEFAULT() antlr.Token { return s._DEFAULT }

func (s *CasesContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *CasesContext) Set_DEFAULT(v antlr.Token) { s._DEFAULT = v }

func (s *CasesContext) Get_expression() IExpressionContext { return s._expression }

func (s *CasesContext) Get_block() IBlockContext { return s._block }

func (s *CasesContext) Get_cases() ICasesContext { return s._cases }

func (s *CasesContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *CasesContext) Set_block(v IBlockContext) { s._block = v }

func (s *CasesContext) Set_cases(v ICasesContext) { s._cases = v }

func (s *CasesContext) GetP() []interface{} { return s.p }

func (s *CasesContext) SetP(v []interface{}) { s.p = v }

func (s *CasesContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCASE, 0)
}

func (s *CasesContext) Expression() IExpressionContext {
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

func (s *CasesContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *CasesContext) Block() IBlockContext {
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

func (s *CasesContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *CasesContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserDEFAULT, 0)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterCases(s)
	}
}

func (s *CasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitCases(s)
	}
}

func (s *CasesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitCases(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Cases() (localctx ICasesContext) {
	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftgrammParserRULE_cases)
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)

			var _m = p.Match(SwiftgrammParserCASE)

			localctx.(*CasesContext)._CASE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)

			var _x = p.expression(0)

			localctx.(*CasesContext)._expression = _x
		}
		{
			p.SetState(156)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)

			var _x = p.Block()

			localctx.(*CasesContext)._block = _x
		}
		{
			p.SetState(158)

			var _x = p.Cases()

			localctx.(*CasesContext)._cases = _x
		}

		value := instructions.NewCaseSwitch(localctx.(*CasesContext).Get_expression().GetP(), localctx.(*CasesContext).Get_block().GetBlk(), (func() int {
			if localctx.(*CasesContext).Get_CASE() == nil {
				return 0
			} else {
				return localctx.(*CasesContext).Get_CASE().GetLine()
			}
		}()), (func() int {
			if localctx.(*CasesContext).Get_CASE() == nil {
				return 0
			} else {
				return localctx.(*CasesContext).Get_CASE().GetColumn()
			}
		}()))
		localctx.(*CasesContext).p = append([]interface{}{value}, localctx.(*CasesContext).Get_cases().GetP()...)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)

			var _m = p.Match(SwiftgrammParserCASE)

			localctx.(*CasesContext)._CASE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)

			var _x = p.expression(0)

			localctx.(*CasesContext)._expression = _x
		}
		{
			p.SetState(163)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)

			var _x = p.Block()

			localctx.(*CasesContext)._block = _x
		}

		value := instructions.NewCaseSwitch(localctx.(*CasesContext).Get_expression().GetP(), localctx.(*CasesContext).Get_block().GetBlk(), (func() int {
			if localctx.(*CasesContext).Get_CASE() == nil {
				return 0
			} else {
				return localctx.(*CasesContext).Get_CASE().GetLine()
			}
		}()), (func() int {
			if localctx.(*CasesContext).Get_CASE() == nil {
				return 0
			} else {
				return localctx.(*CasesContext).Get_CASE().GetColumn()
			}
		}()))
		localctx.(*CasesContext).p = []interface{}{value}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)

			var _m = p.Match(SwiftgrammParserDEFAULT)

			localctx.(*CasesContext)._DEFAULT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)

			var _x = p.Block()

			localctx.(*CasesContext)._block = _x
		}

		value := instructions.NewCaseSwitch(nil, localctx.(*CasesContext).Get_block().GetBlk(), (func() int {
			if localctx.(*CasesContext).Get_DEFAULT() == nil {
				return 0
			} else {
				return localctx.(*CasesContext).Get_DEFAULT().GetLine()
			}
		}()), (func() int {
			if localctx.(*CasesContext).Get_DEFAULT() == nil {
				return 0
			} else {
				return localctx.(*CasesContext).Get_DEFAULT().GetColumn()
			}
		}()))
		localctx.(*CasesContext).p = []interface{}{value}

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

// IIncrement_blContext is an interface to support dynamic dispatch.
type IIncrement_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_INCREMENT returns the _INCREMENT token.
	Get_INCREMENT() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_INCREMENT sets the _INCREMENT token.
	Set_INCREMENT(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	INCREMENT() antlr.TerminalNode
	Expression() IExpressionContext

	// IsIncrement_blContext differentiates from other interfaces.
	IsIncrement_blContext()
}

type Increment_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_ID         antlr.Token
	_INCREMENT  antlr.Token
	_expression IExpressionContext
}

func NewEmptyIncrement_blContext() *Increment_blContext {
	var p = new(Increment_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_increment_bl
	return p
}

func InitEmptyIncrement_blContext(p *Increment_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_increment_bl
}

func (*Increment_blContext) IsIncrement_blContext() {}

func NewIncrement_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Increment_blContext {
	var p = new(Increment_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_increment_bl

	return p
}

func (s *Increment_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Increment_blContext) Get_ID() antlr.Token { return s._ID }

func (s *Increment_blContext) Get_INCREMENT() antlr.Token { return s._INCREMENT }

func (s *Increment_blContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Increment_blContext) Set_INCREMENT(v antlr.Token) { s._INCREMENT = v }

func (s *Increment_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *Increment_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Increment_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Increment_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Increment_blContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Increment_blContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserINCREMENT, 0)
}

func (s *Increment_blContext) Expression() IExpressionContext {
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

func (s *Increment_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Increment_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Increment_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterIncrement_bl(s)
	}
}

func (s *Increment_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitIncrement_bl(s)
	}
}

func (s *Increment_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitIncrement_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Increment_bl() (localctx IIncrement_blContext) {
	localctx = NewIncrement_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftgrammParserRULE_increment_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Increment_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)

		var _m = p.Match(SwiftgrammParserINCREMENT)

		localctx.(*Increment_blContext)._INCREMENT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)

		var _x = p.expression(0)

		localctx.(*Increment_blContext)._expression = _x
	}

	localctx.(*Increment_blContext).instr = instructions.NewIncreDecrem((func() string {
		if localctx.(*Increment_blContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Increment_blContext).Get_ID().GetText()
		}
	}()), (func() string {
		if localctx.(*Increment_blContext).Get_INCREMENT() == nil {
			return ""
		} else {
			return localctx.(*Increment_blContext).Get_INCREMENT().GetText()
		}
	}()), localctx.(*Increment_blContext).Get_expression().GetP(), (func() int {
		if localctx.(*Increment_blContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Increment_blContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*Increment_blContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Increment_blContext).Get_ID().GetColumn()
		}
	}()))

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

// IDecrement_blContext is an interface to support dynamic dispatch.
type IDecrement_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_DECREMENT returns the _DECREMENT token.
	Get_DECREMENT() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_DECREMENT sets the _DECREMENT token.
	Set_DECREMENT(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode
	Expression() IExpressionContext

	// IsDecrement_blContext differentiates from other interfaces.
	IsDecrement_blContext()
}

type Decrement_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_ID         antlr.Token
	_DECREMENT  antlr.Token
	_expression IExpressionContext
}

func NewEmptyDecrement_blContext() *Decrement_blContext {
	var p = new(Decrement_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_decrement_bl
	return p
}

func InitEmptyDecrement_blContext(p *Decrement_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_decrement_bl
}

func (*Decrement_blContext) IsDecrement_blContext() {}

func NewDecrement_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decrement_blContext {
	var p = new(Decrement_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_decrement_bl

	return p
}

func (s *Decrement_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Decrement_blContext) Get_ID() antlr.Token { return s._ID }

func (s *Decrement_blContext) Get_DECREMENT() antlr.Token { return s._DECREMENT }

func (s *Decrement_blContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Decrement_blContext) Set_DECREMENT(v antlr.Token) { s._DECREMENT = v }

func (s *Decrement_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *Decrement_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Decrement_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Decrement_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Decrement_blContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Decrement_blContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserDECREMENT, 0)
}

func (s *Decrement_blContext) Expression() IExpressionContext {
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

func (s *Decrement_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decrement_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decrement_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterDecrement_bl(s)
	}
}

func (s *Decrement_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitDecrement_bl(s)
	}
}

func (s *Decrement_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitDecrement_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Decrement_bl() (localctx IDecrement_blContext) {
	localctx = NewDecrement_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftgrammParserRULE_decrement_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Decrement_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)

		var _m = p.Match(SwiftgrammParserDECREMENT)

		localctx.(*Decrement_blContext)._DECREMENT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)

		var _x = p.expression(0)

		localctx.(*Decrement_blContext)._expression = _x
	}

	localctx.(*Decrement_blContext).instr = instructions.NewIncreDecrem((func() string {
		if localctx.(*Decrement_blContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Decrement_blContext).Get_ID().GetText()
		}
	}()), (func() string {
		if localctx.(*Decrement_blContext).Get_DECREMENT() == nil {
			return ""
		} else {
			return localctx.(*Decrement_blContext).Get_DECREMENT().GetText()
		}
	}()), localctx.(*Decrement_blContext).Get_expression().GetP(), (func() int {
		if localctx.(*Decrement_blContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Decrement_blContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*Decrement_blContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Decrement_blContext).Get_ID().GetColumn()
		}
	}()))

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

// IBreak_blContext is an interface to support dynamic dispatch.
type IBreak_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreak_blContext differentiates from other interfaces.
	IsBreak_blContext()
}

type Break_blContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	instr  abstract.Instruction
}

func NewEmptyBreak_blContext() *Break_blContext {
	var p = new(Break_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_break_bl
	return p
}

func InitEmptyBreak_blContext(p *Break_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_break_bl
}

func (*Break_blContext) IsBreak_blContext() {}

func NewBreak_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_blContext {
	var p = new(Break_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_break_bl

	return p
}

func (s *Break_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Break_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Break_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Break_blContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserBREAK, 0)
}

func (s *Break_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterBreak_bl(s)
	}
}

func (s *Break_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitBreak_bl(s)
	}
}

func (s *Break_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitBreak_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Break_bl() (localctx IBreak_blContext) {
	localctx = NewBreak_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftgrammParserRULE_break_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(SwiftgrammParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*Break_blContext).instr = instructions.NewBreak("break")

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

// IReturn_blContext is an interface to support dynamic dispatch.
type IReturn_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturn_blContext differentiates from other interfaces.
	IsReturn_blContext()
}

type Return_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_expression IExpressionContext
	_RETURN     antlr.Token
}

func NewEmptyReturn_blContext() *Return_blContext {
	var p = new(Return_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_return_bl
	return p
}

func InitEmptyReturn_blContext(p *Return_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_return_bl
}

func (*Return_blContext) IsReturn_blContext() {}

func NewReturn_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_blContext {
	var p = new(Return_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_return_bl

	return p
}

func (s *Return_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_blContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *Return_blContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *Return_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *Return_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Return_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Return_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Return_blContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserRETURN, 0)
}

func (s *Return_blContext) Expression() IExpressionContext {
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

func (s *Return_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterReturn_bl(s)
	}
}

func (s *Return_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitReturn_bl(s)
	}
}

func (s *Return_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitReturn_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Return_bl() (localctx IReturn_blContext) {
	localctx = NewReturn_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftgrammParserRULE_return_bl)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(SwiftgrammParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _x = p.expression(0)

			localctx.(*Return_blContext)._expression = _x
		}

		localctx.(*Return_blContext).instr = instructions.NewReturn(localctx.(*Return_blContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)

			var _m = p.Match(SwiftgrammParserRETURN)

			localctx.(*Return_blContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Return_blContext).instr = instructions.NewReturn(expressions.NewNative(nil, symbol.NIL, (func() int {
			if localctx.(*Return_blContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*Return_blContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*Return_blContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*Return_blContext).Get_RETURN().GetColumn()
			}
		}())))

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

// IContinue_blContext is an interface to support dynamic dispatch.
type IContinue_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinue_blContext differentiates from other interfaces.
	IsContinue_blContext()
}

type Continue_blContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	instr  abstract.Instruction
}

func NewEmptyContinue_blContext() *Continue_blContext {
	var p = new(Continue_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_continue_bl
	return p
}

func InitEmptyContinue_blContext(p *Continue_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_continue_bl
}

func (*Continue_blContext) IsContinue_blContext() {}

func NewContinue_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_blContext {
	var p = new(Continue_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_continue_bl

	return p
}

func (s *Continue_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Continue_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Continue_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Continue_blContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCONTINUE, 0)
}

func (s *Continue_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterContinue_bl(s)
	}
}

func (s *Continue_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitContinue_bl(s)
	}
}

func (s *Continue_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitContinue_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Continue_bl() (localctx IContinue_blContext) {
	localctx = NewContinue_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftgrammParserRULE_continue_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(SwiftgrammParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*Continue_blContext).instr = instructions.NewContinue("continue")

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

// IDeclare_letContext is an interface to support dynamic dispatch.
type IDeclare_letContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_datatype returns the _datatype rule contexts.
	Get_datatype() IDatatypeContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_datatype sets the _datatype rule contexts.
	Set_datatype(IDatatypeContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	LET() antlr.TerminalNode
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Datatype() IDatatypeContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsDeclare_letContext differentiates from other interfaces.
	IsDeclare_letContext()
}

type Declare_letContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_ID         antlr.Token
	_datatype   IDatatypeContext
	_expression IExpressionContext
}

func NewEmptyDeclare_letContext() *Declare_letContext {
	var p = new(Declare_letContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_declare_let
	return p
}

func InitEmptyDeclare_letContext(p *Declare_letContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_declare_let
}

func (*Declare_letContext) IsDeclare_letContext() {}

func NewDeclare_letContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declare_letContext {
	var p = new(Declare_letContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_declare_let

	return p
}

func (s *Declare_letContext) GetParser() antlr.Parser { return s.parser }

func (s *Declare_letContext) Get_ID() antlr.Token { return s._ID }

func (s *Declare_letContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Declare_letContext) Get_datatype() IDatatypeContext { return s._datatype }

func (s *Declare_letContext) Get_expression() IExpressionContext { return s._expression }

func (s *Declare_letContext) Set_datatype(v IDatatypeContext) { s._datatype = v }

func (s *Declare_letContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Declare_letContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Declare_letContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Declare_letContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserLET, 0)
}

func (s *Declare_letContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Declare_letContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *Declare_letContext) Datatype() IDatatypeContext {
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

func (s *Declare_letContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserASSIGN, 0)
}

func (s *Declare_letContext) Expression() IExpressionContext {
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

func (s *Declare_letContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declare_letContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declare_letContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterDeclare_let(s)
	}
}

func (s *Declare_letContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitDeclare_let(s)
	}
}

func (s *Declare_letContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitDeclare_let(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Declare_let() (localctx IDeclare_letContext) {
	localctx = NewDeclare_letContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftgrammParserRULE_declare_let)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_letContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)

			var _x = p.Datatype()

			localctx.(*Declare_letContext)._datatype = _x
		}
		{
			p.SetState(202)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)

			var _x = p.expression(0)

			localctx.(*Declare_letContext)._expression = _x
		}

		localctx.(*Declare_letContext).instr = instructions.NewLet((func() string {
			if localctx.(*Declare_letContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_letContext).Get_ID().GetText()
			}
		}()), localctx.(*Declare_letContext).Get_datatype().GetTd(), localctx.(*Declare_letContext).Get_expression().GetP(), (func() int {
			if localctx.(*Declare_letContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_letContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Declare_letContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_letContext).Get_ID().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_letContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)

			var _x = p.expression(0)

			localctx.(*Declare_letContext)._expression = _x
		}

		localctx.(*Declare_letContext).instr = instructions.NewLet((func() string {
			if localctx.(*Declare_letContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_letContext).Get_ID().GetText()
			}
		}()), symbol.UNDEFINED, localctx.(*Declare_letContext).Get_expression().GetP(), (func() int {
			if localctx.(*Declare_letContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_letContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Declare_letContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_letContext).Get_ID().GetColumn()
			}
		}()))

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

// INative_functionContext is an interface to support dynamic dispatch.
type INative_functionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_INT sets the _INT token.
	Set_INT(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetP returns the p attribute.
	GetP() abstract.Expression

	// SetP sets the p attribute.
	SetP(abstract.Expression)

	// Getter signatures
	STRING() antlr.TerminalNode
	OPEN_PARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	CLOSE_PARENTHESIS() antlr.TerminalNode
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsNative_functionContext differentiates from other interfaces.
	IsNative_functionContext()
}

type Native_functionContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           abstract.Expression
	_STRING     antlr.Token
	_expression IExpressionContext
	_INT        antlr.Token
	_FLOAT      antlr.Token
}

func NewEmptyNative_functionContext() *Native_functionContext {
	var p = new(Native_functionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_native_function
	return p
}

func InitEmptyNative_functionContext(p *Native_functionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_native_function
}

func (*Native_functionContext) IsNative_functionContext() {}

func NewNative_functionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Native_functionContext {
	var p = new(Native_functionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_native_function

	return p
}

func (s *Native_functionContext) GetParser() antlr.Parser { return s.parser }

func (s *Native_functionContext) Get_STRING() antlr.Token { return s._STRING }

func (s *Native_functionContext) Get_INT() antlr.Token { return s._INT }

func (s *Native_functionContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *Native_functionContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *Native_functionContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *Native_functionContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *Native_functionContext) Get_expression() IExpressionContext { return s._expression }

func (s *Native_functionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Native_functionContext) GetP() abstract.Expression { return s.p }

func (s *Native_functionContext) SetP(v abstract.Expression) { s.p = v }

func (s *Native_functionContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSTRING, 0)
}

func (s *Native_functionContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *Native_functionContext) Expression() IExpressionContext {
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

func (s *Native_functionContext) CLOSE_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_PARENTHESIS, 0)
}

func (s *Native_functionContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserINT, 0)
}

func (s *Native_functionContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserFLOAT, 0)
}

func (s *Native_functionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Native_functionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Native_functionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterNative_function(s)
	}
}

func (s *Native_functionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitNative_function(s)
	}
}

func (s *Native_functionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitNative_function(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Native_function() (localctx INative_functionContext) {
	localctx = NewNative_functionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftgrammParserRULE_native_function)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)

			var _m = p.Match(SwiftgrammParserSTRING)

			localctx.(*Native_functionContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)

			var _x = p.expression(0)

			localctx.(*Native_functionContext)._expression = _x
		}
		{
			p.SetState(217)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Native_functionContext).p = expressions.NewNativeFunction((func() string {
			if localctx.(*Native_functionContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Native_functionContext).Get_STRING().GetText()
			}
		}()), localctx.(*Native_functionContext).Get_expression().GetP())

	case SwiftgrammParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)

			var _m = p.Match(SwiftgrammParserINT)

			localctx.(*Native_functionContext)._INT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)

			var _x = p.expression(0)

			localctx.(*Native_functionContext)._expression = _x
		}
		{
			p.SetState(223)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Native_functionContext).p = expressions.NewNativeFunction((func() string {
			if localctx.(*Native_functionContext).Get_INT() == nil {
				return ""
			} else {
				return localctx.(*Native_functionContext).Get_INT().GetText()
			}
		}()), localctx.(*Native_functionContext).Get_expression().GetP())

	case SwiftgrammParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(226)

			var _m = p.Match(SwiftgrammParserFLOAT)

			localctx.(*Native_functionContext)._FLOAT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)

			var _x = p.expression(0)

			localctx.(*Native_functionContext)._expression = _x
		}
		{
			p.SetState(229)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Native_functionContext).p = expressions.NewNativeFunction((func() string {
			if localctx.(*Native_functionContext).Get_FLOAT() == nil {
				return ""
			} else {
				return localctx.(*Native_functionContext).Get_FLOAT().GetText()
			}
		}()), localctx.(*Native_functionContext).Get_expression().GetP())

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

// IDeclare_varContext is an interface to support dynamic dispatch.
type IDeclare_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_QUESTION_MARK returns the _QUESTION_MARK token.
	Get_QUESTION_MARK() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_QUESTION_MARK sets the _QUESTION_MARK token.
	Set_QUESTION_MARK(antlr.Token)

	// Get_datatype returns the _datatype rule contexts.
	Get_datatype() IDatatypeContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_datatype sets the _datatype rule contexts.
	Set_datatype(IDatatypeContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

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
	parser         antlr.Parser
	instr          abstract.Instruction
	_ID            antlr.Token
	_datatype      IDatatypeContext
	_expression    IExpressionContext
	_QUESTION_MARK antlr.Token
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

func (s *Declare_varContext) Get_ID() antlr.Token { return s._ID }

func (s *Declare_varContext) Get_QUESTION_MARK() antlr.Token { return s._QUESTION_MARK }

func (s *Declare_varContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Declare_varContext) Set_QUESTION_MARK(v antlr.Token) { s._QUESTION_MARK = v }

func (s *Declare_varContext) Get_datatype() IDatatypeContext { return s._datatype }

func (s *Declare_varContext) Get_expression() IExpressionContext { return s._expression }

func (s *Declare_varContext) Set_datatype(v IDatatypeContext) { s._datatype = v }

func (s *Declare_varContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Declare_varContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Declare_varContext) SetInstr(v abstract.Instruction) { s.instr = v }

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
	p.EnterRule(localctx, 24, SwiftgrammParserRULE_declare_var)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)

			var _x = p.Datatype()

			localctx.(*Declare_varContext)._datatype = _x
		}
		{
			p.SetState(238)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)

			var _x = p.expression(0)

			localctx.(*Declare_varContext)._expression = _x
		}

		localctx.(*Declare_varContext).instr = instructions.NewDeclareWithValue((func() string {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetText()
			}
		}()), localctx.(*Declare_varContext).Get_datatype().GetTd(), localctx.(*Declare_varContext).Get_expression().GetP(), (func() int {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)

			var _x = p.expression(0)

			localctx.(*Declare_varContext)._expression = _x
		}

		localctx.(*Declare_varContext).instr = instructions.NewDeclareWithValue((func() string {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetText()
			}
		}()), symbol.UNDEFINED, localctx.(*Declare_varContext).Get_expression().GetP(), (func() int {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetColumn()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(248)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)

			var _x = p.Datatype()

			localctx.(*Declare_varContext)._datatype = _x
		}
		{
			p.SetState(252)

			var _m = p.Match(SwiftgrammParserQUESTION_MARK)

			localctx.(*Declare_varContext)._QUESTION_MARK = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Declare_varContext).instr = instructions.NewDeclareWithoutValue((func() string {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetText()
			}
		}()), localctx.(*Declare_varContext).Get_datatype().GetTd(), expressions.NewNative(nil, symbol.NIL, (func() int {
			if localctx.(*Declare_varContext).Get_QUESTION_MARK() == nil {
				return 0
			} else {
				return localctx.(*Declare_varContext).Get_QUESTION_MARK().GetLine()
			}
		}()), (func() int {
			if localctx.(*Declare_varContext).Get_QUESTION_MARK() == nil {
				return 0
			} else {
				return localctx.(*Declare_varContext).Get_QUESTION_MARK().GetColumn()
			}
		}())), (func() int {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetColumn()
			}
		}()))

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

// IAssign_blContext is an interface to support dynamic dispatch.
type IAssign_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssign_blContext differentiates from other interfaces.
	IsAssign_blContext()
}

type Assign_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_ID         antlr.Token
	_expression IExpressionContext
}

func NewEmptyAssign_blContext() *Assign_blContext {
	var p = new(Assign_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_assign_bl
	return p
}

func InitEmptyAssign_blContext(p *Assign_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_assign_bl
}

func (*Assign_blContext) IsAssign_blContext() {}

func NewAssign_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_blContext {
	var p = new(Assign_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_assign_bl

	return p
}

func (s *Assign_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_blContext) Get_ID() antlr.Token { return s._ID }

func (s *Assign_blContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Assign_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *Assign_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Assign_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Assign_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Assign_blContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Assign_blContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserASSIGN, 0)
}

func (s *Assign_blContext) Expression() IExpressionContext {
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

func (s *Assign_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterAssign_bl(s)
	}
}

func (s *Assign_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitAssign_bl(s)
	}
}

func (s *Assign_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitAssign_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Assign_bl() (localctx IAssign_blContext) {
	localctx = NewAssign_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftgrammParserRULE_assign_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Assign_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(SwiftgrammParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)

		var _x = p.expression(0)

		localctx.(*Assign_blContext)._expression = _x
	}

	localctx.(*Assign_blContext).instr = instructions.NewAsignVariable((func() string {
		if localctx.(*Assign_blContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Assign_blContext).Get_ID().GetText()
		}
	}()), localctx.(*Assign_blContext).Get_expression().GetP(), (func() int {
		if localctx.(*Assign_blContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Assign_blContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*Assign_blContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Assign_blContext).Get_ID().GetColumn()
		}
	}()))

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

// IPrint_blContext is an interface to support dynamic dispatch.
type IPrint_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_list_print returns the _list_print rule contexts.
	Get_list_print() IList_printContext

	// Set_list_print sets the _list_print rule contexts.
	Set_list_print(IList_printContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	OPEN_PARENTHESIS() antlr.TerminalNode
	List_print() IList_printContext
	CLOSE_PARENTHESIS() antlr.TerminalNode

	// IsPrint_blContext differentiates from other interfaces.
	IsPrint_blContext()
}

type Print_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_list_print IList_printContext
}

func NewEmptyPrint_blContext() *Print_blContext {
	var p = new(Print_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_print_bl
	return p
}

func InitEmptyPrint_blContext(p *Print_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_print_bl
}

func (*Print_blContext) IsPrint_blContext() {}

func NewPrint_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_blContext {
	var p = new(Print_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_print_bl

	return p
}

func (s *Print_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_blContext) Get_list_print() IList_printContext { return s._list_print }

func (s *Print_blContext) Set_list_print(v IList_printContext) { s._list_print = v }

func (s *Print_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Print_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Print_blContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserPRINT, 0)
}

func (s *Print_blContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *Print_blContext) List_print() IList_printContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_printContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_printContext)
}

func (s *Print_blContext) CLOSE_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_PARENTHESIS, 0)
}

func (s *Print_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterPrint_bl(s)
	}
}

func (s *Print_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitPrint_bl(s)
	}
}

func (s *Print_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitPrint_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Print_bl() (localctx IPrint_blContext) {
	localctx = NewPrint_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftgrammParserRULE_print_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(SwiftgrammParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(SwiftgrammParserOPEN_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)

		var _x = p.List_print()

		localctx.(*Print_blContext)._list_print = _x
	}
	{
		p.SetState(265)
		p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*Print_blContext).instr = instructions.NewPrint(localctx.(*Print_blContext).Get_list_print().GetP())

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

// IList_printContext is an interface to support dynamic dispatch.
type IList_printContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_list_print returns the _list_print rule contexts.
	Get_list_print() IList_printContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_list_print sets the _list_print rule contexts.
	Set_list_print(IList_printContext)

	// GetP returns the p attribute.
	GetP() []interface{}

	// SetP sets the p attribute.
	SetP([]interface{})

	// Getter signatures
	Expression() IExpressionContext
	COMMA() antlr.TerminalNode
	List_print() IList_printContext

	// IsList_printContext differentiates from other interfaces.
	IsList_printContext()
}

type List_printContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           []interface{}
	_expression IExpressionContext
	_list_print IList_printContext
}

func NewEmptyList_printContext() *List_printContext {
	var p = new(List_printContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_list_print
	return p
}

func InitEmptyList_printContext(p *List_printContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_list_print
}

func (*List_printContext) IsList_printContext() {}

func NewList_printContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_printContext {
	var p = new(List_printContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_list_print

	return p
}

func (s *List_printContext) GetParser() antlr.Parser { return s.parser }

func (s *List_printContext) Get_expression() IExpressionContext { return s._expression }

func (s *List_printContext) Get_list_print() IList_printContext { return s._list_print }

func (s *List_printContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *List_printContext) Set_list_print(v IList_printContext) { s._list_print = v }

func (s *List_printContext) GetP() []interface{} { return s.p }

func (s *List_printContext) SetP(v []interface{}) { s.p = v }

func (s *List_printContext) Expression() IExpressionContext {
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

func (s *List_printContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOMMA, 0)
}

func (s *List_printContext) List_print() IList_printContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_printContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_printContext)
}

func (s *List_printContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_printContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_printContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterList_print(s)
	}
}

func (s *List_printContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitList_print(s)
	}
}

func (s *List_printContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitList_print(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) List_print() (localctx IList_printContext) {
	localctx = NewList_printContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftgrammParserRULE_list_print)
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)

			var _x = p.expression(0)

			localctx.(*List_printContext)._expression = _x
		}
		{
			p.SetState(269)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)

			var _x = p.List_print()

			localctx.(*List_printContext)._list_print = _x
		}

		localctx.(*List_printContext).p = append([]interface{}{localctx.(*List_printContext).Get_expression().GetP()}, localctx.(*List_printContext).Get_list_print().GetP()...)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)

			var _x = p.expression(0)

			localctx.(*List_printContext)._expression = _x
		}

		localctx.(*List_printContext).p = []interface{}{localctx.(*List_printContext).Get_expression().GetP()}

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

// IIf_blContext is an interface to support dynamic dispatch.
type IIf_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// GetIfblock returns the ifblock rule contexts.
	GetIfblock() IBlockContext

	// GetElseblock returns the elseblock rule contexts.
	GetElseblock() IBlockContext

	// Get_else_if returns the _else_if rule contexts.
	Get_else_if() IElse_ifContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// SetIfblock sets the ifblock rule contexts.
	SetIfblock(IBlockContext)

	// SetElseblock sets the elseblock rule contexts.
	SetElseblock(IBlockContext)

	// Set_else_if sets the _else_if rule contexts.
	Set_else_if(IElse_ifContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllOPEN_kEY() []antlr.TerminalNode
	OPEN_kEY(i int) antlr.TerminalNode
	AllCLOSE_kEY() []antlr.TerminalNode
	CLOSE_kEY(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode
	Else_if() IElse_ifContext

	// IsIf_blContext differentiates from other interfaces.
	IsIf_blContext()
}

type If_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_IF         antlr.Token
	_expression IExpressionContext
	ifblock     IBlockContext
	elseblock   IBlockContext
	_else_if    IElse_ifContext
}

func NewEmptyIf_blContext() *If_blContext {
	var p = new(If_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_if_bl
	return p
}

func InitEmptyIf_blContext(p *If_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_if_bl
}

func (*If_blContext) IsIf_blContext() {}

func NewIf_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_blContext {
	var p = new(If_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_if_bl

	return p
}

func (s *If_blContext) GetParser() antlr.Parser { return s.parser }

func (s *If_blContext) Get_IF() antlr.Token { return s._IF }

func (s *If_blContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *If_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *If_blContext) GetIfblock() IBlockContext { return s.ifblock }

func (s *If_blContext) GetElseblock() IBlockContext { return s.elseblock }

func (s *If_blContext) Get_else_if() IElse_ifContext { return s._else_if }

func (s *If_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *If_blContext) SetIfblock(v IBlockContext) { s.ifblock = v }

func (s *If_blContext) SetElseblock(v IBlockContext) { s.elseblock = v }

func (s *If_blContext) Set_else_if(v IElse_ifContext) { s._else_if = v }

func (s *If_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *If_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *If_blContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserIF, 0)
}

func (s *If_blContext) Expression() IExpressionContext {
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

func (s *If_blContext) AllOPEN_kEY() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammParserOPEN_kEY)
}

func (s *If_blContext) OPEN_kEY(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, i)
}

func (s *If_blContext) AllCLOSE_kEY() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammParserCLOSE_kEY)
}

func (s *If_blContext) CLOSE_kEY(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, i)
}

func (s *If_blContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *If_blContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *If_blContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserELSE, 0)
}

func (s *If_blContext) Else_if() IElse_ifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_ifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_ifContext)
}

func (s *If_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterIf_bl(s)
	}
}

func (s *If_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitIf_bl(s)
	}
}

func (s *If_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitIf_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) If_bl() (localctx IIf_blContext) {
	localctx = NewIf_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftgrammParserRULE_if_bl)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(278)

			var _m = p.Match(SwiftgrammParserIF)

			localctx.(*If_blContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(280)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(282)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*If_blContext).instr = instructions.NewIf(localctx.(*If_blContext).Get_expression().GetP(), localctx.(*If_blContext).GetIfblock().GetBlk(), nil, (func() int {
			if localctx.(*If_blContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*If_blContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*If_blContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*If_blContext).Get_IF().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)

			var _m = p.Match(SwiftgrammParserIF)

			localctx.(*If_blContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(287)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(289)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)

			var _x = p.Block()

			localctx.(*If_blContext).elseblock = _x
		}
		{
			p.SetState(293)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*If_blContext).instr = instructions.NewIf(localctx.(*If_blContext).Get_expression().GetP(), localctx.(*If_blContext).GetIfblock().GetBlk(), localctx.(*If_blContext).GetElseblock().GetBlk(), (func() int {
			if localctx.(*If_blContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*If_blContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*If_blContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*If_blContext).Get_IF().GetColumn()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)

			var _m = p.Match(SwiftgrammParserIF)

			localctx.(*If_blContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(298)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(300)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)

			var _x = p.Else_if()

			localctx.(*If_blContext)._else_if = _x
		}

		localctx.(*If_blContext).instr = instructions.NewIf(localctx.(*If_blContext).Get_expression().GetP(), localctx.(*If_blContext).GetIfblock().GetBlk(), localctx.(*If_blContext).Get_else_if().GetInstr(), (func() int {
			if localctx.(*If_blContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*If_blContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*If_blContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*If_blContext).Get_IF().GetColumn()
			}
		}()))

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

// IElse_ifContext is an interface to support dynamic dispatch.
type IElse_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ELSE returns the _ELSE token.
	Get_ELSE() antlr.Token

	// Set_ELSE sets the _ELSE token.
	Set_ELSE(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// GetIfblock returns the ifblock rule contexts.
	GetIfblock() IBlockContext

	// GetElseblock returns the elseblock rule contexts.
	GetElseblock() IBlockContext

	// Get_else_if returns the _else_if rule contexts.
	Get_else_if() IElse_ifContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// SetIfblock sets the ifblock rule contexts.
	SetIfblock(IBlockContext)

	// SetElseblock sets the elseblock rule contexts.
	SetElseblock(IBlockContext)

	// Set_else_if sets the _else_if rule contexts.
	Set_else_if(IElse_ifContext)

	// GetInstr returns the instr attribute.
	GetInstr() []interface{}

	// SetInstr sets the instr attribute.
	SetInstr([]interface{})

	// Getter signatures
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllOPEN_kEY() []antlr.TerminalNode
	OPEN_kEY(i int) antlr.TerminalNode
	AllCLOSE_kEY() []antlr.TerminalNode
	CLOSE_kEY(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	Else_if() IElse_ifContext

	// IsElse_ifContext differentiates from other interfaces.
	IsElse_ifContext()
}

type Else_ifContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       []interface{}
	_ELSE       antlr.Token
	_expression IExpressionContext
	ifblock     IBlockContext
	elseblock   IBlockContext
	_else_if    IElse_ifContext
}

func NewEmptyElse_ifContext() *Else_ifContext {
	var p = new(Else_ifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_else_if
	return p
}

func InitEmptyElse_ifContext(p *Else_ifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_else_if
}

func (*Else_ifContext) IsElse_ifContext() {}

func NewElse_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_ifContext {
	var p = new(Else_ifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_else_if

	return p
}

func (s *Else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_ifContext) Get_ELSE() antlr.Token { return s._ELSE }

func (s *Else_ifContext) Set_ELSE(v antlr.Token) { s._ELSE = v }

func (s *Else_ifContext) Get_expression() IExpressionContext { return s._expression }

func (s *Else_ifContext) GetIfblock() IBlockContext { return s.ifblock }

func (s *Else_ifContext) GetElseblock() IBlockContext { return s.elseblock }

func (s *Else_ifContext) Get_else_if() IElse_ifContext { return s._else_if }

func (s *Else_ifContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Else_ifContext) SetIfblock(v IBlockContext) { s.ifblock = v }

func (s *Else_ifContext) SetElseblock(v IBlockContext) { s.elseblock = v }

func (s *Else_ifContext) Set_else_if(v IElse_ifContext) { s._else_if = v }

func (s *Else_ifContext) GetInstr() []interface{} { return s.instr }

func (s *Else_ifContext) SetInstr(v []interface{}) { s.instr = v }

func (s *Else_ifContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammParserELSE)
}

func (s *Else_ifContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserELSE, i)
}

func (s *Else_ifContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserIF, 0)
}

func (s *Else_ifContext) Expression() IExpressionContext {
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

func (s *Else_ifContext) AllOPEN_kEY() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammParserOPEN_kEY)
}

func (s *Else_ifContext) OPEN_kEY(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, i)
}

func (s *Else_ifContext) AllCLOSE_kEY() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammParserCLOSE_kEY)
}

func (s *Else_ifContext) CLOSE_kEY(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, i)
}

func (s *Else_ifContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *Else_ifContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *Else_ifContext) Else_if() IElse_ifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_ifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_ifContext)
}

func (s *Else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterElse_if(s)
	}
}

func (s *Else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitElse_if(s)
	}
}

func (s *Else_ifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitElse_if(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Else_if() (localctx IElse_ifContext) {
	localctx = NewElse_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftgrammParserRULE_else_if)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)

			var _m = p.Match(SwiftgrammParserELSE)

			localctx.(*Else_ifContext)._ELSE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(309)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(311)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Else_ifContext).instr = []interface{}{instructions.NewIf(localctx.(*Else_ifContext).Get_expression().GetP(), localctx.(*Else_ifContext).GetIfblock().GetBlk(), nil, (func() int {
			if localctx.(*Else_ifContext).Get_ELSE() == nil {
				return 0
			} else {
				return localctx.(*Else_ifContext).Get_ELSE().GetLine()
			}
		}()), (func() int {
			if localctx.(*Else_ifContext).Get_ELSE() == nil {
				return 0
			} else {
				return localctx.(*Else_ifContext).Get_ELSE().GetColumn()
			}
		}()))}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(314)

			var _m = p.Match(SwiftgrammParserELSE)

			localctx.(*Else_ifContext)._ELSE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(317)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(319)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)

			var _m = p.Match(SwiftgrammParserELSE)

			localctx.(*Else_ifContext)._ELSE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)

			var _x = p.Block()

			localctx.(*Else_ifContext).elseblock = _x
		}
		{
			p.SetState(323)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Else_ifContext).instr = []interface{}{instructions.NewIf(localctx.(*Else_ifContext).Get_expression().GetP(), localctx.(*Else_ifContext).GetIfblock().GetBlk(), localctx.(*Else_ifContext).GetElseblock().GetBlk(), (func() int {
			if localctx.(*Else_ifContext).Get_ELSE() == nil {
				return 0
			} else {
				return localctx.(*Else_ifContext).Get_ELSE().GetLine()
			}
		}()), (func() int {
			if localctx.(*Else_ifContext).Get_ELSE() == nil {
				return 0
			} else {
				return localctx.(*Else_ifContext).Get_ELSE().GetColumn()
			}
		}()))}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(326)

			var _m = p.Match(SwiftgrammParserELSE)

			localctx.(*Else_ifContext)._ELSE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(329)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(331)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)

			var _x = p.Else_if()

			localctx.(*Else_ifContext)._else_if = _x
		}

		localctx.(*Else_ifContext).instr = []interface{}{instructions.NewIf(localctx.(*Else_ifContext).Get_expression().GetP(), localctx.(*Else_ifContext).GetIfblock().GetBlk(), localctx.(*Else_ifContext).Get_else_if().GetInstr(), (func() int {
			if localctx.(*Else_ifContext).Get_ELSE() == nil {
				return 0
			} else {
				return localctx.(*Else_ifContext).Get_ELSE().GetLine()
			}
		}()), (func() int {
			if localctx.(*Else_ifContext).Get_ELSE() == nil {
				return 0
			} else {
				return localctx.(*Else_ifContext).Get_ELSE().GetColumn()
			}
		}()))}

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

// IWhile_blContext is an interface to support dynamic dispatch.
type IWhile_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	OPEN_kEY() antlr.TerminalNode
	Block() IBlockContext
	CLOSE_kEY() antlr.TerminalNode

	// IsWhile_blContext differentiates from other interfaces.
	IsWhile_blContext()
}

type While_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_WHILE      antlr.Token
	_expression IExpressionContext
	_block      IBlockContext
}

func NewEmptyWhile_blContext() *While_blContext {
	var p = new(While_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_while_bl
	return p
}

func InitEmptyWhile_blContext(p *While_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_while_bl
}

func (*While_blContext) IsWhile_blContext() {}

func NewWhile_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_blContext {
	var p = new(While_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_while_bl

	return p
}

func (s *While_blContext) GetParser() antlr.Parser { return s.parser }

func (s *While_blContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *While_blContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *While_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *While_blContext) Get_block() IBlockContext { return s._block }

func (s *While_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *While_blContext) Set_block(v IBlockContext) { s._block = v }

func (s *While_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *While_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *While_blContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserWHILE, 0)
}

func (s *While_blContext) Expression() IExpressionContext {
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

func (s *While_blContext) OPEN_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, 0)
}

func (s *While_blContext) Block() IBlockContext {
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

func (s *While_blContext) CLOSE_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, 0)
}

func (s *While_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterWhile_bl(s)
	}
}

func (s *While_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitWhile_bl(s)
	}
}

func (s *While_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitWhile_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) While_bl() (localctx IWhile_blContext) {
	localctx = NewWhile_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftgrammParserRULE_while_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)

		var _m = p.Match(SwiftgrammParserWHILE)

		localctx.(*While_blContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)

		var _x = p.expression(0)

		localctx.(*While_blContext)._expression = _x
	}
	{
		p.SetState(339)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)

		var _x = p.Block()

		localctx.(*While_blContext)._block = _x
	}
	{
		p.SetState(341)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*While_blContext).instr = instructions.NewWhile(localctx.(*While_blContext).Get_expression().GetP(), localctx.(*While_blContext).Get_block().GetBlk(), (func() int {
		if localctx.(*While_blContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*While_blContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*While_blContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*While_blContext).Get_WHILE().GetColumn()
		}
	}()))

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

// IFor_blContext is an interface to support dynamic dispatch.
type IFor_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetExpression1 returns the expression1 rule contexts.
	GetExpression1() IExpressionContext

	// GetExpression2 returns the expression2 rule contexts.
	GetExpression2() IExpressionContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// SetExpression1 sets the expression1 rule contexts.
	SetExpression1(IExpressionContext)

	// SetExpression2 sets the expression2 rule contexts.
	SetExpression2(IExpressionContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	RANGE() antlr.TerminalNode
	OPEN_kEY() antlr.TerminalNode
	Block() IBlockContext
	CLOSE_kEY() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsFor_blContext differentiates from other interfaces.
	IsFor_blContext()
}

type For_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_ID         antlr.Token
	expression1 IExpressionContext
	expression2 IExpressionContext
	_block      IBlockContext
}

func NewEmptyFor_blContext() *For_blContext {
	var p = new(For_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_for_bl
	return p
}

func InitEmptyFor_blContext(p *For_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_for_bl
}

func (*For_blContext) IsFor_blContext() {}

func NewFor_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_blContext {
	var p = new(For_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_for_bl

	return p
}

func (s *For_blContext) GetParser() antlr.Parser { return s.parser }

func (s *For_blContext) Get_ID() antlr.Token { return s._ID }

func (s *For_blContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *For_blContext) GetExpression1() IExpressionContext { return s.expression1 }

func (s *For_blContext) GetExpression2() IExpressionContext { return s.expression2 }

func (s *For_blContext) Get_block() IBlockContext { return s._block }

func (s *For_blContext) SetExpression1(v IExpressionContext) { s.expression1 = v }

func (s *For_blContext) SetExpression2(v IExpressionContext) { s.expression2 = v }

func (s *For_blContext) Set_block(v IBlockContext) { s._block = v }

func (s *For_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *For_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *For_blContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserFOR, 0)
}

func (s *For_blContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *For_blContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserIN, 0)
}

func (s *For_blContext) RANGE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserRANGE, 0)
}

func (s *For_blContext) OPEN_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, 0)
}

func (s *For_blContext) Block() IBlockContext {
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

func (s *For_blContext) CLOSE_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, 0)
}

func (s *For_blContext) AllExpression() []IExpressionContext {
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

func (s *For_blContext) Expression(i int) IExpressionContext {
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

func (s *For_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterFor_bl(s)
	}
}

func (s *For_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitFor_bl(s)
	}
}

func (s *For_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitFor_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) For_bl() (localctx IFor_blContext) {
	localctx = NewFor_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftgrammParserRULE_for_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(SwiftgrammParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*For_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.Match(SwiftgrammParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)

		var _x = p.expression(0)

		localctx.(*For_blContext).expression1 = _x
	}
	{
		p.SetState(348)
		p.Match(SwiftgrammParserRANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)

		var _x = p.expression(0)

		localctx.(*For_blContext).expression2 = _x
	}
	{
		p.SetState(350)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)

		var _x = p.Block()

		localctx.(*For_blContext)._block = _x
	}
	{
		p.SetState(352)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*For_blContext).instr = instructions.NewFor((func() string {
		if localctx.(*For_blContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*For_blContext).Get_ID().GetText()
		}
	}()), localctx.(*For_blContext).GetExpression1().GetP(), localctx.(*For_blContext).GetExpression2().GetP(), localctx.(*For_blContext).Get_block().GetBlk(), (func() int {
		if localctx.(*For_blContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*For_blContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*For_blContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*For_blContext).Get_ID().GetColumn()
		}
	}()))

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

// IGuard_blContext is an interface to support dynamic dispatch.
type IGuard_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_GUARD returns the _GUARD token.
	Get_GUARD() antlr.Token

	// Set_GUARD sets the _GUARD token.
	Set_GUARD(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expression() IExpressionContext
	ELSE() antlr.TerminalNode
	OPEN_kEY() antlr.TerminalNode
	Block() IBlockContext
	CLOSE_kEY() antlr.TerminalNode

	// IsGuard_blContext differentiates from other interfaces.
	IsGuard_blContext()
}

type Guard_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_GUARD      antlr.Token
	_expression IExpressionContext
	_block      IBlockContext
}

func NewEmptyGuard_blContext() *Guard_blContext {
	var p = new(Guard_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_guard_bl
	return p
}

func InitEmptyGuard_blContext(p *Guard_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_guard_bl
}

func (*Guard_blContext) IsGuard_blContext() {}

func NewGuard_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Guard_blContext {
	var p = new(Guard_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_guard_bl

	return p
}

func (s *Guard_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Guard_blContext) Get_GUARD() antlr.Token { return s._GUARD }

func (s *Guard_blContext) Set_GUARD(v antlr.Token) { s._GUARD = v }

func (s *Guard_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *Guard_blContext) Get_block() IBlockContext { return s._block }

func (s *Guard_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Guard_blContext) Set_block(v IBlockContext) { s._block = v }

func (s *Guard_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Guard_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Guard_blContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserGUARD, 0)
}

func (s *Guard_blContext) Expression() IExpressionContext {
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

func (s *Guard_blContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserELSE, 0)
}

func (s *Guard_blContext) OPEN_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, 0)
}

func (s *Guard_blContext) Block() IBlockContext {
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

func (s *Guard_blContext) CLOSE_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, 0)
}

func (s *Guard_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Guard_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Guard_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterGuard_bl(s)
	}
}

func (s *Guard_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitGuard_bl(s)
	}
}

func (s *Guard_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitGuard_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Guard_bl() (localctx IGuard_blContext) {
	localctx = NewGuard_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftgrammParserRULE_guard_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)

		var _m = p.Match(SwiftgrammParserGUARD)

		localctx.(*Guard_blContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)

		var _x = p.expression(0)

		localctx.(*Guard_blContext)._expression = _x
	}
	{
		p.SetState(357)
		p.Match(SwiftgrammParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)

		var _x = p.Block()

		localctx.(*Guard_blContext)._block = _x
	}
	{
		p.SetState(360)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*Guard_blContext).instr = instructions.NewGuard(localctx.(*Guard_blContext).Get_expression().GetP(), localctx.(*Guard_blContext).Get_block().GetBlk(), (func() int {
		if localctx.(*Guard_blContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*Guard_blContext).Get_GUARD().GetLine()
		}
	}()), (func() int {
		if localctx.(*Guard_blContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*Guard_blContext).Get_GUARD().GetColumn()
		}
	}()))

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

// IVector_blContext is an interface to support dynamic dispatch.
type IVector_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_datatype returns the _datatype rule contexts.
	Get_datatype() IDatatypeContext

	// Get_array_exp returns the _array_exp rule contexts.
	Get_array_exp() IArray_expContext

	// Set_datatype sets the _datatype rule contexts.
	Set_datatype(IDatatypeContext)

	// Set_array_exp sets the _array_exp rule contexts.
	Set_array_exp(IArray_expContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllOPEN_BRACKET() []antlr.TerminalNode
	OPEN_BRACKET(i int) antlr.TerminalNode
	Datatype() IDatatypeContext
	AllCLOSE_BRACKET() []antlr.TerminalNode
	CLOSE_BRACKET(i int) antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Array_exp() IArray_expContext

	// IsVector_blContext differentiates from other interfaces.
	IsVector_blContext()
}

type Vector_blContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	instr      abstract.Instruction
	_ID        antlr.Token
	_datatype  IDatatypeContext
	_array_exp IArray_expContext
}

func NewEmptyVector_blContext() *Vector_blContext {
	var p = new(Vector_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_vector_bl
	return p
}

func InitEmptyVector_blContext(p *Vector_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_vector_bl
}

func (*Vector_blContext) IsVector_blContext() {}

func NewVector_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_blContext {
	var p = new(Vector_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_vector_bl

	return p
}

func (s *Vector_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_blContext) Get_ID() antlr.Token { return s._ID }

func (s *Vector_blContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Vector_blContext) Get_datatype() IDatatypeContext { return s._datatype }

func (s *Vector_blContext) Get_array_exp() IArray_expContext { return s._array_exp }

func (s *Vector_blContext) Set_datatype(v IDatatypeContext) { s._datatype = v }

func (s *Vector_blContext) Set_array_exp(v IArray_expContext) { s._array_exp = v }

func (s *Vector_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Vector_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Vector_blContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserVAR, 0)
}

func (s *Vector_blContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Vector_blContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *Vector_blContext) AllOPEN_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammParserOPEN_BRACKET)
}

func (s *Vector_blContext) OPEN_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_BRACKET, i)
}

func (s *Vector_blContext) Datatype() IDatatypeContext {
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

func (s *Vector_blContext) AllCLOSE_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammParserCLOSE_BRACKET)
}

func (s *Vector_blContext) CLOSE_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_BRACKET, i)
}

func (s *Vector_blContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserASSIGN, 0)
}

func (s *Vector_blContext) Array_exp() IArray_expContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_expContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_expContext)
}

func (s *Vector_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vector_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterVector_bl(s)
	}
}

func (s *Vector_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitVector_bl(s)
	}
}

func (s *Vector_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitVector_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Vector_bl() (localctx IVector_blContext) {
	localctx = NewVector_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftgrammParserRULE_vector_bl)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Vector_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)

			var _x = p.Datatype()

			localctx.(*Vector_blContext)._datatype = _x
		}
		{
			p.SetState(368)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)

			var _x = p.Array_exp()

			localctx.(*Vector_blContext)._array_exp = _x
		}
		{
			p.SetState(372)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Vector_blContext).instr = instructions.NewVector((func() string {
			if localctx.(*Vector_blContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Vector_blContext).Get_ID().GetText()
			}
		}()), localctx.(*Vector_blContext).Get_datatype().GetTd(), localctx.(*Vector_blContext).Get_array_exp().GetP(), (func() int {
			if localctx.(*Vector_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Vector_blContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Vector_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Vector_blContext).Get_ID().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Vector_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)

			var _x = p.Datatype()

			localctx.(*Vector_blContext)._datatype = _x
		}
		{
			p.SetState(380)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Vector_blContext).instr = instructions.NewVector((func() string {
			if localctx.(*Vector_blContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Vector_blContext).Get_ID().GetText()
			}
		}()), localctx.(*Vector_blContext).Get_datatype().GetTd(), nil, (func() int {
			if localctx.(*Vector_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Vector_blContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Vector_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Vector_blContext).Get_ID().GetColumn()
			}
		}()))

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

// IArray_expContext is an interface to support dynamic dispatch.
type IArray_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_array_exp returns the _array_exp rule contexts.
	Get_array_exp() IArray_expContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_array_exp sets the _array_exp rule contexts.
	Set_array_exp(IArray_expContext)

	// GetP returns the p attribute.
	GetP() []interface{}

	// SetP sets the p attribute.
	SetP([]interface{})

	// Getter signatures
	Expression() IExpressionContext
	COMMA() antlr.TerminalNode
	Array_exp() IArray_expContext

	// IsArray_expContext differentiates from other interfaces.
	IsArray_expContext()
}

type Array_expContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           []interface{}
	_expression IExpressionContext
	_array_exp  IArray_expContext
}

func NewEmptyArray_expContext() *Array_expContext {
	var p = new(Array_expContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_array_exp
	return p
}

func InitEmptyArray_expContext(p *Array_expContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_array_exp
}

func (*Array_expContext) IsArray_expContext() {}

func NewArray_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_expContext {
	var p = new(Array_expContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_array_exp

	return p
}

func (s *Array_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_expContext) Get_expression() IExpressionContext { return s._expression }

func (s *Array_expContext) Get_array_exp() IArray_expContext { return s._array_exp }

func (s *Array_expContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Array_expContext) Set_array_exp(v IArray_expContext) { s._array_exp = v }

func (s *Array_expContext) GetP() []interface{} { return s.p }

func (s *Array_expContext) SetP(v []interface{}) { s.p = v }

func (s *Array_expContext) Expression() IExpressionContext {
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

func (s *Array_expContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOMMA, 0)
}

func (s *Array_expContext) Array_exp() IArray_expContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_expContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_expContext)
}

func (s *Array_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterArray_exp(s)
	}
}

func (s *Array_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitArray_exp(s)
	}
}

func (s *Array_expContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitArray_exp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Array_exp() (localctx IArray_expContext) {
	localctx = NewArray_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftgrammParserRULE_array_exp)
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)

			var _x = p.expression(0)

			localctx.(*Array_expContext)._expression = _x
		}
		{
			p.SetState(389)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)

			var _x = p.Array_exp()

			localctx.(*Array_expContext)._array_exp = _x
		}

		localctx.(*Array_expContext).p = append([]interface{}{localctx.(*Array_expContext).Get_expression().GetP()}, localctx.(*Array_expContext).Get_array_exp().GetP()...)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)

			var _x = p.expression(0)

			localctx.(*Array_expContext)._expression = _x
		}

		localctx.(*Array_expContext).p = []interface{}{localctx.(*Array_expContext).Get_expression().GetP()}

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

// IArray_functionsContext is an interface to support dynamic dispatch.
type IArray_functionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_APPEND returns the _APPEND token.
	Get_APPEND() antlr.Token

	// Get_REMOVELAST returns the _REMOVELAST token.
	Get_REMOVELAST() antlr.Token

	// Get_REMOVE returns the _REMOVE token.
	Get_REMOVE() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_APPEND sets the _APPEND token.
	Set_APPEND(antlr.Token)

	// Set_REMOVELAST sets the _REMOVELAST token.
	Set_REMOVELAST(antlr.Token)

	// Set_REMOVE sets the _REMOVE token.
	Set_REMOVE(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	DOT() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	OPEN_PARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	CLOSE_PARENTHESIS() antlr.TerminalNode
	REMOVELAST() antlr.TerminalNode
	REMOVE() antlr.TerminalNode
	AT() antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsArray_functionsContext differentiates from other interfaces.
	IsArray_functionsContext()
}

type Array_functionsContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_ID         antlr.Token
	_APPEND     antlr.Token
	_expression IExpressionContext
	_REMOVELAST antlr.Token
	_REMOVE     antlr.Token
}

func NewEmptyArray_functionsContext() *Array_functionsContext {
	var p = new(Array_functionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_array_functions
	return p
}

func InitEmptyArray_functionsContext(p *Array_functionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_array_functions
}

func (*Array_functionsContext) IsArray_functionsContext() {}

func NewArray_functionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_functionsContext {
	var p = new(Array_functionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_array_functions

	return p
}

func (s *Array_functionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_functionsContext) Get_ID() antlr.Token { return s._ID }

func (s *Array_functionsContext) Get_APPEND() antlr.Token { return s._APPEND }

func (s *Array_functionsContext) Get_REMOVELAST() antlr.Token { return s._REMOVELAST }

func (s *Array_functionsContext) Get_REMOVE() antlr.Token { return s._REMOVE }

func (s *Array_functionsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Array_functionsContext) Set_APPEND(v antlr.Token) { s._APPEND = v }

func (s *Array_functionsContext) Set_REMOVELAST(v antlr.Token) { s._REMOVELAST = v }

func (s *Array_functionsContext) Set_REMOVE(v antlr.Token) { s._REMOVE = v }

func (s *Array_functionsContext) Get_expression() IExpressionContext { return s._expression }

func (s *Array_functionsContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Array_functionsContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Array_functionsContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Array_functionsContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Array_functionsContext) DOT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserDOT, 0)
}

func (s *Array_functionsContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserAPPEND, 0)
}

func (s *Array_functionsContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *Array_functionsContext) Expression() IExpressionContext {
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

func (s *Array_functionsContext) CLOSE_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_PARENTHESIS, 0)
}

func (s *Array_functionsContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserREMOVELAST, 0)
}

func (s *Array_functionsContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserREMOVE, 0)
}

func (s *Array_functionsContext) AT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserAT, 0)
}

func (s *Array_functionsContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *Array_functionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_functionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_functionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterArray_functions(s)
	}
}

func (s *Array_functionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitArray_functions(s)
	}
}

func (s *Array_functionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitArray_functions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Array_functions() (localctx IArray_functionsContext) {
	localctx = NewArray_functionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftgrammParserRULE_array_functions)
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Array_functionsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.Match(SwiftgrammParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)

			var _m = p.Match(SwiftgrammParserAPPEND)

			localctx.(*Array_functionsContext)._APPEND = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)

			var _x = p.expression(0)

			localctx.(*Array_functionsContext)._expression = _x
		}
		{
			p.SetState(403)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Array_functionsContext).instr = instructions.NewArrayFunctions((func() string {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*Array_functionsContext).Get_APPEND() == nil {
				return ""
			} else {
				return localctx.(*Array_functionsContext).Get_APPEND().GetText()
			}
		}()), localctx.(*Array_functionsContext).Get_expression().GetP(), (func() int {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Array_functionsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(SwiftgrammParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)

			var _m = p.Match(SwiftgrammParserREMOVELAST)

			localctx.(*Array_functionsContext)._REMOVELAST = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Array_functionsContext).instr = instructions.NewArrayFunctions((func() string {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*Array_functionsContext).Get_REMOVELAST() == nil {
				return ""
			} else {
				return localctx.(*Array_functionsContext).Get_REMOVELAST().GetText()
			}
		}()), nil, (func() int {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetColumn()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(412)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Array_functionsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Match(SwiftgrammParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)

			var _m = p.Match(SwiftgrammParserREMOVE)

			localctx.(*Array_functionsContext)._REMOVE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Match(SwiftgrammParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)

			var _x = p.expression(0)

			localctx.(*Array_functionsContext)._expression = _x
		}
		{
			p.SetState(419)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Array_functionsContext).instr = instructions.NewArrayFunctions((func() string {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*Array_functionsContext).Get_REMOVE() == nil {
				return ""
			} else {
				return localctx.(*Array_functionsContext).Get_REMOVE().GetText()
			}
		}()), localctx.(*Array_functionsContext).Get_expression().GetP(), (func() int {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Array_functionsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Array_functionsContext).Get_ID().GetColumn()
			}
		}()))

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

// IAssign_vectorContext is an interface to support dynamic dispatch.
type IAssign_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	CLOSE_BRACKET() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode

	// IsAssign_vectorContext differentiates from other interfaces.
	IsAssign_vectorContext()
}

type Assign_vectorContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_ID         antlr.Token
	_expression IExpressionContext
}

func NewEmptyAssign_vectorContext() *Assign_vectorContext {
	var p = new(Assign_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_assign_vector
	return p
}

func InitEmptyAssign_vectorContext(p *Assign_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_assign_vector
}

func (*Assign_vectorContext) IsAssign_vectorContext() {}

func NewAssign_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_vectorContext {
	var p = new(Assign_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_assign_vector

	return p
}

func (s *Assign_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_vectorContext) Get_ID() antlr.Token { return s._ID }

func (s *Assign_vectorContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Assign_vectorContext) Get_expression() IExpressionContext { return s._expression }

func (s *Assign_vectorContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Assign_vectorContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Assign_vectorContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Assign_vectorContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Assign_vectorContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_BRACKET, 0)
}

func (s *Assign_vectorContext) AllExpression() []IExpressionContext {
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

func (s *Assign_vectorContext) Expression(i int) IExpressionContext {
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

func (s *Assign_vectorContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_BRACKET, 0)
}

func (s *Assign_vectorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserASSIGN, 0)
}

func (s *Assign_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_vectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterAssign_vector(s)
	}
}

func (s *Assign_vectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitAssign_vector(s)
	}
}

func (s *Assign_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitAssign_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Assign_vector() (localctx IAssign_vectorContext) {
	localctx = NewAssign_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftgrammParserRULE_assign_vector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(424)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Assign_vectorContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(425)
		p.Match(SwiftgrammParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)

		var _x = p.expression(0)

		localctx.(*Assign_vectorContext)._expression = _x
	}
	{
		p.SetState(427)
		p.Match(SwiftgrammParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.Match(SwiftgrammParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(429)

		var _x = p.expression(0)

		localctx.(*Assign_vectorContext)._expression = _x
	}

	localctx.(*Assign_vectorContext).instr = instructions.NewAsignVector((func() string {
		if localctx.(*Assign_vectorContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Assign_vectorContext).Get_ID().GetText()
		}
	}()), localctx.(*Assign_vectorContext).Get_expression().GetP(), localctx.(*Assign_vectorContext).Get_expression().GetP(), (func() int {
		if localctx.(*Assign_vectorContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Assign_vectorContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*Assign_vectorContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Assign_vectorContext).Get_ID().GetColumn()
		}
	}()))

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

// IFunction_blContext is an interface to support dynamic dispatch.
type IFunction_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_datatype returns the _datatype rule contexts.
	Get_datatype() IDatatypeContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_params returns the _params rule contexts.
	Get_params() IParamsContext

	// Set_datatype sets the _datatype rule contexts.
	Set_datatype(IDatatypeContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_params sets the _params rule contexts.
	Set_params(IParamsContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	OPEN_PARENTHESIS() antlr.TerminalNode
	CLOSE_PARENTHESIS() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	Datatype() IDatatypeContext
	OPEN_kEY() antlr.TerminalNode
	Block() IBlockContext
	CLOSE_kEY() antlr.TerminalNode
	Params() IParamsContext

	// IsFunction_blContext differentiates from other interfaces.
	IsFunction_blContext()
}

type Function_blContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	instr     abstract.Instruction
	_ID       antlr.Token
	_datatype IDatatypeContext
	_block    IBlockContext
	_params   IParamsContext
}

func NewEmptyFunction_blContext() *Function_blContext {
	var p = new(Function_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_function_bl
	return p
}

func InitEmptyFunction_blContext(p *Function_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_function_bl
}

func (*Function_blContext) IsFunction_blContext() {}

func NewFunction_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_blContext {
	var p = new(Function_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_function_bl

	return p
}

func (s *Function_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_blContext) Get_ID() antlr.Token { return s._ID }

func (s *Function_blContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Function_blContext) Get_datatype() IDatatypeContext { return s._datatype }

func (s *Function_blContext) Get_block() IBlockContext { return s._block }

func (s *Function_blContext) Get_params() IParamsContext { return s._params }

func (s *Function_blContext) Set_datatype(v IDatatypeContext) { s._datatype = v }

func (s *Function_blContext) Set_block(v IBlockContext) { s._block = v }

func (s *Function_blContext) Set_params(v IParamsContext) { s._params = v }

func (s *Function_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Function_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Function_blContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserFUNC, 0)
}

func (s *Function_blContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Function_blContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *Function_blContext) CLOSE_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_PARENTHESIS, 0)
}

func (s *Function_blContext) ARROW() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserARROW, 0)
}

func (s *Function_blContext) Datatype() IDatatypeContext {
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

func (s *Function_blContext) OPEN_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_kEY, 0)
}

func (s *Function_blContext) Block() IBlockContext {
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

func (s *Function_blContext) CLOSE_kEY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_kEY, 0)
}

func (s *Function_blContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *Function_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterFunction_bl(s)
	}
}

func (s *Function_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitFunction_bl(s)
	}
}

func (s *Function_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitFunction_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Function_bl() (localctx IFunction_blContext) {
	localctx = NewFunction_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftgrammParserRULE_function_bl)
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)
			p.Match(SwiftgrammParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Function_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.Match(SwiftgrammParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)

			var _x = p.Datatype()

			localctx.(*Function_blContext)._datatype = _x
		}
		{
			p.SetState(438)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)

			var _x = p.Block()

			localctx.(*Function_blContext)._block = _x
		}
		{
			p.SetState(440)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Function_blContext).instr = instructions.NewDeclareFunction((func() string {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Function_blContext).Get_ID().GetText()
			}
		}()), localctx.(*Function_blContext).Get_datatype().GetTd(), []interface{}{}, localctx.(*Function_blContext).Get_block().GetBlk(), (func() int {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Function_blContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Function_blContext).Get_ID().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)
			p.Match(SwiftgrammParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Function_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)

			var _x = p.Block()

			localctx.(*Function_blContext)._block = _x
		}
		{
			p.SetState(449)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Function_blContext).instr = instructions.NewDeclareFunction((func() string {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Function_blContext).Get_ID().GetText()
			}
		}()), symbol.NIL, []interface{}{}, localctx.(*Function_blContext).Get_block().GetBlk(), (func() int {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Function_blContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Function_blContext).Get_ID().GetColumn()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(452)
			p.Match(SwiftgrammParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Function_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)

			var _x = p.Params()

			localctx.(*Function_blContext)._params = _x
		}
		{
			p.SetState(456)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Match(SwiftgrammParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)

			var _x = p.Datatype()

			localctx.(*Function_blContext)._datatype = _x
		}
		{
			p.SetState(459)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)

			var _x = p.Block()

			localctx.(*Function_blContext)._block = _x
		}
		{
			p.SetState(461)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Function_blContext).instr = instructions.NewDeclareFunction((func() string {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Function_blContext).Get_ID().GetText()
			}
		}()), localctx.(*Function_blContext).Get_datatype().GetTd(), localctx.(*Function_blContext).Get_params().GetP(), localctx.(*Function_blContext).Get_block().GetBlk(), (func() int {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Function_blContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Function_blContext).Get_ID().GetColumn()
			}
		}()))

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(464)
			p.Match(SwiftgrammParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Function_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)

			var _x = p.Params()

			localctx.(*Function_blContext)._params = _x
		}
		{
			p.SetState(468)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)

			var _x = p.Block()

			localctx.(*Function_blContext)._block = _x
		}
		{
			p.SetState(471)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		//fmt.Println(localctx.(*Function_blContext).Get_params().GetP())
		localctx.(*Function_blContext).instr = instructions.NewDeclareFunction((func() string {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Function_blContext).Get_ID().GetText()
			}
		}()), symbol.NIL, localctx.(*Function_blContext).Get_params().GetP(), localctx.(*Function_blContext).Get_block().GetBlk(), (func() int {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Function_blContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Function_blContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Function_blContext).Get_ID().GetColumn()
			}
		}()))

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

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_datatype returns the _datatype rule contexts.
	Get_datatype() IDatatypeContext

	// Get_params returns the _params rule contexts.
	Get_params() IParamsContext

	// Get_extern_params returns the _extern_params rule contexts.
	Get_extern_params() IExtern_paramsContext

	// Set_datatype sets the _datatype rule contexts.
	Set_datatype(IDatatypeContext)

	// Set_params sets the _params rule contexts.
	Set_params(IParamsContext)

	// Set_extern_params sets the _extern_params rule contexts.
	Set_extern_params(IExtern_paramsContext)

	// GetP returns the p attribute.
	GetP() []interface{}

	// SetP sets the p attribute.
	SetP([]interface{})

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Datatype() IDatatypeContext
	COMMA() antlr.TerminalNode
	Params() IParamsContext
	Extern_params() IExtern_paramsContext
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	p              []interface{}
	_ID            antlr.Token
	_datatype      IDatatypeContext
	_params        IParamsContext
	_extern_params IExtern_paramsContext
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) Get_ID() antlr.Token { return s._ID }

func (s *ParamsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ParamsContext) Get_datatype() IDatatypeContext { return s._datatype }

func (s *ParamsContext) Get_params() IParamsContext { return s._params }

func (s *ParamsContext) Get_extern_params() IExtern_paramsContext { return s._extern_params }

func (s *ParamsContext) Set_datatype(v IDatatypeContext) { s._datatype = v }

func (s *ParamsContext) Set_params(v IParamsContext) { s._params = v }

func (s *ParamsContext) Set_extern_params(v IExtern_paramsContext) { s._extern_params = v }

func (s *ParamsContext) GetP() []interface{} { return s.p }

func (s *ParamsContext) SetP(v []interface{}) { s.p = v }

func (s *ParamsContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *ParamsContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *ParamsContext) Datatype() IDatatypeContext {
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

func (s *ParamsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOMMA, 0)
}

func (s *ParamsContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *ParamsContext) Extern_params() IExtern_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtern_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtern_paramsContext)
}

func (s *ParamsContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_BRACKET, 0)
}

func (s *ParamsContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_BRACKET, 0)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SwiftgrammParserRULE_params)
	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}
		{
			p.SetState(479)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)

			var _x = p.Params()

			localctx.(*ParamsContext)._params = _x
		}

		value := instructions.NewParamFunction(instructions.NewDeclareWithoutValue((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), expressions.NewNative(nil, symbol.NIL, (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()))
		localctx.(*ParamsContext).p = append([]interface{}{value}, localctx.(*ParamsContext).Get_params().GetP()...)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(483)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}

		value := instructions.NewParamFunction(instructions.NewDeclareWithoutValue((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), expressions.NewNative(nil, symbol.NIL, (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()))
		localctx.(*ParamsContext).p = []interface{}{value}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(488)

			var _x = p.Extern_params()

			localctx.(*ParamsContext)._extern_params = _x
		}
		{
			p.SetState(489)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}
		{
			p.SetState(492)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)

			var _x = p.Params()

			localctx.(*ParamsContext)._params = _x
		}

		value := instructions.NewParamFunction(instructions.NewDeclareWithoutValue((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), expressions.NewNative(nil, symbol.NIL, (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() string {
			if localctx.(*ParamsContext).Get_extern_params() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*ParamsContext).Get_extern_params().GetStart(), localctx.(*ParamsContext)._extern_params.GetStop())
			}
		}()))
		localctx.(*ParamsContext).p = append([]interface{}{value}, localctx.(*ParamsContext).Get_params().GetP()...)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(496)

			var _x = p.Extern_params()

			localctx.(*ParamsContext)._extern_params = _x
		}
		{
			p.SetState(497)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(499)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}

		value := instructions.NewParamFunction(instructions.NewDeclareWithoutValue((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), expressions.NewNative(nil, symbol.NIL, (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() string {
			if localctx.(*ParamsContext).Get_extern_params() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*ParamsContext).Get_extern_params().GetStart(), localctx.(*ParamsContext)._extern_params.GetStop())
			}
		}()))
		localctx.(*ParamsContext).p = []interface{}{value}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(502)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}
		{
			p.SetState(506)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(507)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(508)

			var _x = p.Params()

			localctx.(*ParamsContext)._params = _x
		}

		value := instructions.NewParamFunction(instructions.NewVector((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), nil, (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()))
		localctx.(*ParamsContext).p = append([]interface{}{value}, localctx.(*ParamsContext).Get_params().GetP()...)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(511)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(513)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(514)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}
		{
			p.SetState(515)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value := instructions.NewParamFunction(instructions.NewVector((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), nil, (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()))
		localctx.(*ParamsContext).p = []interface{}{value}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(518)

			var _x = p.Extern_params()

			localctx.(*ParamsContext)._extern_params = _x
		}
		{
			p.SetState(519)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(521)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}
		{
			p.SetState(523)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)

			var _x = p.Params()

			localctx.(*ParamsContext)._params = _x
		}

		value := instructions.NewParamFunction(instructions.NewVector((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), nil, (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() string {
			if localctx.(*ParamsContext).Get_extern_params() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*ParamsContext).Get_extern_params().GetStart(), localctx.(*ParamsContext)._extern_params.GetStop())
			}
		}()))
		localctx.(*ParamsContext).p = append([]interface{}{value}, localctx.(*ParamsContext).Get_params().GetP()...)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(528)

			var _x = p.Extern_params()

			localctx.(*ParamsContext)._extern_params = _x
		}
		{
			p.SetState(529)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(531)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}
		{
			p.SetState(533)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value := instructions.NewParamFunction(instructions.NewVector((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), nil, (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParamsContext).Get_ID().GetColumn()
			}
		}())), (func() string {
			if localctx.(*ParamsContext).Get_extern_params() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*ParamsContext).Get_extern_params().GetStart(), localctx.(*ParamsContext)._extern_params.GetStop())
			}
		}()))
		localctx.(*ParamsContext).p = []interface{}{value}

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

// IExtern_paramsContext is an interface to support dynamic dispatch.
type IExtern_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_UNDERSCORE returns the _UNDERSCORE token.
	Get_UNDERSCORE() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_UNDERSCORE sets the _UNDERSCORE token.
	Set_UNDERSCORE(antlr.Token)

	// GetP returns the p attribute.
	GetP() string

	// SetP sets the p attribute.
	SetP(string)

	// Getter signatures
	ID() antlr.TerminalNode
	UNDERSCORE() antlr.TerminalNode

	// IsExtern_paramsContext differentiates from other interfaces.
	IsExtern_paramsContext()
}

type Extern_paramsContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           string
	_ID         antlr.Token
	_UNDERSCORE antlr.Token
}

func NewEmptyExtern_paramsContext() *Extern_paramsContext {
	var p = new(Extern_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_extern_params
	return p
}

func InitEmptyExtern_paramsContext(p *Extern_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_extern_params
}

func (*Extern_paramsContext) IsExtern_paramsContext() {}

func NewExtern_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extern_paramsContext {
	var p = new(Extern_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_extern_params

	return p
}

func (s *Extern_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Extern_paramsContext) Get_ID() antlr.Token { return s._ID }

func (s *Extern_paramsContext) Get_UNDERSCORE() antlr.Token { return s._UNDERSCORE }

func (s *Extern_paramsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Extern_paramsContext) Set_UNDERSCORE(v antlr.Token) { s._UNDERSCORE = v }

func (s *Extern_paramsContext) GetP() string { return s.p }

func (s *Extern_paramsContext) SetP(v string) { s.p = v }

func (s *Extern_paramsContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Extern_paramsContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserUNDERSCORE, 0)
}

func (s *Extern_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extern_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extern_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterExtern_params(s)
	}
}

func (s *Extern_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitExtern_params(s)
	}
}

func (s *Extern_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitExtern_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Extern_params() (localctx IExtern_paramsContext) {
	localctx = NewExtern_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftgrammParserRULE_extern_params)
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(538)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Extern_paramsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Extern_paramsContext).p = (func() string {
			if localctx.(*Extern_paramsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Extern_paramsContext).Get_ID().GetText()
			}
		}())

	case SwiftgrammParserUNDERSCORE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(540)

			var _m = p.Match(SwiftgrammParserUNDERSCORE)

			localctx.(*Extern_paramsContext)._UNDERSCORE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Extern_paramsContext).p = (func() string {
			if localctx.(*Extern_paramsContext).Get_UNDERSCORE() == nil {
				return ""
			} else {
				return localctx.(*Extern_paramsContext).Get_UNDERSCORE().GetText()
			}
		}())

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

// ICall_functionContext is an interface to support dynamic dispatch.
type ICall_functionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_list_exp returns the _list_exp rule contexts.
	Get_list_exp() IList_expContext

	// Set_list_exp sets the _list_exp rule contexts.
	Set_list_exp(IList_expContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	OPEN_PARENTHESIS() antlr.TerminalNode
	CLOSE_PARENTHESIS() antlr.TerminalNode
	List_exp() IList_expContext

	// IsCall_functionContext differentiates from other interfaces.
	IsCall_functionContext()
}

type Call_functionContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	instr     abstract.Instruction
	_ID       antlr.Token
	_list_exp IList_expContext
}

func NewEmptyCall_functionContext() *Call_functionContext {
	var p = new(Call_functionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_call_function
	return p
}

func InitEmptyCall_functionContext(p *Call_functionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_call_function
}

func (*Call_functionContext) IsCall_functionContext() {}

func NewCall_functionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_functionContext {
	var p = new(Call_functionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_call_function

	return p
}

func (s *Call_functionContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_functionContext) Get_ID() antlr.Token { return s._ID }

func (s *Call_functionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Call_functionContext) Get_list_exp() IList_expContext { return s._list_exp }

func (s *Call_functionContext) Set_list_exp(v IList_expContext) { s._list_exp = v }

func (s *Call_functionContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Call_functionContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Call_functionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Call_functionContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *Call_functionContext) CLOSE_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_PARENTHESIS, 0)
}

func (s *Call_functionContext) List_exp() IList_expContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_expContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_expContext)
}

func (s *Call_functionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_functionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_functionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterCall_function(s)
	}
}

func (s *Call_functionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitCall_function(s)
	}
}

func (s *Call_functionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitCall_function(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Call_function() (localctx ICall_functionContext) {
	localctx = NewCall_functionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SwiftgrammParserRULE_call_function)
	p.SetState(554)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(544)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Call_functionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(546)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Call_functionContext).instr = instructions.NewCallFunction((func() string {
			if localctx.(*Call_functionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Call_functionContext).Get_ID().GetText()
			}
		}()), []interface{}{}, (func() int {
			if localctx.(*Call_functionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Call_functionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Call_functionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Call_functionContext).Get_ID().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(548)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Call_functionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(550)

			var _x = p.List_exp()

			localctx.(*Call_functionContext)._list_exp = _x
		}
		{
			p.SetState(551)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Call_functionContext).instr = instructions.NewCallFunction((func() string {
			if localctx.(*Call_functionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Call_functionContext).Get_ID().GetText()
			}
		}()), localctx.(*Call_functionContext).Get_list_exp().GetP(), (func() int {
			if localctx.(*Call_functionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Call_functionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Call_functionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Call_functionContext).Get_ID().GetColumn()
			}
		}()))

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

// IList_expContext is an interface to support dynamic dispatch.
type IList_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_list_exp returns the _list_exp rule contexts.
	Get_list_exp() IList_expContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_list_exp sets the _list_exp rule contexts.
	Set_list_exp(IList_expContext)

	// GetP returns the p attribute.
	GetP() []interface{}

	// SetP sets the p attribute.
	SetP([]interface{})

	// Getter signatures
	Expression() IExpressionContext
	COMMA() antlr.TerminalNode
	List_exp() IList_expContext
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsList_expContext differentiates from other interfaces.
	IsList_expContext()
}

type List_expContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           []interface{}
	_expression IExpressionContext
	_list_exp   IList_expContext
	_ID         antlr.Token
}

func NewEmptyList_expContext() *List_expContext {
	var p = new(List_expContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_list_exp
	return p
}

func InitEmptyList_expContext(p *List_expContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_list_exp
}

func (*List_expContext) IsList_expContext() {}

func NewList_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expContext {
	var p = new(List_expContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_list_exp

	return p
}

func (s *List_expContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expContext) Get_ID() antlr.Token { return s._ID }

func (s *List_expContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *List_expContext) Get_expression() IExpressionContext { return s._expression }

func (s *List_expContext) Get_list_exp() IList_expContext { return s._list_exp }

func (s *List_expContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *List_expContext) Set_list_exp(v IList_expContext) { s._list_exp = v }

func (s *List_expContext) GetP() []interface{} { return s.p }

func (s *List_expContext) SetP(v []interface{}) { s.p = v }

func (s *List_expContext) Expression() IExpressionContext {
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

func (s *List_expContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOMMA, 0)
}

func (s *List_expContext) List_exp() IList_expContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_expContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_expContext)
}

func (s *List_expContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *List_expContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *List_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterList_exp(s)
	}
}

func (s *List_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitList_exp(s)
	}
}

func (s *List_expContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitList_exp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) List_exp() (localctx IList_expContext) {
	localctx = NewList_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftgrammParserRULE_list_exp)
	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(556)

			var _x = p.expression(0)

			localctx.(*List_expContext)._expression = _x
		}
		{
			p.SetState(557)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)

			var _x = p.List_exp()

			localctx.(*List_expContext)._list_exp = _x
		}

		value := instructions.NewParamCall(localctx.(*List_expContext).Get_expression().GetP(), "_")
		localctx.(*List_expContext).p = append([]interface{}{value}, localctx.(*List_expContext).Get_list_exp().GetP()...)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(561)

			var _x = p.expression(0)

			localctx.(*List_expContext)._expression = _x
		}

		value := instructions.NewParamCall(localctx.(*List_expContext).Get_expression().GetP(), "_")
		localctx.(*List_expContext).p = []interface{}{value}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(564)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*List_expContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(565)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)

			var _x = p.expression(0)

			localctx.(*List_expContext)._expression = _x
		}
		{
			p.SetState(567)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)

			var _x = p.List_exp()

			localctx.(*List_expContext)._list_exp = _x
		}

		value := instructions.NewParamCall(localctx.(*List_expContext).Get_expression().GetP(), (func() string {
			if localctx.(*List_expContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*List_expContext).Get_ID().GetText()
			}
		}()))
		localctx.(*List_expContext).p = append([]interface{}{value}, localctx.(*List_expContext).Get_list_exp().GetP()...)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(571)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*List_expContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(572)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)

			var _x = p.expression(0)

			localctx.(*List_expContext)._expression = _x
		}

		value := instructions.NewParamCall(localctx.(*List_expContext).Get_expression().GetP(), (func() string {
			if localctx.(*List_expContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*List_expContext).Get_ID().GetText()
			}
		}()))
		localctx.(*List_expContext).p = []interface{}{value}

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

// ICall_function_expContext is an interface to support dynamic dispatch.
type ICall_function_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_call_function returns the _call_function rule contexts.
	Get_call_function() ICall_functionContext

	// Set_call_function sets the _call_function rule contexts.
	Set_call_function(ICall_functionContext)

	// GetP returns the p attribute.
	GetP() abstract.Expression

	// SetP sets the p attribute.
	SetP(abstract.Expression)

	// Getter signatures
	Call_function() ICall_functionContext

	// IsCall_function_expContext differentiates from other interfaces.
	IsCall_function_expContext()
}

type Call_function_expContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	p              abstract.Expression
	_call_function ICall_functionContext
}

func NewEmptyCall_function_expContext() *Call_function_expContext {
	var p = new(Call_function_expContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_call_function_exp
	return p
}

func InitEmptyCall_function_expContext(p *Call_function_expContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_call_function_exp
}

func (*Call_function_expContext) IsCall_function_expContext() {}

func NewCall_function_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_function_expContext {
	var p = new(Call_function_expContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_call_function_exp

	return p
}

func (s *Call_function_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_function_expContext) Get_call_function() ICall_functionContext { return s._call_function }

func (s *Call_function_expContext) Set_call_function(v ICall_functionContext) { s._call_function = v }

func (s *Call_function_expContext) GetP() abstract.Expression { return s.p }

func (s *Call_function_expContext) SetP(v abstract.Expression) { s.p = v }

func (s *Call_function_expContext) Call_function() ICall_functionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_functionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_functionContext)
}

func (s *Call_function_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_function_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_function_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterCall_function_exp(s)
	}
}

func (s *Call_function_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitCall_function_exp(s)
	}
}

func (s *Call_function_expContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitCall_function_exp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Call_function_exp() (localctx ICall_function_expContext) {
	localctx = NewCall_function_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftgrammParserRULE_call_function_exp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(578)

		var _x = p.Call_function()

		localctx.(*Call_function_expContext)._call_function = _x
	}

	localctx.(*Call_function_expContext).p = expressions.NewCallFunctionExp(localctx.(*Call_function_expContext).Get_call_function().GetInstr())

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

// IDeclare_array_blContext is an interface to support dynamic dispatch.
type IDeclare_array_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_type_matrix returns the _type_matrix rule contexts.
	Get_type_matrix() IType_matrixContext

	// Get_exp_matriz returns the _exp_matriz rule contexts.
	Get_exp_matriz() IExp_matrizContext

	// Set_type_matrix sets the _type_matrix rule contexts.
	Set_type_matrix(IType_matrixContext)

	// Set_exp_matriz sets the _exp_matriz rule contexts.
	Set_exp_matriz(IExp_matrizContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_matrix() IType_matrixContext
	ASSIGN() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	Exp_matriz() IExp_matrizContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsDeclare_array_blContext differentiates from other interfaces.
	IsDeclare_array_blContext()
}

type Declare_array_blContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        abstract.Instruction
	_ID          antlr.Token
	_type_matrix IType_matrixContext
	_exp_matriz  IExp_matrizContext
}

func NewEmptyDeclare_array_blContext() *Declare_array_blContext {
	var p = new(Declare_array_blContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_declare_array_bl
	return p
}

func InitEmptyDeclare_array_blContext(p *Declare_array_blContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_declare_array_bl
}

func (*Declare_array_blContext) IsDeclare_array_blContext() {}

func NewDeclare_array_blContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declare_array_blContext {
	var p = new(Declare_array_blContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_declare_array_bl

	return p
}

func (s *Declare_array_blContext) GetParser() antlr.Parser { return s.parser }

func (s *Declare_array_blContext) Get_ID() antlr.Token { return s._ID }

func (s *Declare_array_blContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Declare_array_blContext) Get_type_matrix() IType_matrixContext { return s._type_matrix }

func (s *Declare_array_blContext) Get_exp_matriz() IExp_matrizContext { return s._exp_matriz }

func (s *Declare_array_blContext) Set_type_matrix(v IType_matrixContext) { s._type_matrix = v }

func (s *Declare_array_blContext) Set_exp_matriz(v IExp_matrizContext) { s._exp_matriz = v }

func (s *Declare_array_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Declare_array_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Declare_array_blContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserVAR, 0)
}

func (s *Declare_array_blContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *Declare_array_blContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOLON, 0)
}

func (s *Declare_array_blContext) Type_matrix() IType_matrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_matrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_matrixContext)
}

func (s *Declare_array_blContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserASSIGN, 0)
}

func (s *Declare_array_blContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_BRACKET, 0)
}

func (s *Declare_array_blContext) Exp_matriz() IExp_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_matrizContext)
}

func (s *Declare_array_blContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_BRACKET, 0)
}

func (s *Declare_array_blContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declare_array_blContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declare_array_blContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterDeclare_array_bl(s)
	}
}

func (s *Declare_array_blContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitDeclare_array_bl(s)
	}
}

func (s *Declare_array_blContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitDeclare_array_bl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Declare_array_bl() (localctx IDeclare_array_blContext) {
	localctx = NewDeclare_array_blContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SwiftgrammParserRULE_declare_array_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(581)
		p.Match(SwiftgrammParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(582)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Declare_array_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(583)
		p.Match(SwiftgrammParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(584)

		var _x = p.Type_matrix()

		localctx.(*Declare_array_blContext)._type_matrix = _x
	}
	{
		p.SetState(585)
		p.Match(SwiftgrammParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(586)
		p.Match(SwiftgrammParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(587)

		var _x = p.Exp_matriz()

		localctx.(*Declare_array_blContext)._exp_matriz = _x
	}
	{
		p.SetState(588)
		p.Match(SwiftgrammParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*Declare_array_blContext).instr = instructions.NewDeclareArray((func() string {
		if localctx.(*Declare_array_blContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Declare_array_blContext).Get_ID().GetText()
		}
	}()), localctx.(*Declare_array_blContext).Get_type_matrix().GetTd(), CountM, localctx.(*Declare_array_blContext).Get_exp_matriz().GetP())

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

// IExp_matrizContext is an interface to support dynamic dispatch.
type IExp_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_exp_matriz returns the _exp_matriz rule contexts.
	Get_exp_matriz() IExp_matrizContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_exp_matriz sets the _exp_matriz rule contexts.
	Set_exp_matriz(IExp_matrizContext)

	// GetP returns the p attribute.
	GetP() []interface{}

	// SetP sets the p attribute.
	SetP([]interface{})

	// Getter signatures
	OPEN_BRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	CLOSE_BRACKET() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	Exp_matriz() IExp_matrizContext

	// IsExp_matrizContext differentiates from other interfaces.
	IsExp_matrizContext()
}

type Exp_matrizContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           []interface{}
	_expression IExpressionContext
	_exp_matriz IExp_matrizContext
}

func NewEmptyExp_matrizContext() *Exp_matrizContext {
	var p = new(Exp_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_exp_matriz
	return p
}

func InitEmptyExp_matrizContext(p *Exp_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_exp_matriz
}

func (*Exp_matrizContext) IsExp_matrizContext() {}

func NewExp_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_matrizContext {
	var p = new(Exp_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_exp_matriz

	return p
}

func (s *Exp_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_matrizContext) Get_expression() IExpressionContext { return s._expression }

func (s *Exp_matrizContext) Get_exp_matriz() IExp_matrizContext { return s._exp_matriz }

func (s *Exp_matrizContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Exp_matrizContext) Set_exp_matriz(v IExp_matrizContext) { s._exp_matriz = v }

func (s *Exp_matrizContext) GetP() []interface{} { return s.p }

func (s *Exp_matrizContext) SetP(v []interface{}) { s.p = v }

func (s *Exp_matrizContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_BRACKET, 0)
}

func (s *Exp_matrizContext) Expression() IExpressionContext {
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

func (s *Exp_matrizContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_BRACKET, 0)
}

func (s *Exp_matrizContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOMMA, 0)
}

func (s *Exp_matrizContext) Exp_matriz() IExp_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_matrizContext)
}

func (s *Exp_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_matrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterExp_matriz(s)
	}
}

func (s *Exp_matrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitExp_matriz(s)
	}
}

func (s *Exp_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitExp_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Exp_matriz() (localctx IExp_matrizContext) {
	localctx = NewExp_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SwiftgrammParserRULE_exp_matriz)
	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(591)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(592)

			var _x = p.expression(0)

			localctx.(*Exp_matrizContext)._expression = _x
		}
		{
			p.SetState(593)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value := localctx.(*Exp_matrizContext).Get_expression().GetP()
		localctx.(*Exp_matrizContext).p = []interface{}{value}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(596)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(597)

			var _x = p.expression(0)

			localctx.(*Exp_matrizContext)._expression = _x
		}
		{
			p.SetState(598)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(599)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(600)

			var _x = p.Exp_matriz()

			localctx.(*Exp_matrizContext)._exp_matriz = _x
		}

		value := localctx.(*Exp_matrizContext).Get_expression().GetP()
		localctx.(*Exp_matrizContext).p = append([]interface{}{value}, localctx.(*Exp_matrizContext).Get_exp_matriz().GetP()...)

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

// IType_matrixContext is an interface to support dynamic dispatch.
type IType_matrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_type_matrix returns the _type_matrix rule contexts.
	Get_type_matrix() IType_matrixContext

	// Get_datatype returns the _datatype rule contexts.
	Get_datatype() IDatatypeContext

	// Set_type_matrix sets the _type_matrix rule contexts.
	Set_type_matrix(IType_matrixContext)

	// Set_datatype sets the _datatype rule contexts.
	Set_datatype(IDatatypeContext)

	// GetTd returns the td attribute.
	GetTd() symbol.TypeData

	// SetTd sets the td attribute.
	SetTd(symbol.TypeData)

	// Getter signatures
	OPEN_BRACKET() antlr.TerminalNode
	Type_matrix() IType_matrixContext
	CLOSE_BRACKET() antlr.TerminalNode
	Datatype() IDatatypeContext

	// IsType_matrixContext differentiates from other interfaces.
	IsType_matrixContext()
}

type Type_matrixContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	td           symbol.TypeData
	_type_matrix IType_matrixContext
	_datatype    IDatatypeContext
}

func NewEmptyType_matrixContext() *Type_matrixContext {
	var p = new(Type_matrixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_type_matrix
	return p
}

func InitEmptyType_matrixContext(p *Type_matrixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_type_matrix
}

func (*Type_matrixContext) IsType_matrixContext() {}

func NewType_matrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_matrixContext {
	var p = new(Type_matrixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_type_matrix

	return p
}

func (s *Type_matrixContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_matrixContext) Get_type_matrix() IType_matrixContext { return s._type_matrix }

func (s *Type_matrixContext) Get_datatype() IDatatypeContext { return s._datatype }

func (s *Type_matrixContext) Set_type_matrix(v IType_matrixContext) { s._type_matrix = v }

func (s *Type_matrixContext) Set_datatype(v IDatatypeContext) { s._datatype = v }

func (s *Type_matrixContext) GetTd() symbol.TypeData { return s.td }

func (s *Type_matrixContext) SetTd(v symbol.TypeData) { s.td = v }

func (s *Type_matrixContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_BRACKET, 0)
}

func (s *Type_matrixContext) Type_matrix() IType_matrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_matrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_matrixContext)
}

func (s *Type_matrixContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_BRACKET, 0)
}

func (s *Type_matrixContext) Datatype() IDatatypeContext {
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

func (s *Type_matrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_matrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_matrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterType_matrix(s)
	}
}

func (s *Type_matrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitType_matrix(s)
	}
}

func (s *Type_matrixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitType_matrix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Type_matrix() (localctx IType_matrixContext) {
	localctx = NewType_matrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SwiftgrammParserRULE_type_matrix)
	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(605)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(606)

			var _x = p.Type_matrix()

			localctx.(*Type_matrixContext)._type_matrix = _x
		}
		{
			p.SetState(607)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		CountM++
		localctx.(*Type_matrixContext).td = localctx.(*Type_matrixContext).Get_type_matrix().GetTd()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(610)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(611)

			var _x = p.Datatype()

			localctx.(*Type_matrixContext)._datatype = _x
		}
		{
			p.SetState(612)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		CountM++
		localctx.(*Type_matrixContext).td = localctx.(*Type_matrixContext).Get_datatype().GetTd()

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

// IDefinition_matrixContext is an interface to support dynamic dispatch.
type IDefinition_matrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP returns the p attribute.
	GetP() [][]interface{}

	// SetP sets the p attribute.
	SetP([][]interface{})

	// Getter signatures
	Type_matrix() IType_matrixContext

	// IsDefinition_matrixContext differentiates from other interfaces.
	IsDefinition_matrixContext()
}

type Definition_matrixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	p      [][]interface{}
}

func NewEmptyDefinition_matrixContext() *Definition_matrixContext {
	var p = new(Definition_matrixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_definition_matrix
	return p
}

func InitEmptyDefinition_matrixContext(p *Definition_matrixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammParserRULE_definition_matrix
}

func (*Definition_matrixContext) IsDefinition_matrixContext() {}

func NewDefinition_matrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Definition_matrixContext {
	var p = new(Definition_matrixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammParserRULE_definition_matrix

	return p
}

func (s *Definition_matrixContext) GetParser() antlr.Parser { return s.parser }

func (s *Definition_matrixContext) GetP() [][]interface{} { return s.p }

func (s *Definition_matrixContext) SetP(v [][]interface{}) { s.p = v }

func (s *Definition_matrixContext) Type_matrix() IType_matrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_matrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_matrixContext)
}

func (s *Definition_matrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Definition_matrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Definition_matrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.EnterDefinition_matrix(s)
	}
}

func (s *Definition_matrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammListener); ok {
		listenerT.ExitDefinition_matrix(s)
	}
}

func (s *Definition_matrixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftgrammVisitor:
		return t.VisitDefinition_matrix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftgrammParser) Definition_matrix() (localctx IDefinition_matrixContext) {
	localctx = NewDefinition_matrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SwiftgrammParserRULE_definition_matrix)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(617)
		p.Type_matrix()
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

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_OPEN_BRACKET returns the _OPEN_BRACKET token.
	Get_OPEN_BRACKET() antlr.Token

	// Get_COUNT returns the _COUNT token.
	Get_COUNT() antlr.Token

	// Get_ISEMPTY returns the _ISEMPTY token.
	Get_ISEMPTY() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_FLOATT returns the _FLOATT token.
	Get_FLOATT() antlr.Token

	// Get_STRING_LITERAL returns the _STRING_LITERAL token.
	Get_STRING_LITERAL() antlr.Token

	// Get_CHARACTER_LITERAL returns the _CHARACTER_LITERAL token.
	Get_CHARACTER_LITERAL() antlr.Token

	// Get_TRUE returns the _TRUE token.
	Get_TRUE() antlr.Token

	// Get_FALSE returns the _FALSE token.
	Get_FALSE() antlr.Token

	// Get_NIL returns the _NIL token.
	Get_NIL() antlr.Token

	// SetOper sets the oper token.
	SetOper(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_OPEN_BRACKET sets the _OPEN_BRACKET token.
	Set_OPEN_BRACKET(antlr.Token)

	// Set_COUNT sets the _COUNT token.
	Set_COUNT(antlr.Token)

	// Set_ISEMPTY sets the _ISEMPTY token.
	Set_ISEMPTY(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_FLOATT sets the _FLOATT token.
	Set_FLOATT(antlr.Token)

	// Set_STRING_LITERAL sets the _STRING_LITERAL token.
	Set_STRING_LITERAL(antlr.Token)

	// Set_CHARACTER_LITERAL sets the _CHARACTER_LITERAL token.
	Set_CHARACTER_LITERAL(antlr.Token)

	// Set_TRUE sets the _TRUE token.
	Set_TRUE(antlr.Token)

	// Set_FALSE sets the _FALSE token.
	Set_FALSE(antlr.Token)

	// Set_NIL sets the _NIL token.
	Set_NIL(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_call_function_exp returns the _call_function_exp rule contexts.
	Get_call_function_exp() ICall_function_expContext

	// Get_native_function returns the _native_function rule contexts.
	Get_native_function() INative_functionContext

	// Get_array_exp returns the _array_exp rule contexts.
	Get_array_exp() IArray_expContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_call_function_exp sets the _call_function_exp rule contexts.
	Set_call_function_exp(ICall_function_expContext)

	// Set_native_function sets the _native_function rule contexts.
	Set_native_function(INative_functionContext)

	// Set_array_exp sets the _array_exp rule contexts.
	Set_array_exp(IArray_expContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// GetP returns the p attribute.
	GetP() abstract.Expression

	// SetP sets the p attribute.
	SetP(abstract.Expression)

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	NOT() antlr.TerminalNode
	SUBTRACTION() antlr.TerminalNode
	OPEN_PARENTHESIS() antlr.TerminalNode
	CLOSE_PARENTHESIS() antlr.TerminalNode
	Call_function_exp() ICall_function_expContext
	Native_function() INative_functionContext
	ID() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	DOT() antlr.TerminalNode
	COUNT() antlr.TerminalNode
	ISEMPTY() antlr.TerminalNode
	Array_exp() IArray_expContext
	NUMBER() antlr.TerminalNode
	FLOATT() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	CHARACTER_LITERAL() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NIL() antlr.TerminalNode
	MULTIPLICATION() antlr.TerminalNode
	DIVISION() antlr.TerminalNode
	SUMMATION() antlr.TerminalNode
	MOD() antlr.TerminalNode
	LESS_THAN() antlr.TerminalNode
	LESS_THAN_EQUAL() antlr.TerminalNode
	GREATER_THAN() antlr.TerminalNode
	GREATER_THAN_EQUAL() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	p                  abstract.Expression
	left               IExpressionContext
	oper               antlr.Token
	_expression        IExpressionContext
	_call_function_exp ICall_function_expContext
	_native_function   INative_functionContext
	_ID                antlr.Token
	_OPEN_BRACKET      antlr.Token
	_COUNT             antlr.Token
	_ISEMPTY           antlr.Token
	_array_exp         IArray_expContext
	_NUMBER            antlr.Token
	_FLOATT            antlr.Token
	_STRING_LITERAL    antlr.Token
	_CHARACTER_LITERAL antlr.Token
	_TRUE              antlr.Token
	_FALSE             antlr.Token
	_NIL               antlr.Token
	right              IExpressionContext
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

func (s *ExpressionContext) Get_ID() antlr.Token { return s._ID }

func (s *ExpressionContext) Get_OPEN_BRACKET() antlr.Token { return s._OPEN_BRACKET }

func (s *ExpressionContext) Get_COUNT() antlr.Token { return s._COUNT }

func (s *ExpressionContext) Get_ISEMPTY() antlr.Token { return s._ISEMPTY }

func (s *ExpressionContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExpressionContext) Get_FLOATT() antlr.Token { return s._FLOATT }

func (s *ExpressionContext) Get_STRING_LITERAL() antlr.Token { return s._STRING_LITERAL }

func (s *ExpressionContext) Get_CHARACTER_LITERAL() antlr.Token { return s._CHARACTER_LITERAL }

func (s *ExpressionContext) Get_TRUE() antlr.Token { return s._TRUE }

func (s *ExpressionContext) Get_FALSE() antlr.Token { return s._FALSE }

func (s *ExpressionContext) Get_NIL() antlr.Token { return s._NIL }

func (s *ExpressionContext) SetOper(v antlr.Token) { s.oper = v }

func (s *ExpressionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExpressionContext) Set_OPEN_BRACKET(v antlr.Token) { s._OPEN_BRACKET = v }

func (s *ExpressionContext) Set_COUNT(v antlr.Token) { s._COUNT = v }

func (s *ExpressionContext) Set_ISEMPTY(v antlr.Token) { s._ISEMPTY = v }

func (s *ExpressionContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExpressionContext) Set_FLOATT(v antlr.Token) { s._FLOATT = v }

func (s *ExpressionContext) Set_STRING_LITERAL(v antlr.Token) { s._STRING_LITERAL = v }

func (s *ExpressionContext) Set_CHARACTER_LITERAL(v antlr.Token) { s._CHARACTER_LITERAL = v }

func (s *ExpressionContext) Set_TRUE(v antlr.Token) { s._TRUE = v }

func (s *ExpressionContext) Set_FALSE(v antlr.Token) { s._FALSE = v }

func (s *ExpressionContext) Set_NIL(v antlr.Token) { s._NIL = v }

func (s *ExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *ExpressionContext) Get_expression() IExpressionContext { return s._expression }

func (s *ExpressionContext) Get_call_function_exp() ICall_function_expContext {
	return s._call_function_exp
}

func (s *ExpressionContext) Get_native_function() INative_functionContext { return s._native_function }

func (s *ExpressionContext) Get_array_exp() IArray_expContext { return s._array_exp }

func (s *ExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ExpressionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ExpressionContext) Set_call_function_exp(v ICall_function_expContext) {
	s._call_function_exp = v
}

func (s *ExpressionContext) Set_native_function(v INative_functionContext) { s._native_function = v }

func (s *ExpressionContext) Set_array_exp(v IArray_expContext) { s._array_exp = v }

func (s *ExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ExpressionContext) GetP() abstract.Expression { return s.p }

func (s *ExpressionContext) SetP(v abstract.Expression) { s.p = v }

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

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserNOT, 0)
}

func (s *ExpressionContext) SUBTRACTION() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSUBTRACTION, 0)
}

func (s *ExpressionContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *ExpressionContext) CLOSE_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_PARENTHESIS, 0)
}

func (s *ExpressionContext) Call_function_exp() ICall_function_expContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_function_expContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_function_expContext)
}

func (s *ExpressionContext) Native_function() INative_functionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INative_functionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INative_functionContext)
}

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
}

func (s *ExpressionContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_BRACKET, 0)
}

func (s *ExpressionContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCLOSE_BRACKET, 0)
}

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserDOT, 0)
}

func (s *ExpressionContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCOUNT, 0)
}

func (s *ExpressionContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserISEMPTY, 0)
}

func (s *ExpressionContext) Array_exp() IArray_expContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_expContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_expContext)
}

func (s *ExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserNUMBER, 0)
}

func (s *ExpressionContext) FLOATT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserFLOATT, 0)
}

func (s *ExpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSTRING_LITERAL, 0)
}

func (s *ExpressionContext) CHARACTER_LITERAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCHARACTER_LITERAL, 0)
}

func (s *ExpressionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserTRUE, 0)
}

func (s *ExpressionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserFALSE, 0)
}

func (s *ExpressionContext) NIL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserNIL, 0)
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

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserMOD, 0)
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

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOR, 0)
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
	_startState := 70
	p.EnterRecursionRule(localctx, 70, SwiftgrammParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(674)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(620)

			var _m = p.Match(SwiftgrammParserNOT)

			localctx.(*ExpressionContext).oper = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)

			var _x = p.expression(17)

			localctx.(*ExpressionContext)._expression = _x
		}

		localctx.(*ExpressionContext).p = expressions.NewLogicOperations(nil, localctx.(*ExpressionContext).Get_expression().GetP(), (func() string {
			if localctx.(*ExpressionContext).GetOper() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).GetOper().GetText()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).GetOper() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).GetOper().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).GetOper() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).GetOper().GetColumn()
			}
		}()))

	case 2:
		{
			p.SetState(624)

			var _m = p.Match(SwiftgrammParserSUBTRACTION)

			localctx.(*ExpressionContext).oper = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(625)

			var _x = p.expression(16)

			localctx.(*ExpressionContext)._expression = _x
		}

		localctx.(*ExpressionContext).p = expressions.NewNegative(localctx.(*ExpressionContext).Get_expression().GetP(), (func() int {
			if localctx.(*ExpressionContext).GetOper() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).GetOper().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).GetOper() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).GetOper().GetColumn()
			}
		}()))

	case 3:
		{
			p.SetState(628)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(629)

			var _x = p.expression(0)

			localctx.(*ExpressionContext)._expression = _x
		}
		{
			p.SetState(630)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expression().GetP()

	case 4:
		{
			p.SetState(633)

			var _x = p.Call_function_exp()

			localctx.(*ExpressionContext)._call_function_exp = _x
		}

		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_call_function_exp().GetP()

	case 5:
		{
			p.SetState(636)

			var _x = p.Native_function()

			localctx.(*ExpressionContext)._native_function = _x
		}

		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_native_function().GetP()

	case 6:
		{
			p.SetState(639)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ExpressionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(640)

			var _m = p.Match(SwiftgrammParserOPEN_BRACKET)

			localctx.(*ExpressionContext)._OPEN_BRACKET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(641)

			var _x = p.expression(0)

			localctx.(*ExpressionContext)._expression = _x
		}
		{
			p.SetState(642)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = expressions.NewGetPosVector((func() string {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetText()
			}
		}()), localctx.(*ExpressionContext).Get_expression().GetP(), (func() int {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetColumn()
			}
		}()))

	case 7:
		{
			p.SetState(645)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ExpressionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(646)
			p.Match(SwiftgrammParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(647)

			var _m = p.Match(SwiftgrammParserCOUNT)

			localctx.(*ExpressionContext)._COUNT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = expressions.NewVectorValues((func() string {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*ExpressionContext).Get_COUNT() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_COUNT().GetText()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetColumn()
			}
		}()))

	case 8:
		{
			p.SetState(649)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ExpressionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(650)
			p.Match(SwiftgrammParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(651)

			var _m = p.Match(SwiftgrammParserISEMPTY)

			localctx.(*ExpressionContext)._ISEMPTY = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = expressions.NewVectorValues((func() string {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*ExpressionContext).Get_ISEMPTY() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_ISEMPTY().GetText()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetColumn()
			}
		}()))

	case 9:
		{
			p.SetState(653)

			var _m = p.Match(SwiftgrammParserOPEN_BRACKET)

			localctx.(*ExpressionContext)._OPEN_BRACKET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(654)

			var _x = p.Array_exp()

			localctx.(*ExpressionContext)._array_exp = _x
		}
		{
			p.SetState(655)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = expressions.NewNative(localctx.(*ExpressionContext).Get_array_exp().GetP(), symbol.ARRAY, (func() int {
			if localctx.(*ExpressionContext).Get_OPEN_BRACKET() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_OPEN_BRACKET().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_OPEN_BRACKET() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_OPEN_BRACKET().GetColumn()
			}
		}()))

	case 10:
		{
			p.SetState(658)

			var _m = p.Match(SwiftgrammParserNUMBER)

			localctx.(*ExpressionContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value, err := strconv.Atoi((func() string {
			if localctx.(*ExpressionContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.INT, (func() int {
			if localctx.(*ExpressionContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_NUMBER().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_NUMBER().GetColumn()
			}
		}()))

	case 11:
		{
			p.SetState(660)

			var _m = p.Match(SwiftgrammParserFLOATT)

			localctx.(*ExpressionContext)._FLOATT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value, err := strconv.ParseFloat((func() string {
			if localctx.(*ExpressionContext).Get_FLOATT() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_FLOATT().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.FLOAT, (func() int {
			if localctx.(*ExpressionContext).Get_FLOATT() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_FLOATT().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_FLOATT() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_FLOATT().GetColumn()
			}
		}()))

	case 12:
		{
			p.SetState(662)

			var _m = p.Match(SwiftgrammParserSTRING_LITERAL)

			localctx.(*ExpressionContext)._STRING_LITERAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value := (func() string {
			if localctx.(*ExpressionContext).Get_STRING_LITERAL() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_STRING_LITERAL().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ExpressionContext).Get_STRING_LITERAL() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_STRING_LITERAL().GetText()
			}
		}()))-1]
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.STRING, (func() int {
			if localctx.(*ExpressionContext).Get_STRING_LITERAL() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_STRING_LITERAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_STRING_LITERAL() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_STRING_LITERAL().GetColumn()
			}
		}()))

	case 13:
		{
			p.SetState(664)

			var _m = p.Match(SwiftgrammParserCHARACTER_LITERAL)

			localctx.(*ExpressionContext)._CHARACTER_LITERAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value := (func() string {
			if localctx.(*ExpressionContext).Get_CHARACTER_LITERAL() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_CHARACTER_LITERAL().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ExpressionContext).Get_CHARACTER_LITERAL() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_CHARACTER_LITERAL().GetText()
			}
		}()))-1]
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.CHAR, (func() int {
			if localctx.(*ExpressionContext).Get_CHARACTER_LITERAL() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_CHARACTER_LITERAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_CHARACTER_LITERAL() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_CHARACTER_LITERAL().GetColumn()
			}
		}()))

	case 14:
		{
			p.SetState(666)

			var _m = p.Match(SwiftgrammParserTRUE)

			localctx.(*ExpressionContext)._TRUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value, err := strconv.ParseBool((func() string {
			if localctx.(*ExpressionContext).Get_TRUE() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_TRUE().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.BOOL, (func() int {
			if localctx.(*ExpressionContext).Get_TRUE() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_TRUE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_TRUE() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_TRUE().GetColumn()
			}
		}()))

	case 15:
		{
			p.SetState(668)

			var _m = p.Match(SwiftgrammParserFALSE)

			localctx.(*ExpressionContext)._FALSE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		value, err := strconv.ParseBool((func() string {
			if localctx.(*ExpressionContext).Get_FALSE() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_FALSE().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.BOOL, (func() int {
			if localctx.(*ExpressionContext).Get_FALSE() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_FALSE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_FALSE() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_FALSE().GetColumn()
			}
		}()))

	case 16:
		{
			p.SetState(670)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ExpressionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = expressions.NewIdentifier((func() string {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetText()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_ID().GetColumn()
			}
		}()))

	case 17:
		{
			p.SetState(672)

			var _m = p.Match(SwiftgrammParserNIL)

			localctx.(*ExpressionContext)._NIL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = expressions.NewNative(nil, symbol.NIL, (func() int {
			if localctx.(*ExpressionContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_NIL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpressionContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExpressionContext).Get_NIL().GetColumn()
			}
		}()))

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(713)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(711)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(676)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(677)

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
					p.SetState(678)

					var _x = p.expression(25)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewArithmeticOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetColumn()
					}
				}()))

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(681)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(682)

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
					p.SetState(683)

					var _x = p.expression(24)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewArithmeticOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetColumn()
					}
				}()))

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(686)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(687)

					var _m = p.Match(SwiftgrammParserMOD)

					localctx.(*ExpressionContext).oper = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(688)

					var _x = p.expression(23)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewArithmeticOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetColumn()
					}
				}()))

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(691)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(692)

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
					p.SetState(693)

					var _x = p.expression(22)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewRelationalOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetColumn()
					}
				}()))

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(696)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(697)

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
					p.SetState(698)

					var _x = p.expression(21)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewRelationalOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetColumn()
					}
				}()))

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(701)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(702)

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
					p.SetState(703)

					var _x = p.expression(20)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewRelationalOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetColumn()
					}
				}()))

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(706)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(707)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).oper = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftgrammParserOR || _la == SwiftgrammParserAND) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).oper = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(708)

					var _x = p.expression(19)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewLogicOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return 0
					} else {
						return localctx.(*ExpressionContext).GetOper().GetColumn()
					}
				}()))

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(715)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
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

	// GetTd returns the td attribute.
	GetTd() symbol.TypeData

	// SetTd sets the td attribute.
	SetTd(symbol.TypeData)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode

	// IsDatatypeContext differentiates from other interfaces.
	IsDatatypeContext()
}

type DatatypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	td     symbol.TypeData
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

func (s *DatatypeContext) GetTd() symbol.TypeData { return s.td }

func (s *DatatypeContext) SetTd(v symbol.TypeData) { s.td = v }

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

func (s *DatatypeContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCHARACTER, 0)
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
	p.EnterRule(localctx, 72, SwiftgrammParserRULE_datatype)
	p.SetState(726)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(716)
			p.Match(SwiftgrammParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*DatatypeContext).td = symbol.INT

	case SwiftgrammParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(718)
			p.Match(SwiftgrammParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*DatatypeContext).td = symbol.FLOAT

	case SwiftgrammParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(720)
			p.Match(SwiftgrammParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*DatatypeContext).td = symbol.STRING

	case SwiftgrammParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(722)
			p.Match(SwiftgrammParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*DatatypeContext).td = symbol.BOOL

	case SwiftgrammParserCHARACTER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(724)
			p.Match(SwiftgrammParserCHARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*DatatypeContext).td = symbol.CHAR

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

func (p *SwiftgrammParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 35:
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
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 18)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
