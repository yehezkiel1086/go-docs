package handler

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	conn *grpc.ClientConn
}

func NewGRPCClient(
	addr string,
) (*GRPCClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GRPCClient{conn}, nil
}

func (c *GRPCClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *GRPCClient) Close() error {
	return c.conn.Close()
}
