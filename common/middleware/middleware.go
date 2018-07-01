package commonMiddleware

import (
	"net/http"
)

// Middleware interface for all middlewares
type Middleware interface {
	Init()
	GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc)
}
