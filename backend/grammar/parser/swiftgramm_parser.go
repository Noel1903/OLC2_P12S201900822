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
		"'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE",
		"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK",
		"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC",
		"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "FLOATT", "STRING_LITERAL",
		"ID", "CHARACTER_LITERAL", "INCREMENT", "DECREMENT", "RANGE", "SUMMATION",
		"SUBTRACTION", "MULTIPLICATION", "DIVISION", "MOD", "QUESTION_MARK",
		"OR", "AND", "NOT", "EQUAL", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL",
		"GREATER_THAN", "GREATER_THAN_EQUAL", "ASSIGN", "DOT", "COMMA", "COLON",
		"SEMICOLON", "OPEN_PARENTHESIS", "CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "sentence", "increment_bl", "decrement_bl", "break_bl",
		"return_bl", "declare_let", "declare_var", "print_bl", "if_bl", "else_if",
		"while_bl", "for_bl", "expression", "datatype",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 298, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 38, 8, 1, 11, 1, 12, 1, 39, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 74, 8, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 95, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 111, 8, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 134, 8, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 168, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		3, 11, 199, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 244, 8, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		5, 14, 281, 8, 14, 10, 14, 12, 14, 284, 9, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 296, 8, 15, 1, 15,
		0, 1, 28, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		0, 6, 1, 0, 39, 40, 1, 0, 37, 38, 1, 0, 48, 49, 1, 0, 50, 51, 1, 0, 46,
		47, 1, 0, 43, 44, 319, 0, 32, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 73, 1,
		0, 0, 0, 6, 75, 1, 0, 0, 0, 8, 80, 1, 0, 0, 0, 10, 85, 1, 0, 0, 0, 12,
		94, 1, 0, 0, 0, 14, 110, 1, 0, 0, 0, 16, 133, 1, 0, 0, 0, 18, 135, 1, 0,
		0, 0, 20, 167, 1, 0, 0, 0, 22, 198, 1, 0, 0, 0, 24, 200, 1, 0, 0, 0, 26,
		207, 1, 0, 0, 0, 28, 243, 1, 0, 0, 0, 30, 295, 1, 0, 0, 0, 32, 33, 3, 2,
		1, 0, 33, 34, 5, 0, 0, 1, 34, 35, 6, 0, -1, 0, 35, 1, 1, 0, 0, 0, 36, 38,
		3, 4, 2, 0, 37, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0,
		39, 40, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 6, 1, -1, 0, 42, 3, 1,
		0, 0, 0, 43, 44, 3, 16, 8, 0, 44, 45, 6, 2, -1, 0, 45, 74, 1, 0, 0, 0,
		46, 47, 3, 14, 7, 0, 47, 48, 6, 2, -1, 0, 48, 74, 1, 0, 0, 0, 49, 50, 3,
		18, 9, 0, 50, 51, 6, 2, -1, 0, 51, 74, 1, 0, 0, 0, 52, 53, 3, 20, 10, 0,
		53, 54, 6, 2, -1, 0, 54, 74, 1, 0, 0, 0, 55, 56, 3, 6, 3, 0, 56, 57, 6,
		2, -1, 0, 57, 74, 1, 0, 0, 0, 58, 59, 3, 8, 4, 0, 59, 60, 6, 2, -1, 0,
		60, 74, 1, 0, 0, 0, 61, 62, 3, 24, 12, 0, 62, 63, 6, 2, -1, 0, 63, 74,
		1, 0, 0, 0, 64, 65, 3, 26, 13, 0, 65, 66, 6, 2, -1, 0, 66, 74, 1, 0, 0,
		0, 67, 68, 3, 10, 5, 0, 68, 69, 6, 2, -1, 0, 69, 74, 1, 0, 0, 0, 70, 71,
		3, 12, 6, 0, 71, 72, 6, 2, -1, 0, 72, 74, 1, 0, 0, 0, 73, 43, 1, 0, 0,
		0, 73, 46, 1, 0, 0, 0, 73, 49, 1, 0, 0, 0, 73, 52, 1, 0, 0, 0, 73, 55,
		1, 0, 0, 0, 73, 58, 1, 0, 0, 0, 73, 61, 1, 0, 0, 0, 73, 64, 1, 0, 0, 0,
		73, 67, 1, 0, 0, 0, 73, 70, 1, 0, 0, 0, 74, 5, 1, 0, 0, 0, 75, 76, 5, 32,
		0, 0, 76, 77, 5, 34, 0, 0, 77, 78, 3, 28, 14, 0, 78, 79, 6, 3, -1, 0, 79,
		7, 1, 0, 0, 0, 80, 81, 5, 32, 0, 0, 81, 82, 5, 35, 0, 0, 82, 83, 3, 28,
		14, 0, 83, 84, 6, 4, -1, 0, 84, 9, 1, 0, 0, 0, 85, 86, 5, 16, 0, 0, 86,
		87, 6, 5, -1, 0, 87, 11, 1, 0, 0, 0, 88, 89, 5, 23, 0, 0, 89, 90, 3, 28,
		14, 0, 90, 91, 6, 6, -1, 0, 91, 95, 1, 0, 0, 0, 92, 93, 5, 23, 0, 0, 93,
		95, 6, 6, -1, 0, 94, 88, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 13, 1, 0,
		0, 0, 96, 97, 5, 10, 0, 0, 97, 98, 5, 32, 0, 0, 98, 99, 5, 55, 0, 0, 99,
		100, 3, 30, 15, 0, 100, 101, 5, 52, 0, 0, 101, 102, 3, 28, 14, 0, 102,
		103, 6, 7, -1, 0, 103, 111, 1, 0, 0, 0, 104, 105, 5, 10, 0, 0, 105, 106,
		5, 32, 0, 0, 106, 107, 5, 52, 0, 0, 107, 108, 3, 28, 14, 0, 108, 109, 6,
		7, -1, 0, 109, 111, 1, 0, 0, 0, 110, 96, 1, 0, 0, 0, 110, 104, 1, 0, 0,
		0, 111, 15, 1, 0, 0, 0, 112, 113, 5, 9, 0, 0, 113, 114, 5, 32, 0, 0, 114,
		115, 5, 55, 0, 0, 115, 116, 3, 30, 15, 0, 116, 117, 5, 52, 0, 0, 117, 118,
		3, 28, 14, 0, 118, 119, 6, 8, -1, 0, 119, 134, 1, 0, 0, 0, 120, 121, 5,
		9, 0, 0, 121, 122, 5, 32, 0, 0, 122, 123, 5, 52, 0, 0, 123, 124, 3, 28,
		14, 0, 124, 125, 6, 8, -1, 0, 125, 134, 1, 0, 0, 0, 126, 127, 5, 9, 0,
		0, 127, 128, 5, 32, 0, 0, 128, 129, 5, 55, 0, 0, 129, 130, 3, 30, 15, 0,
		130, 131, 5, 42, 0, 0, 131, 132, 6, 8, -1, 0, 132, 134, 1, 0, 0, 0, 133,
		112, 1, 0, 0, 0, 133, 120, 1, 0, 0, 0, 133, 126, 1, 0, 0, 0, 134, 17, 1,
		0, 0, 0, 135, 136, 5, 11, 0, 0, 136, 137, 5, 57, 0, 0, 137, 138, 3, 28,
		14, 0, 138, 139, 5, 58, 0, 0, 139, 140, 6, 9, -1, 0, 140, 19, 1, 0, 0,
		0, 141, 142, 5, 12, 0, 0, 142, 143, 3, 28, 14, 0, 143, 144, 5, 59, 0, 0,
		144, 145, 3, 2, 1, 0, 145, 146, 5, 60, 0, 0, 146, 147, 6, 10, -1, 0, 147,
		168, 1, 0, 0, 0, 148, 149, 5, 12, 0, 0, 149, 150, 3, 28, 14, 0, 150, 151,
		5, 59, 0, 0, 151, 152, 3, 2, 1, 0, 152, 153, 5, 60, 0, 0, 153, 154, 5,
		13, 0, 0, 154, 155, 5, 59, 0, 0, 155, 156, 3, 2, 1, 0, 156, 157, 5, 60,
		0, 0, 157, 158, 6, 10, -1, 0, 158, 168, 1, 0, 0, 0, 159, 160, 5, 12, 0,
		0, 160, 161, 3, 28, 14, 0, 161, 162, 5, 59, 0, 0, 162, 163, 3, 2, 1, 0,
		163, 164, 5, 60, 0, 0, 164, 165, 3, 22, 11, 0, 165, 166, 6, 10, -1, 0,
		166, 168, 1, 0, 0, 0, 167, 141, 1, 0, 0, 0, 167, 148, 1, 0, 0, 0, 167,
		159, 1, 0, 0, 0, 168, 21, 1, 0, 0, 0, 169, 170, 5, 13, 0, 0, 170, 171,
		5, 12, 0, 0, 171, 172, 3, 28, 14, 0, 172, 173, 5, 59, 0, 0, 173, 174, 3,
		2, 1, 0, 174, 175, 5, 60, 0, 0, 175, 176, 6, 11, -1, 0, 176, 199, 1, 0,
		0, 0, 177, 178, 5, 13, 0, 0, 178, 179, 5, 12, 0, 0, 179, 180, 3, 28, 14,
		0, 180, 181, 5, 59, 0, 0, 181, 182, 3, 2, 1, 0, 182, 183, 5, 60, 0, 0,
		183, 184, 5, 13, 0, 0, 184, 185, 5, 59, 0, 0, 185, 186, 3, 2, 1, 0, 186,
		187, 5, 60, 0, 0, 187, 188, 6, 11, -1, 0, 188, 199, 1, 0, 0, 0, 189, 190,
		5, 13, 0, 0, 190, 191, 5, 12, 0, 0, 191, 192, 3, 28, 14, 0, 192, 193, 5,
		59, 0, 0, 193, 194, 3, 2, 1, 0, 194, 195, 5, 60, 0, 0, 195, 196, 3, 22,
		11, 0, 196, 197, 6, 11, -1, 0, 197, 199, 1, 0, 0, 0, 198, 169, 1, 0, 0,
		0, 198, 177, 1, 0, 0, 0, 198, 189, 1, 0, 0, 0, 199, 23, 1, 0, 0, 0, 200,
		201, 5, 18, 0, 0, 201, 202, 3, 28, 14, 0, 202, 203, 5, 59, 0, 0, 203, 204,
		3, 2, 1, 0, 204, 205, 5, 60, 0, 0, 205, 206, 6, 12, -1, 0, 206, 25, 1,
		0, 0, 0, 207, 208, 5, 19, 0, 0, 208, 209, 5, 32, 0, 0, 209, 210, 5, 20,
		0, 0, 210, 211, 3, 28, 14, 0, 211, 212, 5, 36, 0, 0, 212, 213, 3, 28, 14,
		0, 213, 214, 5, 59, 0, 0, 214, 215, 3, 2, 1, 0, 215, 216, 5, 60, 0, 0,
		216, 217, 6, 13, -1, 0, 217, 27, 1, 0, 0, 0, 218, 219, 6, 14, -1, 0, 219,
		220, 5, 45, 0, 0, 220, 221, 3, 28, 14, 10, 221, 222, 6, 14, -1, 0, 222,
		244, 1, 0, 0, 0, 223, 224, 5, 57, 0, 0, 224, 225, 3, 28, 14, 0, 225, 226,
		5, 58, 0, 0, 226, 244, 1, 0, 0, 0, 227, 228, 5, 29, 0, 0, 228, 244, 6,
		14, -1, 0, 229, 230, 5, 30, 0, 0, 230, 244, 6, 14, -1, 0, 231, 232, 5,
		31, 0, 0, 232, 244, 6, 14, -1, 0, 233, 234, 5, 33, 0, 0, 234, 244, 6, 14,
		-1, 0, 235, 236, 5, 6, 0, 0, 236, 244, 6, 14, -1, 0, 237, 238, 5, 7, 0,
		0, 238, 244, 6, 14, -1, 0, 239, 240, 5, 32, 0, 0, 240, 244, 6, 14, -1,
		0, 241, 242, 5, 8, 0, 0, 242, 244, 6, 14, -1, 0, 243, 218, 1, 0, 0, 0,
		243, 223, 1, 0, 0, 0, 243, 227, 1, 0, 0, 0, 243, 229, 1, 0, 0, 0, 243,
		231, 1, 0, 0, 0, 243, 233, 1, 0, 0, 0, 243, 235, 1, 0, 0, 0, 243, 237,
		1, 0, 0, 0, 243, 239, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 244, 282, 1, 0,
		0, 0, 245, 246, 10, 17, 0, 0, 246, 247, 7, 0, 0, 0, 247, 248, 3, 28, 14,
		18, 248, 249, 6, 14, -1, 0, 249, 281, 1, 0, 0, 0, 250, 251, 10, 16, 0,
		0, 251, 252, 7, 1, 0, 0, 252, 253, 3, 28, 14, 17, 253, 254, 6, 14, -1,
		0, 254, 281, 1, 0, 0, 0, 255, 256, 10, 15, 0, 0, 256, 257, 5, 41, 0, 0,
		257, 258, 3, 28, 14, 16, 258, 259, 6, 14, -1, 0, 259, 281, 1, 0, 0, 0,
		260, 261, 10, 14, 0, 0, 261, 262, 7, 2, 0, 0, 262, 263, 3, 28, 14, 15,
		263, 264, 6, 14, -1, 0, 264, 281, 1, 0, 0, 0, 265, 266, 10, 13, 0, 0, 266,
		267, 7, 3, 0, 0, 267, 268, 3, 28, 14, 14, 268, 269, 6, 14, -1, 0, 269,
		281, 1, 0, 0, 0, 270, 271, 10, 12, 0, 0, 271, 272, 7, 4, 0, 0, 272, 273,
		3, 28, 14, 13, 273, 274, 6, 14, -1, 0, 274, 281, 1, 0, 0, 0, 275, 276,
		10, 11, 0, 0, 276, 277, 7, 5, 0, 0, 277, 278, 3, 28, 14, 12, 278, 279,
		6, 14, -1, 0, 279, 281, 1, 0, 0, 0, 280, 245, 1, 0, 0, 0, 280, 250, 1,
		0, 0, 0, 280, 255, 1, 0, 0, 0, 280, 260, 1, 0, 0, 0, 280, 265, 1, 0, 0,
		0, 280, 270, 1, 0, 0, 0, 280, 275, 1, 0, 0, 0, 281, 284, 1, 0, 0, 0, 282,
		280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 29, 1, 0, 0, 0, 284, 282, 1,
		0, 0, 0, 285, 286, 5, 1, 0, 0, 286, 296, 6, 15, -1, 0, 287, 288, 5, 2,
		0, 0, 288, 296, 6, 15, -1, 0, 289, 290, 5, 31, 0, 0, 290, 296, 6, 15, -1,
		0, 291, 292, 5, 4, 0, 0, 292, 296, 6, 15, -1, 0, 293, 294, 5, 33, 0, 0,
		294, 296, 6, 15, -1, 0, 295, 285, 1, 0, 0, 0, 295, 287, 1, 0, 0, 0, 295,
		289, 1, 0, 0, 0, 295, 291, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 31, 1,
		0, 0, 0, 11, 39, 73, 94, 110, 133, 167, 198, 243, 280, 282, 295,
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
	SwiftgrammParserSTRING_LITERAL     = 31
	SwiftgrammParserID                 = 32
	SwiftgrammParserCHARACTER_LITERAL  = 33
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
	SwiftgrammParserWHITESPACE         = 61
	SwiftgrammParserCOMMENT            = 62
	SwiftgrammParserLINE_COMMENT       = 63
)

