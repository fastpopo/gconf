package gconf

import "errors"

type _ProviderItem struct {
	provider ConfProvider
	token    ReloadToken
}

type _ConfRoot struct {
	providers   []_ProviderItem
	reloadToken ReloadToken
}

func newConfRoot(providers []ConfProvider) ConfRoot {
	root := &_ConfRoot{
		reloadToken: NewReloadToken(),
	}

	for _, p := range providers {
		item := _ProviderItem{
			provider: p,
			token:    p.GetReloadToken(),
		}

		root.providers = append(root.providers, item)
	}

	return root
}

func (c *_ConfRoot) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	for _, p := range c.providers {
		if p.provider.ContainKey(key) == false {
			continue
		}

		return p.provider.Get(key)
	}

	return nil
}

func (c *_ConfRoot) TryGet(key string, defaultValue interface{}) interface{} {
	if key == "" {
		return defaultValue
	}

	result := c.Get(key)

	if result == nil {
		return defaultValue
	}

	return result
}

func (c *_ConfRoot) Set(key string, value interface{}) error {

	if key == "" {
		return errors.New("[_ConfigurationRoot::Set] invalid null argument: key")
	}

	if len(c.providers) == 0 {
		return errors.New("[_ConfigurationRoot::Set] there is no configuration provider")
	}

	for _, p := range c.providers {
		if p.provider.ContainKey(key) == true {
			return p.provider.Set(key, value)
		}
	}

	return c.providers[0].provider.Set(key, value)
}

func (c *_ConfRoot) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	for _, p := range c.providers {
		if p.provider.ContainKey(key) == true {
			return true
		}
	}

	return false
}

func (c *_ConfRoot) Keys() []string {
	var keys []string

	for _, p := range c.providers {
		keys = append(keys, p.provider.Keys()...)
	}

	return keys
}

func (c *_ConfRoot) Values() []interface{} {
	var values []interface{}

	for _, p := range c.providers {
		values = append(values, p.provider.Values()...)
	}

	return values
}

func (c *_ConfRoot) Reload() {
	for _, p := range c.providers {
		if p.token.HasChanged() == false {
			continue
		}

		p.provider.Load()
	}
}
