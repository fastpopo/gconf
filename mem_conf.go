package gconf

import (
	"errors"
)

type _MemConfProvider struct {
	data        map[string]interface{}
	source      ConfSource
	converter   *TypeConverter
	reloadToken ReloadToken
	prefix      string
}

func newMemConfProvider(source ConfSource) ConfProvider {
	p := &_MemConfProvider{
		source:      source,
		reloadToken: NewReloadToken(),
	}

	p.converter = NewTypeConverter(p)
	p.Load()
	return p
}

func (p *_MemConfProvider) GetInt(key string) (int, error) {
	return p.converter.GetInt(key)
}

func (p *_MemConfProvider) GetInt64(key string) (int64, error) {
	return p.converter.GetInt64(key)
}

func (p *_MemConfProvider) GetFloat32(key string) (float32, error) {
	return p.converter.GetFloat32(key)
}

func (p *_MemConfProvider) GetFloat64(key string) (float64, error) {
	return p.converter.GetFloat64(key)
}

func (p *_MemConfProvider) GetByte(key string) (byte, error) {
	return p.converter.GetByte(key)
}

func (p *_MemConfProvider) GetBoolean(key string) (bool, error) {
	return p.converter.GetBoolean(key)
}

func (p *_MemConfProvider) GetString(key string) (string, error) {
	return p.converter.GetString(key)
}

func (p *_MemConfProvider) TryGetInt(key string, defaultValue int) int {
	return p.converter.TryGetInt(key, defaultValue)
}

func (p *_MemConfProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return p.converter.TryGetInt64(key, defaultValue)
}

func (p *_MemConfProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return p.converter.TryGetFloat32(key, defaultValue)
}

func (p *_MemConfProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return p.converter.TryGetFloat64(key, defaultValue)
}

func (p *_MemConfProvider) TryGetByte(key string, defaultValue byte) byte {
	return p.converter.TryGetByte(key, defaultValue)
}

func (p *_MemConfProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return p.converter.TryGetBoolean(key, defaultValue)
}

func (p *_MemConfProvider) TryGetString(key string, defaultValue string) string {
	return p.converter.TryGetString(key, defaultValue)
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

func (p *_MemConfProvider) GetSection(key string) ConfSection {
	return newConfSection(p, key)
}

func (p *_MemConfProvider) Reload() {
	p.Load()
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
