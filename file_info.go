package gconf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type _FileInfo struct {
	exists       bool
	physicalPath string
	fileInfo     os.FileInfo
}

func NewFileInfo(filePath string) FileInfo {
	f := &_FileInfo{
		exists:       false,
		physicalPath: filePath,
	}

	fileInfo, err := os.Stat(filePath)

	if err == nil {
		f.exists = true
		f.fileInfo = fileInfo
	}

	return f
}

func (f *_FileInfo) Exists() bool {
	return f.exists
}

func (f *_FileInfo) IsDirectory() bool {
	if f.exists == false {
		return false
	}

	return f.fileInfo.IsDir()
}

func (f *_FileInfo) LastModified() time.Time {
	if f.exists == false {
		return time.Time{}
	}

	return f.fileInfo.ModTime()
}

func (f *_FileInfo) GetLength() int64 {
	if f.exists == false {
		return 0
	}

	return f.fileInfo.Size()
}

func (f *_FileInfo) GetName() string {
	if f.exists == false {
		return ""
	}

	return f.fileInfo.Name()
}

func (f *_FileInfo) GetPhysicalPath() string {
	return f.physicalPath
}

func (f *_FileInfo) ReadAll() ([]byte, error) {
	if f.exists == false {
		return nil, errors.New(fmt.Sprintf("the file [%s] is not exist.", f.physicalPath))
	}

	return ioutil.ReadFile(f.physicalPath)
}
