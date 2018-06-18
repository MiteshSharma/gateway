package middleware

import (
	"net/http"
)

type Middleware interface {
	Init()
	GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc)
}
