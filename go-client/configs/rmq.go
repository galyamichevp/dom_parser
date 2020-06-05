package configs

import "fmt"

type RMQConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Exchange     string
	ExchangeType string
	Queue        string
	RoutingKey   string
	Concurrency  int
}

func BuildRMQConfig(cfg *Configuration) *RMQConfig {
	rmqConfig := RMQConfig{
		Host:         cfg.RMQ.Host,
		Port:         cfg.RMQ.Port,
		User:         cfg.RMQ.User,
		Password:     cfg.RMQ.Password,
		Exchange:     cfg.RMQ.Exchange,
		ExchangeType: cfg.RMQ.ExchangeType,
		Queue:        cfg.RMQ.Queue,
		RoutingKey:   cfg.RMQ.RoutingKey,
		Concurrency:  cfg.RMQ.Concurrency,
	}

	return &rmqConfig
}

func RMQURL(rqmConfig *RMQConfig) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		rqmConfig.User,
		rqmConfig.Password,
		rqmConfig.Host,
		rqmConfig.Port,
	)
}
