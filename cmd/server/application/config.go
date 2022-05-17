package application

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

const (
	dAddress       = "127.0.0.1:8080"
	dRestore       = true
	dStoreInterval = 10
	dStFile        = "/tmp/devops-metrics-db.json"
)

type Config struct {
	Address       string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	StoreFile     string `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db-ee.json"`
	StoreInterval uint32 `env:"STORE_INTERVAL" envDefault:"10"`
	Restore       bool   `env:"RESTORE" envDefault:"true"`
}

func (c *Config) loadConfiguration() error {
	fAdd := flag.String("a", dAddress, "host:port")
	fRestore := flag.Bool("r", dRestore, "restore previous metrics")
	fStInetrval := flag.Int("i", dStoreInterval, "an interval between metrics storing")
	fStFile := flag.String("f", dStFile, "storage file address")
	flag.Parse()

	if err := env.Parse(c); err != nil {
		return err
	}

	log.Println(c)

	if c.Address == dAddress && fAdd != nil && *fAdd != dAddress {
		c.Address = *fAdd
	}
	if c.StoreFile == dStFile && fStFile != nil && *fStFile != dStFile {
		c.StoreFile = *fStFile
	}
	if c.StoreInterval == dStoreInterval && fStInetrval != nil && *fStInetrval != dStoreInterval {
		c.StoreInterval = uint32(*fStInetrval)
	}
	if c.Restore && fRestore != nil && !*fRestore {
		c.Restore = *fRestore
	}

	log.Println(c)
	return nil
}
