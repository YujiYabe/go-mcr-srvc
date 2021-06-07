package pkg

import (
	"os"
	"path/filepath"
)

var (
	currentPath, _ = os.Getwd()
	StoragePath    = filepath.Join(currentPath, "storage")
	ErrorLogPath   = filepath.Join(StoragePath, "errorlog")
	LogPath        = filepath.Join(StoragePath, "log")
	RegisterPath   = "scripts/order/register"
	YummyPath      = filepath.Join(currentPath, "yummy")
)
