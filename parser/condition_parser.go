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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 33, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 20, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 7, 3, 28, 10, 3, 12, 3, 14, 3, 31, 11, 3, 3, 3, 2, 3, 4, 4, 2,
	4, 2, 3, 3, 2, 3, 8, 2, 34, 2, 6, 3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 7,
	5, 4, 3, 2, 7, 3, 3, 2, 2, 2, 8, 9, 8, 3, 1, 2, 9, 10, 7, 9, 2, 2, 10,
	11, 5, 4, 3, 2, 11, 12, 7, 10, 2, 2, 12, 20, 3, 2, 2, 2, 13, 14, 7, 13,
	2, 2, 14, 15, 9, 2, 2, 2, 15, 20, 7, 14, 2, 2, 16, 17, 7, 13, 2, 2, 17,
	18, 9, 2, 2, 2, 18, 20, 7, 13, 2, 2, 19, 8, 3, 2, 2, 2, 19, 13, 3, 2, 2,
	2, 19, 16, 3, 2, 2, 2, 20, 29, 3, 2, 2, 2, 21, 22, 12, 4, 2, 2, 22, 23,
	7, 12, 2, 2, 23, 28, 5, 4, 3, 5, 24, 25, 12, 3, 2, 2, 25, 26, 7, 11, 2,
	2, 26, 28, 5, 4, 3, 4, 27, 21, 3, 2, 2, 2, 27, 24, 3, 2, 2, 2, 28, 31,
	3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 5, 3, 2, 2, 2,
	31, 29, 3, 2, 2, 2, 5, 19, 27, 29,
}
var literalNames = []string{
	"", "'=='", "'>'", "'>='", "'<'", "'<='", "'in'", "'('", "')'", "'||'",
	"'&&'",
}
var symbolicNames = []string{
	"", "EqualOP", "LargerOp", "LargerEqualOp", "LessOp", "LessEqualOp", "InOP",
	"LeftParen", "RightParen", "OrOP", "AndOP", "Paramater", "Number", "Whitespace",
	"Newline",
}

var ruleNames = []string{
	"start", "expression",
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
	ConditionParserEOF           = antlr.TokenEOF
	ConditionParserEqualOP       = 1
	ConditionParserLargerOp      = 2
	ConditionParserLargerEqualOp = 3
	ConditionParserLessOp        = 4
	ConditionParserLessEqualOp   = 5
	ConditionParserInOP          = 6
	ConditionParserLeftParen     = 7
	ConditionParserRightParen    = 8
	ConditionParserOrOP          = 9
	ConditionParserAndOP         = 10
	ConditionParserParamater     = 11
	ConditionParserNumber        = 12
	ConditionParserWhitespace    = 13
	ConditionParserNewline       = 14
)

// ConditionParser rules.
const (
	ConditionParserRULE_start      = 0
	ConditionParserRULE_expression = 1
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
		p.SetState(4)
		p.expression(0)
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

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareContext struct {
	*ExpressionContext
	left  antlr.Token
	op    antlr.Token
	right antlr.Token
}

func NewCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareContext {
	var p = new(CompareContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompareContext) GetLeft() antlr.Token { return s.left }

func (s *CompareContext) GetOp() antlr.Token { return s.op }

func (s *CompareContext) GetRight() antlr.Token { return s.right }

func (s *CompareContext) SetLeft(v antlr.Token) { s.left = v }

func (s *CompareContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareContext) SetRight(v antlr.Token) { s.right = v }

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) AllParamater() []antlr.TerminalNode {
	return s.GetTokens(ConditionParserParamater)
}

func (s *CompareContext) Paramater(i int) antlr.TerminalNode {
	return s.GetToken(ConditionParserParamater, i)
}

func (s *CompareContext) Number() antlr.TerminalNode {
	return s.GetToken(ConditionParserNumber, 0)
}

func (s *CompareContext) EqualOP() antlr.TerminalNode {
	return s.GetToken(ConditionParserEqualOP, 0)
}

func (s *CompareContext) LargerOp() antlr.TerminalNode {
	return s.GetToken(ConditionParserLargerOp, 0)
}

func (s *CompareContext) LargerEqualOp() antlr.TerminalNode {
	return s.GetToken(ConditionParserLargerEqualOp, 0)
}

func (s *CompareContext) LessOp() antlr.TerminalNode {
	return s.GetToken(ConditionParserLessOp, 0)
}

func (s *CompareContext) LessEqualOp() antlr.TerminalNode {
	return s.GetToken(ConditionParserLessEqualOp, 0)
}

func (s *CompareContext) InOP() antlr.TerminalNode {
	return s.GetToken(ConditionParserInOP, 0)
}

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.ExitCompare(s)
	}
}

