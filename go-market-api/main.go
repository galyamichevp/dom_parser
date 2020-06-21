package main

import (
	"go-dom-parser/api/controllers"
	"go-dom-parser/api/routes"
	"go-dom-parser/api/sockets"
	"go-dom-parser/configs"
	"go-dom-parser/core"
	"go-dom-parser/domain"
	"go-dom-parser/scheduler"
	"strconv"
)

var err error

func main() {

	// load configuration
	cfg := configs.SetupConf()

	// init domain container
	storage := &domain.Storage{}
	storage.Init()

	// setup socket
	socket := sockets.SetupSocket(cfg)

	// setup scheduler
	scheduler := scheduler.SetupScheduler(storage, socket)
	scheduler.RunJobs()

	// setup processor
	processor := core.SetupProcessor(storage, socket)
	processor.RunProcessor()

	// setup controller
	controller := controllers.SetupController(storage)

	// setup routes
	router := routes.SetupRouter(controller)

	// run rest-api server
	router.Run(":" + strconv.Itoa(cfg.Host.Port))
}
