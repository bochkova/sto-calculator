package main

import "sto-calculator/internal/app"

const (
	configPath = "./config.yaml"
)

func main() {
	myApp := app.NewApp(configPath)
	myApp.Run()
}
