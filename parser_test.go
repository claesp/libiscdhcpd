package libiscdhcpd

import (
	"fmt"
	"testing"
)

func TestParserWithNoTokens(t *testing.T) {
	_, err := Parse([]Token{})
	e := fmt.Sprintf("%s", err)
	if e != "no tokens" {
		t.Fatalf("want 'no tokens', got '%s'", e)
	}
}

func TestParse(t *testing.T) {
	var tests = []struct {
		filename string
	}{
		{"examples/000-ubuntu-help.conf"},
		{"examples/001-isc-dhcp6.conf"},
		{"examples/002-share-doc-dhcpd-example.conf"},
		{"examples/003-ibm-dhcpd-example.conf"},
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

			_, parErr := Parse(tokens)
			if parErr != nil {
				t.Fatalf("want 'no error', got '%s'", tokErr)
			}
		})
	}
}
