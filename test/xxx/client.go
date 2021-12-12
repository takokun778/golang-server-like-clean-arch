package xxx

import (
	"os"
	"xxx/proto/xxx"

	"google.golang.org/grpc"
)

type Client struct {
	xxx.XxxServiceClient
}

var client = newClient()

func ClientConnect() *Client {
	return client
}

func newClient() *Client {
	port := os.Getenv("PORT")

	if port == "" {
		port = "19003"
	}

	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	xxxClient := xxx.NewXxxServiceClient(conn)

	return &Client{xxxClient}
}
