package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"package": "utils",
	})
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

func loadConfigFromFile(fileName string, config ConfigParam) {
	filePath := findConfigFile(fileName)

	file, err := os.Open(filePath)

	if err != nil {
		log.WithField("err", err).Fatal("Error occured during config file reading.")
		panic("Error occured during config file reading " + err.Error())
	}

	jsonParser := json.NewDecoder(file)

	log.Debug("Reading config params from file")
	config.LoadConfigFromJsonParser(jsonParser)
	config.SaveDefaultConfigParams()
}
