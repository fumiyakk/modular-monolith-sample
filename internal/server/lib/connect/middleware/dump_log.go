package middleware

import (
	"context"
	"log/slog"
	"reflect"

	"connectrpc.com/connect"
)

func NewDumpLogInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			slog.Info("request dump", slog.Any("body", req.Any()))

			resp, err := next(ctx, req)

			// avoid panics; handlers may return typed nil.
			if !isNilValue(resp) {
				slog.Info("response dump", slog.Any("body", resp.Any()))
			}

			return resp, err
		})
	}
}

func isNilValue(v any) bool {
	switch rv := reflect.ValueOf(v); rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return rv.IsNil()
	case reflect.Invalid:
		return true
	}
	return false
}
