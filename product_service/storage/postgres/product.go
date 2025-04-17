package postgres

import (
	"context"
	"gitlab.com/e-market724/back-end/product_service/genproto/product_service"
	"log"
)

type ProductRepo struct{}

func NewProductRepo() ProductRepoI {
	return &ProductRepo{}
}

func (c *ProductRepo) Ping(ctx context.Context, Ping *product_service.PING) (*product_service.PING, error) {

	log.Println("message: Pong")

	return nil, nil
}
