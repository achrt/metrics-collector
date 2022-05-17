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

func (c *Config) loadConfiguration() error {

	fAdd := flag.String("a", dAddress, "server address - host:port")
	fReportInetrval := flag.Int("r", dReportInterval, "an interval between metrics sendidng")
	fPollInetrval := flag.Int("p", dPollInterval, "an interval between metrics fetching")

	flag.Parse()

	if err := env.Parse(c); err != nil {
		return err
	}

	log.Println(c)

	if c.Address == dAddress && fAdd != nil && *fAdd != dAddress {
		c.Address = *fAdd
	}
	if c.ReportInterval == dReportInterval && fReportInetrval != nil && *fReportInetrval != dReportInterval {
		c.ReportInterval = uint32(*fReportInetrval)
	}
	if c.PollInterval == dPollInterval && fPollInetrval != nil && *fPollInetrval != dPollInterval {
		c.PollInterval = uint32(*fPollInetrval)
	}

	log.Println(c)

	return nil
}
