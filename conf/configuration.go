package conf

type Configuration interface {
	Keys() []string
	Values() []interface{}
}