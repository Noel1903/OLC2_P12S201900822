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
		"'inout'", "", "", "", "", "", "'+='", "'-='", "'...'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'?'", "'||'", "'&&'", "'!'", "'=='", "'!='", "'<'",
		"'<='", "'>'", "'>='", "'='", "'.'", "','", "':'", "';'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'->'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE",
		"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK",
		"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC",
		"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "FLOATT", "ID", "CHARACTER_LITERAL",
		"STRING_LITERAL", "INCREMENT", "DECREMENT", "RANGE", "SUMMATION", "SUBTRACTION",
		"MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT",
		"EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN",
		"GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON",
		"OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", "OPEN_BRACKET",
		"CLOSE_BRACKET", "ARROW", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "sentence", "increment_bl", "decrement_bl", "break_bl",
		"return_bl", "continue_bl", "declare_let", "declare_var", "print_bl",
		"if_bl", "else_if", "while_bl", "for_bl", "guard_bl", "vector_bl", "array_exp",
		"function_bl", "params", "call_function", "list_exp", "call_function_exp",
		"expression", "datatype",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 453, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 4, 1, 56, 8, 1, 11, 1, 12, 1, 57, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 107, 8, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 128, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 147, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 3, 9, 170, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 3, 11, 204, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 235, 8, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 3, 16, 286, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 296, 8, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 330, 8, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 3, 19, 344, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 356, 8, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 366, 8, 21, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 399, 8, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 5, 23, 436, 8, 23, 10, 23, 12, 23, 439, 9, 23, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24,
		451, 8, 24, 1, 24, 0, 1, 46, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 6, 1, 0, 39,
		40, 1, 0, 37, 38, 1, 0, 48, 49, 1, 0, 50, 51, 1, 0, 46, 47, 1, 0, 43, 44,
		478, 0, 50, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0, 4, 106, 1, 0, 0, 0, 6, 108,
		1, 0, 0, 0, 8, 113, 1, 0, 0, 0, 10, 118, 1, 0, 0, 0, 12, 127, 1, 0, 0,
		0, 14, 129, 1, 0, 0, 0, 16, 146, 1, 0, 0, 0, 18, 169, 1, 0, 0, 0, 20, 171,
		1, 0, 0, 0, 22, 203, 1, 0, 0, 0, 24, 234, 1, 0, 0, 0, 26, 236, 1, 0, 0,
		0, 28, 243, 1, 0, 0, 0, 30, 254, 1, 0, 0, 0, 32, 285, 1, 0, 0, 0, 34, 295,
		1, 0, 0, 0, 36, 329, 1, 0, 0, 0, 38, 343, 1, 0, 0, 0, 40, 355, 1, 0, 0,
		0, 42, 365, 1, 0, 0, 0, 44, 367, 1, 0, 0, 0, 46, 398, 1, 0, 0, 0, 48, 450,
		1, 0, 0, 0, 50, 51, 3, 2, 1, 0, 51, 52, 5, 0, 0, 1, 52, 53, 6, 0, -1, 0,
		53, 1, 1, 0, 0, 0, 54, 56, 3, 4, 2, 0, 55, 54, 1, 0, 0, 0, 56, 57, 1, 0,
		0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60,
		6, 1, -1, 0, 60, 3, 1, 0, 0, 0, 61, 62, 3, 18, 9, 0, 62, 63, 6, 2, -1,
		0, 63, 107, 1, 0, 0, 0, 64, 65, 3, 16, 8, 0, 65, 66, 6, 2, -1, 0, 66, 107,
		1, 0, 0, 0, 67, 68, 3, 20, 10, 0, 68, 69, 6, 2, -1, 0, 69, 107, 1, 0, 0,
		0, 70, 71, 3, 22, 11, 0, 71, 72, 6, 2, -1, 0, 72, 107, 1, 0, 0, 0, 73,
		74, 3, 6, 3, 0, 74, 75, 6, 2, -1, 0, 75, 107, 1, 0, 0, 0, 76, 77, 3, 8,
		4, 0, 77, 78, 6, 2, -1, 0, 78, 107, 1, 0, 0, 0, 79, 80, 3, 26, 13, 0, 80,
		81, 6, 2, -1, 0, 81, 107, 1, 0, 0, 0, 82, 83, 3, 28, 14, 0, 83, 84, 6,
		2, -1, 0, 84, 107, 1, 0, 0, 0, 85, 86, 3, 30, 15, 0, 86, 87, 6, 2, -1,
		0, 87, 107, 1, 0, 0, 0, 88, 89, 3, 10, 5, 0, 89, 90, 6, 2, -1, 0, 90, 107,
		1, 0, 0, 0, 91, 92, 3, 12, 6, 0, 92, 93, 6, 2, -1, 0, 93, 107, 1, 0, 0,
		0, 94, 95, 3, 14, 7, 0, 95, 96, 6, 2, -1, 0, 96, 107, 1, 0, 0, 0, 97, 98,
		3, 32, 16, 0, 98, 99, 6, 2, -1, 0, 99, 107, 1, 0, 0, 0, 100, 101, 3, 36,
		18, 0, 101, 102, 6, 2, -1, 0, 102, 107, 1, 0, 0, 0, 103, 104, 3, 40, 20,
		0, 104, 105, 6, 2, -1, 0, 105, 107, 1, 0, 0, 0, 106, 61, 1, 0, 0, 0, 106,
		64, 1, 0, 0, 0, 106, 67, 1, 0, 0, 0, 106, 70, 1, 0, 0, 0, 106, 73, 1, 0,
		0, 0, 106, 76, 1, 0, 0, 0, 106, 79, 1, 0, 0, 0, 106, 82, 1, 0, 0, 0, 106,
		85, 1, 0, 0, 0, 106, 88, 1, 0, 0, 0, 106, 91, 1, 0, 0, 0, 106, 94, 1, 0,
		0, 0, 106, 97, 1, 0, 0, 0, 106, 100, 1, 0, 0, 0, 106, 103, 1, 0, 0, 0,
		107, 5, 1, 0, 0, 0, 108, 109, 5, 31, 0, 0, 109, 110, 5, 34, 0, 0, 110,
		111, 3, 46, 23, 0, 111, 112, 6, 3, -1, 0, 112, 7, 1, 0, 0, 0, 113, 114,
		5, 31, 0, 0, 114, 115, 5, 35, 0, 0, 115, 116, 3, 46, 23, 0, 116, 117, 6,
		4, -1, 0, 117, 9, 1, 0, 0, 0, 118, 119, 5, 16, 0, 0, 119, 120, 6, 5, -1,
		0, 120, 11, 1, 0, 0, 0, 121, 122, 5, 23, 0, 0, 122, 123, 3, 46, 23, 0,
		123, 124, 6, 6, -1, 0, 124, 128, 1, 0, 0, 0, 125, 126, 5, 23, 0, 0, 126,
		128, 6, 6, -1, 0, 127, 121, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 13,
		1, 0, 0, 0, 129, 130, 5, 22, 0, 0, 130, 131, 6, 7, -1, 0, 131, 15, 1, 0,
		0, 0, 132, 133, 5, 10, 0, 0, 133, 134, 5, 31, 0, 0, 134, 135, 5, 55, 0,
		0, 135, 136, 3, 48, 24, 0, 136, 137, 5, 52, 0, 0, 137, 138, 3, 46, 23,
		0, 138, 139, 6, 8, -1, 0, 139, 147, 1, 0, 0, 0, 140, 141, 5, 10, 0, 0,
		141, 142, 5, 31, 0, 0, 142, 143, 5, 52, 0, 0, 143, 144, 3, 46, 23, 0, 144,
		145, 6, 8, -1, 0, 145, 147, 1, 0, 0, 0, 146, 132, 1, 0, 0, 0, 146, 140,
		1, 0, 0, 0, 147, 17, 1, 0, 0, 0, 148, 149, 5, 9, 0, 0, 149, 150, 5, 31,
		0, 0, 150, 151, 5, 55, 0, 0, 151, 152, 3, 48, 24, 0, 152, 153, 5, 52, 0,
		0, 153, 154, 3, 46, 23, 0, 154, 155, 6, 9, -1, 0, 155, 170, 1, 0, 0, 0,
		156, 157, 5, 9, 0, 0, 157, 158, 5, 31, 0, 0, 158, 159, 5, 52, 0, 0, 159,
		160, 3, 46, 23, 0, 160, 161, 6, 9, -1, 0, 161, 170, 1, 0, 0, 0, 162, 163,
		5, 9, 0, 0, 163, 164, 5, 31, 0, 0, 164, 165, 5, 55, 0, 0, 165, 166, 3,
		48, 24, 0, 166, 167, 5, 42, 0, 0, 167, 168, 6, 9, -1, 0, 168, 170, 1, 0,
		0, 0, 169, 148, 1, 0, 0, 0, 169, 156, 1, 0, 0, 0, 169, 162, 1, 0, 0, 0,
		170, 19, 1, 0, 0, 0, 171, 172, 5, 11, 0, 0, 172, 173, 5, 57, 0, 0, 173,
		174, 3, 46, 23, 0, 174, 175, 5, 58, 0, 0, 175, 176, 6, 10, -1, 0, 176,
		21, 1, 0, 0, 0, 177, 178, 5, 12, 0, 0, 178, 179, 3, 46, 23, 0, 179, 180,
		5, 59, 0, 0, 180, 181, 3, 2, 1, 0, 181, 182, 5, 60, 0, 0, 182, 183, 6,
		11, -1, 0, 183, 204, 1, 0, 0, 0, 184, 185, 5, 12, 0, 0, 185, 186, 3, 46,
		23, 0, 186, 187, 5, 59, 0, 0, 187, 188, 3, 2, 1, 0, 188, 189, 5, 60, 0,
		0, 189, 190, 5, 13, 0, 0, 190, 191, 5, 59, 0, 0, 191, 192, 3, 2, 1, 0,
		192, 193, 5, 60, 0, 0, 193, 194, 6, 11, -1, 0, 194, 204, 1, 0, 0, 0, 195,
		196, 5, 12, 0, 0, 196, 197, 3, 46, 23, 0, 197, 198, 5, 59, 0, 0, 198, 199,
		3, 2, 1, 0, 199, 200, 5, 60, 0, 0, 200, 201, 3, 24, 12, 0, 201, 202, 6,
		11, -1, 0, 202, 204, 1, 0, 0, 0, 203, 177, 1, 0, 0, 0, 203, 184, 1, 0,
		0, 0, 203, 195, 1, 0, 0, 0, 204, 23, 1, 0, 0, 0, 205, 206, 5, 13, 0, 0,
		206, 207, 5, 12, 0, 0, 207, 208, 3, 46, 23, 0, 208, 209, 5, 59, 0, 0, 209,
		210, 3, 2, 1, 0, 210, 211, 5, 60, 0, 0, 211, 212, 6, 12, -1, 0, 212, 235,
		1, 0, 0, 0, 213, 214, 5, 13, 0, 0, 214, 215, 5, 12, 0, 0, 215, 216, 3,
		46, 23, 0, 216, 217, 5, 59, 0, 0, 217, 218, 3, 2, 1, 0, 218, 219, 5, 60,
		0, 0, 219, 220, 5, 13, 0, 0, 220, 221, 5, 59, 0, 0, 221, 222, 3, 2, 1,
		0, 222, 223, 5, 60, 0, 0, 223, 224, 6, 12, -1, 0, 224, 235, 1, 0, 0, 0,
		225, 226, 5, 13, 0, 0, 226, 227, 5, 12, 0, 0, 227, 228, 3, 46, 23, 0, 228,
		229, 5, 59, 0, 0, 229, 230, 3, 2, 1, 0, 230, 231, 5, 60, 0, 0, 231, 232,
		3, 24, 12, 0, 232, 233, 6, 12, -1, 0, 233, 235, 1, 0, 0, 0, 234, 205, 1,
		0, 0, 0, 234, 213, 1, 0, 0, 0, 234, 225, 1, 0, 0, 0, 235, 25, 1, 0, 0,
		0, 236, 237, 5, 18, 0, 0, 237, 238, 3, 46, 23, 0, 238, 239, 5, 59, 0, 0,
		239, 240, 3, 2, 1, 0, 240, 241, 5, 60, 0, 0, 241, 242, 6, 13, -1, 0, 242,
		27, 1, 0, 0, 0, 243, 244, 5, 19, 0, 0, 244, 245, 5, 31, 0, 0, 245, 246,
		5, 20, 0, 0, 246, 247, 3, 46, 23, 0, 247, 248, 5, 36, 0, 0, 248, 249, 3,
		46, 23, 0, 249, 250, 5, 59, 0, 0, 250, 251, 3, 2, 1, 0, 251, 252, 5, 60,
		0, 0, 252, 253, 6, 14, -1, 0, 253, 29, 1, 0, 0, 0, 254, 255, 5, 21, 0,
		0, 255, 256, 3, 46, 23, 0, 256, 257, 5, 13, 0, 0, 257, 258, 5, 59, 0, 0,
		258, 259, 3, 2, 1, 0, 259, 260, 5, 60, 0, 0, 260, 261, 6, 15, -1, 0, 261,
		31, 1, 0, 0, 0, 262, 263, 5, 9, 0, 0, 263, 264, 5, 31, 0, 0, 264, 265,
		5, 55, 0, 0, 265, 266, 5, 61, 0, 0, 266, 267, 3, 48, 24, 0, 267, 268, 5,
		62, 0, 0, 268, 269, 5, 52, 0, 0, 269, 270, 5, 61, 0, 0, 270, 271, 3, 34,
		17, 0, 271, 272, 5, 62, 0, 0, 272, 273, 6, 16, -1, 0, 273, 286, 1, 0, 0,
		0, 274, 275, 5, 9, 0, 0, 275, 276, 5, 31, 0, 0, 276, 277, 5, 55, 0, 0,
		277, 278, 5, 61, 0, 0, 278, 279, 3, 48, 24, 0, 279, 280, 5, 62, 0, 0, 280,
		281, 5, 52, 0, 0, 281, 282, 5, 61, 0, 0, 282, 283, 5, 62, 0, 0, 283, 284,
		6, 16, -1, 0, 284, 286, 1, 0, 0, 0, 285, 262, 1, 0, 0, 0, 285, 274, 1,
		0, 0, 0, 286, 33, 1, 0, 0, 0, 287, 288, 3, 46, 23, 0, 288, 289, 5, 54,
		0, 0, 289, 290, 3, 34, 17, 0, 290, 291, 6, 17, -1, 0, 291, 296, 1, 0, 0,
		0, 292, 293, 3, 46, 23, 0, 293, 294, 6, 17, -1, 0, 294, 296, 1, 0, 0, 0,
		295, 287, 1, 0, 0, 0, 295, 292, 1, 0, 0, 0, 296, 35, 1, 0, 0, 0, 297, 298,
		5, 24, 0, 0, 298, 299, 5, 31, 0, 0, 299, 300, 5, 57, 0, 0, 300, 301, 5,
		58, 0, 0, 301, 302, 5, 63, 0, 0, 302, 303, 3, 48, 24, 0, 303, 304, 5, 59,
		0, 0, 304, 305, 3, 2, 1, 0, 305, 306, 5, 60, 0, 0, 306, 307, 6, 18, -1,
		0, 307, 330, 1, 0, 0, 0, 308, 309, 5, 24, 0, 0, 309, 310, 5, 31, 0, 0,
		310, 311, 5, 57, 0, 0, 311, 312, 5, 58, 0, 0, 312, 313, 5, 59, 0, 0, 313,
		314, 3, 2, 1, 0, 314, 315, 5, 60, 0, 0, 315, 316, 6, 18, -1, 0, 316, 330,
		1, 0, 0, 0, 317, 318, 5, 24, 0, 0, 318, 319, 5, 31, 0, 0, 319, 320, 5,
		57, 0, 0, 320, 321, 3, 38, 19, 0, 321, 322, 5, 58, 0, 0, 322, 323, 5, 63,
		0, 0, 323, 324, 3, 48, 24, 0, 324, 325, 5, 59, 0, 0, 325, 326, 3, 2, 1,
		0, 326, 327, 5, 60, 0, 0, 327, 328, 6, 18, -1, 0, 328, 330, 1, 0, 0, 0,
		329, 297, 1, 0, 0, 0, 329, 308, 1, 0, 0, 0, 329, 317, 1, 0, 0, 0, 330,
		37, 1, 0, 0, 0, 331, 332, 5, 31, 0, 0, 332, 333, 5, 55, 0, 0, 333, 334,
		3, 48, 24, 0, 334, 335, 5, 54, 0, 0, 335, 336, 3, 38, 19, 0, 336, 337,
		6, 19, -1, 0, 337, 344, 1, 0, 0, 0, 338, 339, 5, 31, 0, 0, 339, 340, 5,
		55, 0, 0, 340, 341, 3, 48, 24, 0, 341, 342, 6, 19, -1, 0, 342, 344, 1,
		0, 0, 0, 343, 331, 1, 0, 0, 0, 343, 338, 1, 0, 0, 0, 344, 39, 1, 0, 0,
		0, 345, 346, 5, 31, 0, 0, 346, 347, 5, 57, 0, 0, 347, 348, 5, 58, 0, 0,
		348, 356, 6, 20, -1, 0, 349, 350, 5, 31, 0, 0, 350, 351, 5, 57, 0, 0, 351,
		352, 3, 42, 21, 0, 352, 353, 5, 58, 0, 0, 353, 354, 6, 20, -1, 0, 354,
		356, 1, 0, 0, 0, 355, 345, 1, 0, 0, 0, 355, 349, 1, 0, 0, 0, 356, 41, 1,
		0, 0, 0, 357, 358, 3, 46, 23, 0, 358, 359, 5, 54, 0, 0, 359, 360, 3, 42,
		21, 0, 360, 361, 6, 21, -1, 0, 361, 366, 1, 0, 0, 0, 362, 363, 3, 46, 23,
		0, 363, 364, 6, 21, -1, 0, 364, 366, 1, 0, 0, 0, 365, 357, 1, 0, 0, 0,
		365, 362, 1, 0, 0, 0, 366, 43, 1, 0, 0, 0, 367, 368, 3, 40, 20, 0, 368,
		369, 6, 22, -1, 0, 369, 45, 1, 0, 0, 0, 370, 371, 6, 23, -1, 0, 371, 372,
		5, 45, 0, 0, 372, 373, 3, 46, 23, 11, 373, 374, 6, 23, -1, 0, 374, 399,
		1, 0, 0, 0, 375, 376, 5, 57, 0, 0, 376, 377, 3, 46, 23, 0, 377, 378, 5,
		58, 0, 0, 378, 399, 1, 0, 0, 0, 379, 380, 3, 44, 22, 0, 380, 381, 6, 23,
		-1, 0, 381, 399, 1, 0, 0, 0, 382, 383, 5, 29, 0, 0, 383, 399, 6, 23, -1,
		0, 384, 385, 5, 30, 0, 0, 385, 399, 6, 23, -1, 0, 386, 387, 5, 33, 0, 0,
		387, 399, 6, 23, -1, 0, 388, 389, 5, 32, 0, 0, 389, 399, 6, 23, -1, 0,
		390, 391, 5, 6, 0, 0, 391, 399, 6, 23, -1, 0, 392, 393, 5, 7, 0, 0, 393,
		399, 6, 23, -1, 0, 394, 395, 5, 31, 0, 0, 395, 399, 6, 23, -1, 0, 396,
		397, 5, 8, 0, 0, 397, 399, 6, 23, -1, 0, 398, 370, 1, 0, 0, 0, 398, 375,
		1, 0, 0, 0, 398, 379, 1, 0, 0, 0, 398, 382, 1, 0, 0, 0, 398, 384, 1, 0,
		0, 0, 398, 386, 1, 0, 0, 0, 398, 388, 1, 0, 0, 0, 398, 390, 1, 0, 0, 0,
		398, 392, 1, 0, 0, 0, 398, 394, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399,
		437, 1, 0, 0, 0, 400, 401, 10, 18, 0, 0, 401, 402, 7, 0, 0, 0, 402, 403,
		3, 46, 23, 19, 403, 404, 6, 23, -1, 0, 404, 436, 1, 0, 0, 0, 405, 406,
		10, 17, 0, 0, 406, 407, 7, 1, 0, 0, 407, 408, 3, 46, 23, 18, 408, 409,
		6, 23, -1, 0, 409, 436, 1, 0, 0, 0, 410, 411, 10, 16, 0, 0, 411, 412, 5,
		41, 0, 0, 412, 413, 3, 46, 23, 17, 413, 414, 6, 23, -1, 0, 414, 436, 1,
		0, 0, 0, 415, 416, 10, 15, 0, 0, 416, 417, 7, 2, 0, 0, 417, 418, 3, 46,
		23, 16, 418, 419, 6, 23, -1, 0, 419, 436, 1, 0, 0, 0, 420, 421, 10, 14,
		0, 0, 421, 422, 7, 3, 0, 0, 422, 423, 3, 46, 23, 15, 423, 424, 6, 23, -1,
		0, 424, 436, 1, 0, 0, 0, 425, 426, 10, 13, 0, 0, 426, 427, 7, 4, 0, 0,
		427, 428, 3, 46, 23, 14, 428, 429, 6, 23, -1, 0, 429, 436, 1, 0, 0, 0,
		430, 431, 10, 12, 0, 0, 431, 432, 7, 5, 0, 0, 432, 433, 3, 46, 23, 13,
		433, 434, 6, 23, -1, 0, 434, 436, 1, 0, 0, 0, 435, 400, 1, 0, 0, 0, 435,
		405, 1, 0, 0, 0, 435, 410, 1, 0, 0, 0, 435, 415, 1, 0, 0, 0, 435, 420,
		1, 0, 0, 0, 435, 425, 1, 0, 0, 0, 435, 430, 1, 0, 0, 0, 436, 439, 1, 0,
		0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 47, 1, 0, 0, 0,
		439, 437, 1, 0, 0, 0, 440, 441, 5, 1, 0, 0, 441, 451, 6, 24, -1, 0, 442,
		443, 5, 2, 0, 0, 443, 451, 6, 24, -1, 0, 444, 445, 5, 3, 0, 0, 445, 451,
		6, 24, -1, 0, 446, 447, 5, 4, 0, 0, 447, 451, 6, 24, -1, 0, 448, 449, 5,
		5, 0, 0, 449, 451, 6, 24, -1, 0, 450, 440, 1, 0, 0, 0, 450, 442, 1, 0,
		0, 0, 450, 444, 1, 0, 0, 0, 450, 446, 1, 0, 0, 0, 450, 448, 1, 0, 0, 0,
		451, 49, 1, 0, 0, 0, 17, 57, 106, 127, 146, 169, 203, 234, 285, 295, 329,
		343, 355, 365, 398, 435, 437, 450,
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
	SwiftgrammParserFLOATT             = 30
	SwiftgrammParserID                 = 31
	SwiftgrammParserCHARACTER_LITERAL  = 32
	SwiftgrammParserSTRING_LITERAL     = 33
	SwiftgrammParserINCREMENT          = 34
	SwiftgrammParserDECREMENT          = 35
	SwiftgrammParserRANGE              = 36
	SwiftgrammParserSUMMATION          = 37
	SwiftgrammParserSUBTRACTION        = 38
	SwiftgrammParserMULTIPLICATION     = 39
	SwiftgrammParserDIVISION           = 40
	SwiftgrammParserMOD                = 41
	SwiftgrammParserQUESTION_MARK      = 42
	SwiftgrammParserOR                 = 43
	SwiftgrammParserAND                = 44
	SwiftgrammParserNOT                = 45
	SwiftgrammParserEQUAL              = 46
	SwiftgrammParserNOT_EQUAL          = 47
	SwiftgrammParserLESS_THAN          = 48
	SwiftgrammParserLESS_THAN_EQUAL    = 49
	SwiftgrammParserGREATER_THAN       = 50
	SwiftgrammParserGREATER_THAN_EQUAL = 51
	SwiftgrammParserASSIGN             = 52
	SwiftgrammParserDOT                = 53
	SwiftgrammParserCOMMA              = 54
	SwiftgrammParserCOLON              = 55
	SwiftgrammParserSEMICOLON          = 56
	SwiftgrammParserOPEN_PARENTHESIS   = 57
	SwiftgrammParserCLOSE_PARENTHESIS  = 58
	SwiftgrammParserOPEN_kEY           = 59
	SwiftgrammParserCLOSE_kEY          = 60
	SwiftgrammParserOPEN_BRACKET       = 61
	SwiftgrammParserCLOSE_BRACKET      = 62
	SwiftgrammParserARROW              = 63
	SwiftgrammParserWHITESPACE         = 64
	SwiftgrammParserCOMMENT            = 65
	SwiftgrammParserLINE_COMMENT       = 66
)

