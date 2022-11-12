package libiscdhcpd

import (
	"fmt"
	"testing"
)

func TestTokenizeWithNoSpanData(t *testing.T) {
	_, err := Tokenize([]LexerSpan{})
	if err == nil {
		t.Fatalf("want error, got nothing")
	}

	e := fmt.Sprintf("%s", err)
	if e != "no spans" {
		t.Fatalf("want 'no spans', got '%s'", e)
	}
}

func TestTokenize(t *testing.T) {
	var tests = []struct {
		filename string
	}{
		{"examples/000-ubuntu-help.conf"},
		//{"examples/001-isc-dhcp6.conf"},
		{"examples/002-share-doc-dhcpd-example.conf"},
		//{"examples/003-ibm-dhcpd-example.conf"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.filename)
		t.Run(testname, func(t *testing.T) {
			cfg, err := LoadConfigFromFile(tt.filename)
			if err != nil {
				t.Fatalf("want 'no error', got '%s'", err)
			}

			spans, lexErr := Lex(cfg)
			if lexErr != nil {
				t.Fatalf("want 'no error', got '%s'", lexErr)
			}

			tokens, tokErr := Tokenize(spans)
			if tokErr != nil {
				t.Fatalf("want 'no error', got '%s'", tokErr)
			}

			for _, token := range tokens {
				if token.Type == TokenTypeUnknown {
					t.Fatalf("found unknown token '%s'", token)
				}
			}
		})
	}
}
