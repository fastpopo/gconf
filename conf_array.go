package gconf

import "log"

type confArraySection struct {
	path      string
	root      ConfRoot
	converter TypeConverter
}

func NewConfArraySection(root ConfRoot, path string) ConfArraySection {
	if root == nil {
		log.Fatal("[confArraySection::NewConfSection] invalid null argument: root")
	}

	if path == "" {
		log.Fatal("[confArraySection::NewConfSection] invalid empty string argument: path")
	}

	s := &confArraySection{
		root: root,
		path: path,
	}

	s.converter = NewTypeConverter(s)

	if !s.IsArray() {
		return nil
	}

	return s
}

func (c *confArraySection) GetPath() string {
	return c.path
}

func (c *confArraySection) Get(key string) interface{} {
	return c.root.Get(PathCombine(c.path, key))
}

func (c *confArraySection) Set(key string, value interface{}) error {
	return c.root.Set(PathCombine(c.path, key), value)
}

func (c *confArraySection) ContainKey(key string) bool {
	return c.root.ContainKey(PathCombine(c.path, key))
}

func (c *confArraySection) Keys() []string {

	allKeys := c.root.Keys()
	if allKeys == nil || len(allKeys) == 0 {
		return nil
	}

	return FindChildKeys(c.path, allKeys)
}

func (c *confArraySection) Values() []interface{} {
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

func (c *confArraySection) ToKeyValuePairs() []KeyValuePair {
	pairs := c.root.ToKeyValuePairs()

	if pairs == nil || len(pairs) == 0 {
		return nil
	}

	return FindChildPairs(c.path, pairs)
}

func (c *confArraySection) IsEmpty() bool {
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

func (c *confArraySection) IsArray() bool {
	return IsArrayPath(c.path, c.Keys())
}

func (c *confArraySection) GetInt(key string) (int, error) {
	return c.converter.GetInt(key)
}

func (c *confArraySection) GetInt64(key string) (int64, error) {
	return c.converter.GetInt64(key)
}

func (c *confArraySection) GetUint(key string) (uint, error) {
	return c.converter.GetUint(key)
}

func (c *confArraySection) GetUint64(key string) (uint64, error) {
	return c.converter.GetUint64(key)
}

func (c *confArraySection) GetFloat32(key string) (float32, error) {
	return c.converter.GetFloat32(key)
}

func (c *confArraySection) GetFloat64(key string) (float64, error) {
	return c.converter.GetFloat64(key)
}

func (c *confArraySection) GetByte(key string) (byte, error) {
	return c.converter.GetByte(key)
}

func (c *confArraySection) GetBoolean(key string) (bool, error) {
	return c.converter.GetBoolean(key)
}

func (c *confArraySection) GetComplex64(key string) (complex64, error) {
	return c.converter.GetComplex64(key)
}

func (c *confArraySection) GetComplex128(key string) (complex128, error) {
	return c.converter.GetComplex128(key)
}

func (c *confArraySection) GetString(key string) (string, error) {
	return c.converter.GetString(key)
}

func (c *confArraySection) TryGet(key string, defaultValue interface{}) interface{} {
	return c.root.TryGet(PathCombine(c.path, key), defaultValue)
}

func (c *confArraySection) TryGetInt(key string, defaultValue int) int {
	return c.converter.TryGetInt(key, defaultValue)
}

func (c *confArraySection) TryGetInt64(key string, defaultValue int64) int64 {
	return c.converter.TryGetInt64(key, defaultValue)
}

func (c *confArraySection) TryGetUint(key string, defaultValue uint) uint {
	return c.converter.TryGetUint(key, defaultValue)
}

func (c *confArraySection) TryGetUint64(key string, defaultValue uint64) uint64 {
	return c.converter.TryGetUint64(key, defaultValue)
}

func (c *confArraySection) TryGetFloat32(key string, defaultValue float32) float32 {
	return c.converter.TryGetFloat32(key, defaultValue)
}

func (c *confArraySection) TryGetFloat64(key string, defaultValue float64) float64 {
	return c.converter.TryGetFloat64(key, defaultValue)
}

func (c *confArraySection) TryGetByte(key string, defaultValue byte) byte {
	return c.converter.TryGetByte(key, defaultValue)
}

func (c *confArraySection) TryGetBoolean(key string, defaultValue bool) bool {
	return c.converter.TryGetBoolean(key, defaultValue)
}

func (c *confArraySection) TryGetComplex64(key string, defaultValue complex64) complex64 {
	return c.converter.TryGetComplex64(key, defaultValue)
}

func (c *confArraySection) TryGetComplex128(key string, defaultValue complex128) complex128 {
	return c.converter.TryGetComplex128(key, defaultValue)
}

func (c *confArraySection) TryGetString(key string, defaultValue string) string {
	return c.converter.TryGetString(key, defaultValue)
}

func (c *confArraySection) GetSection(key string) ConfSection {
	return c.root.GetSection(PathCombine(c.path, key))
}

func (c *confArraySection) GetArraySection(key string) ConfArraySection {
	return NewConfArraySection(c.root, PathCombine(c.path, key))
}

func (c *confArraySection) Length() int {
	return LengthOfArrayPath(c.path, c.Keys())
}

func (c *confArraySection) GetIndexSection(idx int) ConfSection {
	return c.GetSection(GetArrayIndex(idx))
}
