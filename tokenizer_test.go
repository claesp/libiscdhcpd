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
		filename    string
		firstStart  int
		firstStop   int
		firstString string
		firstValue  string
	}{
		{"examples/000-ubuntu-help.conf", 1, 52, "      1:     52 [comment]: # https://help.ubuntu.com/community/isc-dhcp-server", "# https://help.ubuntu.com/community/isc-dhcp-server\n"},
		{"examples/001-isc-dhcp6.conf", 1, 87, "      1:     87 [comment]: # https://gitlab.isc.org/isc-projects/dhcp/-/raw/master/doc/examples/dhcpd-dhcpv6.conf", "# https://gitlab.isc.org/isc-projects/dhcp/-/raw/master/doc/examples/dhcpd-dhcpv6.conf\n"},
		{"examples/002-share-doc-dhcpd-example.conf", 1, 13, "      1:     13 [comment]: # dhcpd.conf", "# dhcpd.conf\n"},
		{"examples/003-ibm-dhcpd-example.conf", 1, 13, "      1:     13 [comment]: # dhcpd.conf", "# dhcpd.conf\n"},
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

			first := tokens[0]
			if first.Start() != tt.firstStart {
				t.Fatalf("start: want '%d', got '%d'", tt.firstStart, first.Start())
			}

			if first.Stop() != tt.firstStop {
				t.Fatalf("stop: want '%d', got '%d'", tt.firstStop, first.Stop())
			}

			if first.String() != tt.firstString {
				t.Fatalf("string: want '%s', got '%s'", tt.firstString, first.String())
			}

			if first.Value() != tt.firstValue {
				t.Fatalf("string: want '%s', got '%s'", tt.firstValue, first.Value())
			}

			for _, token := range tokens {
				if token.Type == TokenTypeUnknown {
					t.Fatalf("found unknown token '%s'", token)
				}
			}
		})
	}
}
