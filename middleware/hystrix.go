package middleware

import (
	"errors"
	"net/http"

	"github.com/MiteshSharma/gateway/util"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/felixge/httpsnoop"
	"go.uber.org/zap"
)

type HystrixMiddleware struct {
}

func NewHystrixMiddleware() *HystrixMiddleware {
	hystrixMiddleware := &HystrixMiddleware{}
	hystrixMiddleware.Init()
	return hystrixMiddleware
}

func (hm *HystrixMiddleware) Init() {
	services := utils.GetAllServiceNames()
	for index := 0; index < len(services); index++ {
		hystrix.ConfigureCommand(services[index], hystrix.CommandConfig{
			MaxConcurrentRequests: 100,
			Timeout:               60000,
		})
	}
}

func (hm *HystrixMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		serviceName := utils.GetServiceName(r.RequestURI)
		err := hystrix.Do(serviceName, func() error {
			metrix := httpsnoop.CaptureMetrics(next, rw, r)
			if metrix.Code >= 500 && metrix.Code < 600 {
				utils.Logger.Info("Received call in hystrix middleware ", zap.String("hystrixCommand", serviceName),
					zap.Int("response code", metrix.Code))
				return errors.New("Error received ")
			} else {
				utils.Logger.Info("Received call in hystrix middleware ", zap.String("hystrixCommand", serviceName))
			}
			return nil
		}, func(err error) error {
			utils.Logger.Error("Error during call in hystrix middleware ", zap.Error(err))
			return nil
		})
		if err != nil {
			utils.Logger.Error("Error failed call in hystrix middleware ", zap.Error(err))
		}
	}
}
