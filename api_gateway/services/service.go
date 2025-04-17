package services

import (
	"fmt"

	"gitlab.com/e-market724/back-end/api_gateway/genproto/product_service"
	"gitlab.com/e-market724/back-end/api_gateway/genproto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Service() ServiceManagerI {

	userService, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	productService, err := grpc.NewClient("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	serviseManager := &serviceManager{
		userService:    user_service.NewUserServiceClient(userService),
		productService: product_service.NewProductServiceClient(productService),
	}

	return serviseManager
}

type ServiceManagerI interface {
	GetUserSevice() user_service.UserServiceClient
	GetProductSevice() product_service.ProductServiceClient
}

type serviceManager struct {
	userService    user_service.UserServiceClient
	productService product_service.ProductServiceClient
}

func (s *serviceManager) GetUserSevice() user_service.UserServiceClient {

	return s.userService
}

func (s *serviceManager) GetProductSevice() product_service.ProductServiceClient {

	return s.productService
}
