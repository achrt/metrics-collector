package application

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address        string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	ReportInterval uint32 `env:"REPORT_INTERVAL" envDefault:"10"`
	PollInterval   uint32 `env:"POLL_INTERVAL" envDefault:"2"`
}

func loadConfiguration() (cfg Config, err error) {
	err = env.Parse(&cfg)
	return
}
