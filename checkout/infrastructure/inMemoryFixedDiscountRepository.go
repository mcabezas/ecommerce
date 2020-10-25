package infrastructure

import (
	"sync"

	"github.com/mcabezas/ecommerce/checkout/models"
)

type inMemoryFixedDiscountRepository struct {
	fixedDiscounts *sync.Map
}

func NewInMemoryFixedDiscountRepository(discounts []models.FixedDiscount) FixedDiscountRepository {
	fixedDiscounts := &sync.Map{}
	for _, discount := range discounts {
		fixedDiscounts.Store(discount.WatchID(), discount)
	}
	return &inMemoryFixedDiscountRepository{
		fixedDiscounts: fixedDiscounts,
	}
}

func (m *inMemoryFixedDiscountRepository) GetDiscounts(watchIDs []models.WatchID) ([]models.Discount, error) {
	var result []models.Discount
	for _, watchID := range watchIDs {
		if discount, ok := m.fixedDiscounts.Load(watchID); ok {
			fixedDiscount := discount.(models.FixedDiscount)
			result = append(result, &fixedDiscount)
		}
	}
	return result, nil
}
