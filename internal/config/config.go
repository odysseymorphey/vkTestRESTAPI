package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	cfg *koanf.Koanf
}

func NewConfig(configPath string) (*Config, error) {
	cfg := koanf.New(".")

	err := cfg.Load(file.Provider(configPath), yaml.Parser())
	if err != nil {
		return nil, err
	}

	return &Config{
		cfg: cfg,
	}, nil
}

func (c *Config) PostgresDSN() string {
	return c.cfg.String("storage.postgres")
}

func (c *Config) ServerPort() int {
	return c.cfg.Int("server.port")
}
