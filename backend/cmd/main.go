package main

import (
	backend "backend/internal/1_framework"
)

func main() {
	// log.SetFlags(0)
	backend.NewApp().Start()
}
