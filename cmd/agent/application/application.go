package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/achrt/metrics-collector/cmd/agent/metrics"
	sc "github.com/achrt/metrics-collector/internal/controller"
	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/achrt/metrics-collector/internal/sender"
)

type App struct {
	reportInterval      int64
	reportTimerDuration int64
	metricServerAddress string
	reqTimeout          int

	sender *sender.Client
}

const reqTimeout = 2

func New() (*App, error) {
	cfg, err := loadConfiguration()
	if err != nil {
		return nil, err
	}
	return &App{
		reportInterval:      int64(cfg.ReportInterval),
		reportTimerDuration: int64(cfg.ReportInterval),
		sender:              sender.New(),
		metricServerAddress: "http://" + cfg.Address,
		reqTimeout:          reqTimeout,
	}, nil
}

func (a *App) Run(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	ctx2, cancel2 := context.WithCancel(ctx)

	monitor := metrics.New(a.reportTimerDuration)
	go monitor.Run(ctx2, cancel2)
	go sc.Run(ctx2, cancel2)
	go a.report(ctx2, cancel2, monitor)

	<-ctx2.Done()
}

func (a *App) report(ctx context.Context, cancel context.CancelFunc, monitor *metrics.Monitor) {
	defer cancel()

	var status int
	var m models.Metrics
	var err error

	metrics := monitor.MetricCodes()
	for {
		<-time.After(time.Duration(a.reportInterval) * time.Second)

		for _, metric := range metrics {
			m, err = monitor.MetricDataModel(metric)
			if err != nil {
				log.Println(err)
				return
			}
			// TODO: можно ассинхронно отправлять запросы;
			// не очень понятно, нужно ли что-то дополнительно делать с ctx2

			ctx2, cancel2 := context.WithCancel(ctx)

			url := fmt.Sprintf("%s/update/", a.metricServerAddress)

			status, err = sender.New().R().
				SetURL(url).
				SetTimeout(a.reqTimeout).
				SetBody(m).
				SetHeader("Content-Type", "text/plain").
				Post(ctx2, cancel2)

			if err != nil {
				log.Println(err)
			}

			if status != http.StatusOK {
				log.Println("request response status: ", status)
			}
		}
	}
}
