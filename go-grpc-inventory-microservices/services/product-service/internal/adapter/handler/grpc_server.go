package handler

import (
	"net"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/config"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	s *grpc.Server
}

func NewGRPCServer(conf *config.RPC) (*GRPCServer, error) {
	s := grpc.NewServer()

	return &GRPCServer{s}, nil
}

func (s *GRPCServer) GetServer() *grpc.Server {
	return s.s
}

func (s *GRPCServer) Run(conf *config.RPC) error {
	addr := conf.Host + ":" + conf.Port

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return s.s.Serve(lis)
}
