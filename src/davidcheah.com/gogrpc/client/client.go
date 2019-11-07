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

	accountpb "davidcheah.com/gogrpc/proto"
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
	c := accountpb.NewAccountServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//listAccounts(ctx, c)
	//readAccount(ctx, c)
	createAccount(ctx, c)

}
func listAccounts(ctx context.Context, c accountpb.AccountServiceClient) {
	req := &accountpb.ListAccountsReq{}
	stream, err := c.ListAccounts(ctx, req)
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
		log.Printf("Response: %s", res.GetAccount())
	}
}

func readAccount(ctx context.Context, c accountpb.AccountServiceClient) {
	res, err := c.ReadAccount(ctx, &accountpb.ReadAccountReq{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", res.GetAccount())
}

func createAccount(ctx context.Context, c accountpb.AccountServiceClient) {
	account := &accountpb.Account{
		Id:       "999",
		UserId:   "999",
		Email:    "999",
		Password: "999",
	}
	res, err := c.CreateAccount(ctx, &accountpb.CreateAccountReq{Account: account})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", res.GetAccount())
}
