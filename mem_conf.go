package gconf

type MemConfSource struct {
}

func NewMapConfSource(conf map[string]interface{}) *MemConfSource {
	return &MemConfSource{}
}

func (s *MemConfSource) Build(builder ConfBuilder) ConfProvider {
	return NewConfProvider(s)
}

func (s *MemConfSource) Load() (map[string]interface{}, error) {
	data := make(map[string]interface{})

	return data, nil
}
