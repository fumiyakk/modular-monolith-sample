package connect

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/fumiyakk/modular-monolith-sample/gen/sample/v1/samplev1connect"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/connect/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer(
	sample samplev1connect.SampleServiceHandler,
) http.Handler {
	interceptors := connect.WithInterceptors(
		middleware.NewDumpLogInterceptor(),
		middleware.NewErrorLogInterceptor(),
	)

	options := []connect.HandlerOption{
		interceptors,
	}

	mux := http.NewServeMux()

	mux.Handle(samplev1connect.NewSampleServiceHandler(sample, options...))

	h := h2c.NewHandler(mux, &http2.Server{})

	return h
}
