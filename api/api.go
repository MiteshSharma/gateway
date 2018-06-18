package api

import (
	"github.com/MiteshSharma/gateway/middleware"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"package": "api",
	})
)

func InitApi() *negroni.Negroni {
	router := mux.NewRouter()
	InitProxy(router)
	n := negroni.New()
	n.UseFunc(middleware.NewZipkinMiddleware().GetMiddlewareHandler())
	n.UseFunc(middleware.NewLoggerMiddleware().GetMiddlewareHandler())
	n.UseFunc(middleware.NewHystrixMiddleware().GetMiddlewareHandler())
	n.UseHandler(router)
	return n
}
