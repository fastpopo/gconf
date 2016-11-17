package gconf

import (
	"time"
)

type FileInfo interface {
	Exists() bool
	IsDirectory() bool
	LastModified() time.Time
	GetLength() int64
	GetName() string
	GetPhysicalPath() string
	ReadAll() ([]byte, error)
}

// FileWatcher feature will be support in future.
type FileWatcher interface {
	Watch(filter string) ReloadToken
}

type FileConfSource interface {
	ConfSource
	GetFileInfo() FileInfo
	GetPath() string
	EndureIfNotExist() bool
	ReloadOnChange() bool
}

type FileConfProvider interface {
	ConfProvider
	LoadFromStream(stream []byte) error
}
