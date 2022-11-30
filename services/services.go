package services

import (
	"fmt"

	"exam/api-gateway/config"
	pbc "exam/api-gateway/genproto/customer"
	pbp "exam/api-gateway/genproto/post"
	pbr "exam/api-gateway/genproto/reyting"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	CustomerService() pbc.CustomerServiceClient
	PostService() pbp.PostServiceClient
	RankingService() pbr.RankingServiceClient
}

type serviceManager struct {
	customerService pbc.CustomerServiceClient
	postService     pbp.PostServiceClient
	rankingService  pbr.RankingServiceClient
}

func (s *serviceManager) CustomerService() pbc.CustomerServiceClient {
	return s.customerService
}

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}
func (s *serviceManager) RankingService() pbr.RankingServiceClient {
	return s.rankingService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCustomer, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CustomerServiceHost, conf.CustomerServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	fmt.Println("2")
	fmt.Println("post service host va port: ", conf.PostServiceHost, conf.PostServicePort)
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connRanking, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.RankingServiceHost, conf.RankingServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		customerService: pbc.NewCustomerServiceClient(connCustomer),
		postService:     pbp.NewPostServiceClient(connPost),
		rankingService:  pbr.NewRankingServiceClient(connRanking),
	}

	return serviceManager, nil
}
