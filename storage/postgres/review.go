package postgres

import (
	"fmt"

	pb "github.com/Asliddin3/review-servise/genproto/review"
	"github.com/jmoiron/sqlx"
)

type reviewRepo struct {
	db *sqlx.DB
}

func NewReviewRepo(db *sqlx.DB) *reviewRepo {
	return &reviewRepo{db: db}
}
func (r *reviewRepo) DeleteReview(req *pb.PostId) (*pb.Empty, error) {
	fmt.Println(req)
	_, err := r.db.Exec(`
	delete from review where post_id=$1
	`, req.Id)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}
func (r *reviewRepo) GetPostReview(req *pb.PostId) (*pb.PostReview, error) {
	reviewResp := pb.PostReview{}
	fmt.Println(req.Id)
	count := 0
	err := r.db.QueryRow(`
	select count(*) from review where post_id=$1
	`, req.Id).Scan(&count)
	if err != nil {
		return &pb.PostReview{}, err
	}
	fmt.Println(count)
	if count != 0 {
		err = r.db.QueryRow(
			`select ROUND(AVG(review),2),count(*) from review where post_id=$1`, req.Id,
		).Scan(&reviewResp.Review, &reviewResp.Count)
		fmt.Println(err, reviewResp)
		if err != nil {
			return &pb.PostReview{}, err
		}
	}

	fmt.Println(reviewResp)
	return &reviewResp, nil
}

func (r *reviewRepo) CreateReview(req *pb.Review) (*pb.Review, error) {
	postResp := pb.Review{}
	err := r.db.QueryRow(`
	insert into review(review,description,post_id,customer_id)
	values($1,$2,$3,$4) returning review,description,post_id,customer_id
	`, req.Review, req.Description, req.PostId, req.CustomerId).Scan(&postResp.Review, &postResp.Description, &postResp.PostId, &postResp.CustomerId)
	if err != nil {
		return &pb.Review{}, err
	}
	return &postResp, nil
}
