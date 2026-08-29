package handler

import (
	"fmt"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	conn *grpc.ClientConn
}

func NewGRPCClient(
	grpcConf *config.Inventory,
) (*GRPCClient, error) {
	addr := fmt.Sprintf("%s:%s", grpcConf.Host, grpcConf.Port)

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
