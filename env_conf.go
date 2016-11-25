package gconf

import (
	"errors"
	"fmt"
	"github.com/fastpopo/gconf/envconf"
)

type _EnvConfProvider struct {
	data        map[string]string
	source      ConfSource
	converter   *TypeConverter
	reloadToken ReloadToken
	prefix      string
}

func newEnvConfProvider(source ConfSource, prefix string) ConfProvider {
	p := &_EnvConfProvider{
		source:      source,
		prefix:      prefix,
		reloadToken: NewReloadToken(),
	}

	p.converter = NewTypeConverter(p)
	p.Load()

	return p
}

func (p *_EnvConfProvider) GetInt(key string) (int, error) {
	return p.converter.GetInt(key)
}

func (p *_EnvConfProvider) GetInt64(key string) (int64, error) {
	return p.converter.GetInt64(key)
}

func (p *_EnvConfProvider) GetFloat32(key string) (float32, error) {
	return p.converter.GetFloat32(key)
}

func (p *_EnvConfProvider) GetFloat64(key string) (float64, error) {
	return p.converter.GetFloat64(key)
}

func (p *_EnvConfProvider) GetByte(key string) (byte, error) {
	return p.converter.GetByte(key)
}

func (p *_EnvConfProvider) GetBoolean(key string) (bool, error) {
	return p.converter.GetBoolean(key)
}

func (p *_EnvConfProvider) GetString(key string) (string, error) {
	return p.converter.GetString(key)
}

func (p *_EnvConfProvider) TryGetInt(key string, defaultValue int) int {
	return p.converter.TryGetInt(key, defaultValue)
}

func (p *_EnvConfProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return p.converter.TryGetInt64(key, defaultValue)
}

func (p *_EnvConfProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return p.converter.TryGetFloat32(key, defaultValue)
}

func (p *_EnvConfProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return p.converter.TryGetFloat64(key, defaultValue)
}

func (p *_EnvConfProvider) TryGetByte(key string, defaultValue byte) byte {
	return p.converter.TryGetByte(key, defaultValue)
}

func (p *_EnvConfProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return p.converter.TryGetBoolean(key, defaultValue)
}

func (p *_EnvConfProvider) TryGetString(key string, defaultValue string) string {
	return p.converter.TryGetString(key, defaultValue)
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

func (p *_EnvConfProvider) GetSection(key string) ConfSection {
	return newConfSection(p, key)
}

func (p *_EnvConfProvider) Reload() {
	p.Load()
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
