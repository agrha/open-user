package server

import (
	"os"

	"github.com/spf13/viper"
)

const (
	DevEnv        = "development"
	StagingEnv    = "staging"
	ProductionEnv = "production"
)

var (
	Debug bool

	BRIMerchantDbURL      string
	BRIMerchantDbUsername string
	BRIMerchantDbPassword string
	BRIMerchantDbName     string

	BaseUrl    string
	EsbBaseUrl string
)

func GetAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return DevEnv
	}

	return env
}

func InitTest() {
	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Debug = true

	BRIMerchantDbURL = viper.GetString(`database.brimerchant.url`)
	BRIMerchantDbUsername = viper.GetString(`database.brimerchant.user`)
	BRIMerchantDbPassword = viper.GetString(`database.brimerchant.password`)
	BRIMerchantDbName = viper.GetString(`database.brimerchant.name`)

	BaseUrl = viper.GetString(`base_url`)
	EsbBaseUrl = viper.GetString(`esb_url`)

}
