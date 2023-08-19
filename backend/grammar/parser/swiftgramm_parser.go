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
		"'inout'", "", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'?'",
		"'||'", "'&&'", "'!'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
		"'='", "'.'", "','", "':'", "';'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "TRUE", "FALSE",
		"NIL", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "BREAK",
		"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "FUNC",
		"STRUCT", "MUTATING", "SELF", "INOUT", "NUMBER", "FLOATT", "STRING_LITERAL",
		"ID", "CHARACTER_LITERAL", "SUMMATION", "SUBTRACTION", "MULTIPLICATION",
		"DIVISION", "MOD", "QUESTION_MARK", "OR", "AND", "NOT", "EQUAL", "NOT_EQUAL",
		"LESS_THAN", "LESS_THAN_EQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL",
		"ASSIGN", "DOT", "COMMA", "COLON", "SEMICOLON", "OPEN_PARENTHESIS",
		"CLOSE_PARENTHESIS", "OPEN_kEY", "CLOSE_kEY", "WHITESPACE", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "sentence", "declare_var", "print_bl", "expression", "datatype",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 60, 142, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 20, 8, 1,
		11, 1, 12, 1, 21, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		32, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 55,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 88, 8, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 125, 8,
		5, 10, 5, 12, 5, 128, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 140, 8, 6, 1, 6, 0, 1, 10, 7, 0, 2, 4, 6, 8, 10, 12,
		0, 6, 1, 0, 36, 37, 1, 0, 34, 35, 1, 0, 45, 46, 1, 0, 47, 48, 1, 0, 43,
		44, 1, 0, 40, 41, 158, 0, 14, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 31, 1,
		0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 56, 1, 0, 0, 0, 10, 87, 1, 0, 0, 0, 12,
		139, 1, 0, 0, 0, 14, 15, 3, 2, 1, 0, 15, 16, 5, 0, 0, 1, 16, 17, 6, 0,
		-1, 0, 17, 1, 1, 0, 0, 0, 18, 20, 3, 4, 2, 0, 19, 18, 1, 0, 0, 0, 20, 21,
		1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0,
		23, 24, 6, 1, -1, 0, 24, 3, 1, 0, 0, 0, 25, 26, 3, 6, 3, 0, 26, 27, 6,
		2, -1, 0, 27, 32, 1, 0, 0, 0, 28, 29, 3, 8, 4, 0, 29, 30, 6, 2, -1, 0,
		30, 32, 1, 0, 0, 0, 31, 25, 1, 0, 0, 0, 31, 28, 1, 0, 0, 0, 32, 5, 1, 0,
		0, 0, 33, 34, 5, 9, 0, 0, 34, 35, 5, 32, 0, 0, 35, 36, 5, 52, 0, 0, 36,
		37, 3, 12, 6, 0, 37, 38, 5, 49, 0, 0, 38, 39, 3, 10, 5, 0, 39, 40, 6, 3,
		-1, 0, 40, 55, 1, 0, 0, 0, 41, 42, 5, 9, 0, 0, 42, 43, 5, 32, 0, 0, 43,
		44, 5, 49, 0, 0, 44, 45, 3, 10, 5, 0, 45, 46, 6, 3, -1, 0, 46, 55, 1, 0,
		0, 0, 47, 48, 5, 9, 0, 0, 48, 49, 5, 32, 0, 0, 49, 50, 5, 52, 0, 0, 50,
		51, 3, 12, 6, 0, 51, 52, 5, 39, 0, 0, 52, 53, 6, 3, -1, 0, 53, 55, 1, 0,
		0, 0, 54, 33, 1, 0, 0, 0, 54, 41, 1, 0, 0, 0, 54, 47, 1, 0, 0, 0, 55, 7,
		1, 0, 0, 0, 56, 57, 5, 11, 0, 0, 57, 58, 5, 54, 0, 0, 58, 59, 3, 10, 5,
		0, 59, 60, 5, 55, 0, 0, 60, 61, 6, 4, -1, 0, 61, 9, 1, 0, 0, 0, 62, 63,
		6, 5, -1, 0, 63, 64, 5, 42, 0, 0, 64, 65, 3, 10, 5, 10, 65, 66, 6, 5, -1,
		0, 66, 88, 1, 0, 0, 0, 67, 68, 5, 54, 0, 0, 68, 69, 3, 10, 5, 0, 69, 70,
		5, 55, 0, 0, 70, 88, 1, 0, 0, 0, 71, 72, 5, 29, 0, 0, 72, 88, 6, 5, -1,
		0, 73, 74, 5, 30, 0, 0, 74, 88, 6, 5, -1, 0, 75, 76, 5, 31, 0, 0, 76, 88,
		6, 5, -1, 0, 77, 78, 5, 33, 0, 0, 78, 88, 6, 5, -1, 0, 79, 80, 5, 6, 0,
		0, 80, 88, 6, 5, -1, 0, 81, 82, 5, 7, 0, 0, 82, 88, 6, 5, -1, 0, 83, 84,
		5, 32, 0, 0, 84, 88, 6, 5, -1, 0, 85, 86, 5, 8, 0, 0, 86, 88, 6, 5, -1,
		0, 87, 62, 1, 0, 0, 0, 87, 67, 1, 0, 0, 0, 87, 71, 1, 0, 0, 0, 87, 73,
		1, 0, 0, 0, 87, 75, 1, 0, 0, 0, 87, 77, 1, 0, 0, 0, 87, 79, 1, 0, 0, 0,
		87, 81, 1, 0, 0, 0, 87, 83, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 126, 1,
		0, 0, 0, 89, 90, 10, 17, 0, 0, 90, 91, 7, 0, 0, 0, 91, 92, 3, 10, 5, 18,
		92, 93, 6, 5, -1, 0, 93, 125, 1, 0, 0, 0, 94, 95, 10, 16, 0, 0, 95, 96,
		7, 1, 0, 0, 96, 97, 3, 10, 5, 17, 97, 98, 6, 5, -1, 0, 98, 125, 1, 0, 0,
		0, 99, 100, 10, 15, 0, 0, 100, 101, 5, 38, 0, 0, 101, 102, 3, 10, 5, 16,
		102, 103, 6, 5, -1, 0, 103, 125, 1, 0, 0, 0, 104, 105, 10, 14, 0, 0, 105,
		106, 7, 2, 0, 0, 106, 107, 3, 10, 5, 15, 107, 108, 6, 5, -1, 0, 108, 125,
		1, 0, 0, 0, 109, 110, 10, 13, 0, 0, 110, 111, 7, 3, 0, 0, 111, 112, 3,
		10, 5, 14, 112, 113, 6, 5, -1, 0, 113, 125, 1, 0, 0, 0, 114, 115, 10, 12,
		0, 0, 115, 116, 7, 4, 0, 0, 116, 117, 3, 10, 5, 13, 117, 118, 6, 5, -1,
		0, 118, 125, 1, 0, 0, 0, 119, 120, 10, 11, 0, 0, 120, 121, 7, 5, 0, 0,
		121, 122, 3, 10, 5, 12, 122, 123, 6, 5, -1, 0, 123, 125, 1, 0, 0, 0, 124,
		89, 1, 0, 0, 0, 124, 94, 1, 0, 0, 0, 124, 99, 1, 0, 0, 0, 124, 104, 1,
		0, 0, 0, 124, 109, 1, 0, 0, 0, 124, 114, 1, 0, 0, 0, 124, 119, 1, 0, 0,
		0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127,
		11, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 130, 5, 1, 0, 0, 130, 140, 6,
		6, -1, 0, 131, 132, 5, 2, 0, 0, 132, 140, 6, 6, -1, 0, 133, 134, 5, 31,
		0, 0, 134, 140, 6, 6, -1, 0, 135, 136, 5, 4, 0, 0, 136, 140, 6, 6, -1,
		0, 137, 138, 5, 33, 0, 0, 138, 140, 6, 6, -1, 0, 139, 129, 1, 0, 0, 0,
		139, 131, 1, 0, 0, 0, 139, 133, 1, 0, 0, 0, 139, 135, 1, 0, 0, 0, 139,
		137, 1, 0, 0, 0, 140, 13, 1, 0, 0, 0, 7, 21, 31, 54, 87, 124, 126, 139,
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
	SwiftgrammParserSUMMATION          = 34
	SwiftgrammParserSUBTRACTION        = 35
	SwiftgrammParserMULTIPLICATION     = 36
	SwiftgrammParserDIVISION           = 37
	SwiftgrammParserMOD                = 38
	SwiftgrammParserQUESTION_MARK      = 39
	SwiftgrammParserOR                 = 40
	SwiftgrammParserAND                = 41
	SwiftgrammParserNOT                = 42
	SwiftgrammParserEQUAL              = 43
	SwiftgrammParserNOT_EQUAL          = 44
	SwiftgrammParserLESS_THAN          = 45
	SwiftgrammParserLESS_THAN_EQUAL    = 46
	SwiftgrammParserGREATER_THAN       = 47
	SwiftgrammParserGREATER_THAN_EQUAL = 48
	SwiftgrammParserASSIGN             = 49
	SwiftgrammParserDOT                = 50
	SwiftgrammParserCOMMA              = 51
	SwiftgrammParserCOLON              = 52
	SwiftgrammParserSEMICOLON          = 53
	SwiftgrammParserOPEN_PARENTHESIS   = 54
	SwiftgrammParserCLOSE_PARENTHESIS  = 55
	SwiftgrammParserOPEN_kEY           = 56
	SwiftgrammParserCLOSE_kEY          = 57
	SwiftgrammParserWHITESPACE         = 58
	SwiftgrammParserCOMMENT            = 59
	SwiftgrammParserLINE_COMMENT       = 60
)

