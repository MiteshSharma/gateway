package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/MiteshSharma/gateway/common"
	"github.com/MiteshSharma/gateway/gateway/api"
	"github.com/MiteshSharma/gateway/gateway/util"
	"go.uber.org/zap"
)

type GatewayServer struct {
	*common.Server
	GatewayConfig *util.GatewayConfig
}

func NewGatewayServer() *GatewayServer {
	return &GatewayServer{}
}

func (s *GatewayServer) StartServer(configFileName string) {
	s.GatewayConfig = util.Init(configFileName)
	s.Server = &common.Server{}
	s.Server.Init(s.GatewayConfig.Config)

	common.SetServerObj(s.Server)

	go func() {
		err := http.ListenAndServe("localhost:8081", api.InitApi(s.Router))
		if err != nil {
			s.Logger.Error("Error starting server ", zap.Error(err))
			return
		}
	}()

	s.Logger.Debug("Started server.")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	s.Logger.Debug("Stopped server.")
}
