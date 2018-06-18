package middleware

import (
	"net/http"
	"time"

	"github.com/MiteshSharma/gateway/util"
	"github.com/felixge/httpsnoop"
	"go.uber.org/zap"
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

		metrix := httpsnoop.CaptureMetrics(next, rw, r)
		utils.Logger.Info("Request handling completed from logger middleware ", zap.String("Host", r.Host),
			zap.String("Method", r.Method), zap.String("Request", r.RequestURI), zap.String("RemoteAddress", r.RemoteAddr),
			zap.String("Referer", r.Referer()), zap.String("UserAgent", r.UserAgent()), zap.Int("StatusCode", metrix.Code),
			zap.Int("Duration", int(metrix.Duration/time.Millisecond)))
	}
}
