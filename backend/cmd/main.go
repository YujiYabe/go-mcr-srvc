package main

import backend "backend/internal/1_frameworks_driver"

func main() {
	a := backend.NewApp()
	a.Start()
}
