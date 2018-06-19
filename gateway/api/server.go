package api

import (
	"net/http"

	"github.com/MiteshSharma/gateway/common/util"
	"github.com/MiteshSharma/gateway/gateway/util"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
)

type Server struct {
	Router *negroni.Negroni
}

var ServerObj *Server

func InitServer() {
	ServerObj = &Server{}
	ServerObj.Router = InitApi()
}

func StartServer() {
	go func() {
		err := http.ListenAndServe(utils.GatewayConfigParam.ServerConfig.Port, ServerObj.Router)
		if err != nil {
			commomUtil.Logger.Error("Error starting server ", zap.Error(err))
			return
		}
	}()
}

func StopServer() {

}
