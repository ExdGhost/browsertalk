package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var configModel *Model

// Init ...
func Init() {

	env := os.Getenv("TIER")
	path := "config/tier" + env + ".json"

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to bind config")
	}

	json.Unmarshal(bytes, configModel)
}

// Get ...
func Get() *Model {
	return configModel
}
