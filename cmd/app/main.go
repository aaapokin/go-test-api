package main

import (
	config "app/configs"
	"app/internal/app"
	"app/internal/app/routes"
	"flag"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "path", "configs/.env", "path to env")
}

func main() {
	flag.Parse()
	config.SetConfig()
	router := routes.SetupRouter()
	api := app.New(router)
	api.Start()
}
