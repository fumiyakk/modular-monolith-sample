package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"
	samplev1 "github.com/fumiyakk/modular-monolith-sample/gen/sample/v1"
	"github.com/fumiyakk/modular-monolith-sample/gen/sample/v1/samplev1connect"
)

type CreateUserResult struct {
	UserID     string
	ContractID string
}

type Client struct {
	client samplev1connect.SampleServiceClient
}

func NewClient(baseURL string, opts ...connect.ClientOption) *Client {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	defaultOpts := []connect.ClientOption{
		connect.WithGRPC(),
	}
	opts = append(defaultOpts, opts...)

	return &Client{
		client: samplev1connect.NewSampleServiceClient(
			httpClient,
			baseURL,
			opts...,
		),
	}
}

func (c *Client) CreateUser(ctx context.Context, name string) (*CreateUserResult, error) {
	req := connect.NewRequest(&samplev1.CreateUserRequest{
		Name: name,
	})

	resp, err := c.client.CreateUser(ctx, req)
	if err != nil {
		if connectErr, ok := err.(*connect.Error); ok {
			return nil, fmt.Errorf("create user failed with code %v: %v", connectErr.Code(), connectErr.Message())
		}
		return nil, fmt.Errorf("create user failed: %v", err)
	}

	return &CreateUserResult{
		UserID:     resp.Msg.UserId,
		ContractID: resp.Msg.ContractId,
	}, nil
}

func (c *Client) GetUser(ctx context.Context, id string) (*samplev1.User, error) {
	req := connect.NewRequest(&samplev1.GetUserRequest{
		Id: id,
	})

	resp, err := c.client.GetUser(ctx, req)
	if err != nil {
		if connectErr, ok := err.(*connect.Error); ok {
			return nil, fmt.Errorf("get user failed with code %v: %v", connectErr.Code(), connectErr.Message())
		}
		return nil, fmt.Errorf("get user failed: %v", err)
	}

	return resp.Msg.User, nil
}

func (c *Client) GetContract(ctx context.Context, id string) (*samplev1.Contract, error) {
	req := connect.NewRequest(&samplev1.GetContractRequest{
		Id: id,
	})

	resp, err := c.client.GetContract(ctx, req)
	if err != nil {
		if connectErr, ok := err.(*connect.Error); ok {
			return nil, fmt.Errorf("get contract failed with code %v: %v", connectErr.Code(), connectErr.Message())
		}
		return nil, fmt.Errorf("get contract failed: %v", err)
	}

	return resp.Msg.Contract, nil
}
