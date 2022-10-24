package repo

import (
	pb "github.com/Asliddin3/review-servise/genproto/customer"
)

type CustomerStorageI interface {
	GetCustomerInfo(*pb.CustomerId) (*pb.CustomerResponse, error)
}
