//go:generate protoc ./user.proto --go_out=plugins=grpc:./../pb
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"

	userpb "go-grpc/pb"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

const (
	port = ":50051"
)

type User struct {
	ID       string `bson:"_id,omitempty"`
	UserID   string `bson:"author_id"`
	Email    string `bson:"content"`
	Password string `bson:"title"`
}

type UserServiceServer struct {
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserRes, error) {
	user := req.GetUser()
	return &userpb.CreateUserRes{User: user}, nil
}

func (s *UserServiceServer) ReadUser(ctx context.Context, req *userpb.ReadUserReq) (*userpb.ReadUserRes, error) {
	response := &userpb.ReadUserRes{
		User: &userpb.User{
			Id:       "1",
			UserId:   "1",
			Email:    "Test",
			Password: "Test",
		},
	}

	return response, nil
}

func (s *UserServiceServer) ListUsers(req *userpb.ListUsersReq, stream userpb.UserService_ListUsersServer) error {

	for i := 0; i < 100000; i++ {
		stream.Send(&userpb.ListUsersRes{
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
func handleGet(c *gin.Context) {
	c.String(http.StatusOK, "test\n")
}

func main() {

	//rest
	go func() {
		router := gin.Default()
		api := router.Group("/api")
		{
			api.GET("/user/test", handleGet)
		}

		router.Run()
	}()

	//grpc
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := &UserServiceServer{}
	userpb.RegisterUserServiceServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
