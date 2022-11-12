package libiscdhcpd

import (
	"errors"
)

type SpanType int

const (
	SpanTypeUnknown SpanType = iota
	SpanTypeWord
	SpanTypeWhitespace
	SpanTypeComment
	SpanTypeNewline
)

type DhcpdLexSpan struct {
	Start int
	Stop  int
	Value int
	Type  SpanType
}

type DhcpdLexTree struct {
	Spans []DhcpdLexSpan
}

func Lex(cfg IscDhcpdConfig) error {
	if len(cfg.Filedata) == 0 {
		return errors.New("missing filedate, you need to load a configuration")
	}

	/*for _, data := range cfg.Filedata {
		_ := string(data)
	}*/

	return nil
}
