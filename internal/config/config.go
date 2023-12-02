package config

type Config struct {
	Quantity  uint
	Hex       bool
	Uppercase bool
}

var conf = &Config{}

func Get() *Config {
	return conf
}
