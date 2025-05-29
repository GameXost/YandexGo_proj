package server

import (
	"context"

	"github.com/GameXost/YandexGo_proj/USERS/internal/services"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
)

type UserServer struct {
	pb.UnimplementedClientServer
	Service *services.UserService
}

func (s *UserServer) GetUserProfile(ctx context.Context, req *pb.AuthToken) (*pb.User, error) {
	// req.UserId или req.Token — зависит от структуры AuthToken
	return s.Service.GetUserProfile(ctx, req.Token)
}

func (s *UserServer) UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.User, error) {
	return s.Service.UpdateUserProfile(ctx, req)
}

func (s *UserServer) RequestRide(ctx context.Context, req *pb.RideRequest) (*pb.Ride, error) {
	return s.Service.RequestRide(ctx, req)
}

func (s *UserServer) CancelRide(ctx context.Context, req *pb.RideIdRequest) (*pb.StatusResponse, error) {
	return s.Service.CancelRide(ctx, req)
}

func (s *UserServer) GetRideStatus(ctx context.Context, req *pb.UserIdRequest) (*pb.Ride, error) {
	return s.Service.GetRideStatus(ctx, req)
}

func (s *UserServer) GetRideHistory(ctx context.Context, req *pb.UserIdRequest) (*pb.RideHistoryResponse, error) {
	return s.Service.GetRideHistory(ctx, req)
}

func (s *UserServer) GetDriverLocation(ctx context.Context, req *pb.DriverIdRequest) (*pb.Location, error) {
	return s.Service.GetDriverLocation(ctx, req)
}

func (s *UserServer) GetDriverInfo(ctx context.Context, req *pb.DriverIdRequest) (*pb.Driver, error) {
	return s.Service.GetDriverInfo(ctx, req)
}
