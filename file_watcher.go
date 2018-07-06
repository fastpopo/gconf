package gconf

import (
	"errors"
	"fmt"
	"runtime"
	"sync"

	"gopkg.in/fsnotify/fsnotify.v1"
)

type fileWatcher struct {
	filePath    string
	watcher     *fsnotify.Watcher
	changeToken ChangeToken
	isWatching  bool
	done        chan struct{}
}

func NewFileWatcher(filePath string) (Watcher, error) {
	if filePath == "" {
		return nil, errors.New("invalid filePath for file watcher")
	}

	fileInfo := NewFileInfo(filePath)
	if !fileInfo.Exists() {
		return nil, errors.New(fmt.Sprintf("the file[%s] doesn't exist", fileInfo.GetPhysicalPath()))
	}

	fileWatcher := &fileWatcher{
		filePath:    filePath,
		watcher:     nil,
		changeToken: nil,
		isWatching:  false,
	}

	return fileWatcher, nil
}

func (f *fileWatcher) Watch(reloadToken ChangeToken) error {
	if f.IsWatching() {
		f.Close()
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return errors.New("can't create file watcher: " + err.Error())
	}

	f.watcher = watcher
	f.changeToken = reloadToken
	f.watcher.Add(f.filePath)
	f.done = make(chan struct{})

	if runtime.GOOS != "windows" {
		go f.fileWatching()
	} else {
		go f.fileWatchingForWindows()
	}

	return nil
}

func (f *fileWatcher) fileWatching() {
	f.isWatching = true

	defer f.Close()

	for {
		select {
		case <-f.done:
			return

		case <-f.watcher.Errors:
			break

		case event := <-f.watcher.Events:
			if len(event.Name) == 0 {
				break
			}

			if event.Op^fsnotify.Chmod == 0 {
				break
			}

			f.changeToken.OnChanged()
		}
	}
}

func (f *fileWatcher) fileWatchingForWindows() {
	isFired := false
	f.isWatching = true
	var mutex = &sync.Mutex{}

	defer f.Close()

	for f.isWatching {
		select {
		case <-f.done:
			return

		case <-f.watcher.Errors:
			break

		case event := <-f.watcher.Events:
			if len(event.Name) == 0 {
				break
			}

			if event.Op^fsnotify.Chmod == 0 {
				break
			}

			mutex.Lock()
			if !isFired {
				f.changeToken.OnChanged()
				isFired = true
			} else {
				isFired = false
			}
			mutex.Unlock()
		}
	}
}

func (f *fileWatcher) IsWatching() bool {
	return f.isWatching
}

func (f *fileWatcher) Close() error {
	f.isWatching = false
	f.done <- struct{}{}

	if _, ok := <-f.done; ok {
		close(f.done)
	}

	if f.watcher == nil {
		return nil
	}

	err := f.watcher.Close()
	f.watcher = nil

	return err
}
