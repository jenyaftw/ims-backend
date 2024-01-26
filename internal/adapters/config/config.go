package config

type HttpConfig struct {
	Port int
}

type Config struct {
	Http *HttpConfig
}

func NewConfig() (*Config, error) {
	return &Config{}, nil
}
