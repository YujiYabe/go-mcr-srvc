package pkg

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	currentPath, _ = os.Getwd()
	StoragePath    = filepath.Join(currentPath, "storage")
	ErrorLogPath   = filepath.Join(StoragePath, "errorlog")
	LogPath        = filepath.Join(StoragePath, "log")
	YummyPath      = filepath.Join(currentPath, "yummy")
	RegisterPath   = filepath.Join(currentPath, "scripts", "order", "register")
	ReservedPath   = filepath.Join(RegisterPath, "reserved")

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
)

func init() {
	err := godotenv.Load(filepath.Join(currentPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MobilePort = os.Getenv("MOBILE_BACK_PORT")
	PCPort = os.Getenv("PC_BACK_PORT")
	DeliveryPort = os.Getenv("DELIVERY_BACK_PORT")
	MonitorPort = os.Getenv("MONITOR_BACK_PORT")
	DeliveryAddress = backendHost + DeliveryPort
}
