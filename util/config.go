package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ServerConfig ServerConfig
}

type ServerConfig struct {
	Port string
}

var ConfigParam *Config = &Config{}

func findConfigFile(fileName string) string {
	if _, error := os.Stat("./" + fileName); error == nil {
		fileName, _ = filepath.Abs("./" + fileName)
	} else if _, error := os.Stat("./config/" + fileName); error == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, error := os.Stat("./src/github.com/MiteshSharma/gateway/config/" + fileName); error == nil {
		fileName, _ = filepath.Abs("./src/github.com/MiteshSharma/gateway/config/" + fileName)
	}
	return fileName
}

func (o *Config) SaveDefaultConfigParams() {
	if o.ServerConfig.Port == "" {
		o.ServerConfig.Port = ":8080"
	}
}

func LoadConfig(fileName string) {
	filePath := findConfigFile(fileName)

	file, error := os.Open(filePath)

	if error != nil {
		panic("Error occured during config file reading " + error.Error())
	}

	jsonParser := json.NewDecoder(file)

	config := Config{}

	if jsonErr := jsonParser.Decode(&config); jsonErr != nil {
		panic("Json parsing error" + jsonErr.Error())
	}

	config.SaveDefaultConfigParams()

	ConfigParam = &config
}
