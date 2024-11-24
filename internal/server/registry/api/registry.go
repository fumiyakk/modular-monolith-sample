package api

import (
	"github.com/fumiyakk/modular-monolith-sample/gen/sample/v1/samplev1connect"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/unit_of_work"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/presentation/sample"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario"
)

type apiRegistry struct {
	h sample.HandlerInjector
}

func Init() interface {
	GrpcHandler() samplev1connect.SampleServiceHandler
} {
	m := InitModule()

	uow := unit_of_work.NewUnitOfWork()

	s := scenario.New(uow, m)

	h := sample.New(s)

	return &apiRegistry{h}
}

func (r *apiRegistry) GrpcHandler() samplev1connect.SampleServiceHandler {
	return r.h
}
