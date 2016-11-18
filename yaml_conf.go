package gconf

import (
	"errors"
	"github.com/fastpopo/gconf/yamlconf"
	"log"
)

type _YamlFileConfProvider struct {
	data        map[string]interface{}
	source      FileConfSource
	reloadToken ReloadToken
}

func newYamlFileConfProvider(source FileConfSource) FileConfProvider {
	p := &_YamlFileConfProvider{
		source:      source,
		reloadToken: NewReloadToken(),
	}

	p.Load()

	return p
}

func (p *_YamlFileConfProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := p.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (p *_YamlFileConfProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := p.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (p *_YamlFileConfProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[_YamlFileConfProvider::Set] invalid null argument: key")
	}

	p.data[key] = value
	return nil
}

func (p *_YamlFileConfProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := p.data[key]

	return exist
}

func (p *_YamlFileConfProvider) Keys() []string {

	var keys []string

	for k := range p.data {
		keys = append(keys, k)
	}

	return keys
}

func (p *_YamlFileConfProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range p.data {
		values = append(values, v)
	}

	return values
}

func (p *_YamlFileConfProvider) GetSection(key string) ConfSection {
	return newConfSection(p, key)
}

func (p *_YamlFileConfProvider) Reload() {
	p.Load()
}

func (p *_YamlFileConfProvider) GetReloadToken() ReloadToken {
	return p.reloadToken
}

func (p *_YamlFileConfProvider) Load() {

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

func (p *_YamlFileConfProvider) OnReload() {
	prevToken := p.reloadToken
	p.reloadToken = NewReloadToken()

	prevToken.OnReload()
}

func (p *_YamlFileConfProvider) LoadFromStream(stream []byte) error {
	if stream == nil || len(stream) == 0 {
		return nil
	}

	parser := yamlconf.NewYamlConfParser(RootPath, KeyDelimiter)
	err := parser.Parse(stream)

	if err != nil {
		return err
	}

	data := parser.GetDataMap()
	p.data = data

	return nil
}

type _YamlFileConfSource struct {
	path             string
	endureIfNotExist bool
	reloadOnChange   bool
}

func NewYamlFileConfSource(path string, endureIfNotExist bool, reloadOnChange bool) FileConfSource {
	return &_YamlFileConfSource{
		path:             path,
		endureIfNotExist: endureIfNotExist,
		reloadOnChange:   reloadOnChange,
	}
}
func (f *_YamlFileConfSource) Build(builder ConfBuilder) ConfProvider {
	return newYamlFileConfProvider(f)
}

func (f *_YamlFileConfSource) GetFileInfo() FileInfo {
	return NewFileInfo(f.path)
}

func (f *_YamlFileConfSource) GetPath() string {
	return f.path
}

func (f *_YamlFileConfSource) EndureIfNotExist() bool {
	return f.endureIfNotExist
}

func (f *_YamlFileConfSource) ReloadOnChange() bool {
	return f.reloadOnChange
}
