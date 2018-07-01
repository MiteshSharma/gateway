package main

import (
	"flag"

	"github.com/MiteshSharma/gateway/gateway/server"
)

var (
	configFileName string
	GatewayServer  *server.GatewayServer
)

func main() {
	parseCmdParams()
	GatewayServer = server.NewGatewayServer()
	GatewayServer.StartServer(configFileName)
}

func parseCmdParams() {
	flag.StringVar(&configFileName, "config", "gatewayConfig.json", "")
	flag.Parse()
}
