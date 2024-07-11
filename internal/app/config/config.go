package config

import (
	"github.com/caarlos0/env/v8"
)

type Config struct {
	BRIMerchantDbURL      string `env:"BRIMERCHANT_DBURL"`
	BRIMerchantDbUsername string `env:"BRIMERCHANT_DBUSER"`
	BRIMerchantDbPassword string `env:"BRIMERCHANT_DBPASSWORD"`
	BRIMerchantDbName     string `env:"BRIMERCHANT_DBNAME"`

	BaseUrl    string `env:"BASE_URL"`
	EsbBaseUrl string `env:"ESB_URL"`
}

func MustLoadConfig() *Config {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
