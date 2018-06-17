package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"package": "api",
	})
)

func InitApi() *mux.Router {
	router := mux.NewRouter()
	InitProxy(router)
	return router
}
