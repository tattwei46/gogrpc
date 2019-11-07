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
	"strconv"

	accountpb "davidcheah.com/gogrpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Account struct {
	ID       string `bson:"_id,omitempty"`
	UserID   string `bson:"author_id"`
	Email    string `bson:"content"`
	Password string `bson:"title"`
}

type AccountServiceServer struct {
}

func (s *AccountServiceServer) ReadAccount(ctx context.Context, req *accountpb.ReadAccountReq) (*accountpb.ReadAccountRes, error) {
	response := &accountpb.ReadAccountRes{
		Account: &accountpb.Account{
			Id:       "1",
			UserId:   "1",
			Email:    "Test",
			Password: "Test",
		},
	}

	return response, nil
}

func (s *AccountServiceServer) ListAccounts(req *accountpb.ListAccountsReq, stream accountpb.AccountService_ListAccountsServer) error {

	for i := 0; i < 100000; i++ {
		stream.Send(&accountpb.ListAccountsRes{
			Account: &accountpb.Account{
				Id:       strconv.Itoa(i),
				UserId:   strconv.Itoa(i),
				Email:    "Test",
				Password: "Test",
			},
		})
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := &AccountServiceServer{}
	accountpb.RegisterAccountServiceServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
