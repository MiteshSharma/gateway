package api

import (
	"net/http"

	"github.com/MiteshSharma/gateway/util"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

var ServerObj *Server

func InitServer() {
	ServerObj = &Server{}
	ServerObj.Router = InitApi()
}

func StartServer() {
	go func() {
		err := http.ListenAndServe(utils.ConfigParam.ServerConfig.Port, ServerObj.Router)
		if err != nil {
			log.WithField("err", err).Fatal("Server starting failed.")
			return
		}
	}()
}

func StopServer() {

}
