package repository

import (
	pb "avito-tech-internship-2023/proto"
	context "context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository struct {
	*sqlx.DB
	pb.UnimplementedUsersServiceServer
}

func (ur *UserRepository) AddUserToSegment(ctx context.Context, in *pb.AddUserToSegmentRequest) (*pb.AddUserToSegmentResponse, error) {
	user_id := in.UserId
	slugs_add := in.SlugsAdd
	slugs_del := in.SlugsDel

	fmt.Println(user_id, slugs_add, slugs_del)

	if user_id == "" {
		return &pb.AddUserToSegmentResponse{},
			status.Errorf(codes.Aborted, "User ID not specified")
	}

	// Adding user to segments
	// Check if this segment already exist
	for _, slug := range slugs_add {

		_, err := ur.DB.Exec(
			`UPDATE users SET segments = array_append(segments, $1)
					WHERE user_id=$2 AND
    					(
    						SELECT ARRAY[$1] && (SELECT segments::text[] FROM users WHERE user_id=$2)
	    				) = false`,
			slug,
			user_id,
		)
		// This slug currently not in array
		if err != nil {
			fmt.Println(slug)
			fmt.Println(slugs_add)
			fmt.Println(err)
			return &pb.AddUserToSegmentResponse{},
				status.Errorf(codes.Canceled, "Cannot insert slug")
		}
		// TODO : Think about updating TTL
		//_ = ur.DB.QueryRowx(`UPDATE users SET (user_id, segments) = ($1, $2) WHERE user_id=$1`,
		//	user_id,
		//	pq.Array(slugs_add),
		//)
	}

	// Removing slugs from segments
	for _, slug := range slugs_del {
		_ = ur.DB.QueryRowx(`UPDATE users SET segments = ARRAY_REMOVE(segments, $1) WHERE user_id=$2;`,
			slug,
			user_id,
		)
	}

	return &pb.AddUserToSegmentResponse{Message: fmt.Sprintf(`User with ID:"%s" was sucsesfully updated`, user_id)},
		nil
}

func (ur *UserRepository) GetActiveSegments(ctx context.Context, in *pb.GetActiveSegmentsRequest) (*pb.GetActiveSegmentsResponse, error) {
	if in.UserId == "" {
		return &pb.GetActiveSegmentsResponse{}, status.Errorf(codes.InvalidArgument, "something unexpected happened")
	}
	type pgUserSlugs struct {
		userID      string   `json:"user_id"`
		activeSlugs []string `json:"segments"`
	}

	var pgUS pgUserSlugs

	row := ur.DB.QueryRowx("SELECT user_id, segments FROM users WHERE user_id=$1;", in.UserId)

	err := row.Scan(&pgUS.userID, pq.Array(&pgUS.activeSlugs))

	switch err {
	case sql.ErrNoRows:
		return &pb.GetActiveSegmentsResponse{}, status.Errorf(codes.NotFound, "User with specified ID notfound")
	case nil:
		return &pb.GetActiveSegmentsResponse{Slugs: pgUS.activeSlugs}, nil
	default:
		return &pb.GetActiveSegmentsResponse{}, status.Errorf(codes.Unknown, "something unexpected happened")
	}
}
