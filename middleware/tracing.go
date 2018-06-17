package middleware

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"package": "middleware",
	})
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
	log.Info("Creating HTTP collector with endpoint " + zipkinHTTPEndpoint)
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		log.WithField("err", err).Fatal("Creating collector failed.")
		panic("Creating tracer collector failed")
	}
	log.Info("Creating HTTP recorder")
	recorder := zipkin.NewRecorder(collector, debug, zm.ServiceHostPort, zm.ServiceName)
	log.Info("Creating HTTP tracer")
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)
	if err != nil {
		log.WithField("err", err).Fatal("Creating tracer failed.")
		panic("Creating tracer failed")
	}
	log.Info("Setting our tracer as default tracer.")
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
			log.WithField("err", err).Debug("Error encountered while trying to extract span.")
		}
		span := zm.tracer.StartSpan(r.URL.Path, ext.RPCServerOption(wireContext))
		defer span.Finish()
		ctx := opentracing.ContextWithSpan(r.Context(), span)
		r = r.WithContext(ctx)
		next(rw, r)
	}
}
