package libiscdhcpd

import (
	"errors"
	"os"
)

type DhcpdDocument struct {
	Filename string  `json:"filename"`
	Filedata []byte  `json:"filedata"`
	Tokens   []Token `json:"tokens"`
}

func LoadConfigFromFile(filename string) (DhcpdDocument, error) {
	cfg := DhcpdDocument{}
	if filename == "" {
		return cfg, errors.New("filename cannot be empty")
	}
	cfg.Filename = filename

	data, err := os.ReadFile(cfg.Filename)
	if err != nil {
		return cfg, err
	}
	cfg.Filedata = data

	return LoadConfig(cfg)
}

func LoadConfig(config DhcpdDocument) (DhcpdDocument, error) {
	spans, sErr := Lex(config)
	if sErr != nil {
		return config, sErr
	}

	tokens, tErr := Tokenize(spans)
	if tErr != nil {
		return config, tErr
	}
	config.Tokens = tokens

	return config, nil
}
