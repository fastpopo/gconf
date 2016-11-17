# gconf

<pre>
                           __ 
  __ _   ___  ___   _ __   / _|
 / _` | / __|/ _ \ | '_ \ | |_ 
| (_| || (__| (_) || | | ||  _|
 \__, | \___|\___/ |_| |_||_|  
 |___/                         
</pre>


Introduction
------------

gconf is a framework for accessing Key/Value based configuration settings in an pllication. Includes configuration providers for environment variables, JSON files, YAML files.


Installation and usage
----------------------

The import path for the pacakge is *github.com/fastpopo/gconf*.

To install it, run:

    go get github.com/fastpopo/gconf


Example
-------

```Go
package main

import (
	"fmt"

	"github.com/fastpopo/gconf"
)

func main() {
	builder := gconf.NewConfBuilder()
	conf := builder.Add(gconf.NewJsonFileConfSource("test/sample.json", true, false)).Build()

	key := "key"

	v := conf.Get(key)
    fmt.Printf("Key: [%s], Value: [%s]\n", key, v)
    
    return
}

```

Test
----

     go test


License
-------

The gconf packages is licensed under the Apache License 2.0. Please see the LICENSE file for details.
