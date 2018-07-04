package gconf

import (
	"errors"
)

type confRoot struct {
	path      string
	providers []ConfProvider
	converter TypeConverter
}

func newConfRoot(providers []ConfProvider) ConfRoot {
	root := &confRoot{
		path: RootPath,
		providers: providers,
	}

	root.converter = NewTypeConverter(root)

	return root
}

func (c *confRoot) GetPath() string {
	return c.path
}

func (c *confRoot) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	for _, p := range c.providers {
		if !p.ContainKey(key) {
			continue
		}
		return p.Get(key)
	}

	return nil
}

func (c *confRoot) Set(key string, value interface{}) error {

	if key == "" {
		return errors.New("[confRoot::Set] invalid null argument: key")
	}

	if len(c.providers) == 0 {
		return errors.New("[confRoot::Set] there is no configuration provider")
	}

	for _, p := range c.providers {
		if p.ContainKey(key) {
			return p.Set(key, value)
		}
	}

	return c.providers[0].Set(key, value)
}

func (c *confRoot) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	for _, p := range c.providers {
		if p.ContainKey(key) {
			return true
		}
	}

	return false
}

func (c *confRoot) Keys() []string {
	pairMap := c.getCombinedMap()

	var keys []string

	for k := range pairMap {
		keys = append(keys, k)
	}

	return keys
}

func (c *confRoot) Values() []interface{} {
	pairMap := c.getCombinedMap()

	var values []interface{}

	for _, v := range pairMap {
		values = append(values, v)
	}

	return values
}

func (c *confRoot) ToKeyValuePairs() []KeyValuePair {
	pairMap := c.getCombinedMap()

	var pairs []KeyValuePair
	for k, v := range pairMap {
		pair := KeyValuePair{
			Key:   k,
			Value: v,
		}
		pairs = append(pairs, pair)
	}

	return pairs
}

func (c *confRoot) IsEmpty() bool {
	for _, p := range c.providers {
		if !p.IsEmpty() {
			return false
		}
	}

	return true
}

func (c *confRoot) IsArray() bool {
	return c.GetSection(c.path).IsArray()
}

func (c *confRoot) getCombinedMap() map[string]interface{} {
	pairMap := make(map[string]interface{})

	for i := len(c.providers) - 1; i >= 0; i-- {
		subPairs := c.providers[i].ToKeyValuePairs()

		if subPairs == nil || len(subPairs) == 0 {
			continue
		}

		for _, p := range subPairs {
			pairMap[p.Key] = p.Value
		}
	}

	return pairMap
}

func (c *confRoot) GetInt(key string) (int, error) {
	return c.converter.GetInt(key)
}

func (c *confRoot) GetInt64(key string) (int64, error) {
	return c.converter.GetInt64(key)
}

func (c *confRoot) GetUint(key string) (uint, error) {
	return c.converter.GetUint(key)
}

func (c *confRoot) GetUint64(key string) (uint64, error) {
	return c.converter.GetUint64(key)
}

func (c *confRoot) GetFloat32(key string) (float32, error) {
	return c.converter.GetFloat32(key)
}

func (c *confRoot) GetFloat64(key string) (float64, error) {
	return c.converter.GetFloat64(key)
}

func (c *confRoot) GetByte(key string) (byte, error) {
	return c.converter.GetByte(key)
}

func (c *confRoot) GetBoolean(key string) (bool, error) {
	return c.converter.GetBoolean(key)
}

func (c *confRoot) GetComplex64(key string) (complex64, error) {
	return c.converter.GetComplex64(key)
}

func (c *confRoot) GetComplex128(key string) (complex128, error) {
	return c.converter.GetComplex128(key)
}

func (c *confRoot) GetString(key string) (string, error) {
	return c.converter.GetString(key)
}

func (c *confRoot) TryGet(key string, defaultValue interface{}) interface{} {
	if key == "" {
		return defaultValue
	}

	result := c.Get(key)

	if result == nil {
		return defaultValue
	}

	return result
}

func (c *confRoot) TryGetInt(key string, defaultValue int) int {
	return c.converter.TryGetInt(key, defaultValue)
}

func (c *confRoot) TryGetInt64(key string, defaultValue int64) int64 {
	return c.converter.TryGetInt64(key, defaultValue)
}

func (c *confRoot) TryGetUint(key string, defaultValue uint) uint {
	return c.converter.TryGetUint(key, defaultValue)
}

func (c *confRoot) TryGetUint64(key string, defaultValue uint64) uint64 {
	return c.converter.TryGetUint64(key, defaultValue)
}

func (c *confRoot) TryGetFloat32(key string, defaultValue float32) float32 {
	return c.converter.TryGetFloat32(key, defaultValue)
}

func (c *confRoot) TryGetFloat64(key string, defaultValue float64) float64 {
	return c.converter.TryGetFloat64(key, defaultValue)
}

func (c *confRoot) TryGetByte(key string, defaultValue byte) byte {
	return c.converter.TryGetByte(key, defaultValue)
}

func (c *confRoot) TryGetBoolean(key string, defaultValue bool) bool {
	return c.converter.TryGetBoolean(key, defaultValue)
}

func (c *confRoot) TryGetComplex64(key string, defaultValue complex64) complex64 {
	return c.converter.TryGetComplex64(key, defaultValue)
}

func (c *confRoot) TryGetComplex128(key string, defaultValue complex128) complex128 {
	return c.converter.TryGetComplex128(key, defaultValue)
}

func (c *confRoot) TryGetString(key string, defaultValue string) string {
	return c.converter.TryGetString(key, defaultValue)
}

func (c *confRoot) GetSection(key string) ConfSection {
	return NewConfSection(c, PathCombine(c.path, key))
}

func (c *confRoot) GetArraySection(key string) ConfArraySection {
	return NewConfArraySection(c, PathCombine(c.path, key))
}

func (c *confRoot) Reload() {
	for _, p := range c.providers {
		if !p.GetChangeToken().HasChanged() {
			continue
		}

		p.Load()
	}
}

func (c *confRoot) Dispose() {
	for _, p := range c.providers {
		p.Dispose()
	}
}
