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
	sync.RWMutex
	m map[string]models.Metrics

	castTicker   uint32 // кол-во секунд между вызовом Cast(); если 0, то сохранение при каждом обновлении
	saveToDisk   bool
	saveOnUpdate bool
	logFile      string
}

func New(filePath string, castTicker uint32) (s *Storage, err error) {
	s = &Storage{
		m:            map[string]models.Metrics{},
		castTicker:   castTicker,
		saveOnUpdate: castTicker == 0,
		saveToDisk:   filePath != "",
		logFile:      filePath,
	}
	if s.saveToDisk {

		if !s.saveOnUpdate {
			go s.writer()
		}
	}

	return
}

func (s *Storage) Get(code string) (*models.Metrics, error) {
	s.RLock()
	defer s.RUnlock()

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

	s.Lock()
	defer s.Unlock()

	return s.set(code, val)
}

func (s *Storage) set(code string, val models.Metrics) error {
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
		if err := s.cast(); err != nil {
			return err
		}
	}
	
	return nil
}

func (s *Storage) PrintMetrics() map[string]string {
	s.RLock()
	defer s.RUnlock()

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

	consumer, err := newConsumer(s.logFile)
	if err != nil {
		return err
	}
	defer consumer.close()

	m, err := consumer.read()
	if err != nil {
		return err
	}
	for _, mt := range m {
		s.set(mt.ID, *mt)
	}
	return nil
}

// Cast() выгружает данные из in-memory в файл
func (s *Storage) Cast() error {
	s.RLock()
	defer s.RUnlock()

	return s.cast()
}

func (s *Storage) cast() error {
	mtr := []models.Metrics{}
	for _, m := range s.m {
		mtr = append(mtr, m)
	}

	producer, err := newProducer(s.logFile)
	if err != nil {
		return err
	}

	defer producer.close()
	return producer.write(mtr)
}

func (s *Storage) writer() {
	for {
		<-time.After(time.Duration(s.castTicker) * time.Second)
		s.Cast()
	}
}