// SwiftgrammParser rules.
const (
	SwiftgrammParserRULE_s                 = 0
	SwiftgrammParserRULE_block             = 1
	SwiftgrammParserRULE_sentence          = 2
	SwiftgrammParserRULE_increment_bl      = 3
	SwiftgrammParserRULE_decrement_bl      = 4
	SwiftgrammParserRULE_break_bl          = 5
	SwiftgrammParserRULE_return_bl         = 6
	SwiftgrammParserRULE_continue_bl       = 7
	SwiftgrammParserRULE_declare_let       = 8
	SwiftgrammParserRULE_declare_var       = 9
	SwiftgrammParserRULE_print_bl          = 10
	SwiftgrammParserRULE_if_bl             = 11
	SwiftgrammParserRULE_else_if           = 12
	SwiftgrammParserRULE_while_bl          = 13
	SwiftgrammParserRULE_for_bl            = 14
	SwiftgrammParserRULE_guard_bl          = 15
	SwiftgrammParserRULE_vector_bl         = 16
	SwiftgrammParserRULE_array_exp         = 17
	SwiftgrammParserRULE_function_bl       = 18
	SwiftgrammParserRULE_params            = 19
	SwiftgrammParserRULE_call_function     = 20
	SwiftgrammParserRULE_list_exp          = 21
	SwiftgrammParserRULE_call_function_exp = 22
	SwiftgrammParserRULE_expression        = 23
	SwiftgrammParserRULE_datatype          = 24
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
		p.SetState(50)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(51)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2179800576) != 0) {
		{
			p.SetState(54)

			var _x = p.Sentence()

			localctx.(*BlockContext)._sentence = _x
		}
		localctx.(*BlockContext).instr = append(localctx.(*BlockContext).instr, localctx.(*BlockContext)._sentence)

		p.SetState(57)
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

	// Set_declare_var sets the _declare_var rule contexts.
	Set_declare_var(IDeclare_varContext)

	// Set_declare_let sets the _declare_let rule contexts.
	Set_declare_let(IDeclare_letContext)

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

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	Declare_var() IDeclare_varContext
	Declare_let() IDeclare_letContext
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

	// IsSentenceContext differentiates from other interfaces.
	IsSentenceContext()
}

type SentenceContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          abstract.Instruction
	_declare_var   IDeclare_varContext
	_declare_let   IDeclare_letContext
	_print_bl      IPrint_blContext
	_if_bl         IIf_blContext
	_increment_bl  IIncrement_blContext
	_decrement_bl  IDecrement_blContext
	_while_bl      IWhile_blContext
	_for_bl        IFor_blContext
	_guard_bl      IGuard_blContext
	_break_bl      IBreak_blContext
	_return_bl     IReturn_blContext
	_continue_bl   IContinue_blContext
	_vector_bl     IVector_blContext
	_function_bl   IFunction_blContext
	_call_function ICall_functionContext
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

func (s *SentenceContext) Set_declare_var(v IDeclare_varContext) { s._declare_var = v }

func (s *SentenceContext) Set_declare_let(v IDeclare_letContext) { s._declare_let = v }

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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)

			var _x = p.Declare_var()

			localctx.(*SentenceContext)._declare_var = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_declare_var().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)

			var _x = p.Declare_let()

			localctx.(*SentenceContext)._declare_let = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_declare_let().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)

			var _x = p.Print_bl()

			localctx.(*SentenceContext)._print_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_print_bl().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)

			var _x = p.If_bl()

			localctx.(*SentenceContext)._if_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_if_bl().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(73)

			var _x = p.Increment_bl()

			localctx.(*SentenceContext)._increment_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_increment_bl().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(76)

			var _x = p.Decrement_bl()

			localctx.(*SentenceContext)._decrement_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_decrement_bl().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(79)

			var _x = p.While_bl()

			localctx.(*SentenceContext)._while_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_while_bl().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(82)

			var _x = p.For_bl()

			localctx.(*SentenceContext)._for_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_for_bl().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(85)

			var _x = p.Guard_bl()

			localctx.(*SentenceContext)._guard_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_guard_bl().GetInstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(88)

			var _x = p.Break_bl()

			localctx.(*SentenceContext)._break_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_break_bl().GetInstr()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(91)

			var _x = p.Return_bl()

			localctx.(*SentenceContext)._return_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_return_bl().GetInstr()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(94)

			var _x = p.Continue_bl()

			localctx.(*SentenceContext)._continue_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_continue_bl().GetInstr()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(97)

			var _x = p.Vector_bl()

			localctx.(*SentenceContext)._vector_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_vector_bl().GetInstr()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(100)

			var _x = p.Function_bl()

			localctx.(*SentenceContext)._function_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_function_bl().GetInstr()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(103)

			var _x = p.Call_function()

			localctx.(*SentenceContext)._call_function = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_call_function().GetInstr()

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
	p.EnterRule(localctx, 6, SwiftgrammParserRULE_increment_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Increment_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)

		var _m = p.Match(SwiftgrammParserINCREMENT)

		localctx.(*Increment_blContext)._INCREMENT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)

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
	}()), localctx.(*Increment_blContext).Get_expression().GetP())

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
	p.EnterRule(localctx, 8, SwiftgrammParserRULE_decrement_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Decrement_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)

		var _m = p.Match(SwiftgrammParserDECREMENT)

		localctx.(*Decrement_blContext)._DECREMENT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)

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
	}()), localctx.(*Decrement_blContext).Get_expression().GetP())

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
	p.EnterRule(localctx, 10, SwiftgrammParserRULE_break_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
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
	p.EnterRule(localctx, 12, SwiftgrammParserRULE_return_bl)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(SwiftgrammParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)

			var _x = p.expression(0)

			localctx.(*Return_blContext)._expression = _x
		}

		localctx.(*Return_blContext).instr = instructions.NewReturn(localctx.(*Return_blContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(SwiftgrammParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Return_blContext).instr = instructions.NewReturn(expressions.NewNative(nil, symbol.NIL))

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
	p.EnterRule(localctx, 14, SwiftgrammParserRULE_continue_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
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
	p.EnterRule(localctx, 16, SwiftgrammParserRULE_declare_let)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_letContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

			var _x = p.Datatype()

			localctx.(*Declare_letContext)._datatype = _x
		}
		{
			p.SetState(136)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)

			var _x = p.expression(0)

			localctx.(*Declare_letContext)._expression = _x
		}

		localctx.(*Declare_letContext).instr = instructions.NewLet((func() string {
			if localctx.(*Declare_letContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_letContext).Get_ID().GetText()
			}
		}()), localctx.(*Declare_letContext).Get_datatype().GetTd(), localctx.(*Declare_letContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_letContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)

			var _x = p.expression(0)

			localctx.(*Declare_letContext)._expression = _x
		}

		localctx.(*Declare_letContext).instr = instructions.NewLet((func() string {
			if localctx.(*Declare_letContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_letContext).Get_ID().GetText()
			}
		}()), symbol.UNDEFINED, localctx.(*Declare_letContext).Get_expression().GetP())

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

