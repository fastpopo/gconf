package main

import (
	"github.com/fastpopo/gconf"
	"fmt"
)

func main() {
	src := gconf.NewYamlFileConfSource("test.yml").
		SetEndureIfNotExist(false)

	conf := gconf.NewConfBuilder().
		Add(src).
		Build()

	for _, s := range conf.ToArray() {
		fmt.Printf("Key: %s, Value: %v\n", s.Key, s.Value)
	}




}
