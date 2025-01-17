package env

import (
	"log"

	"github.com/spf13/viper"
)

const (
	LOCAL = "local"
	DEV   = "dev"
	PROD  = "prod"
)

var (
	backendHost = "backend"
	GoEchoPort  string
	GRPCPort    string
	GRPCAddress string

	TZ string

	PostgresDSN string
)

var AlcoholIndex = 99

func init() {
	v := viper.New()
	v.AutomaticEnv()
	v.AddConfigPath("internal/env")
	v.SetConfigName(".localenv")
	v.SetConfigType("env")
	if err := v.ReadInConfig(); err != nil {
		log.Println("== == == == == == == == == == ")
		log.Printf("%#v\n", err)
		log.Println("== == == == == == == == == == ")

		// Continue even if file doesn't exist
	}

	TZ = v.GetString("TZ")
	GoEchoPort = v.GetString("GO_ECHO_PORT")
	GRPCPort = v.GetString("GRPC_PORT")
	GRPCAddress = backendHost + ":" + GRPCPort

	PostgresDSN = "host=postgres" +
		" user=" + v.GetString("POSTGRES_USER") +
		" password=" + v.GetString("POSTGRES_PASSWORD") +
		" port=" + v.GetString("POSTGRES_BACK_PORT") +
		" dbname=app" +
		" TimeZone=" + v.GetString("TZ") +
		" sslmode=disable"

}
