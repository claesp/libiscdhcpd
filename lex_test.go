package libiscdhcpd

import (
	"fmt"
	"testing"
)

func TestLexWithNoFiledata(t *testing.T) {
	err := Lex(IscDhcpdConfig{})
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
		filesize int
	}{
		{"examples/000-ubuntu-help.conf", 1428},
		{"../libiscdhcpd/examples/000-ubuntu-help.conf", 1428},
		{"/Users/claes/Code/libiscdhcpd/examples/000-ubuntu-help.conf", 1428},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.filename)
		t.Run(testname, func(t *testing.T) {
			cfg, err := LoadConfigFromFile(tt.filename)
			if err != nil {
				t.Fatalf("want 'no error', got '%s'", err)
			}

			err = Lex(cfg)
			if err != nil {
				t.Fatalf("want 'no error', got '%s'", err)
			}
		})
	}
}
