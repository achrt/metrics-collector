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

func loadConfiguration() (cfg Config, err error) {
	fAdd := flag.String("a", dAddress, "host:port")
	fRestore := flag.Bool("r", dRestore, "restore previous metrics")
	fStInetrval := flag.Int("i", dStoreInterval, "an interval between metrics storing")
	fStFile := flag.String("f", dStFile, "storage file address")
	flag.Parse()

	if err = env.Parse(&cfg); err != nil {
		return
	}

	log.Println(cfg)

	if cfg.Address == dAddress && fAdd != nil && *fAdd != dAddress {
		cfg.Address = *fAdd
	}
	if cfg.StoreFile == dStFile && fStFile != nil && *fStFile != dStFile {
		cfg.StoreFile = *fStFile
	}
	if cfg.StoreInterval == dStoreInterval && fStInetrval != nil && *fStInetrval != dStoreInterval {
		cfg.StoreInterval = uint32(*fStInetrval)
	}
	if cfg.Restore && fRestore != nil && !*fRestore {
		cfg.Restore = *fRestore
	}

	log.Println(cfg)

	return cfg, nil
}
