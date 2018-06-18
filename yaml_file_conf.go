package gconf

import (
	"errors"
	"fmt"
)

type YamlFileConfSource struct {
	path                  string
	endureIfNotExist      bool
	onConfChangedCallback func(ConfChanges)
}

func NewYamlFileConfSource(path string) *YamlFileConfSource {
	return &YamlFileConfSource{
		path:                  path,
		endureIfNotExist:      false,
		onConfChangedCallback: nil,
	}
}

func (s *YamlFileConfSource) Build(builder ConfBuilder) ConfProvider {
	return NewFileConfProvider(s)
}

func (s *YamlFileConfSource) Load() (map[string]interface{}, error) {
	fileInfo := s.GetFileInfo()

	var data map[string]interface{}

	if !fileInfo.Exists() {
		if s.endureIfNotExist {
			return data, nil
		} else {
			return nil, errors.New(fmt.Sprintf("can't find the yaml configuration file[%s]", fileInfo.GetPhysicalPath()))
		}
	}

	stream, err := fileInfo.ReadAll()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't read the yaml configuration file[%s], err: %s", fileInfo.GetPhysicalPath(), err.Error()))
	}

	if stream == nil || len(stream) == 0 {
		return data, nil
	}

	parser := newYamlConfParser(RootPath, KeyDelimiter)
	err = parser.Parse(stream)

	if err != nil {
		return nil, err
	}

	return parser.GetDataMap(), nil
}

func (s *YamlFileConfSource) SetEndureIfNotExist(endureIfNotExist bool) FileConfSource {
	s.endureIfNotExist = endureIfNotExist
	return s
}

func (s *YamlFileConfSource) SetOnConfChangedCallback(onConfChangedCallback func(ConfChanges)) FileConfSource {
	s.onConfChangedCallback = onConfChangedCallback
	return s
}

func (s *YamlFileConfSource) GetOnConfChangedCallback() func(ConfChanges) {
	return s.onConfChangedCallback
}

func (s *YamlFileConfSource) GetFileInfo() FileInfo {
	return NewFileInfo(s.path)
}

func (s *YamlFileConfSource) GetFilePath() string {
	return s.path
}

func (s *YamlFileConfSource) IsEndureIfNotExist() bool {
	return s.endureIfNotExist
}

func (s *YamlFileConfSource) IsFileExist() bool {
	return s.GetFileInfo().Exists()
}
