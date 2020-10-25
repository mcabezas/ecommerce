package infrastructure

import "github.com/mcabezas/ecommerce/checkout/models"

type ComboDiscountRepository interface {
	GetDiscounts(watchIDs []models.WatchID) ([]models.Discount, error)
}
