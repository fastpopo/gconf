package gconf

import (
	"errors"
	"fmt"
	"github.com/fastpopo/gconf/envconf"
	"sync"
)

type _EnvConfProvider struct {
	data        map[string]string
	source      ConfSource
	reloadToken ReloadToken
	prefix      string
	m           sync.Mutex
}

func newEnvConfProvider(source ConfSource, prefix string) ConfProvider {
	p := &_EnvConfProvider{
		source:      source,
		prefix:      prefix,
		reloadToken: NewReloadToken(),
	}

	p.Load()
	return p
}

func (p *_EnvConfProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := p.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (p *_EnvConfProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := p.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (p *_EnvConfProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[_EnvConfProvider::Set] invalid null argument: key")
	}

	p.data[key] = fmt.Sprint(value)
	return nil
}

func (p *_EnvConfProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := p.data[key]

	return exist
}

func (p *_EnvConfProvider) Keys() []string {

	var keys []string

	for k := range p.data {
		keys = append(keys, k)
	}

	return keys
}

func (p *_EnvConfProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range p.data {
		values = append(values, v)
	}

	return values
}

func (p *_EnvConfProvider) GetReloadToken() ReloadToken {
	return p.reloadToken
}

func (p *_EnvConfProvider) Load() {
	parser := envconf.NewEnvConfParser(p.prefix)
	parser.Parse()

	p.data = parser.GetDataMap()
	p.OnReload()
}

func (p *_EnvConfProvider) OnReload() {
	prevToken := p.reloadToken
	p.reloadToken = NewReloadToken()

	prevToken.OnReload()
}

type _EnvConfSource struct {
	prefix string
}

func NewEnvConfSource(prefix string) ConfSource {
	return &_EnvConfSource{
		prefix: prefix,
	}
}
func (f *_EnvConfSource) Build(builder ConfBuilder) ConfProvider {
	return newEnvConfProvider(f, f.prefix)
}
