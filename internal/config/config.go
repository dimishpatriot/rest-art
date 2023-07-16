package config

import (
	"sync"

	"github.com/dimishpatriot/rest-art-of-development/internal/logging"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
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
		Port       string `yaml:"port"`
		Database   string `yaml:"database"`
		Collection string `yaml:"collection"`
		Username   string `env:"MONGO_USER" env-default:"user"`
		Password   string `env:"MONGO_PASSWORD" env-default:"password"`
	} `yaml:"storage" env-required:"true"`
}

var once sync.Once

func GetConfig() *Config {
	var err error
	var cfg *Config

	once.Do(func() {
		logger := logging.GetLogger()

		if err = godotenv.Load(); err != nil {
			logger.Fatalf("can't read .env config file: %s", err)
		}

		cfg = &Config{}
		if err = cleanenv.ReadConfig("config.yaml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			logger.Info(help)
			logger.Fatalf("can't read yaml config file: %s", err)
		}

		showConfigWithSecret(logger, cfg)
	})
	return cfg
}

func showConfigWithSecret(logger *logging.Logger, cfg *Config) {
	cfgPrivatePass := *cfg
	if cfgPrivatePass.Storage.Username != "" {
		cfgPrivatePass.Storage.Username = "some-user"
	} else {
		cfgPrivatePass.Storage.Username = "empty"
	}

	if cfgPrivatePass.Storage.Password != "" {
		cfgPrivatePass.Storage.Password = "*********"
	} else {
		cfgPrivatePass.Storage.Password = "empty"
	}

	logger.Infof("[OK] config created: %+v", cfgPrivatePass)
}