// SwiftgrammParser rules.
const (
	SwiftgrammParserRULE_s           = 0
	SwiftgrammParserRULE_block       = 1
	SwiftgrammParserRULE_sentence    = 2
	SwiftgrammParserRULE_declare_var = 3
	SwiftgrammParserRULE_print_bl    = 4
	SwiftgrammParserRULE_expression  = 5
	SwiftgrammParserRULE_datatype    = 6
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
		p.SetState(14)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(15)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammParserVAR || _la == SwiftgrammParserPRINT {
		{
			p.SetState(18)

			var _x = p.Sentence()

			localctx.(*BlockContext)._sentence = _x
		}
		localctx.(*BlockContext).instr = append(localctx.(*BlockContext).instr, localctx.(*BlockContext)._sentence)

		p.SetState(21)
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

	// Get_print_bl returns the _print_bl rule contexts.
	Get_print_bl() IPrint_blContext

	// Set_declare_var sets the _declare_var rule contexts.
	Set_declare_var(IDeclare_varContext)

	// Set_print_bl sets the _print_bl rule contexts.
	Set_print_bl(IPrint_blContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruction)

	// Getter signatures
	Declare_var() IDeclare_varContext
	Print_bl() IPrint_blContext

	// IsSentenceContext differentiates from other interfaces.
	IsSentenceContext()
}

type SentenceContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        abstract.Instruction
	_declare_var IDeclare_varContext
	_print_bl    IPrint_blContext
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

func (s *SentenceContext) Get_print_bl() IPrint_blContext { return s._print_bl }

func (s *SentenceContext) Set_declare_var(v IDeclare_varContext) { s._declare_var = v }

func (s *SentenceContext) Set_print_bl(v IPrint_blContext) { s._print_bl = v }

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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)

			var _x = p.Declare_var()

			localctx.(*SentenceContext)._declare_var = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_declare_var().GetInstr()

	case SwiftgrammParserPRINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)

			var _x = p.Print_bl()

			localctx.(*SentenceContext)._print_bl = _x
		}
		localctx.(*SentenceContext).instr = localctx.(*SentenceContext).Get_print_bl().GetInstr()

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
	p.EnterRule(localctx, 6, SwiftgrammParserRULE_declare_var)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)

			var _x = p.Datatype()

			localctx.(*Declare_varContext)._datatype = _x
		}
		{
			p.SetState(37)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)

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
			p.SetState(41)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)

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
			p.SetState(47)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)

			var _m = p.Match(SwiftgrammParserID)

			localctx.(*Declare_varContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)

			var _x = p.Datatype()

			localctx.(*Declare_varContext)._datatype = _x
		}
		{
			p.SetState(51)
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
	p.EnterRule(localctx, 8, SwiftgrammParserRULE_print_bl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(SwiftgrammParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(SwiftgrammParserOPEN_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)

		var _x = p.expression(0)

		localctx.(*Print_blContext)._expression = _x
	}
	{
		p.SetState(59)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SwiftgrammParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserNOT:
		{
			p.SetState(63)

			var _m = p.Match(SwiftgrammParserNOT)

			localctx.(*ExpressionContext).oper = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)

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
			p.SetState(67)
			p.Match(SwiftgrammParserOPEN_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)

			var _x = p.expression(0)

			localctx.(*ExpressionContext)._expression = _x
		}
		{
			p.SetState(69)
			p.Match(SwiftgrammParserCLOSE_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftgrammParserNUMBER:
		{
			p.SetState(71)

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
			p.SetState(73)

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
			p.SetState(75)

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
			p.SetState(77)

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
			p.SetState(79)

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
			p.SetState(81)

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
			p.SetState(83)

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
			p.SetState(85)
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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammParserRULE_expression)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(90)

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
					p.SetState(91)

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
				p.SetState(94)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(95)

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
					p.SetState(96)

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
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(100)

					var _m = p.Match(SwiftgrammParserMOD)

					localctx.(*ExpressionContext).oper = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(101)

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
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(105)

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
					p.SetState(106)

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
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(110)

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
					p.SetState(111)

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
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(115)

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
					p.SetState(116)

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
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(120)

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
					p.SetState(121)

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
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 12, SwiftgrammParserRULE_datatype)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
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
			p.SetState(131)
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
			p.SetState(133)
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
			p.SetState(135)
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
			p.SetState(137)
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
	case 5:
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
