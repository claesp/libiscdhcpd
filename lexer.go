package libiscdhcpd

import (
	"errors"
	"fmt"
)

type SpanType int

func (s SpanType) String() string {
	switch s {
	case SpanTypeUnknown:
		return "unknown"
	case SpanTypeWord:
		return "word"
	case SpanTypeTab:
		return "tab"
	case SpanTypeSpace:
		return "space"
	case SpanTypeComment:
		return "comment"
	case SpanTypeNewline:
		return "newline"
	case SpanTypeEquals:
		return "equals"
	case SpanTypeSemicolon:
		return "semicolon"
	case SpanTypeOpenBracket:
		return "openbracket"
	case SpanTypeCloseBracket:
		return "closebracket"
	case SpanTypeOpenParen:
		return "openparen"
	case SpanTypeCloseParen:
		return "closeparen"
	case SpanTypeOpenCurly:
		return "opencurly"
	case SpanTypeCloseCurly:
		return "closecurly"
	case SpanTypeColon:
		return "colon"
	case SpanTypeComma:
		return "comma"
	case SpanTypeDoubleQuote:
		return "doublequote"
	case SpanTypeQuote:
		return "quote"
	default:
		return "unset"
	}
}

const (
	SpanTypeUnknown SpanType = iota
	SpanTypeWord
	SpanTypeTab
	SpanTypeSpace
	SpanTypeComment
	SpanTypeNewline
	SpanTypeEquals
	SpanTypeSemicolon
	SpanTypeOpenBracket
	SpanTypeCloseBracket
	SpanTypeOpenParen
	SpanTypeCloseParen
	SpanTypeOpenCurly
	SpanTypeCloseCurly
	SpanTypeColon
	SpanTypeComma
	SpanTypeDoubleQuote
	SpanTypeQuote
)

type LexerSpan struct {
	Start int      `json:"start"`
	Stop  int      `json:"stop"`
	Value string   `json:"value"`
	Type  SpanType `json:"type"`
}

func (s LexerSpan) String() string {
	return fmt.Sprintf("%d:%d \"%s\" [%s]", s.Start, s.Stop, s.Value, s.Type)
}

func ClassifyCharacter(ch string) SpanType {
	switch ch {
	case "a":
		fallthrough
	case "A":
		fallthrough
	case "b":
		fallthrough
	case "B":
		fallthrough
	case "c":
		fallthrough
	case "C":
		fallthrough
	case "d":
		fallthrough
	case "D":
		fallthrough
	case "e":
		fallthrough
	case "E":
		fallthrough
	case "f":
		fallthrough
	case "F":
		fallthrough
	case "g":
		fallthrough
	case "G":
		fallthrough
	case "h":
		fallthrough
	case "H":
		fallthrough
	case "i":
		fallthrough
	case "I":
		fallthrough
	case "j":
		fallthrough
	case "J":
		fallthrough
	case "k":
		fallthrough
	case "K":
		fallthrough
	case "l":
		fallthrough
	case "L":
		fallthrough
	case "m":
		fallthrough
	case "M":
		fallthrough
	case "n":
		fallthrough
	case "N":
		fallthrough
	case "o":
		fallthrough
	case "O":
		fallthrough
	case "p":
		fallthrough
	case "P":
		fallthrough
	case "q":
		fallthrough
	case "Q":
		fallthrough
	case "r":
		fallthrough
	case "R":
		fallthrough
	case "s":
		fallthrough
	case "S":
		fallthrough
	case "t":
		fallthrough
	case "T":
		fallthrough
	case "u":
		fallthrough
	case "U":
		fallthrough
	case "v":
		fallthrough
	case "V":
		fallthrough
	case "w":
		fallthrough
	case "W":
		fallthrough
	case "x":
		fallthrough
	case "X":
		fallthrough
	case "y":
		fallthrough
	case "Y":
		fallthrough
	case "z":
		fallthrough
	case "Z":
		fallthrough
	case ".":
		fallthrough
	case "-":
		fallthrough
	case "/":
		fallthrough
	case "_":
		fallthrough
	case "!":
		fallthrough
	case "1":
		fallthrough
	case "2":
		fallthrough
	case "3":
		fallthrough
	case "4":
		fallthrough
	case "5":
		fallthrough
	case "6":
		fallthrough
	case "7":
		fallthrough
	case "8":
		fallthrough
	case "9":
		fallthrough
	case "0":
		return SpanTypeWord
	case ",":
		return SpanTypeComma
	case "#":
		return SpanTypeComment
	case "=":
		return SpanTypeEquals
	case ";":
		return SpanTypeSemicolon
	case "\t":
		return SpanTypeTab
	case " ":
		return SpanTypeSpace
	case "\r":
		return SpanTypeNewline
	case "\n":
		return SpanTypeNewline
	case "(":
		return SpanTypeOpenParen
	case ")":
		return SpanTypeCloseParen
	case "[":
		return SpanTypeOpenBracket
	case "]":
		return SpanTypeCloseBracket
	case "}":
		return SpanTypeCloseCurly
	case "{":
		return SpanTypeOpenCurly
	case ":":
		return SpanTypeColon
	case "\"":
		return SpanTypeDoubleQuote
	case "'":
		return SpanTypeQuote
	default:
		return SpanTypeUnknown
	}
}

func Lex(cfg DhcpdDocument) ([]LexerSpan, error) {
	lexSpans := make([]LexerSpan, 0)

	if len(cfg.Filedata) == 0 {
		return lexSpans, errors.New("missing filedate, you need to load a configuration")
	}

	currLexSpan := LexerSpan{}
	currLexSpan.Start = 1
	prevSpanType := SpanTypeUnknown
	for idx, data := range cfg.Filedata {
		pos := idx + 1
		currSpanType := ClassifyCharacter(string(data))
		if currSpanType != prevSpanType && idx != 0 {
			currLexSpan.Stop = pos - 1
			currLexSpan.Type = prevSpanType
			lexSpans = append(lexSpans, currLexSpan)
			currLexSpan = LexerSpan{Start: pos}
		}
		currLexSpan.Value = currLexSpan.Value + string(data)
		prevSpanType = currSpanType
	}

	return lexSpans, nil
}
