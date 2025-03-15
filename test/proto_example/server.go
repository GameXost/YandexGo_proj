package main

import (
	"context"
	pb "github.com/GameXost/YandexGo_proj/tree/gRPCservice/test/proto_example/coffee_shop_proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "BlackCoffee",
		},
		&pb.Item{
			Id:   "2",
			Name: "Americano",
		},
		&pb.Item{
			Id:   "3",
			Name: "Vanilla Soy Chai Latte",
		},
	}
	for i, _ := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}
	return nil
}
func (s *server) PlaceOrder(context context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABCD123",
	}, nil
}
func (s *server) GetOrderStatus(context context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "In Progress",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
