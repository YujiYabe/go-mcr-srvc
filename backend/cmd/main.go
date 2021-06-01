package main

import app "app/internal/1_infrastructure"

func main() {
	a := app.NewApp()
	a.Start()
}
