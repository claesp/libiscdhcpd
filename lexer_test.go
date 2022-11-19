package libiscdhcpd

import (
	"fmt"
	"testing"
)

func TestLexWithNoFiledata(t *testing.T) {
	_, err := Lex(DhcpdDocument{})
	if err == nil {
		t.Fatalf("want error, got nothing")
	}

	e := fmt.Sprintf("%s", err)
	if e != "missing filedate, you need to load a configuration" {
		t.Fatalf("want 'missing filedate, you need to load a configuration', got '%s'", e)
	}
}

func TestLex(t *testing.T) {
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

			for _, span := range spans {
				if span.Type == SpanTypeUnknown {
					t.Fatalf("found unknown span '%s'", span)
				}
			}
		})
	}
}
