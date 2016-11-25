package gconf

import (
	"errors"
	"github.com/fastpopo/gconf/yamlconf"
	"log"
)

type _YamlFileConfProvider struct {
	data        map[string]interface{}
	source      FileConfSource
	converter   *TypeConverter
	reloadToken ReloadToken
}

func newYamlFileConfProvider(source FileConfSource) FileConfProvider {
	p := &_YamlFileConfProvider{
		source:      source,
		reloadToken: NewReloadToken(),
	}

	p.converter = NewTypeConverter(p)
	p.Load()

	return p
}

func (p *_YamlFileConfProvider) GetInt(key string) (int, error) {
	return p.converter.GetInt(key)
}

func (p *_YamlFileConfProvider) GetInt64(key string) (int64, error) {
	return p.converter.GetInt64(key)
}

func (p *_YamlFileConfProvider) GetFloat32(key string) (float32, error) {
	return p.converter.GetFloat32(key)
}

func (p *_YamlFileConfProvider) GetFloat64(key string) (float64, error) {
	return p.converter.GetFloat64(key)
}

func (p *_YamlFileConfProvider) GetByte(key string) (byte, error) {
	return p.converter.GetByte(key)
}

func (p *_YamlFileConfProvider) GetBoolean(key string) (bool, error) {
	return p.converter.GetBoolean(key)
}

func (p *_YamlFileConfProvider) GetString(key string) (string, error) {
	return p.converter.GetString(key)
}

func (p *_YamlFileConfProvider) TryGetInt(key string, defaultValue int) int {
	return p.converter.TryGetInt(key, defaultValue)
}

func (p *_YamlFileConfProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return p.converter.TryGetInt64(key, defaultValue)
}

func (p *_YamlFileConfProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return p.converter.TryGetFloat32(key, defaultValue)
}

func (p *_YamlFileConfProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return p.converter.TryGetFloat64(key, defaultValue)
}

func (p *_YamlFileConfProvider) TryGetByte(key string, defaultValue byte) byte {
	return p.converter.TryGetByte(key, defaultValue)
}

func (p *_YamlFileConfProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return p.converter.TryGetBoolean(key, defaultValue)
}

func (p *_YamlFileConfProvider) TryGetString(key string, defaultValue string) string {
	return p.converter.TryGetString(key, defaultValue)
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