// IDeclare_varContext is an interface to support dynamic dispatch.
type IDeclare_varContext interface {
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
	parser      antlr.Parser
	instr       abstract.Instruction
	_ID         antlr.Token
	_datatype   IDatatypeContext
	_expression IExpressionContext
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

func (s *Declare_varContext) Set_ID(v antlr.Token) { s._ID = v }

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
	p.EnterRule(localctx, 18, SwiftgrammParserRULE_declare_var)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)

			var _x = p.Datatype()

			localctx.(*Declare_varContext)._datatype = _x
		}
		{
			p.SetState(152)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)

			var _x = p.expression(0)

			localctx.(*Declare_varContext)._expression = _x
		}

		localctx.(*Declare_varContext).instr = instructions.NewDeclareWithValue((func() string {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetText()
			}
		}()), localctx.(*Declare_varContext).Get_datatype().GetTd(), localctx.(*Declare_varContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)

			var _x = p.expression(0)

			localctx.(*Declare_varContext)._expression = _x
		}

		localctx.(*Declare_varContext).instr = instructions.NewDeclareWithValue((func() string {
			if localctx.(*Declare_varContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declare_varContext).Get_ID().GetText()
			}
		}()), symbol.UNDEFINED, localctx.(*Declare_varContext).Get_expression().GetP())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)

			var _x = p.Datatype()

			localctx.(*Declare_varContext)._datatype = _x
		}
		{
			p.SetState(166)
			p.Match(SwiftgrammParserQUESTION_MARK)
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
		}()), localctx.(*Declare_varContext).Get_datatype().GetTd(), expressions.NewNative(nil, symbol.NIL))

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

