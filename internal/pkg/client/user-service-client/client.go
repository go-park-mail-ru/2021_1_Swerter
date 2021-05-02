package user_service_client

import (
	"context"
	"fmt"
	desc "gitlab.com/Burunduck/user-service/pkg/user_service_api"
	"google.golang.org/grpc"
	"log"
	"my-motivation/internal/app/models"
)

type Client struct {
	api desc.UserServiceClient
}

func New() (*Client, error) {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:8091",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	//defer grpcConn.Close()
	fmt.Println("service started at port: 8091")
	userAPI := desc.NewUserServiceClient(grpcConn)
	return &Client{api: userAPI}, nil
}

func (c *Client) Register(ctx context.Context, u* models.User) error  {
	fmt.Println("lalalala")
	_, err := c.api.Register(ctx, &desc.RegisterRequest{
		Login:    u.Login,
		Password: u.Password,
	})
	return err
}

func (c *Client) Login(ctx context.Context, u* models.User) (string, error)  {
	fmt.Println("lalalala")
	lr, err := c.api.Login(ctx, &desc.LoginRequest{
		Login:    u.Login,
		Password: u.Password,
	})
	sess := lr.Session
	return sess, err
}