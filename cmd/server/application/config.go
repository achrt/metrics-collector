package application

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address       string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	StoreFile     string `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	StoreInterval uint32 `env:"STORE_INTERVAL" envDefault:"300"`
	Restore       bool   `env:"RESTORE" envDefault:"true"`
}

func loadConfiguration() (cfg Config, err error) {
	err = env.Parse(&cfg)
	return
}
