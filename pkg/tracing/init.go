package tracing

import (
	"io"

	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type Tracing struct {
	tracer opentracing.Tracer
	closer io.Closer
}

var trace *Tracing

// Init returns an instance of Jaeger Tracer.
func Init(service string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, nil, errors.Wrap(err, "Cannot init Jaeger")
	}
	return tracer, closer, nil
}

func InitTracer() error {
	trace = new(Tracing)
	tracer, closer, err := Init(viper.GetString("Tracing.Name"))
	if err != nil {
		return err
	}
	opentracing.SetGlobalTracer(tracer)
	trace.tracer = tracer
	trace.closer = closer
	return nil
}

func CloseTracer() {
	trace.closer.Close()
}

// StartSpanFromRequest extracts the parent span context from the inbound HTTP request
// and starts a new child span if there is a parent span.
func StartSpanFromRequest(r *http.Request) opentracing.Span {
	tracer := trace.tracer
	spanCtx, _ := Extract(tracer, r)
	return tracer.StartSpan("HTTP "+r.Method+":"+r.URL.Path, ext.RPCServerOption(spanCtx))
}

// Inject injects the outbound HTTP request with the given span's context to ensure
// correct propagation of span context throughout the trace.
func Inject(span opentracing.Span, request *http.Request) error {
	return span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
}

// Extract extracts the inbound HTTP request to obtain the parent span's context to ensure
// correct propagation of span context throughout the trace.
func Extract(tracer opentracing.Tracer, r *http.Request) (opentracing.SpanContext, error) {
	return tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
}
