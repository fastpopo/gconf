package gconf

// FileWatcher feature will be support in future.
type _FileWatcher struct {
	filePath    string
	reloadToken ReloadToken
}

func NewFileWatcher() (FileWatcher, error) {
	return &_FileWatcher{}, nil
}

func (f *_FileWatcher) Watch(filePath string) ReloadToken {
	f.reloadToken = NewReloadToken()
	return f.reloadToken
}
