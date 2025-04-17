package service

import (
	"context"
	"gitlab.com/e-market724/back-end/product_service/genproto/product_service"
	"gitlab.com/e-market724/back-end/product_service/storage"
)

type ProductService struct {
	Storage storage.StorageI
	product_service.UnimplementedPingServer
}

func NewProductService(storage storage.StorageI) *ProductService {
	return &ProductService{Storage: storage}
}

func (c ProductService) SendPing(ctx context.Context, req *product_service.PING) (*product_service.PONG, error) {
	c.Storage.GetProductRepo().Ping(ctx, req)

	return &product_service.PONG{
		Msg: "pong",
	}, nil
}
