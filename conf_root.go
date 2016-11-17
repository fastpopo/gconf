package gconf

import "errors"

type _ConfigurationRoot struct {
	providers   []ConfProvider
	reloadToken ReloadToken
}

func newConfRoot(providers []ConfProvider) ConfRoot {
	return &_ConfigurationRoot{
		providers:   providers,
		reloadToken: NewReloadToken(),
	}
}

func (c *_ConfigurationRoot) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	for _, provider := range c.providers {
		if provider.ContainKey(key) == false {
			continue
		}

		return provider.Get(key)
	}

	return nil
}

func (c *_ConfigurationRoot) TryGet(key string, defaultValue interface{}) interface{} {
	if key == "" {
		return defaultValue
	}

	result := c.Get(key)

	if result == nil {
		return defaultValue
	}

	return result
}

func (c *_ConfigurationRoot) Set(key string, value interface{}) error {

	if key == "" {
		return errors.New("[_ConfigurationRoot::Set] invalid null argument: key")
	}

	if len(c.providers) == 0 {
		return errors.New("[_ConfigurationRoot::Set] there is no configuration provider")
	}

	for _, provider := range c.providers {
		if provider.ContainKey(key) == true {
			return provider.Set(key, value)
		}
	}

	return c.providers[0].Set(key, value)
}

func (c *_ConfigurationRoot) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	for _, provider := range c.providers {
		if provider.ContainKey(key) == true {
			return true
		}
	}

	return false
}

func (c *_ConfigurationRoot) Keys() []string {
	var keys []string

	for _, provider := range c.providers {
		keys = append(keys, provider.Keys()...)
	}

	return keys
}

func (c *_ConfigurationRoot) Values() []interface{} {
	var values []interface{}

	for _, provider := range c.providers {
		values = append(values, provider.Values()...)
	}

	return values
}

func (c *_ConfigurationRoot) Reload() {
	for _, provider := range c.providers {
		provider.Load()
	}
}
