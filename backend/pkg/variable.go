package pkg

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	currentPath, _ = os.Getwd()
)
var (
	backendHost     = "backend"
	GoEchoPort      string
	GRPCPort        string
	DeliveryAddress string

	TZ string

	PostgresDSN string

	AssembleNumber int
)

func init() {
	err := godotenv.Load(filepath.Join(currentPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AssembleNumber, err = strconv.Atoi(os.Getenv("ASSEMBLE_NUMBER"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TZ = os.Getenv("TZ")
	GoEchoPort = os.Getenv("GO_ECHO_PORT")
	GRPCPort = os.Getenv("GRPC_PORT")
	DeliveryAddress = backendHost + ":" + GRPCPort

	PostgresDSN = "host=postgres" +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" port=" + os.Getenv("POSTGRES_BACK_PORT") +
		" TimeZone=" + TZ +
		" dbname=app sslmode=disable"

}
