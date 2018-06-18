package utils

import (
	"log"
	"path"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	config := generateConfig(".")
	config.Level.SetLevel(zap.DebugLevel)
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal(err)
	}
	Logger = logger
	defer Logger.Sync()
}

func generateConfig(dir string) zap.Config {
	config := zap.NewProductionConfig()
	destination := path.Join(dir, "log.json")
	config.OutputPaths = []string{"stderr", destination}
	config.ErrorOutputPaths = []string{"stderr", destination}
	return config
}
