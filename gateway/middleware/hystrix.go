package middleware

import (
	"errors"
	"net/http"

	"github.com/MiteshSharma/gateway/common/util"
	"github.com/MiteshSharma/gateway/gateway/model"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/felixge/httpsnoop"
	"go.uber.org/zap"
)

type HystrixMiddleware struct {
	service model.GatewayService
}

func NewHystrixMiddleware(service model.GatewayService) *HystrixMiddleware {
	hystrixMiddleware := &HystrixMiddleware{service: service}
	hystrixMiddleware.Init()
	return hystrixMiddleware
}

func (hm *HystrixMiddleware) Init() {
	services := hm.service.GetServiceNames()
	for index := 0; index < len(services); index++ {
		hystrix.ConfigureCommand(services[index], hystrix.CommandConfig{
			MaxConcurrentRequests: 100,
			Timeout:               60000,
		})
	}
}

func (hm *HystrixMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		serviceName := hm.service.GetServiceNameBy(r.RequestURI)
		err := hystrix.Do(serviceName, func() error {
			metrix := httpsnoop.CaptureMetrics(next, rw, r)
			if metrix.Code >= 500 && metrix.Code < 600 {
				commomUtil.Logger.Info("Received call in hystrix middleware ", zap.String("hystrixCommand", serviceName),
					zap.Int("response code", metrix.Code))
				return errors.New("Error received ")
			} else {
				commomUtil.Logger.Info("Received call in hystrix middleware ", zap.String("hystrixCommand", serviceName))
			}
			return nil
		}, func(err error) error {
			commomUtil.Logger.Error("Error during call in hystrix middleware ", zap.Error(err))
			return nil
		})
		if err != nil {
			commomUtil.Logger.Error("Error failed call in hystrix middleware ", zap.Error(err))
		}
	}
}
