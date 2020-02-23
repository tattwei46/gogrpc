package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"time"
)

func ServerLogIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	today := time.Now()
	peer, _ := peer.FromContext(ctx)
	fmt.Printf("Server call from: %s at %s\n", peer.Addr.String(), today.Format("01/01/2006 15:04:05"))
	fmt.Println("Full method: ", info.FullMethod)
	return handler(ctx, req)
}

func ClientLogIntercept(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	today := time.Now()
	fmt.Println("client call ", method, " service date: ", today.Format("01/01/2006 15:04:05"))
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func ServerRequestInfoIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("failed")
	} else {
		fmt.Println("username: ", meta["username"][0])
	}

	return handler(ctx, req)
}
