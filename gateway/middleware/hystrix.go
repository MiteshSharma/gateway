package middleware

import (
	"errors"
	"net/http"

	"github.com/MiteshSharma/gateway/common"
	"github.com/MiteshSharma/gateway/gateway/model"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/felixge/httpsnoop"
	"go.uber.org/zap"
)

// HystrixMiddleware struct
type HystrixMiddleware struct {
	service model.GatewayService
}

// NewHystrixMiddleware function returns instance of hystrix middleware
func NewHystrixMiddleware(service model.GatewayService) *HystrixMiddleware {
	hystrixMiddleware := &HystrixMiddleware{service: service}
	hystrixMiddleware.Init()
	return hystrixMiddleware
}

// Init function to init anything required for middleware
func (hm *HystrixMiddleware) Init() {
	service := hm.service.GetServiceName()
	hystrix.ConfigureCommand(service, hystrix.CommandConfig{
		MaxConcurrentRequests: 100,
		Timeout:               60000,
	})
}

// GetMiddlewareHandler function returns middleware used to hystrix requests
func (hm *HystrixMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		serviceName := hm.service.GetServiceNameBy(r.RequestURI)
		err := hystrix.Do(serviceName, func() error {
			metrix := httpsnoop.CaptureMetrics(next, rw, r)
			if metrix.Code >= 500 && metrix.Code < 600 {
				common.ServerObj.Logger.Info("Received call in hystrix middleware ", zap.String("hystrixCommand", serviceName),
					zap.Int("response code", metrix.Code))
				return errors.New("Error received ")
			} else {
				common.ServerObj.Logger.Info("Received call in hystrix middleware ", zap.String("hystrixCommand", serviceName))
			}
			return nil
		}, func(err error) error {
			common.ServerObj.Logger.Error("Error during call in hystrix middleware ", zap.Error(err))
			return nil
		})
		if err != nil {
			common.ServerObj.Logger.Error("Error failed call in hystrix middleware ", zap.Error(err))
		}
	}
}
