package service

import (
	"context"

	pb "github.com/Asliddin3/review-servise/genproto/customer"
	l "github.com/Asliddin3/review-servise/pkg/logger"
	grpcclient "github.com/Asliddin3/review-servise/service/grpc_client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerService struct {
	client *grpcclient.ServiceManager
	logger l.Logger
}

func (r *CustomerService) GetCustomerInfo(ctx context.Context, req *pb.CustomerId) (*pb.CustomerResponse, error) {
	customerInfo, err := r.client.CustomerService().GetCustomerInfo(context.Background(), &pb.CustomerId{Id: req.Id})
	if err != nil {
		r.logger.Error("error getting customer info in review", l.Any("error getting customer info ", err))
		return &pb.CustomerResponse{}, status.Error(codes.Internal, "something went wrong")
	}
	return customerInfo, nil
}
