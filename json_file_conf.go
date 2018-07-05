package gconf

import (
	"errors"
	"fmt"
)

type JsonFileConfSource struct {
	path                  string
	endureIfNotExist      bool
	onConfChangedCallback func(ConfChanges)
}

func NewJsonFileConfSource(path string) *JsonFileConfSource {
	return &JsonFileConfSource{
		path:                  path,
		endureIfNotExist:      false,
		onConfChangedCallback: nil,
	}
}

func (s *JsonFileConfSource) Build(builder ConfBuilder) (ConfProvider, error) {
	return NewFileConfProvider(s)
}

func (s *JsonFileConfSource) Load() (map[string]interface{}, error) {
	fileInfo := s.GetFileInfo()

	var data map[string]interface{}

	if !fileInfo.Exists() {
		if s.endureIfNotExist {
			return data, nil
		} else {
			return nil, errors.New(fmt.Sprintf("can't find the json configuration file[%s]", fileInfo.GetPhysicalPath()))
		}
	}

	stream, err := fileInfo.ReadAll()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't read the json configuration file[%s], err: %s", fileInfo.GetPhysicalPath(), err.Error()))
	}

	if stream == nil || len(stream) == 0 {
		return data, nil
	}

	parser := newJsonConfParser(RootPath, PathDelimiter)
	err = parser.Parse(stream)

	if err != nil {
		return nil, err
	}

	return parser.GetDataMap(), nil
}

func (s *JsonFileConfSource) SetEndureIfNotExist(endureIfNotExist bool) FileConfSource {
	s.endureIfNotExist = endureIfNotExist
	return s
}

func (s *JsonFileConfSource) SetOnConfChangedCallback(onConfChangedCallback func(ConfChanges)) FileConfSource {
	s.onConfChangedCallback = onConfChangedCallback
	return s
}

func (s *JsonFileConfSource) GetOnConfChangedCallback() func(ConfChanges) {
	return s.onConfChangedCallback
}

func (s *JsonFileConfSource) GetFileInfo() FileInfo {
	return NewFileInfo(s.path)
}

func (s *JsonFileConfSource) GetFilePath() string {
	return s.path
}

func (s *JsonFileConfSource) IsEndureIfNotExist() bool {
	return s.endureIfNotExist
}

func (s *JsonFileConfSource) IsFileExist() bool {
	return s.GetFileInfo().Exists()
}
