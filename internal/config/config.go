package config

type Config struct {
	AppName   string
	Version   string
	Quantity  uint
	Hex       bool
	Uppercase bool
}

var conf = &Config{
	AppName: "guidgen",
	Version: "2.0.0",
}

func Get() *Config {
	return conf
}
