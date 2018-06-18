package api

import (
	"github.com/MiteshSharma/gateway/middleware"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
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
