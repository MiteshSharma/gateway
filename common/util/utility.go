package util

import (
	"bytes"
	"os"
	"path/filepath"
)

func findConfigFile(fileName string) string {
	if _, error := os.Stat("./../conf/" + fileName); error == nil {
		fileName, _ = filepath.Abs("./../conf/" + fileName)
	} else if _, error := os.Stat("./conf/" + fileName); error == nil {
		fileName, _ = filepath.Abs("./conf/" + fileName)
	}
	return fileName
}

// GetConfigFileContent returns config file content
func GetConfigFileContent(fileName string) string {
	filePath := findConfigFile(fileName)

	file, err := os.Open(filePath)

	if err != nil {
		panic("Error occured during config file reading " + err.Error())
	}

	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	return buf.String()
}
