package service

import (
	"context"

	pb "github.com/Asliddin3/review-servise/genproto/review"
	l "github.com/Asliddin3/review-servise/pkg/logger"
	"github.com/Asliddin3/review-servise/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReviewService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewReviewService(db *sqlx.DB, log l.Logger) *ReviewService {
	return &ReviewService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}
func (s *ReviewService) DeleteReview(ctx context.Context, req *pb.PostId) (*pb.Empty, error) {
	res, err := s.storage.Review().DeleteReview(req)
	if err != nil {
		s.logger.Error("error deleting review", l.Any("error deleting", err))
		return &pb.Empty{}, status.Error(codes.Internal, "error deleting review")
	}
	return res, nil
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.Review) (*pb.Review, error) {
	Review, err := s.storage.Review().CreateReview(req)
	if err != nil {
		s.logger.Error("error while creating Review", l.Any("error creating Review", err))
		return &pb.Review{}, status.Error(codes.Internal, "something went wrong")
	}
	return Review, nil
}

func (s *ReviewService) GetPostReview(ctc context.Context, req *pb.PostId) (*pb.PostReview, error) {
	Review, err := s.storage.Review().GetPostReview(req)
	if err != nil {
		s.logger.Error("error while getting Review", l.Any("error getting Review", err))
		return &pb.PostReview{}, status.Error(codes.Internal, "something went wrong")
	}
	return Review, nil
}
