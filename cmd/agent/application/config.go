package application

import (
	"flag"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

const (
	dAddress        = "127.0.0.1:8080"
	dReportInterval = 10 * time.Second
	dPollInterval   = 2 * time.Second
)

type Config struct {
	Address        string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
}

func (c *Config) loadConfiguration() error {

	fAdd := flag.String("a", dAddress, "server address - host:port")
	fReportInetrval := flag.Duration("r", dReportInterval, "an interval between metrics sendidng")
	fPollInetrval := flag.Duration("p", dPollInterval, "an interval between metrics fetching")

	flag.Parse()

	if err := env.Parse(c); err != nil {
		return err
	}

	log.Println(c)

	if c.Address == dAddress && fAdd != nil && *fAdd != dAddress {
		c.Address = *fAdd
	}
	if c.ReportInterval == dReportInterval && fReportInetrval != nil && *fReportInetrval != dReportInterval {
		c.ReportInterval = *fReportInetrval
	}
	if c.PollInterval == dPollInterval && fPollInetrval != nil && *fPollInetrval != dPollInterval {
		c.PollInterval = *fPollInetrval
	}

	log.Println(c)

	return nil
}
