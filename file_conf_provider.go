package gconf

import (
	"log"
	"errors"
)

type fileConfProvider struct {
	data        map[string]interface{}
	source      FileConfSource
	converter   TypeConverter
	changeToken ChangeToken
	fileWatcher Watcher
}

func NewFileConfProvider(source FileConfSource) FileConfProvider {
	p := &fileConfProvider{
		source:      source,
		changeToken: nil,
		data:        make(map[string]interface{}),
	}

	p.converter = NewTypeConverter(p)
	p.Load()
	p.bindFileWatcher()

	return p
}

func (p *fileConfProvider) bindFileWatcher() {
	if p.fileWatcher != nil {
		p.fileWatcher.Close()
		p.fileWatcher = nil
	}

	p.changeToken = NewChangeToken()

	if p.source.GetOnConfChangedCallback() == nil {
		return
	}

	p.changeToken.SetCallback(p.OnChanged)
	watcher, err := NewFileWatcher(p.source.GetFilePath())

	if err != nil {
		log.Println("can't start the filewatcher: " + err.Error())
	} else {
		p.fileWatcher = watcher
		p.fileWatcher.Watch(p.changeToken)
	}
}

func (p *fileConfProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := p.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (p *fileConfProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[FileConfProvider::Set] invalid null argument: key")
	}

	p.data[key] = value
	return nil
}

func (p *fileConfProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := p.data[key]

	return exist
}

func (p *fileConfProvider) Keys() []string {

	var keys []string

	for k := range p.data {
		keys = append(keys, k)
	}

	return keys
}

func (p *fileConfProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range p.data {
		values = append(values, v)
	}

	return values
}

func (p *fileConfProvider) ToArray() []KeyValuePair {
	var pairs []KeyValuePair

	for k, v := range p.data {

		pair := KeyValuePair{
			Key:   k,
			Value: v,
		}

		pairs = append(pairs, pair)
	}

	return pairs
}

func (p *fileConfProvider) IsEmpty() bool {
	return len(p.data) == 0
}

func (p *fileConfProvider) GetInt(key string) (int, error) {
	return p.converter.GetInt(key)
}

func (p *fileConfProvider) GetInt64(key string) (int64, error) {
	return p.converter.GetInt64(key)
}

func (p *fileConfProvider) GetUint(key string) (uint, error) {
	return p.converter.GetUint(key)
}

func (p *fileConfProvider) GetUint64(key string) (uint64, error) {
	return p.converter.GetUint64(key)
}

func (p *fileConfProvider) GetFloat32(key string) (float32, error) {
	return p.converter.GetFloat32(key)
}

func (p *fileConfProvider) GetFloat64(key string) (float64, error) {
	return p.converter.GetFloat64(key)
}

func (p *fileConfProvider) GetByte(key string) (byte, error) {
	return p.converter.GetByte(key)
}

func (p *fileConfProvider) GetBoolean(key string) (bool, error) {
	return p.converter.GetBoolean(key)
}

func (p *fileConfProvider) GetComplex64(key string) (complex64, error) {
	return p.converter.GetComplex64(key)
}

func (p *fileConfProvider) GetComplex128(key string) (complex128, error) {
	return p.converter.GetComplex128(key)
}

func (p *fileConfProvider) GetString(key string) (string, error) {
	return p.converter.GetString(key)
}

func (p *fileConfProvider) TryGetInt(key string, defaultValue int) int {
	return p.converter.TryGetInt(key, defaultValue)
}

func (p *fileConfProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return p.converter.TryGetInt64(key, defaultValue)
}

func (p *fileConfProvider) TryGetUint(key string, defaultValue uint) uint {
	return p.converter.TryGetUint(key, defaultValue)
}

func (p *fileConfProvider) TryGetUint64(key string, defaultValue uint64) uint64 {
	return p.converter.TryGetUint64(key, defaultValue)
}

func (p *fileConfProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return p.converter.TryGetFloat32(key, defaultValue)
}

func (p *fileConfProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return p.converter.TryGetFloat64(key, defaultValue)
}

func (p *fileConfProvider) TryGetByte(key string, defaultValue byte) byte {
	return p.converter.TryGetByte(key, defaultValue)
}

func (p *fileConfProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return p.converter.TryGetBoolean(key, defaultValue)
}

func (p *fileConfProvider) TryGetComplex64(key string, defaultValue complex64) complex64 {
	return p.converter.TryGetComplex64(key, defaultValue)
}

func (p *fileConfProvider) TryGetComplex128(key string, defaultValue complex128) complex128 {
	return p.converter.TryGetComplex128(key, defaultValue)
}

func (p *fileConfProvider) TryGetString(key string, defaultValue string) string {
	return p.converter.TryGetString(key, defaultValue)
}

func (p *fileConfProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := p.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (p *fileConfProvider) GetSection(key string) ConfSection {
	return NewConfSection(p, key)
}

func (p *fileConfProvider) Reload() {
	p.Load()
}

func (p *fileConfProvider) GetChangeToken() ChangeToken {
	return p.changeToken
}

func (p *fileConfProvider) Load() {
	isExist := p.source.IsFileExist()
	filePath := p.source.GetFilePath()

	if !isExist && !p.source.IsEndureIfNotExist() {
		log.Fatalf("can't find the configuration file: %s\n", filePath)
	}

	if !isExist {
		log.Printf("can't find the configuration file: %s\n", filePath)
		p.data = make(map[string]interface{})
		return
	}

	data, err := p.source.Load()
	if err != nil {
		log.Printf("can't load the configuration file[%s], err: %s\n", filePath, err.Error())
		return
	}

	p.data = data
}

func (p *fileConfProvider) OnChanged() {

	changedData, err := p.source.Load()

	if err != nil {
		log.Println(err.Error())
		return
	}

	changes := CalcConfChanges(changedData, p.data)

	if changes.GetNumOfChanges() == 0 {
		log.Println("there is no changes in file configuration provider")
		return
	}

	p.data = changedData
	p.bindFileWatcher()

	go p.source.GetOnConfChangedCallback()(changes)
}

func (p *fileConfProvider) Dispose() {
	p.fileWatcher.Close()
	p.data = nil
}