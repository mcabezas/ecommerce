package infrastructure

import (
	"sync"

	"github.com/mcabezas/ecommerce/checkout/models"
)

type inMemoryComboDiscountRepository struct {
	comboDiscounts *sync.Map
}

func NewInMemoryComboDiscountRepository(discounts []models.ComboDiscount) ComboDiscountRepository {
	comboDiscounts := &sync.Map{}
	for _, discount := range discounts {
		comboDiscounts.Store(discount.WatchID(), discount)
	}
	return &inMemoryComboDiscountRepository{
		comboDiscounts: comboDiscounts,
	}
}

func (m *inMemoryComboDiscountRepository) GetDiscounts(watchIDs []models.WatchID) ([]models.Discount, error) {
	var result []models.Discount
	for _, watchID := range watchIDs {
		if discount, ok := m.comboDiscounts.Load(watchID); ok {
			comboDiscount := discount.(models.ComboDiscount)
			result = append(result, &comboDiscount)
		}
	}
	return result, nil
}
