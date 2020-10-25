package usecases

import (
	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
)

type DiscountFinder interface {
	findFixedDiscounts(watchIDs []models.WatchID) ([]models.Discount, error)
}

type discountFinder struct {
	repository infrastructure.FixedDiscountRepository
}

func NewDiscountFinder(repository infrastructure.FixedDiscountRepository) DiscountFinder {
	return &discountFinder{repository}
}

func (df *discountFinder) findFixedDiscounts(watchIDs []models.WatchID) ([]models.Discount, error) {
	return df.repository.GetDiscounts(watchIDs)
}
