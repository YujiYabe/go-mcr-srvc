package env

import (
	"github.com/spf13/viper"
)

const (
	LOCAL = "local"
	DEV   = "dev"
	PROD  = "prod"
)

var AlcoholIndex = 99

type env struct {
	EnvDB
}

type EnvDB struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

var Env env

func init() {
	Env = env{
		EnvDB: EnvDB{},
	}

	initEnvDB()
}

// initEnvDBの追加
func initEnvDB() {
	v := viper.New()
	v.AutomaticEnv()

	Env.EnvDB = EnvDB{
		Host:     v.GetString("DB_HOST"),
		Port:     v.GetString("DB_PORT"),
		User:     v.GetString("DB_USER"),
		Password: v.GetString("DB_PASSWORD"),
		Name:     v.GetString("DB_NAME"),
	}
}
