package repo

import (
	pb "github.com/Asliddin3/review-servise/genproto/review"
)

type ReviewStorageI interface {
	// CheckField(*pb.CheckFieldRequest) (*pb.CheckFieldResponse,error)
	CreateReview(*pb.Review) (*pb.Review, error)
	GetReview(*pb.PostId) (*pb.PostReview, error)
}
