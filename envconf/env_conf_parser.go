package envconf

import (
	"os"
	"strings"
)

const (
	EnvKeyValueDelimiter string = "="
)

type _EnvConfParser struct {
	dataMap map[string]string
	prefix  string
}

func NewEnvConfParser(prefix string) *_EnvConfParser {
	return &_EnvConfParser{
		dataMap: make(map[string]string),
		prefix:  prefix,
	}
}

func (p *_EnvConfParser) Parse() {
	env := os.Environ()

	for _, set := range env {
		pair := strings.Split(set, EnvKeyValueDelimiter)

		if len(pair) == 0 {
			continue
		}

		if strings.HasPrefix(pair[0], p.prefix) == false {
			continue
		}

		key := strings.Replace(pair[0], p.prefix, "", 1)
		value := pair[1]
		if len(pair) > 2 {
			value = p.combineValue(pair[1:])
		}

		p.dataMap[key] = value
	}
}

func (p *_EnvConfParser) combineValue(values []string) string {
	if values == nil || len(values) == 0 {
		return ""
	}

	if len(values) == 1 {
		return values[0]
	}

	return strings.Join(values, EnvKeyValueDelimiter)
}

func (p *_EnvConfParser) GetDataMap() map[string]string {
	return p.dataMap
}
