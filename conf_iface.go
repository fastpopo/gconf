package gconf

type KeyValuePair struct {
	Key   string
	Value interface{}
}

type ConfBase interface {
	GetPath() string
	Get(key string) interface{}
	Set(key string, value interface{}) error
	ContainKey(key string) bool
	Keys() []string
	Values() []interface{}
	ToKeyValuePairs() []KeyValuePair
	IsEmpty() bool
	IsArray() bool
}

type Conf interface {
	ConfBase
	GetBoolean(key string) (bool, error)
	GetByte(key string) (byte, error)
	GetInt(key string) (int, error)
	GetInt64(key string) (int64, error)
	GetUint(key string) (uint, error)
	GetUint64(key string) (uint64, error)
	GetFloat32(key string) (float32, error)
	GetFloat64(key string) (float64, error)
	GetComplex64(key string) (complex64, error)
	GetComplex128(key string) (complex128, error)
	GetString(key string) (string, error)
	TryGet(key string, defaultValue interface{}) interface{}
	TryGetBoolean(key string, defaultValue bool) bool
	TryGetByte(key string, defaultValue byte) byte
	TryGetInt(key string, defaultValue int) int
	TryGetInt64(key string, defaultValue int64) int64
	TryGetUint(key string, defaultValue uint) uint
	TryGetUint64(key string, defaultValue uint64) uint64
	TryGetFloat32(key string, defaultValue float32) float32
	TryGetFloat64(key string, defaultValue float64) float64
	TryGetComplex64(key string, defaultValue complex64) complex64
	TryGetComplex128(key string, defaultValue complex128) complex128
	TryGetString(key string, defaultValue string) string
	GetSection(key string) ConfSection
	GetArraySection(key string) ConfArraySection
	//Unmarshal(key string, out interface{}) error  // to be supported in the future
}

type ConfBuilder interface {
	Add(confSource ConfSource) ConfBuilder
	GetSources() []ConfSource
	Build() (ConfRoot, error)
}

type ConfRoot interface {
	Conf
	Reload() error
	Dispose()
}

type ConfSection interface {
	Conf
}

type ConfArraySection interface {
	Conf
	Length() int
	GetIndexSection(idx int) ConfSection
}

type ConfChanges interface {
	GetNumOfChanges() int
	GetChanges() []Change
}

type ConfProvider interface {
	Conf
	Load() error
	Reload() error
	Dispose()
	GetChangeToken() ChangeToken
}

type ConfSource interface {
	Build(confBuilder ConfBuilder) (ConfProvider, error)
	Load() (map[string]interface{}, error)
}

type Watcher interface {
	Watch(reloadToken ChangeToken) error
	IsWatching() bool
	Close() error
}

type ChangeToken interface {
	SetCallback(callback func())
	HasChanged() bool
	SetAsChanged()
	OnChanged()
}
