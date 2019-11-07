/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"

	userpb "davidcheah.com/gogrpc/proto"
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
