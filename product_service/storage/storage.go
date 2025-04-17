package storage

import "gitlab.com/e-market724/back-end/product_service/storage/postgres"

type StorageI interface {
	GetProductRepo() postgres.ProductRepoI
}

type storage struct {
	productRepo postgres.ProductRepoI
}

func NewStorage() StorageI {
	return &storage{
		productRepo: postgres.NewProductRepo(),
	}
}

func (s *storage) GetProductRepo() postgres.ProductRepoI {
	return s.productRepo
}
