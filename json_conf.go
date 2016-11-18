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

func (p *_JsonFileConfProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := p.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (p *_JsonFileConfProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := p.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (p *_JsonFileConfProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[_JsonFileConfigurationProvider::Set] invalid null argument: key")
	}

	p.data[key] = value
	return nil
}

func (p *_JsonFileConfProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := p.data[key]

	return exist
}

func (p *_JsonFileConfProvider) Keys() []string {

	var keys []string

	for k := range p.data {
		keys = append(keys, k)
	}

	return keys
}

func (p *_JsonFileConfProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range p.data {
		values = append(values, v)
	}

	return values
}

func (p *_JsonFileConfProvider) GetSection(key string) ConfSection {
	return newConfSection(p, key)
}

func (p *_JsonFileConfProvider) Reload() {
	p.Load()
}

func (p *_JsonFileConfProvider) GetReloadToken() ReloadToken {
	return p.reloadToken
}

func (p *_JsonFileConfProvider) Load() {

	fileInfo := p.source.GetFileInfo()

	if fileInfo.Exists() == false {
		p.data = make(map[string]interface{})
		p.OnReload()
		return
	}

	stream, err := fileInfo.ReadAll()

	if err != nil {
		log.Print("Can't read the file: ", err)
		return
	}

	p.LoadFromStream(stream)
	p.OnReload()
}

func (p *_JsonFileConfProvider) OnReload() {
	prevToken := p.reloadToken
	p.reloadToken = NewReloadToken()

	prevToken.OnReload()
}

func (p *_JsonFileConfProvider) LoadFromStream(stream []byte) error {
	if stream == nil || len(stream) == 0 {
		return nil
	}

	parser := jsonconf.NewJsonConfParser(RootPath, KeyDelimiter)
	err := parser.Parse(stream)

	if err != nil {
		return err
	}

	data := parser.GetDataMap()
	p.data = data

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
func (s *_JsonFileConfSource) Build(builder ConfBuilder) ConfProvider {
	return newJsonFileConfProvider(s)
}

func (s *_JsonFileConfSource) GetFileInfo() FileInfo {
	return NewFileInfo(s.path)
}

func (s *_JsonFileConfSource) GetPath() string {
	return s.path
}

func (s *_JsonFileConfSource) EndureIfNotExist() bool {
	return s.endureIfNotExist
}

func (s *_JsonFileConfSource) ReloadOnChange() bool {
	return s.reloadOnChange
}
