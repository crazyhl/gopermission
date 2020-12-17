// Code generated from Condition.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Condition

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 40, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 27, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 35, 10, 4, 12,
	4, 14, 4, 38, 11, 4, 3, 4, 2, 3, 6, 5, 2, 4, 6, 2, 3, 4, 2, 3, 7, 10, 10,
	2, 40, 2, 8, 3, 2, 2, 2, 4, 11, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 9, 5,
	6, 4, 2, 9, 10, 7, 2, 2, 3, 10, 3, 3, 2, 2, 2, 11, 12, 9, 2, 2, 2, 12,
	5, 3, 2, 2, 2, 13, 14, 8, 4, 1, 2, 14, 15, 7, 11, 2, 2, 15, 16, 5, 6, 4,
	2, 16, 17, 7, 12, 2, 2, 17, 27, 3, 2, 2, 2, 18, 19, 7, 15, 2, 2, 19, 20,
	5, 4, 3, 2, 20, 21, 7, 15, 2, 2, 21, 27, 3, 2, 2, 2, 22, 23, 7, 15, 2,
	2, 23, 24, 5, 4, 3, 2, 24, 25, 7, 16, 2, 2, 25, 27, 3, 2, 2, 2, 26, 13,
	3, 2, 2, 2, 26, 18, 3, 2, 2, 2, 26, 22, 3, 2, 2, 2, 27, 36, 3, 2, 2, 2,
	28, 29, 12, 7, 2, 2, 29, 30, 7, 13, 2, 2, 30, 35, 5, 6, 4, 8, 31, 32, 12,
	6, 2, 2, 32, 33, 7, 14, 2, 2, 33, 35, 5, 6, 4, 7, 34, 28, 3, 2, 2, 2, 34,
	31, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2,
	2, 37, 7, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 5, 26, 34, 36,
}
var literalNames = []string{
	"", "'='", "'=='", "'>'", "'>='", "'<='", "", "", "'in'", "'('", "')'",
	"'||'", "'&&'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "EqualOP", "RelationalOP", "InOP", "LeftParen",
	"RightParen", "OrOP", "AndOP", "Paramater", "Number", "Whitespace", "Newline",
}

var ruleNames = []string{
	"start", "compareOperator", "expression",
}

type ConditionParser struct {
	*antlr.BaseParser
}

// NewConditionParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ConditionParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewConditionParser(input antlr.TokenStream) *ConditionParser {
	this := new(ConditionParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Condition.g4"

	return this
}

// ConditionParser tokens.
const (
	ConditionParserEOF          = antlr.TokenEOF
	ConditionParserT__0         = 1
	ConditionParserT__1         = 2
	ConditionParserT__2         = 3
	ConditionParserT__3         = 4
	ConditionParserT__4         = 5
	ConditionParserEqualOP      = 6
	ConditionParserRelationalOP = 7
	ConditionParserInOP         = 8
	ConditionParserLeftParen    = 9
	ConditionParserRightParen   = 10
	ConditionParserOrOP         = 11
	ConditionParserAndOP        = 12
	ConditionParserParamater    = 13
	ConditionParserNumber       = 14
	ConditionParserWhitespace   = 15
	ConditionParserNewline      = 16
)

// ConditionParser rules.
const (
	ConditionParserRULE_start           = 0
	ConditionParserRULE_compareOperator = 1
	ConditionParserRULE_expression      = 2
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConditionParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConditionParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConditionParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ConditionParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConditionParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.expression(0)
	}
	{
		p.SetState(7)
		p.Match(ConditionParserEOF)
	}

	return localctx
}

// ICompareOperatorContext is an interface to support dynamic dispatch.
type ICompareOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareOperatorContext differentiates from other interfaces.
	IsCompareOperatorContext()
}

type CompareOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareOperatorContext() *CompareOperatorContext {
	var p = new(CompareOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConditionParserRULE_compareOperator
	return p
}

func (*CompareOperatorContext) IsCompareOperatorContext() {}

func NewCompareOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareOperatorContext {
	var p = new(CompareOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConditionParserRULE_compareOperator

	return p
}

func (s *CompareOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareOperatorContext) InOP() antlr.TerminalNode {
	return s.GetToken(ConditionParserInOP, 0)
}

func (s *CompareOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.EnterCompareOperator(s)
	}
}

func (s *CompareOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.ExitCompareOperator(s)
	}
}

func (p *ConditionParser) CompareOperator() (localctx ICompareOperatorContext) {
	localctx = NewCompareOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConditionParserRULE_compareOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(9)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConditionParserT__0)|(1<<ConditionParserT__1)|(1<<ConditionParserT__2)|(1<<ConditionParserT__3)|(1<<ConditionParserT__4)|(1<<ConditionParserInOP))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConditionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConditionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(ConditionParserLeftParen, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(ConditionParserRightParen, 0)
}

func (s *ExpressionContext) AllParamater() []antlr.TerminalNode {
	return s.GetTokens(ConditionParserParamater)
}

func (s *ExpressionContext) Paramater(i int) antlr.TerminalNode {
	return s.GetToken(ConditionParserParamater, i)
}

func (s *ExpressionContext) CompareOperator() ICompareOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareOperatorContext)
}

func (s *ExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(ConditionParserNumber, 0)
}

func (s *ExpressionContext) OrOP() antlr.TerminalNode {
	return s.GetToken(ConditionParserOrOP, 0)
}

func (s *ExpressionContext) AndOP() antlr.TerminalNode {
	return s.GetToken(ConditionParserAndOP, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConditionParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ConditionParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ConditionParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(12)
			p.Match(ConditionParserLeftParen)
		}
		{
			p.SetState(13)
			p.expression(0)
		}
		{
			p.SetState(14)
			p.Match(ConditionParserRightParen)
		}

	case 2:
		{
			p.SetState(16)
			p.Match(ConditionParserParamater)
		}
		{
			p.SetState(17)
			p.CompareOperator()
		}
		{
			p.SetState(18)
			p.Match(ConditionParserParamater)
		}

	case 3:
		{
			p.SetState(20)
			p.Match(ConditionParserParamater)
		}
		{
			p.SetState(21)
			p.CompareOperator()
		}
		{
			p.SetState(22)
			p.Match(ConditionParserNumber)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConditionParserRULE_expression)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(27)
					p.Match(ConditionParserOrOP)
				}
				{
					p.SetState(28)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConditionParserRULE_expression)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(30)
					p.Match(ConditionParserAndOP)
				}
				{
					p.SetState(31)
					p.expression(5)
				}

			}

		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ConditionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ConditionParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
