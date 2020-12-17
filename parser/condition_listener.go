// Code generated from Condition.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Condition

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConditionListener is a complete listener for a parse tree produced by ConditionParser.
type ConditionListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterCompareOperator is called when entering the compareOperator production.
	EnterCompareOperator(c *CompareOperatorContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitCompareOperator is called when exiting the compareOperator production.
	ExitCompareOperator(c *CompareOperatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
