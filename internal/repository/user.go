package repository

import (
	pb "avito-tech-internship-2023/proto"
	context "context"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	*sqlx.DB
	pb.UnimplementedUsersServiceServer
}

func (ur *UserRepository) AddUserToSegment(context.Context, *pb.AddUserToSegmentRequest) (*pb.AddUserToSegmentResponse, error) {
	return nil, nil
}

func (ur *UserRepository) GetActiveSegments(context.Context, *pb.GetActiveSegmentsRequest) (*pb.GetActiveSegmentsResponse, error) {
	return nil, nil
}
