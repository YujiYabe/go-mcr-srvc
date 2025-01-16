package pkg

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	currentPath, _ = os.Getwd()
)
var (
	backendHost = "backend"
	GoEchoPort  string
	GRPCPort    string
	GRPCAddress string

	TZ string

	PostgresDSN string
)

func init() {
	// env はAWS Secrets Manager で取得
	log.Println("== == == == == == == == == == ")
	log.Printf("%#v\n", currentPath)
	log.Println("== == == == == == == == == == ")

	err := godotenv.Load(
		filepath.Join(
			currentPath,
			"internal",
			"env",
			".localenv",
		),
	)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TZ = os.Getenv("TZ")
	GoEchoPort = os.Getenv("GO_ECHO_PORT")
	GRPCPort = os.Getenv("GRPC_PORT")
	GRPCAddress = backendHost + ":" + GRPCPort

	PostgresDSN = "host=postgres" +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" port=" + os.Getenv("POSTGRES_BACK_PORT") +
		" TimeZone=" + TZ +
		" dbname=app sslmode=disable"

}
