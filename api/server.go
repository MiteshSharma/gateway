package api

import (
	"net/http"

	"github.com/MiteshSharma/gateway/util"
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
			utils.Logger.Error("Error starting server ", zap.Error(err))
			return
		}
	}()
}

func StopServer() {

}
