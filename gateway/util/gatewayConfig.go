package util

import (
	"encoding/json"

	"github.com/MiteshSharma/gateway/common/config"
	"github.com/MiteshSharma/gateway/common/util"
)

// GatewayConfig struct contains all config details needed for this service
type GatewayConfig struct {
	*config.Config
}

// Init function to initialize config variable fetched from file detail provided
func Init(fileName string) *GatewayConfig {
	var config = GatewayConfig{}

	fileContent := util.GetConfigFileContent(fileName)
	err := json.Unmarshal([]byte(fileContent), &config)

	if err != nil {
		panic("Config content not right")
	}
	return &config
}
