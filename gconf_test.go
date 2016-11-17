package gconf_test

import (
	"fmt"
	"testing"

	"github.com/fastpopo/gconf"
)

func TestMain(t *testing.T) {
	builder := gconf.NewConfBuilder()
	conf := builder.Add(gconf.NewJsonFileConfSource("test/sample.json", true, false)).Build()

	key := "key"
	value := "value"

	fmt.Printf("Testing with test/sample.json\n")

	if v := conf.Get(key); v != value {
		fmt.Printf("Error: key, value not matched.\n")
		fmt.Printf("Expected key: [%s], value: [%s]\n", key, value)
		fmt.Printf("Returned key: [%s], value: [%s]\n", key, v)
		t.Error("Test failed.")
	}
}
