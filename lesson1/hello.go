package main

import (
	"fmt"
	"os"

	"github.com/krismp/buat_main_main/open-tracing/lib/tracing"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: expecting argument")
	}
	helloTo := os.Args[1]

	// Note that we are passing a string hello-world to the init method.
	// It is used to mark all spans emitted by the tracer as
	// originating from a hello-world service.
	tracer, closer := tracing.Init("hello-world")
	defer closer.Close()

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	println(helloStr)
	span.LogKV("event", "println", "mykey", "myvalue")

	span.Finish()
}
