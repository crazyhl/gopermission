// Code generated from Condition.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 95, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 67, 10, 12, 3,
	12, 7, 12, 70, 10, 12, 12, 12, 14, 12, 73, 11, 12, 3, 13, 6, 13, 76, 10,
	13, 13, 13, 14, 13, 77, 3, 14, 6, 14, 81, 10, 14, 13, 14, 14, 14, 82, 3,
	14, 3, 14, 3, 15, 3, 15, 5, 15, 89, 10, 15, 3, 15, 5, 15, 92, 10, 15, 3,
	15, 3, 15, 2, 2, 16, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 3, 2, 5, 7, 2, 48, 48,
	50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 2,
	100, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 3, 31, 3, 2, 2, 2, 5,
	34, 3, 2, 2, 2, 7, 36, 3, 2, 2, 2, 9, 39, 3, 2, 2, 2, 11, 41, 3, 2, 2,
	2, 13, 44, 3, 2, 2, 2, 15, 47, 3, 2, 2, 2, 17, 49, 3, 2, 2, 2, 19, 51,
	3, 2, 2, 2, 21, 54, 3, 2, 2, 2, 23, 66, 3, 2, 2, 2, 25, 75, 3, 2, 2, 2,
	27, 80, 3, 2, 2, 2, 29, 91, 3, 2, 2, 2, 31, 32, 7, 63, 2, 2, 32, 33, 7,
	63, 2, 2, 33, 4, 3, 2, 2, 2, 34, 35, 7, 64, 2, 2, 35, 6, 3, 2, 2, 2, 36,
	37, 7, 64, 2, 2, 37, 38, 7, 63, 2, 2, 38, 8, 3, 2, 2, 2, 39, 40, 7, 62,
	2, 2, 40, 10, 3, 2, 2, 2, 41, 42, 7, 62, 2, 2, 42, 43, 7, 63, 2, 2, 43,
	12, 3, 2, 2, 2, 44, 45, 7, 107, 2, 2, 45, 46, 7, 112, 2, 2, 46, 14, 3,
	2, 2, 2, 47, 48, 7, 42, 2, 2, 48, 16, 3, 2, 2, 2, 49, 50, 7, 43, 2, 2,
	50, 18, 3, 2, 2, 2, 51, 52, 7, 126, 2, 2, 52, 53, 7, 126, 2, 2, 53, 20,
	3, 2, 2, 2, 54, 55, 7, 40, 2, 2, 55, 56, 7, 40, 2, 2, 56, 22, 3, 2, 2,
	2, 57, 58, 7, 111, 2, 2, 58, 59, 7, 113, 2, 2, 59, 60, 7, 102, 2, 2, 60,
	61, 7, 103, 2, 2, 61, 67, 7, 110, 2, 2, 62, 63, 7, 119, 2, 2, 63, 64, 7,
	117, 2, 2, 64, 65, 7, 103, 2, 2, 65, 67, 7, 116, 2, 2, 66, 57, 3, 2, 2,
	2, 66, 62, 3, 2, 2, 2, 67, 71, 3, 2, 2, 2, 68, 70, 9, 2, 2, 2, 69, 68,
	3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 24, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 76, 9, 3, 2, 2, 75, 74, 3,
	2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78,
	26, 3, 2, 2, 2, 79, 81, 9, 4, 2, 2, 80, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2,
	2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85,
	8, 14, 2, 2, 85, 28, 3, 2, 2, 2, 86, 88, 7, 15, 2, 2, 87, 89, 7, 12, 2,
	2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 92,
	7, 12, 2, 2, 91, 86, 3, 2, 2, 2, 91, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2,
	93, 94, 8, 15, 2, 2, 94, 30, 3, 2, 2, 2, 10, 2, 66, 69, 71, 77, 82, 88,
	91, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'=='", "'>'", "'>='", "'<'", "'<='", "'in'", "'('", "')'", "'||'",
	"'&&'",
}

var lexerSymbolicNames = []string{
	"", "EqualOP", "LargerOp", "LargerEqualOp", "LessOp", "LessEqualOp", "InOP",
	"LeftParen", "RightParen", "OrOP", "AndOP", "Paramater", "Number", "Whitespace",
	"Newline",
}

var lexerRuleNames = []string{
	"EqualOP", "LargerOp", "LargerEqualOp", "LessOp", "LessEqualOp", "InOP",
	"LeftParen", "RightParen", "OrOP", "AndOP", "Paramater", "Number", "Whitespace",
	"Newline",
}

type ConditionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewConditionLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ConditionLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewConditionLexer(input antlr.CharStream) *ConditionLexer {
	l := new(ConditionLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Condition.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ConditionLexer tokens.
const (
	ConditionLexerEqualOP       = 1
	ConditionLexerLargerOp      = 2
	ConditionLexerLargerEqualOp = 3
	ConditionLexerLessOp        = 4
	ConditionLexerLessEqualOp   = 5
	ConditionLexerInOP          = 6
	ConditionLexerLeftParen     = 7
	ConditionLexerRightParen    = 8
	ConditionLexerOrOP          = 9
	ConditionLexerAndOP         = 10
	ConditionLexerParamater     = 11
	ConditionLexerNumber        = 12
	ConditionLexerWhitespace    = 13
	ConditionLexerNewline       = 14
)