// IPrint_blContext is an interface to support dynamic dispatch.
type IPrint_blContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	OPEN_PARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	CLOSE_PARENTHESIS() antlr.TerminalNode

	// IsPrint_blContext differentiates from other interfaces.
	IsPrint_blContext()
}

type Print_blContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       abstract.Instruction
	_expression IExpressionContext
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

func (s *Print_blContext) Get_expression() IExpressionContext { return s._expression }

func (s *Print_blContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Print_blContext) GetInstr() abstract.Instruction { return s.instr }

func (s *Print_blContext) SetInstr(v abstract.Instruction) { s.instr = v }

func (s *Print_blContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserPRINT, 0)
}

func (s *Print_blContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserOPEN_PARENTHESIS, 0)
}

func (s *Print_blContext) Expression() IExpressionContext {
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
	p.EnterRule(localctx, 20, SwiftgrammParserRULE_print_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(SwiftgrammParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(SwiftgrammParserOPEN_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)

		var _x = p.expression(0)

		localctx.(*Print_blContext)._expression = _x
	}
	{
		p.SetState(174)
		p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*Print_blContext).instr = instructions.NewPrint(localctx.(*Print_blContext).Get_expression().GetP())

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
	p.EnterRule(localctx, 22, SwiftgrammParserRULE_if_bl)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(179)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(181)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*If_blContext).instr = instructions.NewIf(localctx.(*If_blContext).Get_expression().GetP(), localctx.(*If_blContext).GetIfblock().GetBlk(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(186)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(188)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)

			var _x = p.Block()

			localctx.(*If_blContext).elseblock = _x
		}
		{
			p.SetState(192)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*If_blContext).instr = instructions.NewIf(localctx.(*If_blContext).Get_expression().GetP(), localctx.(*If_blContext).GetIfblock().GetBlk(), localctx.(*If_blContext).GetElseblock().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(197)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(199)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)

			var _x = p.Else_if()

			localctx.(*If_blContext)._else_if = _x
		}

		localctx.(*If_blContext).instr = instructions.NewIf(localctx.(*If_blContext).Get_expression().GetP(), localctx.(*If_blContext).GetIfblock().GetBlk(), localctx.(*If_blContext).Get_else_if().GetInstr())

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
	p.EnterRule(localctx, 24, SwiftgrammParserRULE_else_if)
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(208)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(210)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Else_ifContext).instr = []interface{}{instructions.NewIf(localctx.(*Else_ifContext).Get_expression().GetP(), localctx.(*Else_ifContext).GetIfblock().GetBlk(), nil)}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(216)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(218)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)

			var _x = p.Block()

			localctx.(*Else_ifContext).elseblock = _x
		}
		{
			p.SetState(222)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Else_ifContext).instr = []interface{}{instructions.NewIf(localctx.(*Else_ifContext).Get_expression().GetP(), localctx.(*Else_ifContext).GetIfblock().GetBlk(), localctx.(*Else_ifContext).GetElseblock().GetBlk())}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(228)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(230)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)

			var _x = p.Else_if()

			localctx.(*Else_ifContext)._else_if = _x
		}

		localctx.(*Else_ifContext).instr = []interface{}{instructions.NewIf(localctx.(*Else_ifContext).Get_expression().GetP(), localctx.(*Else_ifContext).GetIfblock().GetBlk(), localctx.(*Else_ifContext).Get_else_if().GetInstr())}

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
	p.EnterRule(localctx, 26, SwiftgrammParserRULE_while_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(SwiftgrammParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)

		var _x = p.expression(0)

		localctx.(*While_blContext)._expression = _x
	}
	{
		p.SetState(238)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)

		var _x = p.Block()

		localctx.(*While_blContext)._block = _x
	}
	{
		p.SetState(240)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*While_blContext).instr = instructions.NewWhile(localctx.(*While_blContext).Get_expression().GetP(), localctx.(*While_blContext).Get_block().GetBlk())

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
	p.EnterRule(localctx, 28, SwiftgrammParserRULE_for_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(SwiftgrammParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*For_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(SwiftgrammParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)

		var _x = p.expression(0)

		localctx.(*For_blContext).expression1 = _x
	}
	{
		p.SetState(247)
		p.Match(SwiftgrammParserRANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)

		var _x = p.expression(0)

		localctx.(*For_blContext).expression2 = _x
	}
	{
		p.SetState(249)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)

		var _x = p.Block()

		localctx.(*For_blContext)._block = _x
	}
	{
		p.SetState(251)
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
	}()), localctx.(*For_blContext).GetExpression1().GetP(), localctx.(*For_blContext).GetExpression2().GetP(), localctx.(*For_blContext).Get_block().GetBlk())

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
	p.EnterRule(localctx, 30, SwiftgrammParserRULE_guard_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(SwiftgrammParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)

		var _x = p.expression(0)

		localctx.(*Guard_blContext)._expression = _x
	}
	{
		p.SetState(256)
		p.Match(SwiftgrammParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)

		var _x = p.Block()

		localctx.(*Guard_blContext)._block = _x
	}
	{
		p.SetState(259)
		p.Match(SwiftgrammParserCLOSE_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*Guard_blContext).instr = instructions.NewGuard(localctx.(*Guard_blContext).Get_expression().GetP(), localctx.(*Guard_blContext).Get_block().GetBlk())

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
	p.EnterRule(localctx, 32, SwiftgrammParserRULE_vector_bl)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Vector_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)

			var _x = p.Datatype()

			localctx.(*Vector_blContext)._datatype = _x
		}
		{
			p.SetState(267)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)

			var _x = p.Array_exp()

			localctx.(*Vector_blContext)._array_exp = _x
		}
		{
			p.SetState(271)
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
		}()), localctx.(*Vector_blContext).Get_datatype().GetTd(), localctx.(*Vector_blContext).Get_array_exp().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Vector_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)

			var _x = p.Datatype()

			localctx.(*Vector_blContext)._datatype = _x
		}
		{
			p.SetState(279)
			p.Match(SwiftgrammParserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(SwiftgrammParserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
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
		}()), localctx.(*Vector_blContext).Get_datatype().GetTd(), nil)

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
	p.EnterRule(localctx, 34, SwiftgrammParserRULE_array_exp)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)

			var _x = p.expression(0)

			localctx.(*Array_expContext)._expression = _x
		}
		{
			p.SetState(288)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)

			var _x = p.Array_exp()

			localctx.(*Array_expContext)._array_exp = _x
		}

		localctx.(*Array_expContext).p = append(localctx.(*Array_expContext).Get_array_exp().GetP(), localctx.(*Array_expContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)

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
	p.EnterRule(localctx, 36, SwiftgrammParserRULE_function_bl)
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Match(SwiftgrammParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Function_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Match(SwiftgrammParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)

			var _x = p.Datatype()

			localctx.(*Function_blContext)._datatype = _x
		}
		{
			p.SetState(303)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)

			var _x = p.Block()

			localctx.(*Function_blContext)._block = _x
		}
		{
			p.SetState(305)
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
		}()), localctx.(*Function_blContext).Get_datatype().GetTd(), []interface{}{}, localctx.(*Function_blContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Match(SwiftgrammParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Function_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)

			var _x = p.Block()

			localctx.(*Function_blContext)._block = _x
		}
		{
			p.SetState(314)
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
		}()), symbol.NIL, []interface{}{}, localctx.(*Function_blContext).Get_block().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(317)
			p.Match(SwiftgrammParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Function_blContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)

			var _x = p.Params()

			localctx.(*Function_blContext)._params = _x
		}
		{
			p.SetState(321)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(SwiftgrammParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)

			var _x = p.Datatype()

			localctx.(*Function_blContext)._datatype = _x
		}
		{
			p.SetState(324)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)

			var _x = p.Block()

			localctx.(*Function_blContext)._block = _x
		}
		{
			p.SetState(326)
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
		}()), localctx.(*Function_blContext).Get_datatype().GetTd(), localctx.(*Function_blContext).Get_params().GetP(), localctx.(*Function_blContext).Get_block().GetBlk())

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

	// Set_datatype sets the _datatype rule contexts.
	Set_datatype(IDatatypeContext)

	// Set_params sets the _params rule contexts.
	Set_params(IParamsContext)

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

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         []interface{}
	_ID       antlr.Token
	_datatype IDatatypeContext
	_params   IParamsContext
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

