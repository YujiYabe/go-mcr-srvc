package main

import backend "backend/internal/1_framework"

func main() {
	a := backend.NewApp()
	a.Start()
}
