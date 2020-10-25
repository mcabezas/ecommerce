package infrastructure

import "github.com/mcabezas/ecommerce/checkout/models"

type FixedDiscountRepository interface {
	GetDiscounts(watchIDs []models.WatchID) ([]models.Discount, error)
}
