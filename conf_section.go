package gconf

import (
	"log"
)

type _ConfigurationSection struct {
	root ConfRoot
	path string
	key  string
}

func newConfSection(root ConfRoot, path string) ConfSection {
	if root == nil {
		log.Fatal("[_ConfigurationSection::NewConfigurationSection] invalid null argument: root")
	}

	if path == "" {
		log.Fatal("[_ConfigurationSection::NewConfigurationSection] invalid empty string argument: path")
	}

	return &_ConfigurationSection{
		root: root,
		path: path,
	}
}

func (c *_ConfigurationSection) Get(key string) interface{} {
	return c.root.Get(PathCombine(c.path, key))
}

func (c *_ConfigurationSection) TryGet(key string, defaultValue interface{}) interface{} {
	return c.root.TryGet(PathCombine(c.path, key), defaultValue)
}

func (c *_ConfigurationSection) Set(key string, value interface{}) error {
	return c.root.Set(PathCombine(c.path, key), value)
}

func (c *_ConfigurationSection) ContainKey(key string) bool {
	return c.root.ContainKey(PathCombine(c.path, key))
}

func (c *_ConfigurationSection) Keys() []string {

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

func (c *_ConfigurationSection) Values() []interface{} {
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

func (c *_ConfigurationSection) GetSection(key string) ConfSection {
	return c.root.GetSection(PathCombine(c.path, key))
}

func (c *_ConfigurationSection) GetKey() string {
	return GetSectionKey(c.path)
}

func (c *_ConfigurationSection) GetPath() string {
	return c.path
}
