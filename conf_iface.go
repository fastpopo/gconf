package gconf

type Conf interface {
	GetInt(key string) (int, error)
	GetInt64(key string) (int64, error)
	GetFloat32(key string) (float32, error)
	GetFloat64(key string) (float64, error)
	GetByte(key string) (byte, error)
	GetBoolean(key string) (bool, error)
	GetString(key string) (string, error)
	TryGetInt(key string, defaultValue int) int
	TryGetInt64(key string, defaultValue int64) int64
	TryGetFloat32(key string, defaultValue float32) float32
	TryGetFloat64(key string, defaultValue float64) float64
	TryGetByte(key string, defaultValue byte) byte
	TryGetBoolean(key string, defaultValue bool) bool
	TryGetString(key string, defaultValue string) string
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
