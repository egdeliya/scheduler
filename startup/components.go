package startup

import (
	"github.com/caarlos0/env/v6"
)

func LoadConfig() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
