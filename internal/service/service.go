package service

import (
	pb "github.com/Saser/strecku/api/v1"
	"github.com/Saser/strecku/resources/users"
)

type Service struct {
	pb.UnimplementedStreckUServer

	userRepo *users.Repository
}

func New(userRepo *users.Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}
