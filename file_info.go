package gconf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type fileInfo struct {
	exists       bool
	physicalPath string
	fileInfo     os.FileInfo
}

func NewFileInfo(filePath string) FileInfo {
	f := &fileInfo{
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

func (f *fileInfo) Exists() bool {
	return f.exists
}

func (f *fileInfo) IsDirectory() bool {
	if f.exists == false {
		return false
	}

	return f.fileInfo.IsDir()
}

func (f *fileInfo) LastModified() time.Time {
	if f.exists == false {
		return time.Time{}
	}

	return f.fileInfo.ModTime()
}

func (f *fileInfo) GetLength() int64 {
	if f.exists == false {
		return 0
	}

	return f.fileInfo.Size()
}

func (f *fileInfo) GetName() string {
	if f.exists == false {
		return ""
	}

	return f.fileInfo.Name()
}

func (f *fileInfo) GetPhysicalPath() string {
	return f.physicalPath
}

func (f *fileInfo) ReadAll() ([]byte, error) {
	if f.exists == false {
		return nil, errors.New(fmt.Sprintf("the file [%s] is not exist.", f.physicalPath))
	}

	return ioutil.ReadFile(f.physicalPath)
}
