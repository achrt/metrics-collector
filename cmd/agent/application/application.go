package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/achrt/metrics-collector/cmd/agent/metrics"
	sc "github.com/achrt/metrics-collector/internal/controller"
	"github.com/achrt/metrics-collector/internal/sender"
)

type App struct {
	reportInterval      int64
	reportTimerDuration int64
	metricServerAddress string
	reqTimeout          int

	sender *sender.Client
}

func New(reportInterval, duration int64, metricServerAddress string) *App {
	return &App{
		reportInterval:      reportInterval,
		reportTimerDuration: duration,
		sender:              sender.New(),
		reqTimeout:          2,
		metricServerAddress: metricServerAddress,
	}
}

func (a *App) Run(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	ctx2, cancel2 := context.WithCancel(ctx)

	monitor := metrics.New(a.reportTimerDuration)
	go monitor.Run(ctx2, cancel2)
	go sc.Run(ctx2, cancel2)
	go a.Report(ctx2, cancel2, monitor)

	<-ctx2.Done()
}

func (a *App) Report(ctx context.Context, cancel context.CancelFunc, monitor *metrics.Monitor) {
	defer cancel()

	var mType, name, value string
	var err error
	var status int

	metrics := monitor.MetricCodes()
	for {
		<-time.After(time.Duration(a.reportInterval) * time.Second)

		for _, metric := range metrics {
			mType, name, value, err = monitor.MetricData(metric)
			if err != nil {
				log.Println(err)
				return
			}

			// TODO: можно ассинхронно отправлять запросы;
			// не очень понятно, нужно ли что-то дополнительно делать с ctx2

			ctx2, cancel2 := context.WithCancel(ctx)

			url := fmt.Sprintf("%s/update/%s/%s/%s", a.metricServerAddress, mType, name, value)

			status, err = sender.New().R().
				SetURL(url).
				SetTimeout(a.reqTimeout).
				SetHeader("Content-Type", "text/plain").
				Post(ctx2, cancel2)

			if err != nil {
				log.Println(err)
			}

			if status != http.StatusAccepted {
				log.Println("request response status: ", status)
			}
		}
	}
}
