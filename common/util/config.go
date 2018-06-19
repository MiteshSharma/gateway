package commomUtil

import (
	"encoding/json"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

type ConfigParam interface {
	LoadConfigFromJsonParser(d *json.Decoder)
	SaveDefaultConfigParams()
}

func findConfigFile(fileName string) string {
	if _, error := os.Stat("./" + fileName); error == nil {
		fileName, _ = filepath.Abs("./" + fileName)
	} else if _, error := os.Stat("./config/" + fileName); error == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	}
	return fileName
}

func LoadConfigFromFile(fileName string, config ConfigParam) {
	filePath := findConfigFile(fileName)

	file, err := os.Open(filePath)

	if err != nil {
		Logger.Error("Error occured during config file reading.", zap.Error(err))
		panic("Error occured during config file reading " + err.Error())
	}

	jsonParser := json.NewDecoder(file)

	Logger.Debug("Reading config params from file")
	config.LoadConfigFromJsonParser(jsonParser)
	config.SaveDefaultConfigParams()
}
