package main

import (
	"log/slog"
	"net/http"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/connect"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/registry/api"
)

func main() {
	r := api.Init()

	s := connect.NewServer(r.GrpcHandler())

	server := &http.Server{
		Handler: s,
		Addr:    ":8080",
	}

	slog.Info("Starting server")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
