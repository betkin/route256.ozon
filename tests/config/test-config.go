package test_config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIHost       string `envconfig:"API_HOST" default:"127.0.0.1"`
	APIPort       string `envconfig:"API_PORT" default:"8080"`
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
	return fmt.Sprintf("http://%s:%s", cfg.APIHost, cfg.APIPort)
}

func GetGrpcURL(cfg Config) string {
	return fmt.Sprintf("%s:%s", cfg.APIHost, cfg.GRPCPort)
}
