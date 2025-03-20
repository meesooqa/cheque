package config

type ConfigProvider interface {
	GetConf() (res *Conf, err error)
}

type DefaultConfigProvider struct {
	fname string
}

func NewDefaultConfigProvider() *DefaultConfigProvider {
	return &DefaultConfigProvider{fname: "etc/config.yml"}
}

// GetConf provides Conf from default config file
func (o *DefaultConfigProvider) GetConf() (res *Conf, err error) {
	return load(o.fname)
}
