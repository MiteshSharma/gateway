package commomUtil

import (
	"log"
	"path"

	"github.com/MiteshSharma/gateway/common/model"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger(service model.Service) {
	config := generateConfig(".")
	config.Level.SetLevel(zap.DebugLevel)
	config.InitialFields = map[string]interface{}{
		"service": service.GetServiceName(),
	}
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
