package config

type Config struct {
	AppName   string
	Version   string
	Quantity  uint
	Hex       bool
	Uppercase bool
}

var conf = &Config{
	AppName:   "guidgen",
	Version:   "2.0.1",
	Quantity:  1,
	Hex:       false,
	Uppercase: false,
}

func Get() *Config {
	return conf
}
