package gconf

type configurationBuilder struct {
	sources []ConfSource
}

func NewConfBuilder() ConfBuilder {
	return &configurationBuilder{}
}

func (c *configurationBuilder) Add(source ConfSource) ConfBuilder {
	if source == nil {
		return c
	}

	c.sources = append(c.sources, source)
	return c
}

func (c *configurationBuilder) GetSources() []ConfSource {
	return c.sources
}

func (c *configurationBuilder) Build() (ConfRoot, error) {

	var providers []ConfProvider

	for i := len(c.sources) - 1; i >= 0; i-- {
		provider, err := c.sources[i].Build(c)

		if err != nil {
			return nil, err
		}

		if provider == nil {
			continue
		}

		providers = append(providers, provider)
	}

	return newConfRoot(providers), nil
}
