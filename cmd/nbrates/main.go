package main

import (
	"nbrates/internal/app"
	"nbrates/internal/config"
)

var confPath = "config/config.json"

func main() {
	conf := config.New(confPath)

	app.New(conf)
}
