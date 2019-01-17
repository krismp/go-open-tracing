package tracing

import (
	"fmt"
	"io"

	jaeger "github.com/uber/jaeger-client-go"

	"github.com/uber/jaeger-client-go/config"

	opentracing "github.com/opentracing/opentracing-go"
)

func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: canot init jaeger: %v\n", err))
	}
	return tracer, closer
}
