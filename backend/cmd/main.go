package main

import (
	"fmt"
	"os"

	backend "backend/internal/1_framework"
)

func main() {
	app, err := backend.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := app.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
