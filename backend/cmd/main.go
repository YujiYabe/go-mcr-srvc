package main

import backend "backend/internal/1_frameworks_drivers"

func main() {
	a := backend.NewApp()
	a.Start()
}
