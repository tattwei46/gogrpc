package user

import (
	"context"
	proto "go-grpc/user/pb"
	"io"
	"log"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client proto.UserServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := proto.NewUserServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) ListUsers(ctx context.Context) {
	req := &proto.ListUsersReq{}
	stream, err := c.client.ListUsers(ctx, req)
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
		log.Printf("[GRPC] Server Response: %s", res.GetUser())
	}
}

func (c *Client) GetUser(ctx context.Context, userID string) (User, error) {
	res, err := c.client.GetUser(ctx, &proto.GetUserReq{Id: userID})
	if err != nil {
		return User{}, err
	}
	log.Printf("[GRPC] Server Response: %s", res.GetUser())

	return ToUser(res.GetUser()), nil
}

func (c *Client) CreateUser(ctx context.Context, user User) error {
	res, err := c.client.CreateUser(ctx, &proto.CreateUserReq{User: user.ToProto()})
	if err != nil {
		return err
	}
	log.Printf("[GRPC] Server Response: %s", res.GetUser())
	return nil
}
