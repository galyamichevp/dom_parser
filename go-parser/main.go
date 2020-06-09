package main

import (
	"fmt"
	"go-dom-parser/api/routes"
	"go-dom-parser/api/sockets"
	"go-dom-parser/configs"
	"go-dom-parser/core"
	"go-dom-parser/domain"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error

func main() {

	// load configuration
	cfg := configs.SetupConf()

	// === DB configuration ===
	// Creating a connection to the database

	x := configs.DbURL(configs.BuildDBConfig(cfg))

	configs.DB, err = gorm.Open("mysql", x)

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer configs.DB.Close()

	// run the migrations: todo struct
	configs.DB.AutoMigrate(&domain.Resource{})
	// === DB configuration ===

	// === Processor ===

	p := core.New()
	p.Run()

	// === Processor ===

	// === RMQ configuration ===

	ch := sockets.SetupRMQ(cfg)

	ch.AddProcessor("test", p.ProcessorChan)

	ch.Subscribe(cfg)

	// === RMQ configuration ===

	// define routes
	router := routes.SetupRouter()

	// run server
	router.Run(":" + strconv.Itoa(cfg.Host.Port))
}
