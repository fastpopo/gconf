package gconf

type Conf interface {
	Get(key string) interface{}
	TryGet(key string, defaultValue interface{}) interface{}
	Set(key string, value interface{}) error
	ContainKey(key string) bool
	Keys() []string
	Values() []interface{}
	GetSection(key string) ConfSection
}

type ConfBuilder interface {
	Add(source ConfSource) ConfBuilder
	GetSources() []ConfSource
	Build() ConfRoot
}

type ConfRoot interface {
	Conf
	Reload()
}

type ConfSection interface {
	Conf
	GetKey() string
	GetPath() string
}

type ConfProvider interface {
	ConfRoot
	GetReloadToken() ReloadToken
	Load()
	OnReload()
}

type ConfSource interface {
	Build(builder ConfBuilder) ConfProvider
}

type ReloadToken interface {
	SetCallback(callback func())
	HasChanged() bool
	SetAsChanged()
	OnReload()
}
