package gconf

type _ConfigurationBuilder struct {
	sources []ConfSource
}

func NewConfBuilder() ConfBuilder {
	return &_ConfigurationBuilder{}
}

func (c *_ConfigurationBuilder) Add(source ConfSource) ConfBuilder {
	if source == nil {
		return c
	}

	c.sources = append(c.sources, source)
	return c
}

func (c *_ConfigurationBuilder) GetSources() []ConfSource {
	return c.sources
}

func (c *_ConfigurationBuilder) Build() ConfRoot {

	var providers []ConfProvider

	for i := len(c.sources) - 1; i >= 0; i-- {
		provider := c.sources[i].Build(c)
		if provider == nil {
			continue
		}

		providers = append(providers, provider)
	}

	return newConfRoot(providers)
}
