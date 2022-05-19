package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/achrt/metrics-collector/internal/domain/models"
)

type producer struct {
	file    *os.File
	encoder *json.Encoder
}

func newProducer(filename string) (*producer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	return &producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (p *producer) write(m []models.Metrics) error {
	err := p.encoder.Encode(m)
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
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		if file, err = os.Create(filename); err != nil {
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
	c.decoder.Decode(&m)
	return m, nil
}

func (c *consumer) close() error {
	return c.file.Close()
}
