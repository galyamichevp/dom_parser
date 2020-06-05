package main

import (
	"go-dom-parser/api/routes"
	"go-dom-parser/api/sockets"
	"go-dom-parser/configs"
	"strconv"
)

var err error

func main() {

	// load configuration
	cfg := configs.SetupConf()

	// Creating a connection to the database
	// configs.DB, err = gorm.Open("mysql", configs.DbURL(configs.BuildDBConfig(cfg)))

	// if err != nil {
	// 	fmt.Println("status: ", err)
	// }

	// defer configs.DB.Close()

	// run the migrations: todo struct
	// configs.DB.AutoMigrate(&models.Todo{})

	// === RMQ configuration ===

	data := []byte("hello world --- xxx")

	ch := sockets.SetupRMQ(cfg)

	ch.Publish(cfg, data)

	// === RMQ configuration ===

	// define routes
	router := routes.SetupRouter()

	// run server
	router.Run(":" + strconv.Itoa(cfg.Host.Port))
}
