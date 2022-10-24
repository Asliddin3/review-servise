package grpcClient

import (
	"fmt"

	"github.com/Asliddin3/review-servise/config"
	customerPB "github.com/Asliddin3/review-servise/genproto/customer"
	"google.golang.org/grpc"
)


///GrpcClientI ...
type ServiceManager struct {
	conf            config.Config
	customerService customerPB.CustomerServiceClient
}

func New(cnfg config.Config) (*ServiceManager, error) {
	connCustomer, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cnfg.CustomerServiceHost, cnfg.CustomerServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("error while dial product service: host: %s and port: %d",
			cnfg.CustomerServiceHost, cnfg.CustomerServicePort)
	}

	serviceManager := &ServiceManager{
		conf:            cnfg,
		customerService: customerPB.NewCustomerServiceClient(connCustomer),
	}

	return serviceManager, nil
}

func (s *ServiceManager) CustomerService() customerPB.CustomerServiceClient {
	return s.customerService
}
