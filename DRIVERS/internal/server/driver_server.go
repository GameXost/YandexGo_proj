package server

import (
	"context"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DriverServer struct {
	pb.UnimplementedDriversServer
	Service *services.DriverService
}

func (s *DriverServer) GetDriverProfile(ctx context.Context, req *pb.AuthToken) (*pb.Driver, error) {
	return s.Service.GetDriverProfile(ctx, req.Token)
}

func (s *DriverServer) UpdateDriverProfile(ctx context.Context, req *pb.UpdateDriverProfileRequest) (*pb.Driver, error) {
	return s.Service.UpdateDriverProfile(ctx, req.Driver)
}

func (s *DriverServer) AcceptRide(ctx context.Context, req *pb.RideIdRequest) (*pb.StatusResponse, error) {
	driverID, ok := ctx.Value(DriverIDKey).(string)
	if !ok || driverID == "" {
		return nil, status.Error(codes.Unauthenticated, "driverID not found in context")
	}
	return s.Service.AcceptRide(ctx, req.Id, driverID)
}

func (s *DriverServer) CompleteRide(ctx context.Context, req *pb.RideIdRequest) (*pb.StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteRide not implemented")
}

func (s *DriverServer) CancelRide(ctx context.Context, req *pb.RideIdRequest) (*pb.StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelRide not implemented")
}

func (s *DriverServer) GetCurrentRide(ctx context.Context, req *pb.DriverIdRequest) (*pb.Ride, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentRide not implemented")
}

func (s *DriverServer) UpdateLocation(stream pb.Drivers_UpdateLocationServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}

func (s *DriverServer) GetNearbyRequests(ctx context.Context, req *pb.Location) (*pb.RideRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearbyRequests not implemented")
}

func (s *DriverServer) GetPassengerInfo(ctx context.Context, req *pb.UserIdRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassengerInfo not implemented")
}

func (s *DriverServer) GetRideHistory(ctx context.Context, req *pb.DriverIdRequest) (*pb.RideHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRideHistory not implemented")
}
