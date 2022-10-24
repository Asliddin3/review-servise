package service

import (
	"context"

	"github.com/Asliddin3/review-servise/genproto/customer"
	pb "github.com/Asliddin3/review-servise/genproto/review"
	l "github.com/Asliddin3/review-servise/pkg/logger"
	grpcClient "github.com/Asliddin3/review-servise/service/grpc_client"
	"github.com/Asliddin3/review-servise/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReviewService struct {
	storage storage.IStorage
	logger  l.Logger
	client  *grpcClient.ServiceManager
}

func NewReviewService(cleint *grpcClient.ServiceManager, db *sqlx.DB, log l.Logger) *ReviewService {
	return &ReviewService{
		storage: storage.NewStoragePg(db),
		client:  cleint,
		logger:  log,
	}
}

func (s *ReviewService) GetPostReviews(ctx context.Context, req *pb.PostId) (*pb.ReviewsList, error) {
	res, err := s.storage.Review().GetPostReviews(req)
	for i, reviews := range res.Reviews {
		customerResp, err := s.client.CustomerService().GetCustomerInfo(context.Background(), &customer.CustomerId{Id: reviews.CustomerId})
		if err != nil {
			s.logger.Error("error getting customer info", l.Any("error gettin customer info", err))
			return &pb.ReviewsList{}, status.Error(codes.Internal, "something went wrong")
		}
		reviews.FirstName = customerResp.FirstName
		reviews.LastName = customerResp.LastName
		res.Reviews[i] = reviews
	}
	if err != nil {
		s.logger.Error("error getting list reviews", l.Any("error getting reviews", err))
		return &pb.ReviewsList{}, status.Error(codes.Internal, "errir getting reviews")
	}
	return res, nil
}

func (s *ReviewService) DeleteReview(ctx context.Context, req *pb.PostId) (*pb.Empty, error) {
	res, err := s.storage.Review().DeleteReview(req)
	if err != nil {
		s.logger.Error("error deleting review", l.Any("error deleting", err))
		return &pb.Empty{}, status.Error(codes.Internal, "error deleting review")
	}
	return res, nil
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.ReviewRequest) (*pb.Review, error) {
	Review, err := s.storage.Review().CreateReview(req)
	if err != nil {
		s.logger.Error("error while creating Review", l.Any("error creating Review", err))
		return &pb.Review{}, status.Error(codes.Internal, "something went wrong")
	}
	return Review, nil
}

func (s *ReviewService) GetPostOverall(ctc context.Context, req *pb.PostId) (*pb.PostReview, error) {
	Review, err := s.storage.Review().GetPostOverall(req)
	if err != nil {
		s.logger.Error("error while getting Review", l.Any("error getting Review", err))
		return &pb.PostReview{}, status.Error(codes.Internal, "something went wrong")
	}
	return Review, nil
}
