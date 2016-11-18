package gconf

import (
	"errors"
	"log"

	"github.com/fastpopo/gconf/jsonconf"
)

type _JsonFileConfProvider struct {
	data        map[string]interface{}
	source      FileConfSource
	reloadToken ReloadToken
}

func newJsonFileConfProvider(source FileConfSource) FileConfProvider {
	p := &_JsonFileConfProvider{
		source:      source,
		reloadToken: NewReloadToken(),
	}

	p.Load()

	return p
}

func (f *_JsonFileConfProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := f.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (f *_JsonFileConfProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := f.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (f *_JsonFileConfProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[_JsonFileConfigurationProvider::Set] invalid null argument: key")
	}

	f.data[key] = value
	return nil
}

func (f *_JsonFileConfProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := f.data[key]

	return exist
}

func (f *_JsonFileConfProvider) Keys() []string {

	var keys []string

	for k := range f.data {
		keys = append(keys, k)
	}

	return keys
}

func (f *_JsonFileConfProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range f.data {
		values = append(values, v)
	}

	return values
}

func (f *_JsonFileConfProvider) GetReloadToken() ReloadToken {
	return f.reloadToken
}

func (f *_JsonFileConfProvider) Load() {

	fileInfo := f.source.GetFileInfo()

	if fileInfo.Exists() == false {
		f.data = make(map[string]interface{})
		f.OnReload()
		return
	}

	stream, err := fileInfo.ReadAll()

	if err != nil {
		log.Print("Can't read the file: ", err)
		return
	}

	f.LoadFromStream(stream)
	f.OnReload()
}

func (f *_JsonFileConfProvider) OnReload() {
	prevToken := f.reloadToken
	f.reloadToken = NewReloadToken()

	prevToken.OnReload()
}

func (f *_JsonFileConfProvider) LoadFromStream(stream []byte) error {
	if stream == nil || len(stream) == 0 {
		return nil
	}

	parser := jsonconf.NewJsonConfParser(RootPath, KeyDelimiter)
	err := parser.Parse(stream)

	if err != nil {
		return err
	}

	data := parser.GetDataMap()
	f.data = data

	return nil
}

type _JsonFileConfSource struct {
	path             string
	endureIfNotExist bool
	reloadOnChange   bool
}

func NewJsonFileConfSource(path string, endureIfNotExist bool, reloadOnChange bool) FileConfSource {
	return &_JsonFileConfSource{
		path:             path,
		endureIfNotExist: endureIfNotExist,
		reloadOnChange:   reloadOnChange,
	}
}
func (f *_JsonFileConfSource) Build(builder ConfBuilder) ConfProvider {
	return newJsonFileConfProvider(f)
}

func (f *_JsonFileConfSource) GetFileInfo() FileInfo {
	return NewFileInfo(f.path)
}

func (f *_JsonFileConfSource) GetPath() string {
	return f.path
}

func (f *_JsonFileConfSource) EndureIfNotExist() bool {
	return f.endureIfNotExist
}

func (f *_JsonFileConfSource) ReloadOnChange() bool {
	return f.reloadOnChange
}
