package service

import (
	"avito-tech-internship-2023/internal/repository"
	pb "avito-tech-internship-2023/proto"

	"google.golang.org/grpc"
)

type Service struct {
	*grpc.Server
	*repository.Repository
}

func NewService(s *grpc.Server, repos *repository.Repository) *Service {
	pb.RegisterUsersServiceServer(s, repos.UserRepository)
	pb.RegisterSegmentsServiceServer(s, repos.SegmentRepository)
	return &Service{s, repos}
}
