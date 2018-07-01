package api

import (
	"github.com/MiteshSharma/gateway/common/middleware"
	"github.com/MiteshSharma/gateway/gateway/middleware"
	"github.com/MiteshSharma/gateway/gateway/model"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitApi(router *mux.Router) *negroni.Negroni {
	InitProxy(router)
	n := negroni.New()
	n.UseFunc(commonMiddleware.NewZipkinMiddleware().GetMiddlewareHandler())
	n.UseFunc(commonMiddleware.NewLoggerMiddleware().GetMiddlewareHandler())
	n.UseFunc(middleware.NewHystrixMiddleware(model.NewGatewayService()).GetMiddlewareHandler())
	n.UseHandler(router)
	return n
}
