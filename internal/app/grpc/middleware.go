package grpcapp

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func timeoutInterceptor(duration time.Duration) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, duration)
		defer cancel()

		done := make(chan struct{})
		var resp interface{}
		var err error

		go func() {
			resp, err = handler(ctx, req)
			close(done)
		}()

		select {
		case <-ctx.Done():
			return nil, status.Errorf(codes.DeadlineExceeded, "server timeout: %v", ctx.Err())
		case <-done:
			return resp, err
		}
	}
}
