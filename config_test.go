package libiscdhcpd

import (
	"fmt"
	"testing"
)

func TestLoadConfigFromFileWithNoFilename(t *testing.T) {
	_, err := LoadConfigFromFile("")
	if err == nil {
		t.Fatalf("want error, got nothing")
	}

	e := fmt.Sprintf("%s", err)
	if e != "filename cannot be empty" {
		t.Fatalf("want 'filename cannot be empty', got '%s'", e)
	}
}

func TestLoadConfigFromFileWithMissingFile(t *testing.T) {
	_, err := LoadConfigFromFile("8f8f8f8f8f")
	if err == nil {
		t.Fatalf("want error, got nothing")
	}

	e := fmt.Sprintf("%s", err)
	if e != "open 8f8f8f8f8f: no such file or directory" {
		t.Fatalf("want 'open 8f8f8f8f8f: no such file or directory', got '%s'", e)
	}
}

func TestLoadConfigFromFile(t *testing.T) {
	var tests = []struct {
		filename string
		filesize int
	}{
		{"examples/000-ubuntu-help.conf", 1428},
		{"examples/001-isc-dhcp6.conf", 3456},
		/*{"../libiscdhcpd/examples/000-ubuntu-help.conf", 1428},
		{"/Users/claes/Code/libiscdhcpd/examples/000-ubuntu-help.conf", 1428},*/
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.filename)
		t.Run(testname, func(t *testing.T) {
			cfg, err := LoadConfigFromFile(tt.filename)
			if err != nil {
				t.Fatalf("want 'no error', got '%s'", err)
			}

			if cfg.Filename != tt.filename {
				t.Errorf("filename: want '%s', got '%s'", tt.filename, cfg.Filename)
			}

			if len(cfg.Filedata) != tt.filesize {
				t.Fatalf("filesize: want '%d', got '%d'", tt.filesize, len(cfg.Filedata))
			}
		})
	}
}
