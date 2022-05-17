package application

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

const (
	dAddress        = "127.0.0.1:8080"
	dReportInterval = 10
	dPollInterval   = 2
)

type Config struct {
	Address        string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	ReportInterval uint32 `env:"REPORT_INTERVAL" envDefault:"10"`
	PollInterval   uint32 `env:"POLL_INTERVAL" envDefault:"2"`
}

func loadConfiguration() (cfg Config, err error) {

	fAdd := flag.String("a", dAddress, "server address - host:port")
	fReportInetrval := flag.Int("r", dReportInterval, "an interval between metrics sendidng")
	fPollInetrval := flag.Int("p", dPollInterval, "an interval between metrics fetching")

	flag.Parse()

	if err = env.Parse(&cfg); err != nil {
		return
	}

	log.Println(cfg)

	if cfg.Address == dAddress && fAdd != nil && *fAdd != dAddress {
		cfg.Address = *fAdd
	}
	if cfg.ReportInterval == dReportInterval && fReportInetrval != nil && *fReportInetrval != dReportInterval {
		cfg.ReportInterval = uint32(*fReportInetrval)
	}
	if cfg.PollInterval == dPollInterval && fPollInetrval != nil && *fPollInetrval != dPollInterval {
		cfg.PollInterval = uint32(*fPollInetrval)
	}

	log.Println(cfg)

	return
}
