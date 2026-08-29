package main

import (
	"context"
	"errors"
	"net"

	inventory "github.com/yehezkiel1086/go-docs/go-grpc-order-inventory/services/common/genproto/inventory/protobuf"
	"google.golang.org/grpc"
)

type InventoryServer struct {
	inventory.UnimplementedInventoryServiceServer
}

var stock = map[int64]int32 {
	1: 10,
}

func (s *InventoryServer) CheckStock(ctx context.Context, req *inventory.CheckStockReq) (*inventory.CheckStockRes, error) {
	return &inventory.CheckStockRes{
		ProductId: req.ProductId,
		Quantity: stock[1],
	}, nil
}

func (s *InventoryServer) ReduceStock(ctx context.Context, req *inventory.ReduceStockReq) (*inventory.ReduceStockRes, error) {
	if stock[1] == 0 {
		return &inventory.ReduceStockRes{
			Success: false,
		}, errors.New("RUN OUT OF STOCK")
	}

	stock[1] -= req.Quantity

	if stock[1] <= 0 {
		stock[1] = 0
	}

	return &inventory.ReduceStockRes{
		Success: true,
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()

	inventory.RegisterInventoryServiceServer(
		grpcServer,
		&InventoryServer{},
	)

	grpcServer.Serve(lis)
}
