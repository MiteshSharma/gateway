package middleware

import (
	"net/http"

	"github.com/MiteshSharma/gateway/util"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"go.uber.org/zap"
)

const (
	debug         = false
	sameSpan      = false
	traceID128Bit = true // Tracer generate 128 bit traceID
)

type ZipkinMiddleware struct {
	ServiceHostPort string
	ServiceName     string
	tracer          opentracing.Tracer
}

func NewZipkinMiddleware() *ZipkinMiddleware {
	zipkinMiddleware := &ZipkinMiddleware{
		ServiceHostPort: "127.0.0.1:9411",
		ServiceName:     "gateway",
	}
	zipkinMiddleware.Init()
	return zipkinMiddleware
}

func (zm *ZipkinMiddleware) Init() {
	zipkinHTTPEndpoint := zm.ServiceHostPort + "/api/v1/spans"
	utils.Logger.Info("Init zipkin middleware ", zap.String("zipkinEndpoint", zipkinHTTPEndpoint))
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		utils.Logger.Error("Creating collector failed in zipkin middleware ", zap.Error(err))
		panic("Creating tracer collector failed")
	}
	recorder := zipkin.NewRecorder(collector, debug, zm.ServiceHostPort, zm.ServiceName)
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)
	if err != nil {
		utils.Logger.Error("Creating tracer failed ", zap.Error(err))
		panic("Creating tracer failed")
	}
	opentracing.InitGlobalTracer(tracer)
	zm.tracer = tracer
}

func (zm *ZipkinMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		wireContext, err := zm.tracer.Extract(
			opentracing.TextMap,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
		if err != nil {
			utils.Logger.Debug("Error encountered while trying to extract span ", zap.Error(err))
		}
		span := zm.tracer.StartSpan(r.URL.Path, ext.RPCServerOption(wireContext))
		defer span.Finish()
		ctx := opentracing.ContextWithSpan(r.Context(), span)
		r = r.WithContext(ctx)
		next(rw, r)
	}
}
