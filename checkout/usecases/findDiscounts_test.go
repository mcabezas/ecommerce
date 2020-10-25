package usecases

import (
	"testing"

	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/stretchr/testify/assert"
)

func TestGivenAEmptyListCanFindAUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 3, 200),
		models.NewComboDiscount("002", 3, 120),
	})
	discounts, err := finder.findComboDiscounts([]models.WatchID{})
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 0)
}

func TestCanFindSingleUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 3, 200),
		models.NewComboDiscount("002", 3, 120),
		models.NewComboDiscount("003", 0, 120),
	})
	discounts, err := finder.findComboDiscounts(
		[]models.WatchID{"002"},
	)
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 1)
}

func TestCanFindMultipleUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 3, 200),
		models.NewComboDiscount("002", 3, 120),
		models.NewComboDiscount("003", 0, 120),
	})
	discounts, err := finder.findComboDiscounts(
		[]models.WatchID{"002", "003"},
	)
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 2)
}

func buildDiscountFinder(comboDiscounts []models.ComboDiscount) DiscountFinder {
	return NewDiscountFinder(
		infrastructure.NewInMemoryComboDiscountRepository(comboDiscounts),
	)
}
