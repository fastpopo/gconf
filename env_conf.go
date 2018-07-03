package gconf

import (
	"os"
	"strings"
)

type EnvConfSource struct {
	prefix        string
}

func NewEnvConfSource() *EnvConfSource {
	return &EnvConfSource{
		prefix:        "",
	}
}

func (s *EnvConfSource) Build(builder ConfBuilder) ConfProvider {
	return NewConfProvider(s)
}

func (s *EnvConfSource) Load() (map[string]interface{}, error) {
	parser := newEnvConfParser(s.prefix)
	parser.Parse()

	return parser.GetDataMap(), nil
}

func (s *EnvConfSource) SetPrefix(prefix string) *EnvConfSource {
	s.prefix = prefix
	return s
}

const (
	envKeyValueDelimiter string = "="
)

type envConfParser struct {
	dataMap map[string]interface{}
	prefix  string
}

func newEnvConfParser(prefix string) *envConfParser {
	return &envConfParser{
		dataMap: make(map[string]interface{}),
		prefix:  prefix,
	}
}

func (p *envConfParser) Parse() {
	env := os.Environ()

	for _, set := range env {
		pair := strings.Split(set, envKeyValueDelimiter)

		if len(pair) == 0 {
			continue
		}

		if strings.HasPrefix(pair[0], p.prefix) == false {
			continue
		}

		key := strings.Replace(pair[0], p.prefix, "", 1)
		key = PathCombine(RootPath, key)

		var value string
		if len(pair) > 2 {
			value = p.combineValue(pair[1:])
		} else {
			value = pair[1]
		}

		p.dataMap[key] = value
	}
}

func (p *envConfParser) combineValue(values []string) string {
	if values == nil || len(values) == 0 {
		return ""
	}

	if len(values) == 1 {
		return values[0]
	}

	return strings.Join(values, envKeyValueDelimiter)
}

func (p *envConfParser) GetDataMap() map[string]interface{} {
	return p.dataMap
}
