package repo

import (
	pb "github.com/Asliddin3/review-servise/genproto/review"
)

type ReviewStorageI interface {
	// CheckField(*pb.CheckFieldRequest) (*pb.CheckFieldResponse,error)
	CreateReview(*pb.ReviewRequest) (*pb.Review, error)
	// GetPostReview(*pb.PostId) (*pb.PostReview, error)
	DeleteReview(*pb.PostId) (*pb.Empty, error)
	GetPostOverall(*pb.PostId) (*pb.PostReview, error)
	GetPostReviews(*pb.PostId) (*pb.ReviewsList, error)
	GetReviewById(*pb.ReviewId) (*pb.Review, error)
}
