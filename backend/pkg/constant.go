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
	YummyPath      = filepath.Join(currentPath, "yummy")
	RegisterPath   = filepath.Join(currentPath, "scripts", "order", "register")
	WebPath        = filepath.Join(currentPath, "web")
	IndexPath      = filepath.Join(WebPath, "*.html")
)
