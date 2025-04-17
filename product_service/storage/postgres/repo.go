package postgres

import (
	"context"
	"gitlab.com/e-market724/back-end/product_service/genproto/product_service"
)

type ProductRepoI interface {
	Ping(ctx context.Context, ping *product_service.PING) (*product_service.PING, error)
}
