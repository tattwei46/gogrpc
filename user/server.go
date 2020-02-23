//go:generate protoc ./user.proto --go_out=plugins=grpc:./pb
package user

import (
	"context"
	"fmt"
	"net"
	"strconv"

	userpb "go-grpc/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	userpb.RegisterUserServiceServer(serv, &grpcServer{s})
	fmt.Printf("Listening and serving GRPC on %s\n", lis.Addr())
	return serv.Serve(lis)
}

func (s *grpcServer) CreateUser(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserRes, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("[GRPC] Server Response: Received from ", peer.Addr)
	user := req.GetUser()
	return &userpb.CreateUserRes{User: user}, nil
}

func (s *grpcServer) GetUser(ctx context.Context, req *userpb.GetUserReq) (*userpb.GetUserRes, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("[GRPC] Server Response: Received from ", peer.Addr)

	// Hard coded response
	response := &userpb.GetUserRes{
		User: &userpb.User{
			Id:       "1",
			UserId:   "1",
			Email:    "Test",
			Password: "Test",
		},
	}

	return response, nil
}

func (s *grpcServer) ListUsers(req *userpb.ListUsersReq, stream userpb.UserService_ListUsersServer) error {
	for i := 0; i < 100000; i++ {
		stream.Send(&userpb.ListUsersRes{
			// Hard coded response
			User: &userpb.User{
				Id:       strconv.Itoa(i),
				UserId:   strconv.Itoa(i),
				Email:    "Test",
				Password: "Test",
			},
		})
	}

	return nil
}
