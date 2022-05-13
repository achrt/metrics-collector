package storage

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/achrt/metrics-collector/internal/domain/models"
)

type Storage struct {
	mMutex *sync.RWMutex
	m      map[string]models.Metrics

	castTicker   uint32 // кол-во секунд между вызовом Cast(); если 0, то сохранение при каждом обновлении
	saveToDisk   bool
	saveOnUpdate bool

	*producer
	*consumer
}

func New(filePath string, castTicker uint32) (s *Storage, err error) {
	s = &Storage{
		m:            map[string]models.Metrics{},
		mMutex:       &sync.RWMutex{},
		castTicker:   castTicker,
		saveOnUpdate: castTicker == 0,
		saveToDisk:   filePath != "",
	}
	if s.saveToDisk {
		if s.producer, err = newProducer(filePath); err != nil {
			return
		}
		if s.consumer, err = newConsumer(filePath); err != nil {
			return
		}

		if !s.saveOnUpdate {
			go s.writer()
		}
	}

	return
}

func (s *Storage) Get(code string) (*models.Metrics, error) {
	code = strings.ToLower(code)
	if val, ok := s.m[code]; ok {
		return &val, nil
	}
	return nil, errors.New("unknown metric code")
}

func (s *Storage) Set(code string, val models.Metrics) error {
	if code == "" {
		return errors.New("code is an empty string")
	}

	s.mMutex.RLock()
	defer s.mMutex.RUnlock()

	code = strings.ToLower(code)

	if val.MType == models.TypeCounter {
		if val.Delta == nil {
			return errors.New("val.Delta can not be nil")
		}
		if s.m[code].Delta == nil {
			s.m[code] = val
			return nil
		}
		*s.m[code].Delta += *val.Delta
	}

	if val.MType == models.TypeGauge {
		if val.Value == nil {
			return errors.New("val.Value can not be nil")
		}
		s.m[code] = val
	}

	if s.saveToDisk && s.saveOnUpdate {
		if err := s.Cast(); err != nil {
			return err
		}
	}
	return nil
}

func (s Storage) PrintMetrics() map[string]string {
	res := map[string]string{}
	for code, val := range s.m {
		if val.Delta != nil {
			res[code] = fmt.Sprintf("%v", *val.Delta)
		}
		if val.Value != nil {
			res[code] = fmt.Sprintf("%v", *val.Value)
		}
	}
	return res
}

// Load() загружает метрики из файла в in-memory хранилище
func (s *Storage) Load() error {
	m, err := s.consumer.read()
	if err != nil {
		return err
	}
	for _, mt := range m {
		s.Set(mt.ID, *mt)
	}
	return nil
}

// Cast() выгружает данные из in-memory в файл
func (s Storage) Cast() error {
	mtr := []models.Metrics{}
	for _, m := range s.m {
		mtr = append(mtr, m)
	}
	return s.producer.write(mtr)
}

func (s Storage) Close() {
	if s.saveToDisk {
		s.consumer.close()
		s.producer.close()
	}
}

func (s Storage) writer() {
	for {
		<-time.After(time.Duration(s.castTicker) * time.Second)
		s.Cast()
	}
}
