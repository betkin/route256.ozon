package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

// This data structure describes the configuration of the test object

type Config struct {
	ApiHost       string `envconfig:"API_HOST" default:"127.0.0.1"`
	ApiPort       string `envconfig:"API_PORT" default:"8080"`
	GRPCPort      string `envconfig:"GRPC_PORT" default:"8082"`
	LivecheckPort string `envconfig:"LIVECHECK_PORT" default:"8000"`
	LivecheckURI  string `envconfig:"LIVECHECK_URI" default:"live"`
}

func GetConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	return config, err
}

func GetApiURL(cfg Config) string {
	return fmt.Sprintf("http://%s:%s", cfg.ApiHost, cfg.ApiPort)
}

func GetGrpcURL(cfg Config) string {
	return fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.GRPCPort)
}
