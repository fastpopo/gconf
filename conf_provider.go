package gconf

import (
	"errors"
)

type confProvider struct {
	path        string
	data        map[string]interface{}
	source      ConfSource
	converter   TypeConverter
	changeToken ChangeToken
}

func NewConfProvider(source ConfSource) (ConfProvider, error) {
	p := &confProvider{
		path:        RootPath,
		source:      source,
		changeToken: NewChangeToken(),
	}

	p.converter = NewTypeConverter(p)

	if err := p.Load(); err != nil {
		return nil, err
	}

	return p, nil
}

func (c *confProvider) GetPath() string {
	return c.path
}

func (c *confProvider) Get(key string) interface{} {
	key = PathCombine(c.path, key)

	value, exist := c.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (c *confProvider) Set(key string, value interface{}) error {
	key = PathCombine(c.path, key)

	c.data[key] = value
	return nil
}

func (c *confProvider) ContainKey(key string) bool {
	key = PathCombine(c.path, key)

	_, exist := c.data[key]

	return exist
}

func (c *confProvider) Keys() []string {
	var keys []string

	for k := range c.data {
		keys = append(keys, k)
	}

	return keys
}

func (c *confProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range c.data {
		values = append(values, v)
	}

	return values
}

func (c *confProvider) ToKeyValuePairs() []KeyValuePair {

	var pairs []KeyValuePair

	for k, v := range c.data {
		pair := KeyValuePair{
			Key:   k,
			Value: v,
		}
		pairs = append(pairs, pair)
	}

	return pairs
}

func (c *confProvider) IsEmpty() bool {
	return len(c.data) == 0
}

func (c *confProvider) IsArray() bool {
	return c.GetSection(c.path).IsArray()
}

func (c *confProvider) GetInt(key string) (int, error) {
	return c.converter.GetInt(key)
}

func (c *confProvider) GetInt64(key string) (int64, error) {
	return c.converter.GetInt64(key)
}

func (c *confProvider) GetUint(key string) (uint, error) {
	return c.converter.GetUint(key)
}

func (c *confProvider) GetUint64(key string) (uint64, error) {
	return c.converter.GetUint64(key)
}

func (c *confProvider) GetFloat32(key string) (float32, error) {
	return c.converter.GetFloat32(key)
}

func (c *confProvider) GetFloat64(key string) (float64, error) {
	return c.converter.GetFloat64(key)
}

func (c *confProvider) GetByte(key string) (byte, error) {
	return c.converter.GetByte(key)
}

func (c *confProvider) GetBoolean(key string) (bool, error) {
	return c.converter.GetBoolean(key)
}

func (c *confProvider) GetComplex64(key string) (complex64, error) {
	return c.converter.GetComplex64(key)
}

func (c *confProvider) GetComplex128(key string) (complex128, error) {
	return c.converter.GetComplex128(key)
}

func (c *confProvider) GetString(key string) (string, error) {
	return c.converter.GetString(key)
}

func (c *confProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := c.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (c *confProvider) TryGetInt(key string, defaultValue int) int {
	return c.converter.TryGetInt(key, defaultValue)
}

func (c *confProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return c.converter.TryGetInt64(key, defaultValue)
}

func (c *confProvider) TryGetUint(key string, defaultValue uint) uint {
	return c.converter.TryGetUint(key, defaultValue)
}

func (c *confProvider) TryGetUint64(key string, defaultValue uint64) uint64 {
	return c.converter.TryGetUint64(key, defaultValue)
}

func (c *confProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return c.converter.TryGetFloat32(key, defaultValue)
}

func (c *confProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return c.converter.TryGetFloat64(key, defaultValue)
}

func (c *confProvider) TryGetByte(key string, defaultValue byte) byte {
	return c.converter.TryGetByte(key, defaultValue)
}

func (c *confProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return c.converter.TryGetBoolean(key, defaultValue)
}

func (c *confProvider) TryGetString(key string, defaultValue string) string {
	return c.converter.TryGetString(key, defaultValue)
}

func (c *confProvider) TryGetComplex64(key string, defaultValue complex64) complex64 {
	return c.converter.TryGetComplex64(key, defaultValue)
}

func (c *confProvider) TryGetComplex128(key string, defaultValue complex128) complex128 {
	return c.converter.TryGetComplex128(key, defaultValue)
}

func (c *confProvider) GetSection(key string) ConfSection {
	return NewConfSection(c, PathCombine(c.path, key))
}

func (c *confProvider) GetArraySection(key string) ConfArraySection {
	return NewConfArraySection(c, PathCombine(c.path, key))
}

func (c *confProvider) Reload() error {
	return c.Load()
}

func (c *confProvider) GetCallback() func(ConfChanges) {
	return nil
}

func (c *confProvider) GetChangeToken() ChangeToken {
	return c.changeToken
}

func (c *confProvider) Load() error {
	data, err := c.source.Load()

	if err != nil {
		return errors.New("can't load the contents: " + err.Error())
	}

	c.data = data
	return nil
}

func (c *confProvider) Dispose() {
	c.data = nil
}
