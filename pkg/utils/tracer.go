package utils

import (
	"context"
	"io"
	"time"

	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// NewTracer generate tracer
func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegerCfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory

	mObj := jaeger.NewMetrics(jMetricsFactory, nil)

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender, jaeger.ReporterOptions.Metrics(mObj))
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegerCfg.Logger(jLogger),
		jaegerCfg.Metrics(jMetricsFactory),
		jaegerCfg.Reporter(reporter),
	)
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)

	return tracer, closer, err
}

// ExtractContext extract metadata from tracing context and return new span
func ExtractContext(ctx context.Context, t opentracing.Tracer, name string) opentracing.Span {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := t.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	sp = opentracing.StartSpan(name, opentracing.ChildOf(wireContext))
	return sp
}
