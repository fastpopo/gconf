package consulconf

import (
	consul "github.com/hashicorp/consul/api"
	"github.com/fastpopo/gconf"
)

const (
	RootPrefix               string = ""
	defaultScheme            string = "http"
	defaultWatchIntervalSecs int    = 3
	ConsulKeyDelimiter       string = "/"
)

type consulConfSource struct {
	Address           string
	Scheme            string
	Token             string
	Prefix            string
	WatchIntervalSecs int

	client                *consul.Client
	endureIfNotExist      bool
	onConfChangedCallback func(gconf.ConfChanges)
}

func NewConsulConfSource(address string) *consulConfSource {
	return &consulConfSource{
		Address:               address,
		Scheme:                defaultScheme,
		Token:                 "",
		Prefix:                RootPrefix,
		WatchIntervalSecs:     defaultWatchIntervalSecs,
		endureIfNotExist:      true,
		onConfChangedCallback: nil,
	}
}

func (s *consulConfSource) Build(confBuilder gconf.ConfBuilder) gconf.ConfProvider {
	return NewConsulConfProvider(s)
}

func (s *consulConfSource) Load() (map[string]interface{}, error) {
	c := consul.DefaultConfig()
	c.Address = s.Address
	c.Scheme = s.Scheme
	c.Token = s.Token

	client, err := consul.NewClient(c)
	if err != nil {
		return nil, err
	}

	s.client = client
	kv := s.client.KV()

	pairs, _, err := kv.List(s.Prefix, nil)

	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	for _, p := range pairs {
		if p.Value == nil || len(p.Value) == 0 {
			continue
		}

		data[p.Key] = string(p.Value)
	}

	return data, nil
}

func (s *consulConfSource) SetOnConfChangedCallback(onConfChangedCallback func(gconf.ConfChanges)) ConsulConfSource {
	s.onConfChangedCallback = onConfChangedCallback
	return s
}

func (s *consulConfSource) SetEndureIfNotExist(endureIfNotExist bool) ConsulConfSource {
	s.endureIfNotExist = endureIfNotExist
	return s
}

func (s *consulConfSource) GetOnConfChangedCallback() func(gconf.ConfChanges) {
	return s.onConfChangedCallback
}

func (s *consulConfSource) IsEndureIfNotExist() bool {
	return s.endureIfNotExist
}

func (s *consulConfSource) GetKeyValueStore() *consul.KV {
	return s.client.KV()
}

func (s *consulConfSource) GetKeyPrefix() string {
	return s.Prefix
}

func (s *consulConfSource) GetWatchIntervalSecs() int {
	return s.WatchIntervalSecs
}
