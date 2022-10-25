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

func (r *reviewRepo) GetReviewById(req *pb.ReviewId) (*pb.ReviewResp, error) {
	reviewResp := pb.ReviewResp{}
	err := r.db.QueryRow(`
	select id,post_id,customer_id,description,review,created_at,updated_at from review where id=$1 and deleted_at is null
	`, req.Id).Scan(&reviewResp.Id, &reviewResp.PostId, &reviewResp.CustomerId, &reviewResp.Description, &reviewResp.Review,
		&reviewResp.CreatedAt, &reviewResp.UpdatedAt)
	if err != nil {
		return &pb.ReviewResp{}, err
	}
	return &reviewResp, nil
}

func (s *reviewRepo) GetCustomerReviews(req *pb.CustomerId) (*pb.CustomerReviewList, error) {
	rows, err := s.db.Query(`
	select id,description,review,post_id,created_at,updated_at from review where customer_id=$1 and deleted_at is null
	`, req.Id)
	if err != nil {
		return &pb.CustomerReviewList{}, err
	}
	reviewList := pb.CustomerReviewList{}
	for rows.Next() {
		reviewResp := pb.CustomerReivewResp{}
		err := rows.Scan(&reviewResp.Id, &reviewResp.Description, &reviewResp.Review,
			&reviewResp.PostId, &reviewResp.CreatedAt, &reviewResp.UpdatedAt)
		if err != nil {
			return &pb.CustomerReviewList{}, err
		}
		reviewList.ReviewList = append(reviewList.ReviewList, &reviewResp)
	}
	return &reviewList, nil
}

func (r *reviewRepo) GetPostReviews(req *pb.PostId) (*pb.ReviewsList, error) {
	fmt.Println(req.Id)
	row, err := r.db.Query(`
	select id,customer_id,review,description from review where post_id=$1 and deleted_at is null
	`, req.Id)
	fmt.Println(err)
	if err != nil {
		return &pb.ReviewsList{}, err
	}
	reviewList := pb.ReviewsList{}
	for row.Next() {
		reviewResp := pb.ReviewRespList{}
		err := row.Scan(&reviewResp.Id,
			&reviewResp.CustomerId,
			&reviewResp.Review,
			&reviewResp.Description)
		if err != nil {
			return &pb.ReviewsList{}, err
		}
		reviewList.Reviews = append(reviewList.Reviews, &reviewResp)
	}
	fmt.Println(reviewList)
	return &reviewList, nil
}

func (r *reviewRepo) DeleteReview(req *pb.ReviewId) (*pb.Empty, error) {
	fmt.Println(req)
	_, err := r.db.Exec(`
	update review set deleted_at=current_timestamp where id=$1
	`, req.Id)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}

func (r *reviewRepo) GetPostOverall(req *pb.PostId) (*pb.PostReview, error) {
	reviewResp := pb.PostReview{}
	fmt.Println(req.Id)
	count := 0
	err := r.db.QueryRow(`
	select count(*) from review where post_id=$1 and deleted_at is null
	`, req.Id).Scan(&count)

	if err != nil {
		return &pb.PostReview{}, err
	}

	if count != 0 {
		err = r.db.QueryRow(
			`select ROUND(AVG(review),2),count(*) from review where post_id=$1 and deleted_at is null`, req.Id,
		).Scan(&reviewResp.OveralReview, &reviewResp.Count)
		fmt.Println(err, reviewResp)
		if err != nil {
			return &pb.PostReview{}, err
		}
	}

	return &reviewResp, nil
}

func (r *reviewRepo) CreateReview(req *pb.ReviewRequest) (*pb.Review, error) {
	var id int
	existReview := r.db.QueryRow(`
	select Count(*) from review where customer_id=$1 and post_id=$2
	`, req.CustomerId, req.PostId).Scan(&id)
	if existReview != nil {
		return &pb.Review{}, existReview
	}
	fmt.Println(id)
	if id != 0 {
		return &pb.Review{}, nil
	}
	postResp := pb.Review{}
	err := r.db.QueryRow(`
	insert into review(review,description,post_id,customer_id)
	values($1,$2,$3,$4) returning id,review,description,post_id,customer_id
	`, req.Review, req.Description, req.PostId, req.CustomerId).Scan(&postResp.Id, &postResp.Review, &postResp.Description, &postResp.PostId, &postResp.CustomerId)
	if err != nil {
		return &pb.Review{}, err
	}
	return &postResp, nil
}
