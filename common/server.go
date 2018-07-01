package common

import (
	"fmt"
	"log"
	"os"

	"github.com/MiteshSharma/gateway/common/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var ServerObj *Server

// Server struct is base struct for any kind of server
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
	Logger *zap.Logger
}

func SetServerObj(serverObj *Server) {
	ServerObj = serverObj
}

// Init function is used to initialize router, logger and database to be used
func (s *Server) Init(config *config.Config) {
	s.Router = mux.NewRouter()
	s.Logger = s.getLogger(config)
	s.DB = s.getDb(config)
}

func (s *Server) getDb(config *config.Config) *gorm.DB {
	var db *gorm.DB
	switch config.DatabaseConfig.DbType {
	case "mysql":
		mysqlDb, err := gorm.Open("mysql", s.getMysqlURL(config))
		if err != nil {
			s.Logger.Fatal("Connecting mysql failed due to error ", zap.Error(err))
			os.Exit(1)
		}
		db = mysqlDb
		break
	default:
		break
	}
	return db
}

func (s *Server) getMysqlURL(config *config.Config) string {
	return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseConfig.UserName, config.DatabaseConfig.Password,
		config.DatabaseConfig.DbName)
}

func (s *Server) getLogger(config *config.Config) *zap.Logger {
	loggerConfig := generateConfig(config)
	loggerConfig.Level.SetLevel(zap.DebugLevel)
	loggerConfig.InitialFields = map[string]interface{}{
		"service": config.ServiceName,
	}
	logger, err := loggerConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	return logger
}

func generateConfig(appConfig *config.Config) zap.Config {
	loggerConfig := zap.NewProductionConfig()
	if (appConfig != nil) && ((config.LoggerConfig{}) != appConfig.LoggerConfig) {
		loggerConfig.OutputPaths = []string{"stderr", appConfig.LoggerConfig.LogFilePath}
		loggerConfig.ErrorOutputPaths = []string{"stderr", appConfig.LoggerConfig.LogFilePath}
	}
	return loggerConfig
}
