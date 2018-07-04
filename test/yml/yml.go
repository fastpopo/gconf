package main

import (
	"fmt"
	"github.com/fastpopo/gconf"
)

func main() {
	src := gconf.NewYamlFileConfSource("test.yml").
		SetEndureIfNotExist(false)

	conf := gconf.NewConfBuilder().
		Add(src).
		Build()

	for _, s := range conf.ToKeyValuePairs() {
		fmt.Printf("Key: %s, Value: %v\n", s.Key, s.Value)
	}

}