func (s *ParamsContext) Set_datatype(v IDatatypeContext) { s._datatype = v }

func (s *ParamsContext) Set_params(v IParamsContext) { s._params = v }

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
	p.EnterRule(localctx, 38, SwiftgrammParserRULE_params)
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}
		{
			p.SetState(334)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)

			var _x = p.Params()

			localctx.(*ParamsContext)._params = _x
		}

		localctx.(*ParamsContext).p = append(localctx.(*ParamsContext).Get_params().GetP(), instructions.NewDeclareWithoutValue((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), expressions.NewNative(nil, symbol.NIL)))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(338)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*ParamsContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)

			var _x = p.Datatype()

			localctx.(*ParamsContext)._datatype = _x
		}

		localctx.(*ParamsContext).p = []interface{}{instructions.NewDeclareWithoutValue((func() string {
			if localctx.(*ParamsContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamsContext).Get_ID().GetText()
			}
		}()), localctx.(*ParamsContext).Get_datatype().GetTd(), expressions.NewNative(nil, symbol.NIL))}

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
	p.EnterRule(localctx, 40, SwiftgrammParserRULE_call_function)
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Call_functionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
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
		}()), []interface{}{})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Call_functionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)

			var _x = p.List_exp()

			localctx.(*Call_functionContext)._list_exp = _x
		}
		{
			p.SetState(352)
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
		}()), localctx.(*Call_functionContext).Get_list_exp().GetP())

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

	// IsList_expContext differentiates from other interfaces.
	IsList_expContext()
}

