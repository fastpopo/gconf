package consulconf

import (
	"log"
	"errors"

	"github.com/fastpopo/gconf"
)

type consulConfProvider struct {
	data        map[string]interface{}
	source      ConsulConfSource
	converter   gconf.TypeConverter
	changeToken gconf.ChangeToken
	watcher     gconf.Watcher
}

func NewConsulConfProvider(source ConsulConfSource) ConsulConfProvider {
	p := &consulConfProvider{
		source:      source,
		changeToken: nil,
		data:        make(map[string]interface{}),
	}

	p.converter = gconf.NewTypeConverter(p)
	p.Load()
	p.bindWatcher()

	return p
}

func (p *consulConfProvider) bindWatcher() {
	if p.watcher != nil {
		p.watcher.Close()
		p.watcher = nil
	}

	p.changeToken = gconf.NewChangeToken()

	if p.source.GetOnConfChangedCallback() == nil {
		return
	}

	p.changeToken.SetCallback(p.OnChanged)
	kv := p.source.GetKeyValueStore()

	if kv == nil {
		log.Println("can't retrieve kv instance from consul conf source")
		return
	}

	watcher, err := NewConsulWatcher(kv, p.source.GetKeyPrefix(), p.source.GetWatchIntervalSecs())

	if err != nil {
		log.Println(err.Error())
		return
	}

	p.watcher = watcher
	p.watcher.Watch(p.changeToken)
}

func (p *consulConfProvider) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	value, exist := p.data[key]

	if exist == false {
		return nil
	}

	return value
}

func (p *consulConfProvider) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("[FileConfProvider::Set] invalid null argument: key")
	}

	p.data[key] = value
	return nil
}

func (p *consulConfProvider) ContainKey(key string) bool {
	if key == "" {
		return false
	}

	_, exist := p.data[key]

	return exist
}

func (p *consulConfProvider) Keys() []string {

	var keys []string

	for k := range p.data {
		keys = append(keys, k)
	}

	return keys
}

func (p *consulConfProvider) Values() []interface{} {
	var values []interface{}

	for _, v := range p.data {
		values = append(values, v)
	}

	return values
}

func (p *consulConfProvider) ToArray() []gconf.KeyValuePair {
	var pairs []gconf.KeyValuePair

	for k, v := range p.data {

		pair := gconf.KeyValuePair{
			Key:   k,
			Value: v,
		}

		pairs = append(pairs, pair)
	}

	return pairs
}

func (p *consulConfProvider) IsEmpty() bool {
	return len(p.data) == 0
}

func (p *consulConfProvider) GetInt(key string) (int, error) {
	return p.converter.GetInt(key)
}

func (p *consulConfProvider) GetInt64(key string) (int64, error) {
	return p.converter.GetInt64(key)
}

func (p *consulConfProvider) GetUint(key string) (uint, error) {
	return p.converter.GetUint(key)
}

func (p *consulConfProvider) GetUint64(key string) (uint64, error) {
	return p.converter.GetUint64(key)
}

func (p *consulConfProvider) GetFloat32(key string) (float32, error) {
	return p.converter.GetFloat32(key)
}

func (p *consulConfProvider) GetFloat64(key string) (float64, error) {
	return p.converter.GetFloat64(key)
}

func (p *consulConfProvider) GetByte(key string) (byte, error) {
	return p.converter.GetByte(key)
}

func (p *consulConfProvider) GetBoolean(key string) (bool, error) {
	return p.converter.GetBoolean(key)
}

func (p *consulConfProvider) GetComplex64(key string) (complex64, error) {
	return p.converter.GetComplex64(key)
}

func (p *consulConfProvider) GetComplex128(key string) (complex128, error) {
	return p.converter.GetComplex128(key)
}

func (p *consulConfProvider) GetString(key string) (string, error) {
	return p.converter.GetString(key)
}

func (p *consulConfProvider) TryGetInt(key string, defaultValue int) int {
	return p.converter.TryGetInt(key, defaultValue)
}

func (p *consulConfProvider) TryGetInt64(key string, defaultValue int64) int64 {
	return p.converter.TryGetInt64(key, defaultValue)
}

func (p *consulConfProvider) TryGetUint(key string, defaultValue uint) uint {
	return p.converter.TryGetUint(key, defaultValue)
}

func (p *consulConfProvider) TryGetUint64(key string, defaultValue uint64) uint64 {
	return p.converter.TryGetUint64(key, defaultValue)
}

func (p *consulConfProvider) TryGetFloat32(key string, defaultValue float32) float32 {
	return p.converter.TryGetFloat32(key, defaultValue)
}

func (p *consulConfProvider) TryGetFloat64(key string, defaultValue float64) float64 {
	return p.converter.TryGetFloat64(key, defaultValue)
}

func (p *consulConfProvider) TryGetByte(key string, defaultValue byte) byte {
	return p.converter.TryGetByte(key, defaultValue)
}

func (p *consulConfProvider) TryGetBoolean(key string, defaultValue bool) bool {
	return p.converter.TryGetBoolean(key, defaultValue)
}

func (p *consulConfProvider) TryGetComplex64(key string, defaultValue complex64) complex64 {
	return p.converter.TryGetComplex64(key, defaultValue)
}

func (p *consulConfProvider) TryGetComplex128(key string, defaultValue complex128) complex128 {
	return p.converter.TryGetComplex128(key, defaultValue)
}

func (p *consulConfProvider) TryGetString(key string, defaultValue string) string {
	return p.converter.TryGetString(key, defaultValue)
}

func (p *consulConfProvider) TryGet(key string, defaultValue interface{}) interface{} {
	value := p.Get(key)

	if value == nil {
		return defaultValue
	}

	return value
}

func (p *consulConfProvider) GetSection(key string) gconf.ConfSection {
	return gconf.NewConfSection(p, key)
}

func (p *consulConfProvider) Reload() {
	p.Load()
}

func (p *consulConfProvider) GetChangeToken() gconf.ChangeToken {
	return p.changeToken
}

func (p *consulConfProvider) Load() {
	data, err := p.source.Load()
	if err != nil {
		log.Printf("can't load the configuration on consul, err: %s\n", err.Error())
		return
	}

	p.data = data
}

func (p *consulConfProvider) OnChanged() {
	changedData, err := p.source.Load()

	if err != nil {
		log.Println(err.Error())
		return
	}

	changes := gconf.CalcConfChanges(changedData, p.data)

	if changes.GetNumOfChanges() == 0 {
		log.Println("there is no changes in the consul kv")
		return
	}

	p.data = changedData

	go p.source.GetOnConfChangedCallback()(changes)
	p.bindWatcher()
}

func (p *consulConfProvider) Dispose() {
	p.watcher.Close()
	p.data = nil
}
