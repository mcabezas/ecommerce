package usecases

import (
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/mcabezas/ecommerce/internal/money"
)

type CreateCheckout interface {
	create(watchIDs []models.WatchID) (money.Money, error)
}

func NewCreateCheckout() CreateCheckout {
	return &createCheckout{
	}
}

type createCheckout struct {
}

func (c *createCheckout) create(watchIDs []models.WatchID) (money.Money, error) {
	return money.Money{}, nil
}
