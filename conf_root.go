package gconf

import "errors"

type _ProviderItem struct {
	provider ConfProvider
	token    ReloadToken
}

type _ConfRoot struct {
	providers   []_ProviderItem
	converter   *TypeConverter
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

	root.converter = NewTypeConverter(root)

	return root
}

func (c *_ConfRoot) GetInt(key string) (int, error) {
	return c.converter.GetInt(key)

}

func (c *_ConfRoot) GetInt64(key string) (int64, error) {
	return c.converter.GetInt64(key)
}

func (c *_ConfRoot) GetFloat32(key string) (float32, error) {
	return c.converter.GetFloat32(key)
}

func (c *_ConfRoot) GetFloat64(key string) (float64, error) {
	return c.converter.GetFloat64(key)
}

func (c *_ConfRoot) GetByte(key string) (byte, error) {
	return c.converter.GetByte(key)
}

func (c *_ConfRoot) GetBoolean(key string) (bool, error) {
	return c.converter.GetBoolean(key)
}

func (c *_ConfRoot) GetString(key string) (string, error) {
	return c.converter.GetString(key)
}

func (c *_ConfRoot) TryGetInt(key string, defaultValue int) int {
	return c.converter.TryGetInt(key, defaultValue)
}

func (c *_ConfRoot) TryGetInt64(key string, defaultValue int64) int64 {
	return c.converter.TryGetInt64(key, defaultValue)
}

func (c *_ConfRoot) TryGetFloat32(key string, defaultValue float32) float32 {
	return c.converter.TryGetFloat32(key, defaultValue)
}

func (c *_ConfRoot) TryGetFloat64(key string, defaultValue float64) float64 {
	return c.converter.TryGetFloat64(key, defaultValue)
}

func (c *_ConfRoot) TryGetByte(key string, defaultValue byte) byte {
	return c.converter.TryGetByte(key, defaultValue)
}

func (c *_ConfRoot) TryGetBoolean(key string, defaultValue bool) bool {
	return c.converter.TryGetBoolean(key, defaultValue)
}

func (c *_ConfRoot) TryGetString(key string, defaultValue string) string {
	return c.converter.TryGetString(key, defaultValue)
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

func (c *_ConfRoot) GetSection(key string) ConfSection {
	return newConfSection(c, key)
}

func (c *_ConfRoot) Reload() {
	for _, p := range c.providers {
		if p.token.HasChanged() == false {
			continue
		}

		p.provider.Load()
	}
}
