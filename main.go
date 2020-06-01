package main

import (
	"fmt"
	"go-dom-parser/api/routes"
	"go-dom-parser/configs"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// load configuration
	cfg := configs.SetupConf()

	// Creating a connection to the database
	configs.DB, err = gorm.Open("mysql", configs.DbURL(configs.BuildDBConfig(cfg)))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer configs.DB.Close()

	// run the migrations: todo struct
	configs.DB.AutoMigrate(&models.Todo{})

	// define routes
	router := routes.SetupRouter()

	// run server
	router.Run()
}
