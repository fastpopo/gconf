package gconf

import (
	"log"
)

type confSection struct {
	path      string
	root      ConfRoot
	converter TypeConverter
}

func NewConfSection(root ConfRoot, path string) ConfSection {
	if root == nil {
		log.Fatal("[confSection::NewConfSection] invalid null argument: root")
	}

	if path == "" {
		log.Fatal("[confSection::NewConfSection] invalid empty string argument: path")
	}

	s := &confSection{
		root: root,
		path: path,
	}

	s.converter = NewTypeConverter(s)
	return s
}

func (c *confSection) GetPath() string {
	return c.path
}

func (c *confSection) Get(key string) interface{} {
	return c.root.Get(PathCombine(c.path, key))
}

func (c *confSection) Set(key string, value interface{}) error {
	return c.root.Set(PathCombine(c.path, key), value)
}

func (c *confSection) ContainKey(key string) bool {
	return c.root.ContainKey(PathCombine(c.path, key))
}

func (c *confSection) Keys() []string {

	allKeys := c.root.Keys()
	if allKeys == nil || len(allKeys) == 0 {
		return nil
	}

	return FindChildKeys(c.path, allKeys)
}

func (c *confSection) Values() []interface{} {
	pairs := c.root.ToKeyValuePairs()

	if pairs == nil || len(pairs) == 0 {
		return nil
	}

	var values []interface{}

	for _, p := range pairs {
		if HasPathInKey(c.path, p.Key) {
			values = append(values, p.Value)
		}
	}

	return values
}

func (c *confSection) ToKeyValuePairs() []KeyValuePair {
	pairs := c.root.ToKeyValuePairs()

	if pairs == nil || len(pairs) == 0 {
		return nil
	}

	return FindChildPairs(c.path, pairs)
}

func (c *confSection) IsEmpty() bool {
	pairs := c.root.ToKeyValuePairs()

	if pairs == nil || len(pairs) == 0 {
		return true
	}

	for _, p := range pairs {
		if HasPathInKey(c.path, p.Key) {
			return false
		}
	}

	return true
}

func (c *confSection) IsArray() bool {
	return IsArrayPath(c.path, c.Keys())
}

func (c *confSection) GetInt(key string) (int, error) {
	return c.converter.GetInt(key)
}

func (c *confSection) GetInt64(key string) (int64, error) {
	return c.converter.GetInt64(key)
}

func (c *confSection) GetUint(key string) (uint, error) {
	return c.converter.GetUint(key)
}

func (c *confSection) GetUint64(key string) (uint64, error) {
	return c.converter.GetUint64(key)
}

func (c *confSection) GetFloat32(key string) (float32, error) {
	return c.converter.GetFloat32(key)
}

func (c *confSection) GetFloat64(key string) (float64, error) {
	return c.converter.GetFloat64(key)
}

func (c *confSection) GetByte(key string) (byte, error) {
	return c.converter.GetByte(key)
}

func (c *confSection) GetBoolean(key string) (bool, error) {
	return c.converter.GetBoolean(key)
}

func (c *confSection) GetComplex64(key string) (complex64, error) {
	return c.converter.GetComplex64(key)
}

func (c *confSection) GetComplex128(key string) (complex128, error) {
	return c.converter.GetComplex128(key)
}

func (c *confSection) GetString(key string) (string, error) {
	return c.converter.GetString(key)
}

func (c *confSection) TryGet(key string, defaultValue interface{}) interface{} {
	return c.root.TryGet(PathCombine(c.path, key), defaultValue)
}

func (c *confSection) TryGetInt(key string, defaultValue int) int {
	return c.converter.TryGetInt(key, defaultValue)
}

func (c *confSection) TryGetInt64(key string, defaultValue int64) int64 {
	return c.converter.TryGetInt64(key, defaultValue)
}

func (c *confSection) TryGetUint(key string, defaultValue uint) uint {
	return c.converter.TryGetUint(key, defaultValue)
}

func (c *confSection) TryGetUint64(key string, defaultValue uint64) uint64 {
	return c.converter.TryGetUint64(key, defaultValue)
}

func (c *confSection) TryGetFloat32(key string, defaultValue float32) float32 {
	return c.converter.TryGetFloat32(key, defaultValue)
}

func (c *confSection) TryGetFloat64(key string, defaultValue float64) float64 {
	return c.converter.TryGetFloat64(key, defaultValue)
}

func (c *confSection) TryGetByte(key string, defaultValue byte) byte {
	return c.converter.TryGetByte(key, defaultValue)
}

func (c *confSection) TryGetBoolean(key string, defaultValue bool) bool {
	return c.converter.TryGetBoolean(key, defaultValue)
}

func (c *confSection) TryGetComplex64(key string, defaultValue complex64) complex64 {
	return c.converter.TryGetComplex64(key, defaultValue)
}

func (c *confSection) TryGetComplex128(key string, defaultValue complex128) complex128 {
	return c.converter.TryGetComplex128(key, defaultValue)
}

func (c *confSection) TryGetString(key string, defaultValue string) string {
	return c.converter.TryGetString(key, defaultValue)
}

func (c *confSection) GetSection(key string) ConfSection {
	return c.root.GetSection(PathCombine(c.path, key))
}

func (c *confSection) GetArraySection(key string) ConfArraySection {
	return c.root.GetArraySection(PathCombine(c.path, key))
}
