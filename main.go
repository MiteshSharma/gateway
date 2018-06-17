package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/MiteshSharma/gateway/api"
	"github.com/MiteshSharma/gateway/util"
	"github.com/Sirupsen/logrus"
)

var (
	configFileName string
	log            = logrus.WithFields(logrus.Fields{
		"package": "main",
	})
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	parseCmdParams()
	log.Debug("Command param read.")
	utils.LoadConfig(configFileName)
	log.Debug("Configration read.")
	api.InitServer()
	log.Debug("Init server done.")
	api.StartServer()
	log.Debug("Server started.")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	api.StopServer()
	log.Debug("Server stopped.")
}

func parseCmdParams() {
	flag.StringVar(&configFileName, "config", "config.json", "")
	flag.Parse()
}
