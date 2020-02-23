package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"time"
)

const timeFormat = "01/01/2006 15:04:05"

func ServerLogIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	today := time.Now()
	peer, _ := peer.FromContext(ctx)
	fmt.Printf("[INTERCEPT] Server call method %s from: %s at %s\n",
		info.FullMethod,
		peer.Addr.String(),
		today.Format(timeFormat))
	return handler(ctx, req)
}

func ClientLogIntercept(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	today := time.Now()
	fmt.Printf("[INTERCEPT] Client call method %s at %s\n", method, today.Format(timeFormat))
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
