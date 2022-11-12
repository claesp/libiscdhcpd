package libiscdhcpd

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type TokenType int

func (t TokenType) String() string {
	switch t {
	case TokenTypeUnknown:
		return "unknown"
	case TokenTypeComment:
		return "comment"
	case TokenTypeDirective:
		return "directive"
	case TokenTypeStartBlock:
		return "startblock"
	case TokenTypeEndBlock:
		return "endblock"
	case TokenTypeEmptyLine:
		return "emptyline"
	default:
		return "unset"
	}
}

const (
	TokenTypeUnknown TokenType = iota
	TokenTypeComment
	TokenTypeDirective
	TokenTypeStartBlock
	TokenTypeEndBlock
	TokenTypeEmptyLine
)

type Token struct {
	Type  TokenType
	Spans []LexerSpan
}

func (t *Token) Start() int {
	return t.Spans[0].Start
}

func (t *Token) Stop() int {
	return t.Spans[len(t.Spans)-1].Stop
}

func (t *Token) Value() string {
	s := ""
	for _, span := range t.Spans {
		s = s + span.Value
	}
	return s
}

func (t Token) String() string {
	return fmt.Sprintf("%7d:%7d [%s]: %s", t.Start(), t.Stop(), t.Type, strings.TrimSuffix(t.Value(), "\n"))
}

func TokenizeWords(start int, spans []LexerSpan) (Token, int, error) {
	token := Token{}
	token.Type = TokenTypeUnknown
	token.Spans = make([]LexerSpan, 0)
	progress := start

	done := false
	for i := start; i < len(spans); i++ {
		span := spans[i]
		token.Spans = append(token.Spans, span)
		progress = progress + 1

		switch span.Type {
		case SpanTypeComment:
			token.Type = TokenTypeComment
		case SpanTypeSemicolon:
			if token.Type == TokenTypeUnknown {
				token.Type = TokenTypeDirective
			}
		case SpanTypeOpenCurly:
			if token.Type == TokenTypeUnknown {
				token.Type = TokenTypeStartBlock
			}
		case SpanTypeCloseCurly:
			if token.Type == TokenTypeUnknown {
				token.Type = TokenTypeEndBlock
			}
		case SpanTypeNewline:
			if token.Type == TokenTypeUnknown {
				token.Type = TokenTypeEmptyLine
			}
			done = true
		}

		if done {
			break
		}
	}

	log.Println(" -- token", token, "--")

	return token, progress, nil
}

func Tokenize(spans []LexerSpan) ([]Token, error) {
	tokens := make([]Token, 0)
	if len(spans) == 0 {
		return tokens, errors.New("no spans")
	}

	for i := 0; i < len(spans); i++ {
		span := spans[i]
		switch span.Type {
		default:
			token, progress, err := TokenizeWords(i, spans)
			if err != nil {
				log.Println(err)
				continue
			}
			tokens = append(tokens, token)
			i = progress - 1
		}
	}

	return tokens, nil
}
