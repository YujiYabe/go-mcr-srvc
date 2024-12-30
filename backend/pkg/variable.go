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
	StoragePath    = filepath.Join(currentPath, "storage")
	ErrorLogPath   = filepath.Join(StoragePath, "errorlog")
	LogPath        = filepath.Join(StoragePath, "log")
	YummyPath      = filepath.Join(currentPath, "yummy")

	WebPath   = filepath.Join(currentPath, "web")
	IndexPath = filepath.Join(WebPath, "*.html")
)
var (
	backendHost     = "backend"
	MobilePort      string
	PCPort          string
	DeliveryPort    string
	MonitorPort     string
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
	MobilePort = os.Getenv("MOBILE_PORT")
	PCPort = os.Getenv("PC_PORT")
	DeliveryPort = os.Getenv("DELIVERY_PORT")
	MonitorPort = os.Getenv("MONITOR_PORT")
	DeliveryAddress = backendHost + ":" + DeliveryPort

	PostgresDSN = "host=postgres" +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" port=" + os.Getenv("POSTGRES_BACK_PORT") +
		" TimeZone=" + TZ +
		" dbname=app sslmode=disable"

}
