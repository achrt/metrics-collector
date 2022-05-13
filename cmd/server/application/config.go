package application

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address        string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
}

func loadConfiguration() (cfg Config, err error) {
	err = env.Parse(&cfg)
	return
}
