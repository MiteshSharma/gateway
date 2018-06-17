package middleware

import (
	"errors"
	"net/http"

	"github.com/MiteshSharma/gateway/util"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/felixge/httpsnoop"
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
			log.WithField("response code", metrix.Code).Debug("Request received")
			if metrix.Code >= 500 && metrix.Code < 600 {
				log.WithField("hystrixCommand", serviceName).
					WithField("StatusCode", metrix.Code).Warn("Backend failure for service: " + serviceName)
				return errors.New("Error received ")
			}
			return nil
		}, func(err error) error {
			log.WithError(err).WithField("hystrixCommand", serviceName).Warn("hystrix error handling")
			return nil
		})
		if err != nil {
			log.WithError(err).WithField("hystrixCommand", serviceName).Error("hystrix request return with error")
		}
	}
}
