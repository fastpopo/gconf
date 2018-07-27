package main

import (
	"fmt"
	"github.com/fastpopo/gconf"
)

func main() {
	src := gconf.NewYamlFileConfSource("test.yml").
		SetEndureIfNotExist(false)

	conf, err := gconf.NewConfBuilder().Add(src).Build()

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, s := range conf.ToKeyValuePairs() {
		fmt.Printf("Key: %s\n", s.Key)
	}
}
