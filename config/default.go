package config

type config struct {
	opts Options
}

func newConfig(opts ...Option) (Config, error) {
	var c config

	err := c.Init(opts...)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *config) Init(opts ...Option) error {
	c.opts = Options{}
	for _, o := range opts {
		o(&c.opts)
	}

	return nil
}

func (c *config) Options() Options {
	return c.opts
}

func (c *config) String() string {
	return "config"
}
