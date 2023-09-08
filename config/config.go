package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Greeting string
}

func LoadConfig() (*Config, error) {
	conf := &Config{}

	if _, err := toml.DecodeFile("config.toml", conf); err != nil {
		return nil, err
	}

	if envGreeting := os.Getenv("GREETING"); envGreeting != "" {
		conf.Greeting = envGreeting
	}

	return conf, nil
}
