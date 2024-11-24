package create_user

import (
	"context"

	"connectrpc.com/connect"
	samplev1 "github.com/fumiyakk/modular-monolith-sample/gen/sample/v1"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario"
)

type Handler interface {
	CreateUser(context.Context, *connect.Request[samplev1.CreateUserRequest]) (*connect.Response[samplev1.CreateUserResponse], error)
}

type handler struct {
	s scenario.CreateUserScenario
}

func New(s scenario.CreateUserScenario) Handler {
	return &handler{s}
}

func (h *handler) CreateUser(ctx context.Context, req *connect.Request[samplev1.CreateUserRequest]) (*connect.Response[samplev1.CreateUserResponse], error) {
	userID, contractID, err := h.s.CreateUser(ctx, req.Msg.GetName())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&samplev1.CreateUserResponse{
		UserId:     userID.String(),
		ContractId: contractID.String(),
	}), nil
}
