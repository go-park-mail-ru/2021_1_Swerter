package user_service_client

import (
	"context"
	"fmt"
	desc "gitlab.com/Burunduck/user-service/pkg/user_service_api"
	"google.golang.org/grpc"
	"log"
	"my-motivation/internal/app/models"
	"strconv"
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
	_, err := c.api.Register(ctx, &desc.RegisterRequest{
		Login:    u.Login,
		Password: u.Password,
		FirstName: u.FirstName,
		LastName: u.LastName,
	})
	return err
}

func (c *Client) Login(ctx context.Context, u* models.User) (string, error)  {
	lr, err := c.api.Login(ctx, &desc.LoginRequest{
		Login:    u.Login,
		Password: u.Password,
	})
	sess := lr.Session
	return sess, err
}

func (c *Client) GetUserBySession(ctx context.Context, session string) (*models.User, error)  {
	sr, err := c.api.GetUserBySession(ctx, &desc.UserBySessionRequest{
		Session: session,
	})
	id,_ := strconv.Atoi(sr.User.Id)
	u := models.User{
		ID: id,
		FirstName: sr.User.FirstName,
		LastName: sr.User.LastName,
	}
	return &u, err
}

func (c *Client) Logout(ctx context.Context, session string) error {
	_, err := c.api.Logout(ctx, &desc.LogoutRequest{Session: session})
	if err != nil {
		return err
	}
	return nil
}
