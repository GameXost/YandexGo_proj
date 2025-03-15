package main

import (
	"context"
	"log"
	"net"
)


pb "github.com/GameXost/YandexGo_proj/tree/gRPCservice/test/proto_examplez"
    "google.golang.org/grpc"


type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(*pb.MenuRequest, pb.CoffeeShop_GetMenuServer) error {
	items:= []*pb.Item {
		&pb.Item{
			Id: "1"
		}
	}
}
func (s *server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
}
func (s *server) GetOerder(context.Context, *pb.Order) (*pb.OrderStatus, error) {
}
