package consulconf

import (
	"time"

	consul "github.com/hashicorp/consul/api"

	"github.com/fastpopo/gconf"
	"log"
)

type consulWatcher struct {
	kv           *consul.KV
	prefix       string
	lastIndex    uint64
	ticker       *time.Ticker
	intervalSecs int
	changeToken  gconf.ChangeToken
	isWatching   bool
}

func NewConsulWatcher(kv *consul.KV, prefix string, intervalSecs int) (gconf.Watcher, error) {
	consulWatcher := &consulWatcher{
		kv:           kv,
		prefix:       prefix,
		intervalSecs: intervalSecs,
		changeToken:  nil,
		isWatching:   false,
	}

	return consulWatcher, nil
}

func (w *consulWatcher) Watch(reloadToken gconf.ChangeToken) error {
	if w.IsWatching() {
		if err := w.Close(); err != nil {
			return err
		}
	}

	lastIndex, err := w.getLastIndex()

	if err != nil {
		return err
	}

	w.lastIndex = lastIndex
	w.changeToken = reloadToken

	go w.watchConsul()
	return nil
}

func (w *consulWatcher) watchConsul() {
	w.isWatching = true
	w.ticker = time.NewTicker(time.Second * time.Duration(w.intervalSecs))

	ticker := w.ticker.C
	defer w.Close()

	for {
		select {
		case <-ticker:
			if w.isChanged() {
				w.changeToken.OnChanged()
				return
			}
		}
	}
}

func (w *consulWatcher) getLastIndex() (uint64, error) {
	_, meta, err := w.kv.Keys(w.prefix, "", nil)

	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return meta.LastIndex, nil
}

func (w *consulWatcher) isChanged() bool {

	lastIdx, err := w.getLastIndex()

	if err != nil {
		return false
	}

	if w.lastIndex != lastIdx {
		return true
	}

	return false
}

func (w *consulWatcher) IsWatching() bool {
	return w.isWatching
}

func (w *consulWatcher) Close() error {
	w.isWatching = false
	w.ticker.Stop()

	return nil
}
