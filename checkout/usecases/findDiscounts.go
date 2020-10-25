package usecases

import (
	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
)

type DiscountFinder interface {
	findComboDiscounts(watchIDs []models.WatchID) ([]models.Discount, error)
}

type discountFinder struct {
	repository infrastructure.ComboDiscountRepository
}

func NewDiscountFinder(repository infrastructure.ComboDiscountRepository) DiscountFinder {
	return &discountFinder{repository}
}

func (df *discountFinder) findComboDiscounts(watchIDs []models.WatchID) ([]models.Discount, error) {
	return df.repository.GetDiscounts(watchIDs)
}
