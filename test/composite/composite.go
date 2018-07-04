package main

import (
	"fmt"
	"github.com/fastpopo/gconf"
)

func main() {

	envSource := gconf.NewEnvConfSource()
	jsonSource := gconf.NewJsonFileConfSource("real.json")
	yamlSource := gconf.NewJsonFileConfSource("real.yml")

	conf := gconf.NewConfBuilder().
		Add(envSource).
		Add(jsonSource).
		Add(yamlSource).
		Build()

	pairs := conf.ToKeyValuePairs()
	for _, p := range pairs {
		fmt.Printf("Key: %v, Value: %v\n", p.Key, p.Value)
	}
}