type OrCompareContext struct {
	*ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewOrCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrCompareContext {
	var p = new(OrCompareContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrCompareContext) GetLeft() IExpressionContext { return s.left }

func (s *OrCompareContext) GetRight() IExpressionContext { return s.right }

func (s *OrCompareContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *OrCompareContext) SetRight(v IExpressionContext) { s.right = v }

func (s *OrCompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrCompareContext) OrOP() antlr.TerminalNode {
	return s.GetToken(ConditionParserOrOP, 0)
}

func (s *OrCompareContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrCompareContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrCompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.EnterOrCompare(s)
	}
}

func (s *OrCompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.ExitOrCompare(s)
	}
}

type ParensExpressionContext struct {
	*ExpressionContext
}

func NewParensExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExpressionContext {
	var p = new(ParensExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParensExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(ConditionParserLeftParen, 0)
}

func (s *ParensExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParensExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(ConditionParserRightParen, 0)
}

func (s *ParensExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.EnterParensExpression(s)
	}
}

func (s *ParensExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.ExitParensExpression(s)
	}
}

type AndCompareContext struct {
	*ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewAndCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndCompareContext {
	var p = new(AndCompareContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndCompareContext) GetLeft() IExpressionContext { return s.left }

func (s *AndCompareContext) GetRight() IExpressionContext { return s.right }

func (s *AndCompareContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *AndCompareContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AndCompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndCompareContext) AndOP() antlr.TerminalNode {
	return s.GetToken(ConditionParserAndOP, 0)
}

func (s *AndCompareContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndCompareContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndCompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.EnterAndCompare(s)
	}
}

func (s *AndCompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConditionListener); ok {
		listenerT.ExitAndCompare(s)
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
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ConditionParserRULE_expression, _p)
	var _la int

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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParensExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(7)
			p.Match(ConditionParserLeftParen)
		}
		{
			p.SetState(8)
			p.expression(0)
		}
		{
			p.SetState(9)
			p.Match(ConditionParserRightParen)
		}

	case 2:
		localctx = NewCompareContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)

			var _m = p.Match(ConditionParserParamater)

			localctx.(*CompareContext).left = _m
		}
		{
			p.SetState(12)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConditionParserEqualOP)|(1<<ConditionParserLargerOp)|(1<<ConditionParserLargerEqualOp)|(1<<ConditionParserLessOp)|(1<<ConditionParserLessEqualOp)|(1<<ConditionParserInOP))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(13)

			var _m = p.Match(ConditionParserNumber)

			localctx.(*CompareContext).right = _m
		}

	case 3:
		localctx = NewCompareContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)

			var _m = p.Match(ConditionParserParamater)

			localctx.(*CompareContext).left = _m
		}
		{
			p.SetState(15)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConditionParserEqualOP)|(1<<ConditionParserLargerOp)|(1<<ConditionParserLargerEqualOp)|(1<<ConditionParserLessOp)|(1<<ConditionParserLessEqualOp)|(1<<ConditionParserInOP))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(16)

			var _m = p.Match(ConditionParserParamater)

			localctx.(*CompareContext).right = _m
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(25)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndCompareContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AndCompareContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ConditionParserRULE_expression)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(20)
					p.Match(ConditionParserAndOP)
				}
				{
					p.SetState(21)

					var _x = p.expression(3)

					localctx.(*AndCompareContext).right = _x
				}

			case 2:
				localctx = NewOrCompareContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OrCompareContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ConditionParserRULE_expression)
				p.SetState(22)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(23)
					p.Match(ConditionParserOrOP)
				}
				{
					p.SetState(24)

					var _x = p.expression(2)

					localctx.(*OrCompareContext).right = _x
				}

			}

		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ConditionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
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
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