type List_expContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           []interface{}
	_expression IExpressionContext
	_list_exp   IList_expContext
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
	p.EnterRule(localctx, 42, SwiftgrammParserRULE_list_exp)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)

			var _x = p.expression(0)

			localctx.(*List_expContext)._expression = _x
		}
		{
			p.SetState(358)
			p.Match(SwiftgrammParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)

			var _x = p.List_exp()

			localctx.(*List_expContext)._list_exp = _x
		}

		localctx.(*List_expContext).p = append(localctx.(*List_expContext).Get_list_exp().GetP(), localctx.(*List_expContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)

			var _x = p.expression(0)

			localctx.(*List_expContext)._expression = _x
		}

		localctx.(*List_expContext).p = []interface{}{localctx.(*List_expContext).Get_expression().GetP()}

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
	p.EnterRule(localctx, 44, SwiftgrammParserRULE_call_function_exp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)

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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOper returns the oper token.
	GetOper() antlr.Token

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

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// SetOper sets the oper token.
	SetOper(antlr.Token)

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

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_call_function_exp returns the _call_function_exp rule contexts.
	Get_call_function_exp() ICall_function_expContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_call_function_exp sets the _call_function_exp rule contexts.
	Set_call_function_exp(ICall_function_expContext)

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
	OPEN_PARENTHESIS() antlr.TerminalNode
	CLOSE_PARENTHESIS() antlr.TerminalNode
	Call_function_exp() ICall_function_expContext
	NUMBER() antlr.TerminalNode
	FLOATT() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	CHARACTER_LITERAL() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	ID() antlr.TerminalNode
	NIL() antlr.TerminalNode
	MULTIPLICATION() antlr.TerminalNode
	DIVISION() antlr.TerminalNode
	SUMMATION() antlr.TerminalNode
	SUBTRACTION() antlr.TerminalNode
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
	_NUMBER            antlr.Token
	_FLOATT            antlr.Token
	_STRING_LITERAL    antlr.Token
	_CHARACTER_LITERAL antlr.Token
	_TRUE              antlr.Token
	_FALSE             antlr.Token
	_ID                antlr.Token
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

