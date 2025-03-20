package main

import (
	"context"
	"github.com/GameXost/YandexGo_proj/Clients/clients_proto"
	proto_files "github.com/GameXost/YandexGo_proj/Clients/clients_proto"
	"log"
	"net"
)

type ClientServer struct {
	proto_files.UnimplementedClientServer
}

func (cs ClientServer) GetUserProfile(ctx context.Context, AT *proto_files.AuthToken) (*proto_files.User, error) {
	// тут должен быть запрос из БД данных польователя по токену
}

func (cs ClientServer) UpdateUserProfile(ctx context.Context, UPR *proto_files.UpdateProfileRequest) (*proto_files.User, error) {
	// тут должен быть запрос в БД на  изменение профиля польователя
}

func (cs ClientServer) RequestRide(ctx context.Context, RR *proto_files.RideRequest) (*proto_files.Ride, error) {
	// создание нового заказа и запись его в БД
}

func (cs ClientServer) CancelRide(ctx context.Context, RIR *proto_files.RideIdRequest) (*proto_files.StatusResponse, error) {
	// отмена заказа и удаление его из БД
}

func (cs ClientServer) GetRideStatus(ctx context.Context, UID *proto_files.UserIdRequest) (*proto_files.Ride, error) {
	// получение статуса заказа и его данные из БД
}

func (cs ClientServer) GetRideHistory(ctx context.Context, UIR *proto_files.UserIdRequest) (*proto_files.RideHistoryResponse, error) {
	// история последних поездок пользоваетеля из кэша
}

func (cs ClientServer) GetDriverLocation(ctx context.Context, DIR *proto_files.DriverIdRequest) (*proto_files.Location, error) {
	// получение координат водителя из БД - редис
}
func (cs ClientServer) GetDriverInfo(ctx context.Context, DIR *proto_files.DriverIdRequest) (*proto_files.Driver, error) {
	// получение информации о водителе из БД
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

}
