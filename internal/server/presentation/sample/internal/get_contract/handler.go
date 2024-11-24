package get_contract

import (
	"context"

	"connectrpc.com/connect"
	samplev1 "github.com/fumiyakk/modular-monolith-sample/gen/sample/v1"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/entity"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario"
	"github.com/google/uuid"
)

type Handler interface {
	GetContract(context.Context, *connect.Request[samplev1.GetContractRequest]) (*connect.Response[samplev1.GetContractResponse], error)
}

type handler struct {
	s scenario.GetContractScenario
}

func New(s scenario.GetContractScenario) Handler {
	return &handler{s}
}

func (h *handler) GetContract(ctx context.Context, req *connect.Request[samplev1.GetContractRequest]) (*connect.Response[samplev1.GetContractResponse], error) {
	stringID := req.Msg.GetId()
	id, err := uuid.Parse(stringID)
	if err != nil {
		return nil, err
	}

	contract, err := h.s.GetContract(ctx, id)
	if err != nil {
		return nil, err
	}

	var status samplev1.ContractStatus
	switch contract.Status {
	case entity.ContractStatusActive:
		status = samplev1.ContractStatus_CONTRACT_STATUS_ACTIVE
	case entity.ContractStatusInactive:
		status = samplev1.ContractStatus_CONTRACT_STATUS_INACTIVE
	default:
		status = samplev1.ContractStatus_CONTRACT_STATUS_UNSPECIFIED
	}

	return connect.NewResponse(&samplev1.GetContractResponse{
		Contract: &samplev1.Contract{
			Id:     contract.ID.String(),
			UserId: contract.UserID.String(),
			Status: status,
		},
	}), nil
}
