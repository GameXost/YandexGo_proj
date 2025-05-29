package server

import (
    "context"
    "fmt"

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
    fmt.Printf("driverID in gRPC method: %v\n", ctx.Value(DriverIDKey))
    //fmt.Println("GetDriverProfile called!")
    driverID, ok := ctx.Value(DriverIDKey).(string)
    if !ok || driverID == "" {
        return nil, status.Error(codes.Unauthenticated, "driverID not found in context")
    }
    return s.Service.GetDriverProfile(ctx, driverID)
}

func (s *DriverServer) UpdateDriverProfile(ctx context.Context, req *pb.UpdateDriverProfileRequest) (*pb.Driver, error) {
    driverID, ok := ctx.Value(DriverIDKey).(string)
    if !ok || driverID == "" {
        return nil, status.Error(codes.Unauthenticated, "driverID not found in context")
    }
    if req.Driver == nil || req.Driver.Id != driverID {
        return nil, status.Error(codes.PermissionDenied, "cannot update another driver's profile")
    }
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
    driverID, ok := ctx.Value(DriverIDKey).(string)
    if !ok || driverID == "" {
        return nil, status.Error(codes.Unauthenticated, "driverID not found in context")
    }
    return s.Service.CompleteRide(ctx, req.Id, driverID)
}

func (s *DriverServer) CancelRide(ctx context.Context, req *pb.RideIdRequest) (*pb.StatusResponse, error) {
    driverID, ok := ctx.Value(DriverIDKey).(string)
    if !ok || driverID == "" {
        return nil, status.Error(codes.Unauthenticated, "driverID not found in context")
    }
    return s.Service.CancelRide(ctx, req.Id, driverID)
}

func (s *DriverServer) GetCurrentRide(ctx context.Context, req *pb.DriverIdRequest) (*pb.Ride, error) {
    driverID, ok := ctx.Value(DriverIDKey).(string)
    if !ok || driverID == "" {
        return nil, status.Error(codes.Unauthenticated, "driverID not found in context")
    }
    if req.Id != driverID {
        return nil, status.Error(codes.PermissionDenied, "cannot get another driver's current ride")
    }
    return s.Service.GetCurrentRide(ctx, driverID)
}

func (s *DriverServer) UpdateLocation(stream pb.Drivers_UpdateLocationServer) error {
    ctx := stream.Context()
    driverID, ok := ctx.Value(DriverIDKey).(string)
    if !ok || driverID == "" {
        return status.Error(codes.Unauthenticated, "driverID not found in context")
    }
    updates := make(chan *pb.LocationUpdateRequest)
    go func() {
        defer close(updates)
        for {
            req, err := stream.Recv()
            if err != nil {
                break
            }
            // Ensure only the authenticated driver can update their own location
            if req.DriverId == driverID {
                updates <- req
            }
        }
    }()
    resp, err := s.Service.UpdateLocation(ctx, updates)
    if err != nil {
        return err
    }
    return stream.SendAndClose(resp)
}

func (s *DriverServer) GetNearbyRequests(ctx context.Context, req *pb.Location) (*pb.RideRequestsResponse, error) {
    driverID, ok := ctx.Value(DriverIDKey).(string)
    if !ok || driverID == "" {
        return nil, status.Error(codes.Unauthenticated, "driverID not found in context")
    }
    // Optionally, you could log or use driverID for auditing
    return s.Service.GetNearbyRequests(ctx, req)
}

func (s *DriverServer) GetPassengerInfo(ctx context.Context, req *pb.UserIdRequest) (*pb.User, error) {
    return s.Service.GetPassengerInfo(ctx, req.Id)
}

func (s *DriverServer) GetRideHistory(ctx context.Context, req *pb.DriverIdRequest) (*pb.RideHistoryResponse, error) {
    driverID, ok := ctx.Value(DriverIDKey).(string)
    if !ok || driverID == "" {
        return nil, status.Error(codes.Unauthenticated, "driverID not found in context")
    }
    if req.Id != driverID {
        return nil, status.Error(codes.PermissionDenied, "cannot get another driver's ride history")
    }
    return s.Service.GetRideHistory(ctx, driverID)
}
