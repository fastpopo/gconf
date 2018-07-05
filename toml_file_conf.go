package gconf

import (
	"errors"
	"fmt"
)

type TomlFileConfSource struct {
	path                  string
	endureIfNotExist      bool
	onConfChangedCallback func(ConfChanges)
}

func NewTomlFileConfSource(path string) *TomlFileConfSource {
	return &TomlFileConfSource{
		path:                  path,
		endureIfNotExist:      false,
		onConfChangedCallback: nil,
	}
}

func (s *TomlFileConfSource) Build(builder ConfBuilder) (ConfProvider, error) {
	return NewFileConfProvider(s)
}

func (s *TomlFileConfSource) Load() (map[string]interface{}, error) {
	fileInfo := s.GetFileInfo()

	var data map[string]interface{}

	if !fileInfo.Exists() {
		if s.endureIfNotExist {
			return data, nil
		} else {
			return nil, errors.New(fmt.Sprintf("can't find the toml configuration file[%s]", fileInfo.GetPhysicalPath()))
		}
	}

	stream, err := fileInfo.ReadAll()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't read the toml configuration file[%s], err: %s", fileInfo.GetPhysicalPath(), err.Error()))
	}

	if stream == nil || len(stream) == 0 {
		return data, nil
	}

	parser := newTomlConfParser(RootPath, PathDelimiter)
	err = parser.Parse(stream)

	if err != nil {
		return nil, err
	}

	return parser.GetDataMap(), nil
}

func (s *TomlFileConfSource) SetEndureIfNotExist(endureIfNotExist bool) FileConfSource {
	s.endureIfNotExist = endureIfNotExist
	return s
}

func (s *TomlFileConfSource) SetOnConfChangedCallback(onConfChangedCallback func(ConfChanges)) FileConfSource {
	s.onConfChangedCallback = onConfChangedCallback
	return s
}

func (s *TomlFileConfSource) GetOnConfChangedCallback() func(ConfChanges) {
	return s.onConfChangedCallback
}

func (s *TomlFileConfSource) GetFileInfo() FileInfo {
	return NewFileInfo(s.path)
}

func (s *TomlFileConfSource) GetFilePath() string {
	return s.path
}

func (s *TomlFileConfSource) IsEndureIfNotExist() bool {
	return s.endureIfNotExist
}

func (s *TomlFileConfSource) IsFileExist() bool {
	return s.GetFileInfo().Exists()
}
