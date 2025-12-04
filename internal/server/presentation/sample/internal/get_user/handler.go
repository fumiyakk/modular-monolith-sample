package get_user

import (
	"context"

	"connectrpc.com/connect"
	samplev1 "github.com/fumiyakk/modular-monolith-sample/gen/sample/v1"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario"
	"github.com/google/uuid"
)

type Handler interface {
	GetUser(context.Context, *connect.Request[samplev1.GetUserRequest]) (*connect.Response[samplev1.GetUserResponse], error)
}

type handler struct {
	s scenario.GetUserScenario
}

func New(s scenario.GetUserScenario) Handler {
	return &handler{s}
}

func (h *handler) GetUser(ctx context.Context, req *connect.Request[samplev1.GetUserRequest]) (*connect.Response[samplev1.GetUserResponse], error) {
	stringID := req.Msg.GetId()
	id, err := uuid.Parse(stringID)
	if err != nil {
		return nil, err
	}

	user, err := h.s.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&samplev1.GetUserResponse{
		User: &samplev1.User{
			Id:   user.ID.String(),
			Name: user.Name,
		},
	}), nil
}
