package configs

import "fmt"

type RMQConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	ExchangeIn      string
	ExchangeTypeIn  string
	QueueIn         string
	RoutingKeyIn    string
	ExchangeOut     string
	ExchangeTypeOut string
	QueueOut        string
	RoutingKeyOut   string
	Concurrency     int
}

func BuildRMQConfig(cfg *Configuration) *RMQConfig {
	rmqConfig := RMQConfig{
		Host:            cfg.RMQ.Host,
		Port:            cfg.RMQ.Port,
		User:            cfg.RMQ.User,
		Password:        cfg.RMQ.Password,
		ExchangeIn:      cfg.RMQ.ExchangeIn,
		ExchangeTypeIn:  cfg.RMQ.ExchangeTypeIn,
		QueueIn:         cfg.RMQ.QueueIn,
		RoutingKeyIn:    cfg.RMQ.RoutingKeyIn,
		ExchangeOut:     cfg.RMQ.ExchangeOut,
		ExchangeTypeOut: cfg.RMQ.ExchangeTypeOut,
		QueueOut:        cfg.RMQ.QueueOut,
		RoutingKeyOut:   cfg.RMQ.RoutingKeyOut,
		Concurrency:     cfg.RMQ.Concurrency,
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