func (s *ExpressionContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExpressionContext) Get_FLOATT() antlr.Token { return s._FLOATT }

func (s *ExpressionContext) Get_STRING_LITERAL() antlr.Token { return s._STRING_LITERAL }

func (s *ExpressionContext) Get_CHARACTER_LITERAL() antlr.Token { return s._CHARACTER_LITERAL }

func (s *ExpressionContext) Get_TRUE() antlr.Token { return s._TRUE }

func (s *ExpressionContext) Get_FALSE() antlr.Token { return s._FALSE }

func (s *ExpressionContext) Get_ID() antlr.Token { return s._ID }

func (s *ExpressionContext) SetOper(v antlr.Token) { s.oper = v }

func (s *ExpressionContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExpressionContext) Set_FLOATT(v antlr.Token) { s._FLOATT = v }

func (s *ExpressionContext) Set_STRING_LITERAL(v antlr.Token) { s._STRING_LITERAL = v }

func (s *ExpressionContext) Set_CHARACTER_LITERAL(v antlr.Token) { s._CHARACTER_LITERAL = v }

func (s *ExpressionContext) Set_TRUE(v antlr.Token) { s._TRUE = v }

func (s *ExpressionContext) Set_FALSE(v antlr.Token) { s._FALSE = v }

func (s *ExpressionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *ExpressionContext) Get_expression() IExpressionContext { return s._expression }

func (s *ExpressionContext) Get_call_function_exp() ICall_function_expContext {
	return s._call_function_exp
}

func (s *ExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ExpressionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ExpressionContext) Set_call_function_exp(v ICall_function_expContext) {
	s._call_function_exp = v
}

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

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserID, 0)
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

func (s *ExpressionContext) SUBTRACTION() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSUBTRACTION, 0)
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
	_startState := 46
	p.EnterRecursionRule(localctx, 46, SwiftgrammParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(371)

			var _m = p.Match(SwiftgrammParserNOT)

			localctx.(*ExpressionContext).oper = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)

			var _x = p.expression(11)

			localctx.(*ExpressionContext)._expression = _x
		}

		localctx.(*ExpressionContext).p = expressions.NewLogicOperations(nil, localctx.(*ExpressionContext).Get_expression().GetP(), (func() string {
			if localctx.(*ExpressionContext).GetOper() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).GetOper().GetText()
			}
		}()))

	case 2:
		{
			p.SetState(375)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)

			var _x = p.expression(0)

			localctx.(*ExpressionContext)._expression = _x
		}
		{
			p.SetState(377)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(379)

			var _x = p.Call_function_exp()

			localctx.(*ExpressionContext)._call_function_exp = _x
		}

		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_call_function_exp().GetP()

	case 4:
		{
			p.SetState(382)

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
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.INT)

	case 5:
		{
			p.SetState(384)

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
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.FLOAT)

	case 6:
		{
			p.SetState(386)

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
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.STRING)

	case 7:
		{
			p.SetState(388)

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
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.CHAR)

	case 8:
		{
			p.SetState(390)

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
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.BOOL)

	case 9:
		{
			p.SetState(392)

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
		localctx.(*ExpressionContext).p = expressions.NewNative(value, symbol.BOOL)

	case 10:
		{
			p.SetState(394)

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
		}()))

	case 11:
		{
			p.SetState(396)
			p.Match(SwiftgrammParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = expressions.NewNative(nil, symbol.NIL)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(435)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(400)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(401)

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
					p.SetState(402)

					var _x = p.expression(19)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewArithmeticOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()))

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(405)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(406)

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
					p.SetState(407)

					var _x = p.expression(18)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewArithmeticOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()))

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(410)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(411)

					var _m = p.Match(SwiftgrammParserMOD)

					localctx.(*ExpressionContext).oper = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(412)

					var _x = p.expression(17)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewArithmeticOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()))

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(415)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(416)

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
					p.SetState(417)

					var _x = p.expression(16)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewRelationalOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()))

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(421)

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
					p.SetState(422)

					var _x = p.expression(15)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewRelationalOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()))

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(425)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(426)

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
					p.SetState(427)

					var _x = p.expression(14)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewRelationalOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()))

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(430)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(431)

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
					p.SetState(432)

					var _x = p.expression(13)

					localctx.(*ExpressionContext).right = _x
					localctx.(*ExpressionContext)._expression = _x
				}

				localctx.(*ExpressionContext).p = expressions.NewLogicOperations(localctx.(*ExpressionContext).GetLeft().GetP(), localctx.(*ExpressionContext).GetRight().GetP(), (func() string {
					if localctx.(*ExpressionContext).GetOper() == nil {
						return ""
					} else {
						return localctx.(*ExpressionContext).GetOper().GetText()
					}
				}()))

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 48, SwiftgrammParserRULE_datatype)
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(440)
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
			p.SetState(442)
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
			p.SetState(444)
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
			p.SetState(446)
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
			p.SetState(448)
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
	case 23:
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
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
