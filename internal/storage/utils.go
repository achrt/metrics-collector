package storage

import (
	"encoding/json"
	"os"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/labstack/gommon/log"
)

type producer struct {
	file    *os.File
	encoder *json.Encoder
}

func newProducer(filename string) (*producer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (p *producer) write(m []models.Metrics) error {
	err := p.encoder.Encode(m)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (p *producer) close() error {
	return p.file.Close()
}

// *****************************************************************

type consumer struct {
	file    *os.File
	decoder *json.Decoder
}

func newConsumer(filename string) (*consumer, error) {
	var file *os.File
	var err error

	file, err = os.OpenFile(filename, os.O_RDONLY, 0400)
	if err != nil {
		log.Error(err)
		if file, err = os.Create(filename); err != nil {
			log.Error(err)
			return nil, err
		}
	}

	return &consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (c *consumer) read() ([]*models.Metrics, error) {
	m := []*models.Metrics{}
	if err := c.decoder.Decode(&m); err != nil {
		log.Error(err)
	}
	return m, nil
}

func (c *consumer) close() error {
	return c.file.Close()
}
