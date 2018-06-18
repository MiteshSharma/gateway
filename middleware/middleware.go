package middleware

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"package": "middleware",
	})
)

type Middleware interface {
	Init()
	GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc)
}
