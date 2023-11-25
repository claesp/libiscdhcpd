package libiscdhcpd

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type ConfigNode interface {
	Name() string
}

type RootNode struct {
	Options []ConfigNode
}

func (n RootNode) Name() string {
	return "root"
}

type OptionNode struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OptionRoutersNode struct {
	IPAddresses []net.IP
}

func (n OptionRoutersNode) Name() string {
	return "Routers"
}

func (n OptionNode) Name() string {
	return n.Key
}

type Config struct {
	Root RootNode `json:"dhcpd"`
}

func ParseOption(token Token) (ConfigNode, error) {
	var opt OptionNode

	for _, span := range token.Spans {
		switch span.Type {
		case SpanTypeWord:
			if opt.Key == "" && span.Value != "option" {
				opt.Key = span.Value
			} else {
				opt.Value = span.Value
			}
		}
	}

	switch opt.Key {
	case "routers":
		ip := net.ParseIP(opt.Value)
		return OptionRoutersNode{IPAddresses: []net.IP{ip}}, nil
	}

	return opt, nil
}

func ParseDirective(token Token) (ConfigNode, error) {
	log.Println(token.Spans)
	for _, span := range token.Spans {
		switch span.Type {
		case SpanTypeWord:
			switch span.Value {
			case "option":
				node, err := ParseOption(token)
				return node, err
			}
		}
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
			switch v := node.(type) {
			case OptionNode:
				cfg.Root.Options = append(cfg.Root.Options, node)
			default:
				fmt.Printf("unexpected type %T", v)
			}
		}
	}

	return cfg, nil
}
