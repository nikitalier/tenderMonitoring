package main

import (
	"log"
	golog "log"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/nikitalier/tenderMonitoring/config"
	"github.com/nikitalier/tenderMonitoring/pkg/application"
	"github.com/nikitalier/tenderMonitoring/pkg/elastic"
	"github.com/nikitalier/tenderMonitoring/pkg/provider"
	"github.com/nikitalier/tenderMonitoring/pkg/repository"
	"github.com/nikitalier/tenderMonitoring/pkg/service"
)

func main() {
	var (
		appConfig config.Config
		logger    zerolog.Logger
	)

	err := appConfig.Load()
	if err != nil {
		golog.Fatalf("%v", err)
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "logLevel"

	if appConfig.Env != "local" {
		es, err := elastic.New("logs"+strings.ToLower(appConfig.AppName), appConfig.AppName, &appConfig.Logging)
		if err != nil {
			golog.Fatalf("%v", err)
		}
		logger = zerolog.New(es).With().Timestamp().CallerWithSkipFrameCount(2).Logger()
	} else {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().CallerWithSkipFrameCount(2).Timestamp().Logger()
	}

	prov := provider.New(&appConfig.DataBase[0], logger)

	err = prov.Open()
	if err != nil {
		log.Panic(err)
	}

	rep := repository.New(prov.GetCon(), logger)
	rep.PingDB()
	logger.Info().Msg("DB connected")

	serv := service.New(rep, &logger)

	app := application.New(&application.Options{Serv: appConfig.ServerOpt, Svc: serv, Logger: logger})

	app.Start()
}
