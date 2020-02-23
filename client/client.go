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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"io"
	"log"

	userpb "davidcheah.com/gogrpc/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listUsers(ctx, c)
	//readUser(ctx, c)
	//createUser(ctx, c)

}
func listUsers(ctx context.Context, c userpb.UserServiceClient) {
	req := &userpb.ListUsersReq{}
	stream, err := c.ListUsers(ctx, req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		log.Printf("Response: %s", res.GetUser())
	}
}

func readUser(ctx context.Context, c userpb.UserServiceClient) {
	res, err := c.ReadUser(ctx, &userpb.ReadUserReq{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", res.GetUser())
}

func createUser(ctx context.Context, c userpb.UserServiceClient) {
	User := &userpb.User{
		Id:       "999",
		UserId:   "999",
		Email:    "999",
		Password: "999",
	}
	res, err := c.CreateUser(ctx, &userpb.CreateUserReq{User: User})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", res.GetUser())
}
