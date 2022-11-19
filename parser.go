package libiscdhcpd

import (
	"errors"
	"log"
	"strings"
)

type ConfigNode interface {
	Name() string
}

type RootNode struct {
	Options []OptionNode
}

func (n RootNode) Name() string {
	return "root"
}

type OptionNode struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (n OptionNode) Name() string {
	return n.Key
}

type Config struct {
	Root RootNode `json:"dhcpd"`
}

func ParseUnknown(parts []string) (ConfigNode, error) {
	return OptionNode{}, nil
}

func ParseDirective(token Token) (ConfigNode, error) {
	log.Println(token)
	v := strings.Trim(token.Value(), " \n")
	parts := strings.Split(v, " ")

	firstPart := parts[0]

	switch firstPart {
	case "option":
		node, err := ParseUnknown(parts)
		return node, err
	default:
		node, err := ParseUnknown(parts)
		return node, err
	}

	return OptionNode{}, nil
}

func Parse(tokens []Token) (Config, error) {
	cfg := Config{}
	if len(tokens) == 0 {
		return cfg, errors.New("no tokens")
	}

	cfg.Root = RootNode{}
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Type {
		case TokenTypeDirective:
			node, err := ParseDirective(token)
			if err != nil {
				break
			}
			cfg.Root.Options = append(cfg.Root.Options, node.(OptionNode))
		}
	}

	return Config{}, nil
}
