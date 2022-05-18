package application

import (
	"flag"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

const (
	dAddress       = "127.0.0.1:8080"
	dRestore       = true
	dStoreInterval = 300 * time.Second
	dStFile        = "/tmp/devops-metrics-db.json"
)

type Config struct {
	Address       string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	Restore       bool          `env:"RESTORE"`
}

func (c *Config) loadConfiguration() error {
	fAdd := flag.String("a", dAddress, "host:port")

	fStInetrval := flag.Duration("i", dStoreInterval, "an interval between metrics storing")
	fStFile := flag.String("f", dStFile, "storage file address")

	flag.BoolVar(&c.Restore, "r", true, "restore data from a file in -f")
	flag.Parse()

	if err := env.Parse(c); err != nil {
		return err
	}
	if c.Address == dAddress && fAdd != nil && *fAdd != dAddress {
		c.Address = *fAdd
	}

	if c.StoreFile == dStFile && fStFile != nil && *fStFile != dStFile {
		c.StoreFile = *fStFile
	}
	if c.StoreInterval == dStoreInterval && fStInetrval != nil && *fStInetrval != dStoreInterval {
		c.StoreInterval = *fStInetrval
	}

	log.Println(c)
	return nil
}
