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

func (c *configurationBuilder) Build() ConfRoot {

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
