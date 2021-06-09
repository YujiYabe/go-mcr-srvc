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

	TZ string

	MysqlDatabase     string
	MysqlUser         string
	MysqlPassword     string
	MysqlRootPassword string

	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresPort     string
	PostgresDSN      string

	MongoDatabase     string
	MongoRootUsername string
	MongoRootPassword string
)

func init() {
	err := godotenv.Load(filepath.Join(currentPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TZ = os.Getenv("TZ")
	MobilePort = os.Getenv("MOBILE_BACK_PORT")
	PCPort = os.Getenv("PC_BACK_PORT")
	DeliveryPort = os.Getenv("DELIVERY_BACK_PORT")
	MonitorPort = os.Getenv("MONITOR_BACK_PORT")
	DeliveryAddress = backendHost + DeliveryPort

	MysqlDatabase = os.Getenv("MYSQL_DATABASE")
	MysqlUser = os.Getenv("MYSQL_USER")
	MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	MysqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")

	PostgresDB = os.Getenv("POSTGRES_DB")

	MongoDatabase = os.Getenv("MONGO_INITDB_DATABASE")
	MongoRootUsername = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	MongoRootPassword = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	PostgresDSN = "host=postgres" +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" port=" + os.Getenv("POSTGRES_BACK_PORT") +
		" TimeZone=" + TZ +
		" dbname=app sslmode=disable"

}
