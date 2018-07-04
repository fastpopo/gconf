package gconf

import (
	"errors"
	"log"
)

type fileConfProvider struct {
	path        string
	data        map[string]interface{}
	source      FileConfSource
	converter   TypeConverter
	changeToken ChangeToken
	fileWatcher Watcher
}

func NewFileConfProvider(source FileConfSource) FileConfProvider {
	p := &fileConfProvider{
		path:        RootPath,
		source:      source,
		changeToken: nil,
		data:        make(map[string]interface{}),
	}

	p.converter = NewTypeConverter(p)
	p.Load()
	p.bindFileWatcher()

	return p
}

func (c *fileConfProvider) bindFileWatcher() {
	if c.fileWatcher != nil {
		c.fileWatcher.Close()
		c.fileWatcher = nil
	}

	c.changeToken = NewChangeToken()

	if c.source.GetOnConfChangedCallback() == nil {
		return
	}

	c.changeToken.SetCallback(c.OnChanged)
	watcher, err := NewFileWatcher(c.source.GetFilePath())

	if err != nil {
		log.Println("can't start the filewatcher: " + err.Error())
	} else {
		c.fileWatcher = watcher
		c.fileWatcher.Watch(c.changeToken)
	}
}

func (c *fileConfProvider) GetPath() string {
	return c.path
}

func (c *fileConfProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := c.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (c *fileConfProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[FileConfProvider::Set] invalid null argument: key")
	}

	c.data[key] = value
	return nil
}

func (c *fileConfProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := c.data[key]

	return exist
}

func (c *fileConfProvider) Keys() []string {

	var keys []string

	for k := range c.data {
		keys = append(keys, k)
	}

	return keys
}

func (c *fileConfProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range c.data {
		values = append(values, v)
	}

	return values
}

func (c *fileConfProvider) ToKeyValuePairs() []KeyValuePair {
	var pairs []KeyValuePair

	for k, v := range c.data {

		pair := KeyValuePair{
			Key:   k,
			Value: v,
		}

		pairs = append(pairs, pair)
	}

	return pairs
}

func (c *fileConfProvider) IsEmpty() bool {
	return len(c.data) == 0
}

func (c *fileConfProvider) IsArray() bool {
	return c.GetSection(c.path).IsArray()
}

func (c *fileConfProvider) GetInt(key string) (int, error) {
	return c.converter.GetInt(key)
}

func (c *fileConfProvider) GetInt64(key string) (int64, error) {
	return c.converter.GetInt64(key)
}

func (c *fileConfProvider) GetUint(key string) (uint, error) {
	return c.converter.GetUint(key)
}

func (c *fileConfProvider) GetUint64(key string) (uint64, error) {
	return c.converter.GetUint64(key)
}

func (c *fileConfProvider) GetFloat32(key string) (float32, error) {
	return c.converter.GetFloat32(key)
}

func (c *fileConfProvider) GetFloat64(key string) (float64, error) {
	return c.converter.GetFloat64(key)
}

func (c *fileConfProvider) GetByte(key string) (byte, error) {
	return c.converter.GetByte(key)
}

func (c *fileConfProvider) GetBoolean(key string) (bool, error) {
	return c.converter.GetBoolean(key)
}

func (c *fileConfProvider) GetComplex64(key string) (complex64, error) {
	return c.converter.GetComplex64(key)
}

func (c *fileConfProvider) GetComplex128(key string) (complex128, error) {
	return c.converter.GetComplex128(key)
}

func (c *fileConfProvider) GetString(key string) (string, error) {
	return c.converter.GetString(key)
}

func (c *fileConfProvider) TryGetInt(key string, defaultValue int) int {
	return c.converter.TryGetInt(key, defaultValue)
}

func (c *fileConfProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return c.converter.TryGetInt64(key, defaultValue)
}

func (c *fileConfProvider) TryGetUint(key string, defaultValue uint) uint {
	return c.converter.TryGetUint(key, defaultValue)
}

func (c *fileConfProvider) TryGetUint64(key string, defaultValue uint64) uint64 {
	return c.converter.TryGetUint64(key, defaultValue)
}

func (c *fileConfProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return c.converter.TryGetFloat32(key, defaultValue)
}

func (c *fileConfProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return c.converter.TryGetFloat64(key, defaultValue)
}

func (c *fileConfProvider) TryGetByte(key string, defaultValue byte) byte {
	return c.converter.TryGetByte(key, defaultValue)
}

func (c *fileConfProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return c.converter.TryGetBoolean(key, defaultValue)
}

func (c *fileConfProvider) TryGetComplex64(key string, defaultValue complex64) complex64 {
	return c.converter.TryGetComplex64(key, defaultValue)
}

func (c *fileConfProvider) TryGetComplex128(key string, defaultValue complex128) complex128 {
	return c.converter.TryGetComplex128(key, defaultValue)
}

func (c *fileConfProvider) TryGetString(key string, defaultValue string) string {
	return c.converter.TryGetString(key, defaultValue)
}

func (c *fileConfProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := c.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (c *fileConfProvider) GetSection(key string) ConfSection {
	return NewConfSection(c, PathCombine(c.path, key))
}

func (c *fileConfProvider) GetArraySection(key string) ConfArraySection {
	return NewConfArraySection(c, PathCombine(c.path, key))
}

func (c *fileConfProvider) Reload() {
	c.Load()
}

func (c *fileConfProvider) GetChangeToken() ChangeToken {
	return c.changeToken
}

func (c *fileConfProvider) Load() {
	isExist := c.source.IsFileExist()
	filePath := c.source.GetFilePath()

	if !isExist && !c.source.IsEndureIfNotExist() {
		log.Fatalf("can't find the configuration file: %s\n", filePath)
	}

	if !isExist {
		log.Printf("can't find the configuration file: %s\n", filePath)
		c.data = make(map[string]interface{})
		return
	}

	data, err := c.source.Load()
	if err != nil {
		log.Printf("can't load the configuration file[%s], err: %s\n", filePath, err.Error())
		return
	}

	c.data = data
}

func (c *fileConfProvider) OnChanged() {

	changedData, err := c.source.Load()

	if err != nil {
		log.Println(err.Error())
		return
	}

	changes := CalcConfChanges(changedData, c.data)

	if changes.GetNumOfChanges() == 0 {
		log.Println("there is no changes in file configuration provider")
		return
	}

	c.data = changedData
	c.bindFileWatcher()

	go c.source.GetOnConfChangedCallback()(changes)
}

func (c *fileConfProvider) Dispose() {
	c.fileWatcher.Close()
	c.data = nil
}
