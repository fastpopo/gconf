package gconf

import (
	"errors"
	"reflect"
	"strconv"
)

type TypeConverter struct {
	conf Conf
}

var (
	defaultInt32   int     = 0
	defaultInt64   int64   = 0
	defaultFloat32 float32 = 0.0
	defaultFloat64 float64 = 0.0
	defaultBool    bool    = false
	defaultString  string  = ""
	defaultByte    byte    = 0

	Error_Invalid_Argument error = errors.New("Invalid argument")
	Error_Cant_Find_Key    error = errors.New("Can't find the key in configurations")
	Error_Cant_Convert     error = errors.New("Can't convert the type")
)

func NewTypeConverter(conf Conf) *TypeConverter {
	return &TypeConverter{
		conf: conf,
	}
}

func (t *TypeConverter) GetInt(key string) (int, error) {
	if key == "" {
		return defaultInt32, Error_Invalid_Argument
	}

	i := t.conf.Get(key)

	if i == nil {
		return defaultInt32, Error_Cant_Find_Key
	}

	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Int:
		return int(v.Int()), nil
	case reflect.Int8:
		return int(v.Int()), nil
	case reflect.Int16:
		return int(v.Int()), nil
	case reflect.Int32:
		return int(v.Int()), nil
	case reflect.Int64:
		return int(v.Int()), nil
	case reflect.Float32:
		return int(v.Float()), nil
	case reflect.Float64:
		return int(v.Float()), nil
	case reflect.String:
		return t.stringToInt32(v.String())
	}

	return defaultInt32, Error_Cant_Convert
}

func (t *TypeConverter) stringToInt32(val string) (int, error) {
	if val == "" {
		return defaultInt32, Error_Invalid_Argument
	}

	return strconv.Atoi(val)
}

func (t *TypeConverter) GetInt64(key string) (int64, error) {
	if key == "" {
		return defaultInt64, Error_Invalid_Argument
	}

	i := t.conf.Get(key)

	if i == nil {
		return defaultInt64, Error_Cant_Find_Key
	}

	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Int:
		return int64(v.Int()), nil
	case reflect.Int8:
		return int64(v.Int()), nil
	case reflect.Int16:
		return int64(v.Int()), nil
	case reflect.Int32:
		return int64(v.Int()), nil
	case reflect.Int64:
		return int64(v.Int()), nil
	case reflect.Float32:
		return int64(v.Float()), nil
	case reflect.Float64:
		return int64(v.Float()), nil
	case reflect.String:
		return t.stringToInt64(v.String())
	}

	return defaultInt64, Error_Cant_Convert
}

func (t *TypeConverter) stringToInt64(val string) (int64, error) {
	if val == "" {
		return defaultInt64, Error_Invalid_Argument
	}

	return strconv.ParseInt(val, 10, 64)
}

func (t *TypeConverter) GetFloat32(key string) (float32, error) {
	if key == "" {
		return defaultFloat32, Error_Invalid_Argument
	}

	i := t.conf.Get(key)

	if i == nil {
		return defaultFloat32, Error_Cant_Find_Key
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {

	case reflect.Float32:
		return float32(v.Float()), nil
	case reflect.Float64:
		return float32(v.Float()), nil
	case reflect.String:
		return t.stringToFloat32(v.String())
	}

	return defaultFloat32, Error_Cant_Convert
}

func (t *TypeConverter) stringToFloat32(val string) (float32, error) {
	if val == "" {
		return defaultFloat32, Error_Invalid_Argument
	}

	v, err := strconv.ParseFloat(val, 32)

	if err != nil {
		return defaultFloat32, err
	}

	return float32(v), nil
}

func (t *TypeConverter) GetFloat64(key string) (float64, error) {
	if key == "" {
		return defaultFloat64, Error_Invalid_Argument
	}

	i := t.conf.Get(key)

	if i == nil {
		return defaultFloat64, Error_Cant_Find_Key
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {

	case reflect.Float32:
		return float64(v.Float()), nil
	case reflect.Float64:
		return float64(v.Float()), nil
	case reflect.String:
		return t.stringToFloat64(v.String())
	}

	return defaultFloat64, Error_Cant_Convert
}

func (t *TypeConverter) stringToFloat64(val string) (float64, error) {
	if val == "" {
		return defaultFloat64, Error_Invalid_Argument
	}

	return strconv.ParseFloat(val, 64)
}

func (t *TypeConverter) GetByte(key string) (byte, error) {
	if key == "" {
		return defaultByte, Error_Invalid_Argument
	}

	i := t.conf.Get(key)

	if i == nil {
		return defaultByte, Error_Cant_Find_Key
	}

	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Int:
		return byte(v.Int()), nil
	case reflect.Int8:
		return byte(v.Int()), nil
	case reflect.Int16:
		return byte(v.Int()), nil
	case reflect.Int32:
		return byte(v.Int()), nil
	case reflect.Int64:
		return byte(v.Int()), nil
	case reflect.Float32:
		return byte(v.Float()), nil
	case reflect.Float64:
		return byte(v.Float()), nil
	case reflect.String:
		return t.stringToByte(v.String())
	}

	return defaultByte, Error_Cant_Convert
}

func (t *TypeConverter) stringToByte(val string) (byte, error) {
	if val == "" {
		return defaultByte, Error_Invalid_Argument
	}

	v, err := strconv.ParseInt(val, 10, 8)

	if err != nil {
		return defaultByte, err
	}

	return byte(v), nil
}

func (t *TypeConverter) GetBoolean(key string) (bool, error) {
	if key == "" {
		return defaultBool, Error_Invalid_Argument
	}

	i := t.conf.Get(key)

	if i == nil {
		return defaultBool, Error_Cant_Find_Key
	}

	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.String:
		return t.stringToBoolean(v.String())
	}

	return defaultBool, Error_Cant_Convert
}

func (t *TypeConverter) stringToBoolean(val string) (bool, error) {
	if val == "" {
		return defaultBool, Error_Invalid_Argument
	}

	return strconv.ParseBool(val)
}

func (t *TypeConverter) GetString(key string) (string, error) {
	if key == "" {
		return defaultString, Error_Invalid_Argument
	}

	i := t.conf.Get(key)

	if i == nil {
		return defaultString, Error_Cant_Find_Key
	}

	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	}

	return defaultString, Error_Cant_Convert
}

func (t *TypeConverter) TryGetInt(key string, defaultValue int) int {
	if key == "" {
		return defaultValue
	}

	v, err := t.conf.GetInt(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetInt64(key string, defaultValue int64) int64 {
	if key == "" {
		return defaultValue
	}

	v, err := t.conf.GetInt64(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetFloat32(key string, defaultValue float32) float32 {
	if key == "" {
		return defaultValue
	}

	v, err := t.conf.GetFloat32(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetFloat64(key string, defaultValue float64) float64 {
	if key == "" {
		return defaultValue
	}

	v, err := t.conf.GetFloat64(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetByte(key string, defaultValue byte) byte {
	if key == "" {
		return defaultValue
	}

	v, err := t.conf.GetByte(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetBoolean(key string, defaultValue bool) bool {
	if key == "" {
		return defaultValue
	}

	v, err := t.conf.GetBoolean(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetString(key string, defaultValue string) string {
	if key == "" {
		return defaultValue
	}

	v, err := t.conf.GetString(key)

	if err != nil {
		return defaultValue
	}

	return v
}
