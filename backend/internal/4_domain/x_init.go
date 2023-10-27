package domain

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var AssembleNumber int

func init() {
	currentPath, _ := os.Getwd()

	err := godotenv.Load(filepath.Join(currentPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AssembleNumber, err = strconv.Atoi(os.Getenv("ASSEMBLE_NUMBER"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
