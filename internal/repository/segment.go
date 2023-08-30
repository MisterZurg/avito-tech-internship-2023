package repository

import (
	pb "avito-tech-internship-2023/proto"
	context "context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SegmentRepository struct {
	*sqlx.DB
	pb.UnimplementedSegmentsServiceServer
}

// CreateSegment()  POST /v1/example/createSegment
func (sr *SegmentRepository) CreateSegment(ctx context.Context, in *pb.CreateSegmentRequest) (*pb.CreateSegmentResponse, error) {
	slug := in.Slug
	if slug == "" {
		return &pb.CreateSegmentResponse{},
			status.Errorf(codes.Aborted, "Slug not specified")
	}

	var existedSlug string
	_ = sr.DB.QueryRowx("SELECT slug FROM segments WHERE slug=$1", slug).Scan(&existedSlug)
	if existedSlug == slug {
		return &pb.CreateSegmentResponse{},
			status.Errorf(codes.AlreadyExists, "Slug already exist")
	}

	row := sr.DB.QueryRowx(`INSERT INTO segments (slug) VALUES ($1) RETURNING slug`,
		in.Slug,
	)

	var insertedSlug string
	if err := row.Scan(&insertedSlug); err != nil {
		return &pb.CreateSegmentResponse{},
			status.Errorf(codes.Canceled, "Cannot insert slug")
	}

	return &pb.CreateSegmentResponse{Message: fmt.Sprintf(`Slug:"%s" was sucsesfully created`, slug)}, nil
}

// DELETE /v1/example/deleteSegment/{slug}
func (sr *SegmentRepository) DeleteSegment(ctx context.Context, in *pb.DeleteSegmentRequest) (*pb.DeleteSegmentResponse, error) {
	slug := in.Slug
	if slug == "" {
		return &pb.DeleteSegmentResponse{},
			status.Errorf(codes.Aborted, "Slug not specified")
	}
	_, err := sr.DB.Exec("DELETE FROM segments WHERE slug=$1", slug)
	if err != nil {
		return &pb.DeleteSegmentResponse{},
			status.Errorf(codes.Canceled, "Cannot delete specified slug")
	}
	return &pb.DeleteSegmentResponse{Message: fmt.Sprintf(`Slug:"%s" was sucsesfully deleted`, slug)}, nil
}
