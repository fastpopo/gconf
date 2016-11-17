package main

import (
	"fmt"
	"github.com/fastpopo/gconf"
)

func main() {
	builder := gconf.NewConfBuilder()
	builder.Add(gconf.NewJsonFileConfSource("test.json", true, false))
	builder.Add(gconf.NewYamlFileConfSource("test.yml", true, false))
	builder.Add(gconf.NewEnvConfSource("SCREAM_"))

	conf := builder.Build()

	for _, s := range conf.Keys() {
		fmt.Printf("Key: %s, Value: %v\n", s, conf.Get(s))
	}
}
