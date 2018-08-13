package gconf

import (
	"errors"
	"reflect"
	"strconv"
)

const (
	defaultInt32      int        = 0
	defaultInt64      int64      = 0
	defaultUint32     uint       = 0
	defaultUint64     uint64     = 0
	defaultFloat32    float32    = 0.0
	defaultFloat64    float64    = 0.0
	defaultComplex64  complex64  = 0
	defaultComplex128 complex128 = 0
	defaultBool       bool       = false
	defaultString     string     = ""
	defaultByte       byte       = 0
)

var (
	errorInvalidArgument = errors.New("invalid argument")
	errorCantFindKey     = errors.New("can't find the key in configurations")
	errorCantConvert     = errors.New("can't convert the type")
)

type TypeConverter struct {
	confBase ConfBase
}

func NewTypeConverter(confBase ConfBase) TypeConverter {
	return TypeConverter{
		confBase: confBase,
	}
}

func (t *TypeConverter) GetInt(key string) (int, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultInt32, errorCantFindKey
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return int(v.Float()), nil
	case reflect.String:
		return StringToInt32(v.String())
	}

	return defaultInt32, errorCantConvert
}

func (t *TypeConverter) GetInt64(key string) (int64, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultInt64, errorCantFindKey
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return int64(v.Float()), nil
	case reflect.String:
		return StringToInt64(v.String())
	}

	return defaultInt64, errorCantConvert
}

func (t *TypeConverter) GetUint(key string) (uint, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultUint32, errorCantFindKey
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return uint(v.Float()), nil
	case reflect.String:
		return StringToUint32(v.String())
	}

	return defaultUint32, errorCantConvert
}

func (t *TypeConverter) GetUint64(key string) (uint64, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultUint64, errorCantFindKey
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint64(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return uint64(v.Float()), nil
	case reflect.String:
		return StringToUint64(v.String())
	}

	return defaultUint64, errorCantConvert
}

func (t *TypeConverter) GetFloat32(key string) (float32, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultFloat32, errorCantFindKey
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float32(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float32(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return float32(v.Float()), nil
	case reflect.String:
		return StringToFloat32(v.String())
	}

	return defaultFloat32, errorCantConvert
}

func (t *TypeConverter) GetFloat64(key string) (float64, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultFloat64, errorCantFindKey
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return float64(v.Float()), nil
	case reflect.String:
		return StringToFloat64(v.String())
	}

	return defaultFloat64, errorCantConvert
}

func (t *TypeConverter) GetComplex64(key string) (complex64, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultComplex64, errorCantFindKey
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return complex64(complex(float64(v.Int()), 0)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return complex64(complex(float64(v.Uint()), 0)), nil
	case reflect.Float32, reflect.Float64:
		return complex64(complex(float64(v.Float()), 0)), nil
	case reflect.String:
		temp, err := StringToFloat64(v.String())
		return complex64(complex(temp, 0)), err
	}

	return defaultComplex64, errorCantConvert
}

func (t *TypeConverter) GetComplex128(key string) (complex128, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultComplex128, errorCantFindKey
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return complex(float64(v.Int()), 0), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return complex(float64(v.Uint()), 0), nil
	case reflect.Float32, reflect.Float64:
		return complex(float64(v.Float()), 0), nil
	case reflect.String:
		temp, err := StringToFloat64(v.String())
		return complex(temp, 0), err
	}

	return defaultComplex128, errorCantConvert
}

func (t *TypeConverter) GetByte(key string) (byte, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultByte, errorCantFindKey
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return byte(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return byte(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return byte(v.Float()), nil
	case reflect.String:
		return StringToByte(v.String())
	}

	return defaultByte, errorCantConvert
}

func (t *TypeConverter) GetBoolean(key string) (bool, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultBool, errorCantFindKey
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.String:
		return StringToBoolean(v.String())
	}

	return defaultBool, errorCantConvert
}

func (t *TypeConverter) GetString(key string) (string, error) {
	value := t.confBase.Get(key)

	if value == nil {
		return defaultString, errorCantFindKey
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int64ToString(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Uint64ToString(v.Uint()), nil
	case reflect.Float32:
		return Float32ToString(float32(v.Float())), nil
	case reflect.Float64:
		return Float64ToString(v.Float()), nil
	case reflect.String:
		return v.String(), nil
	}

	return defaultString, errorCantConvert
}

func (t *TypeConverter) TryGetInt(key string, defaultValue int) int {
	v, err := t.GetInt(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetInt64(key string, defaultValue int64) int64 {
	v, err := t.GetInt64(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetUint(key string, defaultValue uint) uint {
	v, err := t.GetUint(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetUint64(key string, defaultValue uint64) uint64 {
	v, err := t.GetUint64(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetFloat32(key string, defaultValue float32) float32 {
	v, err := t.GetFloat32(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetFloat64(key string, defaultValue float64) float64 {
	v, err := t.GetFloat64(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetComplex64(key string, defaultValue complex64) complex64 {
	v, err := t.GetComplex64(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetComplex128(key string, defaultValue complex128) complex128 {
	v, err := t.GetComplex128(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetByte(key string, defaultValue byte) byte {
	v, err := t.GetByte(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetBoolean(key string, defaultValue bool) bool {
	v, err := t.GetBoolean(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func (t *TypeConverter) TryGetString(key string, defaultValue string) string {
	v, err := t.GetString(key)

	if err != nil {
		return defaultValue
	}

	return v
}

func StringToInt32(val string) (int, error) {
	if val == "" {
		return defaultInt32, errorInvalidArgument
	}

	return strconv.Atoi(val)
}

func Int32ToString(val int) string {
	return Int64ToString(int64(val))
}

func StringToInt64(val string) (int64, error) {
	if val == "" {
		return defaultInt64, errorInvalidArgument
	}

	return strconv.ParseInt(val, 10, 64)
}

func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

func StringToUint32(val string) (uint, error) {
	if val == "" {
		return defaultUint32, errorInvalidArgument
	}

	v, err := strconv.ParseUint(val, 10, 32)

	return uint(v), err
}

func Uint32ToString(val uint32) string {
	return Uint64ToString(uint64(val))
}

func StringToUint64(val string) (uint64, error) {
	if val == "" {
		return defaultUint64, errorInvalidArgument
	}

	return strconv.ParseUint(val, 10, 64)
}

func Uint64ToString(val uint64) string {
	return strconv.FormatUint(val, 10)
}

func StringToFloat32(val string) (float32, error) {
	if val == "" {
		return defaultFloat32, errorInvalidArgument
	}

	v, err := strconv.ParseFloat(val, 32)

	if err != nil {
		return defaultFloat32, err
	}

	return float32(v), nil
}

func Float32ToString(val float32) string {
	return strconv.FormatFloat(float64(val), 'f', -1, 32)
}

func StringToFloat64(val string) (float64, error) {
	if val == "" {
		return defaultFloat64, errorInvalidArgument
	}

	return strconv.ParseFloat(val, 64)
}

func Float64ToString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

func StringToByte(val string) (byte, error) {
	if val == "" {
		return defaultByte, errorInvalidArgument
	}

	v, err := strconv.ParseInt(val, 10, 8)

	if err != nil {
		return defaultByte, err
	}

	return byte(v), nil
}

func ByteToString(val byte) string {
	return Int64ToString(int64(val))
}

func StringToBoolean(val string) (bool, error) {
	if val == "" {
		return defaultBool, errorInvalidArgument
	}

	return strconv.ParseBool(val)
}

func BooleanToString(val bool) string {
	return strconv.FormatBool(val)
}
