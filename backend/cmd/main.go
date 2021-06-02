package main

import backend "backend/internal/1_infrastructure"

func main() {
	a := backend.NewApp()
	a.Start()
}
