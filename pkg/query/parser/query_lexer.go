// Code generated from /Users/i322053/goworkspace/src/github.com/Peripli/service-manager/pkg/query/Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 86, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 37, 10, 
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 55, 10, 6, 3, 7, 6, 7, 58, 10, 7, 13, 
	7, 14, 7, 59, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 68, 10, 8, 12, 
	8, 14, 8, 71, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 78, 10, 9, 3, 
	10, 6, 10, 81, 10, 10, 13, 10, 14, 10, 82, 3, 10, 3, 10, 2, 2, 11, 3, 3, 
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 5, 8, 2, 47, 
	47, 49, 59, 67, 92, 94, 94, 97, 97, 99, 124, 4, 2, 41, 41, 94, 94, 5, 2, 
	11, 12, 15, 15, 34, 34, 2, 96, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 
	5, 25, 3, 2, 2, 2, 7, 27, 3, 2, 2, 2, 9, 36, 3, 2, 2, 2, 11, 54, 3, 2, 
	2, 2, 13, 57, 3, 2, 2, 2, 15, 61, 3, 2, 2, 2, 17, 77, 3, 2, 2, 2, 19, 80, 
	3, 2, 2, 2, 21, 22, 7, 99, 2, 2, 22, 23, 7, 112, 2, 2, 23, 24, 7, 102, 
	2, 2, 24, 4, 3, 2, 2, 2, 25, 26, 7, 93, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 
	7, 95, 2, 2, 28, 8, 3, 2, 2, 2, 29, 30, 7, 107, 2, 2, 30, 37, 7, 112, 2, 
	2, 31, 32, 7, 112, 2, 2, 32, 33, 7, 113, 2, 2, 33, 34, 7, 118, 2, 2, 34, 
	35, 7, 107, 2, 2, 35, 37, 7, 112, 2, 2, 36, 29, 3, 2, 2, 2, 36, 31, 3, 
	2, 2, 2, 37, 10, 3, 2, 2, 2, 38, 39, 7, 103, 2, 2, 39, 55, 7, 115, 2, 2, 
	40, 41, 7, 112, 2, 2, 41, 42, 7, 103, 2, 2, 42, 55, 7, 115, 2, 2, 43, 44, 
	7, 103, 2, 2, 44, 45, 7, 115, 2, 2, 45, 46, 7, 113, 2, 2, 46, 47, 7, 116, 
	2, 2, 47, 48, 7, 112, 2, 2, 48, 49, 7, 107, 2, 2, 49, 55, 7, 110, 2, 2, 
	50, 51, 7, 105, 2, 2, 51, 55, 7, 118, 2, 2, 52, 53, 7, 110, 2, 2, 53, 55, 
	7, 118, 2, 2, 54, 38, 3, 2, 2, 2, 54, 40, 3, 2, 2, 2, 54, 43, 3, 2, 2, 
	2, 54, 50, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 12, 3, 2, 2, 2, 56, 58, 
	9, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 
	59, 60, 3, 2, 2, 2, 60, 14, 3, 2, 2, 2, 61, 69, 7, 41, 2, 2, 62, 63, 7, 
	94, 2, 2, 63, 68, 11, 2, 2, 2, 64, 65, 7, 41, 2, 2, 65, 68, 7, 41, 2, 2, 
	66, 68, 10, 3, 2, 2, 67, 62, 3, 2, 2, 2, 67, 64, 3, 2, 2, 2, 67, 66, 3, 
	2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 
	72, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 7, 41, 2, 2, 73, 16, 3, 2, 
	2, 2, 74, 78, 7, 46, 2, 2, 75, 76, 7, 46, 2, 2, 76, 78, 7, 34, 2, 2, 77, 
	74, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 18, 3, 2, 2, 2, 79, 81, 9, 4, 2, 
	2, 80, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 
	3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 8, 10, 2, 2, 85, 20, 3, 2, 2, 2, 
	10, 2, 36, 54, 59, 67, 69, 77, 82, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'and'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "MultiOp", "UniOp", "Key", "Value", "ValueSeparator", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "MultiOp", "UniOp", "Key", "Value", "ValueSeparator", 
	"WS",
}

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewQueryLexer(input antlr.CharStream) *QueryLexer {

	l := new(QueryLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerT__0 = 1
	QueryLexerT__1 = 2
	QueryLexerT__2 = 3
	QueryLexerMultiOp = 4
	QueryLexerUniOp = 5
	QueryLexerKey = 6
	QueryLexerValue = 7
	QueryLexerValueSeparator = 8
	QueryLexerWS = 9
)
