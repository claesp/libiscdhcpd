package libiscdhcpd

import (
	"errors"
	"os"
)

type IscDhcpdConfig struct {
	Filename string `json:"filename"`
	Filedata []byte `json:"filedata"`
}

func LoadConfigFromFile(filename string) (IscDhcpdConfig, error) {
	cfg := IscDhcpdConfig{}
	if filename == "" {
		return cfg, errors.New("filename cannot be empty")
	}
	cfg.Filename = filename

	data, err := os.ReadFile(cfg.Filename)
	if err != nil {
		return cfg, err
	}
	cfg.Filedata = data

	return cfg, nil
}
