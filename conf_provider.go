package gconf

import (
	"errors"
	"log"
)

type confProvider struct {
	data        map[string]interface{}
	source      ConfSource
	converter   TypeConverter
	changeToken ChangeToken
}

func NewConfProvider(source ConfSource) ConfProvider {
	p := &confProvider {
		source: source,
		changeToken: NewChangeToken(),
	}

	p.converter = NewTypeConverter(p)
	p.Load()

	return p
}


func (p *confProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := p.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (p *confProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[envConfProvider::Set] invalid null argument: key")
	}

	p.data[key] = value
	return nil
}

func (p *confProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := p.data[key]

	return exist
}

func (p *confProvider) Keys() []string {
	var keys []string

	for k := range p.data {
		keys = append(keys, k)
	}

	return keys
}

func (p *confProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range p.data {
		values = append(values, v)
	}

	return values
}

func (p *confProvider) ToArray() []KeyValuePair {

	var pairs []KeyValuePair

	for k, v := range p.data {
		pair := KeyValuePair {
			Key: k,
			Value: v,
		}
		pairs = append(pairs, pair)
	}

	return pairs
}

func (p *confProvider) IsEmpty() bool {
	return len(p.data) == 0
}

func (p *confProvider) GetInt(key string) (int, error) {
	return p.converter.GetInt(key)
}

func (p *confProvider) GetInt64(key string) (int64, error) {
	return p.converter.GetInt64(key)
}

func (p *confProvider) GetUint(key string) (uint, error) {
	return p.converter.GetUint(key)
}

func (p *confProvider) GetUint64(key string) (uint64, error) {
	return p.converter.GetUint64(key)
}

func (p *confProvider) GetFloat32(key string) (float32, error) {
	return p.converter.GetFloat32(key)
}

func (p *confProvider) GetFloat64(key string) (float64, error) {
	return p.converter.GetFloat64(key)
}

func (p *confProvider) GetByte(key string) (byte, error) {
	return p.converter.GetByte(key)
}

func (p *confProvider) GetBoolean(key string) (bool, error) {
	return p.converter.GetBoolean(key)
}

func (p *confProvider) GetComplex64(key string) (complex64, error) {
	return p.converter.GetComplex64(key)
}

func (p *confProvider) GetComplex128(key string) (complex128, error) {
	return p.converter.GetComplex128(key)
}

func (p *confProvider) GetString(key string) (string, error) {
	return p.converter.GetString(key)
}

func (p *confProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := p.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (p *confProvider) TryGetInt(key string, defaultValue int) int {
	return p.converter.TryGetInt(key, defaultValue)
}

func (p *confProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return p.converter.TryGetInt64(key, defaultValue)
}

func (p *confProvider) TryGetUint(key string, defaultValue uint) uint {
	return p.converter.TryGetUint(key, defaultValue)
}

func (p *confProvider) TryGetUint64(key string, defaultValue uint64) uint64 {
	return p.converter.TryGetUint64(key, defaultValue)
}

func (p *confProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return p.converter.TryGetFloat32(key, defaultValue)
}

func (p *confProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return p.converter.TryGetFloat64(key, defaultValue)
}

func (p *confProvider) TryGetByte(key string, defaultValue byte) byte {
	return p.converter.TryGetByte(key, defaultValue)
}

func (p *confProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return p.converter.TryGetBoolean(key, defaultValue)
}

func (p *confProvider) TryGetString(key string, defaultValue string) string {
	return p.converter.TryGetString(key, defaultValue)
}

func (p *confProvider) TryGetComplex64(key string, defaultValue complex64) complex64 {
	return p.converter.TryGetComplex64(key, defaultValue)
}

func (p *confProvider) TryGetComplex128(key string, defaultValue complex128) complex128 {
	return p.converter.TryGetComplex128(key, defaultValue)
}

func (p *confProvider) GetSection(key string) ConfSection {
	return NewConfSection(p, key)
}

func (p *confProvider) Reload() {
	p.Load()
}

func (p *confProvider) GetCallback() func(ConfChanges) {
	return nil
}

func (p *confProvider) GetChangeToken() ChangeToken {
	return p.changeToken
}

func (p *confProvider) Load() {
	data, err := p.source.Load()

	if err != nil {
		log.Printf("can't load the configuration map from ConfSource: " + err.Error())
		return
	}

	p.data = data
}

func (p *confProvider) Dispose() {
	p.data = nil
}