package main

import (
	"fmt"
	"time"
	"github.com/fastpopo/gconf"
	"github.com/fastpopo/consulconf"
)

func main() {
	jsonSource := gconf.NewJsonFileConfSource("appsettings.json")

	source := consulconf.NewConsulConfSource("165.213.198.242:8500").
		SetOnConfChangedCallback(consulChanged)

	conf := gconf.NewConfBuilder().
		Add(jsonSource).
		Add(source).
		Build()

	for i := 0; i < 30; i++ {
		fmt.Println("\nKey-Value Store Dump")
		fmt.Println("=============================================")

		pairs := conf.ToArray()
		for _, p := range pairs {
			fmt.Printf("Key: %v, Value: %v\n", p.Key, p.Value)
		}

		fmt.Printf("=============================================\n\n")

		time.Sleep(time.Second * 10)
	}
}

func consulChanged(confChanges gconf.ConfChanges) {
	fmt.Printf("[consulChanged] Changed: %v\n", confChanges.GetNumOfChanges())

	if confChanges.GetNumOfChanges() == 0 {
		return
	}

	changes := confChanges.GetChanges()

	for _, c := range changes {
		fmt.Printf("%v\n", c.String())
	}
}
