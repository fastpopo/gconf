package gconf

import (
	"log"
)

type _ConfSection struct {
	root      ConfRoot
	converter *TypeConverter
	path      string
	key       string
}

func newConfSection(root ConfRoot, path string) ConfSection {
	if root == nil {
		log.Fatal("[_ConfigurationSection::NewConfigurationSection] invalid null argument: root")
	}

	if path == "" {
		log.Fatal("[_ConfigurationSection::NewConfigurationSection] invalid empty string argument: path")
	}

	s := &_ConfSection{
		root: root,
		path: path,
	}

	s.converter = NewTypeConverter(s)
	return s
}

func (c *_ConfSection) GetInt(key string) (int, error) {
	return c.root.GetInt(PathCombine(c.path, key))
}

func (c *_ConfSection) GetInt64(key string) (int64, error) {
	return c.root.GetInt64(PathCombine(c.path, key))
}

func (c *_ConfSection) GetFloat32(key string) (float32, error) {
	return c.root.GetFloat32(PathCombine(c.path, key))
}

func (c *_ConfSection) GetFloat64(key string) (float64, error) {
	return c.root.GetFloat64(PathCombine(c.path, key))
}

func (c *_ConfSection) GetByte(key string) (byte, error) {
	return c.root.GetByte(PathCombine(c.path, key))
}

func (c *_ConfSection) GetBoolean(key string) (bool, error) {
	return c.root.GetBoolean(PathCombine(c.path, key))
}

func (c *_ConfSection) GetString(key string) (string, error) {
	return c.root.GetString(PathCombine(c.path, key))
}

func (c *_ConfSection) TryGetInt(key string, defaultValue int) int {
	return c.root.TryGetInt(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfSection) TryGetInt64(key string, defaultValue int64) int64 {
	return c.root.TryGetInt64(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfSection) TryGetFloat32(key string, defaultValue float32) float32 {
	return c.root.TryGetFloat32(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfSection) TryGetFloat64(key string, defaultValue float64) float64 {
	return c.root.TryGetFloat64(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfSection) TryGetByte(key string, defaultValue byte) byte {
	return c.root.TryGetByte(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfSection) TryGetBoolean(key string, defaultValue bool) bool {
	return c.root.TryGetBoolean(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfSection) TryGetString(key string, defaultValue string) string {
	return c.root.TryGetString(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfSection) Get(key string) interface{} {
	return c.root.Get(PathCombine(c.path, key))
}

func (c *_ConfSection) TryGet(key string, defaultValue interface{}) interface{} {
	return c.root.TryGet(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfSection) Set(key string, value interface{}) error {
	return c.root.Set(PathCombine(c.path, key), value)
}

func (c *_ConfSection) ContainKey(key string) bool {
	return c.root.ContainKey(PathCombine(c.path, key))
}

func (c *_ConfSection) Keys() []string {

	allkeys := c.root.Keys()
	if allkeys == nil || len(allkeys) == 0 {
		return nil
	}

	var newKeys []string

	for _, k := range allkeys {
		if HasPathInKey(c.path, k) == true {
			newKeys = append(newKeys, k)
		}
	}

	return newKeys
}

func (c *_ConfSection) Values() []interface{} {
	keys := c.Keys()

	if keys == nil || len(keys) == 0 {
		return nil
	}

	var values []interface{}

	for _, k := range keys {
		values = append(values, c.root.Get(k))
	}

	return values
}

func (c *_ConfSection) GetSection(key string) ConfSection {
	return c.root.GetSection(PathCombine(c.path, key))
}

func (c *_ConfSection) GetKey() string {
	return GetSectionKey(c.path)
}

func (c *_ConfSection) GetPath() string {
	return c.path
}
