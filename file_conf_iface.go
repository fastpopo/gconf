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

type FileConfSource interface {
	ConfSource
	SetEndureIfNotExist(bool) FileConfSource
	SetOnConfChangedCallback(func(ConfChanges)) FileConfSource
	GetOnConfChangedCallback() func(ConfChanges)
	GetFileInfo() FileInfo
	GetFilePath() string
	IsEndureIfNotExist() bool
	IsFileExist() bool
}

type FileConfProvider interface {
	ConfProvider
	OnChanged()
}