// SwiftgrammParser rules.
const (
	SwiftgrammParserRULE_s            = 0
	SwiftgrammParserRULE_block        = 1
	SwiftgrammParserRULE_sentence     = 2
	SwiftgrammParserRULE_increment_bl = 3
	SwiftgrammParserRULE_decrement_bl = 4
	SwiftgrammParserRULE_break_bl     = 5
	SwiftgrammParserRULE_return_bl    = 6
	SwiftgrammParserRULE_declare_let  = 7
	SwiftgrammParserRULE_declare_var  = 8
	SwiftgrammParserRULE_print_bl     = 9
	SwiftgrammParserRULE_if_bl        = 10
	SwiftgrammParserRULE_else_if      = 11
	SwiftgrammParserRULE_while_bl     = 12
	SwiftgrammParserRULE_for_bl       = 13
	SwiftgrammParserRULE_expression   = 14
	SwiftgrammParserRULE_datatype     = 15
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
		p.SetState(32)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(33)
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4304215552) != 0) {
		{
			p.SetState(36)

			var _x = p.Sentence()

			localctx.(*BlockContext)._sentence = _x
		}
		localctx.(*BlockContext).instr = append(localctx.(*BlockContext).instr, localctx.(*BlockContext)._sentence)

		p.SetState(39)
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

	// Get_break_bl returns the _break_bl rule contexts.
	Get_break_bl() IBreak_blContext

	// Get_return_bl returns the _return_bl rule contexts.
	Get_return_bl() IReturn_blContext

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

	// Set_break_bl sets the _break_bl rule contexts.
	Set_break_bl(IBreak_blContext)

	// Set_return_bl sets the _return_bl rule contexts.
	Set_return_bl(IReturn_blContext)

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
	Break_bl() IBreak_blContext
	Return_bl() IReturn_blContext

	// IsSentenceContext differentiates from other interfaces.
	IsSentenceContext()
}

type SentenceContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	instr         abstract.Instruction
	_declare_var  IDeclare_varContext
	_declare_let  IDeclare_letContext
	_print_bl     IPrint_blContext
	_if_bl        IIf_blContext
	_increment_bl IIncrement_blContext
	_decrement_bl IDecrement_blContext
	_while_bl     IWhile_blContext
	_for_bl       IFor_blContext
	_break_bl     IBreak_blContext
	_return_bl    IReturn_blContext
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

func (s *SentenceContext) Get_break_bl() IBreak_blContext { return s._break_bl }

func (s *SentenceContext) Get_return_bl() IReturn_blContext { return s._return_bl }

func (s *SentenceContext) Set_declare_var(v IDeclare_varContext) { s._declare_var = v }

func (s *SentenceContext) Set_declare_let(v IDeclare_letContext) { s._declare_let = v }

func (s *SentenceContext) Set_print_bl(v IPrint_blContext) { s._print_bl = v }

func (s *SentenceContext) Set_if_bl(v IIf_blContext) { s._if_bl = v }

func (s *SentenceContext) Set_increment_bl(v IIncrement_blContext) { s._increment_bl = v }

func (s *SentenceContext) Set_decrement_bl(v IDecrement_blContext) { s._decrement_bl = v }

func (s *SentenceContext) Set_while_bl(v IWhile_blContext) { s._while_bl = v }

