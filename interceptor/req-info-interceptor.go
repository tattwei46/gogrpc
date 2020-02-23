package interceptor

import (
	"context"
	"errors"
	"go-grpc/constant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var ErrRequestFail = errors.New("server request rejected")

func ServerRequestInfoIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrRequestFail
	}

	if _, ok := meta[constant.UsernameKey]; !ok {
		return nil, ErrRequestFail
	}

	if _, ok := meta[constant.PasswordKey]; !ok {
		return nil, ErrRequestFail
	}

	if constant.Authenticate(meta[constant.UsernameKey][0], meta[constant.PasswordKey][0]) {
		return handler(ctx, req)
	}
	return nil, ErrRequestFail
}
