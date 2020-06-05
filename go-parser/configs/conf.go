package configs

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

//Configuration - application configuration
type Configuration struct {
	Db   DBConfig
	RMQ  RMQConfig
	Host HostConfig
}

//SetupConf - load configuration
func SetupConf() *Configuration {
	c := flag.String("c", "app.conf", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}

	return &Config
}
