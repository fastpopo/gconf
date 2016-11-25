package gconf

import (
	"errors"
	"log"

	"github.com/fastpopo/gconf/jsonconf"
)

type _JsonFileConfProvider struct {
	data        map[string]interface{}
	source      FileConfSource
	converter   *TypeConverter
	reloadToken ReloadToken
}

func newJsonFileConfProvider(source FileConfSource) FileConfProvider {
	p := &_JsonFileConfProvider{
		source:      source,
		reloadToken: NewReloadToken(),
	}

	p.converter = NewTypeConverter(p)
	p.Load()

	return p
}

func (p *_JsonFileConfProvider) GetInt(key string) (int, error) {
	return p.converter.GetInt(key)
}

func (p *_JsonFileConfProvider) GetInt64(key string) (int64, error) {
	return p.converter.GetInt64(key)
}

func (p *_JsonFileConfProvider) GetFloat32(key string) (float32, error) {
	return p.converter.GetFloat32(key)
}

func (p *_JsonFileConfProvider) GetFloat64(key string) (float64, error) {
	return p.converter.GetFloat64(key)
}

func (p *_JsonFileConfProvider) GetByte(key string) (byte, error) {
	return p.converter.GetByte(key)
}

func (p *_JsonFileConfProvider) GetBoolean(key string) (bool, error) {
	return p.converter.GetBoolean(key)
}

func (p *_JsonFileConfProvider) GetString(key string) (string, error) {
	return p.converter.GetString(key)
}

func (p *_JsonFileConfProvider) TryGetInt(key string, defaultValue int) int {
	return p.converter.TryGetInt(key, defaultValue)
}

func (p *_JsonFileConfProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return p.converter.TryGetInt64(key, defaultValue)
}

func (p *_JsonFileConfProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return p.converter.TryGetFloat32(key, defaultValue)
}

func (p *_JsonFileConfProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return p.converter.TryGetFloat64(key, defaultValue)
}

func (p *_JsonFileConfProvider) TryGetByte(key string, defaultValue byte) byte {
	return p.converter.TryGetByte(key, defaultValue)
}

func (p *_JsonFileConfProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return p.converter.TryGetBoolean(key, defaultValue)
}

func (p *_JsonFileConfProvider) TryGetString(key string, defaultValue string) string {
	return p.converter.TryGetString(key, defaultValue)
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
