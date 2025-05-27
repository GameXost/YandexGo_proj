package server

import (
	"USERS/internal/services"
	"context"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
)

type UserServer struct {
	pb.UnimplementedClientServer
	Service *services.UserService
}

func (s *UserServer) GetUserProfile(ctx context.Context, req *pb.AuthToken) (*pb.User, error) {
	// req.UserId или req.Token — зависит от структуры AuthToken
	return s.Service.GetUserProfile(ctx, req.UserId)
}
