package server

import (
	"context"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/services"
)

type DriverServer struct {
	Service *services.DriverService
}

func (s *DriverServer) UpdateDriverProfile(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	return s.Service.UpdateDriverProfile(ctx, req)
}
