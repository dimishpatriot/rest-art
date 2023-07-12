package config

import (
	"sync"

	"github.com/dimishpatriot/rest-art-of-development/internal/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen" env-required:"true"`
	Storage struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"storage_port"`
		Database   string `yaml:"database"`
		Collection string `yaml:"collection"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
	} `yaml:"storage" env-required:"true"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}

		logger.Debugf("config created: %+v", instance)
	})
	return instance
}
