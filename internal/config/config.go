package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type AppConf struct {
	Route   Route   `json:"route"`
	Storage Storage `json:"storage"`
}

func New(path string) *AppConf {
	configFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file config: %s", err.Error())
	}
	defer configFile.Close()

	byteValue, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatalf("read file config: %s", err.Error())
	}

	var c AppConf

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		log.Fatalf("unmarshal file config: %s", err.Error())
	}

	if c.Route.validate() != nil {
		log.Fatalf("validate: %s", err.Error())
	}

	if c.Storage.validate() != nil {
		log.Fatalf("validate: %s", err.Error())
	}

	return &c
}
