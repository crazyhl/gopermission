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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 112,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 5, 7, 52, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 5, 8, 60, 10, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 5, 14, 84, 10, 14, 3, 14, 7, 14, 87, 10, 14, 12,
	14, 14, 14, 90, 11, 14, 3, 15, 6, 15, 93, 10, 15, 13, 15, 14, 15, 94, 3,
	16, 6, 16, 98, 10, 16, 13, 16, 14, 16, 99, 3, 16, 3, 16, 3, 17, 3, 17,
	5, 17, 106, 10, 17, 3, 17, 5, 17, 109, 10, 17, 3, 17, 3, 17, 2, 2, 18,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 3, 2, 5, 7, 2, 48, 48, 50,
	59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 2, 121,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 37, 3, 2, 2, 2, 7, 40, 3, 2, 2, 2, 9, 42,
	3, 2, 2, 2, 11, 45, 3, 2, 2, 2, 13, 51, 3, 2, 2, 2, 15, 59, 3, 2, 2, 2,
	17, 61, 3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 66, 3, 2, 2, 2, 23, 68, 3,
	2, 2, 2, 25, 71, 3, 2, 2, 2, 27, 83, 3, 2, 2, 2, 29, 92, 3, 2, 2, 2, 31,
	97, 3, 2, 2, 2, 33, 108, 3, 2, 2, 2, 35, 36, 7, 63, 2, 2, 36, 4, 3, 2,
	2, 2, 37, 38, 7, 63, 2, 2, 38, 39, 7, 63, 2, 2, 39, 6, 3, 2, 2, 2, 40,
	41, 7, 64, 2, 2, 41, 8, 3, 2, 2, 2, 42, 43, 7, 64, 2, 2, 43, 44, 7, 63,
	2, 2, 44, 10, 3, 2, 2, 2, 45, 46, 7, 62, 2, 2, 46, 47, 7, 63, 2, 2, 47,
	12, 3, 2, 2, 2, 48, 49, 7, 63, 2, 2, 49, 52, 7, 63, 2, 2, 50, 52, 7, 63,
	2, 2, 51, 48, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 14, 3, 2, 2, 2, 53, 60,
	7, 64, 2, 2, 54, 55, 7, 64, 2, 2, 55, 60, 7, 63, 2, 2, 56, 60, 7, 62, 2,
	2, 57, 58, 7, 62, 2, 2, 58, 60, 7, 63, 2, 2, 59, 53, 3, 2, 2, 2, 59, 54,
	3, 2, 2, 2, 59, 56, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 16, 3, 2, 2, 2,
	61, 62, 7, 107, 2, 2, 62, 63, 7, 112, 2, 2, 63, 18, 3, 2, 2, 2, 64, 65,
	7, 42, 2, 2, 65, 20, 3, 2, 2, 2, 66, 67, 7, 43, 2, 2, 67, 22, 3, 2, 2,
	2, 68, 69, 7, 126, 2, 2, 69, 70, 7, 126, 2, 2, 70, 24, 3, 2, 2, 2, 71,
	72, 7, 40, 2, 2, 72, 73, 7, 40, 2, 2, 73, 26, 3, 2, 2, 2, 74, 75, 7, 111,
	2, 2, 75, 76, 7, 113, 2, 2, 76, 77, 7, 102, 2, 2, 77, 78, 7, 103, 2, 2,
	78, 84, 7, 110, 2, 2, 79, 80, 7, 119, 2, 2, 80, 81, 7, 117, 2, 2, 81, 82,
	7, 103, 2, 2, 82, 84, 7, 116, 2, 2, 83, 74, 3, 2, 2, 2, 83, 79, 3, 2, 2,
	2, 84, 88, 3, 2, 2, 2, 85, 87, 9, 2, 2, 2, 86, 85, 3, 2, 2, 2, 87, 90,
	3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 28, 3, 2, 2, 2,
	90, 88, 3, 2, 2, 2, 91, 93, 9, 3, 2, 2, 92, 91, 3, 2, 2, 2, 93, 94, 3,
	2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 30, 3, 2, 2, 2, 96,
	98, 9, 4, 2, 2, 97, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2,
	2, 99, 100, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 8, 16, 2, 2, 102,
	32, 3, 2, 2, 2, 103, 105, 7, 15, 2, 2, 104, 106, 7, 12, 2, 2, 105, 104,
	3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 109, 7, 12,
	2, 2, 108, 103, 3, 2, 2, 2, 108, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2,
	110, 111, 8, 17, 2, 2, 111, 34, 3, 2, 2, 2, 12, 2, 51, 59, 83, 86, 88,
	94, 99, 105, 108, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'=='", "'>'", "'>='", "'<='", "", "", "'in'", "'('", "')'",
	"'||'", "'&&'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "EqualOP", "RelationalOP", "InOP", "LeftParen",
	"RightParen", "OrOP", "AndOP", "Paramater", "Number", "Whitespace", "Newline",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "EqualOP", "RelationalOP", "InOP",
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
	ConditionLexerT__0         = 1
	ConditionLexerT__1         = 2
	ConditionLexerT__2         = 3
	ConditionLexerT__3         = 4
	ConditionLexerT__4         = 5
	ConditionLexerEqualOP      = 6
	ConditionLexerRelationalOP = 7
	ConditionLexerInOP         = 8
	ConditionLexerLeftParen    = 9
	ConditionLexerRightParen   = 10
	ConditionLexerOrOP         = 11
	ConditionLexerAndOP        = 12
	ConditionLexerParamater    = 13
	ConditionLexerNumber       = 14
	ConditionLexerWhitespace   = 15
	ConditionLexerNewline      = 16
)
