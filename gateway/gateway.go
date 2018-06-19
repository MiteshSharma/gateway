package gateway

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/MiteshSharma/gateway/common/util"
	"github.com/MiteshSharma/gateway/gateway/api"
	"github.com/MiteshSharma/gateway/gateway/model"
	"github.com/MiteshSharma/gateway/gateway/util"
)

var (
	configFileName string
)

func Server() {
	commomUtil.InitLogger(model.NewGatewayService())
	parseCmdParams()
	utils.LoadConfig(configFileName)
	api.InitServer()
	api.StartServer()

	commomUtil.Logger.Debug("Started server.")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	api.StopServer()
	commomUtil.Logger.Debug("Stopped server.")
}

func parseCmdParams() {
	flag.StringVar(&configFileName, "config", "gatewayConfig.json", "")
	flag.Parse()
}
