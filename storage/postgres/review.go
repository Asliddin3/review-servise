package postgres

import (
	pb "github.com/Asliddin3/review-servise/genproto/review"
	"github.com/jmoiron/sqlx"
)

type reviewRepo struct {
	db *sqlx.DB
}

func NewReviewRepo(db *sqlx.DB) *reviewRepo {
	return &reviewRepo{db: db}
}

func (r *reviewRepo) GetReview(req *pb.PostId) (*pb.PostReview, error) {
	reviewResp := pb.PostReview{}
	err := r.db.QueryRow(
		`select AVG(review),count(*) from reviews where post_id=$1`, req.Id,
	).Scan(&reviewResp.Review, &reviewResp.Count)
	if err != nil {
		return &pb.PostReview{}, err
	}
	return &reviewResp, nil
}

func (r *reviewRepo) CreateReview(req *pb.Review) (*pb.Review, error) {
	postResp := pb.Review{}
	err := r.db.QueryRow(`
	insert into review(review,description,post_id,customer_id)
	values($1,$2,$3,$4) retuning review,description,post_id,customer_id
	`).Scan(&postResp.Review, &postResp.Description, &postResp.PostId, &postResp.CustomerId)
	if err != nil {
		return &pb.Review{}, err
	}
	return &postResp, nil
}
