package templates

const Middlewares = `
package handlers

import (
	"{{.ImportPath -}} /svc"
	pb "{{.PBImportPath -}}"

	"github.com/opentracing/opentracing-go"
	"seller/basis/tracing"
	"seller/basis/log"
	"os"
)

const (
	SERVICE_NAME = "{{.Service.Name}}Service"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/metaverse/truss/_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)
	
	in.WrapAllLabeledExcept(log.LogServer(SERVICE_NAME), "Ping")
	tracer := opentracing.GlobalTracer()
	in.WrapAllLabeledExcept(tracing.TraceServer(tracer), "Ping")
	return in
}

func WrapService(in pb.{{.Service.Name}}Server) pb.{{.Service.Name}}Server {
	_, _, err := tracing.NewJaegerTracer(SERVICE_NAME)
	if err != nil {
		log.GetInstance(SERVICE_NAME).ELog("main", "启动失败! 链路跟踪异常:"+err.Error())
		os.Exit(1)
	}
	return in
}
`
