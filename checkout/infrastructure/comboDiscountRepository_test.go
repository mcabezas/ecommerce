package infrastructure

import (
	"testing"

	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/stretchr/testify/assert"
)

func TestCanGetDiscounts(t *testing.T) {
	repository := NewInMemoryComboDiscountRepository([]models.ComboDiscount{
		models.NewComboDiscount("100", 2, 200),
		models.NewComboDiscount("10", 2, 20),
		models.NewComboDiscount("1", 2, 2),
	})
	discounts, err := repository.GetDiscounts([]models.WatchID{"1"})
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 1)
}

func TestCanGetMorThanOneDiscounts(t *testing.T) {
	repository := NewInMemoryComboDiscountRepository([]models.ComboDiscount{
		models.NewComboDiscount("100", 2, 200),
		models.NewComboDiscount("10", 2, 20),
		models.NewComboDiscount("1", 2, 2),
	})
	discounts, err := repository.GetDiscounts([]models.WatchID{"wrongID", "1", "10"})
	assert.Nil(t, err)
	assert.True(t, len(discounts)==2)
}