func (s *SentenceContext) Set_for_bl(v IFor_blContext) { s._for_bl = v }

func (s *SentenceContext) Set_break_bl(v IBreak_blContext) { s._break_bl = v }

func (s *SentenceContext) Set_return_bl(v IReturn_blContext) { s._return_bl = v }

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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)

			var _x = p.Declare_var()

			localctx.(*SentenceContext)._declare_var = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_declare_var().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)

			var _x = p.Declare_let()

			localctx.(*SentenceContext)._declare_let = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_declare_let().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)

			var _x = p.Print_bl()

			localctx.(*SentenceContext)._print_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_print_bl().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(52)

			var _x = p.If_bl()

			localctx.(*SentenceContext)._if_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_if_bl().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(55)

			var _x = p.Increment_bl()

			localctx.(*SentenceContext)._increment_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_increment_bl().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(58)

			var _x = p.Decrement_bl()

			localctx.(*SentenceContext)._decrement_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_decrement_bl().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(61)

			var _x = p.While_bl()

			localctx.(*SentenceContext)._while_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_while_bl().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(64)

			var _x = p.For_bl()

			localctx.(*SentenceContext)._for_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_for_bl().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(67)

			var _x = p.Break_bl()

			localctx.(*SentenceContext)._break_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_break_bl().GetInstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(70)

			var _x = p.Return_bl()

			localctx.(*SentenceContext)._return_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_return_bl().GetInstr()

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
		p.SetState(75)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Increment_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)

		var _m = p.Match(SwiftgrammParserINCREMENT)

		localctx.(*Increment_blContext)._INCREMENT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)

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
		p.SetState(80)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*Decrement_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)

		var _m = p.Match(SwiftgrammParserDECREMENT)

		localctx.(*Decrement_blContext)._DECREMENT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)

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
		p.SetState(85)
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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(SwiftgrammParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)

			var _x = p.expression(0)

			localctx.(*Return_blContext)._expression = _x
		}

		localctx.(*Return_blContext).instr = instructions.NewReturn(localctx.(*Return_blContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
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
	p.EnterRule(localctx, 14, SwiftgrammParserRULE_declare_let)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
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

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_letContext)._ID = _m
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

			var _x = p.Datatype()

			localctx.(*Declare_letContext)._datatype = _x
		}
		{
			p.SetState(100)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)

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
			p.SetState(104)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_letContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)

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
	p.EnterRule(localctx, 16, SwiftgrammParserRULE_declare_var)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)

			var _x = p.Datatype()

			localctx.(*Declare_varContext)._datatype = _x
		}
		{
			p.SetState(116)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)

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
			p.SetState(120)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)

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
			p.SetState(126)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)

			var _x = p.Datatype()

			localctx.(*Declare_varContext)._datatype = _x
		}
		{
			p.SetState(130)
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
	p.EnterRule(localctx, 18, SwiftgrammParserRULE_print_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(SwiftgrammParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(SwiftgrammParserOPEN_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)

		var _x = p.expression(0)

		localctx.(*Print_blContext)._expression = _x
	}
	{
		p.SetState(138)
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
	p.EnterRule(localctx, 20, SwiftgrammParserRULE_if_bl)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(143)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(145)
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
			p.SetState(148)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(150)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(152)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)

			var _x = p.Block()

			localctx.(*If_blContext).elseblock = _x
		}
		{
			p.SetState(156)
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
			p.SetState(159)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)

			var _x = p.expression(0)

			localctx.(*If_blContext)._expression = _x
		}
		{
			p.SetState(161)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)

			var _x = p.Block()

			localctx.(*If_blContext).ifblock = _x
		}
		{
			p.SetState(163)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)

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
	p.EnterRule(localctx, 22, SwiftgrammParserRULE_else_if)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(172)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(174)
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
			p.SetState(177)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(180)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(182)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)

			var _x = p.Block()

			localctx.(*Else_ifContext).elseblock = _x
		}
		{
			p.SetState(186)
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
			p.SetState(189)
			p.Match(SwiftgrammParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(SwiftgrammParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)

			var _x = p.expression(0)

			localctx.(*Else_ifContext)._expression = _x
		}
		{
			p.SetState(192)
			p.Match(SwiftgrammParserOPEN_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)

			var _x = p.Block()

			localctx.(*Else_ifContext).ifblock = _x
		}
		{
			p.SetState(194)
			p.Match(SwiftgrammParserCLOSE_kEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)

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
	p.EnterRule(localctx, 24, SwiftgrammParserRULE_while_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(SwiftgrammParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)

		var _x = p.expression(0)

		localctx.(*While_blContext)._expression = _x
	}
	{
		p.SetState(202)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)

		var _x = p.Block()

		localctx.(*While_blContext)._block = _x
	}
	{
		p.SetState(204)
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
	p.EnterRule(localctx, 26, SwiftgrammParserRULE_for_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(SwiftgrammParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)

		var _m = p.Match(SwiftgrammParserID)

		localctx.(*For_blContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(SwiftgrammParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)

		var _x = p.expression(0)

		localctx.(*For_blContext).expression1 = _x
	}
	{
		p.SetState(211)
		p.Match(SwiftgrammParserRANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)

		var _x = p.expression(0)

		localctx.(*For_blContext).expression2 = _x
	}
	{
		p.SetState(213)
		p.Match(SwiftgrammParserOPEN_kEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)

		var _x = p.Block()

		localctx.(*For_blContext)._block = _x
	}
	{
		p.SetState(215)
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

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

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

func (s *ExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ExpressionContext) Set_expression(v IExpressionContext) { s._expression = v }

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
	_startState := 28
	p.EnterRecursionRule(localctx, 28, SwiftgrammParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserNOT:
		{
			p.SetState(219)

			var _m = p.Match(SwiftgrammParserNOT)

			localctx.(*ExpressionContext).oper = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)

			var _x = p.expression(10)

			localctx.(*ExpressionContext)._expression = _x
		}

		localctx.(*ExpressionContext).p = expressions.NewLogicOperations(nil, localctx.(*ExpressionContext).Get_expression().GetP(), (func() string {
			if localctx.(*ExpressionContext).GetOper() == nil {
				return ""
			} else {
				return localctx.(*ExpressionContext).GetOper().GetText()
			}
		}()))

	case SwiftgrammParserOPEN_PARENTHESIS:
		{
			p.SetState(223)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)

			var _x = p.expression(0)

			localctx.(*ExpressionContext)._expression = _x
		}
		{
			p.SetState(225)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftgrammParserNUMBER:
		{
			p.SetState(227)

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

	case SwiftgrammParserFLOATT:
		{
			p.SetState(229)

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

	case SwiftgrammParserSTRING_LITERAL:
		{
			p.SetState(231)

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

	case SwiftgrammParserCHARACTER_LITERAL:
		{
			p.SetState(233)

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

	case SwiftgrammParserTRUE:
		{
			p.SetState(235)

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

	case SwiftgrammParserFALSE:
		{
			p.SetState(237)

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

	case SwiftgrammParserID:
		{
			p.SetState(239)

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

	case SwiftgrammParserNIL:
		{
			p.SetState(241)
			p.Match(SwiftgrammParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExpressionContext).p = expressions.NewNative(nil, symbol.NIL)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(246)

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
					p.SetState(247)

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

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(251)

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
					p.SetState(252)

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

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(256)

					var _m = p.Match(SwiftgrammParserMOD)

					localctx.(*ExpressionContext).oper = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(257)

					var _x = p.expression(16)

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
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(261)

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
					p.SetState(262)

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

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(266)

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
					p.SetState(267)

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

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(270)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(271)

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
					p.SetState(272)

					var _x = p.expression(13)

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
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(276)

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
					p.SetState(277)

					var _x = p.expression(12)

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
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	STRING_LITERAL() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHARACTER_LITERAL() antlr.TerminalNode

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

func (s *DatatypeContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserSTRING_LITERAL, 0)
}

func (s *DatatypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserBOOL, 0)
}

func (s *DatatypeContext) CHARACTER_LITERAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammParserCHARACTER_LITERAL, 0)
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
	p.EnterRule(localctx, 30, SwiftgrammParserRULE_datatype)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
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
			p.SetState(287)
			p.Match(SwiftgrammParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*DatatypeContext).td = symbol.FLOAT

	case SwiftgrammParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(289)
			p.Match(SwiftgrammParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*DatatypeContext).td = symbol.STRING

	case SwiftgrammParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(291)
			p.Match(SwiftgrammParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*DatatypeContext).td = symbol.BOOL

	case SwiftgrammParserCHARACTER_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(293)
			p.Match(SwiftgrammParserCHARACTER_LITERAL)
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
	case 14:
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
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
