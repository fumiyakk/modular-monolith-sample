package middleware

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

func NewErrorLogInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)

			if err != nil {
				slog.Error("error occurred during processing request", err)
			}

			return resp, err
		})
	}
}
