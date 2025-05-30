package server

import (
	"context"

	"github.com/GameXost/YandexGo_proj/USERS/internal/services"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	pb.UnimplementedClientServer
	Service *services.UserService
}

func (s *UserServer) GetUserProfile(ctx context.Context, _ *emptypb.Empty) (*pb.User, error) {
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok || userID == "" {
		return nil, status.Error(codes.Unauthenticated, "userID not found in context")
	}
	return s.Service.GetUserProfile(ctx, userID)
}

func (s *UserServer) UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.User, error) {
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok || userID == "" {
		return nil, status.Error(codes.Unauthenticated, "userID not found in context")
	}
	if req.Id != userID {
		return nil, status.Error(codes.PermissionDenied, "cannot update another user's profile")
	}
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
