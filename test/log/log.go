package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/fastpopo/gconf"
)

func main() {
	stage := "prod"

	baseSrc := gconf.NewJsonFileConfSource("appsettings.json").
		SetEndureIfNotExist(false)

	prodSrc := gconf.NewJsonFileConfSource(fmt.Sprintf("appsettings-%v.json", stage)).
		SetEndureIfNotExist(false).
		SetOnConfChangedCallback(prodJsonEdited)

	conf := gconf.NewConfBuilder().
		Add(baseSrc).
		Add(prodSrc).
		Build()

	for _, s := range conf.ToKeyValuePairs() {
		fmt.Printf("Key: %s, Value: %v\n", s.Key, s.Value)
	}

	level, err := log.ParseLevel(conf.TryGetString("logging/printLevel", "debug"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.SetLevel(level)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{})

	go logLoop()

	time.Sleep(time.Second * 300)
}

func logLoop() {

	for i := 0; i < 300; i++ {
		log.Debugf("[Debug] index: %v\n", i)
		log.Infof("[Info] index: %v\n", i)
		log.Warnf("[Warn] index: %v\n", i)
		log.Errorf("[Error] index: %v\n", i)
		time.Sleep(time.Second)
	}
}

func prodJsonEdited(confChanges gconf.ConfChanges) {
	fmt.Printf("[prodJsonEdited] Changed: %v\n", confChanges.GetNumOfChanges())

	if confChanges.GetNumOfChanges() == 0 {
		return
	}

	changes := confChanges.GetChanges()

	for _, c := range changes {
		fmt.Printf("%v\n", c.String())

		if c.KeyName != "logging/printLevel" {
			continue
		}

		level, err := log.ParseLevel(c.Current.(string))

		if err != nil {
			log.Error(err)
			return
		}

		log.SetLevel(level)
	}
}
