package commonMiddleware

import (
	"net/http"
	"time"

	"github.com/MiteshSharma/gateway/common"
	"github.com/felixge/httpsnoop"
	"go.uber.org/zap"
)

// LoggerMiddleware struct
type LoggerMiddleware struct {
}

// NewLoggerMiddleware function returns instance of logger middleware
func NewLoggerMiddleware() *LoggerMiddleware {
	loggerMiddleware := &LoggerMiddleware{}
	loggerMiddleware.Init()
	return loggerMiddleware
}

// Init function to init anything required for middleware
func (lm *LoggerMiddleware) Init() {
}

// GetMiddlewareHandler function returns middleware used to log requests
func (lm *LoggerMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		metrix := httpsnoop.CaptureMetrics(next, rw, r)
		common.ServerObj.Logger.Info("Request handling completed from logger middleware ", zap.String("Host", r.Host),
			zap.String("Method", r.Method), zap.String("Request", r.RequestURI), zap.String("RemoteAddress", r.RemoteAddr),
			zap.String("Referer", r.Referer()), zap.String("UserAgent", r.UserAgent()), zap.Int("StatusCode", metrix.Code),
			zap.Int("Duration", int(metrix.Duration/time.Millisecond)))
	}
}
