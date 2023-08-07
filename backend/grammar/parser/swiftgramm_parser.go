// Code generated from Swiftgramm.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgramm
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

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
		"s", "block", "sentences", "sentence", "declare_var", "declare_constant",
		"assign_var", "expression", "datatype",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 85, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		1, 5, 1, 23, 8, 1, 10, 1, 12, 1, 26, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		32, 8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 37, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 56, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 75, 8, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 0, 2, 1, 0, 29, 32, 1, 0, 1, 4, 83, 0, 18, 1, 0, 0, 0, 2, 24,
		1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 36, 1, 0, 0, 0, 8, 55, 1, 0, 0, 0, 10,
		74, 1, 0, 0, 0, 12, 76, 1, 0, 0, 0, 14, 80, 1, 0, 0, 0, 16, 82, 1, 0, 0,
		0, 18, 19, 3, 2, 1, 0, 19, 20, 5, 0, 0, 1, 20, 1, 1, 0, 0, 0, 21, 23, 3,
		4, 2, 0, 22, 21, 1, 0, 0, 0, 23, 26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24,
		25, 1, 0, 0, 0, 25, 3, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 27, 28, 3, 6, 3,
		0, 28, 29, 3, 4, 2, 0, 29, 32, 1, 0, 0, 0, 30, 32, 3, 6, 3, 0, 31, 27,
		1, 0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 5, 1, 0, 0, 0, 33, 37, 3, 8, 4, 0,
		34, 37, 3, 10, 5, 0, 35, 37, 3, 12, 6, 0, 36, 33, 1, 0, 0, 0, 36, 34, 1,
		0, 0, 0, 36, 35, 1, 0, 0, 0, 37, 7, 1, 0, 0, 0, 38, 39, 5, 9, 0, 0, 39,
		40, 5, 31, 0, 0, 40, 41, 5, 50, 0, 0, 41, 42, 3, 16, 8, 0, 42, 43, 5, 47,
		0, 0, 43, 44, 3, 14, 7, 0, 44, 56, 1, 0, 0, 0, 45, 46, 5, 9, 0, 0, 46,
		47, 5, 31, 0, 0, 47, 48, 5, 47, 0, 0, 48, 56, 3, 14, 7, 0, 49, 50, 5, 9,
		0, 0, 50, 51, 5, 31, 0, 0, 51, 52, 5, 50, 0, 0, 52, 53, 3, 16, 8, 0, 53,
		54, 5, 37, 0, 0, 54, 56, 1, 0, 0, 0, 55, 38, 1, 0, 0, 0, 55, 45, 1, 0,
		0, 0, 55, 49, 1, 0, 0, 0, 56, 9, 1, 0, 0, 0, 57, 58, 5, 10, 0, 0, 58, 59,
		5, 31, 0, 0, 59, 60, 5, 50, 0, 0, 60, 61, 3, 16, 8, 0, 61, 62, 5, 47, 0,
		0, 62, 63, 3, 14, 7, 0, 63, 75, 1, 0, 0, 0, 64, 65, 5, 10, 0, 0, 65, 66,
		5, 31, 0, 0, 66, 67, 5, 47, 0, 0, 67, 75, 3, 14, 7, 0, 68, 69, 5, 10, 0,
		0, 69, 70, 5, 31, 0, 0, 70, 71, 5, 50, 0, 0, 71, 72, 3, 16, 8, 0, 72, 73,
		5, 37, 0, 0, 73, 75, 1, 0, 0, 0, 74, 57, 1, 0, 0, 0, 74, 64, 1, 0, 0, 0,
		74, 68, 1, 0, 0, 0, 75, 11, 1, 0, 0, 0, 76, 77, 5, 31, 0, 0, 77, 78, 5,
		47, 0, 0, 78, 79, 3, 14, 7, 0, 79, 13, 1, 0, 0, 0, 80, 81, 7, 0, 0, 0,
		81, 15, 1, 0, 0, 0, 82, 83, 7, 1, 0, 0, 83, 17, 1, 0, 0, 0, 5, 24, 31,
		36, 55, 74,
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
	SwiftgrammParserQUESTION_MARK      = 37
	SwiftgrammParserOR                 = 38
	SwiftgrammParserAND                = 39
	SwiftgrammParserNOT                = 40
	SwiftgrammParserEQUAL              = 41
	SwiftgrammParserNOT_EQUAL          = 42
	SwiftgrammParserLESS_THAN          = 43
	SwiftgrammParserLESS_THAN_EQUAL    = 44
	SwiftgrammParserGREATER_THAN       = 45
	SwiftgrammParserGREATER_THAN_EQUAL = 46
	SwiftgrammParserASSIGN             = 47
	SwiftgrammParserDOT                = 48
	SwiftgrammParserCOMMA              = 49
	SwiftgrammParserCOLON              = 50
	SwiftgrammParserSEMICOLON          = 51
	SwiftgrammParserOPEN_PARENTHESIS   = 52
	SwiftgrammParserCLOSE_PARENTHESIS  = 53
	SwiftgrammParserOPEN_kEY           = 54
	SwiftgrammParserCLOSE_kEY          = 55
	SwiftgrammParserWHITESPACE         = 56
	SwiftgrammParserCOMMENT            = 57
	SwiftgrammParserLINE_COMMENT       = 58
)

// SwiftgrammParser rules.
const (
	SwiftgrammParserRULE_s                = 0
	SwiftgrammParserRULE_block            = 1
	SwiftgrammParserRULE_sentences        = 2
	SwiftgrammParserRULE_sentence         = 3
	SwiftgrammParserRULE_declare_var      = 4
	SwiftgrammParserRULE_declare_constant = 5
	SwiftgrammParserRULE_assign_var       = 6
	SwiftgrammParserRULE_expression       = 7
	SwiftgrammParserRULE_datatype         = 8
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
		p.SetState(18)
		p.Block()
	}
	{
		p.SetState(19)
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
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147485184) != 0 {
		{
			p.SetState(21)
			p.Sentences()
		}

		p.SetState(26)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Sentence()
		}
		{
			p.SetState(28)
			p.Sentences()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Declare_var()
		}

	case SwiftgrammParserLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Declare_constant()
		}

	case SwiftgrammParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
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
	p.EnterRule(localctx, 8, SwiftgrammParserRULE_declare_var)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Datatype()
		}
		{
			p.SetState(42)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.Match(SwiftgrammParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Datatype()
		}
		{
			p.SetState(53)
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
	p.EnterRule(localctx, 10, SwiftgrammParserRULE_declare_constant)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Datatype()
		}
		{
			p.SetState(61)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(SwiftgrammParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Match(SwiftgrammParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(SwiftgrammParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(SwiftgrammParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Datatype()
		}
		{
			p.SetState(72)
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
	p.EnterRule(localctx, 12, SwiftgrammParserRULE_assign_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(SwiftgrammParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(SwiftgrammParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Expression()
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

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	CHARACTER_LITERAL() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftgrammParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8053063680) != 0) {
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
	p.EnterRule(localctx, 16, SwiftgrammParserRULE_datatype)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
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
