package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/MiteshSharma/gateway/api"
	"github.com/MiteshSharma/gateway/util"
)

var (
	configFileName string
)

func main() {
	utils.InitLogger()
	utils.Logger.Debug("Starting server.")
	parseCmdParams()
	utils.LoadConfig(configFileName)
	api.InitServer()
	api.StartServer()

	utils.Logger.Debug("Started server.")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	api.StopServer()

	utils.Logger.Debug("Stopped server.")
}

func parseCmdParams() {
	flag.StringVar(&configFileName, "config", "config.json", "")
	flag.Parse()
}
