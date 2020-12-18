// Code generated from Condition.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Condition

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConditionListener is a complete listener for a parse tree produced by ConditionParser.
type ConditionListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterCompare is called when entering the compare production.
	EnterCompare(c *CompareContext)

	// EnterOrCompare is called when entering the orCompare production.
	EnterOrCompare(c *OrCompareContext)

	// EnterParensExpression is called when entering the parensExpression production.
	EnterParensExpression(c *ParensExpressionContext)

	// EnterAndCompare is called when entering the andCompare production.
	EnterAndCompare(c *AndCompareContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitCompare is called when exiting the compare production.
	ExitCompare(c *CompareContext)

	// ExitOrCompare is called when exiting the orCompare production.
	ExitOrCompare(c *OrCompareContext)

	// ExitParensExpression is called when exiting the parensExpression production.
	ExitParensExpression(c *ParensExpressionContext)

	// ExitAndCompare is called when exiting the andCompare production.
	ExitAndCompare(c *AndCompareContext)
}
