package gconf

import (
	"errors"
)

type _MemConfProvider struct {
	data        map[string]interface{}
	source      ConfSource
	reloadToken ReloadToken
	prefix      string
}

func newMemConfProvider(source ConfSource) ConfProvider {
	p := &_MemConfProvider{
		source:      source,
		reloadToken: NewReloadToken(),
	}

	p.Load()
	return p
}

func (p *_MemConfProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := p.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (p *_MemConfProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := p.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (p *_MemConfProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[_EnvConfProvider::Set] invalid null argument: key")
	}

	p.data[key] = value
	return nil
}

func (p *_MemConfProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := p.data[key]

	return exist
}

func (p *_MemConfProvider) Keys() []string {

	var keys []string

	for k := range p.data {
		keys = append(keys, k)
	}

	return keys
}

func (p *_MemConfProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range p.data {
		values = append(values, v)
	}

	return values
}

func (p *_MemConfProvider) GetReloadToken() ReloadToken {
	return p.reloadToken
}

func (p *_MemConfProvider) Load() {
	p.data = make(map[string]interface{})
}

func (p *_MemConfProvider) OnReload() {
	prevToken := p.reloadToken
	p.reloadToken = NewReloadToken()

	prevToken.OnReload()
}

type _MemConfSource struct {
}

func NewMemConfSource() ConfSource {
	return &_MemConfSource{}
}
func (f *_MemConfSource) Build(builder ConfBuilder) ConfProvider {
	return newMemConfProvider(f)
}
