package config

import (
	"context"
	"os"
	"testWorkmate/internal/common/types/error_with_codes"
	http_config "testWorkmate/internal/config/http"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Http      http_config.HttpConfig `yaml:"http"`
	path      string
	envReader envReader
}

func (c *Config) GetHttp() *http_config.HttpConfig {
	return &c.Http
}

type envReader interface {
	EnvReadConfig(addr string, cfg interface{}) error
}

func MustLoad(ctx context.Context, configPath string, envReader envReader) *Config {
	operation := "config.MustLoad()"

	cfg := new(Config)
	cfg.envReader = envReader

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		logrus.WithFields(logrus.Fields{
			"config_path": configPath,
		}).WithError(err).Fatal(error_with_codes.ErrorFailedToFindConfig.SetOperation(operation)) // set operation on custom error
	}

	err = envReader.EnvReadConfig(configPath, cfg)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"config_path": configPath,
		}).WithError(err).Fatal(error_with_codes.ErrorFailedToReadConfig.SetOperation(operation)) // set operation on custom error
	}

	return cfg
}
