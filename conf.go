package gconf

import "errors"

type Configuration interface {
	Get(key string) interface{}
	TryGet(key string, defaultValue interface{}) interface{}
	Set(key string, value interface{}) error
	ContainKey(key string) bool
	Keys() []string
	Values() []interface{}
}

type ConfigurationBuilder interface {
	Add(conf ConfigurationSource) ConfigurationBuilder
}

type ConfigurationSource interface {
}

type ConfigurationRoot interface {
	Configuration
	Reload() (bool, error)
}

type _ConfigurationBuilder struct {
	Sources []ConfigurationSource
}

func NewConfBuilder() ConfigurationBuilder {
	return &_ConfigurationBuilder{}
}

func (c *_ConfigurationBuilder) Add(conf ConfigurationSource) *_ConfigurationBuilder {
	if conf == nil {
		return errors.New("invalid null argument: conf")
	}

	c.Sources = append(c.Sources, conf)
	return c
}
