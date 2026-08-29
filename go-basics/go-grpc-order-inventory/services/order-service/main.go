package main

import (
	"context"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
	inventory "github.com/yehezkiel1086/go-docs/go-grpc-order-inventory/services/common/genproto/inventory/protobuf"
	order "github.com/yehezkiel1086/go-docs/go-grpc-order-inventory/services/common/genproto/order/protobuf"
	"google.golang.org/grpc"
)

type OrderServer struct {
	order.UnimplementedOrderServiceServer
	inventoryClient inventory.InventoryServiceClient
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (*order.CreateOrderRes, error) {
	stock, err := s.inventoryClient.CheckStock(ctx, &inventory.CheckStockReq{
		ProductId: req.ProductId,
	})
	if err != nil {
		return nil, err
	}

	if stock.Quantity < req.Quantity {
		return &order.CreateOrderRes{
			Status: "OUT OF STOCK",
		}, nil
	}

	res, err := s.inventoryClient.ReduceStock(ctx, &inventory.ReduceStockReq{
		ProductId: req.ProductId,
		Quantity: req.Quantity,
	})
	if err != nil {
		return nil, err
	}

	if !res.Success {
		return &order.CreateOrderRes{
			Status: "FAILED",
		}, nil
	}

	return &order.CreateOrderRes{
		Status: "SUCCESS",
	}, nil
}

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	inventoryClient := inventory.NewInventoryServiceClient(conn)

	lis, _ := net.Listen("tcp", ":50052")
	grpcServer := grpc.NewServer()

	orderServer := &OrderServer{
		inventoryClient: inventoryClient,
	}

	order.RegisterOrderServiceServer(
		grpcServer,
		orderServer,
	)

	go grpcServer.Serve(lis)

	r := gin.Default()

	r.POST("/orders", func(c *gin.Context) {
		productID, _ := strconv.ParseInt(c.Query("product_id"), 10, 64)
    qty, _ := strconv.Atoi(c.Query("quantity"))

    res, _ := orderServer.CreateOrder(
      context.Background(),
      &order.CreateOrderReq{
        ProductId: productID,
        Quantity:  int32(qty),
      },
    )

    c.JSON(200, gin.H{"status": res.Status})
	})

	r.Run(":8080")
}
