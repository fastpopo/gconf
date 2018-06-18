package consulconf

import (
	consul "github.com/hashicorp/consul/api"
	"github.com/fastpopo/gconf"
)

type ConsulConfSource interface {
	gconf.ConfSource
	SetOnConfChangedCallback(func(gconf.ConfChanges)) ConsulConfSource
	SetEndureIfNotExist(bool) ConsulConfSource
	GetOnConfChangedCallback() func(gconf.ConfChanges)
	IsEndureIfNotExist() bool
	GetKeyValueStore() *consul.KV
	GetKeyPrefix() string
	GetWatchIntervalSecs() int
}

type ConsulConfProvider interface {
	gconf.ConfProvider
	OnChanged()
}
