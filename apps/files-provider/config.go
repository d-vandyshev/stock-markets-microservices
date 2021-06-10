package main

import (
	"fmt"
	"os"
)

type Config struct {
	DataPath    string
	RabbitmqUrl string
}

func (c *Config) SetFromEnvOrDie() {
	c.DataPath = c.GetEnv("DATA_PATH")
	c.RabbitmqUrl = c.GetEnv("RABBITMQ_URL")
}

func (c Config) GetEnv(param string) string {
	value := os.Getenv(param)
	if value == "" {
		fmt.Printf("ERROR. Environment var %s is empty", param)
		os.Exit(2)
	}
	return value
}
