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

// EnterCompareOperator is called when production compareOperator is entered.
func (s *BaseConditionListener) EnterCompareOperator(ctx *CompareOperatorContext) {}

// ExitCompareOperator is called when production compareOperator is exited.
func (s *BaseConditionListener) ExitCompareOperator(ctx *CompareOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConditionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConditionListener) ExitExpression(ctx *ExpressionContext) {}
