// Code generated from Condition.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Condition

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConditionListener is a complete listener for a parse tree produced by ConditionParser.
type BaseConditionListener struct{}

var _ ConditionListener = &BaseConditionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConditionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConditionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConditionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConditionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseConditionListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseConditionListener) ExitStart(ctx *StartContext) {}

// EnterCompare is called when production compare is entered.
func (s *BaseConditionListener) EnterCompare(ctx *CompareContext) {}

// ExitCompare is called when production compare is exited.
func (s *BaseConditionListener) ExitCompare(ctx *CompareContext) {}

// EnterOrCompare is called when production orCompare is entered.
func (s *BaseConditionListener) EnterOrCompare(ctx *OrCompareContext) {}

// ExitOrCompare is called when production orCompare is exited.
func (s *BaseConditionListener) ExitOrCompare(ctx *OrCompareContext) {}

// EnterParensExpression is called when production parensExpression is entered.
func (s *BaseConditionListener) EnterParensExpression(ctx *ParensExpressionContext) {}

// ExitParensExpression is called when production parensExpression is exited.
func (s *BaseConditionListener) ExitParensExpression(ctx *ParensExpressionContext) {}

// EnterAndCompare is called when production andCompare is entered.
func (s *BaseConditionListener) EnterAndCompare(ctx *AndCompareContext) {}

// ExitAndCompare is called when production andCompare is exited.
func (s *BaseConditionListener) ExitAndCompare(ctx *AndCompareContext) {}
