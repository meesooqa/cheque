package config

type LoadConfigFunc func(string) (*Conf, error)

type ConfigProvider interface {
	GetConf() (res *Conf, err error)
}

type DefaultConfigProvider struct {
	fname    string
	loadFunc LoadConfigFunc
}

func NewDefaultConfigProvider() *DefaultConfigProvider {
	return NewDefaultConfigProviderWithCustomLoader("etc/config.yml", load)
}

func NewDefaultConfigProviderWithCustomLoader(fname string, loader LoadConfigFunc) *DefaultConfigProvider {
	return &DefaultConfigProvider{
		fname:    fname,
		loadFunc: loader,
	}
}

// GetConf provides Conf from default config file
func (o *DefaultConfigProvider) GetConf() (res *Conf, err error) {
	return o.loadFunc(o.fname)
}
