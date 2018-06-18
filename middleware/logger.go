package middleware

import (
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
)

type LoggerMiddleware struct {
}

func NewLoggerMiddleware() *LoggerMiddleware {
	loggerMiddleware := &LoggerMiddleware{}
	loggerMiddleware.Init()
	return loggerMiddleware
}

func (lm *LoggerMiddleware) Init() {
}

func (lm *LoggerMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		metrixFields := map[string]interface{}{
			"Host":          r.Host,
			"Method":        r.Method,
			"Request":       r.RequestURI,
			"RemoteAddress": r.RemoteAddr,
			"Referer":       r.Referer(),
			"UserAgent":     r.UserAgent(),
		}

		metrix := httpsnoop.CaptureMetrics(next, rw, r)

		metrixFields["StatusCode"] = metrix.Code
		metrixFields["Duration"] = int(metrix.Duration / time.Millisecond)

		log.WithFields(metrixFields).Info("Request handling completed.")
	}
}
