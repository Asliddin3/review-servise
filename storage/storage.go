package storage

import (
	"github.com/Asliddin3/review-servise/storage/postgres"
	"github.com/Asliddin3/review-servise/storage/repo"
	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Review() repo.ReviewStorageI
}

type storagePg struct {
	db       *sqlx.DB
	reviewRepo repo.ReviewStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		reviewRepo: postgres.NewReviewRepo(db),
	}
}
func (s storagePg) Review() repo.ReviewStorageI {
	return s.reviewRepo
}
