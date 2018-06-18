package utils

import (
	"encoding/json"

	"go.uber.org/zap"
)

type GatewayConfig struct {
	ServerConfig ServerConfig
}

type ServerConfig struct {
	Port string
}

var GatewayConfigParam GatewayConfig = GatewayConfig{}

func (o GatewayConfig) SaveDefaultConfigParams() {
	if o.ServerConfig.Port == "" {
		o.ServerConfig.Port = ":8080"
	}
}

func (o GatewayConfig) LoadConfigFromJsonParser(jsonParser *json.Decoder) {
	if jsonErr := jsonParser.Decode(&GatewayConfigParam); jsonErr != nil {
		Logger.Error("Json parsing error: ", zap.Error(jsonErr))
	}
}

func LoadConfig(fileName string) {
	loadConfigFromFile(fileName, GatewayConfigParam)
}
