package main

import (
	"log"

	"github.com/nikitalier/tenderMonitoring/config"
	"github.com/nikitalier/tenderMonitoring/pkg/application"
	"github.com/nikitalier/tenderMonitoring/pkg/provider"
	"github.com/nikitalier/tenderMonitoring/pkg/repository"
	"github.com/nikitalier/tenderMonitoring/pkg/service"
)

func main() {
	var appConfig config.Config

	err := appConfig.Load()
	if err != nil {
		log.Panic(err)
	}

	prov := provider.New(&appConfig.DataBase[0])

	err = prov.Open()
	if err != nil {
		log.Panic(err)
	}

	rep := repository.New(prov.GetCon())
	rep.PingDB()
	log.Println("DB connected")

	serv := service.New(rep)

	app := application.New(&application.Options{Serv: appConfig.ServerOpt, Svc: serv})

	app.Start()
}
